// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FeatureFlagStalenessRecommendedAction An action suggested for a feature flag based on its staleness state.
type FeatureFlagStalenessRecommendedAction struct {
	// The action to consider. Values include `remove_from_code`, `archive_flag`, `mark_as_permanent`, `snooze`, and `check_sdk_config`.
	Action *string `json:"action,omitempty"`
	// An explanation of the suggested action.
	Description *string `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFeatureFlagStalenessRecommendedAction instantiates a new FeatureFlagStalenessRecommendedAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFeatureFlagStalenessRecommendedAction() *FeatureFlagStalenessRecommendedAction {
	this := FeatureFlagStalenessRecommendedAction{}
	return &this
}

// NewFeatureFlagStalenessRecommendedActionWithDefaults instantiates a new FeatureFlagStalenessRecommendedAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFeatureFlagStalenessRecommendedActionWithDefaults() *FeatureFlagStalenessRecommendedAction {
	this := FeatureFlagStalenessRecommendedAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *FeatureFlagStalenessRecommendedAction) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagStalenessRecommendedAction) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *FeatureFlagStalenessRecommendedAction) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *FeatureFlagStalenessRecommendedAction) SetAction(v string) {
	o.Action = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FeatureFlagStalenessRecommendedAction) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureFlagStalenessRecommendedAction) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FeatureFlagStalenessRecommendedAction) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FeatureFlagStalenessRecommendedAction) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FeatureFlagStalenessRecommendedAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FeatureFlagStalenessRecommendedAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action      *string `json:"action,omitempty"`
		Description *string `json:"description,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "description"})
	} else {
		return err
	}
	o.Action = all.Action
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
