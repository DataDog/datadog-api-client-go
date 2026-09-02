// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionDataQualityEntityMetricConfig Measure configuration for one side of a source to target comparison.
type MonitorFormulaAndFunctionDataQualityEntityMetricConfig struct {
	// Custom SQL query used to compute the measure for this entity.
	CustomSql *string `json:"custom_sql,omitempty"`
	// Custom WHERE clause applied when computing the measure for this entity.
	CustomWhere *string `json:"custom_where,omitempty"`
	// Identifier of the data entity to measure.
	EntityId string `json:"entity_id"`
	// Type of the data entity to measure.
	EntityType string `json:"entity_type"`
	// Columns to group results by when computing the measure for this entity.
	GroupByColumns []string `json:"group_by_columns,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionDataQualityEntityMetricConfig instantiates a new MonitorFormulaAndFunctionDataQualityEntityMetricConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionDataQualityEntityMetricConfig(entityId string, entityType string) *MonitorFormulaAndFunctionDataQualityEntityMetricConfig {
	this := MonitorFormulaAndFunctionDataQualityEntityMetricConfig{}
	this.EntityId = entityId
	this.EntityType = entityType
	return &this
}

// NewMonitorFormulaAndFunctionDataQualityEntityMetricConfigWithDefaults instantiates a new MonitorFormulaAndFunctionDataQualityEntityMetricConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionDataQualityEntityMetricConfigWithDefaults() *MonitorFormulaAndFunctionDataQualityEntityMetricConfig {
	this := MonitorFormulaAndFunctionDataQualityEntityMetricConfig{}
	return &this
}

// GetCustomSql returns the CustomSql field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetCustomSql() string {
	if o == nil || o.CustomSql == nil {
		var ret string
		return ret
	}
	return *o.CustomSql
}

// GetCustomSqlOk returns a tuple with the CustomSql field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetCustomSqlOk() (*string, bool) {
	if o == nil || o.CustomSql == nil {
		return nil, false
	}
	return o.CustomSql, true
}

// HasCustomSql returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) HasCustomSql() bool {
	return o != nil && o.CustomSql != nil
}

// SetCustomSql gets a reference to the given string and assigns it to the CustomSql field.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) SetCustomSql(v string) {
	o.CustomSql = &v
}

// GetCustomWhere returns the CustomWhere field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetCustomWhere() string {
	if o == nil || o.CustomWhere == nil {
		var ret string
		return ret
	}
	return *o.CustomWhere
}

// GetCustomWhereOk returns a tuple with the CustomWhere field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetCustomWhereOk() (*string, bool) {
	if o == nil || o.CustomWhere == nil {
		return nil, false
	}
	return o.CustomWhere, true
}

// HasCustomWhere returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) HasCustomWhere() bool {
	return o != nil && o.CustomWhere != nil
}

// SetCustomWhere gets a reference to the given string and assigns it to the CustomWhere field.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) SetCustomWhere(v string) {
	o.CustomWhere = &v
}

// GetEntityId returns the EntityId field value.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) SetEntityType(v string) {
	o.EntityType = v
}

// GetGroupByColumns returns the GroupByColumns field value if set, zero value otherwise.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetGroupByColumns() []string {
	if o == nil || o.GroupByColumns == nil {
		var ret []string
		return ret
	}
	return o.GroupByColumns
}

// GetGroupByColumnsOk returns a tuple with the GroupByColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) GetGroupByColumnsOk() (*[]string, bool) {
	if o == nil || o.GroupByColumns == nil {
		return nil, false
	}
	return &o.GroupByColumns, true
}

// HasGroupByColumns returns a boolean if a field has been set.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) HasGroupByColumns() bool {
	return o != nil && o.GroupByColumns != nil
}

// SetGroupByColumns gets a reference to the given []string and assigns it to the GroupByColumns field.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) SetGroupByColumns(v []string) {
	o.GroupByColumns = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionDataQualityEntityMetricConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomSql != nil {
		toSerialize["custom_sql"] = o.CustomSql
	}
	if o.CustomWhere != nil {
		toSerialize["custom_where"] = o.CustomWhere
	}
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	if o.GroupByColumns != nil {
		toSerialize["group_by_columns"] = o.GroupByColumns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionDataQualityEntityMetricConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomSql      *string  `json:"custom_sql,omitempty"`
		CustomWhere    *string  `json:"custom_where,omitempty"`
		EntityId       *string  `json:"entity_id"`
		EntityType     *string  `json:"entity_type"`
		GroupByColumns []string `json:"group_by_columns,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EntityId == nil {
		return fmt.Errorf("required field entity_id missing")
	}
	if all.EntityType == nil {
		return fmt.Errorf("required field entity_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_sql", "custom_where", "entity_id", "entity_type", "group_by_columns"})
	} else {
		return err
	}
	o.CustomSql = all.CustomSql
	o.CustomWhere = all.CustomWhere
	o.EntityId = *all.EntityId
	o.EntityType = *all.EntityType
	o.GroupByColumns = all.GroupByColumns

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
