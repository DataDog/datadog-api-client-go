// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXReference External reference for a vulnerability.
type CycloneDXReference struct {
	// Identifier of the reference.
	Id *string `json:"id,omitempty"`
	// Source information for a reference.
	Source *CycloneDXReferenceSource `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXReference instantiates a new CycloneDXReference object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXReference() *CycloneDXReference {
	this := CycloneDXReference{}
	return &this
}

// NewCycloneDXReferenceWithDefaults instantiates a new CycloneDXReference object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXReferenceWithDefaults() *CycloneDXReference {
	this := CycloneDXReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CycloneDXReference) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycloneDXReference) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CycloneDXReference) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CycloneDXReference) SetId(v string) {
	o.Id = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *CycloneDXReference) GetSource() CycloneDXReferenceSource {
	if o == nil || o.Source == nil {
		var ret CycloneDXReferenceSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycloneDXReference) GetSourceOk() (*CycloneDXReferenceSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *CycloneDXReference) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given CycloneDXReferenceSource and assigns it to the Source field.
func (o *CycloneDXReference) SetSource(v CycloneDXReferenceSource) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *CycloneDXReference) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string                   `json:"id,omitempty"`
		Source *CycloneDXReferenceSource `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "source"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
