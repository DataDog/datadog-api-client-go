// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateExposureSQLModelV2ResponseMeta Removed model entries. Present only when the update removes an entry.
type ExperimentsUpdateExposureSQLModelV2ResponseMeta struct {
	// Property names removed from the model.
	RemovedPropertyNames []string `json:"removed_property_names,omitempty"`
	// Subject type IDs removed from the model.
	RemovedSubjectTypeIds []string `json:"removed_subject_type_ids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsUpdateExposureSQLModelV2ResponseMeta instantiates a new ExperimentsUpdateExposureSQLModelV2ResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsUpdateExposureSQLModelV2ResponseMeta() *ExperimentsUpdateExposureSQLModelV2ResponseMeta {
	this := ExperimentsUpdateExposureSQLModelV2ResponseMeta{}
	return &this
}

// NewExperimentsUpdateExposureSQLModelV2ResponseMetaWithDefaults instantiates a new ExperimentsUpdateExposureSQLModelV2ResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsUpdateExposureSQLModelV2ResponseMetaWithDefaults() *ExperimentsUpdateExposureSQLModelV2ResponseMeta {
	this := ExperimentsUpdateExposureSQLModelV2ResponseMeta{}
	return &this
}

// GetRemovedPropertyNames returns the RemovedPropertyNames field value if set, zero value otherwise.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) GetRemovedPropertyNames() []string {
	if o == nil || o.RemovedPropertyNames == nil {
		var ret []string
		return ret
	}
	return o.RemovedPropertyNames
}

// GetRemovedPropertyNamesOk returns a tuple with the RemovedPropertyNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) GetRemovedPropertyNamesOk() (*[]string, bool) {
	if o == nil || o.RemovedPropertyNames == nil {
		return nil, false
	}
	return &o.RemovedPropertyNames, true
}

// HasRemovedPropertyNames returns a boolean if a field has been set.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) HasRemovedPropertyNames() bool {
	return o != nil && o.RemovedPropertyNames != nil
}

// SetRemovedPropertyNames gets a reference to the given []string and assigns it to the RemovedPropertyNames field.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) SetRemovedPropertyNames(v []string) {
	o.RemovedPropertyNames = v
}

// GetRemovedSubjectTypeIds returns the RemovedSubjectTypeIds field value if set, zero value otherwise.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) GetRemovedSubjectTypeIds() []string {
	if o == nil || o.RemovedSubjectTypeIds == nil {
		var ret []string
		return ret
	}
	return o.RemovedSubjectTypeIds
}

// GetRemovedSubjectTypeIdsOk returns a tuple with the RemovedSubjectTypeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) GetRemovedSubjectTypeIdsOk() (*[]string, bool) {
	if o == nil || o.RemovedSubjectTypeIds == nil {
		return nil, false
	}
	return &o.RemovedSubjectTypeIds, true
}

// HasRemovedSubjectTypeIds returns a boolean if a field has been set.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) HasRemovedSubjectTypeIds() bool {
	return o != nil && o.RemovedSubjectTypeIds != nil
}

// SetRemovedSubjectTypeIds gets a reference to the given []string and assigns it to the RemovedSubjectTypeIds field.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) SetRemovedSubjectTypeIds(v []string) {
	o.RemovedSubjectTypeIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsUpdateExposureSQLModelV2ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RemovedPropertyNames != nil {
		toSerialize["removed_property_names"] = o.RemovedPropertyNames
	}
	if o.RemovedSubjectTypeIds != nil {
		toSerialize["removed_subject_type_ids"] = o.RemovedSubjectTypeIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsUpdateExposureSQLModelV2ResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RemovedPropertyNames  []string `json:"removed_property_names,omitempty"`
		RemovedSubjectTypeIds []string `json:"removed_subject_type_ids,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"removed_property_names", "removed_subject_type_ids"})
	} else {
		return err
	}
	o.RemovedPropertyNames = all.RemovedPropertyNames
	o.RemovedSubjectTypeIds = all.RemovedSubjectTypeIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
