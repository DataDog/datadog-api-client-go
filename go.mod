module github.com/DataDog/datadog-api-client-go/v2

go 1.17

retract (
	// Version used to retract v2.0.0 and v2.0.1. DO NOT USE.
	v2.0.1
	// Premature major version v2 release. DO NOT USE.
	v2.0.0
)

require (
	github.com/DataDog/zstd v1.5.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
)

require (
	github.com/golang/protobuf v1.2.0 // indirect
	golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	google.golang.org/appengine v1.4.0 // indirect
)
