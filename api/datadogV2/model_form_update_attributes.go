// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FormUpdateAttributes Attributes for updating a form.
type FormUpdateAttributes struct {
	// Update parameters for the form.
	FormUpdate *FormUpdateAttributesFormUpdate `json:"form_update,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFormUpdateAttributes instantiates a new FormUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFormUpdateAttributes() *FormUpdateAttributes {
	this := FormUpdateAttributes{}
	return &this
}

// NewFormUpdateAttributesWithDefaults instantiates a new FormUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFormUpdateAttributesWithDefaults() *FormUpdateAttributes {
	this := FormUpdateAttributes{}
	return &this
}

// GetFormUpdate returns the FormUpdate field value if set, zero value otherwise.
func (o *FormUpdateAttributes) GetFormUpdate() FormUpdateAttributesFormUpdate {
	if o == nil || o.FormUpdate == nil {
		var ret FormUpdateAttributesFormUpdate
		return ret
	}
	return *o.FormUpdate
}

// GetFormUpdateOk returns a tuple with the FormUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormUpdateAttributes) GetFormUpdateOk() (*FormUpdateAttributesFormUpdate, bool) {
	if o == nil || o.FormUpdate == nil {
		return nil, false
	}
	return o.FormUpdate, true
}

// HasFormUpdate returns a boolean if a field has been set.
func (o *FormUpdateAttributes) HasFormUpdate() bool {
	return o != nil && o.FormUpdate != nil
}

// SetFormUpdate gets a reference to the given FormUpdateAttributesFormUpdate and assigns it to the FormUpdate field.
func (o *FormUpdateAttributes) SetFormUpdate(v FormUpdateAttributesFormUpdate) {
	o.FormUpdate = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FormUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FormUpdate != nil {
		toSerialize["form_update"] = o.FormUpdate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FormUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FormUpdate *FormUpdateAttributesFormUpdate `json:"form_update,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"form_update"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.FormUpdate != nil && all.FormUpdate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.FormUpdate = all.FormUpdate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
