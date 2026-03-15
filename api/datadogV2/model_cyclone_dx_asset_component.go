// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXAssetComponent The asset component represents the system or host being scanned.
type CycloneDXAssetComponent struct {
	// Optional reference to a component in the components list.
	BomRef *string `json:"bom-ref,omitempty"`
	// The name of the asset.
	Name string `json:"name"`
	// The type of the asset component.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXAssetComponent instantiates a new CycloneDXAssetComponent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXAssetComponent(name string) *CycloneDXAssetComponent {
	this := CycloneDXAssetComponent{}
	this.Name = name
	return &this
}

// NewCycloneDXAssetComponentWithDefaults instantiates a new CycloneDXAssetComponent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXAssetComponentWithDefaults() *CycloneDXAssetComponent {
	this := CycloneDXAssetComponent{}
	return &this
}

// GetBomRef returns the BomRef field value if set, zero value otherwise.
func (o *CycloneDXAssetComponent) GetBomRef() string {
	if o == nil || o.BomRef == nil {
		var ret string
		return ret
	}
	return *o.BomRef
}

// GetBomRefOk returns a tuple with the BomRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycloneDXAssetComponent) GetBomRefOk() (*string, bool) {
	if o == nil || o.BomRef == nil {
		return nil, false
	}
	return o.BomRef, true
}

// HasBomRef returns a boolean if a field has been set.
func (o *CycloneDXAssetComponent) HasBomRef() bool {
	return o != nil && o.BomRef != nil
}

// SetBomRef gets a reference to the given string and assigns it to the BomRef field.
func (o *CycloneDXAssetComponent) SetBomRef(v string) {
	o.BomRef = &v
}

// GetName returns the Name field value.
func (o *CycloneDXAssetComponent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CycloneDXAssetComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CycloneDXAssetComponent) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CycloneDXAssetComponent) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycloneDXAssetComponent) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CycloneDXAssetComponent) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CycloneDXAssetComponent) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXAssetComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BomRef != nil {
		toSerialize["bom-ref"] = o.BomRef
	}
	toSerialize["name"] = o.Name
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CycloneDXAssetComponent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BomRef *string `json:"bom-ref,omitempty"`
		Name   *string `json:"name"`
		Type   *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bom-ref", "name", "type"})
	} else {
		return err
	}
	o.BomRef = all.BomRef
	o.Name = *all.Name
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
