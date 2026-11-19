// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentReferenceTableBlock Selects users or accounts from a reference table.
type SegmentReferenceTableBlock struct {
	// Columns to include.
	Columns []SegmentReferenceTableColumn `json:"columns,omitempty"`
	// Query that filters the reference table rows.
	FilterQuery *string `json:"filter_query,omitempty"`
	// How reference table rows are matched to entities.
	JoinCondition SegmentReferenceTableJoinCondition `json:"join_condition"`
	// Name of the block, used in `combination`.
	Name *string `json:"name,omitempty"`
	// Name of the reference table.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentReferenceTableBlock instantiates a new SegmentReferenceTableBlock object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentReferenceTableBlock(joinCondition SegmentReferenceTableJoinCondition, tableName string) *SegmentReferenceTableBlock {
	this := SegmentReferenceTableBlock{}
	this.JoinCondition = joinCondition
	this.TableName = tableName
	return &this
}

// NewSegmentReferenceTableBlockWithDefaults instantiates a new SegmentReferenceTableBlock object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentReferenceTableBlockWithDefaults() *SegmentReferenceTableBlock {
	this := SegmentReferenceTableBlock{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *SegmentReferenceTableBlock) GetColumns() []SegmentReferenceTableColumn {
	if o == nil || o.Columns == nil {
		var ret []SegmentReferenceTableColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentReferenceTableBlock) GetColumnsOk() (*[]SegmentReferenceTableColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *SegmentReferenceTableBlock) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []SegmentReferenceTableColumn and assigns it to the Columns field.
func (o *SegmentReferenceTableBlock) SetColumns(v []SegmentReferenceTableColumn) {
	o.Columns = v
}

// GetFilterQuery returns the FilterQuery field value if set, zero value otherwise.
func (o *SegmentReferenceTableBlock) GetFilterQuery() string {
	if o == nil || o.FilterQuery == nil {
		var ret string
		return ret
	}
	return *o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentReferenceTableBlock) GetFilterQueryOk() (*string, bool) {
	if o == nil || o.FilterQuery == nil {
		return nil, false
	}
	return o.FilterQuery, true
}

// HasFilterQuery returns a boolean if a field has been set.
func (o *SegmentReferenceTableBlock) HasFilterQuery() bool {
	return o != nil && o.FilterQuery != nil
}

// SetFilterQuery gets a reference to the given string and assigns it to the FilterQuery field.
func (o *SegmentReferenceTableBlock) SetFilterQuery(v string) {
	o.FilterQuery = &v
}

// GetJoinCondition returns the JoinCondition field value.
func (o *SegmentReferenceTableBlock) GetJoinCondition() SegmentReferenceTableJoinCondition {
	if o == nil {
		var ret SegmentReferenceTableJoinCondition
		return ret
	}
	return o.JoinCondition
}

// GetJoinConditionOk returns a tuple with the JoinCondition field value
// and a boolean to check if the value has been set.
func (o *SegmentReferenceTableBlock) GetJoinConditionOk() (*SegmentReferenceTableJoinCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinCondition, true
}

// SetJoinCondition sets field value.
func (o *SegmentReferenceTableBlock) SetJoinCondition(v SegmentReferenceTableJoinCondition) {
	o.JoinCondition = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentReferenceTableBlock) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentReferenceTableBlock) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentReferenceTableBlock) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentReferenceTableBlock) SetName(v string) {
	o.Name = &v
}

// GetTableName returns the TableName field value.
func (o *SegmentReferenceTableBlock) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *SegmentReferenceTableBlock) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *SegmentReferenceTableBlock) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentReferenceTableBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.FilterQuery != nil {
		toSerialize["filter_query"] = o.FilterQuery
	}
	toSerialize["join_condition"] = o.JoinCondition
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["table_name"] = o.TableName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SegmentReferenceTableBlock) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns       []SegmentReferenceTableColumn       `json:"columns,omitempty"`
		FilterQuery   *string                             `json:"filter_query,omitempty"`
		JoinCondition *SegmentReferenceTableJoinCondition `json:"join_condition"`
		Name          *string                             `json:"name,omitempty"`
		TableName     *string                             `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JoinCondition == nil {
		return fmt.Errorf("required field join_condition missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "filter_query", "join_condition", "name", "table_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = all.Columns
	o.FilterQuery = all.FilterQuery
	if all.JoinCondition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinCondition = *all.JoinCondition
	o.Name = all.Name
	o.TableName = *all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
