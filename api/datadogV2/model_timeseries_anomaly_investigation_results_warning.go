// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesAnomalyInvestigationResultsWarning Non-fatal warning produced while executing the investigation.
type TimeseriesAnomalyInvestigationResultsWarning struct {
	// Human-readable warning message.
	Message string `json:"message"`
	// Machine-readable warning name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesAnomalyInvestigationResultsWarning instantiates a new TimeseriesAnomalyInvestigationResultsWarning object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesAnomalyInvestigationResultsWarning(message string, name string) *TimeseriesAnomalyInvestigationResultsWarning {
	this := TimeseriesAnomalyInvestigationResultsWarning{}
	this.Message = message
	this.Name = name
	return &this
}

// NewTimeseriesAnomalyInvestigationResultsWarningWithDefaults instantiates a new TimeseriesAnomalyInvestigationResultsWarning object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesAnomalyInvestigationResultsWarningWithDefaults() *TimeseriesAnomalyInvestigationResultsWarning {
	this := TimeseriesAnomalyInvestigationResultsWarning{}
	return &this
}

// GetMessage returns the Message field value.
func (o *TimeseriesAnomalyInvestigationResultsWarning) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResultsWarning) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *TimeseriesAnomalyInvestigationResultsWarning) SetMessage(v string) {
	o.Message = v
}

// GetName returns the Name field value.
func (o *TimeseriesAnomalyInvestigationResultsWarning) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TimeseriesAnomalyInvestigationResultsWarning) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *TimeseriesAnomalyInvestigationResultsWarning) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesAnomalyInvestigationResultsWarning) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["message"] = o.Message
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesAnomalyInvestigationResultsWarning) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Message *string `json:"message"`
		Name    *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"message", "name"})
	} else {
		return err
	}
	o.Message = *all.Message
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
