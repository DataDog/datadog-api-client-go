// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RUMApplicationListAttributes RUM application list attributes.
type RUMApplicationListAttributes struct {
	// ID of the API key associated with the application.
	ApiKeyId *int32 `json:"api_key_id,omitempty"`
	// Indicates if flat sampling of Session Replay for APM-correlated RUM sessions is enabled.
	ApmRumFlatSamplingReplayEnabled *bool `json:"apm_rum_flat_sampling_replay_enabled,omitempty"`
	// Sample rate (0-100) of Session Replay for APM-correlated RUM sessions when flat sampling is enabled.
	ApmRumFlatSamplingReplaySampleRate *float64 `json:"apm_rum_flat_sampling_replay_sample_rate,omitempty"`
	// ID of the RUM application.
	ApplicationId string `json:"application_id"`
	// Timestamp in ms of the creation date.
	CreatedAt int64 `json:"created_at"`
	// Handle of the creator user.
	CreatedByHandle string `json:"created_by_handle"`
	// Timestamp in ms of the date the RUM application was disabled. Only present if the application was disabled.
	DisabledAt *int64 `json:"disabled_at,omitempty"`
	// Handle of the user who disabled the RUM application. Only present if the application was disabled.
	DisabledByHandle *string `json:"disabled_by_handle,omitempty"`
	// Indicates if Error Tracking exclusion filters are enabled for the application.
	ErrorTrackingExclusionFilterEnabled *bool `json:"error_tracking_exclusion_filter_enabled,omitempty"`
	// Non-native ingestion paths the RUM application sends its data through. `otel` means the application sends RUM data through OpenTelemetry instead of the Datadog RUM SDK.
	ExternalFeeds []string `json:"external_feeds,omitempty"`
	// Hash of the RUM application. Optional.
	Hash *string `json:"hash,omitempty"`
	// Indicates if the RUM application is active.
	IsActive *bool `json:"is_active,omitempty"`
	// Name of the RUM application.
	Name string `json:"name"`
	// Indicates if the out-of-the-box RUM metrics are installed for the application.
	OotbMetricsInstalled *bool `json:"ootb_metrics_installed,omitempty"`
	// Org ID of the RUM application.
	OrgId int32 `json:"org_id"`
	// Indicates that the Product Analytics preview is disabled for the application. Only present (and `true`) when the Product Analytics retention state is not preview.
	ProductAnalyticsPreviewDisabled *bool `json:"product_analytics_preview_disabled,omitempty"`
	// Sample rate (0-100) of Session Replay for Product Analytics sessions.
	ProductAnalyticsReplaySampleRate *int32 `json:"product_analytics_replay_sample_rate,omitempty"`
	// Product Scales configuration for the RUM application.
	ProductScales *RUMProductScales `json:"product_scales,omitempty"`
	// ID of the RUM SDK remote configuration for the application, if one exists.
	RemoteConfigId *string `json:"remote_config_id,omitempty"`
	// Short name of the RUM application.
	ShortName *string `json:"short_name,omitempty"`
	// List of tags associated with the RUM application.
	Tags []string `json:"tags,omitempty"`
	// Type of the RUM application. Possible values are `browser`, `ios`, `android`, `react-native`, `flutter`, `roku`, `electron`, `unity`, `kotlin-multiplatform`, `react`, `nextjs`, `angular`, `vuejs`, `nuxt`, `tanstack`, `maui`, `salesforce-lwc`, `shopify`, `cpp`, `rum-cpp`, `unreal`.
	Type string `json:"type"`
	// Timestamp in ms of the last update date.
	UpdatedAt int64 `json:"updated_at"`
	// Handle of the updater user.
	UpdatedByHandle string `json:"updated_by_handle"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRUMApplicationListAttributes instantiates a new RUMApplicationListAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRUMApplicationListAttributes(applicationId string, createdAt int64, createdByHandle string, name string, orgId int32, typeVar string, updatedAt int64, updatedByHandle string) *RUMApplicationListAttributes {
	this := RUMApplicationListAttributes{}
	this.ApplicationId = applicationId
	this.CreatedAt = createdAt
	this.CreatedByHandle = createdByHandle
	this.Name = name
	this.OrgId = orgId
	this.Type = typeVar
	this.UpdatedAt = updatedAt
	this.UpdatedByHandle = updatedByHandle
	return &this
}

// NewRUMApplicationListAttributesWithDefaults instantiates a new RUMApplicationListAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRUMApplicationListAttributesWithDefaults() *RUMApplicationListAttributes {
	this := RUMApplicationListAttributes{}
	return &this
}

// GetApiKeyId returns the ApiKeyId field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetApiKeyId() int32 {
	if o == nil || o.ApiKeyId == nil {
		var ret int32
		return ret
	}
	return *o.ApiKeyId
}

// GetApiKeyIdOk returns a tuple with the ApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetApiKeyIdOk() (*int32, bool) {
	if o == nil || o.ApiKeyId == nil {
		return nil, false
	}
	return o.ApiKeyId, true
}

// HasApiKeyId returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasApiKeyId() bool {
	return o != nil && o.ApiKeyId != nil
}

// SetApiKeyId gets a reference to the given int32 and assigns it to the ApiKeyId field.
func (o *RUMApplicationListAttributes) SetApiKeyId(v int32) {
	o.ApiKeyId = &v
}

// GetApmRumFlatSamplingReplayEnabled returns the ApmRumFlatSamplingReplayEnabled field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetApmRumFlatSamplingReplayEnabled() bool {
	if o == nil || o.ApmRumFlatSamplingReplayEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ApmRumFlatSamplingReplayEnabled
}

// GetApmRumFlatSamplingReplayEnabledOk returns a tuple with the ApmRumFlatSamplingReplayEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetApmRumFlatSamplingReplayEnabledOk() (*bool, bool) {
	if o == nil || o.ApmRumFlatSamplingReplayEnabled == nil {
		return nil, false
	}
	return o.ApmRumFlatSamplingReplayEnabled, true
}

// HasApmRumFlatSamplingReplayEnabled returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasApmRumFlatSamplingReplayEnabled() bool {
	return o != nil && o.ApmRumFlatSamplingReplayEnabled != nil
}

// SetApmRumFlatSamplingReplayEnabled gets a reference to the given bool and assigns it to the ApmRumFlatSamplingReplayEnabled field.
func (o *RUMApplicationListAttributes) SetApmRumFlatSamplingReplayEnabled(v bool) {
	o.ApmRumFlatSamplingReplayEnabled = &v
}

// GetApmRumFlatSamplingReplaySampleRate returns the ApmRumFlatSamplingReplaySampleRate field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetApmRumFlatSamplingReplaySampleRate() float64 {
	if o == nil || o.ApmRumFlatSamplingReplaySampleRate == nil {
		var ret float64
		return ret
	}
	return *o.ApmRumFlatSamplingReplaySampleRate
}

// GetApmRumFlatSamplingReplaySampleRateOk returns a tuple with the ApmRumFlatSamplingReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetApmRumFlatSamplingReplaySampleRateOk() (*float64, bool) {
	if o == nil || o.ApmRumFlatSamplingReplaySampleRate == nil {
		return nil, false
	}
	return o.ApmRumFlatSamplingReplaySampleRate, true
}

// HasApmRumFlatSamplingReplaySampleRate returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasApmRumFlatSamplingReplaySampleRate() bool {
	return o != nil && o.ApmRumFlatSamplingReplaySampleRate != nil
}

// SetApmRumFlatSamplingReplaySampleRate gets a reference to the given float64 and assigns it to the ApmRumFlatSamplingReplaySampleRate field.
func (o *RUMApplicationListAttributes) SetApmRumFlatSamplingReplaySampleRate(v float64) {
	o.ApmRumFlatSamplingReplaySampleRate = &v
}

// GetApplicationId returns the ApplicationId field value.
func (o *RUMApplicationListAttributes) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *RUMApplicationListAttributes) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *RUMApplicationListAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *RUMApplicationListAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetCreatedByHandle returns the CreatedByHandle field value.
func (o *RUMApplicationListAttributes) GetCreatedByHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetCreatedByHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByHandle, true
}

// SetCreatedByHandle sets field value.
func (o *RUMApplicationListAttributes) SetCreatedByHandle(v string) {
	o.CreatedByHandle = v
}

// GetDisabledAt returns the DisabledAt field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetDisabledAt() int64 {
	if o == nil || o.DisabledAt == nil {
		var ret int64
		return ret
	}
	return *o.DisabledAt
}

// GetDisabledAtOk returns a tuple with the DisabledAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetDisabledAtOk() (*int64, bool) {
	if o == nil || o.DisabledAt == nil {
		return nil, false
	}
	return o.DisabledAt, true
}

// HasDisabledAt returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasDisabledAt() bool {
	return o != nil && o.DisabledAt != nil
}

// SetDisabledAt gets a reference to the given int64 and assigns it to the DisabledAt field.
func (o *RUMApplicationListAttributes) SetDisabledAt(v int64) {
	o.DisabledAt = &v
}

// GetDisabledByHandle returns the DisabledByHandle field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetDisabledByHandle() string {
	if o == nil || o.DisabledByHandle == nil {
		var ret string
		return ret
	}
	return *o.DisabledByHandle
}

// GetDisabledByHandleOk returns a tuple with the DisabledByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetDisabledByHandleOk() (*string, bool) {
	if o == nil || o.DisabledByHandle == nil {
		return nil, false
	}
	return o.DisabledByHandle, true
}

// HasDisabledByHandle returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasDisabledByHandle() bool {
	return o != nil && o.DisabledByHandle != nil
}

// SetDisabledByHandle gets a reference to the given string and assigns it to the DisabledByHandle field.
func (o *RUMApplicationListAttributes) SetDisabledByHandle(v string) {
	o.DisabledByHandle = &v
}

// GetErrorTrackingExclusionFilterEnabled returns the ErrorTrackingExclusionFilterEnabled field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetErrorTrackingExclusionFilterEnabled() bool {
	if o == nil || o.ErrorTrackingExclusionFilterEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ErrorTrackingExclusionFilterEnabled
}

// GetErrorTrackingExclusionFilterEnabledOk returns a tuple with the ErrorTrackingExclusionFilterEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetErrorTrackingExclusionFilterEnabledOk() (*bool, bool) {
	if o == nil || o.ErrorTrackingExclusionFilterEnabled == nil {
		return nil, false
	}
	return o.ErrorTrackingExclusionFilterEnabled, true
}

// HasErrorTrackingExclusionFilterEnabled returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasErrorTrackingExclusionFilterEnabled() bool {
	return o != nil && o.ErrorTrackingExclusionFilterEnabled != nil
}

// SetErrorTrackingExclusionFilterEnabled gets a reference to the given bool and assigns it to the ErrorTrackingExclusionFilterEnabled field.
func (o *RUMApplicationListAttributes) SetErrorTrackingExclusionFilterEnabled(v bool) {
	o.ErrorTrackingExclusionFilterEnabled = &v
}

// GetExternalFeeds returns the ExternalFeeds field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetExternalFeeds() []string {
	if o == nil || o.ExternalFeeds == nil {
		var ret []string
		return ret
	}
	return o.ExternalFeeds
}

// GetExternalFeedsOk returns a tuple with the ExternalFeeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetExternalFeedsOk() (*[]string, bool) {
	if o == nil || o.ExternalFeeds == nil {
		return nil, false
	}
	return &o.ExternalFeeds, true
}

// HasExternalFeeds returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasExternalFeeds() bool {
	return o != nil && o.ExternalFeeds != nil
}

// SetExternalFeeds gets a reference to the given []string and assigns it to the ExternalFeeds field.
func (o *RUMApplicationListAttributes) SetExternalFeeds(v []string) {
	o.ExternalFeeds = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasHash() bool {
	return o != nil && o.Hash != nil
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *RUMApplicationListAttributes) SetHash(v string) {
	o.Hash = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasIsActive() bool {
	return o != nil && o.IsActive != nil
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *RUMApplicationListAttributes) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetName returns the Name field value.
func (o *RUMApplicationListAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RUMApplicationListAttributes) SetName(v string) {
	o.Name = v
}

// GetOotbMetricsInstalled returns the OotbMetricsInstalled field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetOotbMetricsInstalled() bool {
	if o == nil || o.OotbMetricsInstalled == nil {
		var ret bool
		return ret
	}
	return *o.OotbMetricsInstalled
}

// GetOotbMetricsInstalledOk returns a tuple with the OotbMetricsInstalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetOotbMetricsInstalledOk() (*bool, bool) {
	if o == nil || o.OotbMetricsInstalled == nil {
		return nil, false
	}
	return o.OotbMetricsInstalled, true
}

// HasOotbMetricsInstalled returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasOotbMetricsInstalled() bool {
	return o != nil && o.OotbMetricsInstalled != nil
}

// SetOotbMetricsInstalled gets a reference to the given bool and assigns it to the OotbMetricsInstalled field.
func (o *RUMApplicationListAttributes) SetOotbMetricsInstalled(v bool) {
	o.OotbMetricsInstalled = &v
}

// GetOrgId returns the OrgId field value.
func (o *RUMApplicationListAttributes) GetOrgId() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetOrgIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *RUMApplicationListAttributes) SetOrgId(v int32) {
	o.OrgId = v
}

// GetProductAnalyticsPreviewDisabled returns the ProductAnalyticsPreviewDisabled field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetProductAnalyticsPreviewDisabled() bool {
	if o == nil || o.ProductAnalyticsPreviewDisabled == nil {
		var ret bool
		return ret
	}
	return *o.ProductAnalyticsPreviewDisabled
}

// GetProductAnalyticsPreviewDisabledOk returns a tuple with the ProductAnalyticsPreviewDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetProductAnalyticsPreviewDisabledOk() (*bool, bool) {
	if o == nil || o.ProductAnalyticsPreviewDisabled == nil {
		return nil, false
	}
	return o.ProductAnalyticsPreviewDisabled, true
}

// HasProductAnalyticsPreviewDisabled returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasProductAnalyticsPreviewDisabled() bool {
	return o != nil && o.ProductAnalyticsPreviewDisabled != nil
}

// SetProductAnalyticsPreviewDisabled gets a reference to the given bool and assigns it to the ProductAnalyticsPreviewDisabled field.
func (o *RUMApplicationListAttributes) SetProductAnalyticsPreviewDisabled(v bool) {
	o.ProductAnalyticsPreviewDisabled = &v
}

// GetProductAnalyticsReplaySampleRate returns the ProductAnalyticsReplaySampleRate field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetProductAnalyticsReplaySampleRate() int32 {
	if o == nil || o.ProductAnalyticsReplaySampleRate == nil {
		var ret int32
		return ret
	}
	return *o.ProductAnalyticsReplaySampleRate
}

// GetProductAnalyticsReplaySampleRateOk returns a tuple with the ProductAnalyticsReplaySampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetProductAnalyticsReplaySampleRateOk() (*int32, bool) {
	if o == nil || o.ProductAnalyticsReplaySampleRate == nil {
		return nil, false
	}
	return o.ProductAnalyticsReplaySampleRate, true
}

// HasProductAnalyticsReplaySampleRate returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasProductAnalyticsReplaySampleRate() bool {
	return o != nil && o.ProductAnalyticsReplaySampleRate != nil
}

// SetProductAnalyticsReplaySampleRate gets a reference to the given int32 and assigns it to the ProductAnalyticsReplaySampleRate field.
func (o *RUMApplicationListAttributes) SetProductAnalyticsReplaySampleRate(v int32) {
	o.ProductAnalyticsReplaySampleRate = &v
}

// GetProductScales returns the ProductScales field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetProductScales() RUMProductScales {
	if o == nil || o.ProductScales == nil {
		var ret RUMProductScales
		return ret
	}
	return *o.ProductScales
}

// GetProductScalesOk returns a tuple with the ProductScales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetProductScalesOk() (*RUMProductScales, bool) {
	if o == nil || o.ProductScales == nil {
		return nil, false
	}
	return o.ProductScales, true
}

// HasProductScales returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasProductScales() bool {
	return o != nil && o.ProductScales != nil
}

// SetProductScales gets a reference to the given RUMProductScales and assigns it to the ProductScales field.
func (o *RUMApplicationListAttributes) SetProductScales(v RUMProductScales) {
	o.ProductScales = &v
}

// GetRemoteConfigId returns the RemoteConfigId field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetRemoteConfigId() string {
	if o == nil || o.RemoteConfigId == nil {
		var ret string
		return ret
	}
	return *o.RemoteConfigId
}

// GetRemoteConfigIdOk returns a tuple with the RemoteConfigId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetRemoteConfigIdOk() (*string, bool) {
	if o == nil || o.RemoteConfigId == nil {
		return nil, false
	}
	return o.RemoteConfigId, true
}

// HasRemoteConfigId returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasRemoteConfigId() bool {
	return o != nil && o.RemoteConfigId != nil
}

// SetRemoteConfigId gets a reference to the given string and assigns it to the RemoteConfigId field.
func (o *RUMApplicationListAttributes) SetRemoteConfigId(v string) {
	o.RemoteConfigId = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasShortName() bool {
	return o != nil && o.ShortName != nil
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *RUMApplicationListAttributes) SetShortName(v string) {
	o.ShortName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RUMApplicationListAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RUMApplicationListAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *RUMApplicationListAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value.
func (o *RUMApplicationListAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RUMApplicationListAttributes) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *RUMApplicationListAttributes) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *RUMApplicationListAttributes) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetUpdatedByHandle returns the UpdatedByHandle field value.
func (o *RUMApplicationListAttributes) GetUpdatedByHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedByHandle
}

// GetUpdatedByHandleOk returns a tuple with the UpdatedByHandle field value
// and a boolean to check if the value has been set.
func (o *RUMApplicationListAttributes) GetUpdatedByHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedByHandle, true
}

// SetUpdatedByHandle sets field value.
func (o *RUMApplicationListAttributes) SetUpdatedByHandle(v string) {
	o.UpdatedByHandle = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RUMApplicationListAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiKeyId != nil {
		toSerialize["api_key_id"] = o.ApiKeyId
	}
	if o.ApmRumFlatSamplingReplayEnabled != nil {
		toSerialize["apm_rum_flat_sampling_replay_enabled"] = o.ApmRumFlatSamplingReplayEnabled
	}
	if o.ApmRumFlatSamplingReplaySampleRate != nil {
		toSerialize["apm_rum_flat_sampling_replay_sample_rate"] = o.ApmRumFlatSamplingReplaySampleRate
	}
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["created_by_handle"] = o.CreatedByHandle
	if o.DisabledAt != nil {
		toSerialize["disabled_at"] = o.DisabledAt
	}
	if o.DisabledByHandle != nil {
		toSerialize["disabled_by_handle"] = o.DisabledByHandle
	}
	if o.ErrorTrackingExclusionFilterEnabled != nil {
		toSerialize["error_tracking_exclusion_filter_enabled"] = o.ErrorTrackingExclusionFilterEnabled
	}
	if o.ExternalFeeds != nil {
		toSerialize["external_feeds"] = o.ExternalFeeds
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	toSerialize["name"] = o.Name
	if o.OotbMetricsInstalled != nil {
		toSerialize["ootb_metrics_installed"] = o.OotbMetricsInstalled
	}
	toSerialize["org_id"] = o.OrgId
	if o.ProductAnalyticsPreviewDisabled != nil {
		toSerialize["product_analytics_preview_disabled"] = o.ProductAnalyticsPreviewDisabled
	}
	if o.ProductAnalyticsReplaySampleRate != nil {
		toSerialize["product_analytics_replay_sample_rate"] = o.ProductAnalyticsReplaySampleRate
	}
	if o.ProductScales != nil {
		toSerialize["product_scales"] = o.ProductScales
	}
	if o.RemoteConfigId != nil {
		toSerialize["remote_config_id"] = o.RemoteConfigId
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["updated_by_handle"] = o.UpdatedByHandle

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RUMApplicationListAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKeyId                            *int32            `json:"api_key_id,omitempty"`
		ApmRumFlatSamplingReplayEnabled     *bool             `json:"apm_rum_flat_sampling_replay_enabled,omitempty"`
		ApmRumFlatSamplingReplaySampleRate  *float64          `json:"apm_rum_flat_sampling_replay_sample_rate,omitempty"`
		ApplicationId                       *string           `json:"application_id"`
		CreatedAt                           *int64            `json:"created_at"`
		CreatedByHandle                     *string           `json:"created_by_handle"`
		DisabledAt                          *int64            `json:"disabled_at,omitempty"`
		DisabledByHandle                    *string           `json:"disabled_by_handle,omitempty"`
		ErrorTrackingExclusionFilterEnabled *bool             `json:"error_tracking_exclusion_filter_enabled,omitempty"`
		ExternalFeeds                       []string          `json:"external_feeds,omitempty"`
		Hash                                *string           `json:"hash,omitempty"`
		IsActive                            *bool             `json:"is_active,omitempty"`
		Name                                *string           `json:"name"`
		OotbMetricsInstalled                *bool             `json:"ootb_metrics_installed,omitempty"`
		OrgId                               *int32            `json:"org_id"`
		ProductAnalyticsPreviewDisabled     *bool             `json:"product_analytics_preview_disabled,omitempty"`
		ProductAnalyticsReplaySampleRate    *int32            `json:"product_analytics_replay_sample_rate,omitempty"`
		ProductScales                       *RUMProductScales `json:"product_scales,omitempty"`
		RemoteConfigId                      *string           `json:"remote_config_id,omitempty"`
		ShortName                           *string           `json:"short_name,omitempty"`
		Tags                                []string          `json:"tags,omitempty"`
		Type                                *string           `json:"type"`
		UpdatedAt                           *int64            `json:"updated_at"`
		UpdatedByHandle                     *string           `json:"updated_by_handle"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedByHandle == nil {
		return fmt.Errorf("required field created_by_handle missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	if all.UpdatedByHandle == nil {
		return fmt.Errorf("required field updated_by_handle missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key_id", "apm_rum_flat_sampling_replay_enabled", "apm_rum_flat_sampling_replay_sample_rate", "application_id", "created_at", "created_by_handle", "disabled_at", "disabled_by_handle", "error_tracking_exclusion_filter_enabled", "external_feeds", "hash", "is_active", "name", "ootb_metrics_installed", "org_id", "product_analytics_preview_disabled", "product_analytics_replay_sample_rate", "product_scales", "remote_config_id", "short_name", "tags", "type", "updated_at", "updated_by_handle"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKeyId = all.ApiKeyId
	o.ApmRumFlatSamplingReplayEnabled = all.ApmRumFlatSamplingReplayEnabled
	o.ApmRumFlatSamplingReplaySampleRate = all.ApmRumFlatSamplingReplaySampleRate
	o.ApplicationId = *all.ApplicationId
	o.CreatedAt = *all.CreatedAt
	o.CreatedByHandle = *all.CreatedByHandle
	o.DisabledAt = all.DisabledAt
	o.DisabledByHandle = all.DisabledByHandle
	o.ErrorTrackingExclusionFilterEnabled = all.ErrorTrackingExclusionFilterEnabled
	o.ExternalFeeds = all.ExternalFeeds
	o.Hash = all.Hash
	o.IsActive = all.IsActive
	o.Name = *all.Name
	o.OotbMetricsInstalled = all.OotbMetricsInstalled
	o.OrgId = *all.OrgId
	o.ProductAnalyticsPreviewDisabled = all.ProductAnalyticsPreviewDisabled
	o.ProductAnalyticsReplaySampleRate = all.ProductAnalyticsReplaySampleRate
	if all.ProductScales != nil && all.ProductScales.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ProductScales = all.ProductScales
	o.RemoteConfigId = all.RemoteConfigId
	o.ShortName = all.ShortName
	o.Tags = all.Tags
	o.Type = *all.Type
	o.UpdatedAt = *all.UpdatedAt
	o.UpdatedByHandle = *all.UpdatedByHandle

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
