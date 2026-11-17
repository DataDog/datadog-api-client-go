// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsBanditV2DTO Response containing a bandit.
type ExperimentsBanditV2DTO struct {
	// JSON:API resource containing a bandit.
	Data ExperimentsBanditV2DTOData `json:"data"`
	// Information about the bandit response.
	Meta *ExperimentsBanditV2ReadMetaDTO `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsBanditV2DTO instantiates a new ExperimentsBanditV2DTO object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsBanditV2DTO(data ExperimentsBanditV2DTOData) *ExperimentsBanditV2DTO {
	this := ExperimentsBanditV2DTO{}
	this.Data = data
	return &this
}

// NewExperimentsBanditV2DTOWithDefaults instantiates a new ExperimentsBanditV2DTO object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsBanditV2DTOWithDefaults() *ExperimentsBanditV2DTO {
	this := ExperimentsBanditV2DTO{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsBanditV2DTO) GetData() ExperimentsBanditV2DTOData {
	if o == nil {
		var ret ExperimentsBanditV2DTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTO) GetDataOk() (*ExperimentsBanditV2DTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsBanditV2DTO) SetData(v ExperimentsBanditV2DTOData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExperimentsBanditV2DTO) GetMeta() ExperimentsBanditV2ReadMetaDTO {
	if o == nil || o.Meta == nil {
		var ret ExperimentsBanditV2ReadMetaDTO
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsBanditV2DTO) GetMetaOk() (*ExperimentsBanditV2ReadMetaDTO, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExperimentsBanditV2DTO) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ExperimentsBanditV2ReadMetaDTO and assigns it to the Meta field.
func (o *ExperimentsBanditV2DTO) SetMeta(v ExperimentsBanditV2ReadMetaDTO) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsBanditV2DTO) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsBanditV2DTO) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *ExperimentsBanditV2DTOData     `json:"data"`
		Meta *ExperimentsBanditV2ReadMetaDTO `json:"meta,omitempty"`
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
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
