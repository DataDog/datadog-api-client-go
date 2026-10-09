// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SeverityOverrideResponse Response for the severity override request.
type SeverityOverrideResponse struct {
	// Data of the severity override response.
	Data SeverityOverrideResponseData `json:"data"`
	// Security findings skipped while processing the severity override request.
	Meta *SeverityOverrideResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSeverityOverrideResponse instantiates a new SeverityOverrideResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSeverityOverrideResponse(data SeverityOverrideResponseData) *SeverityOverrideResponse {
	this := SeverityOverrideResponse{}
	this.Data = data
	return &this
}

// NewSeverityOverrideResponseWithDefaults instantiates a new SeverityOverrideResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSeverityOverrideResponseWithDefaults() *SeverityOverrideResponse {
	this := SeverityOverrideResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *SeverityOverrideResponse) GetData() SeverityOverrideResponseData {
	if o == nil {
		var ret SeverityOverrideResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SeverityOverrideResponse) GetDataOk() (*SeverityOverrideResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *SeverityOverrideResponse) SetData(v SeverityOverrideResponseData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SeverityOverrideResponse) GetMeta() SeverityOverrideResponseMeta {
	if o == nil || o.Meta == nil {
		var ret SeverityOverrideResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeverityOverrideResponse) GetMetaOk() (*SeverityOverrideResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SeverityOverrideResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given SeverityOverrideResponseMeta and assigns it to the Meta field.
func (o *SeverityOverrideResponse) SetMeta(v SeverityOverrideResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SeverityOverrideResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SeverityOverrideResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *SeverityOverrideResponseData `json:"data"`
		Meta *SeverityOverrideResponseMeta `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
