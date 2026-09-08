// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UnitCostRequestAttributes The attributes of a unit cost create or replace request.
type UnitCostRequestAttributes struct {
	// A timeseries object containing `queries` and `formulas` arrays.
	DenominatorQuery UnitCostQueryDefinition `json:"denominator_query"`
	// An optional description of the unit cost. At most 2000 characters.
	Description datadog.NullableString `json:"description,omitempty"`
	// The name of the unit cost. At most 200 characters.
	Name string `json:"name"`
	// A timeseries object containing `queries` and `formulas` arrays.
	NumeratorQuery UnitCostQueryDefinition `json:"numerator_query"`
	// The label describing the denominator unit, for example `user`. At most 100 characters.
	UnitLabel string `json:"unit_label"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUnitCostRequestAttributes instantiates a new UnitCostRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUnitCostRequestAttributes(denominatorQuery UnitCostQueryDefinition, name string, numeratorQuery UnitCostQueryDefinition, unitLabel string) *UnitCostRequestAttributes {
	this := UnitCostRequestAttributes{}
	this.DenominatorQuery = denominatorQuery
	this.Name = name
	this.NumeratorQuery = numeratorQuery
	this.UnitLabel = unitLabel
	return &this
}

// NewUnitCostRequestAttributesWithDefaults instantiates a new UnitCostRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUnitCostRequestAttributesWithDefaults() *UnitCostRequestAttributes {
	this := UnitCostRequestAttributes{}
	return &this
}

// GetDenominatorQuery returns the DenominatorQuery field value.
func (o *UnitCostRequestAttributes) GetDenominatorQuery() UnitCostQueryDefinition {
	if o == nil {
		var ret UnitCostQueryDefinition
		return ret
	}
	return o.DenominatorQuery
}

// GetDenominatorQueryOk returns a tuple with the DenominatorQuery field value
// and a boolean to check if the value has been set.
func (o *UnitCostRequestAttributes) GetDenominatorQueryOk() (*UnitCostQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DenominatorQuery, true
}

// SetDenominatorQuery sets field value.
func (o *UnitCostRequestAttributes) SetDenominatorQuery(v UnitCostQueryDefinition) {
	o.DenominatorQuery = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UnitCostRequestAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UnitCostRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *UnitCostRequestAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *UnitCostRequestAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *UnitCostRequestAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *UnitCostRequestAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value.
func (o *UnitCostRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UnitCostRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UnitCostRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetNumeratorQuery returns the NumeratorQuery field value.
func (o *UnitCostRequestAttributes) GetNumeratorQuery() UnitCostQueryDefinition {
	if o == nil {
		var ret UnitCostQueryDefinition
		return ret
	}
	return o.NumeratorQuery
}

// GetNumeratorQueryOk returns a tuple with the NumeratorQuery field value
// and a boolean to check if the value has been set.
func (o *UnitCostRequestAttributes) GetNumeratorQueryOk() (*UnitCostQueryDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumeratorQuery, true
}

// SetNumeratorQuery sets field value.
func (o *UnitCostRequestAttributes) SetNumeratorQuery(v UnitCostQueryDefinition) {
	o.NumeratorQuery = v
}

// GetUnitLabel returns the UnitLabel field value.
func (o *UnitCostRequestAttributes) GetUnitLabel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UnitLabel
}

// GetUnitLabelOk returns a tuple with the UnitLabel field value
// and a boolean to check if the value has been set.
func (o *UnitCostRequestAttributes) GetUnitLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnitLabel, true
}

// SetUnitLabel sets field value.
func (o *UnitCostRequestAttributes) SetUnitLabel(v string) {
	o.UnitLabel = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UnitCostRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["denominator_query"] = o.DenominatorQuery
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	toSerialize["name"] = o.Name
	toSerialize["numerator_query"] = o.NumeratorQuery
	toSerialize["unit_label"] = o.UnitLabel

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UnitCostRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DenominatorQuery *UnitCostQueryDefinition `json:"denominator_query"`
		Description      datadog.NullableString   `json:"description,omitempty"`
		Name             *string                  `json:"name"`
		NumeratorQuery   *UnitCostQueryDefinition `json:"numerator_query"`
		UnitLabel        *string                  `json:"unit_label"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DenominatorQuery == nil {
		return fmt.Errorf("required field denominator_query missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.NumeratorQuery == nil {
		return fmt.Errorf("required field numerator_query missing")
	}
	if all.UnitLabel == nil {
		return fmt.Errorf("required field unit_label missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"denominator_query", "description", "name", "numerator_query", "unit_label"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DenominatorQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DenominatorQuery = *all.DenominatorQuery
	o.Description = all.Description
	o.Name = *all.Name
	if all.NumeratorQuery.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumeratorQuery = *all.NumeratorQuery
	o.UnitLabel = *all.UnitLabel

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
