package api

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

func TestContainsUnparsedObject(t *testing.T) {
	assert := tests.Assert(context.Background(), t)
	testCases := []struct {
		name                   string
		value                  interface{}
		expectedBool           bool
		expectedUnparsedObject interface{}
	}{
		{
			"top level unparsed struct",
			datadogV1.Dashboard{UnparsedObject: map[string]interface{}{"foo": "bar"}},
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"nested unparsed struct",
			datadogV1.SyntheticsAPITest{Name: "foo", Config: datadogV1.SyntheticsAPITestConfig{UnparsedObject: map[string]interface{}{"foo": "bar"}}},
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"unparsed struct in array",
			datadogV1.Dashboard{LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_FREE, Widgets: []datadogV1.Widget{{Definition: datadogV1.WidgetDefinition{}}, {UnparsedObject: map[string]interface{}{"foo": "bar"}}}},
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"unparsed enum in array",
			datadogV2.DowntimeResponseAttributes{NotifyEndStates: []datadogV2.DowntimeNotifyEndStateTypes{"alert", "foobar"}},
			true,
			datadogV2.DowntimeNotifyEndStateTypes("foobar"),
		},
		{
			"unparsed enum in map",
			map[string]datadogV2.DowntimeNotifyEndStateTypes{"foo": "alert", "bar": "foobar"},
			true,
			datadogV2.DowntimeNotifyEndStateTypes("foobar"),
		},
		{
			"unparsed nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{UnparsedObject: map[string]interface{}{"foo": "bar"}}),
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"unparsed nested in nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{UnparsedObject: map[string]interface{}{"foo": "bar"}}}),
			true,
			map[string]interface{}{"foo": "bar"},
		},
		{
			"valid nullable",
			datadogV2.NewNullableLogsArchiveDestination(&datadogV2.LogsArchiveDestination{LogsArchiveDestinationAzure: &datadogV2.LogsArchiveDestinationAzure{Type: datadogV2.LOGSARCHIVEDESTINATIONAZURETYPE_AZURE}}),
			false,
			nil,
		},
		{
			"valid struct",
			datadogV1.SyntheticsAPITest{Name: "foo", Type: datadogV1.SYNTHETICSAPITESTTYPE_API, Config: datadogV1.SyntheticsAPITestConfig{Assertions: []datadogV1.SyntheticsAssertion{{SyntheticsAssertionTarget: &datadogV1.SyntheticsAssertionTarget{Type: datadogV1.SYNTHETICSASSERTIONTYPE_BODY, Operator: datadogV1.SYNTHETICSASSERTIONOPERATOR_CONTAINS}}}}},
			false,
			nil,
		},
		{
			"valid simple type",
			"a simple string",
			false,
			nil,
		},
		{
			"valid simple pointer",
			datadog.PtrString("a simple pointer to string"),
			false,
			nil,
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(tc.name, func(t *testing.T) {
			n, m := datadog.ContainsUnparsedObject(c.value)
			assert.Equal(c.expectedUnparsedObject, m)
			assert.Equal(c.expectedBool, n)
		})
	}
}
