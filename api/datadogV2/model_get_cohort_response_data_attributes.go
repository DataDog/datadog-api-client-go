// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortResponseDataAttributes
type GetCohortResponseDataAttributes struct {
	//
	Cohorts []GetCohortResponseDataAttributesCohortsItems `json:"cohorts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetCohortResponseDataAttributes instantiates a new GetCohortResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetCohortResponseDataAttributes() *GetCohortResponseDataAttributes {
	this := GetCohortResponseDataAttributes{}
	return &this
}

// NewGetCohortResponseDataAttributesWithDefaults instantiates a new GetCohortResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetCohortResponseDataAttributesWithDefaults() *GetCohortResponseDataAttributes {
	this := GetCohortResponseDataAttributes{}
	return &this
}

// GetCohorts returns the Cohorts field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributes) GetCohorts() []GetCohortResponseDataAttributesCohortsItems {
	if o == nil || o.Cohorts == nil {
		var ret []GetCohortResponseDataAttributesCohortsItems
		return ret
	}
	return o.Cohorts
}

// GetCohortsOk returns a tuple with the Cohorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributes) GetCohortsOk() (*[]GetCohortResponseDataAttributesCohortsItems, bool) {
	if o == nil || o.Cohorts == nil {
		return nil, false
	}
	return &o.Cohorts, true
}

// HasCohorts returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributes) HasCohorts() bool {
	return o != nil && o.Cohorts != nil
}

// SetCohorts gets a reference to the given []GetCohortResponseDataAttributesCohortsItems and assigns it to the Cohorts field.
func (o *GetCohortResponseDataAttributes) SetCohorts(v []GetCohortResponseDataAttributesCohortsItems) {
	o.Cohorts = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetCohortResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cohorts != nil {
		toSerialize["cohorts"] = o.Cohorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetCohortResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cohorts []GetCohortResponseDataAttributesCohortsItems `json:"cohorts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cohorts"})
	} else {
		return err
	}
	o.Cohorts = all.Cohorts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
