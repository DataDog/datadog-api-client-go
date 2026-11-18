// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostRecommendationCustomRuleAttributes Attributes of a custom recommendation rule.
type CostRecommendationCustomRuleAttributes struct {
	// The cloud provider of the recommender, such as `aws`, `azure`, or `gcp`.
	CloudProvider string `json:"cloud_provider"`
	// When the rule was created.
	CreatedAt time.Time `json:"created_at"`
	// The UUID of the user who created the rule.
	CreatedBy string `json:"created_by"`
	// Whether the rule is applied when recommendations are generated.
	Enabled bool `json:"enabled"`
	// Tag filters that select the resources the rule applies to. An empty list selects every resource.
	Filters []CostRecommendationCustomRuleFilter `json:"filters"`
	// When the rule was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// The UUID of the user who last modified the rule.
	ModifiedBy string `json:"modified_by"`
	// The name of the rule.
	Name string `json:"name"`
	// Free-form notes, such as the reason for the customization.
	Notes string `json:"notes"`
	// The order in which the rules of a recommender are evaluated, lowest first. It is assigned when the rule is created.
	Priority int64 `json:"priority"`
	// The recommender the rule customizes.
	RecommenderId string `json:"recommender_id"`
	// Setting overrides, keyed by setting name. Settings left out keep the recommender default. The accepted
	// settings and their ranges depend on the recommender.
	Settings map[string]float64 `json:"settings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostRecommendationCustomRuleAttributes instantiates a new CostRecommendationCustomRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostRecommendationCustomRuleAttributes(cloudProvider string, createdAt time.Time, createdBy string, enabled bool, filters []CostRecommendationCustomRuleFilter, modifiedAt time.Time, modifiedBy string, name string, notes string, priority int64, recommenderId string, settings map[string]float64) *CostRecommendationCustomRuleAttributes {
	this := CostRecommendationCustomRuleAttributes{}
	this.CloudProvider = cloudProvider
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Enabled = enabled
	this.Filters = filters
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.Notes = notes
	this.Priority = priority
	this.RecommenderId = recommenderId
	this.Settings = settings
	return &this
}

// NewCostRecommendationCustomRuleAttributesWithDefaults instantiates a new CostRecommendationCustomRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostRecommendationCustomRuleAttributesWithDefaults() *CostRecommendationCustomRuleAttributes {
	this := CostRecommendationCustomRuleAttributes{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value.
func (o *CostRecommendationCustomRuleAttributes) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CostRecommendationCustomRuleAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CostRecommendationCustomRuleAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetEnabled returns the Enabled field value.
func (o *CostRecommendationCustomRuleAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFilters returns the Filters field value.
func (o *CostRecommendationCustomRuleAttributes) GetFilters() []CostRecommendationCustomRuleFilter {
	if o == nil {
		var ret []CostRecommendationCustomRuleFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetFiltersOk() (*[]CostRecommendationCustomRuleFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetFilters(v []CostRecommendationCustomRuleFilter) {
	o.Filters = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *CostRecommendationCustomRuleAttributes) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *CostRecommendationCustomRuleAttributes) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *CostRecommendationCustomRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetNotes returns the Notes field value.
func (o *CostRecommendationCustomRuleAttributes) GetNotes() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notes, true
}

// SetNotes sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetNotes(v string) {
	o.Notes = v
}

// GetPriority returns the Priority field value.
func (o *CostRecommendationCustomRuleAttributes) GetPriority() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetPriorityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetPriority(v int64) {
	o.Priority = v
}

// GetRecommenderId returns the RecommenderId field value.
func (o *CostRecommendationCustomRuleAttributes) GetRecommenderId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RecommenderId
}

// GetRecommenderIdOk returns a tuple with the RecommenderId field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetRecommenderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommenderId, true
}

// SetRecommenderId sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetRecommenderId(v string) {
	o.RecommenderId = v
}

// GetSettings returns the Settings field value.
func (o *CostRecommendationCustomRuleAttributes) GetSettings() map[string]float64 {
	if o == nil {
		var ret map[string]float64
		return ret
	}
	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *CostRecommendationCustomRuleAttributes) GetSettingsOk() (*map[string]float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value.
func (o *CostRecommendationCustomRuleAttributes) SetSettings(v map[string]float64) {
	o.Settings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostRecommendationCustomRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cloud_provider"] = o.CloudProvider
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["enabled"] = o.Enabled
	toSerialize["filters"] = o.Filters
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["notes"] = o.Notes
	toSerialize["priority"] = o.Priority
	toSerialize["recommender_id"] = o.RecommenderId
	toSerialize["settings"] = o.Settings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostRecommendationCustomRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudProvider *string                               `json:"cloud_provider"`
		CreatedAt     *time.Time                            `json:"created_at"`
		CreatedBy     *string                               `json:"created_by"`
		Enabled       *bool                                 `json:"enabled"`
		Filters       *[]CostRecommendationCustomRuleFilter `json:"filters"`
		ModifiedAt    *time.Time                            `json:"modified_at"`
		ModifiedBy    *string                               `json:"modified_by"`
		Name          *string                               `json:"name"`
		Notes         *string                               `json:"notes"`
		Priority      *int64                                `json:"priority"`
		RecommenderId *string                               `json:"recommender_id"`
		Settings      *map[string]float64                   `json:"settings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CloudProvider == nil {
		return fmt.Errorf("required field cloud_provider missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Notes == nil {
		return fmt.Errorf("required field notes missing")
	}
	if all.Priority == nil {
		return fmt.Errorf("required field priority missing")
	}
	if all.RecommenderId == nil {
		return fmt.Errorf("required field recommender_id missing")
	}
	if all.Settings == nil {
		return fmt.Errorf("required field settings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud_provider", "created_at", "created_by", "enabled", "filters", "modified_at", "modified_by", "name", "notes", "priority", "recommender_id", "settings"})
	} else {
		return err
	}
	o.CloudProvider = *all.CloudProvider
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Enabled = *all.Enabled
	o.Filters = *all.Filters
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	o.Notes = *all.Notes
	o.Priority = *all.Priority
	o.RecommenderId = *all.RecommenderId
	o.Settings = *all.Settings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
