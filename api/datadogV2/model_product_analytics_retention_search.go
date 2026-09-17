// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionSearch Defines the cohort and return criteria that make up a retention query.
type ProductAnalyticsRetentionSearch struct {
	// Defines the event that places an entity into a cohort, and how cohorts are bucketed over time.
	CohortCriteria ProductAnalyticsRetentionCohortCriteria `json:"cohort_criteria"`
	// Filters narrowing the events considered by a retention query.
	Filters *ProductAnalyticsRetentionFilters `json:"filters,omitempty"`
	// The entity whose retention is measured.
	RetentionEntity ProductAnalyticsRetentionEntity `json:"retention_entity"`
	// When an entity counts as having returned. Use `conversion_on` to count only entities that
	// returned during the period itself, or `conversion_on_or_after` to also count later returns.
	ReturnCondition ProductAnalyticsRetentionReturnCondition `json:"return_condition"`
	// Defines the event that counts as a return, and the window in which it must occur.
	ReturnCriteria *ProductAnalyticsRetentionReturnCriteria `json:"return_criteria,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionSearch instantiates a new ProductAnalyticsRetentionSearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionSearch(cohortCriteria ProductAnalyticsRetentionCohortCriteria, retentionEntity ProductAnalyticsRetentionEntity, returnCondition ProductAnalyticsRetentionReturnCondition) *ProductAnalyticsRetentionSearch {
	this := ProductAnalyticsRetentionSearch{}
	this.CohortCriteria = cohortCriteria
	this.RetentionEntity = retentionEntity
	this.ReturnCondition = returnCondition
	return &this
}

// NewProductAnalyticsRetentionSearchWithDefaults instantiates a new ProductAnalyticsRetentionSearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionSearchWithDefaults() *ProductAnalyticsRetentionSearch {
	this := ProductAnalyticsRetentionSearch{}
	return &this
}

// GetCohortCriteria returns the CohortCriteria field value.
func (o *ProductAnalyticsRetentionSearch) GetCohortCriteria() ProductAnalyticsRetentionCohortCriteria {
	if o == nil {
		var ret ProductAnalyticsRetentionCohortCriteria
		return ret
	}
	return o.CohortCriteria
}

// GetCohortCriteriaOk returns a tuple with the CohortCriteria field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionSearch) GetCohortCriteriaOk() (*ProductAnalyticsRetentionCohortCriteria, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CohortCriteria, true
}

// SetCohortCriteria sets field value.
func (o *ProductAnalyticsRetentionSearch) SetCohortCriteria(v ProductAnalyticsRetentionCohortCriteria) {
	o.CohortCriteria = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionSearch) GetFilters() ProductAnalyticsRetentionFilters {
	if o == nil || o.Filters == nil {
		var ret ProductAnalyticsRetentionFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionSearch) GetFiltersOk() (*ProductAnalyticsRetentionFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionSearch) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given ProductAnalyticsRetentionFilters and assigns it to the Filters field.
func (o *ProductAnalyticsRetentionSearch) SetFilters(v ProductAnalyticsRetentionFilters) {
	o.Filters = &v
}

// GetRetentionEntity returns the RetentionEntity field value.
func (o *ProductAnalyticsRetentionSearch) GetRetentionEntity() ProductAnalyticsRetentionEntity {
	if o == nil {
		var ret ProductAnalyticsRetentionEntity
		return ret
	}
	return o.RetentionEntity
}

// GetRetentionEntityOk returns a tuple with the RetentionEntity field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionSearch) GetRetentionEntityOk() (*ProductAnalyticsRetentionEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionEntity, true
}

// SetRetentionEntity sets field value.
func (o *ProductAnalyticsRetentionSearch) SetRetentionEntity(v ProductAnalyticsRetentionEntity) {
	o.RetentionEntity = v
}

// GetReturnCondition returns the ReturnCondition field value.
func (o *ProductAnalyticsRetentionSearch) GetReturnCondition() ProductAnalyticsRetentionReturnCondition {
	if o == nil {
		var ret ProductAnalyticsRetentionReturnCondition
		return ret
	}
	return o.ReturnCondition
}

// GetReturnConditionOk returns a tuple with the ReturnCondition field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionSearch) GetReturnConditionOk() (*ProductAnalyticsRetentionReturnCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReturnCondition, true
}

// SetReturnCondition sets field value.
func (o *ProductAnalyticsRetentionSearch) SetReturnCondition(v ProductAnalyticsRetentionReturnCondition) {
	o.ReturnCondition = v
}

// GetReturnCriteria returns the ReturnCriteria field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionSearch) GetReturnCriteria() ProductAnalyticsRetentionReturnCriteria {
	if o == nil || o.ReturnCriteria == nil {
		var ret ProductAnalyticsRetentionReturnCriteria
		return ret
	}
	return *o.ReturnCriteria
}

// GetReturnCriteriaOk returns a tuple with the ReturnCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionSearch) GetReturnCriteriaOk() (*ProductAnalyticsRetentionReturnCriteria, bool) {
	if o == nil || o.ReturnCriteria == nil {
		return nil, false
	}
	return o.ReturnCriteria, true
}

// HasReturnCriteria returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionSearch) HasReturnCriteria() bool {
	return o != nil && o.ReturnCriteria != nil
}

// SetReturnCriteria gets a reference to the given ProductAnalyticsRetentionReturnCriteria and assigns it to the ReturnCriteria field.
func (o *ProductAnalyticsRetentionSearch) SetReturnCriteria(v ProductAnalyticsRetentionReturnCriteria) {
	o.ReturnCriteria = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cohort_criteria"] = o.CohortCriteria
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	toSerialize["retention_entity"] = o.RetentionEntity
	toSerialize["return_condition"] = o.ReturnCondition
	if o.ReturnCriteria != nil {
		toSerialize["return_criteria"] = o.ReturnCriteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionSearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CohortCriteria  *ProductAnalyticsRetentionCohortCriteria  `json:"cohort_criteria"`
		Filters         *ProductAnalyticsRetentionFilters         `json:"filters,omitempty"`
		RetentionEntity *ProductAnalyticsRetentionEntity          `json:"retention_entity"`
		ReturnCondition *ProductAnalyticsRetentionReturnCondition `json:"return_condition"`
		ReturnCriteria  *ProductAnalyticsRetentionReturnCriteria  `json:"return_criteria,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CohortCriteria == nil {
		return fmt.Errorf("required field cohort_criteria missing")
	}
	if all.RetentionEntity == nil {
		return fmt.Errorf("required field retention_entity missing")
	}
	if all.ReturnCondition == nil {
		return fmt.Errorf("required field return_condition missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cohort_criteria", "filters", "retention_entity", "return_condition", "return_criteria"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CohortCriteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CohortCriteria = *all.CohortCriteria
	if all.Filters != nil && all.Filters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filters = all.Filters
	if !all.RetentionEntity.IsValid() {
		hasInvalidField = true
	} else {
		o.RetentionEntity = *all.RetentionEntity
	}
	if !all.ReturnCondition.IsValid() {
		hasInvalidField = true
	} else {
		o.ReturnCondition = *all.ReturnCondition
	}
	if all.ReturnCriteria != nil && all.ReturnCriteria.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ReturnCriteria = all.ReturnCriteria

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
