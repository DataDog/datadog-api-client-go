// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SuggestionMetadata Optional metadata for a suggestion.
type SuggestionMetadata struct {
	// Variant ID for variant delete suggestions.
	VariantId *string `json:"variant_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSuggestionMetadata instantiates a new SuggestionMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSuggestionMetadata() *SuggestionMetadata {
	this := SuggestionMetadata{}
	return &this
}

// NewSuggestionMetadataWithDefaults instantiates a new SuggestionMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSuggestionMetadataWithDefaults() *SuggestionMetadata {
	this := SuggestionMetadata{}
	return &this
}

// GetVariantId returns the VariantId field value if set, zero value otherwise.
func (o *SuggestionMetadata) GetVariantId() string {
	if o == nil || o.VariantId == nil {
		var ret string
		return ret
	}
	return *o.VariantId
}

// GetVariantIdOk returns a tuple with the VariantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuggestionMetadata) GetVariantIdOk() (*string, bool) {
	if o == nil || o.VariantId == nil {
		return nil, false
	}
	return o.VariantId, true
}

// HasVariantId returns a boolean if a field has been set.
func (o *SuggestionMetadata) HasVariantId() bool {
	return o != nil && o.VariantId != nil
}

// SetVariantId gets a reference to the given string and assigns it to the VariantId field.
func (o *SuggestionMetadata) SetVariantId(v string) {
	o.VariantId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SuggestionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.VariantId != nil {
		toSerialize["variant_id"] = o.VariantId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SuggestionMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		VariantId *string `json:"variant_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"variant_id"})
	} else {
		return err
	}
	o.VariantId = all.VariantId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
