// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumArchiveConfigAttributes Archive configuration.
type RumArchiveConfigAttributes struct {
	// Whether archiving is enabled.
	Enabled bool `json:"enabled"`
	// Organization ID.
	OrgId int64 `json:"org_id"`
	// Configured retention in days; omitted when unset. When archiving is enabled, this is 30, 90, or 180.
	RetentionDays *int32 `json:"retention_days,omitempty"`
	// Last update time.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Handle of the user who last updated the configuration.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumArchiveConfigAttributes instantiates a new RumArchiveConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumArchiveConfigAttributes(enabled bool, orgId int64) *RumArchiveConfigAttributes {
	this := RumArchiveConfigAttributes{}
	this.Enabled = enabled
	this.OrgId = orgId
	return &this
}

// NewRumArchiveConfigAttributesWithDefaults instantiates a new RumArchiveConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumArchiveConfigAttributesWithDefaults() *RumArchiveConfigAttributes {
	this := RumArchiveConfigAttributes{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *RumArchiveConfigAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RumArchiveConfigAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RumArchiveConfigAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOrgId returns the OrgId field value.
func (o *RumArchiveConfigAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RumArchiveConfigAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *RumArchiveConfigAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetRetentionDays returns the RetentionDays field value if set, zero value otherwise.
func (o *RumArchiveConfigAttributes) GetRetentionDays() int32 {
	if o == nil || o.RetentionDays == nil {
		var ret int32
		return ret
	}
	return *o.RetentionDays
}

// GetRetentionDaysOk returns a tuple with the RetentionDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumArchiveConfigAttributes) GetRetentionDaysOk() (*int32, bool) {
	if o == nil || o.RetentionDays == nil {
		return nil, false
	}
	return o.RetentionDays, true
}

// HasRetentionDays returns a boolean if a field has been set.
func (o *RumArchiveConfigAttributes) HasRetentionDays() bool {
	return o != nil && o.RetentionDays != nil
}

// SetRetentionDays gets a reference to the given int32 and assigns it to the RetentionDays field.
func (o *RumArchiveConfigAttributes) SetRetentionDays(v int32) {
	o.RetentionDays = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RumArchiveConfigAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumArchiveConfigAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RumArchiveConfigAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RumArchiveConfigAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RumArchiveConfigAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumArchiveConfigAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RumArchiveConfigAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *RumArchiveConfigAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumArchiveConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["org_id"] = o.OrgId
	if o.RetentionDays != nil {
		toSerialize["retention_days"] = o.RetentionDays
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumArchiveConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled       *bool      `json:"enabled"`
		OrgId         *int64     `json:"org_id"`
		RetentionDays *int32     `json:"retention_days,omitempty"`
		UpdatedAt     *time.Time `json:"updated_at,omitempty"`
		UpdatedBy     *string    `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "org_id", "retention_days", "updated_at", "updated_by"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.OrgId = *all.OrgId
	o.RetentionDays = all.RetentionDays
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
