// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	CustomIDPIssuer         = "datadog-custom-idp"
	customIDPTokenExpiration = 5 * time.Minute
)

// CustomIDPAuth implements delegated authentication using custom IDP with JWT.
type CustomIDPAuth struct {
	// PrivateKey is the PEM-encoded private key used to sign the JWT
	PrivateKey []byte
	// Algorithm is the signing algorithm to use (RS256 or ES256)
	Algorithm SigningAlgorithm
	// KeyUUID is the key identifier
	KeyUUID uuid.UUID
	// Audience is the JWT audience
	Audience string
}

// Authenticate implements the DelegatedTokenProvider interface.
// It generates a signed JWT and uses it to obtain a delegated token.
func (c *CustomIDPAuth) Authenticate(ctx context.Context, config *DelegatedTokenConfig) (*DelegatedTokenCredentials, error) {
	if config == nil || config.OrgUUID == "" {
		return nil, fmt.Errorf("missing org UUID in config")
	}

	if err := c.validate(); err != nil {
		return nil, err
	}

	orgUUID, err := uuid.Parse(config.OrgUUID)
	if err != nil {
		return nil, fmt.Errorf("invalid org UUID: %w", err)
	}

	jwtToken, err := c.GenerateJWT(ctx, orgUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate JWT: %w", err)
	}

	signedJWT, err := c.SignJWT(jwtToken)
	if err != nil {
		return nil, fmt.Errorf("failed to sign JWT: %w", err)
	}

	authResponse, err := GetDelegatedToken(ctx, config.OrgUUID, signedJWT)
	return authResponse, err
}

func (c *CustomIDPAuth) validate() error {
	if len(c.PrivateKey) == 0 {
		return fmt.Errorf("missing private key")
	}
	if c.Algorithm != SigningAlgorithmRS256 && c.Algorithm != SigningAlgorithmES256 {
		return fmt.Errorf("invalid algorithm: must be RS256 or ES256")
	}
	if c.KeyUUID == uuid.Nil {
		return fmt.Errorf("missing key UUID")
	}
	if c.Audience == "" {
		return fmt.Errorf("missing audience")
	}
	return nil
}

// GenerateJWT creates a JWT structure with the current timestamp and configured values.
// The orgUUID parameter is used as the JWT subject.
func (c *CustomIDPAuth) GenerateJWT(ctx context.Context, orgUUID uuid.UUID) (*JWT, error) {
	now := time.Now().UTC()

	jwt := &JWT{
		Header: JWTHeader{
			Algorithm: c.Algorithm,
			KeyUUID:   c.KeyUUID,
			Type:      "JWT",
		},
		Body: JWTBody{
			Audience:            c.Audience,
			IssuedAtTimestamp:   now,
			Iss:                 CustomIDPIssuer,
			ExpirationTimestamp: now.Add(customIDPTokenExpiration),
			Subject:             orgUUID,
		},
	}

	return jwt, nil
}

// SignJWT signs a JWT and returns the complete JWT token string in the format: header.body.signature
func (c *CustomIDPAuth) SignJWT(jwt *JWT) (string, error) {
	headerJSON, err := json.Marshal(jwt.Header)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JWT header: %w", err)
	}
	headerEncoded := base64.RawURLEncoding.EncodeToString(headerJSON)

	bodyMap := map[string]interface{}{
		"aud": jwt.Body.Audience,
		"iat": jwt.Body.IssuedAtTimestamp.Unix(),
		"iss": jwt.Body.Iss,
		"exp": jwt.Body.ExpirationTimestamp.Unix(),
		"sub": jwt.Body.Subject.String(),
	}
	bodyJSON, err := json.Marshal(bodyMap)
	if err != nil {
		return "", fmt.Errorf("failed to marshal JWT body: %w", err)
	}
	bodyEncoded := base64.RawURLEncoding.EncodeToString(bodyJSON)

	signingInput := headerEncoded + "." + bodyEncoded

	signature, err := SignString(signingInput, c.PrivateKey, c.Algorithm)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return strings.Join([]string{headerEncoded, bodyEncoded, signature}, "."), nil
}
