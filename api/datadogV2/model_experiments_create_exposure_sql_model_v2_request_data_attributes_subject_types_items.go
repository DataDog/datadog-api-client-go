// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems Mapping between a subject type and its identifier column in the SQL model.
type ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems struct {
	// SQL result column that contains the subject identifier.
	ColumnName string `json:"column_name"`
	// Identifier of the subject type mapped to this column.
	SubjectTypeId string `json:"subject_type_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems instantiates a new ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems(columnName string, subjectTypeId string) *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems {
	this := ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems{}
	this.ColumnName = columnName
	this.SubjectTypeId = subjectTypeId
	return &this
}

// NewExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItemsWithDefaults instantiates a new ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItemsWithDefaults() *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems {
	this := ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems{}
	return &this
}

// GetColumnName returns the ColumnName field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) GetColumnName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) GetColumnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnName, true
}

// SetColumnName sets field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) SetColumnName(v string) {
	o.ColumnName = v
}

// GetSubjectTypeId returns the SubjectTypeId field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) GetSubjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SubjectTypeId
}

// GetSubjectTypeIdOk returns a tuple with the SubjectTypeId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) GetSubjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubjectTypeId, true
}

// SetSubjectTypeId sets field value.
func (o *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) SetSubjectTypeId(v string) {
	o.SubjectTypeId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["column_name"] = o.ColumnName
	toSerialize["subject_type_id"] = o.SubjectTypeId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExposureSQLModelV2RequestDataAttributesSubjectTypesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ColumnName    *string `json:"column_name"`
		SubjectTypeId *string `json:"subject_type_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ColumnName == nil {
		return fmt.Errorf("required field column_name missing")
	}
	if all.SubjectTypeId == nil {
		return fmt.Errorf("required field subject_type_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"column_name", "subject_type_id"})
	} else {
		return err
	}
	o.ColumnName = *all.ColumnName
	o.SubjectTypeId = *all.SubjectTypeId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
