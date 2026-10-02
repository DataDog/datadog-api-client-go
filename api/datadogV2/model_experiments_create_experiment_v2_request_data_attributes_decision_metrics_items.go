// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems Metric used to make an experiment decision, with its primary metric designation.
type ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems struct {
	// Whether this is the experiment primary metric.
	IsPrimary bool `json:"is_primary"`
	// Decision metric UUID.
	MetricId uuid.UUID `json:"metric_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems(isPrimary bool, metricId uuid.UUID) *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems{}
	this.IsPrimary = isPrimary
	this.MetricId = metricId
	return &this
}

// NewExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItemsWithDefaults instantiates a new ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItemsWithDefaults() *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems {
	this := ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems{}
	return &this
}

// GetIsPrimary returns the IsPrimary field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) GetIsPrimary() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) GetIsPrimaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPrimary, true
}

// SetIsPrimary sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) SetIsPrimary(v bool) {
	o.IsPrimary = v
}

// GetMetricId returns the MetricId field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) GetMetricId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) GetMetricIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) SetMetricId(v uuid.UUID) {
	o.MetricId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_primary"] = o.IsPrimary
	toSerialize["metric_id"] = o.MetricId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateExperimentV2RequestDataAttributesDecisionMetricsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsPrimary *bool      `json:"is_primary"`
		MetricId  *uuid.UUID `json:"metric_id"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_primary", "metric_id"})
	} else {
		return err
	}
	o.IsPrimary = *all.IsPrimary
	o.MetricId = *all.MetricId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
