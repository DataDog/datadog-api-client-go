// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DefaultRuleResponseAttributes Default rule attributes.
type DefaultRuleResponseAttributes struct {
	// The description of the default rule.
	Description *string `json:"description,omitempty"`
	// The maturity level of the rule.
	Level *int64 `json:"level,omitempty"`
	// The name of the default rule.
	Name string `json:"name"`
	// Required scope for the rule.
	ScopeRequired *string `json:"scope_required,omitempty"`
	// The description of the scorecard.
	ScorecardDescription *string `json:"scorecard_description,omitempty"`
	// The scorecard this rule belongs to.
	ScorecardName string `json:"scorecard_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDefaultRuleResponseAttributes instantiates a new DefaultRuleResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDefaultRuleResponseAttributes(name string, scorecardName string) *DefaultRuleResponseAttributes {
	this := DefaultRuleResponseAttributes{}
	this.Name = name
	this.ScorecardName = scorecardName
	return &this
}

// NewDefaultRuleResponseAttributesWithDefaults instantiates a new DefaultRuleResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDefaultRuleResponseAttributesWithDefaults() *DefaultRuleResponseAttributes {
	this := DefaultRuleResponseAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DefaultRuleResponseAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultRuleResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DefaultRuleResponseAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DefaultRuleResponseAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *DefaultRuleResponseAttributes) GetLevel() int64 {
	if o == nil || o.Level == nil {
		var ret int64
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultRuleResponseAttributes) GetLevelOk() (*int64, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *DefaultRuleResponseAttributes) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given int64 and assigns it to the Level field.
func (o *DefaultRuleResponseAttributes) SetLevel(v int64) {
	o.Level = &v
}

// GetName returns the Name field value.
func (o *DefaultRuleResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DefaultRuleResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DefaultRuleResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetScopeRequired returns the ScopeRequired field value if set, zero value otherwise.
func (o *DefaultRuleResponseAttributes) GetScopeRequired() string {
	if o == nil || o.ScopeRequired == nil {
		var ret string
		return ret
	}
	return *o.ScopeRequired
}

// GetScopeRequiredOk returns a tuple with the ScopeRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultRuleResponseAttributes) GetScopeRequiredOk() (*string, bool) {
	if o == nil || o.ScopeRequired == nil {
		return nil, false
	}
	return o.ScopeRequired, true
}

// HasScopeRequired returns a boolean if a field has been set.
func (o *DefaultRuleResponseAttributes) HasScopeRequired() bool {
	return o != nil && o.ScopeRequired != nil
}

// SetScopeRequired gets a reference to the given string and assigns it to the ScopeRequired field.
func (o *DefaultRuleResponseAttributes) SetScopeRequired(v string) {
	o.ScopeRequired = &v
}

// GetScorecardDescription returns the ScorecardDescription field value if set, zero value otherwise.
func (o *DefaultRuleResponseAttributes) GetScorecardDescription() string {
	if o == nil || o.ScorecardDescription == nil {
		var ret string
		return ret
	}
	return *o.ScorecardDescription
}

// GetScorecardDescriptionOk returns a tuple with the ScorecardDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultRuleResponseAttributes) GetScorecardDescriptionOk() (*string, bool) {
	if o == nil || o.ScorecardDescription == nil {
		return nil, false
	}
	return o.ScorecardDescription, true
}

// HasScorecardDescription returns a boolean if a field has been set.
func (o *DefaultRuleResponseAttributes) HasScorecardDescription() bool {
	return o != nil && o.ScorecardDescription != nil
}

// SetScorecardDescription gets a reference to the given string and assigns it to the ScorecardDescription field.
func (o *DefaultRuleResponseAttributes) SetScorecardDescription(v string) {
	o.ScorecardDescription = &v
}

// GetScorecardName returns the ScorecardName field value.
func (o *DefaultRuleResponseAttributes) GetScorecardName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ScorecardName
}

// GetScorecardNameOk returns a tuple with the ScorecardName field value
// and a boolean to check if the value has been set.
func (o *DefaultRuleResponseAttributes) GetScorecardNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScorecardName, true
}

// SetScorecardName sets field value.
func (o *DefaultRuleResponseAttributes) SetScorecardName(v string) {
	o.ScorecardName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DefaultRuleResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	toSerialize["name"] = o.Name
	if o.ScopeRequired != nil {
		toSerialize["scope_required"] = o.ScopeRequired
	}
	if o.ScorecardDescription != nil {
		toSerialize["scorecard_description"] = o.ScorecardDescription
	}
	toSerialize["scorecard_name"] = o.ScorecardName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DefaultRuleResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description          *string `json:"description,omitempty"`
		Level                *int64  `json:"level,omitempty"`
		Name                 *string `json:"name"`
		ScopeRequired        *string `json:"scope_required,omitempty"`
		ScorecardDescription *string `json:"scorecard_description,omitempty"`
		ScorecardName        *string `json:"scorecard_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ScorecardName == nil {
		return fmt.Errorf("required field scorecard_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "level", "name", "scope_required", "scorecard_description", "scorecard_name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Level = all.Level
	o.Name = *all.Name
	o.ScopeRequired = all.ScopeRequired
	o.ScorecardDescription = all.ScorecardDescription
	o.ScorecardName = *all.ScorecardName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
