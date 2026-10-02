// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTOArray List of variant result resources.
type ExperimentsVariantResultsV2DTOArray struct {
	// Resources returned in this response.
	Data []ExperimentsVariantResultsV2DTOData `json:"data"`
	// Information about when experiment results were updated and whether they are stale.
	Meta *ExperimentsExperimentResultsV2MetaDTO `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsVariantResultsV2DTOArray instantiates a new ExperimentsVariantResultsV2DTOArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsVariantResultsV2DTOArray(data []ExperimentsVariantResultsV2DTOData) *ExperimentsVariantResultsV2DTOArray {
	this := ExperimentsVariantResultsV2DTOArray{}
	this.Data = data
	return &this
}

// NewExperimentsVariantResultsV2DTOArrayWithDefaults instantiates a new ExperimentsVariantResultsV2DTOArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsVariantResultsV2DTOArrayWithDefaults() *ExperimentsVariantResultsV2DTOArray {
	this := ExperimentsVariantResultsV2DTOArray{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsVariantResultsV2DTOArray) GetData() []ExperimentsVariantResultsV2DTOData {
	if o == nil {
		var ret []ExperimentsVariantResultsV2DTOData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTOArray) GetDataOk() (*[]ExperimentsVariantResultsV2DTOData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsVariantResultsV2DTOArray) SetData(v []ExperimentsVariantResultsV2DTOData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTOArray) GetMeta() ExperimentsExperimentResultsV2MetaDTO {
	if o == nil || o.Meta == nil {
		var ret ExperimentsExperimentResultsV2MetaDTO
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTOArray) GetMetaOk() (*ExperimentsExperimentResultsV2MetaDTO, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTOArray) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ExperimentsExperimentResultsV2MetaDTO and assigns it to the Meta field.
func (o *ExperimentsVariantResultsV2DTOArray) SetMeta(v ExperimentsExperimentResultsV2MetaDTO) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsVariantResultsV2DTOArray) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsVariantResultsV2DTOArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ExperimentsVariantResultsV2DTOData  `json:"data"`
		Meta *ExperimentsExperimentResultsV2MetaDTO `json:"meta,omitempty"`
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
