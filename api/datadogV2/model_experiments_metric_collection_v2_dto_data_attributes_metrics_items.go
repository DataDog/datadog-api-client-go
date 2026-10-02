// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsMetricCollectionV2DTODataAttributesMetricsItems A metric included in the collection.
type ExperimentsMetricCollectionV2DTODataAttributesMetricsItems struct {
	// ID of the metric represented by this entry.
	MetricId *string `json:"metric_id,omitempty"`
	// Display name of the metric represented by this entry.
	MetricName *string `json:"metric_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsMetricCollectionV2DTODataAttributesMetricsItems instantiates a new ExperimentsMetricCollectionV2DTODataAttributesMetricsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsMetricCollectionV2DTODataAttributesMetricsItems() *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems {
	this := ExperimentsMetricCollectionV2DTODataAttributesMetricsItems{}
	return &this
}

// NewExperimentsMetricCollectionV2DTODataAttributesMetricsItemsWithDefaults instantiates a new ExperimentsMetricCollectionV2DTODataAttributesMetricsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsMetricCollectionV2DTODataAttributesMetricsItemsWithDefaults() *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems {
	this := ExperimentsMetricCollectionV2DTODataAttributesMetricsItems{}
	return &this
}

// GetMetricId returns the MetricId field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) GetMetricId() string {
	if o == nil || o.MetricId == nil {
		var ret string
		return ret
	}
	return *o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) GetMetricIdOk() (*string, bool) {
	if o == nil || o.MetricId == nil {
		return nil, false
	}
	return o.MetricId, true
}

// HasMetricId returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) HasMetricId() bool {
	return o != nil && o.MetricId != nil
}

// SetMetricId gets a reference to the given string and assigns it to the MetricId field.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) SetMetricId(v string) {
	o.MetricId = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) GetMetricName() string {
	if o == nil || o.MetricName == nil {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) GetMetricNameOk() (*string, bool) {
	if o == nil || o.MetricName == nil {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) HasMetricName() bool {
	return o != nil && o.MetricName != nil
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) SetMetricName(v string) {
	o.MetricName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MetricId != nil {
		toSerialize["metric_id"] = o.MetricId
	}
	if o.MetricName != nil {
		toSerialize["metric_name"] = o.MetricName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsMetricCollectionV2DTODataAttributesMetricsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricId   *string `json:"metric_id,omitempty"`
		MetricName *string `json:"metric_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metric_id", "metric_name"})
	} else {
		return err
	}
	o.MetricId = all.MetricId
	o.MetricName = all.MetricName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
