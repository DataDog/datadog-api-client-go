// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems Reference to a metric to include in the collection.
type ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems struct {
	// Identifier of the metric to include in the collection.
	MetricId uuid.UUID `json:"metric_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems instantiates a new ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems(metricId uuid.UUID) *ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems {
	this := ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems{}
	this.MetricId = metricId
	return &this
}

// NewExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItemsWithDefaults instantiates a new ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItemsWithDefaults() *ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems {
	this := ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems{}
	return &this
}

// GetMetricId returns the MetricId field value.
func (o *ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems) GetMetricId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems) GetMetricIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value.
func (o *ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems) SetMetricId(v uuid.UUID) {
	o.MetricId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["metric_id"] = o.MetricId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricId *uuid.UUID `json:"metric_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MetricId == nil {
		return fmt.Errorf("required field metric_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metric_id"})
	} else {
		return err
	}
	o.MetricId = *all.MetricId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
