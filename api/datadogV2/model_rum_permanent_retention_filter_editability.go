// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMPermanentRetentionFilterEditability Indicates which cross-product fields of a permanent RUM retention filter can be updated.
type RUMPermanentRetentionFilterEditability struct {
	// Whether the APM trace cross-product configuration of the filter can be updated.
	TraceEditable *bool `json:"trace_editable,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMPermanentRetentionFilterEditability instantiates a new RUMPermanentRetentionFilterEditability object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMPermanentRetentionFilterEditability() *RUMPermanentRetentionFilterEditability {
	this := RUMPermanentRetentionFilterEditability{}
	return &this
}

// NewRUMPermanentRetentionFilterEditabilityWithDefaults instantiates a new RUMPermanentRetentionFilterEditability object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMPermanentRetentionFilterEditabilityWithDefaults() *RUMPermanentRetentionFilterEditability {
	this := RUMPermanentRetentionFilterEditability{}
	return &this
}

// GetTraceEditable returns the TraceEditable field value if set, zero value otherwise.
func (o *RUMPermanentRetentionFilterEditability) GetTraceEditable() bool {
	if o == nil || o.TraceEditable == nil {
		var ret bool
		return ret
	}
	return *o.TraceEditable
}

// GetTraceEditableOk returns a tuple with the TraceEditable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMPermanentRetentionFilterEditability) GetTraceEditableOk() (*bool, bool) {
	if o == nil || o.TraceEditable == nil {
		return nil, false
	}
	return o.TraceEditable, true
}

// HasTraceEditable returns a boolean if a field has been set.
func (o *RUMPermanentRetentionFilterEditability) HasTraceEditable() bool {
	return o != nil && o.TraceEditable != nil
}

// SetTraceEditable gets a reference to the given bool and assigns it to the TraceEditable field.
func (o *RUMPermanentRetentionFilterEditability) SetTraceEditable(v bool) {
	o.TraceEditable = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMPermanentRetentionFilterEditability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TraceEditable != nil {
		toSerialize["trace_editable"] = o.TraceEditable
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMPermanentRetentionFilterEditability) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TraceEditable *bool `json:"trace_editable,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"trace_editable"})
	} else {
		return err
	}
	o.TraceEditable = all.TraceEditable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
