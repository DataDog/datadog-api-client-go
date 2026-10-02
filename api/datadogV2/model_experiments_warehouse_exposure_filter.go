// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsWarehouseExposureFilter A comparison that selects warehouse exposure data by a property.
type ExperimentsWarehouseExposureFilter struct {
	// Comparison applied by the warehouse entry-point filter.
	Operation ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation `json:"operation"`
	// Warehouse property UUID.
	PropertyId string `json:"property_id"`
	// Ordered values used by the filter.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsWarehouseExposureFilter instantiates a new ExperimentsWarehouseExposureFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsWarehouseExposureFilter(operation ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation, propertyId string, values []string) *ExperimentsWarehouseExposureFilter {
	this := ExperimentsWarehouseExposureFilter{}
	this.Operation = operation
	this.PropertyId = propertyId
	this.Values = values
	return &this
}

// NewExperimentsWarehouseExposureFilterWithDefaults instantiates a new ExperimentsWarehouseExposureFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsWarehouseExposureFilterWithDefaults() *ExperimentsWarehouseExposureFilter {
	this := ExperimentsWarehouseExposureFilter{}
	return &this
}

// GetOperation returns the Operation field value.
func (o *ExperimentsWarehouseExposureFilter) GetOperation() ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation {
	if o == nil {
		var ret ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation
		return ret
	}
	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseExposureFilter) GetOperationOk() (*ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value.
func (o *ExperimentsWarehouseExposureFilter) SetOperation(v ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation) {
	o.Operation = v
}

// GetPropertyId returns the PropertyId field value.
func (o *ExperimentsWarehouseExposureFilter) GetPropertyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PropertyId
}

// GetPropertyIdOk returns a tuple with the PropertyId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseExposureFilter) GetPropertyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyId, true
}

// SetPropertyId sets field value.
func (o *ExperimentsWarehouseExposureFilter) SetPropertyId(v string) {
	o.PropertyId = v
}

// GetValues returns the Values field value.
func (o *ExperimentsWarehouseExposureFilter) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ExperimentsWarehouseExposureFilter) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *ExperimentsWarehouseExposureFilter) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsWarehouseExposureFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["operation"] = o.Operation
	toSerialize["property_id"] = o.PropertyId
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsWarehouseExposureFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Operation  *ExperimentsPatchExperimentV2ResponseDataAttributesWarehouseExposureConfigurationEntryPointFiltersItemsOperation `json:"operation"`
		PropertyId *string                                                                                                          `json:"property_id"`
		Values     *[]string                                                                                                        `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Operation == nil {
		return fmt.Errorf("required field operation missing")
	}
	if all.PropertyId == nil {
		return fmt.Errorf("required field property_id missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"operation", "property_id", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Operation.IsValid() {
		hasInvalidField = true
	} else {
		o.Operation = *all.Operation
	}
	o.PropertyId = *all.PropertyId
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
