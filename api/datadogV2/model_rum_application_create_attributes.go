// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationCreateAttributes RUM application creation attributes.
type RUMApplicationCreateAttributes struct {
	// Name of the RUM application.
	Name string `json:"name"`
	// Configures which RUM events are processed and stored for the application.
	RumEventProcessingState *RUMEventProcessingState `json:"rum_event_processing_state,omitempty"`
	// Controls the retention policy for product analytics data derived from RUM events.
	RumProductAnalyticsRetentionState *RUMProductAnalyticsRetentionState `json:"rum_product_analytics_retention_state,omitempty"`
	// Type of the RUM application. Supported values are `browser`, `ios`, `android`, `react-native`, `flutter`, `roku`, `electron`, `unity`, `kotlin-multiplatform`.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMApplicationCreateAttributes instantiates a new RUMApplicationCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationCreateAttributes(name string) *RUMApplicationCreateAttributes {
	this := RUMApplicationCreateAttributes{}
	this.Name = name
	return &this
}

// NewRUMApplicationCreateAttributesWithDefaults instantiates a new RUMApplicationCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationCreateAttributesWithDefaults() *RUMApplicationCreateAttributes {
	this := RUMApplicationCreateAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *RUMApplicationCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMApplicationCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetRumEventProcessingState returns the RumEventProcessingState field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetRumEventProcessingState() RUMEventProcessingState {
	if o == nil || o.RumEventProcessingState == nil {
		var ret RUMEventProcessingState
		return ret
	}
	return *o.RumEventProcessingState
}

// GetRumEventProcessingStateOk returns a tuple with the RumEventProcessingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetRumEventProcessingStateOk() (*RUMEventProcessingState, bool) {
	if o == nil || o.RumEventProcessingState == nil {
		return nil, false
	}
	return o.RumEventProcessingState, true
}

// HasRumEventProcessingState returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasRumEventProcessingState() bool {
	return o != nil && o.RumEventProcessingState != nil
}

// SetRumEventProcessingState gets a reference to the given RUMEventProcessingState and assigns it to the RumEventProcessingState field.
func (o *RUMApplicationCreateAttributes) SetRumEventProcessingState(v RUMEventProcessingState) {
	o.RumEventProcessingState = &v
}

// GetRumProductAnalyticsRetentionState returns the RumProductAnalyticsRetentionState field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetRumProductAnalyticsRetentionState() RUMProductAnalyticsRetentionState {
	if o == nil || o.RumProductAnalyticsRetentionState == nil {
		var ret RUMProductAnalyticsRetentionState
		return ret
	}
	return *o.RumProductAnalyticsRetentionState
}

// GetRumProductAnalyticsRetentionStateOk returns a tuple with the RumProductAnalyticsRetentionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetRumProductAnalyticsRetentionStateOk() (*RUMProductAnalyticsRetentionState, bool) {
	if o == nil || o.RumProductAnalyticsRetentionState == nil {
		return nil, false
	}
	return o.RumProductAnalyticsRetentionState, true
}

// HasRumProductAnalyticsRetentionState returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasRumProductAnalyticsRetentionState() bool {
	return o != nil && o.RumProductAnalyticsRetentionState != nil
}

// SetRumProductAnalyticsRetentionState gets a reference to the given RUMProductAnalyticsRetentionState and assigns it to the RumProductAnalyticsRetentionState field.
func (o *RUMApplicationCreateAttributes) SetRumProductAnalyticsRetentionState(v RUMProductAnalyticsRetentionState) {
	o.RumProductAnalyticsRetentionState = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RUMApplicationCreateAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationCreateAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RUMApplicationCreateAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RUMApplicationCreateAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.RumEventProcessingState != nil {
		toSerialize["rum_event_processing_state"] = o.RumEventProcessingState
	}
	if o.RumProductAnalyticsRetentionState != nil {
		toSerialize["rum_product_analytics_retention_state"] = o.RumProductAnalyticsRetentionState
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMApplicationCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                              *string                            `json:"name"`
		RumEventProcessingState           *RUMEventProcessingState           `json:"rum_event_processing_state,omitempty"`
		RumProductAnalyticsRetentionState *RUMProductAnalyticsRetentionState `json:"rum_product_analytics_retention_state,omitempty"`
		Type                              *string                            `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "rum_event_processing_state", "rum_product_analytics_retention_state", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if all.RumEventProcessingState != nil && !all.RumEventProcessingState.IsValid() {
		hasInvalidField = true
	} else {
		o.RumEventProcessingState = all.RumEventProcessingState
	}
	if all.RumProductAnalyticsRetentionState != nil && !all.RumProductAnalyticsRetentionState.IsValid() {
		hasInvalidField = true
	} else {
		o.RumProductAnalyticsRetentionState = all.RumProductAnalyticsRetentionState
	}
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
