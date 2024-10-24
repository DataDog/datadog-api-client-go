// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileTestBinding Objects describing the binding used for a mobile test.
type SyntheticsMobileTestBinding struct {
	// List of principals for a mobile test binding.
	Principals []string `json:"principals,omitempty"`
	// The type of relation for the binding.
	Relation *SyntheticsMobileTestBindingRelation `json:"relation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileTestBinding instantiates a new SyntheticsMobileTestBinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileTestBinding() *SyntheticsMobileTestBinding {
	this := SyntheticsMobileTestBinding{}
	return &this
}

// NewSyntheticsMobileTestBindingWithDefaults instantiates a new SyntheticsMobileTestBinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileTestBindingWithDefaults() *SyntheticsMobileTestBinding {
	this := SyntheticsMobileTestBinding{}
	return &this
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *SyntheticsMobileTestBinding) GetPrincipals() []string {
	if o == nil || o.Principals == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestBinding) GetPrincipalsOk() (*[]string, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return &o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *SyntheticsMobileTestBinding) HasPrincipals() bool {
	return o != nil && o.Principals != nil
}

// SetPrincipals gets a reference to the given []string and assigns it to the Principals field.
func (o *SyntheticsMobileTestBinding) SetPrincipals(v []string) {
	o.Principals = v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *SyntheticsMobileTestBinding) GetRelation() SyntheticsMobileTestBindingRelation {
	if o == nil || o.Relation == nil {
		var ret SyntheticsMobileTestBindingRelation
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileTestBinding) GetRelationOk() (*SyntheticsMobileTestBindingRelation, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *SyntheticsMobileTestBinding) HasRelation() bool {
	return o != nil && o.Relation != nil
}

// SetRelation gets a reference to the given SyntheticsMobileTestBindingRelation and assigns it to the Relation field.
func (o *SyntheticsMobileTestBinding) SetRelation(v SyntheticsMobileTestBindingRelation) {
	o.Relation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileTestBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileTestBinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Principals []string                             `json:"principals,omitempty"`
		Relation   *SyntheticsMobileTestBindingRelation `json:"relation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"principals", "relation"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Principals = all.Principals
	if all.Relation != nil && !all.Relation.IsValid() {
		hasInvalidField = true
	} else {
		o.Relation = all.Relation
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
