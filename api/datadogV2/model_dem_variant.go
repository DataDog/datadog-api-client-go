// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DemVariant A variant (sub-funnel) of a DEM journey with its own steps.
type DemVariant struct {
	// An optional RUM query filter to scope this variant.
	Filter *string `json:"filter,omitempty"`
	// The unique identifier of the variant.
	Id *string `json:"id,omitempty"`
	// The name of the variant.
	Name string `json:"name"`
	// List of RUM journey steps.
	RumSteps []DemRumStep `json:"rum_steps"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDemVariant instantiates a new DemVariant object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDemVariant(name string, rumSteps []DemRumStep) *DemVariant {
	this := DemVariant{}
	this.Name = name
	this.RumSteps = rumSteps
	return &this
}

// NewDemVariantWithDefaults instantiates a new DemVariant object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDemVariantWithDefaults() *DemVariant {
	this := DemVariant{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DemVariant) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemVariant) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DemVariant) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *DemVariant) SetFilter(v string) {
	o.Filter = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DemVariant) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DemVariant) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DemVariant) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DemVariant) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value.
func (o *DemVariant) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DemVariant) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DemVariant) SetName(v string) {
	o.Name = v
}

// GetRumSteps returns the RumSteps field value.
func (o *DemVariant) GetRumSteps() []DemRumStep {
	if o == nil {
		var ret []DemRumStep
		return ret
	}
	return o.RumSteps
}

// GetRumStepsOk returns a tuple with the RumSteps field value
// and a boolean to check if the value has been set.
func (o *DemVariant) GetRumStepsOk() (*[]DemRumStep, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RumSteps, true
}

// SetRumSteps sets field value.
func (o *DemVariant) SetRumSteps(v []DemRumStep) {
	o.RumSteps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DemVariant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["rum_steps"] = o.RumSteps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DemVariant) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filter   *string       `json:"filter,omitempty"`
		Id       *string       `json:"id,omitempty"`
		Name     *string       `json:"name"`
		RumSteps *[]DemRumStep `json:"rum_steps"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.RumSteps == nil {
		return fmt.Errorf("required field rum_steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filter", "id", "name", "rum_steps"})
	} else {
		return err
	}
	o.Filter = all.Filter
	o.Id = all.Id
	o.Name = *all.Name
	o.RumSteps = *all.RumSteps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
