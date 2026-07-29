// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumConfigAttributes Attributes of the RUM configuration.
type RumConfigAttributes struct {
	// Whether the RUM configuration is disabled for the organization.
	Disabled *bool `json:"disabled,omitempty"`
	// Whether application tags are enforced for the RUM applications in the organization.
	EnforcedApplicationTags bool `json:"enforced_application_tags"`
	// Timestamp of when the enforced application tags setting was last updated.
	EnforcedApplicationTagsUpdatedAt *time.Time `json:"enforced_application_tags_updated_at,omitempty"`
	// Handle of the user who last updated the enforced application tags setting.
	EnforcedApplicationTagsUpdatedBy *string `json:"enforced_application_tags_updated_by,omitempty"`
	// Version of the out-of-the-box metrics installed for the organization.
	OotbMetricsVersion *int64 `json:"ootb_metrics_version,omitempty"`
	// Timestamp of when the out-of-the-box metrics version was installed.
	OotbMetricsVersionInstalledAt *time.Time `json:"ootb_metrics_version_installed_at,omitempty"`
	// Whether retention filters are enabled for the organization.
	RetentionFiltersEnabled bool `json:"retention_filters_enabled"`
	// Timestamp of when the retention filters setting was last updated.
	RetentionFiltersEnabledUpdatedAt *time.Time `json:"retention_filters_enabled_updated_at,omitempty"`
	// Handle of the user or job who last updated the retention filters setting.
	RetentionFiltersEnabledUpdatedBy *string `json:"retention_filters_enabled_updated_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumConfigAttributes instantiates a new RumConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumConfigAttributes(enforcedApplicationTags bool, retentionFiltersEnabled bool) *RumConfigAttributes {
	this := RumConfigAttributes{}
	this.EnforcedApplicationTags = enforcedApplicationTags
	this.RetentionFiltersEnabled = retentionFiltersEnabled
	return &this
}

// NewRumConfigAttributesWithDefaults instantiates a new RumConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumConfigAttributesWithDefaults() *RumConfigAttributes {
	this := RumConfigAttributes{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *RumConfigAttributes) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *RumConfigAttributes) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *RumConfigAttributes) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetEnforcedApplicationTags returns the EnforcedApplicationTags field value.
func (o *RumConfigAttributes) GetEnforcedApplicationTags() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EnforcedApplicationTags
}

// GetEnforcedApplicationTagsOk returns a tuple with the EnforcedApplicationTags field value
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetEnforcedApplicationTagsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnforcedApplicationTags, true
}

// SetEnforcedApplicationTags sets field value.
func (o *RumConfigAttributes) SetEnforcedApplicationTags(v bool) {
	o.EnforcedApplicationTags = v
}

// GetEnforcedApplicationTagsUpdatedAt returns the EnforcedApplicationTagsUpdatedAt field value if set, zero value otherwise.
func (o *RumConfigAttributes) GetEnforcedApplicationTagsUpdatedAt() time.Time {
	if o == nil || o.EnforcedApplicationTagsUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.EnforcedApplicationTagsUpdatedAt
}

// GetEnforcedApplicationTagsUpdatedAtOk returns a tuple with the EnforcedApplicationTagsUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetEnforcedApplicationTagsUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.EnforcedApplicationTagsUpdatedAt == nil {
		return nil, false
	}
	return o.EnforcedApplicationTagsUpdatedAt, true
}

// HasEnforcedApplicationTagsUpdatedAt returns a boolean if a field has been set.
func (o *RumConfigAttributes) HasEnforcedApplicationTagsUpdatedAt() bool {
	return o != nil && o.EnforcedApplicationTagsUpdatedAt != nil
}

// SetEnforcedApplicationTagsUpdatedAt gets a reference to the given time.Time and assigns it to the EnforcedApplicationTagsUpdatedAt field.
func (o *RumConfigAttributes) SetEnforcedApplicationTagsUpdatedAt(v time.Time) {
	o.EnforcedApplicationTagsUpdatedAt = &v
}

// GetEnforcedApplicationTagsUpdatedBy returns the EnforcedApplicationTagsUpdatedBy field value if set, zero value otherwise.
func (o *RumConfigAttributes) GetEnforcedApplicationTagsUpdatedBy() string {
	if o == nil || o.EnforcedApplicationTagsUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.EnforcedApplicationTagsUpdatedBy
}

// GetEnforcedApplicationTagsUpdatedByOk returns a tuple with the EnforcedApplicationTagsUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetEnforcedApplicationTagsUpdatedByOk() (*string, bool) {
	if o == nil || o.EnforcedApplicationTagsUpdatedBy == nil {
		return nil, false
	}
	return o.EnforcedApplicationTagsUpdatedBy, true
}

// HasEnforcedApplicationTagsUpdatedBy returns a boolean if a field has been set.
func (o *RumConfigAttributes) HasEnforcedApplicationTagsUpdatedBy() bool {
	return o != nil && o.EnforcedApplicationTagsUpdatedBy != nil
}

// SetEnforcedApplicationTagsUpdatedBy gets a reference to the given string and assigns it to the EnforcedApplicationTagsUpdatedBy field.
func (o *RumConfigAttributes) SetEnforcedApplicationTagsUpdatedBy(v string) {
	o.EnforcedApplicationTagsUpdatedBy = &v
}

// GetOotbMetricsVersion returns the OotbMetricsVersion field value if set, zero value otherwise.
func (o *RumConfigAttributes) GetOotbMetricsVersion() int64 {
	if o == nil || o.OotbMetricsVersion == nil {
		var ret int64
		return ret
	}
	return *o.OotbMetricsVersion
}

// GetOotbMetricsVersionOk returns a tuple with the OotbMetricsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetOotbMetricsVersionOk() (*int64, bool) {
	if o == nil || o.OotbMetricsVersion == nil {
		return nil, false
	}
	return o.OotbMetricsVersion, true
}

// HasOotbMetricsVersion returns a boolean if a field has been set.
func (o *RumConfigAttributes) HasOotbMetricsVersion() bool {
	return o != nil && o.OotbMetricsVersion != nil
}

// SetOotbMetricsVersion gets a reference to the given int64 and assigns it to the OotbMetricsVersion field.
func (o *RumConfigAttributes) SetOotbMetricsVersion(v int64) {
	o.OotbMetricsVersion = &v
}

// GetOotbMetricsVersionInstalledAt returns the OotbMetricsVersionInstalledAt field value if set, zero value otherwise.
func (o *RumConfigAttributes) GetOotbMetricsVersionInstalledAt() time.Time {
	if o == nil || o.OotbMetricsVersionInstalledAt == nil {
		var ret time.Time
		return ret
	}
	return *o.OotbMetricsVersionInstalledAt
}

// GetOotbMetricsVersionInstalledAtOk returns a tuple with the OotbMetricsVersionInstalledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetOotbMetricsVersionInstalledAtOk() (*time.Time, bool) {
	if o == nil || o.OotbMetricsVersionInstalledAt == nil {
		return nil, false
	}
	return o.OotbMetricsVersionInstalledAt, true
}

// HasOotbMetricsVersionInstalledAt returns a boolean if a field has been set.
func (o *RumConfigAttributes) HasOotbMetricsVersionInstalledAt() bool {
	return o != nil && o.OotbMetricsVersionInstalledAt != nil
}

// SetOotbMetricsVersionInstalledAt gets a reference to the given time.Time and assigns it to the OotbMetricsVersionInstalledAt field.
func (o *RumConfigAttributes) SetOotbMetricsVersionInstalledAt(v time.Time) {
	o.OotbMetricsVersionInstalledAt = &v
}

// GetRetentionFiltersEnabled returns the RetentionFiltersEnabled field value.
func (o *RumConfigAttributes) GetRetentionFiltersEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RetentionFiltersEnabled
}

// GetRetentionFiltersEnabledOk returns a tuple with the RetentionFiltersEnabled field value
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetRetentionFiltersEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RetentionFiltersEnabled, true
}

// SetRetentionFiltersEnabled sets field value.
func (o *RumConfigAttributes) SetRetentionFiltersEnabled(v bool) {
	o.RetentionFiltersEnabled = v
}

// GetRetentionFiltersEnabledUpdatedAt returns the RetentionFiltersEnabledUpdatedAt field value if set, zero value otherwise.
func (o *RumConfigAttributes) GetRetentionFiltersEnabledUpdatedAt() time.Time {
	if o == nil || o.RetentionFiltersEnabledUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RetentionFiltersEnabledUpdatedAt
}

// GetRetentionFiltersEnabledUpdatedAtOk returns a tuple with the RetentionFiltersEnabledUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetRetentionFiltersEnabledUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.RetentionFiltersEnabledUpdatedAt == nil {
		return nil, false
	}
	return o.RetentionFiltersEnabledUpdatedAt, true
}

// HasRetentionFiltersEnabledUpdatedAt returns a boolean if a field has been set.
func (o *RumConfigAttributes) HasRetentionFiltersEnabledUpdatedAt() bool {
	return o != nil && o.RetentionFiltersEnabledUpdatedAt != nil
}

// SetRetentionFiltersEnabledUpdatedAt gets a reference to the given time.Time and assigns it to the RetentionFiltersEnabledUpdatedAt field.
func (o *RumConfigAttributes) SetRetentionFiltersEnabledUpdatedAt(v time.Time) {
	o.RetentionFiltersEnabledUpdatedAt = &v
}

// GetRetentionFiltersEnabledUpdatedBy returns the RetentionFiltersEnabledUpdatedBy field value if set, zero value otherwise.
func (o *RumConfigAttributes) GetRetentionFiltersEnabledUpdatedBy() string {
	if o == nil || o.RetentionFiltersEnabledUpdatedBy == nil {
		var ret string
		return ret
	}
	return *o.RetentionFiltersEnabledUpdatedBy
}

// GetRetentionFiltersEnabledUpdatedByOk returns a tuple with the RetentionFiltersEnabledUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumConfigAttributes) GetRetentionFiltersEnabledUpdatedByOk() (*string, bool) {
	if o == nil || o.RetentionFiltersEnabledUpdatedBy == nil {
		return nil, false
	}
	return o.RetentionFiltersEnabledUpdatedBy, true
}

// HasRetentionFiltersEnabledUpdatedBy returns a boolean if a field has been set.
func (o *RumConfigAttributes) HasRetentionFiltersEnabledUpdatedBy() bool {
	return o != nil && o.RetentionFiltersEnabledUpdatedBy != nil
}

// SetRetentionFiltersEnabledUpdatedBy gets a reference to the given string and assigns it to the RetentionFiltersEnabledUpdatedBy field.
func (o *RumConfigAttributes) SetRetentionFiltersEnabledUpdatedBy(v string) {
	o.RetentionFiltersEnabledUpdatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	toSerialize["enforced_application_tags"] = o.EnforcedApplicationTags
	if o.EnforcedApplicationTagsUpdatedAt != nil {
		if o.EnforcedApplicationTagsUpdatedAt.Nanosecond() == 0 {
			toSerialize["enforced_application_tags_updated_at"] = o.EnforcedApplicationTagsUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["enforced_application_tags_updated_at"] = o.EnforcedApplicationTagsUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.EnforcedApplicationTagsUpdatedBy != nil {
		toSerialize["enforced_application_tags_updated_by"] = o.EnforcedApplicationTagsUpdatedBy
	}
	if o.OotbMetricsVersion != nil {
		toSerialize["ootb_metrics_version"] = o.OotbMetricsVersion
	}
	if o.OotbMetricsVersionInstalledAt != nil {
		if o.OotbMetricsVersionInstalledAt.Nanosecond() == 0 {
			toSerialize["ootb_metrics_version_installed_at"] = o.OotbMetricsVersionInstalledAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["ootb_metrics_version_installed_at"] = o.OotbMetricsVersionInstalledAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["retention_filters_enabled"] = o.RetentionFiltersEnabled
	if o.RetentionFiltersEnabledUpdatedAt != nil {
		if o.RetentionFiltersEnabledUpdatedAt.Nanosecond() == 0 {
			toSerialize["retention_filters_enabled_updated_at"] = o.RetentionFiltersEnabledUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["retention_filters_enabled_updated_at"] = o.RetentionFiltersEnabledUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.RetentionFiltersEnabledUpdatedBy != nil {
		toSerialize["retention_filters_enabled_updated_by"] = o.RetentionFiltersEnabledUpdatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Disabled                         *bool      `json:"disabled,omitempty"`
		EnforcedApplicationTags          *bool      `json:"enforced_application_tags"`
		EnforcedApplicationTagsUpdatedAt *time.Time `json:"enforced_application_tags_updated_at,omitempty"`
		EnforcedApplicationTagsUpdatedBy *string    `json:"enforced_application_tags_updated_by,omitempty"`
		OotbMetricsVersion               *int64     `json:"ootb_metrics_version,omitempty"`
		OotbMetricsVersionInstalledAt    *time.Time `json:"ootb_metrics_version_installed_at,omitempty"`
		RetentionFiltersEnabled          *bool      `json:"retention_filters_enabled"`
		RetentionFiltersEnabledUpdatedAt *time.Time `json:"retention_filters_enabled_updated_at,omitempty"`
		RetentionFiltersEnabledUpdatedBy *string    `json:"retention_filters_enabled_updated_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EnforcedApplicationTags == nil {
		return fmt.Errorf("required field enforced_application_tags missing")
	}
	if all.RetentionFiltersEnabled == nil {
		return fmt.Errorf("required field retention_filters_enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"disabled", "enforced_application_tags", "enforced_application_tags_updated_at", "enforced_application_tags_updated_by", "ootb_metrics_version", "ootb_metrics_version_installed_at", "retention_filters_enabled", "retention_filters_enabled_updated_at", "retention_filters_enabled_updated_by"})
	} else {
		return err
	}
	o.Disabled = all.Disabled
	o.EnforcedApplicationTags = *all.EnforcedApplicationTags
	o.EnforcedApplicationTagsUpdatedAt = all.EnforcedApplicationTagsUpdatedAt
	o.EnforcedApplicationTagsUpdatedBy = all.EnforcedApplicationTagsUpdatedBy
	o.OotbMetricsVersion = all.OotbMetricsVersion
	o.OotbMetricsVersionInstalledAt = all.OotbMetricsVersionInstalledAt
	o.RetentionFiltersEnabled = *all.RetentionFiltersEnabled
	o.RetentionFiltersEnabledUpdatedAt = all.RetentionFiltersEnabledUpdatedAt
	o.RetentionFiltersEnabledUpdatedBy = all.RetentionFiltersEnabledUpdatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
