package api

import (
	"encoding/json"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/stretchr/testify/require"
)

func TestWIFCapabilityUpdateCanClearAllCapabilities(t *testing.T) {
	attributes := datadogV2.NewAWSCloudAuthIntakeMappingUpdateAttributes([]string{})
	data := datadogV2.NewAWSCloudAuthIntakeMappingUpdateData(*attributes, datadogV2.AWSCLOUDAUTHINTAKEMAPPINGTYPE_AWS_CLOUD_AUTH_INTAKE_MAPPING)
	request := datadogV2.NewAWSCloudAuthIntakeMappingUpdateRequest(*data)
	encoded, err := json.Marshal(request)
	require.NoError(t, err)
	var wire struct {
		Data struct {
			Attributes map[string]json.RawMessage `json:"attributes"`
		} `json:"data"`
	}
	require.NoError(t, json.Unmarshal(encoded, &wire))
	require.JSONEq(t, "[]", string(wire.Data.Attributes["capabilities"]))
}

func TestWIFCapabilityModelsPreserveFutureNamesAndLegacyResponses(t *testing.T) {
	var attributes datadogV2.AWSCloudAuthIntakeMappingAttributesResponse
	require.NoError(t, json.Unmarshal([]byte(`{"arn_pattern":"arn:aws:sts::123456789012:assumed-role/agent/*","capabilities":["future_capability"]}`), &attributes))
	require.Nil(t, attributes.UnparsedObject)
	require.Equal(t, []string{"future_capability"}, []string(attributes.GetCapabilities()))

	var legacy datadogV2.AWSCloudAuthIntakeMappingAttributesResponse
	require.NoError(t, json.Unmarshal([]byte(`{"arn_pattern":"arn:aws:sts::123456789012:assumed-role/agent/*","par_enrollment_enabled":true}`), &legacy))
	require.Nil(t, legacy.UnparsedObject)
	require.True(t, legacy.GetParEnrollmentEnabled())
}

func TestWIFMappingCreateCanRequestCapabilities(t *testing.T) {
	attributes := datadogV2.NewAWSCloudAuthIntakeMappingCreateAttributes("arn:aws:sts::123456789012:assumed-role/agent/*")
	attributes.SetCapabilities([]string{"par_enrollment"})
	encoded, err := json.Marshal(attributes)
	require.NoError(t, err)
	require.JSONEq(t, `{"arn_pattern":"arn:aws:sts::123456789012:assumed-role/agent/*","capabilities":["par_enrollment"]}`, string(encoded))
}
