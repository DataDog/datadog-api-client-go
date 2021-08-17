package api

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestContainsUnparsedObject(t *testing.T) {
	assert := tests.Assert(context.Background(), t)
	testCases := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{
			"top level unparsed struct",
			datadog.Dashboard{UnparsedObject: map[string]interface{}{"foo": "bar"}},
			true,
		},
		{
			"nested unparsed struct",
			datadog.SyntheticsAPITest{Name: datadog.PtrString("foo"), Config: &datadog.SyntheticsAPITestConfig{UnparsedObject: map[string]interface{}{"foo": "bar"}}},
			true,
		},
		{
			"unparsed struct in array",
			datadog.Dashboard{Widgets: []datadog.Widget{{Definition: datadog.WidgetDefinition{}}, {UnparsedObject: map[string]interface{}{"foo": "bar"}}}},
			true,
		},
		{
			"unparsed enum in array",
			datadog.SyntheticsTestOptions{DeviceIds: &[]datadog.SyntheticsDeviceID{"edge", "foobar"}},
			true,
		},
		{
			"unparsed enum in map",
			map[string]datadog.SyntheticsDeviceID{"foo": "edge", "bar": "foobar"},
			true,
		},
		{
			"valid struct",
			datadog.SyntheticsAPITest{Name: datadog.PtrString("foo"), Config: &datadog.SyntheticsAPITestConfig{Assertions: &[]datadog.SyntheticsAssertion{{SyntheticsAssertionTarget: &datadog.SyntheticsAssertionTarget{Type: datadog.SYNTHETICSASSERTIONTYPE_BODY}}}}},
			false,
		},
		{
			"valid simple type",
			"a simple string",
			false,
		},
		{
			"valid simple pointer",
			datadog.PtrString("a simple pointer to string"),
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(tc.expected, datadog.ContainsUnparsedObject(tc.value))
		})
	}
}
