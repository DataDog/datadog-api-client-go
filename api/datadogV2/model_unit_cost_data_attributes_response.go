// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UnitCostDataAttributesResponse The attributes of a unit cost.
type UnitCostDataAttributesResponse struct {
	// The time the unit cost was created.
	CreatedAt time.Time `json:"created_at"`
	// The UUID of the user who created the unit cost.
	CreatedBy uuid.UUID `json:"created_by"`
	// A timeseries object containing `queries` and `formulas` arrays.
	DenominatorQuery UnitCostQueryDefinition `json:"denominator_query"`
	// The data source of the denominator queries, or `multisource` when the denominator
	// queries span more than one data source.
	DenominatorType string `json:"denominator_type"`
	// The description of the unit cost. Omitted when the unit cost has no description.
	Description datadog.NullableString `json:"description,omitempty"`
	// The name of the unit cost.
	Name string `json:"name"`
	// A timeseries object containing `queries` and `formulas` arrays.
	NumeratorQuery UnitCostQueryDefinition `json:"numerator_query"`
	// The ID of the organization the unit cost belongs to.
	OrgId int64 `json:"org_id"`
	// The label describing the denominator unit.
	UnitLabel string `json:"unit_label"`
	// The time the unit cost was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The UUID of the user who last updated the unit cost.
	UpdatedBy uuid.UUID `json:"updated_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUnitCostDataAttributesResponse instantiates a new UnitCostDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUnitCostDataAttributesResponse(createdAt time.Time, createdBy uuid.UUID, denominatorQuery UnitCostQueryDefinition, denominatorType string, name string, numeratorQuery UnitCostQueryDefinition, orgId int64, unitLabel string, updatedAt time.Time, updatedBy uuid.UUID) *UnitCostDataAttributesResponse {
	this := UnitCostDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.DenominatorQuery = denominatorQuery
	this.DenominatorType = denominatorType
	this.Name = name
	this.NumeratorQuery = numeratorQuery
	this.OrgId = orgId
	this.UnitLabel = unitLabel
	this.UpdatedAt = updatedAt
	this.UpdatedBy = updatedBy
	return &this
}

// NewUnitCostDataAttributesResponseWithDefaults instantiates a new UnitCostDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUnitCostDataAttributesResponseWithDefaults() *UnitCostDataAttributesResponse {
	this := UnitCostDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *UnitCostDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *UnitCostDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *UnitCostDataAttributesResponse) GetCreatedBy() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetCreatedByOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *UnitCostDataAttributesResponse) SetCreatedBy(v uuid.UUID) {
	o.CreatedBy = v
}

// GetDenominatorQuery returns the DenominatorQuery field value.
func (o *UnitCostDataAttributesResponse) GetDenominatorQuery() UnitCostQueryDefinition {
	if o == nil {
		var ret UnitCostQueryDefinition
		return ret
	}
	return o.DenominatorQuery
}

// GetDenominatorQueryOk returns a tuple with the DenominatorQuery field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetDenominatorQueryOk() (*UnitCostQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DenominatorQuery, true
}

// SetDenominatorQuery sets field value.
func (o *UnitCostDataAttributesResponse) SetDenominatorQuery(v UnitCostQueryDefinition) {
	o.DenominatorQuery = v
}

// GetDenominatorType returns the DenominatorType field value.
func (o *UnitCostDataAttributesResponse) GetDenominatorType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DenominatorType
}

// GetDenominatorTypeOk returns a tuple with the DenominatorType field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetDenominatorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DenominatorType, true
}

// SetDenominatorType sets field value.
func (o *UnitCostDataAttributesResponse) SetDenominatorType(v string) {
	o.DenominatorType = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnitCostDataAttributesResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UnitCostDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UnitCostDataAttributesResponse) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *UnitCostDataAttributesResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *UnitCostDataAttributesResponse) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *UnitCostDataAttributesResponse) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value.
func (o *UnitCostDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UnitCostDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetNumeratorQuery returns the NumeratorQuery field value.
func (o *UnitCostDataAttributesResponse) GetNumeratorQuery() UnitCostQueryDefinition {
	if o == nil {
		var ret UnitCostQueryDefinition
		return ret
	}
	return o.NumeratorQuery
}

// GetNumeratorQueryOk returns a tuple with the NumeratorQuery field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetNumeratorQueryOk() (*UnitCostQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeratorQuery, true
}

// SetNumeratorQuery sets field value.
func (o *UnitCostDataAttributesResponse) SetNumeratorQuery(v UnitCostQueryDefinition) {
	o.NumeratorQuery = v
}

// GetOrgId returns the OrgId field value.
func (o *UnitCostDataAttributesResponse) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *UnitCostDataAttributesResponse) SetOrgId(v int64) {
	o.OrgId = v
}

// GetUnitLabel returns the UnitLabel field value.
func (o *UnitCostDataAttributesResponse) GetUnitLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UnitLabel
}

// GetUnitLabelOk returns a tuple with the UnitLabel field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetUnitLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitLabel, true
}

// SetUnitLabel sets field value.
func (o *UnitCostDataAttributesResponse) SetUnitLabel(v string) {
	o.UnitLabel = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *UnitCostDataAttributesResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *UnitCostDataAttributesResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value.
func (o *UnitCostDataAttributesResponse) GetUpdatedBy() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value
// and a boolean to check if the value has been set.
func (o *UnitCostDataAttributesResponse) GetUpdatedByOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedBy, true
}

// SetUpdatedBy sets field value.
func (o *UnitCostDataAttributesResponse) SetUpdatedBy(v uuid.UUID) {
	o.UpdatedBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UnitCostDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["denominator_query"] = o.DenominatorQuery
	toSerialize["denominator_type"] = o.DenominatorType
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["numerator_query"] = o.NumeratorQuery
	toSerialize["org_id"] = o.OrgId
	toSerialize["unit_label"] = o.UnitLabel
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["updated_by"] = o.UpdatedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UnitCostDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time               `json:"created_at"`
		CreatedBy        *uuid.UUID               `json:"created_by"`
		DenominatorQuery *UnitCostQueryDefinition `json:"denominator_query"`
		DenominatorType  *string                  `json:"denominator_type"`
		Description      datadog.NullableString   `json:"description,omitempty"`
		Name             *string                  `json:"name"`
		NumeratorQuery   *UnitCostQueryDefinition `json:"numerator_query"`
		OrgId            *int64                   `json:"org_id"`
		UnitLabel        *string                  `json:"unit_label"`
		UpdatedAt        *time.Time               `json:"updated_at"`
		UpdatedBy        *uuid.UUID               `json:"updated_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.DenominatorQuery == nil {
		return fmt.Errorf("required field denominator_query missing")
	}
	if all.DenominatorType == nil {
		return fmt.Errorf("required field denominator_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.NumeratorQuery == nil {
		return fmt.Errorf("required field numerator_query missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.UnitLabel == nil {
		return fmt.Errorf("required field unit_label missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedBy == nil {
		return fmt.Errorf("required field updated_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "denominator_query", "denominator_type", "description", "name", "numerator_query", "org_id", "unit_label", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	if all.DenominatorQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DenominatorQuery = *all.DenominatorQuery
	o.DenominatorType = *all.DenominatorType
	o.Description = all.Description
	o.Name = *all.Name
	if all.NumeratorQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumeratorQuery = *all.NumeratorQuery
	o.OrgId = *all.OrgId
	o.UnitLabel = *all.UnitLabel
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedBy = *all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
