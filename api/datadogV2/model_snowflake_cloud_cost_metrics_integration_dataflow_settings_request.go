// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest Settings of the Cloud Cost Management dataflow. Only the fields provided are changed.
type SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest struct {
	// Snowflake query tags to ingest, so that cost data can be broken down by them in Cloud Cost Management. Provide the tag names as a comma-separated list without spaces, using only letters, digits, underscores, dots, and hyphens. Datadog does not collect query tags by default.
	QueryTags *string `json:"query_tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewSnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest instantiates a new SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest() *SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest {
	this := SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest{}
	return &this
}

// NewSnowflakeCloudCostMetricsIntegrationDataflowSettingsRequestWithDefaults instantiates a new SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSnowflakeCloudCostMetricsIntegrationDataflowSettingsRequestWithDefaults() *SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest {
	this := SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest{}
	return &this
}

// GetQueryTags returns the QueryTags field value if set, zero value otherwise.
func (o *SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest) GetQueryTags() string {
	if o == nil || o.QueryTags == nil {
		var ret string
		return ret
	}
	return *o.QueryTags
}

// GetQueryTagsOk returns a tuple with the QueryTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest) GetQueryTagsOk() (*string, bool) {
	if o == nil || o.QueryTags == nil {
		return nil, false
	}
	return o.QueryTags, true
}

// HasQueryTags returns a boolean if a field has been set.
func (o *SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest) HasQueryTags() bool {
	return o != nil && o.QueryTags != nil
}

// SetQueryTags gets a reference to the given string and assigns it to the QueryTags field.
func (o *SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest) SetQueryTags(v string) {
	o.QueryTags = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.QueryTags != nil {
		toSerialize["query_tags"] = o.QueryTags
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SnowflakeCloudCostMetricsIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		QueryTags *string `json:"query_tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.QueryTags = all.QueryTags

	return nil
}
