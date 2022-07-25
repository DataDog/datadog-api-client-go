module github.com/DataDog/datadog-api-client-go/v2/tests

go 1.14

require (
	github.com/DataDog/datadog-api-client-go/v2 v2.0.0-20220725190118-a72751af03fe
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/DataDog/dd-sdk-go-testing v0.0.0-20210929140144-5d69f0a9bd49
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/cucumber/messages-go/v12 v12.0.0
	github.com/dnaeon/go-vcr v1.0.1
	github.com/go-bdd/gobdd v1.1.4-0.20211209204431-ca566a78d075
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jonboulle/clockwork v0.1.0
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

replace github.com/DataDog/datadog-api-client-go/v2 => ../
