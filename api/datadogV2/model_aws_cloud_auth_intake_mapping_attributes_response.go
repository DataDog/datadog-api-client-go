// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCloudAuthIntakeMappingAttributesResponse Attributes for AWS cloud authentication intake mapping response.
type AWSCloudAuthIntakeMappingAttributesResponse struct {
	// AWS caller ARN pattern allowed to authenticate for telemetry submission. ARN syntax and AWS integration membership are validated by the server. For an assumed role, use the STS assumed-role ARN with a trailing /* wildcard to match its sessions.
	ArnPattern string `json:"arn_pattern"`
	// Whether the mapping allows Private Action Runner enrollment. This attribute may be omitted when enrollment is disabled. Enrollment permission is managed separately from the ARN pattern.
	ParEnrollmentEnabled *bool `json:"par_enrollment_enabled,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCloudAuthIntakeMappingAttributesResponse instantiates a new AWSCloudAuthIntakeMappingAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCloudAuthIntakeMappingAttributesResponse(arnPattern string) *AWSCloudAuthIntakeMappingAttributesResponse {
	this := AWSCloudAuthIntakeMappingAttributesResponse{}
	this.ArnPattern = arnPattern
	return &this
}

// NewAWSCloudAuthIntakeMappingAttributesResponseWithDefaults instantiates a new AWSCloudAuthIntakeMappingAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCloudAuthIntakeMappingAttributesResponseWithDefaults() *AWSCloudAuthIntakeMappingAttributesResponse {
	this := AWSCloudAuthIntakeMappingAttributesResponse{}
	return &this
}

// GetArnPattern returns the ArnPattern field value.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) GetArnPattern() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ArnPattern
}

// GetArnPatternOk returns a tuple with the ArnPattern field value
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) GetArnPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArnPattern, true
}

// SetArnPattern sets field value.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) SetArnPattern(v string) {
	o.ArnPattern = v
}

// GetParEnrollmentEnabled returns the ParEnrollmentEnabled field value if set, zero value otherwise.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) GetParEnrollmentEnabled() bool {
	if o == nil || o.ParEnrollmentEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ParEnrollmentEnabled
}

// GetParEnrollmentEnabledOk returns a tuple with the ParEnrollmentEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) GetParEnrollmentEnabledOk() (*bool, bool) {
	if o == nil || o.ParEnrollmentEnabled == nil {
		return nil, false
	}
	return o.ParEnrollmentEnabled, true
}

// HasParEnrollmentEnabled returns a boolean if a field has been set.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) HasParEnrollmentEnabled() bool {
	return o != nil && o.ParEnrollmentEnabled != nil
}

// SetParEnrollmentEnabled gets a reference to the given bool and assigns it to the ParEnrollmentEnabled field.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) SetParEnrollmentEnabled(v bool) {
	o.ParEnrollmentEnabled = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCloudAuthIntakeMappingAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["arn_pattern"] = o.ArnPattern
	if o.ParEnrollmentEnabled != nil {
		toSerialize["par_enrollment_enabled"] = o.ParEnrollmentEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCloudAuthIntakeMappingAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ArnPattern           *string `json:"arn_pattern"`
		ParEnrollmentEnabled *bool   `json:"par_enrollment_enabled,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ArnPattern == nil {
		return fmt.Errorf("required field arn_pattern missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arn_pattern", "par_enrollment_enabled"})
	} else {
		return err
	}
	o.ArnPattern = *all.ArnPattern
	o.ParEnrollmentEnabled = all.ParEnrollmentEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
