// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCloudAuthIntakeMappingUpdateAttributes The complete set of workload capabilities to allow.
type AWSCloudAuthIntakeMappingUpdateAttributes struct {
	// Explicit workload capabilities enabled on an intake mapping. The supported capability is par_enrollment. Enabling or disabling it requires `on_prem_runner_write` in addition to mapping write permission. An omitted list on a new mapping enables no capabilities. Unknown and duplicate values are rejected.
	Capabilities []string `json:"capabilities"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCloudAuthIntakeMappingUpdateAttributes instantiates a new AWSCloudAuthIntakeMappingUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCloudAuthIntakeMappingUpdateAttributes(capabilities []string) *AWSCloudAuthIntakeMappingUpdateAttributes {
	this := AWSCloudAuthIntakeMappingUpdateAttributes{}
	this.Capabilities = capabilities
	return &this
}

// NewAWSCloudAuthIntakeMappingUpdateAttributesWithDefaults instantiates a new AWSCloudAuthIntakeMappingUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCloudAuthIntakeMappingUpdateAttributesWithDefaults() *AWSCloudAuthIntakeMappingUpdateAttributes {
	this := AWSCloudAuthIntakeMappingUpdateAttributes{}
	return &this
}

// GetCapabilities returns the Capabilities field value.
func (o *AWSCloudAuthIntakeMappingUpdateAttributes) GetCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingUpdateAttributes) GetCapabilitiesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// SetCapabilities sets field value.
func (o *AWSCloudAuthIntakeMappingUpdateAttributes) SetCapabilities(v []string) {
	o.Capabilities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCloudAuthIntakeMappingUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["capabilities"] = o.Capabilities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCloudAuthIntakeMappingUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Capabilities *[]string `json:"capabilities"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Capabilities == nil {
		return fmt.Errorf("required field capabilities missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"capabilities"})
	} else {
		return err
	}
	o.Capabilities = *all.Capabilities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
