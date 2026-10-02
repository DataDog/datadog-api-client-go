// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsUpdateMetricSQLModelV2ResponseMeta Model entries removed by the update. Empty arrays mean no entries were removed.
type ExperimentsUpdateMetricSQLModelV2ResponseMeta struct {
	// Measure column names removed from the model.
	DeletedMeasures []string `json:"deleted_measures,omitempty"`
	// Property names removed from the model.
	DeletedProperties []string `json:"deleted_properties,omitempty"`
	// Subject type IDs removed from the model.
	DeletedSubjectTypes []string `json:"deleted_subject_types,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsUpdateMetricSQLModelV2ResponseMeta instantiates a new ExperimentsUpdateMetricSQLModelV2ResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsUpdateMetricSQLModelV2ResponseMeta() *ExperimentsUpdateMetricSQLModelV2ResponseMeta {
	this := ExperimentsUpdateMetricSQLModelV2ResponseMeta{}
	return &this
}

// NewExperimentsUpdateMetricSQLModelV2ResponseMetaWithDefaults instantiates a new ExperimentsUpdateMetricSQLModelV2ResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsUpdateMetricSQLModelV2ResponseMetaWithDefaults() *ExperimentsUpdateMetricSQLModelV2ResponseMeta {
	this := ExperimentsUpdateMetricSQLModelV2ResponseMeta{}
	return &this
}

// GetDeletedMeasures returns the DeletedMeasures field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) GetDeletedMeasures() []string {
	if o == nil || o.DeletedMeasures == nil {
		var ret []string
		return ret
	}
	return o.DeletedMeasures
}

// GetDeletedMeasuresOk returns a tuple with the DeletedMeasures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) GetDeletedMeasuresOk() (*[]string, bool) {
	if o == nil || o.DeletedMeasures == nil {
		return nil, false
	}
	return &o.DeletedMeasures, true
}

// HasDeletedMeasures returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) HasDeletedMeasures() bool {
	return o != nil && o.DeletedMeasures != nil
}

// SetDeletedMeasures gets a reference to the given []string and assigns it to the DeletedMeasures field.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) SetDeletedMeasures(v []string) {
	o.DeletedMeasures = v
}

// GetDeletedProperties returns the DeletedProperties field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) GetDeletedProperties() []string {
	if o == nil || o.DeletedProperties == nil {
		var ret []string
		return ret
	}
	return o.DeletedProperties
}

// GetDeletedPropertiesOk returns a tuple with the DeletedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) GetDeletedPropertiesOk() (*[]string, bool) {
	if o == nil || o.DeletedProperties == nil {
		return nil, false
	}
	return &o.DeletedProperties, true
}

// HasDeletedProperties returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) HasDeletedProperties() bool {
	return o != nil && o.DeletedProperties != nil
}

// SetDeletedProperties gets a reference to the given []string and assigns it to the DeletedProperties field.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) SetDeletedProperties(v []string) {
	o.DeletedProperties = v
}

// GetDeletedSubjectTypes returns the DeletedSubjectTypes field value if set, zero value otherwise.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) GetDeletedSubjectTypes() []string {
	if o == nil || o.DeletedSubjectTypes == nil {
		var ret []string
		return ret
	}
	return o.DeletedSubjectTypes
}

// GetDeletedSubjectTypesOk returns a tuple with the DeletedSubjectTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) GetDeletedSubjectTypesOk() (*[]string, bool) {
	if o == nil || o.DeletedSubjectTypes == nil {
		return nil, false
	}
	return &o.DeletedSubjectTypes, true
}

// HasDeletedSubjectTypes returns a boolean if a field has been set.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) HasDeletedSubjectTypes() bool {
	return o != nil && o.DeletedSubjectTypes != nil
}

// SetDeletedSubjectTypes gets a reference to the given []string and assigns it to the DeletedSubjectTypes field.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) SetDeletedSubjectTypes(v []string) {
	o.DeletedSubjectTypes = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsUpdateMetricSQLModelV2ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedMeasures != nil {
		toSerialize["deleted_measures"] = o.DeletedMeasures
	}
	if o.DeletedProperties != nil {
		toSerialize["deleted_properties"] = o.DeletedProperties
	}
	if o.DeletedSubjectTypes != nil {
		toSerialize["deleted_subject_types"] = o.DeletedSubjectTypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsUpdateMetricSQLModelV2ResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedMeasures     []string `json:"deleted_measures,omitempty"`
		DeletedProperties   []string `json:"deleted_properties,omitempty"`
		DeletedSubjectTypes []string `json:"deleted_subject_types,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted_measures", "deleted_properties", "deleted_subject_types"})
	} else {
		return err
	}
	o.DeletedMeasures = all.DeletedMeasures
	o.DeletedProperties = all.DeletedProperties
	o.DeletedSubjectTypes = all.DeletedSubjectTypes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
