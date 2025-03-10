// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFilterMeta The object describing metadata of a RUM retention filter.
type RumRetentionFilterMeta struct {
	// The origin source of the retention filter, which may be created or updated through the Datadog UI, terraform, or other channels such as direct API calls.
	Source *RumRetentionFilterSource `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRetentionFilterMeta instantiates a new RumRetentionFilterMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRetentionFilterMeta() *RumRetentionFilterMeta {
	this := RumRetentionFilterMeta{}
	return &this
}

// NewRumRetentionFilterMetaWithDefaults instantiates a new RumRetentionFilterMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRetentionFilterMetaWithDefaults() *RumRetentionFilterMeta {
	this := RumRetentionFilterMeta{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *RumRetentionFilterMeta) GetSource() RumRetentionFilterSource {
	if o == nil || o.Source == nil {
		var ret RumRetentionFilterSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFilterMeta) GetSourceOk() (*RumRetentionFilterSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *RumRetentionFilterMeta) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given RumRetentionFilterSource and assigns it to the Source field.
func (o *RumRetentionFilterMeta) SetSource(v RumRetentionFilterSource) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRetentionFilterMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRetentionFilterMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Source *RumRetentionFilterSource `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"source"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Source != nil && !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = all.Source
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
