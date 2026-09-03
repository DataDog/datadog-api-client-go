// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatabricksIntegrationAccountSettingsRequest Settings for creating the Databricks integration account.
type DatabricksIntegrationAccountSettingsRequest struct {
	// ID of the SQL warehouse used to query the Databricks system tables.
	SystemTablesSqlWarehouseId *string `json:"system_tables_sql_warehouse_id,omitempty"`
	// URL of the Databricks workspace.
	WorkspaceUrl string `json:"workspace_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDatabricksIntegrationAccountSettingsRequest instantiates a new DatabricksIntegrationAccountSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatabricksIntegrationAccountSettingsRequest(workspaceUrl string) *DatabricksIntegrationAccountSettingsRequest {
	this := DatabricksIntegrationAccountSettingsRequest{}
	this.WorkspaceUrl = workspaceUrl
	return &this
}

// NewDatabricksIntegrationAccountSettingsRequestWithDefaults instantiates a new DatabricksIntegrationAccountSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatabricksIntegrationAccountSettingsRequestWithDefaults() *DatabricksIntegrationAccountSettingsRequest {
	this := DatabricksIntegrationAccountSettingsRequest{}
	return &this
}

// GetSystemTablesSqlWarehouseId returns the SystemTablesSqlWarehouseId field value if set, zero value otherwise.
func (o *DatabricksIntegrationAccountSettingsRequest) GetSystemTablesSqlWarehouseId() string {
	if o == nil || o.SystemTablesSqlWarehouseId == nil {
		var ret string
		return ret
	}
	return *o.SystemTablesSqlWarehouseId
}

// GetSystemTablesSqlWarehouseIdOk returns a tuple with the SystemTablesSqlWarehouseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountSettingsRequest) GetSystemTablesSqlWarehouseIdOk() (*string, bool) {
	if o == nil || o.SystemTablesSqlWarehouseId == nil {
		return nil, false
	}
	return o.SystemTablesSqlWarehouseId, true
}

// HasSystemTablesSqlWarehouseId returns a boolean if a field has been set.
func (o *DatabricksIntegrationAccountSettingsRequest) HasSystemTablesSqlWarehouseId() bool {
	return o != nil && o.SystemTablesSqlWarehouseId != nil
}

// SetSystemTablesSqlWarehouseId gets a reference to the given string and assigns it to the SystemTablesSqlWarehouseId field.
func (o *DatabricksIntegrationAccountSettingsRequest) SetSystemTablesSqlWarehouseId(v string) {
	o.SystemTablesSqlWarehouseId = &v
}

// GetWorkspaceUrl returns the WorkspaceUrl field value.
func (o *DatabricksIntegrationAccountSettingsRequest) GetWorkspaceUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.WorkspaceUrl
}

// GetWorkspaceUrlOk returns a tuple with the WorkspaceUrl field value
// and a boolean to check if the value has been set.
func (o *DatabricksIntegrationAccountSettingsRequest) GetWorkspaceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceUrl, true
}

// SetWorkspaceUrl sets field value.
func (o *DatabricksIntegrationAccountSettingsRequest) SetWorkspaceUrl(v string) {
	o.WorkspaceUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatabricksIntegrationAccountSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.SystemTablesSqlWarehouseId != nil {
		toSerialize["system_tables_sql_warehouse_id"] = o.SystemTablesSqlWarehouseId
	}
	toSerialize["workspace_url"] = o.WorkspaceUrl
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatabricksIntegrationAccountSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SystemTablesSqlWarehouseId *string `json:"system_tables_sql_warehouse_id,omitempty"`
		WorkspaceUrl               *string `json:"workspace_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.WorkspaceUrl == nil {
		return fmt.Errorf("required field workspace_url missing")
	}
	o.SystemTablesSqlWarehouseId = all.SystemTablesSqlWarehouseId
	o.WorkspaceUrl = *all.WorkspaceUrl

	return nil
}
