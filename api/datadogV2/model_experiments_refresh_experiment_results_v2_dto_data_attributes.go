// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsRefreshExperimentResultsV2DTODataAttributes Details of the experiment refresh result.
type ExperimentsRefreshExperimentResultsV2DTODataAttributes struct {
	// Whether the refresh request succeeded.
	Success *bool `json:"success,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsRefreshExperimentResultsV2DTODataAttributes instantiates a new ExperimentsRefreshExperimentResultsV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsRefreshExperimentResultsV2DTODataAttributes() *ExperimentsRefreshExperimentResultsV2DTODataAttributes {
	this := ExperimentsRefreshExperimentResultsV2DTODataAttributes{}
	return &this
}

// NewExperimentsRefreshExperimentResultsV2DTODataAttributesWithDefaults instantiates a new ExperimentsRefreshExperimentResultsV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsRefreshExperimentResultsV2DTODataAttributesWithDefaults() *ExperimentsRefreshExperimentResultsV2DTODataAttributes {
	this := ExperimentsRefreshExperimentResultsV2DTODataAttributes{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ExperimentsRefreshExperimentResultsV2DTODataAttributes) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsRefreshExperimentResultsV2DTODataAttributes) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ExperimentsRefreshExperimentResultsV2DTODataAttributes) HasSuccess() bool {
	return o != nil && o.Success != nil
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ExperimentsRefreshExperimentResultsV2DTODataAttributes) SetSuccess(v bool) {
	o.Success = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsRefreshExperimentResultsV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsRefreshExperimentResultsV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Success *bool `json:"success,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"success"})
	} else {
		return err
	}
	o.Success = all.Success

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
