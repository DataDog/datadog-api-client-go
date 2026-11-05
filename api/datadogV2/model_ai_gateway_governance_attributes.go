// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayGovernanceAttributes Organization-wide AI Gateway governance controls.
type AiGatewayGovernanceAttributes struct {
	// Controls whether AI Gateway captures request and response content in addition to metadata.
	DataCaptureMode AiGatewayDataCaptureMode `json:"data_capture_mode"`
	// Items Metadata-only capture adds, each listed once. An item requires the items whose information it already reveals: `tool_output_size`, `tool_names` and `conversation_shape` require `tool_usage`, and `input_size_by_role` requires `tool_output_size`. A set that leaves out a required item is rejected. When several Metadata-only overrides match a request, it gets only the items all of them select.
	MetadataCapture []AiGatewayMetadataCaptureItem `json:"metadata_capture,omitempty"`
	// Zero data retention requirement for the organization's requests.
	ZdrMode AiGatewayZDRMode `json:"zdr_mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayGovernanceAttributes instantiates a new AiGatewayGovernanceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayGovernanceAttributes(dataCaptureMode AiGatewayDataCaptureMode, zdrMode AiGatewayZDRMode) *AiGatewayGovernanceAttributes {
	this := AiGatewayGovernanceAttributes{}
	this.DataCaptureMode = dataCaptureMode
	this.ZdrMode = zdrMode
	return &this
}

// NewAiGatewayGovernanceAttributesWithDefaults instantiates a new AiGatewayGovernanceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayGovernanceAttributesWithDefaults() *AiGatewayGovernanceAttributes {
	this := AiGatewayGovernanceAttributes{}
	return &this
}

// GetDataCaptureMode returns the DataCaptureMode field value.
func (o *AiGatewayGovernanceAttributes) GetDataCaptureMode() AiGatewayDataCaptureMode {
	if o == nil {
		var ret AiGatewayDataCaptureMode
		return ret
	}
	return o.DataCaptureMode
}

// GetDataCaptureModeOk returns a tuple with the DataCaptureMode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayGovernanceAttributes) GetDataCaptureModeOk() (*AiGatewayDataCaptureMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataCaptureMode, true
}

// SetDataCaptureMode sets field value.
func (o *AiGatewayGovernanceAttributes) SetDataCaptureMode(v AiGatewayDataCaptureMode) {
	o.DataCaptureMode = v
}

// GetMetadataCapture returns the MetadataCapture field value if set, zero value otherwise.
func (o *AiGatewayGovernanceAttributes) GetMetadataCapture() []AiGatewayMetadataCaptureItem {
	if o == nil || o.MetadataCapture == nil {
		var ret []AiGatewayMetadataCaptureItem
		return ret
	}
	return o.MetadataCapture
}

// GetMetadataCaptureOk returns a tuple with the MetadataCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayGovernanceAttributes) GetMetadataCaptureOk() (*[]AiGatewayMetadataCaptureItem, bool) {
	if o == nil || o.MetadataCapture == nil {
		return nil, false
	}
	return &o.MetadataCapture, true
}

// HasMetadataCapture returns a boolean if a field has been set.
func (o *AiGatewayGovernanceAttributes) HasMetadataCapture() bool {
	return o != nil && o.MetadataCapture != nil
}

// SetMetadataCapture gets a reference to the given []AiGatewayMetadataCaptureItem and assigns it to the MetadataCapture field.
func (o *AiGatewayGovernanceAttributes) SetMetadataCapture(v []AiGatewayMetadataCaptureItem) {
	o.MetadataCapture = v
}

// GetZdrMode returns the ZdrMode field value.
func (o *AiGatewayGovernanceAttributes) GetZdrMode() AiGatewayZDRMode {
	if o == nil {
		var ret AiGatewayZDRMode
		return ret
	}
	return o.ZdrMode
}

// GetZdrModeOk returns a tuple with the ZdrMode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayGovernanceAttributes) GetZdrModeOk() (*AiGatewayZDRMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ZdrMode, true
}

// SetZdrMode sets field value.
func (o *AiGatewayGovernanceAttributes) SetZdrMode(v AiGatewayZDRMode) {
	o.ZdrMode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayGovernanceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data_capture_mode"] = o.DataCaptureMode
	if o.MetadataCapture != nil {
		toSerialize["metadata_capture"] = o.MetadataCapture
	}
	toSerialize["zdr_mode"] = o.ZdrMode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayGovernanceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DataCaptureMode *AiGatewayDataCaptureMode      `json:"data_capture_mode"`
		MetadataCapture []AiGatewayMetadataCaptureItem `json:"metadata_capture,omitempty"`
		ZdrMode         *AiGatewayZDRMode              `json:"zdr_mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DataCaptureMode == nil {
		return fmt.Errorf("required field data_capture_mode missing")
	}
	if all.ZdrMode == nil {
		return fmt.Errorf("required field zdr_mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data_capture_mode", "metadata_capture", "zdr_mode"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.DataCaptureMode.IsValid() {
		hasInvalidField = true
	} else {
		o.DataCaptureMode = *all.DataCaptureMode
	}
	o.MetadataCapture = all.MetadataCapture
	if !all.ZdrMode.IsValid() {
		hasInvalidField = true
	} else {
		o.ZdrMode = *all.ZdrMode
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
