// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceMitigationRequestAttributes The attributes of a governance mitigation request.
type GovernanceMitigationRequestAttributes struct {
	// The identifiers of the detections to mitigate in this request.
	DetectionIds []string `json:"detection_ids,omitempty"`
	// The detection type whose detections should be mitigated.
	DetectionType *string `json:"detection_type,omitempty"`
	// A free-form map of parameter names to their configured values.
	MitigationParameters map[string]interface{} `json:"mitigation_parameters,omitempty"`
	// The mitigation to apply to the selected detections. Defaults to the control's configured mitigation when omitted.
	MitigationType *string `json:"mitigation_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceMitigationRequestAttributes instantiates a new GovernanceMitigationRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceMitigationRequestAttributes() *GovernanceMitigationRequestAttributes {
	this := GovernanceMitigationRequestAttributes{}
	return &this
}

// NewGovernanceMitigationRequestAttributesWithDefaults instantiates a new GovernanceMitigationRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceMitigationRequestAttributesWithDefaults() *GovernanceMitigationRequestAttributes {
	this := GovernanceMitigationRequestAttributes{}
	return &this
}

// GetDetectionIds returns the DetectionIds field value if set, zero value otherwise.
func (o *GovernanceMitigationRequestAttributes) GetDetectionIds() []string {
	if o == nil || o.DetectionIds == nil {
		var ret []string
		return ret
	}
	return o.DetectionIds
}

// GetDetectionIdsOk returns a tuple with the DetectionIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceMitigationRequestAttributes) GetDetectionIdsOk() (*[]string, bool) {
	if o == nil || o.DetectionIds == nil {
		return nil, false
	}
	return &o.DetectionIds, true
}

// HasDetectionIds returns a boolean if a field has been set.
func (o *GovernanceMitigationRequestAttributes) HasDetectionIds() bool {
	return o != nil && o.DetectionIds != nil
}

// SetDetectionIds gets a reference to the given []string and assigns it to the DetectionIds field.
func (o *GovernanceMitigationRequestAttributes) SetDetectionIds(v []string) {
	o.DetectionIds = v
}

// GetDetectionType returns the DetectionType field value if set, zero value otherwise.
func (o *GovernanceMitigationRequestAttributes) GetDetectionType() string {
	if o == nil || o.DetectionType == nil {
		var ret string
		return ret
	}
	return *o.DetectionType
}

// GetDetectionTypeOk returns a tuple with the DetectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceMitigationRequestAttributes) GetDetectionTypeOk() (*string, bool) {
	if o == nil || o.DetectionType == nil {
		return nil, false
	}
	return o.DetectionType, true
}

// HasDetectionType returns a boolean if a field has been set.
func (o *GovernanceMitigationRequestAttributes) HasDetectionType() bool {
	return o != nil && o.DetectionType != nil
}

// SetDetectionType gets a reference to the given string and assigns it to the DetectionType field.
func (o *GovernanceMitigationRequestAttributes) SetDetectionType(v string) {
	o.DetectionType = &v
}

// GetMitigationParameters returns the MitigationParameters field value if set, zero value otherwise.
func (o *GovernanceMitigationRequestAttributes) GetMitigationParameters() map[string]interface{} {
	if o == nil || o.MitigationParameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MitigationParameters
}

// GetMitigationParametersOk returns a tuple with the MitigationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceMitigationRequestAttributes) GetMitigationParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.MitigationParameters == nil {
		return nil, false
	}
	return &o.MitigationParameters, true
}

// HasMitigationParameters returns a boolean if a field has been set.
func (o *GovernanceMitigationRequestAttributes) HasMitigationParameters() bool {
	return o != nil && o.MitigationParameters != nil
}

// SetMitigationParameters gets a reference to the given map[string]interface{} and assigns it to the MitigationParameters field.
func (o *GovernanceMitigationRequestAttributes) SetMitigationParameters(v map[string]interface{}) {
	o.MitigationParameters = v
}

// GetMitigationType returns the MitigationType field value if set, zero value otherwise.
func (o *GovernanceMitigationRequestAttributes) GetMitigationType() string {
	if o == nil || o.MitigationType == nil {
		var ret string
		return ret
	}
	return *o.MitigationType
}

// GetMitigationTypeOk returns a tuple with the MitigationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceMitigationRequestAttributes) GetMitigationTypeOk() (*string, bool) {
	if o == nil || o.MitigationType == nil {
		return nil, false
	}
	return o.MitigationType, true
}

// HasMitigationType returns a boolean if a field has been set.
func (o *GovernanceMitigationRequestAttributes) HasMitigationType() bool {
	return o != nil && o.MitigationType != nil
}

// SetMitigationType gets a reference to the given string and assigns it to the MitigationType field.
func (o *GovernanceMitigationRequestAttributes) SetMitigationType(v string) {
	o.MitigationType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceMitigationRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DetectionIds != nil {
		toSerialize["detection_ids"] = o.DetectionIds
	}
	if o.DetectionType != nil {
		toSerialize["detection_type"] = o.DetectionType
	}
	if o.MitigationParameters != nil {
		toSerialize["mitigation_parameters"] = o.MitigationParameters
	}
	if o.MitigationType != nil {
		toSerialize["mitigation_type"] = o.MitigationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceMitigationRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DetectionIds         []string               `json:"detection_ids,omitempty"`
		DetectionType        *string                `json:"detection_type,omitempty"`
		MitigationParameters map[string]interface{} `json:"mitigation_parameters,omitempty"`
		MitigationType       *string                `json:"mitigation_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detection_ids", "detection_type", "mitigation_parameters", "mitigation_type"})
	} else {
		return err
	}
	o.DetectionIds = all.DetectionIds
	o.DetectionType = all.DetectionType
	o.MitigationParameters = all.MitigationParameters
	o.MitigationType = all.MitigationType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
