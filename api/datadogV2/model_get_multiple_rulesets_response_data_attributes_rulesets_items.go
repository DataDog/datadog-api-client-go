// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetMultipleRulesetsResponseDataAttributesRulesetsItems A ruleset returned in the response, containing its metadata and associated rules.
type GetMultipleRulesetsResponseDataAttributesRulesetsItems struct {
	// A detailed description of the ruleset's purpose and the types of issues it targets.
	Description string `json:"description"`
	// The unique identifier of the ruleset, which is the same as its name.
	Id string `json:"id"`
	// The unique name of the ruleset.
	Name string `json:"name"`
	// The list of static analysis rules included in this ruleset.
	Rules []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems `json:"rules"`
	// A brief summary of the ruleset, suitable for display in listings.
	ShortDescription string `json:"short_description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItems instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItems(description string, id string, name string, rules []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems, shortDescription string) *GetMultipleRulesetsResponseDataAttributesRulesetsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItems{}
	this.Description = description
	this.Id = id
	this.Name = name
	this.Rules = rules
	this.ShortDescription = shortDescription
	return &this
}

// NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsWithDefaults instantiates a new GetMultipleRulesetsResponseDataAttributesRulesetsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetMultipleRulesetsResponseDataAttributesRulesetsItemsWithDefaults() *GetMultipleRulesetsResponseDataAttributesRulesetsItems {
	this := GetMultipleRulesetsResponseDataAttributesRulesetsItems{}
	return &this
}

// GetDescription returns the Description field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetName(v string) {
	o.Name = v
}

// GetRules returns the Rules field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetRules() []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems {
	if o == nil {
		var ret []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetRulesOk() (*[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rules, true
}

// SetRules sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetRules(v []GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems) {
	o.Rules = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) SetShortDescription(v string) {
	o.ShortDescription = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetMultipleRulesetsResponseDataAttributesRulesetsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["rules"] = o.Rules
	toSerialize["short_description"] = o.ShortDescription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetMultipleRulesetsResponseDataAttributesRulesetsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string                                                             `json:"description"`
		Id               *string                                                             `json:"id"`
		Name             *string                                                             `json:"name"`
		Rules            *[]GetMultipleRulesetsResponseDataAttributesRulesetsItemsRulesItems `json:"rules"`
		ShortDescription *string                                                             `json:"short_description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rules == nil {
		return fmt.Errorf("required field rules missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "id", "name", "rules", "short_description"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Id = *all.Id
	o.Name = *all.Name
	o.Rules = *all.Rules
	o.ShortDescription = *all.ShortDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
