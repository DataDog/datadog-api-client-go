// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetCohortResponseDataAttributesCohortsItems
type GetCohortResponseDataAttributesCohortsItems struct {
	//
	Cohort *string `json:"cohort,omitempty"`
	//
	CohortSize *int64 `json:"cohort_size,omitempty"`
	//
	StartTime *int64 `json:"start_time,omitempty"`
	//
	Values []GetCohortResponseDataAttributesCohortsItemsValuesItems `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetCohortResponseDataAttributesCohortsItems instantiates a new GetCohortResponseDataAttributesCohortsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetCohortResponseDataAttributesCohortsItems() *GetCohortResponseDataAttributesCohortsItems {
	this := GetCohortResponseDataAttributesCohortsItems{}
	return &this
}

// NewGetCohortResponseDataAttributesCohortsItemsWithDefaults instantiates a new GetCohortResponseDataAttributesCohortsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetCohortResponseDataAttributesCohortsItemsWithDefaults() *GetCohortResponseDataAttributesCohortsItems {
	this := GetCohortResponseDataAttributesCohortsItems{}
	return &this
}

// GetCohort returns the Cohort field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItems) GetCohort() string {
	if o == nil || o.Cohort == nil {
		var ret string
		return ret
	}
	return *o.Cohort
}

// GetCohortOk returns a tuple with the Cohort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) GetCohortOk() (*string, bool) {
	if o == nil || o.Cohort == nil {
		return nil, false
	}
	return o.Cohort, true
}

// HasCohort returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) HasCohort() bool {
	return o != nil && o.Cohort != nil
}

// SetCohort gets a reference to the given string and assigns it to the Cohort field.
func (o *GetCohortResponseDataAttributesCohortsItems) SetCohort(v string) {
	o.Cohort = &v
}

// GetCohortSize returns the CohortSize field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItems) GetCohortSize() int64 {
	if o == nil || o.CohortSize == nil {
		var ret int64
		return ret
	}
	return *o.CohortSize
}

// GetCohortSizeOk returns a tuple with the CohortSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) GetCohortSizeOk() (*int64, bool) {
	if o == nil || o.CohortSize == nil {
		return nil, false
	}
	return o.CohortSize, true
}

// HasCohortSize returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) HasCohortSize() bool {
	return o != nil && o.CohortSize != nil
}

// SetCohortSize gets a reference to the given int64 and assigns it to the CohortSize field.
func (o *GetCohortResponseDataAttributesCohortsItems) SetCohortSize(v int64) {
	o.CohortSize = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItems) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) HasStartTime() bool {
	return o != nil && o.StartTime != nil
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *GetCohortResponseDataAttributesCohortsItems) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *GetCohortResponseDataAttributesCohortsItems) GetValues() []GetCohortResponseDataAttributesCohortsItemsValuesItems {
	if o == nil || o.Values == nil {
		var ret []GetCohortResponseDataAttributesCohortsItemsValuesItems
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) GetValuesOk() (*[]GetCohortResponseDataAttributesCohortsItemsValuesItems, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *GetCohortResponseDataAttributesCohortsItems) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []GetCohortResponseDataAttributesCohortsItemsValuesItems and assigns it to the Values field.
func (o *GetCohortResponseDataAttributesCohortsItems) SetValues(v []GetCohortResponseDataAttributesCohortsItemsValuesItems) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetCohortResponseDataAttributesCohortsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cohort != nil {
		toSerialize["cohort"] = o.Cohort
	}
	if o.CohortSize != nil {
		toSerialize["cohort_size"] = o.CohortSize
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetCohortResponseDataAttributesCohortsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cohort     *string                                                  `json:"cohort,omitempty"`
		CohortSize *int64                                                   `json:"cohort_size,omitempty"`
		StartTime  *int64                                                   `json:"start_time,omitempty"`
		Values     []GetCohortResponseDataAttributesCohortsItemsValuesItems `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cohort", "cohort_size", "start_time", "values"})
	} else {
		return err
	}
	o.Cohort = all.Cohort
	o.CohortSize = all.CohortSize
	o.StartTime = all.StartTime
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
