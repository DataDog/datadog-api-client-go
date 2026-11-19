// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCloudAuthIntakeMappingCreateAttributes Attributes for creating an AWS cloud authentication intake mapping.
type AWSCloudAuthIntakeMappingCreateAttributes struct {
	// AWS caller ARN pattern allowed to authenticate for telemetry submission. ARN syntax and AWS integration membership are validated by the server. For an assumed role, use the STS assumed-role ARN with a trailing /* wildcard to match its sessions.
	ArnPattern string `json:"arn_pattern"`
	// Explicit workload capabilities enabled on an intake mapping. The supported capability is par_enrollment. Enabling or disabling it requires `on_prem_runner_write` in addition to mapping write permission. An omitted list on a new mapping enables no capabilities. Unknown and duplicate values are rejected.
	Capabilities []string `json:"capabilities,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCloudAuthIntakeMappingCreateAttributes instantiates a new AWSCloudAuthIntakeMappingCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCloudAuthIntakeMappingCreateAttributes(arnPattern string) *AWSCloudAuthIntakeMappingCreateAttributes {
	this := AWSCloudAuthIntakeMappingCreateAttributes{}
	this.ArnPattern = arnPattern
	return &this
}

// NewAWSCloudAuthIntakeMappingCreateAttributesWithDefaults instantiates a new AWSCloudAuthIntakeMappingCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCloudAuthIntakeMappingCreateAttributesWithDefaults() *AWSCloudAuthIntakeMappingCreateAttributes {
	this := AWSCloudAuthIntakeMappingCreateAttributes{}
	return &this
}

// GetArnPattern returns the ArnPattern field value.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) GetArnPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArnPattern
}

// GetArnPatternOk returns a tuple with the ArnPattern field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) GetArnPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArnPattern, true
}

// SetArnPattern sets field value.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) SetArnPattern(v string) {
	o.ArnPattern = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) GetCapabilities() []string {
	if o == nil || o.Capabilities == nil {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) GetCapabilitiesOk() (*[]string, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) HasCapabilities() bool {
	return o != nil && o.Capabilities != nil
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) SetCapabilities(v []string) {
	o.Capabilities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCloudAuthIntakeMappingCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arn_pattern"] = o.ArnPattern
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCloudAuthIntakeMappingCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArnPattern   *string  `json:"arn_pattern"`
		Capabilities []string `json:"capabilities,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ArnPattern == nil {
		return fmt.Errorf("required field arn_pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arn_pattern", "capabilities"})
	} else {
		return err
	}
	o.ArnPattern = *all.ArnPattern
	o.Capabilities = all.Capabilities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
