// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchStepsRequest Steps of the launch. The tracking plan is required.
type ProductAnalyticsLaunchStepsRequest struct {
	// Dashboard linked to the launch.
	Dashboard *ProductAnalyticsLaunchDashboardStep `json:"dashboard,omitempty"`
	// Experiment linked to the launch. The draft fields are cleared when `experiment_id` is set.
	Experiment *ProductAnalyticsLaunchExperimentStep `json:"experiment,omitempty"`
	// Funnel linked to the launch.
	Funnel *ProductAnalyticsLaunchFunnelStep `json:"funnel,omitempty"`
	// Tracking plan of a launch, listing the events and properties to instrument.
	TrackingPlan ProductAnalyticsLaunchTrackingPlanRequest `json:"tracking_plan"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchStepsRequest instantiates a new ProductAnalyticsLaunchStepsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchStepsRequest(trackingPlan ProductAnalyticsLaunchTrackingPlanRequest) *ProductAnalyticsLaunchStepsRequest {
	this := ProductAnalyticsLaunchStepsRequest{}
	this.TrackingPlan = trackingPlan
	return &this
}

// NewProductAnalyticsLaunchStepsRequestWithDefaults instantiates a new ProductAnalyticsLaunchStepsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchStepsRequestWithDefaults() *ProductAnalyticsLaunchStepsRequest {
	this := ProductAnalyticsLaunchStepsRequest{}
	return &this
}

// GetDashboard returns the Dashboard field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchStepsRequest) GetDashboard() ProductAnalyticsLaunchDashboardStep {
	if o == nil || o.Dashboard == nil {
		var ret ProductAnalyticsLaunchDashboardStep
		return ret
	}
	return *o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchStepsRequest) GetDashboardOk() (*ProductAnalyticsLaunchDashboardStep, bool) {
	if o == nil || o.Dashboard == nil {
		return nil, false
	}
	return o.Dashboard, true
}

// HasDashboard returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchStepsRequest) HasDashboard() bool {
	return o != nil && o.Dashboard != nil
}

// SetDashboard gets a reference to the given ProductAnalyticsLaunchDashboardStep and assigns it to the Dashboard field.
func (o *ProductAnalyticsLaunchStepsRequest) SetDashboard(v ProductAnalyticsLaunchDashboardStep) {
	o.Dashboard = &v
}

// GetExperiment returns the Experiment field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchStepsRequest) GetExperiment() ProductAnalyticsLaunchExperimentStep {
	if o == nil || o.Experiment == nil {
		var ret ProductAnalyticsLaunchExperimentStep
		return ret
	}
	return *o.Experiment
}

// GetExperimentOk returns a tuple with the Experiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchStepsRequest) GetExperimentOk() (*ProductAnalyticsLaunchExperimentStep, bool) {
	if o == nil || o.Experiment == nil {
		return nil, false
	}
	return o.Experiment, true
}

// HasExperiment returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchStepsRequest) HasExperiment() bool {
	return o != nil && o.Experiment != nil
}

// SetExperiment gets a reference to the given ProductAnalyticsLaunchExperimentStep and assigns it to the Experiment field.
func (o *ProductAnalyticsLaunchStepsRequest) SetExperiment(v ProductAnalyticsLaunchExperimentStep) {
	o.Experiment = &v
}

// GetFunnel returns the Funnel field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchStepsRequest) GetFunnel() ProductAnalyticsLaunchFunnelStep {
	if o == nil || o.Funnel == nil {
		var ret ProductAnalyticsLaunchFunnelStep
		return ret
	}
	return *o.Funnel
}

// GetFunnelOk returns a tuple with the Funnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchStepsRequest) GetFunnelOk() (*ProductAnalyticsLaunchFunnelStep, bool) {
	if o == nil || o.Funnel == nil {
		return nil, false
	}
	return o.Funnel, true
}

// HasFunnel returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchStepsRequest) HasFunnel() bool {
	return o != nil && o.Funnel != nil
}

// SetFunnel gets a reference to the given ProductAnalyticsLaunchFunnelStep and assigns it to the Funnel field.
func (o *ProductAnalyticsLaunchStepsRequest) SetFunnel(v ProductAnalyticsLaunchFunnelStep) {
	o.Funnel = &v
}

// GetTrackingPlan returns the TrackingPlan field value.
func (o *ProductAnalyticsLaunchStepsRequest) GetTrackingPlan() ProductAnalyticsLaunchTrackingPlanRequest {
	if o == nil {
		var ret ProductAnalyticsLaunchTrackingPlanRequest
		return ret
	}
	return o.TrackingPlan
}

// GetTrackingPlanOk returns a tuple with the TrackingPlan field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchStepsRequest) GetTrackingPlanOk() (*ProductAnalyticsLaunchTrackingPlanRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingPlan, true
}

// SetTrackingPlan sets field value.
func (o *ProductAnalyticsLaunchStepsRequest) SetTrackingPlan(v ProductAnalyticsLaunchTrackingPlanRequest) {
	o.TrackingPlan = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchStepsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Dashboard != nil {
		toSerialize["dashboard"] = o.Dashboard
	}
	if o.Experiment != nil {
		toSerialize["experiment"] = o.Experiment
	}
	if o.Funnel != nil {
		toSerialize["funnel"] = o.Funnel
	}
	toSerialize["tracking_plan"] = o.TrackingPlan

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchStepsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dashboard    *ProductAnalyticsLaunchDashboardStep       `json:"dashboard,omitempty"`
		Experiment   *ProductAnalyticsLaunchExperimentStep      `json:"experiment,omitempty"`
		Funnel       *ProductAnalyticsLaunchFunnelStep          `json:"funnel,omitempty"`
		TrackingPlan *ProductAnalyticsLaunchTrackingPlanRequest `json:"tracking_plan"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.TrackingPlan == nil {
		return fmt.Errorf("required field tracking_plan missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dashboard", "experiment", "funnel", "tracking_plan"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Dashboard != nil && all.Dashboard.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dashboard = all.Dashboard
	if all.Experiment != nil && all.Experiment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Experiment = all.Experiment
	if all.Funnel != nil && all.Funnel.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Funnel = all.Funnel
	if all.TrackingPlan.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TrackingPlan = *all.TrackingPlan

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
