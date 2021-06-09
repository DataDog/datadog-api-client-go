module github.com/DataDog/datadog-api-client-go/tests

go 1.14

require (
	github.com/DataDog/datadog-api-client-go v1.0.0-beta.22
	github.com/dnaeon/go-vcr v1.0.1
	github.com/go-bdd/gobdd v1.1.3
	github.com/jonboulle/clockwork v0.1.0
	github.com/mcuadros/go-lookup v0.0.0-20200831155250-80f87a4fa5ee
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	gopkg.in/DataDog/dd-trace-go.v1 v1.30.0-rc.1.0.20210420124628-f63633f38e8f
	gopkg.in/h2non/gock.v1 v1.0.15
	gotest.tools v2.2.0+incompatible
)

replace gopkg.in/DataDog/dd-trace-go.v1 => github.com/DataDog/dd-trace-go v1.30.0-rc.1.0.20210420124628-f63633f38e8f

replace github.com/mcuadros/go-lookup => github.com/zippolyte/go-lookup v0.0.0-20210331103309-2b027b989715

replace github.com/DataDog/datadog-api-client-go => ../
