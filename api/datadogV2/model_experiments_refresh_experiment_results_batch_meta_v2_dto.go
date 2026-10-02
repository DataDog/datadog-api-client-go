// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsRefreshExperimentResultsBatchMetaV2DTO Summary of refresh outcomes across the organization's experiments.
type ExperimentsRefreshExperimentResultsBatchMetaV2DTO struct {
	// Number of experiments updated by the refresh request.
	ExperimentsUpdated *int64 `json:"experiments_updated,omitempty"`
	// Refresh outcome reported for each experiment.
	Results []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems `json:"results,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsRefreshExperimentResultsBatchMetaV2DTO instantiates a new ExperimentsRefreshExperimentResultsBatchMetaV2DTO object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsRefreshExperimentResultsBatchMetaV2DTO() *ExperimentsRefreshExperimentResultsBatchMetaV2DTO {
	this := ExperimentsRefreshExperimentResultsBatchMetaV2DTO{}
	return &this
}

// NewExperimentsRefreshExperimentResultsBatchMetaV2DTOWithDefaults instantiates a new ExperimentsRefreshExperimentResultsBatchMetaV2DTO object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsRefreshExperimentResultsBatchMetaV2DTOWithDefaults() *ExperimentsRefreshExperimentResultsBatchMetaV2DTO {
	this := ExperimentsRefreshExperimentResultsBatchMetaV2DTO{}
	return &this
}

// GetExperimentsUpdated returns the ExperimentsUpdated field value if set, zero value otherwise.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) GetExperimentsUpdated() int64 {
	if o == nil || o.ExperimentsUpdated == nil {
		var ret int64
		return ret
	}
	return *o.ExperimentsUpdated
}

// GetExperimentsUpdatedOk returns a tuple with the ExperimentsUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) GetExperimentsUpdatedOk() (*int64, bool) {
	if o == nil || o.ExperimentsUpdated == nil {
		return nil, false
	}
	return o.ExperimentsUpdated, true
}

// HasExperimentsUpdated returns a boolean if a field has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) HasExperimentsUpdated() bool {
	return o != nil && o.ExperimentsUpdated != nil
}

// SetExperimentsUpdated gets a reference to the given int64 and assigns it to the ExperimentsUpdated field.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) SetExperimentsUpdated(v int64) {
	o.ExperimentsUpdated = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) GetResults() []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems {
	if o == nil || o.Results == nil {
		var ret []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) GetResultsOk() (*[]ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return &o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) HasResults() bool {
	return o != nil && o.Results != nil
}

// SetResults gets a reference to the given []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems and assigns it to the Results field.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) SetResults(v []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems) {
	o.Results = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsRefreshExperimentResultsBatchMetaV2DTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExperimentsUpdated != nil {
		toSerialize["experiments_updated"] = o.ExperimentsUpdated
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsRefreshExperimentResultsBatchMetaV2DTO) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExperimentsUpdated *int64                                                          `json:"experiments_updated,omitempty"`
		Results            []ExperimentsRefreshExperimentResultsBatchMetaV2DTOResultsItems `json:"results,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"experiments_updated", "results"})
	} else {
		return err
	}
	o.ExperimentsUpdated = all.ExperimentsUpdated
	o.Results = all.Results

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
