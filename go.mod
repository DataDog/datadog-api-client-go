module github.com/DataDog/datadog-api-client-go

go 1.14

require (
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/dnaeon/go-vcr v1.0.1
	github.com/go-bdd/gobdd v1.1.3
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/jonboulle/clockwork v0.1.0
	github.com/mcuadros/go-lookup v0.0.0-20200831155250-80f87a4fa5ee
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20190108225652-1e06a53dbb7e
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20200420163511-1957bb5e6d1f // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.30.0-rc.1.0.20210420124628-f63633f38e8f
	gopkg.in/h2non/gock.v1 v1.0.15
	gotest.tools v2.2.0+incompatible
)

replace gopkg.in/DataDog/dd-trace-go.v1 => github.com/DataDog/dd-trace-go v1.30.0-rc.1.0.20210420124628-f63633f38e8f

replace github.com/mcuadros/go-lookup => github.com/zippolyte/go-lookup v0.0.0-20210331103309-2b027b989715
