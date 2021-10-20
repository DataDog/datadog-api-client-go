module github.com/DataDog/datadog-api-client-go/tests

go 1.14

require (
	github.com/DataDog/datadog-api-client-go v1.0.0-beta.22
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/DataDog/dd-sdk-go-testing v0.0.0-20210929140144-5d69f0a9bd49
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/dnaeon/go-vcr v1.0.1
	github.com/go-bdd/gobdd v1.1.3
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.4.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jonboulle/clockwork v0.1.0
	github.com/mcuadros/go-lookup v0.0.0-20200831155250-80f87a4fa5ee
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/tinylib/msgp v1.1.6 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4
	golang.org/x/sys v0.0.0-20211020174200-9d6173849985 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.33.0
	gopkg.in/h2non/gock.v1 v1.0.15
	gotest.tools v2.2.0+incompatible
)

replace gopkg.in/DataDog/dd-trace-go.v1 => github.com/DataDog/dd-trace-go v1.30.0-rc.1.0.20210420124628-f63633f38e8f

replace github.com/mcuadros/go-lookup => github.com/therve/go-lookup v0.0.0-20210911125823-5508a75dec6f

replace github.com/DataDog/datadog-api-client-go => ../
