// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package awsauth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

const (
	defaultRegion     = "us-east-1"
	getCallerIdentity = "Action=GetCallerIdentity&Version=2011-06-15"
	orgIDHeader       = "x-ddog-org-id"
	stsService        = "sts"
	formContentType   = "application/x-www-form-urlencoded; charset=utf-8"
)

// Option configures a Provider.
type Option func(*providerOptions) error

type providerOptions struct {
	loadOptions []func(*awsconfig.LoadOptions) error
	httpClient  *http.Client
}

// WithConfigOptions adds options passed to awsconfig.LoadDefaultConfig.
// This is an escape hatch for AWS SDK settings not covered by the convenience
// options in this package.
func WithConfigOptions(options ...func(*awsconfig.LoadOptions) error) Option {
	return func(config *providerOptions) error {
		config.loadOptions = append(config.loadOptions, options...)
		return nil
	}
}

// WithRegion sets the AWS region used for STS endpoint resolution and request
// signing. When omitted, the AWS SDK configuration is used, falling back to
// us-east-1 when the credential source does not specify a region.
func WithRegion(region string) Option {
	return func(config *providerOptions) error {
		if region != "" {
			config.loadOptions = append(config.loadOptions, awsconfig.WithRegion(region))
		}
		return nil
	}
}

// WithStaticCredentials uses explicitly supplied AWS credentials instead of
// the default credential chain. A session token is optional.
func WithStaticCredentials(accessKeyID, secretAccessKey, sessionToken string) Option {
	return func(config *providerOptions) error {
		if accessKeyID == "" || secretAccessKey == "" {
			return errors.New("AWS access key ID and secret access key must both be set")
		}
		provider := credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, sessionToken)
		config.loadOptions = append(config.loadOptions, awsconfig.WithCredentialsProvider(provider))
		return nil
	}
}

// WithHTTPClient sets the client used for AWS credential retrieval and the
// Datadog delegated-token exchange. This allows callers to preserve custom
// proxy, TLS, and timeout configuration across the entire authentication flow.
func WithHTTPClient(client *http.Client) Option {
	return func(config *providerOptions) error {
		if client == nil {
			return errors.New("HTTP client must not be nil")
		}
		config.httpClient = client
		config.loadOptions = append(config.loadOptions, awsconfig.WithHTTPClient(client))
		return nil
	}
}

// Provider implements datadog.DelegatedTokenProvider using the AWS SDK default
// credential chain. Configuration and credential retrieval are lazy so clients
// that disable validation do not contact AWS during construction.
type Provider struct {
	loadOptions []func(*awsconfig.LoadOptions) error
	httpClient  *http.Client

	configMu  sync.Mutex
	awsConfig *aws.Config
}

var _ datadog.DelegatedTokenProvider = (*Provider)(nil)

// New creates an AWS delegated-token provider.
func New(options ...Option) (*Provider, error) {
	config := providerOptions{
		httpClient: http.DefaultClient,
	}
	for _, option := range options {
		if option == nil {
			continue
		}
		if err := option(&config); err != nil {
			return nil, err
		}
	}

	return &Provider{
		loadOptions: config.loadOptions,
		httpClient:  config.httpClient,
	}, nil
}

// Authenticate exchanges an AWS-signed GetCallerIdentity proof for a Datadog
// delegated token.
func (provider *Provider) Authenticate(ctx context.Context, config *datadog.DelegatedTokenConfig) (*datadog.DelegatedTokenCredentials, error) {
	if config == nil || config.OrgUUID == "" {
		return nil, errors.New("missing org UUID in delegated token configuration")
	}

	awsConfig, err := provider.loadAWSConfig(ctx)
	if err != nil {
		return nil, err
	}
	credentials, err := awsConfig.Credentials.Retrieve(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving AWS credentials: %w", err)
	}
	log.Printf("[INFO] Datadog AWS delegated auth using credential source %q in region %q", credentials.Source, awsConfig.Region)

	proof, err := provider.generateProof(ctx, config.OrgUUID, *awsConfig, credentials)
	if err != nil {
		return nil, err
	}
	tokenURL, err := datadog.GetDelegatedTokenUrl(ctx)
	if err != nil {
		return nil, fmt.Errorf("resolving Datadog delegated token URL: %w", err)
	}
	token, err := provider.exchangeDelegatedToken(ctx, tokenURL, config.OrgUUID, proof)
	if err != nil {
		return nil, fmt.Errorf("exchanging AWS identity proof for Datadog delegated token: %w", err)
	}
	if token == nil {
		return nil, errors.New("Datadog delegated token endpoint returned no credentials")
	}

	return token, nil
}

func (provider *Provider) exchangeDelegatedToken(ctx context.Context, tokenURL, orgUUID, proof string) (*datadog.DelegatedTokenCredentials, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("creating Datadog delegated token request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Delegated "+proof)

	response, err := provider.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("requesting Datadog delegated token: %w", err)
	}
	defer response.Body.Close()

	const maximumResponseSize = 1 << 20
	body, err := io.ReadAll(io.LimitReader(response.Body, maximumResponseSize+1))
	if err != nil {
		return nil, fmt.Errorf("reading Datadog delegated token response: %w", err)
	}
	if len(body) > maximumResponseSize {
		return nil, errors.New("Datadog delegated token response exceeds 1 MiB")
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Datadog delegated token endpoint returned %s", response.Status)
	}

	credentials, err := datadog.ParseDelegatedTokenResponse(body, orgUUID, proof)
	if err != nil {
		return nil, fmt.Errorf("parsing Datadog delegated token response: %w", err)
	}
	return credentials, nil
}

func (provider *Provider) loadAWSConfig(ctx context.Context) (*aws.Config, error) {
	provider.configMu.Lock()
	defer provider.configMu.Unlock()

	if provider.awsConfig != nil {
		return provider.awsConfig, nil
	}

	config, err := awsconfig.LoadDefaultConfig(ctx, provider.loadOptions...)
	if err != nil {
		return nil, fmt.Errorf("loading AWS default configuration: %w", err)
	}
	if config.Credentials == nil {
		return nil, errors.New("AWS default configuration returned no credentials provider")
	}
	if config.Region == "" {
		config.Region = defaultRegion
	}

	provider.awsConfig = &config
	return provider.awsConfig, nil
}

func (provider *Provider) generateProof(ctx context.Context, orgUUID string, config aws.Config, credentials aws.Credentials) (string, error) {
	stsOptions := sts.NewFromConfig(config).Options()
	endpoint, err := stsOptions.EndpointResolverV2.ResolveEndpoint(ctx, sts.EndpointParameters{
		Region:       aws.String(stsOptions.Region),
		UseDualStack: aws.Bool(stsOptions.EndpointOptions.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled),
		UseFIPS:      aws.Bool(stsOptions.EndpointOptions.UseFIPSEndpoint == aws.FIPSEndpointStateEnabled),
		Endpoint:     stsOptions.BaseEndpoint,
	})
	if err != nil {
		return "", fmt.Errorf("resolving AWS STS endpoint: %w", err)
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.URI.String(), strings.NewReader(getCallerIdentity))
	if err != nil {
		return "", fmt.Errorf("creating AWS STS identity request: %w", err)
	}
	request.Header.Set("Content-Type", formContentType)
	request.Header.Set("Content-Length", strconv.Itoa(len(getCallerIdentity)))
	request.Header.Set(orgIDHeader, orgUUID)

	payloadHash := sha256.Sum256([]byte(getCallerIdentity))
	if err := v4.NewSigner().SignHTTP(
		ctx,
		credentials,
		request,
		hex.EncodeToString(payloadHash[:]),
		stsService,
		stsOptions.Region,
		time.Now(),
	); err != nil {
		return "", fmt.Errorf("signing AWS STS identity request: %w", err)
	}

	headers := request.Header.Clone()
	headers.Set("host", request.URL.Host)
	headers.Set("User-Agent", datadog.GetUserAgent())
	headersJSON, err := json.Marshal(headers)
	if err != nil {
		return "", fmt.Errorf("encoding AWS identity proof headers: %w", err)
	}

	parts := []string{
		base64.StdEncoding.EncodeToString([]byte(getCallerIdentity)),
		base64.StdEncoding.EncodeToString(headersJSON),
		http.MethodPost,
		base64.StdEncoding.EncodeToString([]byte(request.URL.String())),
	}
	return strings.Join(parts, "|"), nil
}
