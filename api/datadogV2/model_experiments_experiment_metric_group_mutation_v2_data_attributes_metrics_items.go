// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems Metric in an experiment metric group, with its name and primary metric designation.
type ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems struct {
	// Whether this is the experiment primary metric.
	IsPrimary bool `json:"is_primary"`
	// Identifier of the metric in the group.
	MetricId string `json:"metric_id"`
	// Display name of the metric in the group.
	MetricName string `json:"metric_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems instantiates a new ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems(isPrimary bool, metricId string, metricName string) *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems {
	this := ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems{}
	this.IsPrimary = isPrimary
	this.MetricId = metricId
	this.MetricName = metricName
	return &this
}

// NewExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItemsWithDefaults instantiates a new ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItemsWithDefaults() *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems {
	this := ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems{}
	return &this
}

// GetIsPrimary returns the IsPrimary field value.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetMetricId returns the MetricId field value.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) GetMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) SetMetricId(v string) {
	o.MetricId = v
}

// GetMetricName returns the MetricName field value.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) SetMetricName(v string) {
	o.MetricName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_primary"] = o.IsPrimary
	toSerialize["metric_id"] = o.MetricId
	toSerialize["metric_name"] = o.MetricName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentMetricGroupMutationV2DataAttributesMetricsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsPrimary  *bool   `json:"is_primary"`
		MetricId   *string `json:"metric_id"`
		MetricName *string `json:"metric_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsPrimary == nil {
		return fmt.Errorf("required field is_primary missing")
	}
	if all.MetricId == nil {
		return fmt.Errorf("required field metric_id missing")
	}
	if all.MetricName == nil {
		return fmt.Errorf("required field metric_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_primary", "metric_id", "metric_name"})
	} else {
		return err
	}
	o.IsPrimary = *all.IsPrimary
	o.MetricId = *all.MetricId
	o.MetricName = *all.MetricName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
