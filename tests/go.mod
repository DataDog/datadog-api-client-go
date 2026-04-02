module github.com/DataDog/datadog-api-client-go/v2/tests

go 1.23.0

toolchain go1.23.7

require (
	github.com/DataDog/datadog-api-client-go/v2 v2.14.0
	github.com/DataDog/dd-sdk-go-testing v0.0.3
	github.com/cucumber/messages-go/v12 v12.0.0
	github.com/go-bdd/gobdd v1.1.4-0.20211209204431-ca566a78d075
	github.com/golang/mock v1.6.0
	github.com/jonboulle/clockwork v0.1.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.36.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.33.0
	gopkg.in/dnaeon/go-vcr.v3 v3.1.2
	gopkg.in/h2non/gock.v1 v1.0.15
	gotest.tools v2.2.0+incompatible
)

require (
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/DataDog/sketches-go v1.0.0 // indirect
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/cucumber/gherkin-go/v13 v13.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/h2non/parth v0.0.0-20190131123155-b4df798d6542 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	golang.org/x/oauth2 v0.10.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/DataDog/datadog-api-client-go/v2 => ../
