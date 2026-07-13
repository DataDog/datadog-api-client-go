// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricNameFilterPreviewMetric A CloudWatch metric and the Datadog metric names it produces.
type AWSMetricNameFilterPreviewMetric struct {
	// The CloudWatch metric name.
	CwName string `json:"cw_name"`
	// The Datadog metric names produced from this CloudWatch metric.
	DdNames []AWSMetricNameFilterPreviewDDName `json:"dd_names"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSMetricNameFilterPreviewMetric instantiates a new AWSMetricNameFilterPreviewMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSMetricNameFilterPreviewMetric(cwName string, ddNames []AWSMetricNameFilterPreviewDDName) *AWSMetricNameFilterPreviewMetric {
	this := AWSMetricNameFilterPreviewMetric{}
	this.CwName = cwName
	this.DdNames = ddNames
	return &this
}

// NewAWSMetricNameFilterPreviewMetricWithDefaults instantiates a new AWSMetricNameFilterPreviewMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSMetricNameFilterPreviewMetricWithDefaults() *AWSMetricNameFilterPreviewMetric {
	this := AWSMetricNameFilterPreviewMetric{}
	return &this
}

// GetCwName returns the CwName field value.
func (o *AWSMetricNameFilterPreviewMetric) GetCwName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CwName
}

// GetCwNameOk returns a tuple with the CwName field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewMetric) GetCwNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CwName, true
}

// SetCwName sets field value.
func (o *AWSMetricNameFilterPreviewMetric) SetCwName(v string) {
	o.CwName = v
}

// GetDdNames returns the DdNames field value.
func (o *AWSMetricNameFilterPreviewMetric) GetDdNames() []AWSMetricNameFilterPreviewDDName {
	if o == nil {
		var ret []AWSMetricNameFilterPreviewDDName
		return ret
	}
	return o.DdNames
}

// GetDdNamesOk returns a tuple with the DdNames field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewMetric) GetDdNamesOk() (*[]AWSMetricNameFilterPreviewDDName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DdNames, true
}

// SetDdNames sets field value.
func (o *AWSMetricNameFilterPreviewMetric) SetDdNames(v []AWSMetricNameFilterPreviewDDName) {
	o.DdNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSMetricNameFilterPreviewMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cw_name"] = o.CwName
	toSerialize["dd_names"] = o.DdNames

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSMetricNameFilterPreviewMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CwName  *string                             `json:"cw_name"`
		DdNames *[]AWSMetricNameFilterPreviewDDName `json:"dd_names"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CwName == nil {
		return fmt.Errorf("required field cw_name missing")
	}
	if all.DdNames == nil {
		return fmt.Errorf("required field dd_names missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cw_name", "dd_names"})
	} else {
		return err
	}
	o.CwName = *all.CwName
	o.DdNames = *all.DdNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
