// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems
type GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems struct {
	//
	Description *string `json:"description,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems{}
	return &this
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItemsWithDefaults instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItemsWithDefaults() *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItemsArgumentsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string `json:"description,omitempty"`
		Name        *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
