package delegated_auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	tokenUrlEndpoint    = "https://%s.%s/api/v2/delegated-token"
	authorizationHeader = "Authorization"
	authorizationType   = "Delegated"

	contentTypeHeader   = "Content-Type"
	contextLengthHeader = "Content-Length"
	applicationJson     = "application/json"
)

func getDelegatedToken(ctx context.Context, orgUUID, delegatedAuthProof string) (*datadog.DelegatedTokenCredentials, error) {
	// Get site, api subdomain, and orgUUID from the context
	serverVars := getServerVariables(ctx)
	site, err := getSite(serverVars)
	if err != nil {
		return nil, err
	}
	subdomain := getApiSubdomain(serverVars)

	// Call the token endpoint to get a temporary token
	url := fmt.Sprintf(tokenUrlEndpoint, subdomain, site)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}
	req.Header.Set(contentTypeHeader, applicationJson)
	req.Header.Set(authorizationHeader, fmt.Sprintf("%s %s", authorizationType, delegatedAuthProof))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	tokenBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get token: %s", resp.Status)
	}

	// Parse the response to get the token
	var tokenResponse map[string]interface{}
	err = json.Unmarshal(tokenBytes, &tokenResponse)
	if err != nil {
		return nil, err
	}

	// Get attributes from the response
	dataResponse, ok := tokenResponse["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to get data from response: %v", tokenResponse)
	}

	// Get the attributes from the data
	attributes, ok := dataResponse["attributes"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to get attributes from response: %v", tokenResponse)
	}

	// Get the access token from the response
	token, ok := attributes["access_token"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get token from response: %v", tokenResponse)
	}

	// Get the expiration time from the response
	// Default to 15 minutes if the expiration time is not set
	expirationTime := time.Now().Add(time.Duration(15) * time.Minute)
	expirationStr, ok := attributes["expires"].(string)
	if ok {
		expirationInt, err := strconv.ParseInt(expirationStr, 10, 64)
		if err == nil {
			// Convert the expiration time to a time.Time object
			expirationTime = time.Unix(expirationInt, 0)
		}
	}

	return &datadog.DelegatedTokenCredentials{
		OrgUUID:        orgUUID,
		DelegatedToken: token,
		DelegatedProof: delegatedAuthProof,
		Expiration:     expirationTime,
	}, nil
}

func getServerVariables(ctx context.Context) map[string]string {
	vars := ctx.Value(datadog.ContextServerVariables)
	if vars == nil {
		return map[string]string{}
	}
	return vars.(map[string]string)
}

func getSite(serverVars map[string]string) (string, error) {
	site := serverVars["site"]
	if site == "" {
		name := serverVars["name"]
		if name != "" {
			name = strings.TrimPrefix(name, "api.")
			return name, nil
		}
		return "", datadog.ReportError("site not found in server variables")
	}
	return site, nil
}

func getApiSubdomain(serverVars map[string]string) string {
	apiName := serverVars["subdomain"]
	if apiName == "" {
		return "api"
	}
	return apiName
}
