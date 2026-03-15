// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXAffect Reference to a component affected by a vulnerability.
type CycloneDXAffect struct {
	// Reference to a component's bom-ref.
	Ref string `json:"ref"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXAffect instantiates a new CycloneDXAffect object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXAffect(ref string) *CycloneDXAffect {
	this := CycloneDXAffect{}
	this.Ref = ref
	return &this
}

// NewCycloneDXAffectWithDefaults instantiates a new CycloneDXAffect object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXAffectWithDefaults() *CycloneDXAffect {
	this := CycloneDXAffect{}
	return &this
}

// GetRef returns the Ref field value.
func (o *CycloneDXAffect) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *CycloneDXAffect) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value.
func (o *CycloneDXAffect) SetRef(v string) {
	o.Ref = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXAffect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ref"] = o.Ref

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CycloneDXAffect) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ref *string `json:"ref"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Ref == nil {
		return fmt.Errorf("required field ref missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ref"})
	} else {
		return err
	}
	o.Ref = *all.Ref

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
