// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsLaunchPatchSteps Steps to update. Steps that are not set are left unchanged.
// A `tracking_plan` replaces the whole stored tracking plan, so include every event and property to keep.
type ProductAnalyticsLaunchPatchSteps struct {
	// Dashboard linked to the launch.
	Dashboard *ProductAnalyticsLaunchDashboardStep `json:"dashboard,omitempty"`
	// Experiment linked to the launch. The draft fields are cleared when `experiment_id` is set.
	Experiment *ProductAnalyticsLaunchExperimentStep `json:"experiment,omitempty"`
	// Funnel linked to the launch.
	Funnel *ProductAnalyticsLaunchFunnelStep `json:"funnel,omitempty"`
	// Tracking plan of a launch, listing the events and properties to instrument.
	TrackingPlan *ProductAnalyticsLaunchTrackingPlanRequest `json:"tracking_plan,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsLaunchPatchSteps instantiates a new ProductAnalyticsLaunchPatchSteps object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsLaunchPatchSteps() *ProductAnalyticsLaunchPatchSteps {
	this := ProductAnalyticsLaunchPatchSteps{}
	return &this
}

// NewProductAnalyticsLaunchPatchStepsWithDefaults instantiates a new ProductAnalyticsLaunchPatchSteps object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsLaunchPatchStepsWithDefaults() *ProductAnalyticsLaunchPatchSteps {
	this := ProductAnalyticsLaunchPatchSteps{}
	return &this
}

// GetDashboard returns the Dashboard field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchSteps) GetDashboard() ProductAnalyticsLaunchDashboardStep {
	if o == nil || o.Dashboard == nil {
		var ret ProductAnalyticsLaunchDashboardStep
		return ret
	}
	return *o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchSteps) GetDashboardOk() (*ProductAnalyticsLaunchDashboardStep, bool) {
	if o == nil || o.Dashboard == nil {
		return nil, false
	}
	return o.Dashboard, true
}

// HasDashboard returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchSteps) HasDashboard() bool {
	return o != nil && o.Dashboard != nil
}

// SetDashboard gets a reference to the given ProductAnalyticsLaunchDashboardStep and assigns it to the Dashboard field.
func (o *ProductAnalyticsLaunchPatchSteps) SetDashboard(v ProductAnalyticsLaunchDashboardStep) {
	o.Dashboard = &v
}

// GetExperiment returns the Experiment field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchSteps) GetExperiment() ProductAnalyticsLaunchExperimentStep {
	if o == nil || o.Experiment == nil {
		var ret ProductAnalyticsLaunchExperimentStep
		return ret
	}
	return *o.Experiment
}

// GetExperimentOk returns a tuple with the Experiment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchSteps) GetExperimentOk() (*ProductAnalyticsLaunchExperimentStep, bool) {
	if o == nil || o.Experiment == nil {
		return nil, false
	}
	return o.Experiment, true
}

// HasExperiment returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchSteps) HasExperiment() bool {
	return o != nil && o.Experiment != nil
}

// SetExperiment gets a reference to the given ProductAnalyticsLaunchExperimentStep and assigns it to the Experiment field.
func (o *ProductAnalyticsLaunchPatchSteps) SetExperiment(v ProductAnalyticsLaunchExperimentStep) {
	o.Experiment = &v
}

// GetFunnel returns the Funnel field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchSteps) GetFunnel() ProductAnalyticsLaunchFunnelStep {
	if o == nil || o.Funnel == nil {
		var ret ProductAnalyticsLaunchFunnelStep
		return ret
	}
	return *o.Funnel
}

// GetFunnelOk returns a tuple with the Funnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchSteps) GetFunnelOk() (*ProductAnalyticsLaunchFunnelStep, bool) {
	if o == nil || o.Funnel == nil {
		return nil, false
	}
	return o.Funnel, true
}

// HasFunnel returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchSteps) HasFunnel() bool {
	return o != nil && o.Funnel != nil
}

// SetFunnel gets a reference to the given ProductAnalyticsLaunchFunnelStep and assigns it to the Funnel field.
func (o *ProductAnalyticsLaunchPatchSteps) SetFunnel(v ProductAnalyticsLaunchFunnelStep) {
	o.Funnel = &v
}

// GetTrackingPlan returns the TrackingPlan field value if set, zero value otherwise.
func (o *ProductAnalyticsLaunchPatchSteps) GetTrackingPlan() ProductAnalyticsLaunchTrackingPlanRequest {
	if o == nil || o.TrackingPlan == nil {
		var ret ProductAnalyticsLaunchTrackingPlanRequest
		return ret
	}
	return *o.TrackingPlan
}

// GetTrackingPlanOk returns a tuple with the TrackingPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsLaunchPatchSteps) GetTrackingPlanOk() (*ProductAnalyticsLaunchTrackingPlanRequest, bool) {
	if o == nil || o.TrackingPlan == nil {
		return nil, false
	}
	return o.TrackingPlan, true
}

// HasTrackingPlan returns a boolean if a field has been set.
func (o *ProductAnalyticsLaunchPatchSteps) HasTrackingPlan() bool {
	return o != nil && o.TrackingPlan != nil
}

// SetTrackingPlan gets a reference to the given ProductAnalyticsLaunchTrackingPlanRequest and assigns it to the TrackingPlan field.
func (o *ProductAnalyticsLaunchPatchSteps) SetTrackingPlan(v ProductAnalyticsLaunchTrackingPlanRequest) {
	o.TrackingPlan = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsLaunchPatchSteps) MarshalJSON() ([]byte, error) {
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
	if o.TrackingPlan != nil {
		toSerialize["tracking_plan"] = o.TrackingPlan
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsLaunchPatchSteps) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dashboard    *ProductAnalyticsLaunchDashboardStep       `json:"dashboard,omitempty"`
		Experiment   *ProductAnalyticsLaunchExperimentStep      `json:"experiment,omitempty"`
		Funnel       *ProductAnalyticsLaunchFunnelStep          `json:"funnel,omitempty"`
		TrackingPlan *ProductAnalyticsLaunchTrackingPlanRequest `json:"tracking_plan,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
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
	if all.TrackingPlan != nil && all.TrackingPlan.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TrackingPlan = all.TrackingPlan

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
