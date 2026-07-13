// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricNameFilterPreviewNamespace The metric name filter preview for a single namespace.
type AWSMetricNameFilterPreviewNamespace struct {
	// The metric name filter patterns evaluated for this namespace and how many metrics they matched.
	Filters []AWSMetricNameFilterPreviewFilterMatch `json:"filters"`
	// The CloudWatch metrics collected for this namespace and whether each resulting
	// Datadog metric is filtered.
	Metrics []AWSMetricNameFilterPreviewMetric `json:"metrics"`
	// The AWS CloudWatch namespace.
	Namespace string `json:"namespace"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSMetricNameFilterPreviewNamespace instantiates a new AWSMetricNameFilterPreviewNamespace object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSMetricNameFilterPreviewNamespace(filters []AWSMetricNameFilterPreviewFilterMatch, metrics []AWSMetricNameFilterPreviewMetric, namespace string) *AWSMetricNameFilterPreviewNamespace {
	this := AWSMetricNameFilterPreviewNamespace{}
	this.Filters = filters
	this.Metrics = metrics
	this.Namespace = namespace
	return &this
}

// NewAWSMetricNameFilterPreviewNamespaceWithDefaults instantiates a new AWSMetricNameFilterPreviewNamespace object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSMetricNameFilterPreviewNamespaceWithDefaults() *AWSMetricNameFilterPreviewNamespace {
	this := AWSMetricNameFilterPreviewNamespace{}
	return &this
}

// GetFilters returns the Filters field value.
func (o *AWSMetricNameFilterPreviewNamespace) GetFilters() []AWSMetricNameFilterPreviewFilterMatch {
	if o == nil {
		var ret []AWSMetricNameFilterPreviewFilterMatch
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewNamespace) GetFiltersOk() (*[]AWSMetricNameFilterPreviewFilterMatch, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *AWSMetricNameFilterPreviewNamespace) SetFilters(v []AWSMetricNameFilterPreviewFilterMatch) {
	o.Filters = v
}

// GetMetrics returns the Metrics field value.
func (o *AWSMetricNameFilterPreviewNamespace) GetMetrics() []AWSMetricNameFilterPreviewMetric {
	if o == nil {
		var ret []AWSMetricNameFilterPreviewMetric
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewNamespace) GetMetricsOk() (*[]AWSMetricNameFilterPreviewMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value.
func (o *AWSMetricNameFilterPreviewNamespace) SetMetrics(v []AWSMetricNameFilterPreviewMetric) {
	o.Metrics = v
}

// GetNamespace returns the Namespace field value.
func (o *AWSMetricNameFilterPreviewNamespace) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewNamespace) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *AWSMetricNameFilterPreviewNamespace) SetNamespace(v string) {
	o.Namespace = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSMetricNameFilterPreviewNamespace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filters"] = o.Filters
	toSerialize["metrics"] = o.Metrics
	toSerialize["namespace"] = o.Namespace

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSMetricNameFilterPreviewNamespace) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filters   *[]AWSMetricNameFilterPreviewFilterMatch `json:"filters"`
		Metrics   *[]AWSMetricNameFilterPreviewMetric      `json:"metrics"`
		Namespace *string                                  `json:"namespace"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	if all.Metrics == nil {
		return fmt.Errorf("required field metrics missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filters", "metrics", "namespace"})
	} else {
		return err
	}
	o.Filters = *all.Filters
	o.Metrics = *all.Metrics
	o.Namespace = *all.Namespace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
