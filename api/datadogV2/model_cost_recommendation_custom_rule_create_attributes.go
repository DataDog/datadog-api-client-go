// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationCustomRuleCreateAttributes Attributes of a new custom recommendation rule.
type CostRecommendationCustomRuleCreateAttributes struct {
	// Whether the rule is applied when recommendations are generated. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty"`
	// Tag filters that select the resources the rule applies to. An empty list selects every resource.
	Filters []CostRecommendationCustomRuleFilter `json:"filters,omitempty"`
	// The name of the rule.
	Name string `json:"name"`
	// Free-form notes, such as the reason for the customization.
	Notes *string `json:"notes,omitempty"`
	// The recommender to customize. It must be a recommender that supports custom rules.
	RecommenderId string `json:"recommender_id"`
	// Setting overrides, keyed by setting name. Settings left out keep the recommender default. The accepted
	// settings and their ranges depend on the recommender.
	Settings map[string]float64 `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationCustomRuleCreateAttributes instantiates a new CostRecommendationCustomRuleCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationCustomRuleCreateAttributes(name string, recommenderId string) *CostRecommendationCustomRuleCreateAttributes {
	this := CostRecommendationCustomRuleCreateAttributes{}
	this.Name = name
	this.RecommenderId = recommenderId
	return &this
}

// NewCostRecommendationCustomRuleCreateAttributesWithDefaults instantiates a new CostRecommendationCustomRuleCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationCustomRuleCreateAttributesWithDefaults() *CostRecommendationCustomRuleCreateAttributes {
	this := CostRecommendationCustomRuleCreateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CostRecommendationCustomRuleCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleCreateAttributes) GetFilters() []CostRecommendationCustomRuleFilter {
	if o == nil || o.Filters == nil {
		var ret []CostRecommendationCustomRuleFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) GetFiltersOk() (*[]CostRecommendationCustomRuleFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []CostRecommendationCustomRuleFilter and assigns it to the Filters field.
func (o *CostRecommendationCustomRuleCreateAttributes) SetFilters(v []CostRecommendationCustomRuleFilter) {
	o.Filters = v
}

// GetName returns the Name field value.
func (o *CostRecommendationCustomRuleCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CostRecommendationCustomRuleCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleCreateAttributes) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CostRecommendationCustomRuleCreateAttributes) SetNotes(v string) {
	o.Notes = &v
}

// GetRecommenderId returns the RecommenderId field value.
func (o *CostRecommendationCustomRuleCreateAttributes) GetRecommenderId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RecommenderId
}

// GetRecommenderIdOk returns a tuple with the RecommenderId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) GetRecommenderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommenderId, true
}

// SetRecommenderId sets field value.
func (o *CostRecommendationCustomRuleCreateAttributes) SetRecommenderId(v string) {
	o.RecommenderId = v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleCreateAttributes) GetSettings() map[string]float64 {
	if o == nil || o.Settings == nil {
		var ret map[string]float64
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) GetSettingsOk() (*map[string]float64, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleCreateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]float64 and assigns it to the Settings field.
func (o *CostRecommendationCustomRuleCreateAttributes) SetSettings(v map[string]float64) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationCustomRuleCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["name"] = o.Name
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["recommender_id"] = o.RecommenderId
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationCustomRuleCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       *bool                                `json:"enabled,omitempty"`
		Filters       []CostRecommendationCustomRuleFilter `json:"filters,omitempty"`
		Name          *string                              `json:"name"`
		Notes         *string                              `json:"notes,omitempty"`
		RecommenderId *string                              `json:"recommender_id"`
		Settings      map[string]float64                   `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RecommenderId == nil {
		return fmt.Errorf("required field recommender_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "filters", "name", "notes", "recommender_id", "settings"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.Filters = all.Filters
	o.Name = *all.Name
	o.Notes = all.Notes
	o.RecommenderId = *all.RecommenderId
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
