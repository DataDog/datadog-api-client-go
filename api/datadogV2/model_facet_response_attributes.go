// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FacetResponseAttributes Facet attributes.
type FacetResponseAttributes struct {
	// Array of facet values.
	Items []FacetItem `json:"items"`
	// The key of the facet.
	Key string `json:"key"`
	// The display name of the facet.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFacetResponseAttributes instantiates a new FacetResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFacetResponseAttributes(items []FacetItem, key string) *FacetResponseAttributes {
	this := FacetResponseAttributes{}
	this.Items = items
	this.Key = key
	return &this
}

// NewFacetResponseAttributesWithDefaults instantiates a new FacetResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFacetResponseAttributesWithDefaults() *FacetResponseAttributes {
	this := FacetResponseAttributes{}
	return &this
}

// GetItems returns the Items field value.
func (o *FacetResponseAttributes) GetItems() []FacetItem {
	if o == nil {
		var ret []FacetItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *FacetResponseAttributes) GetItemsOk() (*[]FacetItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *FacetResponseAttributes) SetItems(v []FacetItem) {
	o.Items = v
}

// GetKey returns the Key field value.
func (o *FacetResponseAttributes) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *FacetResponseAttributes) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *FacetResponseAttributes) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FacetResponseAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FacetResponseAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FacetResponseAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FacetResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	toSerialize["key"] = o.Key
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FacetResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items *[]FacetItem `json:"items"`
		Key   *string      `json:"key"`
		Name  *string      `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"items", "key", "name"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.Key = *all.Key
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
