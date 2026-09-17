// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridResponseAttributes Attributes of a retention grid response, containing the cohort rows and the period columns.
type ProductAnalyticsRetentionGridResponseAttributes struct {
	// The cohorts forming the rows of the grid.
	Cohorts []ProductAnalyticsRetentionGridCohort `json:"cohorts,omitempty"`
	// The entity whose retention was measured.
	RetentionEntity *string `json:"retention_entity,omitempty"`
	// The return periods forming the columns of the grid.
	RetentionPeriods []ProductAnalyticsRetentionPeriod `json:"retention_periods,omitempty"`
	// Unit definitions for the grid values.
	Unit []ProductAnalyticsUnit `json:"unit,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionGridResponseAttributes instantiates a new ProductAnalyticsRetentionGridResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionGridResponseAttributes() *ProductAnalyticsRetentionGridResponseAttributes {
	this := ProductAnalyticsRetentionGridResponseAttributes{}
	return &this
}

// NewProductAnalyticsRetentionGridResponseAttributesWithDefaults instantiates a new ProductAnalyticsRetentionGridResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionGridResponseAttributesWithDefaults() *ProductAnalyticsRetentionGridResponseAttributes {
	this := ProductAnalyticsRetentionGridResponseAttributes{}
	return &this
}

// GetCohorts returns the Cohorts field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetCohorts() []ProductAnalyticsRetentionGridCohort {
	if o == nil || o.Cohorts == nil {
		var ret []ProductAnalyticsRetentionGridCohort
		return ret
	}
	return o.Cohorts
}

// GetCohortsOk returns a tuple with the Cohorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetCohortsOk() (*[]ProductAnalyticsRetentionGridCohort, bool) {
	if o == nil || o.Cohorts == nil {
		return nil, false
	}
	return &o.Cohorts, true
}

// HasCohorts returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) HasCohorts() bool {
	return o != nil && o.Cohorts != nil
}

// SetCohorts gets a reference to the given []ProductAnalyticsRetentionGridCohort and assigns it to the Cohorts field.
func (o *ProductAnalyticsRetentionGridResponseAttributes) SetCohorts(v []ProductAnalyticsRetentionGridCohort) {
	o.Cohorts = v
}

// GetRetentionEntity returns the RetentionEntity field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetRetentionEntity() string {
	if o == nil || o.RetentionEntity == nil {
		var ret string
		return ret
	}
	return *o.RetentionEntity
}

// GetRetentionEntityOk returns a tuple with the RetentionEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetRetentionEntityOk() (*string, bool) {
	if o == nil || o.RetentionEntity == nil {
		return nil, false
	}
	return o.RetentionEntity, true
}

// HasRetentionEntity returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) HasRetentionEntity() bool {
	return o != nil && o.RetentionEntity != nil
}

// SetRetentionEntity gets a reference to the given string and assigns it to the RetentionEntity field.
func (o *ProductAnalyticsRetentionGridResponseAttributes) SetRetentionEntity(v string) {
	o.RetentionEntity = &v
}

// GetRetentionPeriods returns the RetentionPeriods field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetRetentionPeriods() []ProductAnalyticsRetentionPeriod {
	if o == nil || o.RetentionPeriods == nil {
		var ret []ProductAnalyticsRetentionPeriod
		return ret
	}
	return o.RetentionPeriods
}

// GetRetentionPeriodsOk returns a tuple with the RetentionPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetRetentionPeriodsOk() (*[]ProductAnalyticsRetentionPeriod, bool) {
	if o == nil || o.RetentionPeriods == nil {
		return nil, false
	}
	return &o.RetentionPeriods, true
}

// HasRetentionPeriods returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) HasRetentionPeriods() bool {
	return o != nil && o.RetentionPeriods != nil
}

// SetRetentionPeriods gets a reference to the given []ProductAnalyticsRetentionPeriod and assigns it to the RetentionPeriods field.
func (o *ProductAnalyticsRetentionGridResponseAttributes) SetRetentionPeriods(v []ProductAnalyticsRetentionPeriod) {
	o.RetentionPeriods = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetUnit() []ProductAnalyticsUnit {
	if o == nil || o.Unit == nil {
		var ret []ProductAnalyticsUnit
		return ret
	}
	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) GetUnitOk() (*[]ProductAnalyticsUnit, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return &o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridResponseAttributes) HasUnit() bool {
	return o != nil && o.Unit != nil
}

// SetUnit gets a reference to the given []ProductAnalyticsUnit and assigns it to the Unit field.
func (o *ProductAnalyticsRetentionGridResponseAttributes) SetUnit(v []ProductAnalyticsUnit) {
	o.Unit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionGridResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cohorts != nil {
		toSerialize["cohorts"] = o.Cohorts
	}
	if o.RetentionEntity != nil {
		toSerialize["retention_entity"] = o.RetentionEntity
	}
	if o.RetentionPeriods != nil {
		toSerialize["retention_periods"] = o.RetentionPeriods
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
func (o *ProductAnalyticsRetentionGridResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cohorts          []ProductAnalyticsRetentionGridCohort `json:"cohorts,omitempty"`
		RetentionEntity  *string                               `json:"retention_entity,omitempty"`
		RetentionPeriods []ProductAnalyticsRetentionPeriod     `json:"retention_periods,omitempty"`
		Unit             []ProductAnalyticsUnit                `json:"unit,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cohorts", "retention_entity", "retention_periods", "unit"})
	} else {
		return err
	}
	o.Cohorts = all.Cohorts
	o.RetentionEntity = all.RetentionEntity
	o.RetentionPeriods = all.RetentionPeriods
	o.Unit = all.Unit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
