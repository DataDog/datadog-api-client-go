// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SegmentDataAttributesDataQueryReferenceTableItems
type SegmentDataAttributesDataQueryReferenceTableItems struct {
	//
	Columns []SegmentDataAttributesDataQueryReferenceTableItemsColumnsItems `json:"columns,omitempty"`
	//
	FilterQuery *string `json:"filter_query,omitempty"`
	//
	JoinCondition SegmentDataAttributesDataQueryReferenceTableItemsJoinCondition `json:"join_condition"`
	//
	Name *string `json:"name,omitempty"`
	//
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSegmentDataAttributesDataQueryReferenceTableItems instantiates a new SegmentDataAttributesDataQueryReferenceTableItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSegmentDataAttributesDataQueryReferenceTableItems(joinCondition SegmentDataAttributesDataQueryReferenceTableItemsJoinCondition, tableName string) *SegmentDataAttributesDataQueryReferenceTableItems {
	this := SegmentDataAttributesDataQueryReferenceTableItems{}
	this.JoinCondition = joinCondition
	this.TableName = tableName
	return &this
}

// NewSegmentDataAttributesDataQueryReferenceTableItemsWithDefaults instantiates a new SegmentDataAttributesDataQueryReferenceTableItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSegmentDataAttributesDataQueryReferenceTableItemsWithDefaults() *SegmentDataAttributesDataQueryReferenceTableItems {
	this := SegmentDataAttributesDataQueryReferenceTableItems{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetColumns() []SegmentDataAttributesDataQueryReferenceTableItemsColumnsItems {
	if o == nil || o.Columns == nil {
		var ret []SegmentDataAttributesDataQueryReferenceTableItemsColumnsItems
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetColumnsOk() (*[]SegmentDataAttributesDataQueryReferenceTableItemsColumnsItems, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []SegmentDataAttributesDataQueryReferenceTableItemsColumnsItems and assigns it to the Columns field.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) SetColumns(v []SegmentDataAttributesDataQueryReferenceTableItemsColumnsItems) {
	o.Columns = v
}

// GetFilterQuery returns the FilterQuery field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetFilterQuery() string {
	if o == nil || o.FilterQuery == nil {
		var ret string
		return ret
	}
	return *o.FilterQuery
}

// GetFilterQueryOk returns a tuple with the FilterQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetFilterQueryOk() (*string, bool) {
	if o == nil || o.FilterQuery == nil {
		return nil, false
	}
	return o.FilterQuery, true
}

// HasFilterQuery returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) HasFilterQuery() bool {
	return o != nil && o.FilterQuery != nil
}

// SetFilterQuery gets a reference to the given string and assigns it to the FilterQuery field.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) SetFilterQuery(v string) {
	o.FilterQuery = &v
}

// GetJoinCondition returns the JoinCondition field value.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetJoinCondition() SegmentDataAttributesDataQueryReferenceTableItemsJoinCondition {
	if o == nil {
		var ret SegmentDataAttributesDataQueryReferenceTableItemsJoinCondition
		return ret
	}
	return o.JoinCondition
}

// GetJoinConditionOk returns a tuple with the JoinCondition field value
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetJoinConditionOk() (*SegmentDataAttributesDataQueryReferenceTableItemsJoinCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinCondition, true
}

// SetJoinCondition sets field value.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) SetJoinCondition(v SegmentDataAttributesDataQueryReferenceTableItemsJoinCondition) {
	o.JoinCondition = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) SetName(v string) {
	o.Name = &v
}

// GetTableName returns the TableName field value.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *SegmentDataAttributesDataQueryReferenceTableItems) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SegmentDataAttributesDataQueryReferenceTableItems) MarshalJSON() ([]byte, error) {
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
func (o *SegmentDataAttributesDataQueryReferenceTableItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Columns       []SegmentDataAttributesDataQueryReferenceTableItemsColumnsItems `json:"columns,omitempty"`
		FilterQuery   *string                                                         `json:"filter_query,omitempty"`
		JoinCondition *SegmentDataAttributesDataQueryReferenceTableItemsJoinCondition `json:"join_condition"`
		Name          *string                                                         `json:"name,omitempty"`
		TableName     *string                                                         `json:"table_name"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
