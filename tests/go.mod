module github.com/DataDog/datadog-api-client-go/tests

go 1.23.0

toolchain go1.23.7

require (
	github.com/DataDog/dd-sdk-go-testing v0.0.3
	github.com/cucumber/messages-go/v12 v12.0.0
	github.com/go-bdd/gobdd v1.1.4-0.20211209204431-ca566a78d075
	github.com/jonboulle/clockwork v0.1.0
	github.com/stretchr/testify v1.7.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.33.0
	gopkg.in/dnaeon/go-vcr.v3 v3.1.2
	gotest.tools v2.2.0+incompatible
)

require golang.org/x/net v0.36.0 // indirect

require (
	github.com/DataDog/datadog-api-client-go/packages/account-management v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/actions v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/ai-guard v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/ai-observability v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/apm v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/app-api-protection v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/app-builder v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/audit-trail v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/bits-ai v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/ci-visibility v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/cloud-cost v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/cloud-integrations v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/cloud-network-monitoring v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/code-coverage v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/code-security v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/containers v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/core v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/dashboards v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/data-observability v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/ddsql v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/disaster-recovery v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/end-user-device-monitoring v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/error-tracking v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/events v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/experiments v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/feature-flags v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/fleet-automation v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/governance v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/incident-management v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/incident-response v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/infrastructure v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/integrations v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/logs v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/metrics v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/monitors v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/multi-org v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/network-device-monitoring v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/notebooks v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/observability-pipelines v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/on-call v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/product-analytics v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/reference-tables v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/rum v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/security-detection v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/security-findings v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/sensitive-data-scanner v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/service-level-objectives v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/session-replay v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/software-catalog v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/software-delivery v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/storage-management v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/synthetics v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/teams v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/test-optimization v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/usage-metering v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/widgets v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/workflow-automation v0.0.1
	github.com/DataDog/datadog-api-client-go/packages/workload-protection v0.0.1
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

replace github.com/DataDog/datadog-api-client-go/packages/account-management => ../packages/account-management

replace github.com/DataDog/datadog-api-client-go/packages/actions => ../packages/actions

replace github.com/DataDog/datadog-api-client-go/packages/ai-guard => ../packages/ai-guard

replace github.com/DataDog/datadog-api-client-go/packages/ai-observability => ../packages/ai-observability

replace github.com/DataDog/datadog-api-client-go/packages/apm => ../packages/apm

replace github.com/DataDog/datadog-api-client-go/packages/app-api-protection => ../packages/app-api-protection

replace github.com/DataDog/datadog-api-client-go/packages/app-builder => ../packages/app-builder

replace github.com/DataDog/datadog-api-client-go/packages/audit-trail => ../packages/audit-trail

replace github.com/DataDog/datadog-api-client-go/packages/bits-ai => ../packages/bits-ai

replace github.com/DataDog/datadog-api-client-go/packages/ci-visibility => ../packages/ci-visibility

replace github.com/DataDog/datadog-api-client-go/packages/cloud-cost => ../packages/cloud-cost

replace github.com/DataDog/datadog-api-client-go/packages/cloud-integrations => ../packages/cloud-integrations

replace github.com/DataDog/datadog-api-client-go/packages/cloud-network-monitoring => ../packages/cloud-network-monitoring

replace github.com/DataDog/datadog-api-client-go/packages/code-coverage => ../packages/code-coverage

replace github.com/DataDog/datadog-api-client-go/packages/code-security => ../packages/code-security

replace github.com/DataDog/datadog-api-client-go/packages/containers => ../packages/containers

replace github.com/DataDog/datadog-api-client-go/packages/core => ../packages/core

replace github.com/DataDog/datadog-api-client-go/packages/dashboards => ../packages/dashboards

replace github.com/DataDog/datadog-api-client-go/packages/data-observability => ../packages/data-observability

replace github.com/DataDog/datadog-api-client-go/packages/ddsql => ../packages/ddsql

replace github.com/DataDog/datadog-api-client-go/packages/disaster-recovery => ../packages/disaster-recovery

replace github.com/DataDog/datadog-api-client-go/packages/end-user-device-monitoring => ../packages/end-user-device-monitoring

replace github.com/DataDog/datadog-api-client-go/packages/error-tracking => ../packages/error-tracking

replace github.com/DataDog/datadog-api-client-go/packages/events => ../packages/events

replace github.com/DataDog/datadog-api-client-go/packages/experiments => ../packages/experiments

replace github.com/DataDog/datadog-api-client-go/packages/feature-flags => ../packages/feature-flags

replace github.com/DataDog/datadog-api-client-go/packages/fleet-automation => ../packages/fleet-automation

replace github.com/DataDog/datadog-api-client-go/packages/governance => ../packages/governance

replace github.com/DataDog/datadog-api-client-go/packages/incident-management => ../packages/incident-management

replace github.com/DataDog/datadog-api-client-go/packages/incident-response => ../packages/incident-response

replace github.com/DataDog/datadog-api-client-go/packages/infrastructure => ../packages/infrastructure

replace github.com/DataDog/datadog-api-client-go/packages/integrations => ../packages/integrations

replace github.com/DataDog/datadog-api-client-go/packages/logs => ../packages/logs

replace github.com/DataDog/datadog-api-client-go/packages/metrics => ../packages/metrics

replace github.com/DataDog/datadog-api-client-go/packages/monitors => ../packages/monitors

replace github.com/DataDog/datadog-api-client-go/packages/multi-org => ../packages/multi-org

replace github.com/DataDog/datadog-api-client-go/packages/network-device-monitoring => ../packages/network-device-monitoring

replace github.com/DataDog/datadog-api-client-go/packages/notebooks => ../packages/notebooks

replace github.com/DataDog/datadog-api-client-go/packages/observability-pipelines => ../packages/observability-pipelines

replace github.com/DataDog/datadog-api-client-go/packages/on-call => ../packages/on-call

replace github.com/DataDog/datadog-api-client-go/packages/product-analytics => ../packages/product-analytics

replace github.com/DataDog/datadog-api-client-go/packages/reference-tables => ../packages/reference-tables

replace github.com/DataDog/datadog-api-client-go/packages/rum => ../packages/rum

replace github.com/DataDog/datadog-api-client-go/packages/security-detection => ../packages/security-detection

replace github.com/DataDog/datadog-api-client-go/packages/security-findings => ../packages/security-findings

replace github.com/DataDog/datadog-api-client-go/packages/sensitive-data-scanner => ../packages/sensitive-data-scanner

replace github.com/DataDog/datadog-api-client-go/packages/service-level-objectives => ../packages/service-level-objectives

replace github.com/DataDog/datadog-api-client-go/packages/session-replay => ../packages/session-replay

replace github.com/DataDog/datadog-api-client-go/packages/software-catalog => ../packages/software-catalog

replace github.com/DataDog/datadog-api-client-go/packages/software-delivery => ../packages/software-delivery

replace github.com/DataDog/datadog-api-client-go/packages/storage-management => ../packages/storage-management

replace github.com/DataDog/datadog-api-client-go/packages/synthetics => ../packages/synthetics

replace github.com/DataDog/datadog-api-client-go/packages/teams => ../packages/teams

replace github.com/DataDog/datadog-api-client-go/packages/test-optimization => ../packages/test-optimization

replace github.com/DataDog/datadog-api-client-go/packages/usage-metering => ../packages/usage-metering

replace github.com/DataDog/datadog-api-client-go/packages/widgets => ../packages/widgets

replace github.com/DataDog/datadog-api-client-go/packages/workflow-automation => ../packages/workflow-automation

replace github.com/DataDog/datadog-api-client-go/packages/workload-protection => ../packages/workload-protection
