// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GitHubCostsIntegrationDataflowsResponse Data Datadog collects from GitHub Costs, keyed by dataflow id.
type GitHubCostsIntegrationDataflowsResponse struct {
	// GitHub costs for Actions, Packages, Copilot, storage, and other billed products, estimated from list prices and usage, including discounts when available. Used by [Cloud Cost Management](https://docs.datadoghq.com/cloud_cost_management/). Requires an enterprise or organization on GitHub's enhanced billing platform; without it this dataflow collects no data. This dataflow is always on and cannot be turned off, which is why it appears in responses only.
	GithubCloudCostMetrics *GitHubCostsCloudCostMetricsIntegrationDataflowResponse `json:"github-cloud-cost-metrics,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGitHubCostsIntegrationDataflowsResponse instantiates a new GitHubCostsIntegrationDataflowsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGitHubCostsIntegrationDataflowsResponse() *GitHubCostsIntegrationDataflowsResponse {
	this := GitHubCostsIntegrationDataflowsResponse{}
	return &this
}

// NewGitHubCostsIntegrationDataflowsResponseWithDefaults instantiates a new GitHubCostsIntegrationDataflowsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGitHubCostsIntegrationDataflowsResponseWithDefaults() *GitHubCostsIntegrationDataflowsResponse {
	this := GitHubCostsIntegrationDataflowsResponse{}
	return &this
}

// GetGithubCloudCostMetrics returns the GithubCloudCostMetrics field value if set, zero value otherwise.
func (o *GitHubCostsIntegrationDataflowsResponse) GetGithubCloudCostMetrics() GitHubCostsCloudCostMetricsIntegrationDataflowResponse {
	if o == nil || o.GithubCloudCostMetrics == nil {
		var ret GitHubCostsCloudCostMetricsIntegrationDataflowResponse
		return ret
	}
	return *o.GithubCloudCostMetrics
}

// GetGithubCloudCostMetricsOk returns a tuple with the GithubCloudCostMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitHubCostsIntegrationDataflowsResponse) GetGithubCloudCostMetricsOk() (*GitHubCostsCloudCostMetricsIntegrationDataflowResponse, bool) {
	if o == nil || o.GithubCloudCostMetrics == nil {
		return nil, false
	}
	return o.GithubCloudCostMetrics, true
}

// HasGithubCloudCostMetrics returns a boolean if a field has been set.
func (o *GitHubCostsIntegrationDataflowsResponse) HasGithubCloudCostMetrics() bool {
	return o != nil && o.GithubCloudCostMetrics != nil
}

// SetGithubCloudCostMetrics gets a reference to the given GitHubCostsCloudCostMetricsIntegrationDataflowResponse and assigns it to the GithubCloudCostMetrics field.
func (o *GitHubCostsIntegrationDataflowsResponse) SetGithubCloudCostMetrics(v GitHubCostsCloudCostMetricsIntegrationDataflowResponse) {
	o.GithubCloudCostMetrics = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GitHubCostsIntegrationDataflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GithubCloudCostMetrics != nil {
		toSerialize["github-cloud-cost-metrics"] = o.GithubCloudCostMetrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GitHubCostsIntegrationDataflowsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GithubCloudCostMetrics *GitHubCostsCloudCostMetricsIntegrationDataflowResponse `json:"github-cloud-cost-metrics,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"github-cloud-cost-metrics"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GithubCloudCostMetrics != nil && all.GithubCloudCostMetrics.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GithubCloudCostMetrics = all.GithubCloudCostMetrics

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
