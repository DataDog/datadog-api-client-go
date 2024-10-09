// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeEventCreateResponseAttributesAttributes The definition of `ChangeEventCreateResponseAttributesAttributes` object.
type ChangeEventCreateResponseAttributesAttributes struct {
	// The definition of `ChangeEventCreateResponseAttributesAttributesEvt` object.
	Evt *ChangeEventCreateResponseAttributesAttributesEvt `json:"evt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeEventCreateResponseAttributesAttributes instantiates a new ChangeEventCreateResponseAttributesAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeEventCreateResponseAttributesAttributes() *ChangeEventCreateResponseAttributesAttributes {
	this := ChangeEventCreateResponseAttributesAttributes{}
	return &this
}

// NewChangeEventCreateResponseAttributesAttributesWithDefaults instantiates a new ChangeEventCreateResponseAttributesAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeEventCreateResponseAttributesAttributesWithDefaults() *ChangeEventCreateResponseAttributesAttributes {
	this := ChangeEventCreateResponseAttributesAttributes{}
	return &this
}

// GetEvt returns the Evt field value if set, zero value otherwise.
func (o *ChangeEventCreateResponseAttributesAttributes) GetEvt() ChangeEventCreateResponseAttributesAttributesEvt {
	if o == nil || o.Evt == nil {
		var ret ChangeEventCreateResponseAttributesAttributesEvt
		return ret
	}
	return *o.Evt
}

// GetEvtOk returns a tuple with the Evt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeEventCreateResponseAttributesAttributes) GetEvtOk() (*ChangeEventCreateResponseAttributesAttributesEvt, bool) {
	if o == nil || o.Evt == nil {
		return nil, false
	}
	return o.Evt, true
}

// HasEvt returns a boolean if a field has been set.
func (o *ChangeEventCreateResponseAttributesAttributes) HasEvt() bool {
	return o != nil && o.Evt != nil
}

// SetEvt gets a reference to the given ChangeEventCreateResponseAttributesAttributesEvt and assigns it to the Evt field.
func (o *ChangeEventCreateResponseAttributesAttributes) SetEvt(v ChangeEventCreateResponseAttributesAttributesEvt) {
	o.Evt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeEventCreateResponseAttributesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Evt != nil {
		toSerialize["evt"] = o.Evt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeEventCreateResponseAttributesAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Evt *ChangeEventCreateResponseAttributesAttributesEvt `json:"evt,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"evt"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Evt != nil && all.Evt.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Evt = all.Evt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
