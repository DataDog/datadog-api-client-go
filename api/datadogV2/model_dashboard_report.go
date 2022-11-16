// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DashboardReport Dashboard report schema.
type DashboardReport struct {
	// Attributes for the dashboard report schema.
	Attributes *DashboardReportAttributes `json:"attributes,omitempty"`
	// ID of the dashboard report.
	Id string `json:"id"`
	// Relationships of a dashboard report.
	Relationships *DashboardReportRelationships `json:"relationships,omitempty"`
	// JSON:API type of the dashboard report.
	Type DashboardReportType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReport instantiates a new DashboardReport object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReport(id string, typeVar DashboardReportType) *DashboardReport {
	this := DashboardReport{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewDashboardReportWithDefaults instantiates a new DashboardReport object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportWithDefaults() *DashboardReport {
	this := DashboardReport{}
	var typeVar DashboardReportType = DASHBOARDREPORTTYPE_DASHBOARD_REPORTS_TYPE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DashboardReport) GetAttributes() DashboardReportAttributes {
	if o == nil || o.Attributes == nil {
		var ret DashboardReportAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReport) GetAttributesOk() (*DashboardReportAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DashboardReport) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DashboardReportAttributes and assigns it to the Attributes field.
func (o *DashboardReport) SetAttributes(v DashboardReportAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *DashboardReport) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardReport) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DashboardReport) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *DashboardReport) GetRelationships() DashboardReportRelationships {
	if o == nil || o.Relationships == nil {
		var ret DashboardReportRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReport) GetRelationshipsOk() (*DashboardReportRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *DashboardReport) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given DashboardReportRelationships and assigns it to the Relationships field.
func (o *DashboardReport) SetRelationships(v DashboardReportRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *DashboardReport) GetType() DashboardReportType {
	if o == nil {
		var ret DashboardReportType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardReport) GetTypeOk() (*DashboardReportType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DashboardReport) SetType(v DashboardReportType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReport) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Id   *string              `json:"id"`
		Type *DashboardReportType `json:"type"`
	}{}
	all := struct {
		Attributes    *DashboardReportAttributes    `json:"attributes,omitempty"`
		Id            string                        `json:"id"`
		Relationships *DashboardReportRelationships `json:"relationships,omitempty"`
		Type          DashboardReportType           `json:"type"`
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Relationships = all.Relationships
	o.Type = all.Type
	return nil
}
