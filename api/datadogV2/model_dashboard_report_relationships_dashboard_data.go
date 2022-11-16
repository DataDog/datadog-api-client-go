// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReportRelationshipsDashboardData JSON:API data representing a report's associated dashboard.
type DashboardReportRelationshipsDashboardData struct {
	// ID of a report's associated dashboard. Screenboards (dashboards with layout_type=FREE) are not supported.
	Id string `json:"id"`
	// JSON:API type of the report's associated dashboard.
	Type DashboardReportResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportRelationshipsDashboardData instantiates a new DashboardReportRelationshipsDashboardData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportRelationshipsDashboardData(id string, typeVar DashboardReportResourceType) *DashboardReportRelationshipsDashboardData {
	this := DashboardReportRelationshipsDashboardData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDashboardReportRelationshipsDashboardDataWithDefaults instantiates a new DashboardReportRelationshipsDashboardData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportRelationshipsDashboardDataWithDefaults() *DashboardReportRelationshipsDashboardData {
	this := DashboardReportRelationshipsDashboardData{}
	var typeVar DashboardReportResourceType = DASHBOARDREPORTRESOURCETYPE_DASHBOARDS_TYPE
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *DashboardReportRelationshipsDashboardData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardReportRelationshipsDashboardData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DashboardReportRelationshipsDashboardData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *DashboardReportRelationshipsDashboardData) GetType() DashboardReportResourceType {
	if o == nil {
		var ret DashboardReportResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardReportRelationshipsDashboardData) GetTypeOk() (*DashboardReportResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DashboardReportRelationshipsDashboardData) SetType(v DashboardReportResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportRelationshipsDashboardData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportRelationshipsDashboardData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Id   *string                      `json:"id"`
		Type *DashboardReportResourceType `json:"type"`
	}{}
	all := struct {
		Id   string                      `json:"id"`
		Type DashboardReportResourceType `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Id = all.Id
	o.Type = all.Type
	return nil
}
