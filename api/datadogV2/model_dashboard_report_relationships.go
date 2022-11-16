// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DashboardReportRelationships Relationships of a dashboard report.
type DashboardReportRelationships struct {
	// Dashboard associated with the report.
	Dashboard *DashboardReportRelationshipsDashboard `json:"dashboard,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportRelationships instantiates a new DashboardReportRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportRelationships() *DashboardReportRelationships {
	this := DashboardReportRelationships{}
	return &this
}

// NewDashboardReportRelationshipsWithDefaults instantiates a new DashboardReportRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportRelationshipsWithDefaults() *DashboardReportRelationships {
	this := DashboardReportRelationships{}
	return &this
}

// GetDashboard returns the Dashboard field value if set, zero value otherwise.
func (o *DashboardReportRelationships) GetDashboard() DashboardReportRelationshipsDashboard {
	if o == nil || o.Dashboard == nil {
		var ret DashboardReportRelationshipsDashboard
		return ret
	}
	return *o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportRelationships) GetDashboardOk() (*DashboardReportRelationshipsDashboard, bool) {
	if o == nil || o.Dashboard == nil {
		return nil, false
	}
	return o.Dashboard, true
}

// HasDashboard returns a boolean if a field has been set.
func (o *DashboardReportRelationships) HasDashboard() bool {
	return o != nil && o.Dashboard != nil
}

// SetDashboard gets a reference to the given DashboardReportRelationshipsDashboard and assigns it to the Dashboard field.
func (o *DashboardReportRelationships) SetDashboard(v DashboardReportRelationshipsDashboard) {
	o.Dashboard = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Dashboard != nil {
		toSerialize["dashboard"] = o.Dashboard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportRelationships) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Dashboard *DashboardReportRelationshipsDashboard `json:"dashboard,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if all.Dashboard != nil && all.Dashboard.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Dashboard = all.Dashboard
	return nil
}
