// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelSuggestionResponseDataAttributes
type FunnelSuggestionResponseDataAttributes struct {
	//
	Actions []FunnelSuggestionResponseDataAttributesActionsItems `json:"actions,omitempty"`
	//
	Views []FunnelSuggestionResponseDataAttributesViewsItems `json:"views,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFunnelSuggestionResponseDataAttributes instantiates a new FunnelSuggestionResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelSuggestionResponseDataAttributes() *FunnelSuggestionResponseDataAttributes {
	this := FunnelSuggestionResponseDataAttributes{}
	return &this
}

// NewFunnelSuggestionResponseDataAttributesWithDefaults instantiates a new FunnelSuggestionResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelSuggestionResponseDataAttributesWithDefaults() *FunnelSuggestionResponseDataAttributes {
	this := FunnelSuggestionResponseDataAttributes{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *FunnelSuggestionResponseDataAttributes) GetActions() []FunnelSuggestionResponseDataAttributesActionsItems {
	if o == nil || o.Actions == nil {
		var ret []FunnelSuggestionResponseDataAttributesActionsItems
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionResponseDataAttributes) GetActionsOk() (*[]FunnelSuggestionResponseDataAttributesActionsItems, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return &o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *FunnelSuggestionResponseDataAttributes) HasActions() bool {
	return o != nil && o.Actions != nil
}

// SetActions gets a reference to the given []FunnelSuggestionResponseDataAttributesActionsItems and assigns it to the Actions field.
func (o *FunnelSuggestionResponseDataAttributes) SetActions(v []FunnelSuggestionResponseDataAttributesActionsItems) {
	o.Actions = v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *FunnelSuggestionResponseDataAttributes) GetViews() []FunnelSuggestionResponseDataAttributesViewsItems {
	if o == nil || o.Views == nil {
		var ret []FunnelSuggestionResponseDataAttributesViewsItems
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelSuggestionResponseDataAttributes) GetViewsOk() (*[]FunnelSuggestionResponseDataAttributesViewsItems, bool) {
	if o == nil || o.Views == nil {
		return nil, false
	}
	return &o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *FunnelSuggestionResponseDataAttributes) HasViews() bool {
	return o != nil && o.Views != nil
}

// SetViews gets a reference to the given []FunnelSuggestionResponseDataAttributesViewsItems and assigns it to the Views field.
func (o *FunnelSuggestionResponseDataAttributes) SetViews(v []FunnelSuggestionResponseDataAttributesViewsItems) {
	o.Views = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelSuggestionResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Views != nil {
		toSerialize["views"] = o.Views
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelSuggestionResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actions []FunnelSuggestionResponseDataAttributesActionsItems `json:"actions,omitempty"`
		Views   []FunnelSuggestionResponseDataAttributesViewsItems   `json:"views,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actions", "views"})
	} else {
		return err
	}
	o.Actions = all.Actions
	o.Views = all.Views

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
