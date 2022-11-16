// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DashboardReportDestination Report destination-specific configuration. This determines how reports are sent. Only email destinations are supported.
type DashboardReportDestination struct {
	// Email destinations for a report.
	Email *DashboardReportDestinationEmail `json:"email,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportDestination instantiates a new DashboardReportDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportDestination() *DashboardReportDestination {
	this := DashboardReportDestination{}
	return &this
}

// NewDashboardReportDestinationWithDefaults instantiates a new DashboardReportDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportDestinationWithDefaults() *DashboardReportDestination {
	this := DashboardReportDestination{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DashboardReportDestination) GetEmail() DashboardReportDestinationEmail {
	if o == nil || o.Email == nil {
		var ret DashboardReportDestinationEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportDestination) GetEmailOk() (*DashboardReportDestinationEmail, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DashboardReportDestination) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given DashboardReportDestinationEmail and assigns it to the Email field.
func (o *DashboardReportDestination) SetEmail(v DashboardReportDestinationEmail) {
	o.Email = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportDestination) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Email *DashboardReportDestinationEmail `json:"email,omitempty"`
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
	if all.Email != nil && all.Email.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Email = all.Email
	return nil
}
