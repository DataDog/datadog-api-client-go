// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems The definition of `ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems` object.
type ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems struct {
	// The `items` `name`.
	Name *string `json:"name,omitempty"`
	// The `items` `type`.
	Type *string `json:"type,omitempty"`
	// The `items` `value`.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems instantiates a new ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems() *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems {
	this := ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems{}
	return &this
}

// NewResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItemsWithDefaults instantiates a new ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItemsWithDefaults() *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems {
	this := ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ResolveVulnerableSymbolsResponseDataAttributesResultsItemsVulnerableSymbolsItemsSymbolsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name  *string `json:"name,omitempty"`
		Type  *string `json:"type,omitempty"`
		Value *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "type", "value"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
