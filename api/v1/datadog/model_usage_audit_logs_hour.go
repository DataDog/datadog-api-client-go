/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package datadog

import (
	"encoding/json"
	"time"
)

// UsageAuditLogsHour Audit logs usage for a given organization for a given hour.
type UsageAuditLogsHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The total number of audit logs lines indexed during a given hour.
	LinesIndexed *int64 `json:"lines_indexed,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:-`
	AdditionalProperties map[string]interface{}
}

// NewUsageAuditLogsHour instantiates a new UsageAuditLogsHour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageAuditLogsHour() *UsageAuditLogsHour {
	this := UsageAuditLogsHour{}
	return &this
}

// NewUsageAuditLogsHourWithDefaults instantiates a new UsageAuditLogsHour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageAuditLogsHourWithDefaults() *UsageAuditLogsHour {
	this := UsageAuditLogsHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageAuditLogsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAuditLogsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageAuditLogsHour) HasHour() bool {
	if o != nil && o.Hour != nil {
		return true
	}

	return false
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageAuditLogsHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetLinesIndexed returns the LinesIndexed field value if set, zero value otherwise.
func (o *UsageAuditLogsHour) GetLinesIndexed() int64 {
	if o == nil || o.LinesIndexed == nil {
		var ret int64
		return ret
	}
	return *o.LinesIndexed
}

// GetLinesIndexedOk returns a tuple with the LinesIndexed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAuditLogsHour) GetLinesIndexedOk() (*int64, bool) {
	if o == nil || o.LinesIndexed == nil {
		return nil, false
	}
	return o.LinesIndexed, true
}

// HasLinesIndexed returns a boolean if a field has been set.
func (o *UsageAuditLogsHour) HasLinesIndexed() bool {
	if o != nil && o.LinesIndexed != nil {
		return true
	}

	return false
}

// SetLinesIndexed gets a reference to the given int64 and assigns it to the LinesIndexed field.
func (o *UsageAuditLogsHour) SetLinesIndexed(v int64) {
	o.LinesIndexed = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageAuditLogsHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAuditLogsHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageAuditLogsHour) HasOrgName() bool {
	if o != nil && o.OrgName != nil {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageAuditLogsHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageAuditLogsHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAuditLogsHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageAuditLogsHour) HasPublicId() bool {
	if o != nil && o.PublicId != nil {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageAuditLogsHour) SetPublicId(v string) {
	o.PublicId = &v
}

func (o UsageAuditLogsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LinesIndexed != nil {
		toSerialize["lines_indexed"] = o.LinesIndexed
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

func (o *UsageAuditLogsHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Hour         *time.Time `json:"hour,omitempty"`
		LinesIndexed *int64     `json:"lines_indexed,omitempty"`
		OrgName      *string    `json:"org_name,omitempty"`
		PublicId     *string    `json:"public_id,omitempty"`
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
	o.Hour = all.Hour
	o.LinesIndexed = all.LinesIndexed
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	return nil
}
