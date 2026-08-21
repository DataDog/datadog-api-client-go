// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridCohort One row of the retention grid, holding the results for a single cohort.
type ProductAnalyticsRetentionGridCohort struct {
	// The cells of the row, one per return period.
	Cells []ProductAnalyticsRetentionGridCohortCell `json:"cells,omitempty"`
	// End of the cohort window, in epoch milliseconds.
	CohortEndTime *int64 `json:"cohort_end_time,omitempty"`
	// Zero-based index of the cohort in the grid.
	CohortIndex *int64 `json:"cohort_index,omitempty"`
	// Number of entities in the cohort.
	CohortSize *int64 `json:"cohort_size,omitempty"`
	// Start of the cohort window, in epoch milliseconds.
	CohortStartTime *int64 `json:"cohort_start_time,omitempty"`
	// The group-by facet values that identify this row.
	GroupTags []string `json:"group_tags,omitempty"`
	// Label identifying the cohort, such as the week it started.
	Name *string `json:"name,omitempty"`
	// Whether the row holds one cohort's own numbers, or the weighted roll-up across every cohort.
	Type *ProductAnalyticsRetentionGridCohortType `json:"type,omitempty"`
	// Unit definitions for the cell values.
	Unit []ProductAnalyticsUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionGridCohort instantiates a new ProductAnalyticsRetentionGridCohort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionGridCohort() *ProductAnalyticsRetentionGridCohort {
	this := ProductAnalyticsRetentionGridCohort{}
	return &this
}

// NewProductAnalyticsRetentionGridCohortWithDefaults instantiates a new ProductAnalyticsRetentionGridCohort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionGridCohortWithDefaults() *ProductAnalyticsRetentionGridCohort {
	this := ProductAnalyticsRetentionGridCohort{}
	return &this
}

// GetCells returns the Cells field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetCells() []ProductAnalyticsRetentionGridCohortCell {
	if o == nil || o.Cells == nil {
		var ret []ProductAnalyticsRetentionGridCohortCell
		return ret
	}
	return o.Cells
}

// GetCellsOk returns a tuple with the Cells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetCellsOk() (*[]ProductAnalyticsRetentionGridCohortCell, bool) {
	if o == nil || o.Cells == nil {
		return nil, false
	}
	return &o.Cells, true
}

// HasCells returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasCells() bool {
	return o != nil && o.Cells != nil
}

// SetCells gets a reference to the given []ProductAnalyticsRetentionGridCohortCell and assigns it to the Cells field.
func (o *ProductAnalyticsRetentionGridCohort) SetCells(v []ProductAnalyticsRetentionGridCohortCell) {
	o.Cells = v
}

// GetCohortEndTime returns the CohortEndTime field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortEndTime() int64 {
	if o == nil || o.CohortEndTime == nil {
		var ret int64
		return ret
	}
	return *o.CohortEndTime
}

// GetCohortEndTimeOk returns a tuple with the CohortEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortEndTimeOk() (*int64, bool) {
	if o == nil || o.CohortEndTime == nil {
		return nil, false
	}
	return o.CohortEndTime, true
}

// HasCohortEndTime returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasCohortEndTime() bool {
	return o != nil && o.CohortEndTime != nil
}

// SetCohortEndTime gets a reference to the given int64 and assigns it to the CohortEndTime field.
func (o *ProductAnalyticsRetentionGridCohort) SetCohortEndTime(v int64) {
	o.CohortEndTime = &v
}

// GetCohortIndex returns the CohortIndex field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortIndex() int64 {
	if o == nil || o.CohortIndex == nil {
		var ret int64
		return ret
	}
	return *o.CohortIndex
}

// GetCohortIndexOk returns a tuple with the CohortIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortIndexOk() (*int64, bool) {
	if o == nil || o.CohortIndex == nil {
		return nil, false
	}
	return o.CohortIndex, true
}

// HasCohortIndex returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasCohortIndex() bool {
	return o != nil && o.CohortIndex != nil
}

// SetCohortIndex gets a reference to the given int64 and assigns it to the CohortIndex field.
func (o *ProductAnalyticsRetentionGridCohort) SetCohortIndex(v int64) {
	o.CohortIndex = &v
}

// GetCohortSize returns the CohortSize field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortSize() int64 {
	if o == nil || o.CohortSize == nil {
		var ret int64
		return ret
	}
	return *o.CohortSize
}

// GetCohortSizeOk returns a tuple with the CohortSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortSizeOk() (*int64, bool) {
	if o == nil || o.CohortSize == nil {
		return nil, false
	}
	return o.CohortSize, true
}

// HasCohortSize returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasCohortSize() bool {
	return o != nil && o.CohortSize != nil
}

// SetCohortSize gets a reference to the given int64 and assigns it to the CohortSize field.
func (o *ProductAnalyticsRetentionGridCohort) SetCohortSize(v int64) {
	o.CohortSize = &v
}

// GetCohortStartTime returns the CohortStartTime field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortStartTime() int64 {
	if o == nil || o.CohortStartTime == nil {
		var ret int64
		return ret
	}
	return *o.CohortStartTime
}

// GetCohortStartTimeOk returns a tuple with the CohortStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetCohortStartTimeOk() (*int64, bool) {
	if o == nil || o.CohortStartTime == nil {
		return nil, false
	}
	return o.CohortStartTime, true
}

// HasCohortStartTime returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasCohortStartTime() bool {
	return o != nil && o.CohortStartTime != nil
}

// SetCohortStartTime gets a reference to the given int64 and assigns it to the CohortStartTime field.
func (o *ProductAnalyticsRetentionGridCohort) SetCohortStartTime(v int64) {
	o.CohortStartTime = &v
}

// GetGroupTags returns the GroupTags field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetGroupTags() []string {
	if o == nil || o.GroupTags == nil {
		var ret []string
		return ret
	}
	return o.GroupTags
}

// GetGroupTagsOk returns a tuple with the GroupTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetGroupTagsOk() (*[]string, bool) {
	if o == nil || o.GroupTags == nil {
		return nil, false
	}
	return &o.GroupTags, true
}

// HasGroupTags returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasGroupTags() bool {
	return o != nil && o.GroupTags != nil
}

// SetGroupTags gets a reference to the given []string and assigns it to the GroupTags field.
func (o *ProductAnalyticsRetentionGridCohort) SetGroupTags(v []string) {
	o.GroupTags = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsRetentionGridCohort) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetType() ProductAnalyticsRetentionGridCohortType {
	if o == nil || o.Type == nil {
		var ret ProductAnalyticsRetentionGridCohortType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetTypeOk() (*ProductAnalyticsRetentionGridCohortType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given ProductAnalyticsRetentionGridCohortType and assigns it to the Type field.
func (o *ProductAnalyticsRetentionGridCohort) SetType(v ProductAnalyticsRetentionGridCohortType) {
	o.Type = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridCohort) GetUnit() []ProductAnalyticsUnit {
	if o == nil || o.Unit == nil {
		var ret []ProductAnalyticsUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridCohort) GetUnitOk() (*[]ProductAnalyticsUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return &o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridCohort) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given []ProductAnalyticsUnit and assigns it to the Unit field.
func (o *ProductAnalyticsRetentionGridCohort) SetUnit(v []ProductAnalyticsUnit) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionGridCohort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cells != nil {
		toSerialize["cells"] = o.Cells
	}
	if o.CohortEndTime != nil {
		toSerialize["cohort_end_time"] = o.CohortEndTime
	}
	if o.CohortIndex != nil {
		toSerialize["cohort_index"] = o.CohortIndex
	}
	if o.CohortSize != nil {
		toSerialize["cohort_size"] = o.CohortSize
	}
	if o.CohortStartTime != nil {
		toSerialize["cohort_start_time"] = o.CohortStartTime
	}
	if o.GroupTags != nil {
		toSerialize["group_tags"] = o.GroupTags
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionGridCohort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cells           []ProductAnalyticsRetentionGridCohortCell `json:"cells,omitempty"`
		CohortEndTime   *int64                                    `json:"cohort_end_time,omitempty"`
		CohortIndex     *int64                                    `json:"cohort_index,omitempty"`
		CohortSize      *int64                                    `json:"cohort_size,omitempty"`
		CohortStartTime *int64                                    `json:"cohort_start_time,omitempty"`
		GroupTags       []string                                  `json:"group_tags,omitempty"`
		Name            *string                                   `json:"name,omitempty"`
		Type            *ProductAnalyticsRetentionGridCohortType  `json:"type,omitempty"`
		Unit            []ProductAnalyticsUnit                    `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cells", "cohort_end_time", "cohort_index", "cohort_size", "cohort_start_time", "group_tags", "name", "type", "unit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cells = all.Cells
	o.CohortEndTime = all.CohortEndTime
	o.CohortIndex = all.CohortIndex
	o.CohortSize = all.CohortSize
	o.CohortStartTime = all.CohortStartTime
	o.GroupTags = all.GroupTags
	o.Name = all.Name
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
