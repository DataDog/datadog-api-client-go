// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSDSHour Sensitive Data Scanner usage for a given organization for a given hour.
type UsageSDSHour struct {
	// The total number of bytes scanned of APM usage across all usage types by the Sensitive Data Scanner from the start of the given hour’s month until the given hour.
	ApmScannedBytes *int64 `json:"apm_scanned_bytes,omitempty"`
	// The total number of bytes scanned of Events usage across all usage types by the Sensitive Data Scanner from the start of the given hour’s month until the given hour.
	EventsScannedBytes *int64 `json:"events_scanned_bytes,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The total number of bytes scanned of logs usage by the Sensitive Data Scanner from the start of the given hour’s month until the given hour.
	LogsScannedBytes *int64 `json:"logs_scanned_bytes,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// The total number of bytes scanned of RUM usage across all usage types by the Sensitive Data Scanner from the start of the given hour’s month until the given hour.
	RumScannedBytes *int64 `json:"rum_scanned_bytes,omitempty"`
	// The total number of bytes scanned across all usage types by the Sensitive Data Scanner from the start of the given hour’s month until the given hour.
	TotalScannedBytes *int64 `json:"total_scanned_bytes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageSDSHour instantiates a new UsageSDSHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSDSHour() *UsageSDSHour {
	this := UsageSDSHour{}
	return &this
}

// NewUsageSDSHourWithDefaults instantiates a new UsageSDSHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSDSHourWithDefaults() *UsageSDSHour {
	this := UsageSDSHour{}
	return &this
}

// GetApmScannedBytes returns the ApmScannedBytes field value if set, zero value otherwise.
func (o *UsageSDSHour) GetApmScannedBytes() int64 {
	if o == nil || o.ApmScannedBytes == nil {
		var ret int64
		return ret
	}
	return *o.ApmScannedBytes
}

// GetApmScannedBytesOk returns a tuple with the ApmScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetApmScannedBytesOk() (*int64, bool) {
	if o == nil || o.ApmScannedBytes == nil {
		return nil, false
	}
	return o.ApmScannedBytes, true
}

// HasApmScannedBytes returns a boolean if a field has been set.
func (o *UsageSDSHour) HasApmScannedBytes() bool {
	return o != nil && o.ApmScannedBytes != nil
}

// SetApmScannedBytes gets a reference to the given int64 and assigns it to the ApmScannedBytes field.
func (o *UsageSDSHour) SetApmScannedBytes(v int64) {
	o.ApmScannedBytes = &v
}

// GetEventsScannedBytes returns the EventsScannedBytes field value if set, zero value otherwise.
func (o *UsageSDSHour) GetEventsScannedBytes() int64 {
	if o == nil || o.EventsScannedBytes == nil {
		var ret int64
		return ret
	}
	return *o.EventsScannedBytes
}

// GetEventsScannedBytesOk returns a tuple with the EventsScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetEventsScannedBytesOk() (*int64, bool) {
	if o == nil || o.EventsScannedBytes == nil {
		return nil, false
	}
	return o.EventsScannedBytes, true
}

// HasEventsScannedBytes returns a boolean if a field has been set.
func (o *UsageSDSHour) HasEventsScannedBytes() bool {
	return o != nil && o.EventsScannedBytes != nil
}

// SetEventsScannedBytes gets a reference to the given int64 and assigns it to the EventsScannedBytes field.
func (o *UsageSDSHour) SetEventsScannedBytes(v int64) {
	o.EventsScannedBytes = &v
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageSDSHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageSDSHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageSDSHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetLogsScannedBytes returns the LogsScannedBytes field value if set, zero value otherwise.
func (o *UsageSDSHour) GetLogsScannedBytes() int64 {
	if o == nil || o.LogsScannedBytes == nil {
		var ret int64
		return ret
	}
	return *o.LogsScannedBytes
}

// GetLogsScannedBytesOk returns a tuple with the LogsScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetLogsScannedBytesOk() (*int64, bool) {
	if o == nil || o.LogsScannedBytes == nil {
		return nil, false
	}
	return o.LogsScannedBytes, true
}

// HasLogsScannedBytes returns a boolean if a field has been set.
func (o *UsageSDSHour) HasLogsScannedBytes() bool {
	return o != nil && o.LogsScannedBytes != nil
}

// SetLogsScannedBytes gets a reference to the given int64 and assigns it to the LogsScannedBytes field.
func (o *UsageSDSHour) SetLogsScannedBytes(v int64) {
	o.LogsScannedBytes = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageSDSHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageSDSHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageSDSHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageSDSHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageSDSHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageSDSHour) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRumScannedBytes returns the RumScannedBytes field value if set, zero value otherwise.
func (o *UsageSDSHour) GetRumScannedBytes() int64 {
	if o == nil || o.RumScannedBytes == nil {
		var ret int64
		return ret
	}
	return *o.RumScannedBytes
}

// GetRumScannedBytesOk returns a tuple with the RumScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetRumScannedBytesOk() (*int64, bool) {
	if o == nil || o.RumScannedBytes == nil {
		return nil, false
	}
	return o.RumScannedBytes, true
}

// HasRumScannedBytes returns a boolean if a field has been set.
func (o *UsageSDSHour) HasRumScannedBytes() bool {
	return o != nil && o.RumScannedBytes != nil
}

// SetRumScannedBytes gets a reference to the given int64 and assigns it to the RumScannedBytes field.
func (o *UsageSDSHour) SetRumScannedBytes(v int64) {
	o.RumScannedBytes = &v
}

// GetTotalScannedBytes returns the TotalScannedBytes field value if set, zero value otherwise.
func (o *UsageSDSHour) GetTotalScannedBytes() int64 {
	if o == nil || o.TotalScannedBytes == nil {
		var ret int64
		return ret
	}
	return *o.TotalScannedBytes
}

// GetTotalScannedBytesOk returns a tuple with the TotalScannedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSDSHour) GetTotalScannedBytesOk() (*int64, bool) {
	if o == nil || o.TotalScannedBytes == nil {
		return nil, false
	}
	return o.TotalScannedBytes, true
}

// HasTotalScannedBytes returns a boolean if a field has been set.
func (o *UsageSDSHour) HasTotalScannedBytes() bool {
	return o != nil && o.TotalScannedBytes != nil
}

// SetTotalScannedBytes gets a reference to the given int64 and assigns it to the TotalScannedBytes field.
func (o *UsageSDSHour) SetTotalScannedBytes(v int64) {
	o.TotalScannedBytes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSDSHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.ApmScannedBytes != nil {
		toSerialize["apm_scanned_bytes"] = o.ApmScannedBytes
	}
	if o.EventsScannedBytes != nil {
		toSerialize["events_scanned_bytes"] = o.EventsScannedBytes
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LogsScannedBytes != nil {
		toSerialize["logs_scanned_bytes"] = o.LogsScannedBytes
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.RumScannedBytes != nil {
		toSerialize["rum_scanned_bytes"] = o.RumScannedBytes
	}
	if o.TotalScannedBytes != nil {
		toSerialize["total_scanned_bytes"] = o.TotalScannedBytes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageSDSHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		ApmScannedBytes    *int64     `json:"apm_scanned_bytes,omitempty"`
		EventsScannedBytes *int64     `json:"events_scanned_bytes,omitempty"`
		Hour               *time.Time `json:"hour,omitempty"`
		LogsScannedBytes   *int64     `json:"logs_scanned_bytes,omitempty"`
		OrgName            *string    `json:"org_name,omitempty"`
		PublicId           *string    `json:"public_id,omitempty"`
		RumScannedBytes    *int64     `json:"rum_scanned_bytes,omitempty"`
		TotalScannedBytes  *int64     `json:"total_scanned_bytes,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"apm_scanned_bytes", "events_scanned_bytes", "hour", "logs_scanned_bytes", "org_name", "public_id", "rum_scanned_bytes", "total_scanned_bytes"})
	} else {
		return err
	}
	o.ApmScannedBytes = all.ApmScannedBytes
	o.EventsScannedBytes = all.EventsScannedBytes
	o.Hour = all.Hour
	o.LogsScannedBytes = all.LogsScannedBytes
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.RumScannedBytes = all.RumScannedBytes
	o.TotalScannedBytes = all.TotalScannedBytes
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
