// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"encoding/json"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageAnalyzedLogsHour The number of analyzed logs for each hour for a given organization.
type UsageAnalyzedLogsHour struct {
	// Contains the number of analyzed logs.
	AnalyzedLogs datadog.NullableInt64 `json:"analyzed_logs,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageAnalyzedLogsHour instantiates a new UsageAnalyzedLogsHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageAnalyzedLogsHour() *UsageAnalyzedLogsHour {
	this := UsageAnalyzedLogsHour{}
	return &this
}

// NewUsageAnalyzedLogsHourWithDefaults instantiates a new UsageAnalyzedLogsHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageAnalyzedLogsHourWithDefaults() *UsageAnalyzedLogsHour {
	this := UsageAnalyzedLogsHour{}
	return &this
}

// GetAnalyzedLogs returns the AnalyzedLogs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageAnalyzedLogsHour) GetAnalyzedLogs() int64 {
	if o == nil || o.AnalyzedLogs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AnalyzedLogs.Get()
}

// GetAnalyzedLogsOk returns a tuple with the AnalyzedLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageAnalyzedLogsHour) GetAnalyzedLogsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnalyzedLogs.Get(), o.AnalyzedLogs.IsSet()
}

// HasAnalyzedLogs returns a boolean if a field has been set.
func (o *UsageAnalyzedLogsHour) HasAnalyzedLogs() bool {
	return o != nil && o.AnalyzedLogs.IsSet()
}

// SetAnalyzedLogs gets a reference to the given datadog.NullableInt64 and assigns it to the AnalyzedLogs field.
func (o *UsageAnalyzedLogsHour) SetAnalyzedLogs(v int64) {
	o.AnalyzedLogs.Set(&v)
}

// SetAnalyzedLogsNil sets the value for AnalyzedLogs to be an explicit nil.
func (o *UsageAnalyzedLogsHour) SetAnalyzedLogsNil() {
	o.AnalyzedLogs.Set(nil)
}

// UnsetAnalyzedLogs ensures that no value is present for AnalyzedLogs, not even an explicit nil.
func (o *UsageAnalyzedLogsHour) UnsetAnalyzedLogs() {
	o.AnalyzedLogs.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageAnalyzedLogsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAnalyzedLogsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageAnalyzedLogsHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageAnalyzedLogsHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageAnalyzedLogsHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAnalyzedLogsHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageAnalyzedLogsHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageAnalyzedLogsHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageAnalyzedLogsHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAnalyzedLogsHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageAnalyzedLogsHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageAnalyzedLogsHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageAnalyzedLogsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.AnalyzedLogs.IsSet() {
		toSerialize["analyzed_logs"] = o.AnalyzedLogs.Get()
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
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

// UnmarshalJSON deserializes the given payload.
func (o *UsageAnalyzedLogsHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		AnalyzedLogs datadog.NullableInt64 `json:"analyzed_logs,omitempty"`
		Hour         *time.Time            `json:"hour,omitempty"`
		OrgName      *string               `json:"org_name,omitempty"`
		PublicId     *string               `json:"public_id,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"analyzed_logs", "hour", "org_name", "public_id"})
	} else {
		return err
	}
	o.AnalyzedLogs = all.AnalyzedLogs
	o.Hour = all.Hour
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
