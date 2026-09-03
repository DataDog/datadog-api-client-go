// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentAveragedMetrics Averaged DORA and delivery metrics computed across the commits and pull requests included in the deployment.
type DORADeploymentAveragedMetrics struct {
	// The averaged change lead time, in seconds.
	ChangeLeadTime *int64 `json:"change_lead_time,omitempty"`
	// The averaged merge time, in seconds.
	MergeTime *int64 `json:"merge_time,omitempty"`
	// The averaged review time, in seconds.
	ReviewTime *int64 `json:"review_time,omitempty"`
	// The averaged time to deploy, in seconds.
	TimeToDeploy *int64 `json:"time_to_deploy,omitempty"`
	// The averaged time until the pull request was ready for review, in seconds.
	TimeToPrReady *int64 `json:"time_to_pr_ready,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORADeploymentAveragedMetrics instantiates a new DORADeploymentAveragedMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentAveragedMetrics() *DORADeploymentAveragedMetrics {
	this := DORADeploymentAveragedMetrics{}
	return &this
}

// NewDORADeploymentAveragedMetricsWithDefaults instantiates a new DORADeploymentAveragedMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentAveragedMetricsWithDefaults() *DORADeploymentAveragedMetrics {
	this := DORADeploymentAveragedMetrics{}
	return &this
}

// GetChangeLeadTime returns the ChangeLeadTime field value if set, zero value otherwise.
func (o *DORADeploymentAveragedMetrics) GetChangeLeadTime() int64 {
	if o == nil || o.ChangeLeadTime == nil {
		var ret int64
		return ret
	}
	return *o.ChangeLeadTime
}

// GetChangeLeadTimeOk returns a tuple with the ChangeLeadTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentAveragedMetrics) GetChangeLeadTimeOk() (*int64, bool) {
	if o == nil || o.ChangeLeadTime == nil {
		return nil, false
	}
	return o.ChangeLeadTime, true
}

// HasChangeLeadTime returns a boolean if a field has been set.
func (o *DORADeploymentAveragedMetrics) HasChangeLeadTime() bool {
	return o != nil && o.ChangeLeadTime != nil
}

// SetChangeLeadTime gets a reference to the given int64 and assigns it to the ChangeLeadTime field.
func (o *DORADeploymentAveragedMetrics) SetChangeLeadTime(v int64) {
	o.ChangeLeadTime = &v
}

// GetMergeTime returns the MergeTime field value if set, zero value otherwise.
func (o *DORADeploymentAveragedMetrics) GetMergeTime() int64 {
	if o == nil || o.MergeTime == nil {
		var ret int64
		return ret
	}
	return *o.MergeTime
}

// GetMergeTimeOk returns a tuple with the MergeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentAveragedMetrics) GetMergeTimeOk() (*int64, bool) {
	if o == nil || o.MergeTime == nil {
		return nil, false
	}
	return o.MergeTime, true
}

// HasMergeTime returns a boolean if a field has been set.
func (o *DORADeploymentAveragedMetrics) HasMergeTime() bool {
	return o != nil && o.MergeTime != nil
}

// SetMergeTime gets a reference to the given int64 and assigns it to the MergeTime field.
func (o *DORADeploymentAveragedMetrics) SetMergeTime(v int64) {
	o.MergeTime = &v
}

// GetReviewTime returns the ReviewTime field value if set, zero value otherwise.
func (o *DORADeploymentAveragedMetrics) GetReviewTime() int64 {
	if o == nil || o.ReviewTime == nil {
		var ret int64
		return ret
	}
	return *o.ReviewTime
}

// GetReviewTimeOk returns a tuple with the ReviewTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentAveragedMetrics) GetReviewTimeOk() (*int64, bool) {
	if o == nil || o.ReviewTime == nil {
		return nil, false
	}
	return o.ReviewTime, true
}

// HasReviewTime returns a boolean if a field has been set.
func (o *DORADeploymentAveragedMetrics) HasReviewTime() bool {
	return o != nil && o.ReviewTime != nil
}

// SetReviewTime gets a reference to the given int64 and assigns it to the ReviewTime field.
func (o *DORADeploymentAveragedMetrics) SetReviewTime(v int64) {
	o.ReviewTime = &v
}

// GetTimeToDeploy returns the TimeToDeploy field value if set, zero value otherwise.
func (o *DORADeploymentAveragedMetrics) GetTimeToDeploy() int64 {
	if o == nil || o.TimeToDeploy == nil {
		var ret int64
		return ret
	}
	return *o.TimeToDeploy
}

// GetTimeToDeployOk returns a tuple with the TimeToDeploy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentAveragedMetrics) GetTimeToDeployOk() (*int64, bool) {
	if o == nil || o.TimeToDeploy == nil {
		return nil, false
	}
	return o.TimeToDeploy, true
}

// HasTimeToDeploy returns a boolean if a field has been set.
func (o *DORADeploymentAveragedMetrics) HasTimeToDeploy() bool {
	return o != nil && o.TimeToDeploy != nil
}

// SetTimeToDeploy gets a reference to the given int64 and assigns it to the TimeToDeploy field.
func (o *DORADeploymentAveragedMetrics) SetTimeToDeploy(v int64) {
	o.TimeToDeploy = &v
}

// GetTimeToPrReady returns the TimeToPrReady field value if set, zero value otherwise.
func (o *DORADeploymentAveragedMetrics) GetTimeToPrReady() int64 {
	if o == nil || o.TimeToPrReady == nil {
		var ret int64
		return ret
	}
	return *o.TimeToPrReady
}

// GetTimeToPrReadyOk returns a tuple with the TimeToPrReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentAveragedMetrics) GetTimeToPrReadyOk() (*int64, bool) {
	if o == nil || o.TimeToPrReady == nil {
		return nil, false
	}
	return o.TimeToPrReady, true
}

// HasTimeToPrReady returns a boolean if a field has been set.
func (o *DORADeploymentAveragedMetrics) HasTimeToPrReady() bool {
	return o != nil && o.TimeToPrReady != nil
}

// SetTimeToPrReady gets a reference to the given int64 and assigns it to the TimeToPrReady field.
func (o *DORADeploymentAveragedMetrics) SetTimeToPrReady(v int64) {
	o.TimeToPrReady = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentAveragedMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeLeadTime != nil {
		toSerialize["change_lead_time"] = o.ChangeLeadTime
	}
	if o.MergeTime != nil {
		toSerialize["merge_time"] = o.MergeTime
	}
	if o.ReviewTime != nil {
		toSerialize["review_time"] = o.ReviewTime
	}
	if o.TimeToDeploy != nil {
		toSerialize["time_to_deploy"] = o.TimeToDeploy
	}
	if o.TimeToPrReady != nil {
		toSerialize["time_to_pr_ready"] = o.TimeToPrReady
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentAveragedMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeLeadTime *int64 `json:"change_lead_time,omitempty"`
		MergeTime      *int64 `json:"merge_time,omitempty"`
		ReviewTime     *int64 `json:"review_time,omitempty"`
		TimeToDeploy   *int64 `json:"time_to_deploy,omitempty"`
		TimeToPrReady  *int64 `json:"time_to_pr_ready,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_lead_time", "merge_time", "review_time", "time_to_deploy", "time_to_pr_ready"})
	} else {
		return err
	}
	o.ChangeLeadTime = all.ChangeLeadTime
	o.MergeTime = all.MergeTime
	o.ReviewTime = all.ReviewTime
	o.TimeToDeploy = all.TimeToDeploy
	o.TimeToPrReady = all.TimeToPrReady

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
