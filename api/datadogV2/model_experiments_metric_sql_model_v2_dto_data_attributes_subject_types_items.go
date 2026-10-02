// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems A mapping between a subject type and its identifying SQL column.
type ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems struct {
	// Name of the SQL result column that identifies subjects of this type.
	ColumnName *string `json:"column_name,omitempty"`
	// ID of the subject type used by this configuration.
	SubjectTypeId *string `json:"subject_type_id,omitempty"`
	// Read-only measure ID. Pass it as warehouse_metric_measure.id when the metric operation is `uniqueSubjects`.
	UniqueSubjectCountMeasureId *string `json:"unique_subject_count_measure_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems instantiates a new ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems() *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems {
	this := ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems{}
	return &this
}

// NewExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItemsWithDefaults instantiates a new ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItemsWithDefaults() *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems {
	this := ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems{}
	return &this
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) HasColumnName() bool {
	return o != nil && o.ColumnName != nil
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetSubjectTypeId returns the SubjectTypeId field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) GetSubjectTypeId() string {
	if o == nil || o.SubjectTypeId == nil {
		var ret string
		return ret
	}
	return *o.SubjectTypeId
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil || o.SubjectTypeId == nil {
		return nil, false
	}
	return o.SubjectTypeId, true
}

// HasSubjectTypeId returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) HasSubjectTypeId() bool {
	return o != nil && o.SubjectTypeId != nil
}

// SetSubjectTypeId gets a reference to the given string and assigns it to the SubjectTypeId field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) SetSubjectTypeId(v string) {
	o.SubjectTypeId = &v
}

// GetUniqueSubjectCountMeasureId returns the UniqueSubjectCountMeasureId field value if set, zero value otherwise.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) GetUniqueSubjectCountMeasureId() string {
	if o == nil || o.UniqueSubjectCountMeasureId == nil {
		var ret string
		return ret
	}
	return *o.UniqueSubjectCountMeasureId
}

// GetUniqueSubjectCountMeasureIdOk returns a tuple with the UniqueSubjectCountMeasureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) GetUniqueSubjectCountMeasureIdOk() (*string, bool) {
	if o == nil || o.UniqueSubjectCountMeasureId == nil {
		return nil, false
	}
	return o.UniqueSubjectCountMeasureId, true
}

// HasUniqueSubjectCountMeasureId returns a boolean if a field has been set.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) HasUniqueSubjectCountMeasureId() bool {
	return o != nil && o.UniqueSubjectCountMeasureId != nil
}

// SetUniqueSubjectCountMeasureId gets a reference to the given string and assigns it to the UniqueSubjectCountMeasureId field.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) SetUniqueSubjectCountMeasureId(v string) {
	o.UniqueSubjectCountMeasureId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ColumnName != nil {
		toSerialize["column_name"] = o.ColumnName
	}
	if o.SubjectTypeId != nil {
		toSerialize["subject_type_id"] = o.SubjectTypeId
	}
	if o.UniqueSubjectCountMeasureId != nil {
		toSerialize["unique_subject_count_measure_id"] = o.UniqueSubjectCountMeasureId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricSQLModelV2DTODataAttributesSubjectTypesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName                  *string `json:"column_name,omitempty"`
		SubjectTypeId               *string `json:"subject_type_id,omitempty"`
		UniqueSubjectCountMeasureId *string `json:"unique_subject_count_measure_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "subject_type_id", "unique_subject_count_measure_id"})
	} else {
		return err
	}
	o.ColumnName = all.ColumnName
	o.SubjectTypeId = all.SubjectTypeId
	o.UniqueSubjectCountMeasureId = all.UniqueSubjectCountMeasureId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
