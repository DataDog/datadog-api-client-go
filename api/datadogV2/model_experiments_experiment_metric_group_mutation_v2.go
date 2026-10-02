// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentMetricGroupMutationV2 Response containing the saved experiment metric group and result refresh information.
type ExperimentsExperimentMetricGroupMutationV2 struct {
	// Experiment metric group resource with its identifier and metric selection.
	Data ExperimentsExperimentMetricGroupMutationV2Data `json:"data"`
	// Refresh requirements and warnings returned by an experiment update.
	Meta *ExperimentsPatchExperimentV2MetaDTO `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentMetricGroupMutationV2 instantiates a new ExperimentsExperimentMetricGroupMutationV2 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentMetricGroupMutationV2(data ExperimentsExperimentMetricGroupMutationV2Data) *ExperimentsExperimentMetricGroupMutationV2 {
	this := ExperimentsExperimentMetricGroupMutationV2{}
	this.Data = data
	return &this
}

// NewExperimentsExperimentMetricGroupMutationV2WithDefaults instantiates a new ExperimentsExperimentMetricGroupMutationV2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentMetricGroupMutationV2WithDefaults() *ExperimentsExperimentMetricGroupMutationV2 {
	this := ExperimentsExperimentMetricGroupMutationV2{}
	return &this
}

// GetData returns the Data field value.
func (o *ExperimentsExperimentMetricGroupMutationV2) GetData() ExperimentsExperimentMetricGroupMutationV2Data {
	if o == nil {
		var ret ExperimentsExperimentMetricGroupMutationV2Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2) GetDataOk() (*ExperimentsExperimentMetricGroupMutationV2Data, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ExperimentsExperimentMetricGroupMutationV2) SetData(v ExperimentsExperimentMetricGroupMutationV2Data) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ExperimentsExperimentMetricGroupMutationV2) GetMeta() ExperimentsPatchExperimentV2MetaDTO {
	if o == nil || o.Meta == nil {
		var ret ExperimentsPatchExperimentV2MetaDTO
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2) GetMetaOk() (*ExperimentsPatchExperimentV2MetaDTO, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ExperimentsPatchExperimentV2MetaDTO and assigns it to the Meta field.
func (o *ExperimentsExperimentMetricGroupMutationV2) SetMeta(v ExperimentsPatchExperimentV2MetaDTO) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentMetricGroupMutationV2) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsExperimentMetricGroupMutationV2) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *ExperimentsExperimentMetricGroupMutationV2Data `json:"data"`
		Meta *ExperimentsPatchExperimentV2MetaDTO            `json:"meta,omitempty"`
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
