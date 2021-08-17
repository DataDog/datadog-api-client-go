package api

import (
	"context"
	"testing"

	datadogV1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	datadogV2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
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
			datadogV1.Dashboard{UnparsedObject: map[string]interface{}{"foo": "bar"}},
			true,
		},
		{
			"nested unparsed struct",
			datadogV1.SyntheticsAPITest{Name: datadogV1.PtrString("foo"), Config: &datadogV1.SyntheticsAPITestConfig{UnparsedObject: map[string]interface{}{"foo": "bar"}}},
			true,
		},
		{
			"unparsed struct in array",
			datadogV1.Dashboard{Widgets: []datadogV1.Widget{{Definition: datadogV1.WidgetDefinition{}}, {UnparsedObject: map[string]interface{}{"foo": "bar"}}}},
			true,
		},
		{
			"unparsed enum in array",
			datadogV1.SyntheticsTestOptions{DeviceIds: &[]datadogV1.SyntheticsDeviceID{"edge", "foobar"}},
			true,
		},
		{
			"unparsed enum in map",
			map[string]datadogV1.SyntheticsDeviceID{"foo": "edge", "bar": "foobar"},
			true,
		},
		{
			"unparsed nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{UnparsedObject: map[string]interface{}{"foo": "bar"}}),
			true,
		},
		{
			"unparsed nested in nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{UnparsedObject: map[string]interface{}{"foo": "bar"}}}),
			true,
		},
		{
			"valid nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{Type: datadogV2.LOGSARCHIVEDESTINATIONAZURETYPE_AZURE}}),
			false,
		},
		{
			"valid struct",
			datadogV1.SyntheticsAPITest{Name: datadogV1.PtrString("foo"), Config: &datadogV1.SyntheticsAPITestConfig{Assertions: &[]datadogV1.SyntheticsAssertion{{SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{Type: datadogV1.SYNTHETICSASSERTIONTYPE_BODY}}}}},
			false,
		},
		{
			"valid simple type",
			"a simple string",
			false,
		},
		{
			"valid simple pointer",
			datadogV1.PtrString("a simple pointer to string"),
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(tc.expected, datadogV1.ContainsUnparsedObject(tc.value))
		})
	}
}
