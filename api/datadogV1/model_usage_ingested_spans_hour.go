// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageIngestedSpansHour Ingested spans usage for a given organization for a given hour.
type UsageIngestedSpansHour struct {
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the total number of bytes ingested for APM spans during a given hour.
	IngestedEventsBytes datadog.NullableInt64 `json:"ingested_events_bytes,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewUsageIngestedSpansHour instantiates a new UsageIngestedSpansHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageIngestedSpansHour() *UsageIngestedSpansHour {
	this := UsageIngestedSpansHour{}
	return &this
}

// NewUsageIngestedSpansHourWithDefaults instantiates a new UsageIngestedSpansHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageIngestedSpansHourWithDefaults() *UsageIngestedSpansHour {
	this := UsageIngestedSpansHour{}
	return &this
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageIngestedSpansHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIngestedSpansHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageIngestedSpansHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageIngestedSpansHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetIngestedEventsBytes returns the IngestedEventsBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageIngestedSpansHour) GetIngestedEventsBytes() int64 {
	if o == nil || o.IngestedEventsBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytes.Get()
}

// GetIngestedEventsBytesOk returns a tuple with the IngestedEventsBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageIngestedSpansHour) GetIngestedEventsBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngestedEventsBytes.Get(), o.IngestedEventsBytes.IsSet()
}

// HasIngestedEventsBytes returns a boolean if a field has been set.
func (o *UsageIngestedSpansHour) HasIngestedEventsBytes() bool {
	return o != nil && o.IngestedEventsBytes.IsSet()
}

// SetIngestedEventsBytes gets a reference to the given datadog.NullableInt64 and assigns it to the IngestedEventsBytes field.
func (o *UsageIngestedSpansHour) SetIngestedEventsBytes(v int64) {
	o.IngestedEventsBytes.Set(&v)
}

// SetIngestedEventsBytesNil sets the value for IngestedEventsBytes to be an explicit nil.
func (o *UsageIngestedSpansHour) SetIngestedEventsBytesNil() {
	o.IngestedEventsBytes.Set(nil)
}

// UnsetIngestedEventsBytes ensures that no value is present for IngestedEventsBytes, not even an explicit nil.
func (o *UsageIngestedSpansHour) UnsetIngestedEventsBytes() {
	o.IngestedEventsBytes.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageIngestedSpansHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIngestedSpansHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageIngestedSpansHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageIngestedSpansHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageIngestedSpansHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageIngestedSpansHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageIngestedSpansHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageIngestedSpansHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageIngestedSpansHour) MarshalJSON() ([]byte, error) {
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
	if o.IngestedEventsBytes.IsSet() {
		toSerialize["ingested_events_bytes"] = o.IngestedEventsBytes.Get()
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
func (o *UsageIngestedSpansHour) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Hour                *time.Time            `json:"hour,omitempty"`
		IngestedEventsBytes datadog.NullableInt64 `json:"ingested_events_bytes,omitempty"`
		OrgName             *string               `json:"org_name,omitempty"`
		PublicId            *string               `json:"public_id,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hour", "ingested_events_bytes", "org_name", "public_id"})
	} else {
		return err
	}
	o.Hour = all.Hour
	o.IngestedEventsBytes = all.IngestedEventsBytes
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
