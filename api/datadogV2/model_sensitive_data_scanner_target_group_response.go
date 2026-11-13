// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerTargetGroupResponse A scanning group of a target, with its rules included.
type SensitiveDataScannerTargetGroupResponse struct {
	// A scanning group of a target.
	Data SensitiveDataScannerTargetGroupData `json:"data"`
	// Rules of the returned groups.
	Included []SensitiveDataScannerRuleIncludedItem `json:"included,omitempty"`
	// Meta response. Always an empty object.
	Meta map[string]interface{} `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerTargetGroupResponse instantiates a new SensitiveDataScannerTargetGroupResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerTargetGroupResponse(data SensitiveDataScannerTargetGroupData, meta map[string]interface{}) *SensitiveDataScannerTargetGroupResponse {
	this := SensitiveDataScannerTargetGroupResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewSensitiveDataScannerTargetGroupResponseWithDefaults instantiates a new SensitiveDataScannerTargetGroupResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerTargetGroupResponseWithDefaults() *SensitiveDataScannerTargetGroupResponse {
	this := SensitiveDataScannerTargetGroupResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *SensitiveDataScannerTargetGroupResponse) GetData() SensitiveDataScannerTargetGroupData {
	if o == nil {
		var ret SensitiveDataScannerTargetGroupData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupResponse) GetDataOk() (*SensitiveDataScannerTargetGroupData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *SensitiveDataScannerTargetGroupResponse) SetData(v SensitiveDataScannerTargetGroupData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *SensitiveDataScannerTargetGroupResponse) GetIncluded() []SensitiveDataScannerRuleIncludedItem {
	if o == nil || o.Included == nil {
		var ret []SensitiveDataScannerRuleIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupResponse) GetIncludedOk() (*[]SensitiveDataScannerRuleIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *SensitiveDataScannerTargetGroupResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []SensitiveDataScannerRuleIncludedItem and assigns it to the Included field.
func (o *SensitiveDataScannerTargetGroupResponse) SetIncluded(v []SensitiveDataScannerRuleIncludedItem) {
	o.Included = v
}

// GetMeta returns the Meta field value.
func (o *SensitiveDataScannerTargetGroupResponse) GetMeta() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupResponse) GetMetaOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *SensitiveDataScannerTargetGroupResponse) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerTargetGroupResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerTargetGroupResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *SensitiveDataScannerTargetGroupData   `json:"data"`
		Included []SensitiveDataScannerRuleIncludedItem `json:"included,omitempty"`
		Meta     *map[string]interface{}                `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Included = all.Included
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
