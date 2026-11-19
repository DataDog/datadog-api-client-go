// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditRefreshV2DTO Response acknowledging a refresh that has started.
type ExperimentsBanditRefreshV2DTO struct {
	// The refresh that started.
	Data ExperimentsBanditRefreshV2DTOData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditRefreshV2DTO instantiates a new ExperimentsBanditRefreshV2DTO object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditRefreshV2DTO(data ExperimentsBanditRefreshV2DTOData) *ExperimentsBanditRefreshV2DTO {
	this := ExperimentsBanditRefreshV2DTO{}
	this.Data = data
	return &this
}

// NewExperimentsBanditRefreshV2DTOWithDefaults instantiates a new ExperimentsBanditRefreshV2DTO object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditRefreshV2DTOWithDefaults() *ExperimentsBanditRefreshV2DTO {
	this := ExperimentsBanditRefreshV2DTO{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsBanditRefreshV2DTO) GetData() ExperimentsBanditRefreshV2DTOData {
	if o == nil {
		var ret ExperimentsBanditRefreshV2DTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditRefreshV2DTO) GetDataOk() (*ExperimentsBanditRefreshV2DTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsBanditRefreshV2DTO) SetData(v ExperimentsBanditRefreshV2DTOData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditRefreshV2DTO) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsBanditRefreshV2DTO) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *ExperimentsBanditRefreshV2DTOData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
