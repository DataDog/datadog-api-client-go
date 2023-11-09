package tests

import (
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func BenchmarkDatadogModels(b *testing.B) {
	b.ReportAllocs()
	jb, err := os.ReadFile("/tmp/1.json")
	if err != nil {
		b.Fatal(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var db datadogV1.Dashboard
		err := datadog.Unmarshal(jb, &db)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkJSONParsing(b *testing.B) {
	b.ReportAllocs()
	jb, err := os.ReadFile("/tmp/1.json")
	if err != nil {
		b.Fatal(err.Error())
	}
	for i := 0; i < b.N; i++ {
		var db map[string]any
		err := datadog.Unmarshal(jb, &db)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
