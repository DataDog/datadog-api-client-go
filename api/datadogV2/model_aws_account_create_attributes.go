// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountCreateAttributes AWS Account attributes for creation
type AWSAccountCreateAttributes struct {
	// Tags to apply to all metrics in the account
	AccountTags []string `json:"account_tags,omitempty"`
	// AWS Authentication config
	AuthConfig *AWSAuthConfig `json:"auth_config,omitempty"`
	// AWS Account ID
	AwsAccountId string `json:"aws_account_id"`
	// AWS Regions to collect data from
	AwsRegions *AWSRegionsList `json:"aws_regions,omitempty"`
	// AWS Logs config
	LogsConfig *AWSLogs `json:"logs_config,omitempty"`
	// AWS Metrics config
	MetricsConfig *AWSMetrics `json:"metrics_config,omitempty"`
	// AWS Resources config
	ResourcesConfig *AWSResources `json:"resources_config,omitempty"`
	// AWS Traces config
	TracesConfig *AWSTraces `json:"traces_config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSAccountCreateAttributes instantiates a new AWSAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAccountCreateAttributes(awsAccountId string) *AWSAccountCreateAttributes {
	this := AWSAccountCreateAttributes{}
	this.AwsAccountId = awsAccountId
	return &this
}

// NewAWSAccountCreateAttributesWithDefaults instantiates a new AWSAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAccountCreateAttributesWithDefaults() *AWSAccountCreateAttributes {
	this := AWSAccountCreateAttributes{}
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise.
func (o *AWSAccountCreateAttributes) GetAccountTags() []string {
	if o == nil || o.AccountTags == nil {
		var ret []string
		return ret
	}
	return o.AccountTags
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetAccountTagsOk() (*[]string, bool) {
	if o == nil || o.AccountTags == nil {
		return nil, false
	}
	return &o.AccountTags, true
}

// HasAccountTags returns a boolean if a field has been set.
func (o *AWSAccountCreateAttributes) HasAccountTags() bool {
	return o != nil && o.AccountTags != nil
}

// SetAccountTags gets a reference to the given []string and assigns it to the AccountTags field.
func (o *AWSAccountCreateAttributes) SetAccountTags(v []string) {
	o.AccountTags = v
}

// GetAuthConfig returns the AuthConfig field value if set, zero value otherwise.
func (o *AWSAccountCreateAttributes) GetAuthConfig() AWSAuthConfig {
	if o == nil || o.AuthConfig == nil {
		var ret AWSAuthConfig
		return ret
	}
	return *o.AuthConfig
}

// GetAuthConfigOk returns a tuple with the AuthConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetAuthConfigOk() (*AWSAuthConfig, bool) {
	if o == nil || o.AuthConfig == nil {
		return nil, false
	}
	return o.AuthConfig, true
}

// HasAuthConfig returns a boolean if a field has been set.
func (o *AWSAccountCreateAttributes) HasAuthConfig() bool {
	return o != nil && o.AuthConfig != nil
}

// SetAuthConfig gets a reference to the given AWSAuthConfig and assigns it to the AuthConfig field.
func (o *AWSAccountCreateAttributes) SetAuthConfig(v AWSAuthConfig) {
	o.AuthConfig = &v
}

// GetAwsAccountId returns the AwsAccountId field value.
func (o *AWSAccountCreateAttributes) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value.
func (o *AWSAccountCreateAttributes) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *AWSAccountCreateAttributes) GetAwsRegions() AWSRegionsList {
	if o == nil || o.AwsRegions == nil {
		var ret AWSRegionsList
		return ret
	}
	return *o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetAwsRegionsOk() (*AWSRegionsList, bool) {
	if o == nil || o.AwsRegions == nil {
		return nil, false
	}
	return o.AwsRegions, true
}

// HasAwsRegions returns a boolean if a field has been set.
func (o *AWSAccountCreateAttributes) HasAwsRegions() bool {
	return o != nil && o.AwsRegions != nil
}

// SetAwsRegions gets a reference to the given AWSRegionsList and assigns it to the AwsRegions field.
func (o *AWSAccountCreateAttributes) SetAwsRegions(v AWSRegionsList) {
	o.AwsRegions = &v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *AWSAccountCreateAttributes) GetLogsConfig() AWSLogs {
	if o == nil || o.LogsConfig == nil {
		var ret AWSLogs
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetLogsConfigOk() (*AWSLogs, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *AWSAccountCreateAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given AWSLogs and assigns it to the LogsConfig field.
func (o *AWSAccountCreateAttributes) SetLogsConfig(v AWSLogs) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *AWSAccountCreateAttributes) GetMetricsConfig() AWSMetrics {
	if o == nil || o.MetricsConfig == nil {
		var ret AWSMetrics
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetMetricsConfigOk() (*AWSMetrics, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *AWSAccountCreateAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given AWSMetrics and assigns it to the MetricsConfig field.
func (o *AWSAccountCreateAttributes) SetMetricsConfig(v AWSMetrics) {
	o.MetricsConfig = &v
}

// GetResourcesConfig returns the ResourcesConfig field value if set, zero value otherwise.
func (o *AWSAccountCreateAttributes) GetResourcesConfig() AWSResources {
	if o == nil || o.ResourcesConfig == nil {
		var ret AWSResources
		return ret
	}
	return *o.ResourcesConfig
}

// GetResourcesConfigOk returns a tuple with the ResourcesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetResourcesConfigOk() (*AWSResources, bool) {
	if o == nil || o.ResourcesConfig == nil {
		return nil, false
	}
	return o.ResourcesConfig, true
}

// HasResourcesConfig returns a boolean if a field has been set.
func (o *AWSAccountCreateAttributes) HasResourcesConfig() bool {
	return o != nil && o.ResourcesConfig != nil
}

// SetResourcesConfig gets a reference to the given AWSResources and assigns it to the ResourcesConfig field.
func (o *AWSAccountCreateAttributes) SetResourcesConfig(v AWSResources) {
	o.ResourcesConfig = &v
}

// GetTracesConfig returns the TracesConfig field value if set, zero value otherwise.
func (o *AWSAccountCreateAttributes) GetTracesConfig() AWSTraces {
	if o == nil || o.TracesConfig == nil {
		var ret AWSTraces
		return ret
	}
	return *o.TracesConfig
}

// GetTracesConfigOk returns a tuple with the TracesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountCreateAttributes) GetTracesConfigOk() (*AWSTraces, bool) {
	if o == nil || o.TracesConfig == nil {
		return nil, false
	}
	return o.TracesConfig, true
}

// HasTracesConfig returns a boolean if a field has been set.
func (o *AWSAccountCreateAttributes) HasTracesConfig() bool {
	return o != nil && o.TracesConfig != nil
}

// SetTracesConfig gets a reference to the given AWSTraces and assigns it to the TracesConfig field.
func (o *AWSAccountCreateAttributes) SetTracesConfig(v AWSTraces) {
	o.TracesConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAccountCreateAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["aws_account_id"] = o.AwsAccountId
	if o.AwsRegions != nil {
		toSerialize["aws_regions"] = o.AwsRegions
	}
	if o.LogsConfig != nil {
		toSerialize["logs_config"] = o.LogsConfig
	}
	if o.MetricsConfig != nil {
		toSerialize["metrics_config"] = o.MetricsConfig
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
func (o *AWSAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags     []string        `json:"account_tags,omitempty"`
		AuthConfig      *AWSAuthConfig  `json:"auth_config,omitempty"`
		AwsAccountId    *string         `json:"aws_account_id"`
		AwsRegions      *AWSRegionsList `json:"aws_regions,omitempty"`
		LogsConfig      *AWSLogs        `json:"logs_config,omitempty"`
		MetricsConfig   *AWSMetrics     `json:"metrics_config,omitempty"`
		ResourcesConfig *AWSResources   `json:"resources_config,omitempty"`
		TracesConfig    *AWSTraces      `json:"traces_config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AwsAccountId == nil {
		return fmt.Errorf("required field aws_account_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_tags", "auth_config", "aws_account_id", "aws_regions", "logs_config", "metrics_config", "resources_config", "traces_config"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountTags = all.AccountTags
	if all.AuthConfig != nil && all.AuthConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AuthConfig = all.AuthConfig
	o.AwsAccountId = *all.AwsAccountId
	if all.AwsRegions != nil && all.AwsRegions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AwsRegions = all.AwsRegions
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
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
