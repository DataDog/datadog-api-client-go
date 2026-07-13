// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricNameFiltersIncludeOnly Include only metric names matching one of these patterns for a single namespace.
type AWSMetricNameFiltersIncludeOnly struct {
	// Include only metric names matching one of these patterns.
	IncludeOnly []string `json:"include_only"`
	// The AWS CloudWatch namespace to which this metric name filter applies.
	Namespace string `json:"namespace"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSMetricNameFiltersIncludeOnly instantiates a new AWSMetricNameFiltersIncludeOnly object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSMetricNameFiltersIncludeOnly(includeOnly []string, namespace string) *AWSMetricNameFiltersIncludeOnly {
	this := AWSMetricNameFiltersIncludeOnly{}
	this.IncludeOnly = includeOnly
	this.Namespace = namespace
	return &this
}

// NewAWSMetricNameFiltersIncludeOnlyWithDefaults instantiates a new AWSMetricNameFiltersIncludeOnly object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSMetricNameFiltersIncludeOnlyWithDefaults() *AWSMetricNameFiltersIncludeOnly {
	this := AWSMetricNameFiltersIncludeOnly{}
	return &this
}

// GetIncludeOnly returns the IncludeOnly field value.
func (o *AWSMetricNameFiltersIncludeOnly) GetIncludeOnly() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IncludeOnly
}

// GetIncludeOnlyOk returns a tuple with the IncludeOnly field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFiltersIncludeOnly) GetIncludeOnlyOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeOnly, true
}

// SetIncludeOnly sets field value.
func (o *AWSMetricNameFiltersIncludeOnly) SetIncludeOnly(v []string) {
	o.IncludeOnly = v
}

// GetNamespace returns the Namespace field value.
func (o *AWSMetricNameFiltersIncludeOnly) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFiltersIncludeOnly) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value.
func (o *AWSMetricNameFiltersIncludeOnly) SetNamespace(v string) {
	o.Namespace = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSMetricNameFiltersIncludeOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["include_only"] = o.IncludeOnly
	toSerialize["namespace"] = o.Namespace

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSMetricNameFiltersIncludeOnly) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IncludeOnly *[]string `json:"include_only"`
		Namespace   *string   `json:"namespace"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IncludeOnly == nil {
		return fmt.Errorf("required field include_only missing")
	}
	if all.Namespace == nil {
		return fmt.Errorf("required field namespace missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"include_only", "namespace"})
	} else {
		return err
	}
	o.IncludeOnly = *all.IncludeOnly
	o.Namespace = *all.Namespace

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
