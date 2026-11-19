// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchEvent An event to instrument for the launch.
type ProductAnalyticsLaunchEvent struct {
	// What the event records.
	Description *string `json:"description,omitempty"`
	// IDs of the launch KPIs that the event measures.
	Kpis []string `json:"kpis,omitempty"`
	// Name of the event, unique within the tracking plan.
	Name string `json:"name"`
	// Properties that the event carries.
	Properties []ProductAnalyticsLaunchEventProperty `json:"properties,omitempty"`
	// What causes the event to be sent.
	Trigger *string `json:"trigger,omitempty"`
	// Views where the event is sent. Each value must be unique and not blank.
	Views []string `json:"views,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchEvent instantiates a new ProductAnalyticsLaunchEvent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchEvent(name string) *ProductAnalyticsLaunchEvent {
	this := ProductAnalyticsLaunchEvent{}
	this.Name = name
	return &this
}

// NewProductAnalyticsLaunchEventWithDefaults instantiates a new ProductAnalyticsLaunchEvent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchEventWithDefaults() *ProductAnalyticsLaunchEvent {
	this := ProductAnalyticsLaunchEvent{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchEvent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEvent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchEvent) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProductAnalyticsLaunchEvent) SetDescription(v string) {
	o.Description = &v
}

// GetKpis returns the Kpis field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchEvent) GetKpis() []string {
	if o == nil || o.Kpis == nil {
		var ret []string
		return ret
	}
	return o.Kpis
}

// GetKpisOk returns a tuple with the Kpis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEvent) GetKpisOk() (*[]string, bool) {
	if o == nil || o.Kpis == nil {
		return nil, false
	}
	return &o.Kpis, true
}

// HasKpis returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchEvent) HasKpis() bool {
	return o != nil && o.Kpis != nil
}

// SetKpis gets a reference to the given []string and assigns it to the Kpis field.
func (o *ProductAnalyticsLaunchEvent) SetKpis(v []string) {
	o.Kpis = v
}

// GetName returns the Name field value.
func (o *ProductAnalyticsLaunchEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ProductAnalyticsLaunchEvent) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchEvent) GetProperties() []ProductAnalyticsLaunchEventProperty {
	if o == nil || o.Properties == nil {
		var ret []ProductAnalyticsLaunchEventProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEvent) GetPropertiesOk() (*[]ProductAnalyticsLaunchEventProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchEvent) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []ProductAnalyticsLaunchEventProperty and assigns it to the Properties field.
func (o *ProductAnalyticsLaunchEvent) SetProperties(v []ProductAnalyticsLaunchEventProperty) {
	o.Properties = v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchEvent) GetTrigger() string {
	if o == nil || o.Trigger == nil {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEvent) GetTriggerOk() (*string, bool) {
	if o == nil || o.Trigger == nil {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchEvent) HasTrigger() bool {
	return o != nil && o.Trigger != nil
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *ProductAnalyticsLaunchEvent) SetTrigger(v string) {
	o.Trigger = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchEvent) GetViews() []string {
	if o == nil || o.Views == nil {
		var ret []string
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchEvent) GetViewsOk() (*[]string, bool) {
	if o == nil || o.Views == nil {
		return nil, false
	}
	return &o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchEvent) HasViews() bool {
	return o != nil && o.Views != nil
}

// SetViews gets a reference to the given []string and assigns it to the Views field.
func (o *ProductAnalyticsLaunchEvent) SetViews(v []string) {
	o.Views = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Kpis != nil {
		toSerialize["kpis"] = o.Kpis
	}
	toSerialize["name"] = o.Name
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Trigger != nil {
		toSerialize["trigger"] = o.Trigger
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
func (o *ProductAnalyticsLaunchEvent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                               `json:"description,omitempty"`
		Kpis        []string                              `json:"kpis,omitempty"`
		Name        *string                               `json:"name"`
		Properties  []ProductAnalyticsLaunchEventProperty `json:"properties,omitempty"`
		Trigger     *string                               `json:"trigger,omitempty"`
		Views       []string                              `json:"views,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "kpis", "name", "properties", "trigger", "views"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Kpis = all.Kpis
	o.Name = *all.Name
	o.Properties = all.Properties
	o.Trigger = all.Trigger
	o.Views = all.Views

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
