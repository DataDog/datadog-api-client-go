// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSegmentReferenceTable A reference table query block within a segment data query.
type RumSegmentReferenceTable struct {
	// The columns to include from the reference table.
	Columns []RumSegmentReferenceTableColumn `json:"columns"`
	// An optional filter query for the reference table data.
	FilterQuery *string `json:"filter_query,omitempty"`
	// The join condition for a reference table query block.
	JoinCondition RumSegmentReferenceTableJoinCondition `json:"join_condition"`
	// The name of this query block.
	Name string `json:"name"`
	// The name of the reference table.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSegmentReferenceTable instantiates a new RumSegmentReferenceTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSegmentReferenceTable(columns []RumSegmentReferenceTableColumn, joinCondition RumSegmentReferenceTableJoinCondition, name string, tableName string) *RumSegmentReferenceTable {
	this := RumSegmentReferenceTable{}
	this.Columns = columns
	this.JoinCondition = joinCondition
	this.Name = name
	this.TableName = tableName
	return &this
}

// NewRumSegmentReferenceTableWithDefaults instantiates a new RumSegmentReferenceTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSegmentReferenceTableWithDefaults() *RumSegmentReferenceTable {
	this := RumSegmentReferenceTable{}
	return &this
}

// GetColumns returns the Columns field value.
func (o *RumSegmentReferenceTable) GetColumns() []RumSegmentReferenceTableColumn {
	if o == nil {
		var ret []RumSegmentReferenceTableColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *RumSegmentReferenceTable) GetColumnsOk() (*[]RumSegmentReferenceTableColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value.
func (o *RumSegmentReferenceTable) SetColumns(v []RumSegmentReferenceTableColumn) {
	o.Columns = v
}

// GetFilterQuery returns the FilterQuery field value if set, zero value otherwise.
func (o *RumSegmentReferenceTable) GetFilterQuery() string {
	if o == nil || o.FilterQuery == nil {
		var ret string
		return ret
	}
	return *o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSegmentReferenceTable) GetFilterQueryOk() (*string, bool) {
	if o == nil || o.FilterQuery == nil {
		return nil, false
	}
	return o.FilterQuery, true
}

// HasFilterQuery returns a boolean if a field has been set.
func (o *RumSegmentReferenceTable) HasFilterQuery() bool {
	return o != nil && o.FilterQuery != nil
}

// SetFilterQuery gets a reference to the given string and assigns it to the FilterQuery field.
func (o *RumSegmentReferenceTable) SetFilterQuery(v string) {
	o.FilterQuery = &v
}

// GetJoinCondition returns the JoinCondition field value.
func (o *RumSegmentReferenceTable) GetJoinCondition() RumSegmentReferenceTableJoinCondition {
	if o == nil {
		var ret RumSegmentReferenceTableJoinCondition
		return ret
	}
	return o.JoinCondition
}

// GetJoinConditionOk returns a tuple with the JoinCondition field value
// and a boolean to check if the value has been set.
func (o *RumSegmentReferenceTable) GetJoinConditionOk() (*RumSegmentReferenceTableJoinCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinCondition, true
}

// SetJoinCondition sets field value.
func (o *RumSegmentReferenceTable) SetJoinCondition(v RumSegmentReferenceTableJoinCondition) {
	o.JoinCondition = v
}

// GetName returns the Name field value.
func (o *RumSegmentReferenceTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RumSegmentReferenceTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RumSegmentReferenceTable) SetName(v string) {
	o.Name = v
}

// GetTableName returns the TableName field value.
func (o *RumSegmentReferenceTable) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *RumSegmentReferenceTable) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *RumSegmentReferenceTable) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSegmentReferenceTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["columns"] = o.Columns
	if o.FilterQuery != nil {
		toSerialize["filter_query"] = o.FilterQuery
	}
	toSerialize["join_condition"] = o.JoinCondition
	toSerialize["name"] = o.Name
	toSerialize["table_name"] = o.TableName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumSegmentReferenceTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns       *[]RumSegmentReferenceTableColumn      `json:"columns"`
		FilterQuery   *string                                `json:"filter_query,omitempty"`
		JoinCondition *RumSegmentReferenceTableJoinCondition `json:"join_condition"`
		Name          *string                                `json:"name"`
		TableName     *string                                `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Columns == nil {
		return fmt.Errorf("required field columns missing")
	}
	if all.JoinCondition == nil {
		return fmt.Errorf("required field join_condition missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"columns", "filter_query", "join_condition", "name", "table_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Columns = *all.Columns
	o.FilterQuery = all.FilterQuery
	if all.JoinCondition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinCondition = *all.JoinCondition
	o.Name = *all.Name
	o.TableName = *all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
