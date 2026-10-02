// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsRefreshExperimentResultsV2DTOArray List of experiment refresh result resources.
type ExperimentsRefreshExperimentResultsV2DTOArray struct {
	// Resources returned in this response.
	Data []ExperimentsRefreshExperimentResultsV2DTOData `json:"data"`
	// Summary of refresh outcomes across the organization's experiments.
	Meta *ExperimentsRefreshExperimentResultsBatchMetaV2DTO `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsRefreshExperimentResultsV2DTOArray instantiates a new ExperimentsRefreshExperimentResultsV2DTOArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsRefreshExperimentResultsV2DTOArray(data []ExperimentsRefreshExperimentResultsV2DTOData) *ExperimentsRefreshExperimentResultsV2DTOArray {
	this := ExperimentsRefreshExperimentResultsV2DTOArray{}
	this.Data = data
	return &this
}

// NewExperimentsRefreshExperimentResultsV2DTOArrayWithDefaults instantiates a new ExperimentsRefreshExperimentResultsV2DTOArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsRefreshExperimentResultsV2DTOArrayWithDefaults() *ExperimentsRefreshExperimentResultsV2DTOArray {
	this := ExperimentsRefreshExperimentResultsV2DTOArray{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) GetData() []ExperimentsRefreshExperimentResultsV2DTOData {
	if o == nil {
		var ret []ExperimentsRefreshExperimentResultsV2DTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) GetDataOk() (*[]ExperimentsRefreshExperimentResultsV2DTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) SetData(v []ExperimentsRefreshExperimentResultsV2DTOData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) GetMeta() ExperimentsRefreshExperimentResultsBatchMetaV2DTO {
	if o == nil || o.Meta == nil {
		var ret ExperimentsRefreshExperimentResultsBatchMetaV2DTO
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) GetMetaOk() (*ExperimentsRefreshExperimentResultsBatchMetaV2DTO, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ExperimentsRefreshExperimentResultsBatchMetaV2DTO and assigns it to the Meta field.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) SetMeta(v ExperimentsRefreshExperimentResultsBatchMetaV2DTO) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsRefreshExperimentResultsV2DTOArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsRefreshExperimentResultsV2DTOArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ExperimentsRefreshExperimentResultsV2DTOData    `json:"data"`
		Meta *ExperimentsRefreshExperimentResultsBatchMetaV2DTO `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
