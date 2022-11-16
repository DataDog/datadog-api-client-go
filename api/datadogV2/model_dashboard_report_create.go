// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReportCreate Schema for request body to create a dashboard report.
type DashboardReportCreate struct {
	// Attributes for the dashboard report schema used to create a dashboard report.
	Attributes *DashboardReportCreateAttributes `json:"attributes,omitempty"`
	// JSON:API type of the dashboard report.
	Type DashboardReportType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportCreate instantiates a new DashboardReportCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportCreate(typeVar DashboardReportType) *DashboardReportCreate {
	this := DashboardReportCreate{}
	this.Type = typeVar
	return &this
}

// NewDashboardReportCreateWithDefaults instantiates a new DashboardReportCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportCreateWithDefaults() *DashboardReportCreate {
	this := DashboardReportCreate{}
	var typeVar DashboardReportType = DASHBOARDREPORTTYPE_DASHBOARD_REPORTS_TYPE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DashboardReportCreate) GetAttributes() DashboardReportCreateAttributes {
	if o == nil || o.Attributes == nil {
		var ret DashboardReportCreateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportCreate) GetAttributesOk() (*DashboardReportCreateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DashboardReportCreate) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DashboardReportCreateAttributes and assigns it to the Attributes field.
func (o *DashboardReportCreate) SetAttributes(v DashboardReportCreateAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *DashboardReportCreate) GetType() DashboardReportType {
	if o == nil {
		var ret DashboardReportType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardReportCreate) GetTypeOk() (*DashboardReportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DashboardReportCreate) SetType(v DashboardReportType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportCreate) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Type *DashboardReportType `json:"type"`
	}{}
	all := struct {
		Attributes *DashboardReportCreateAttributes `json:"attributes,omitempty"`
		Type       DashboardReportType              `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Type = all.Type
	return nil
}
