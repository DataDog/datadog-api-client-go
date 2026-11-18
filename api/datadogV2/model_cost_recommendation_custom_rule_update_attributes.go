// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationCustomRuleUpdateAttributes Attributes to change. Attributes left out keep their value. `filters` and `settings` replace the stored list
// and map when sent.
type CostRecommendationCustomRuleUpdateAttributes struct {
	// Whether the rule is applied when recommendations are generated.
	Enabled *bool `json:"enabled,omitempty"`
	// Tag filters that select the resources the rule applies to. An empty list selects every resource.
	Filters []CostRecommendationCustomRuleFilter `json:"filters,omitempty"`
	// The name of the rule.
	Name *string `json:"name,omitempty"`
	// Free-form notes, such as the reason for the customization.
	Notes *string `json:"notes,omitempty"`
	// Setting overrides, keyed by setting name. Settings left out keep the recommender default. The accepted
	// settings and their ranges depend on the recommender.
	Settings map[string]float64 `json:"settings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationCustomRuleUpdateAttributes instantiates a new CostRecommendationCustomRuleUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationCustomRuleUpdateAttributes() *CostRecommendationCustomRuleUpdateAttributes {
	this := CostRecommendationCustomRuleUpdateAttributes{}
	return &this
}

// NewCostRecommendationCustomRuleUpdateAttributesWithDefaults instantiates a new CostRecommendationCustomRuleUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationCustomRuleUpdateAttributesWithDefaults() *CostRecommendationCustomRuleUpdateAttributes {
	this := CostRecommendationCustomRuleUpdateAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CostRecommendationCustomRuleUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetFilters() []CostRecommendationCustomRuleFilter {
	if o == nil || o.Filters == nil {
		var ret []CostRecommendationCustomRuleFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetFiltersOk() (*[]CostRecommendationCustomRuleFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return &o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given []CostRecommendationCustomRuleFilter and assigns it to the Filters field.
func (o *CostRecommendationCustomRuleUpdateAttributes) SetFilters(v []CostRecommendationCustomRuleFilter) {
	o.Filters = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CostRecommendationCustomRuleUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetNotes() string {
	if o == nil || o.Notes == nil {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetNotesOk() (*string, bool) {
	if o == nil || o.Notes == nil {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) HasNotes() bool {
	return o != nil && o.Notes != nil
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *CostRecommendationCustomRuleUpdateAttributes) SetNotes(v string) {
	o.Notes = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetSettings() map[string]float64 {
	if o == nil || o.Settings == nil {
		var ret map[string]float64
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) GetSettingsOk() (*map[string]float64, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return &o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CostRecommendationCustomRuleUpdateAttributes) HasSettings() bool {
	return o != nil && o.Settings != nil
}

// SetSettings gets a reference to the given map[string]float64 and assigns it to the Settings field.
func (o *CostRecommendationCustomRuleUpdateAttributes) SetSettings(v map[string]float64) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationCustomRuleUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationCustomRuleUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled  *bool                                `json:"enabled,omitempty"`
		Filters  []CostRecommendationCustomRuleFilter `json:"filters,omitempty"`
		Name     *string                              `json:"name,omitempty"`
		Notes    *string                              `json:"notes,omitempty"`
		Settings map[string]float64                   `json:"settings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "filters", "name", "notes", "settings"})
	} else {
		return err
	}
	o.Enabled = all.Enabled
	o.Filters = all.Filters
	o.Name = all.Name
	o.Notes = all.Notes
	o.Settings = all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
