// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountPatchAttributes AWS Account attributes for patching
type AWSAccountPatchAttributes struct {
	// Tags to apply to all metrics in the account
	AccountTags []string `json:"account_tags,omitempty"`
	// AWS Authentication config
	AuthConfig *AWSAuthConfig `json:"auth_config,omitempty"`
	// AWS Account ID
	AwsAccountId *string `json:"aws_account_id,omitempty"`
	// AWS Account name
	AwsAccountName *string `json:"aws_account_name,omitempty"`
	// AWS Regions to collect data from
	AwsRegions *AWSRegionsList `json:"aws_regions,omitempty"`
	// Creation date of the account
	CreatedAt *string `json:"created_at,omitempty"`
	// AWS Logs config
	LogsConfig *AWSLogs `json:"logs_config,omitempty"`
	// AWS Metrics config
	MetricsConfig *AWSMetrics `json:"metrics_config,omitempty"`
	// Last modification date of the account
	ModifiedAt *string `json:"modified_at,omitempty"`
	// AWS Resources config
	ResourcesConfig *AWSResources `json:"resources_config,omitempty"`
	// AWS Traces config
	TracesConfig *AWSTraces `json:"traces_config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSAccountPatchAttributes instantiates a new AWSAccountPatchAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAccountPatchAttributes() *AWSAccountPatchAttributes {
	this := AWSAccountPatchAttributes{}
	return &this
}

// NewAWSAccountPatchAttributesWithDefaults instantiates a new AWSAccountPatchAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAccountPatchAttributesWithDefaults() *AWSAccountPatchAttributes {
	this := AWSAccountPatchAttributes{}
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetAccountTags() []string {
	if o == nil || o.AccountTags == nil {
		var ret []string
		return ret
	}
	return o.AccountTags
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetAccountTagsOk() (*[]string, bool) {
	if o == nil || o.AccountTags == nil {
		return nil, false
	}
	return &o.AccountTags, true
}

// HasAccountTags returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasAccountTags() bool {
	return o != nil && o.AccountTags != nil
}

// SetAccountTags gets a reference to the given []string and assigns it to the AccountTags field.
func (o *AWSAccountPatchAttributes) SetAccountTags(v []string) {
	o.AccountTags = v
}

// GetAuthConfig returns the AuthConfig field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetAuthConfig() AWSAuthConfig {
	if o == nil || o.AuthConfig == nil {
		var ret AWSAuthConfig
		return ret
	}
	return *o.AuthConfig
}

// GetAuthConfigOk returns a tuple with the AuthConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetAuthConfigOk() (*AWSAuthConfig, bool) {
	if o == nil || o.AuthConfig == nil {
		return nil, false
	}
	return o.AuthConfig, true
}

// HasAuthConfig returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasAuthConfig() bool {
	return o != nil && o.AuthConfig != nil
}

// SetAuthConfig gets a reference to the given AWSAuthConfig and assigns it to the AuthConfig field.
func (o *AWSAccountPatchAttributes) SetAuthConfig(v AWSAuthConfig) {
	o.AuthConfig = &v
}

// GetAwsAccountId returns the AwsAccountId field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetAwsAccountId() string {
	if o == nil || o.AwsAccountId == nil {
		var ret string
		return ret
	}
	return *o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetAwsAccountIdOk() (*string, bool) {
	if o == nil || o.AwsAccountId == nil {
		return nil, false
	}
	return o.AwsAccountId, true
}

// HasAwsAccountId returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasAwsAccountId() bool {
	return o != nil && o.AwsAccountId != nil
}

// SetAwsAccountId gets a reference to the given string and assigns it to the AwsAccountId field.
func (o *AWSAccountPatchAttributes) SetAwsAccountId(v string) {
	o.AwsAccountId = &v
}

// GetAwsAccountName returns the AwsAccountName field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetAwsAccountName() string {
	if o == nil || o.AwsAccountName == nil {
		var ret string
		return ret
	}
	return *o.AwsAccountName
}

// GetAwsAccountNameOk returns a tuple with the AwsAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetAwsAccountNameOk() (*string, bool) {
	if o == nil || o.AwsAccountName == nil {
		return nil, false
	}
	return o.AwsAccountName, true
}

// HasAwsAccountName returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasAwsAccountName() bool {
	return o != nil && o.AwsAccountName != nil
}

// SetAwsAccountName gets a reference to the given string and assigns it to the AwsAccountName field.
func (o *AWSAccountPatchAttributes) SetAwsAccountName(v string) {
	o.AwsAccountName = &v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetAwsRegions() AWSRegionsList {
	if o == nil || o.AwsRegions == nil {
		var ret AWSRegionsList
		return ret
	}
	return *o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetAwsRegionsOk() (*AWSRegionsList, bool) {
	if o == nil || o.AwsRegions == nil {
		return nil, false
	}
	return o.AwsRegions, true
}

// HasAwsRegions returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasAwsRegions() bool {
	return o != nil && o.AwsRegions != nil
}

// SetAwsRegions gets a reference to the given AWSRegionsList and assigns it to the AwsRegions field.
func (o *AWSAccountPatchAttributes) SetAwsRegions(v AWSRegionsList) {
	o.AwsRegions = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AWSAccountPatchAttributes) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetLogsConfig() AWSLogs {
	if o == nil || o.LogsConfig == nil {
		var ret AWSLogs
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetLogsConfigOk() (*AWSLogs, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given AWSLogs and assigns it to the LogsConfig field.
func (o *AWSAccountPatchAttributes) SetLogsConfig(v AWSLogs) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetMetricsConfig() AWSMetrics {
	if o == nil || o.MetricsConfig == nil {
		var ret AWSMetrics
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetMetricsConfigOk() (*AWSMetrics, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given AWSMetrics and assigns it to the MetricsConfig field.
func (o *AWSAccountPatchAttributes) SetMetricsConfig(v AWSMetrics) {
	o.MetricsConfig = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *AWSAccountPatchAttributes) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetResourcesConfig returns the ResourcesConfig field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetResourcesConfig() AWSResources {
	if o == nil || o.ResourcesConfig == nil {
		var ret AWSResources
		return ret
	}
	return *o.ResourcesConfig
}

// GetResourcesConfigOk returns a tuple with the ResourcesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetResourcesConfigOk() (*AWSResources, bool) {
	if o == nil || o.ResourcesConfig == nil {
		return nil, false
	}
	return o.ResourcesConfig, true
}

// HasResourcesConfig returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasResourcesConfig() bool {
	return o != nil && o.ResourcesConfig != nil
}

// SetResourcesConfig gets a reference to the given AWSResources and assigns it to the ResourcesConfig field.
func (o *AWSAccountPatchAttributes) SetResourcesConfig(v AWSResources) {
	o.ResourcesConfig = &v
}

// GetTracesConfig returns the TracesConfig field value if set, zero value otherwise.
func (o *AWSAccountPatchAttributes) GetTracesConfig() AWSTraces {
	if o == nil || o.TracesConfig == nil {
		var ret AWSTraces
		return ret
	}
	return *o.TracesConfig
}

// GetTracesConfigOk returns a tuple with the TracesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountPatchAttributes) GetTracesConfigOk() (*AWSTraces, bool) {
	if o == nil || o.TracesConfig == nil {
		return nil, false
	}
	return o.TracesConfig, true
}

// HasTracesConfig returns a boolean if a field has been set.
func (o *AWSAccountPatchAttributes) HasTracesConfig() bool {
	return o != nil && o.TracesConfig != nil
}

// SetTracesConfig gets a reference to the given AWSTraces and assigns it to the TracesConfig field.
func (o *AWSAccountPatchAttributes) SetTracesConfig(v AWSTraces) {
	o.TracesConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAccountPatchAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountTags != nil {
		toSerialize["account_tags"] = o.AccountTags
	}
	if o.AuthConfig != nil {
		toSerialize["auth_config"] = o.AuthConfig
	}
	if o.AwsAccountId != nil {
		toSerialize["aws_account_id"] = o.AwsAccountId
	}
	if o.AwsAccountName != nil {
		toSerialize["aws_account_name"] = o.AwsAccountName
	}
	if o.AwsRegions != nil {
		toSerialize["aws_regions"] = o.AwsRegions
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LogsConfig != nil {
		toSerialize["logs_config"] = o.LogsConfig
	}
	if o.MetricsConfig != nil {
		toSerialize["metrics_config"] = o.MetricsConfig
	}
	if o.ModifiedAt != nil {
		toSerialize["modified_at"] = o.ModifiedAt
	}
	if o.ResourcesConfig != nil {
		toSerialize["resources_config"] = o.ResourcesConfig
	}
	if o.TracesConfig != nil {
		toSerialize["traces_config"] = o.TracesConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAccountPatchAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags     []string        `json:"account_tags,omitempty"`
		AuthConfig      *AWSAuthConfig  `json:"auth_config,omitempty"`
		AwsAccountId    *string         `json:"aws_account_id,omitempty"`
		AwsAccountName  *string         `json:"aws_account_name,omitempty"`
		AwsRegions      *AWSRegionsList `json:"aws_regions,omitempty"`
		CreatedAt       *string         `json:"created_at,omitempty"`
		LogsConfig      *AWSLogs        `json:"logs_config,omitempty"`
		MetricsConfig   *AWSMetrics     `json:"metrics_config,omitempty"`
		ModifiedAt      *string         `json:"modified_at,omitempty"`
		ResourcesConfig *AWSResources   `json:"resources_config,omitempty"`
		TracesConfig    *AWSTraces      `json:"traces_config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_tags", "auth_config", "aws_account_id", "aws_account_name", "aws_regions", "created_at", "logs_config", "metrics_config", "modified_at", "resources_config", "traces_config"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountTags = all.AccountTags
	if all.AuthConfig != nil && all.AuthConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuthConfig = all.AuthConfig
	o.AwsAccountId = all.AwsAccountId
	o.AwsAccountName = all.AwsAccountName
	if all.AwsRegions != nil && all.AwsRegions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AwsRegions = all.AwsRegions
	o.CreatedAt = all.CreatedAt
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
	o.ModifiedAt = all.ModifiedAt
	if all.ResourcesConfig != nil && all.ResourcesConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ResourcesConfig = all.ResourcesConfig
	if all.TracesConfig != nil && all.TracesConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TracesConfig = all.TracesConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
