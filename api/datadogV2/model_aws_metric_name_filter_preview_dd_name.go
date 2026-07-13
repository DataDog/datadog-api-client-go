// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSMetricNameFilterPreviewDDName A Datadog metric name and whether it is filtered.
type AWSMetricNameFilterPreviewDDName struct {
	// Whether this Datadog metric name is filtered out.
	Filtered bool `json:"filtered"`
	// The Datadog metric name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSMetricNameFilterPreviewDDName instantiates a new AWSMetricNameFilterPreviewDDName object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSMetricNameFilterPreviewDDName(filtered bool, name string) *AWSMetricNameFilterPreviewDDName {
	this := AWSMetricNameFilterPreviewDDName{}
	this.Filtered = filtered
	this.Name = name
	return &this
}

// NewAWSMetricNameFilterPreviewDDNameWithDefaults instantiates a new AWSMetricNameFilterPreviewDDName object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSMetricNameFilterPreviewDDNameWithDefaults() *AWSMetricNameFilterPreviewDDName {
	this := AWSMetricNameFilterPreviewDDName{}
	return &this
}

// GetFiltered returns the Filtered field value.
func (o *AWSMetricNameFilterPreviewDDName) GetFiltered() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Filtered
}

// GetFilteredOk returns a tuple with the Filtered field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewDDName) GetFilteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filtered, true
}

// SetFiltered sets field value.
func (o *AWSMetricNameFilterPreviewDDName) SetFiltered(v bool) {
	o.Filtered = v
}

// GetName returns the Name field value.
func (o *AWSMetricNameFilterPreviewDDName) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AWSMetricNameFilterPreviewDDName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AWSMetricNameFilterPreviewDDName) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSMetricNameFilterPreviewDDName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["filtered"] = o.Filtered
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSMetricNameFilterPreviewDDName) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filtered *bool   `json:"filtered"`
		Name     *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Filtered == nil {
		return fmt.Errorf("required field filtered missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"filtered", "name"})
	} else {
		return err
	}
	o.Filtered = *all.Filtered
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
