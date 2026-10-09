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

func TestNullableInterface(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	var unset datadog.NullableInterface
	assert.False(unset.IsSet())
	assert.Nil(unset.Get())

	testCases := []struct {
		name     string
		json     string
		expected interface{}
	}{
		{"null", "null", nil},
		{"false", "false", false},
		{"zero", "0", float64(0)},
		{"empty string", `""`, ""},
		{"string", `"abc"`, "abc"},
		{"object", `{"key":"value"}`, map[string]interface{}{"key": "value"}},
		{"list", `["a",1]`, []interface{}{"a", float64(1)}},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			var holder struct {
				Value datadog.NullableInterface `json:"value"`
			}
			assert.NoError(datadog.Unmarshal([]byte(`{"value":`+c.json+`}`), &holder))
			assert.True(holder.Value.IsSet())
			if c.expected == nil {
				assert.Nil(holder.Value.Get())
			} else {
				assert.Equal(c.expected, *holder.Value.Get())
			}

			serialized, err := datadog.Marshal(holder.Value)
			assert.NoError(err)
			var roundTrip interface{}
			assert.NoError(datadog.Unmarshal(serialized, &roundTrip))
			assert.Equal(c.expected, roundTrip)
		})
	}

	t.Run("absent", func(t *testing.T) {
		var holder struct {
			Value datadog.NullableInterface `json:"value"`
		}
		assert.NoError(datadog.Unmarshal([]byte(`{}`), &holder))
		assert.False(holder.Value.IsSet())
		assert.Nil(holder.Value.Get())
	})

	t.Run("set and unset", func(t *testing.T) {
		var value interface{} = "abc"
		nullable := datadog.NewNullableInterface(&value)
		assert.True(nullable.IsSet())
		assert.Equal("abc", *nullable.Get())

		nullable.Set(nil)
		assert.True(nullable.IsSet())
		assert.Nil(nullable.Get())
		serialized, err := datadog.Marshal(nullable)
		assert.NoError(err)
		assert.Equal("null", string(serialized))

		nullable.Unset()
		assert.False(nullable.IsSet())
		assert.Nil(nullable.Get())
	})
}
