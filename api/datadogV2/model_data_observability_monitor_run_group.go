// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataObservabilityMonitorRunGroup A monitor group that is not ok after a run.
type DataObservabilityMonitorRunGroup struct {
	// The account of the entity, when the platform has one.
	Account *string `json:"account,omitempty"`
	// The column of the entity, when it is a column.
	Column *string `json:"column,omitempty"`
	// The database of the entity.
	Database string `json:"database"`
	// The ID of the table or column the group checks.
	EntityId string `json:"entity_id"`
	// The type of the entity, such as table or column.
	EntityType string `json:"entity_type"`
	// The group's name in Datadog monitors, as the monitor API and UI show it. Pass it as `filter[group]` to list the group's evaluations.
	Group string `json:"group"`
	// The values of the monitor's group-by columns for this group. `N/A` means the column had no value.
	GroupByValues map[string]string `json:"group_by_values,omitempty"`
	// The data platform of the entity.
	Platform string `json:"platform"`
	// The schema of the entity, when it has one.
	Schema *string `json:"schema,omitempty"`
	// The state of a monitor group, from the alerting engine.
	Status DataObservabilityMonitorGroupStatus `json:"status"`
	// The table of the entity, when it is a table or column.
	Table *string `json:"table,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataObservabilityMonitorRunGroup instantiates a new DataObservabilityMonitorRunGroup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataObservabilityMonitorRunGroup(database string, entityId string, entityType string, group string, platform string, status DataObservabilityMonitorGroupStatus) *DataObservabilityMonitorRunGroup {
	this := DataObservabilityMonitorRunGroup{}
	this.Database = database
	this.EntityId = entityId
	this.EntityType = entityType
	this.Group = group
	this.Platform = platform
	this.Status = status
	return &this
}

// NewDataObservabilityMonitorRunGroupWithDefaults instantiates a new DataObservabilityMonitorRunGroup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataObservabilityMonitorRunGroupWithDefaults() *DataObservabilityMonitorRunGroup {
	this := DataObservabilityMonitorRunGroup{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *DataObservabilityMonitorRunGroup) GetAccount() string {
	if o == nil || o.Account == nil {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetAccountOk() (*string, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *DataObservabilityMonitorRunGroup) HasAccount() bool {
	return o != nil && o.Account != nil
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *DataObservabilityMonitorRunGroup) SetAccount(v string) {
	o.Account = &v
}

// GetColumn returns the Column field value if set, zero value otherwise.
func (o *DataObservabilityMonitorRunGroup) GetColumn() string {
	if o == nil || o.Column == nil {
		var ret string
		return ret
	}
	return *o.Column
}

// GetColumnOk returns a tuple with the Column field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetColumnOk() (*string, bool) {
	if o == nil || o.Column == nil {
		return nil, false
	}
	return o.Column, true
}

// HasColumn returns a boolean if a field has been set.
func (o *DataObservabilityMonitorRunGroup) HasColumn() bool {
	return o != nil && o.Column != nil
}

// SetColumn gets a reference to the given string and assigns it to the Column field.
func (o *DataObservabilityMonitorRunGroup) SetColumn(v string) {
	o.Column = &v
}

// GetDatabase returns the Database field value.
func (o *DataObservabilityMonitorRunGroup) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *DataObservabilityMonitorRunGroup) SetDatabase(v string) {
	o.Database = v
}

// GetEntityId returns the EntityId field value.
func (o *DataObservabilityMonitorRunGroup) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value.
func (o *DataObservabilityMonitorRunGroup) SetEntityId(v string) {
	o.EntityId = v
}

// GetEntityType returns the EntityType field value.
func (o *DataObservabilityMonitorRunGroup) GetEntityType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetEntityTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityType, true
}

// SetEntityType sets field value.
func (o *DataObservabilityMonitorRunGroup) SetEntityType(v string) {
	o.EntityType = v
}

// GetGroup returns the Group field value.
func (o *DataObservabilityMonitorRunGroup) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value.
func (o *DataObservabilityMonitorRunGroup) SetGroup(v string) {
	o.Group = v
}

// GetGroupByValues returns the GroupByValues field value if set, zero value otherwise.
func (o *DataObservabilityMonitorRunGroup) GetGroupByValues() map[string]string {
	if o == nil || o.GroupByValues == nil {
		var ret map[string]string
		return ret
	}
	return o.GroupByValues
}

// GetGroupByValuesOk returns a tuple with the GroupByValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetGroupByValuesOk() (*map[string]string, bool) {
	if o == nil || o.GroupByValues == nil {
		return nil, false
	}
	return &o.GroupByValues, true
}

// HasGroupByValues returns a boolean if a field has been set.
func (o *DataObservabilityMonitorRunGroup) HasGroupByValues() bool {
	return o != nil && o.GroupByValues != nil
}

// SetGroupByValues gets a reference to the given map[string]string and assigns it to the GroupByValues field.
func (o *DataObservabilityMonitorRunGroup) SetGroupByValues(v map[string]string) {
	o.GroupByValues = v
}

// GetPlatform returns the Platform field value.
func (o *DataObservabilityMonitorRunGroup) GetPlatform() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Platform, true
}

// SetPlatform sets field value.
func (o *DataObservabilityMonitorRunGroup) SetPlatform(v string) {
	o.Platform = v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DataObservabilityMonitorRunGroup) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DataObservabilityMonitorRunGroup) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DataObservabilityMonitorRunGroup) SetSchema(v string) {
	o.Schema = &v
}

// GetStatus returns the Status field value.
func (o *DataObservabilityMonitorRunGroup) GetStatus() DataObservabilityMonitorGroupStatus {
	if o == nil {
		var ret DataObservabilityMonitorGroupStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetStatusOk() (*DataObservabilityMonitorGroupStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *DataObservabilityMonitorRunGroup) SetStatus(v DataObservabilityMonitorGroupStatus) {
	o.Status = v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *DataObservabilityMonitorRunGroup) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataObservabilityMonitorRunGroup) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *DataObservabilityMonitorRunGroup) HasTable() bool {
	return o != nil && o.Table != nil
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *DataObservabilityMonitorRunGroup) SetTable(v string) {
	o.Table = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataObservabilityMonitorRunGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	if o.Column != nil {
		toSerialize["column"] = o.Column
	}
	toSerialize["database"] = o.Database
	toSerialize["entity_id"] = o.EntityId
	toSerialize["entity_type"] = o.EntityType
	toSerialize["group"] = o.Group
	if o.GroupByValues != nil {
		toSerialize["group_by_values"] = o.GroupByValues
	}
	toSerialize["platform"] = o.Platform
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	toSerialize["status"] = o.Status
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataObservabilityMonitorRunGroup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Account       *string                              `json:"account,omitempty"`
		Column        *string                              `json:"column,omitempty"`
		Database      *string                              `json:"database"`
		EntityId      *string                              `json:"entity_id"`
		EntityType    *string                              `json:"entity_type"`
		Group         *string                              `json:"group"`
		GroupByValues map[string]string                    `json:"group_by_values,omitempty"`
		Platform      *string                              `json:"platform"`
		Schema        *string                              `json:"schema,omitempty"`
		Status        *DataObservabilityMonitorGroupStatus `json:"status"`
		Table         *string                              `json:"table,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.EntityId == nil {
		return fmt.Errorf("required field entity_id missing")
	}
	if all.EntityType == nil {
		return fmt.Errorf("required field entity_type missing")
	}
	if all.Group == nil {
		return fmt.Errorf("required field group missing")
	}
	if all.Platform == nil {
		return fmt.Errorf("required field platform missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account", "column", "database", "entity_id", "entity_type", "group", "group_by_values", "platform", "schema", "status", "table"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Account = all.Account
	o.Column = all.Column
	o.Database = *all.Database
	o.EntityId = *all.EntityId
	o.EntityType = *all.EntityType
	o.Group = *all.Group
	o.GroupByValues = all.GroupByValues
	o.Platform = *all.Platform
	o.Schema = all.Schema
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.Table = all.Table

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
