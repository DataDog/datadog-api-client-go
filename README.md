# datadog-api-client-go

This repository contains a Go API client for [Datadog API](https://docs.datadoghq.com/api/).  
The code is generated using [openapi-generator](https://github.com/OpenAPITools/openapi-generator)
and [apigentools](https://github.com/DataDog/apigentools).

## Layout

It is the intention of this repository to contain per-major-version API client packages. Right
now, Datadog only has one API version - `v1` - so that is the only module present. If/when
there is another version of the API, we will add a different package for it. The end goal is
to be able to use API clients for different major versions of the API side-by-side.

## The Datadog API v1 Client

The client library for Datadog API v1 is located in the `datadog_v1` directory. Import it with

```go
import github.com/datadog/datadog-api-client-go/datadog_v1
```

All the documentation for this package is available through `datadog_v1/README.md`.

## Contributing

As most of the code in this repository is generated, we will only accept PRs for files
that are not modified by our code-generation machinery (changes to the generated files
would get overwritten). A list of files that we will happily accept contributions for
follows:

* `datadog_v1/.gitignore`
* `datadog_v1/*_test.go`
