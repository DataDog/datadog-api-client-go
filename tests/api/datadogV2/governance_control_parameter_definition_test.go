package test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/DataDog/datadog-api-client-go/v2/tests"
)

const governanceWorkflowParameter = `{
	"name": "workflowId",
	"display_name": "Workflow",
	"description": "The ID of the workflow to use for the mitigation.",
	"type": "string",
	"required": true,
	"supported_values": null,
	"default_value": null
}`

func governanceParameter(defaultValue string) string {
	field := ""
	if defaultValue != "" {
		field = `, "default_value": ` + defaultValue
	}
	return `{
		"name": "param",
		"display_name": "Param",
		"description": "A parameter.",
		"type": "string",
		"required": false,
		"supported_values": null` + field + `
	}`
}

func TestGovernanceControlParameterDefinitionDefaultValue(t *testing.T) {
	testCases := []struct {
		name     string
		json     string
		expected interface{}
	}{
		{"null", "null", nil},
		{"false", "false", false},
		{"zero", "0", json.Number("0")},
		{"large integer", "9007199254740993", json.Number("9007199254740993")},
		{"empty string", `""`, ""},
		{"string", `"abc"`, "abc"},
		{"object", `{"key": "value"}`, map[string]interface{}{"key": "value"}},
		{"nested large integer", `{"limit": 9007199254740993}`, map[string]interface{}{"limit": json.Number("9007199254740993")}},
		{"list", `["a", 1]`, []interface{}{"a", json.Number("1")}},
		{"list with large integer", `[9007199254740993]`, []interface{}{json.Number("9007199254740993")}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert := tests.Assert(context.Background(), t)
			var parameter datadogV2.GovernanceControlParameterDefinition
			assert.NoError(datadog.Unmarshal([]byte(governanceParameter(tc.json)), &parameter))
			assert.Nil(parameter.UnparsedObject)
			assert.Equal(tc.expected, parameter.GetDefaultValue())

			value, ok := parameter.GetDefaultValueOk()
			assert.True(ok)
			if tc.expected == nil {
				assert.Nil(value)
			} else {
				assert.Equal(tc.expected, *value)
			}

			serialized, err := datadog.Marshal(parameter)
			assert.NoError(err)
			var roundTrip map[string]interface{}
			assert.NoError(datadog.UnmarshalUseNumber(serialized, &roundTrip))
			assert.Contains(roundTrip, "default_value")
			assert.Equal(tc.expected, roundTrip["default_value"])
		})
	}
}

func TestGovernanceControlParameterDefinitionMissingDefaultValue(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	var parameter datadogV2.GovernanceControlParameterDefinition
	err := datadog.Unmarshal([]byte(governanceParameter("")), &parameter)
	assert.EqualError(err, "required field default_value missing")

	var mitigation datadogV2.GovernanceControlMitigationDefinition
	assert.NoError(datadog.Unmarshal([]byte(`{
		"id": "run_workflow",
		"title": "Run workflow",
		"description": "Run a workflow.",
		"execution_modes": ["manual"],
		"permissions": [],
		"supported_parameters": [`+governanceParameter("")+`]
	}`), &mitigation))
	assert.NotNil(mitigation.UnparsedObject)
}

func TestGovernanceControlMitigationDefinitionWorkflowParameter(t *testing.T) {
	assert := tests.Assert(context.Background(), t)

	var mitigation datadogV2.GovernanceControlMitigationDefinition
	assert.NoError(datadog.Unmarshal([]byte(`{
		"id": "run_workflow",
		"title": "Run workflow",
		"description": "Run a workflow.",
		"execution_modes": ["manual"],
		"permissions": [],
		"supported_parameters": [`+governanceWorkflowParameter+`]
	}`), &mitigation))
	assert.Nil(mitigation.UnparsedObject)

	unparsed, _ := datadog.ContainsUnparsedObject(mitigation)
	assert.False(unparsed)

	parameters := mitigation.GetSupportedParameters()
	assert.Len(parameters, 1)
	parameter := parameters[0]
	assert.Nil(parameter.UnparsedObject)
	assert.Equal("workflowId", parameter.GetName())
	assert.True(parameter.GetRequired())
	assert.Nil(parameter.GetSupportedValues())
	assert.Nil(parameter.GetDefaultValue())
	value, ok := parameter.GetDefaultValueOk()
	assert.Nil(value)
	assert.True(ok)
}
