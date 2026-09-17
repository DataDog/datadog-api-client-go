// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationQueryStatus Execution status for one named query.
type TimeseriesAnomalyInvestigationQueryStatus struct {
	// Query name from the request.
	Name string `json:"name"`
	// Current execution status for a named query.
	Status TimeseriesAnomalyInvestigationQueryExecutionStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationQueryStatus instantiates a new TimeseriesAnomalyInvestigationQueryStatus object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationQueryStatus(name string, status TimeseriesAnomalyInvestigationQueryExecutionStatus) *TimeseriesAnomalyInvestigationQueryStatus {
	this := TimeseriesAnomalyInvestigationQueryStatus{}
	this.Name = name
	this.Status = status
	return &this
}

// NewTimeseriesAnomalyInvestigationQueryStatusWithDefaults instantiates a new TimeseriesAnomalyInvestigationQueryStatus object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationQueryStatusWithDefaults() *TimeseriesAnomalyInvestigationQueryStatus {
	this := TimeseriesAnomalyInvestigationQueryStatus{}
	return &this
}

// GetName returns the Name field value.
func (o *TimeseriesAnomalyInvestigationQueryStatus) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationQueryStatus) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TimeseriesAnomalyInvestigationQueryStatus) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value.
func (o *TimeseriesAnomalyInvestigationQueryStatus) GetStatus() TimeseriesAnomalyInvestigationQueryExecutionStatus {
	if o == nil {
		var ret TimeseriesAnomalyInvestigationQueryExecutionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationQueryStatus) GetStatusOk() (*TimeseriesAnomalyInvestigationQueryExecutionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *TimeseriesAnomalyInvestigationQueryStatus) SetStatus(v TimeseriesAnomalyInvestigationQueryExecutionStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationQueryStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationQueryStatus) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name   *string                                             `json:"name"`
		Status *TimeseriesAnomalyInvestigationQueryExecutionStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = *all.Name
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
