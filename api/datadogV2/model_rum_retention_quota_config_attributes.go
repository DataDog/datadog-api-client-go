// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMRetentionQuotaConfigAttributes The RUM retention quota configuration properties.
type RUMRetentionQuotaConfigAttributes struct {
	// The configuration used when `mode` is `custom`.
	Custom *RUMRetentionQuotaCustomConfig `json:"custom,omitempty"`
	// The retention quota mode. `custom` enforces a fixed session limit.
	// `custom` is the only supported mode.
	Mode RUMRetentionQuotaMode `json:"mode"`
	// The ID of the organization the retention quota configuration belongs to.
	OrgId int64 `json:"org_id"`
	// The date the retention quota configuration was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The handle of the user who last updated the retention quota configuration.
	UpdatedBy *string `json:"updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMRetentionQuotaConfigAttributes instantiates a new RUMRetentionQuotaConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMRetentionQuotaConfigAttributes(mode RUMRetentionQuotaMode, orgId int64) *RUMRetentionQuotaConfigAttributes {
	this := RUMRetentionQuotaConfigAttributes{}
	this.Mode = mode
	this.OrgId = orgId
	return &this
}

// NewRUMRetentionQuotaConfigAttributesWithDefaults instantiates a new RUMRetentionQuotaConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMRetentionQuotaConfigAttributesWithDefaults() *RUMRetentionQuotaConfigAttributes {
	this := RUMRetentionQuotaConfigAttributes{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *RUMRetentionQuotaConfigAttributes) GetCustom() RUMRetentionQuotaCustomConfig {
	if o == nil || o.Custom == nil {
		var ret RUMRetentionQuotaCustomConfig
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionQuotaConfigAttributes) GetCustomOk() (*RUMRetentionQuotaCustomConfig, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *RUMRetentionQuotaConfigAttributes) HasCustom() bool {
	return o != nil && o.Custom != nil
}

// SetCustom gets a reference to the given RUMRetentionQuotaCustomConfig and assigns it to the Custom field.
func (o *RUMRetentionQuotaConfigAttributes) SetCustom(v RUMRetentionQuotaCustomConfig) {
	o.Custom = &v
}

// GetMode returns the Mode field value.
func (o *RUMRetentionQuotaConfigAttributes) GetMode() RUMRetentionQuotaMode {
	if o == nil {
		var ret RUMRetentionQuotaMode
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *RUMRetentionQuotaConfigAttributes) GetModeOk() (*RUMRetentionQuotaMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *RUMRetentionQuotaConfigAttributes) SetMode(v RUMRetentionQuotaMode) {
	o.Mode = v
}

// GetOrgId returns the OrgId field value.
func (o *RUMRetentionQuotaConfigAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RUMRetentionQuotaConfigAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *RUMRetentionQuotaConfigAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RUMRetentionQuotaConfigAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionQuotaConfigAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RUMRetentionQuotaConfigAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RUMRetentionQuotaConfigAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RUMRetentionQuotaConfigAttributes) GetUpdatedBy() string {
	if o == nil || o.UpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMRetentionQuotaConfigAttributes) GetUpdatedByOk() (*string, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RUMRetentionQuotaConfigAttributes) HasUpdatedBy() bool {
	return o != nil && o.UpdatedBy != nil
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *RUMRetentionQuotaConfigAttributes) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMRetentionQuotaConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	toSerialize["mode"] = o.Mode
	toSerialize["org_id"] = o.OrgId
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
func (o *RUMRetentionQuotaConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Custom    *RUMRetentionQuotaCustomConfig `json:"custom,omitempty"`
		Mode      *RUMRetentionQuotaMode         `json:"mode"`
		OrgId     *int64                         `json:"org_id"`
		UpdatedAt *time.Time                     `json:"updated_at,omitempty"`
		UpdatedBy *string                        `json:"updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom", "mode", "org_id", "updated_at", "updated_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Custom != nil && all.Custom.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Custom = all.Custom
	if !all.Mode.IsValid() {
		hasInvalidField = true
	} else {
		o.Mode = *all.Mode
	}
	o.OrgId = *all.OrgId
	o.UpdatedAt = all.UpdatedAt
	o.UpdatedBy = all.UpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
