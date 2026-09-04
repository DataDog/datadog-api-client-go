// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationFindingSynonym Tag grouped under an influential tag by synonym analysis.
type TimeseriesAnomalyInvestigationFindingSynonym struct {
	// Synonymous tag key.
	Key string `json:"key"`
	// Values associated with the synonymous tag.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationFindingSynonym instantiates a new TimeseriesAnomalyInvestigationFindingSynonym object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationFindingSynonym(key string, values []string) *TimeseriesAnomalyInvestigationFindingSynonym {
	this := TimeseriesAnomalyInvestigationFindingSynonym{}
	this.Key = key
	this.Values = values
	return &this
}

// NewTimeseriesAnomalyInvestigationFindingSynonymWithDefaults instantiates a new TimeseriesAnomalyInvestigationFindingSynonym object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationFindingSynonymWithDefaults() *TimeseriesAnomalyInvestigationFindingSynonym {
	this := TimeseriesAnomalyInvestigationFindingSynonym{}
	return &this
}

// GetKey returns the Key field value.
func (o *TimeseriesAnomalyInvestigationFindingSynonym) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFindingSynonym) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value.
func (o *TimeseriesAnomalyInvestigationFindingSynonym) SetKey(v string) {
	o.Key = v
}

// GetValues returns the Values field value.
func (o *TimeseriesAnomalyInvestigationFindingSynonym) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationFindingSynonym) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *TimeseriesAnomalyInvestigationFindingSynonym) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationFindingSynonym) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["key"] = o.Key
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationFindingSynonym) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key    *string   `json:"key"`
		Values *[]string `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Key == nil {
		return fmt.Errorf("required field key missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "values"})
	} else {
		return err
	}
	o.Key = *all.Key
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
