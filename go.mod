module github.com/DataDog/datadog-api-client-go

go 1.14

require (
	github.com/DataDog/datadog-go v3.6.0+incompatible // indirect
	github.com/dnaeon/go-vcr v1.0.1
	github.com/go-bdd/gobdd v1.1.3-0.20210205100305-4910f932a786
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/jonboulle/clockwork v0.1.0
	github.com/mcuadros/go-lookup v0.0.0-20200831155250-80f87a4fa5ee
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/philhofer/fwd v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.29.0-alpha.1.0.20210212162759-6903fc5e42cd
	gopkg.in/h2non/gock.v1 v1.0.15
	gotest.tools v2.2.0+incompatible
)

replace gopkg.in/DataDog/dd-trace-go.v1 => github.com/DataDog/dd-trace-go v1.29.0-alpha.1.0.20210212162759-6903fc5e42cd
