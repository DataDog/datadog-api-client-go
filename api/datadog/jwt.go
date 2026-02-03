// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadog

import (
	"time"

	"github.com/google/uuid"
)

// JWTHeader represents the header portion of a JWT token.
type JWTHeader struct {
	// Algorithm is the signing algorithm used for the JWT
	Algorithm SigningAlgorithm `json:"alg"`
	// KeyUUID is the key identifier
	KeyUUID uuid.UUID `json:"kid"`
	// Type is the token type (typically "JWT")
	Type string `json:"typ"`
}

// JWTBody represents the body (claims) portion of a JWT token.
type JWTBody struct {
	// Audience identifies the recipients that the JWT is intended for
	Audience string `json:"aud"`
	// IssuedAtTimestamp is the time at which the JWT was issued
	IssuedAtTimestamp time.Time `json:"iat"`
	// Iss is the issuer of the JWT
	Iss string `json:"iss"`
	// ExpirationTimestamp is the expiration time on or after which the JWT must not be accepted
	ExpirationTimestamp time.Time `json:"exp"`
	// Subject identifies the principal that is the subject of the JWT
	Subject uuid.UUID `json:"sub"`
}

// JWT represents a complete JSON Web Token structure.
type JWT struct {
	// Header contains the JWT header information
	Header JWTHeader `json:"header"`
	// Body contains the JWT claims
	Body JWTBody `json:"body"`
}
