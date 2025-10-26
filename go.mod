module github.com/DataDog/datadog-api-client-go/v2

go 1.23.0

retract (
	// Version used to retract v2.0.0 and v2.0.1. DO NOT USE.
	v2.0.1
	// Premature major version v2 release. DO NOT USE.
	v2.0.0
)

require (
	github.com/DataDog/zstd v1.5.2
	github.com/goccy/go-json v0.10.2
	github.com/google/uuid v1.5.0
	golang.org/x/oauth2 v0.27.0
)
