// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsProtocolMetricGroupInput A metric group supplied by the protocol. Empty decision groups are removed.
type ExperimentsProtocolMetricGroupInput struct {
	// Whether this group contains decision metrics.
	IsDecision *bool `json:"is_decision,omitempty"`
	// Metrics in the group.
	Metrics []ExperimentsProtocolMetricInput `json:"metrics,omitempty"`
	// Metric group name.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsProtocolMetricGroupInput instantiates a new ExperimentsProtocolMetricGroupInput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsProtocolMetricGroupInput() *ExperimentsProtocolMetricGroupInput {
	this := ExperimentsProtocolMetricGroupInput{}
	return &this
}

// NewExperimentsProtocolMetricGroupInputWithDefaults instantiates a new ExperimentsProtocolMetricGroupInput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsProtocolMetricGroupInputWithDefaults() *ExperimentsProtocolMetricGroupInput {
	this := ExperimentsProtocolMetricGroupInput{}
	return &this
}

// GetIsDecision returns the IsDecision field value if set, zero value otherwise.
func (o *ExperimentsProtocolMetricGroupInput) GetIsDecision() bool {
	if o == nil || o.IsDecision == nil {
		var ret bool
		return ret
	}
	return *o.IsDecision
}

// GetIsDecisionOk returns a tuple with the IsDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolMetricGroupInput) GetIsDecisionOk() (*bool, bool) {
	if o == nil || o.IsDecision == nil {
		return nil, false
	}
	return o.IsDecision, true
}

// HasIsDecision returns a boolean if a field has been set.
func (o *ExperimentsProtocolMetricGroupInput) HasIsDecision() bool {
	return o != nil && o.IsDecision != nil
}

// SetIsDecision gets a reference to the given bool and assigns it to the IsDecision field.
func (o *ExperimentsProtocolMetricGroupInput) SetIsDecision(v bool) {
	o.IsDecision = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsProtocolMetricGroupInput) GetMetrics() []ExperimentsProtocolMetricInput {
	if o == nil {
		var ret []ExperimentsProtocolMetricInput
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsProtocolMetricGroupInput) GetMetricsOk() (*[]ExperimentsProtocolMetricInput, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsProtocolMetricGroupInput) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsProtocolMetricInput and assigns it to the Metrics field.
func (o *ExperimentsProtocolMetricGroupInput) SetMetrics(v []ExperimentsProtocolMetricInput) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsProtocolMetricGroupInput) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsProtocolMetricGroupInput) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsProtocolMetricGroupInput) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsProtocolMetricGroupInput) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsProtocolMetricGroupInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IsDecision != nil {
		toSerialize["is_decision"] = o.IsDecision
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsProtocolMetricGroupInput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsDecision *bool                            `json:"is_decision,omitempty"`
		Metrics    []ExperimentsProtocolMetricInput `json:"metrics,omitempty"`
		Name       *string                          `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_decision", "metrics", "name"})
	} else {
		return err
	}
	o.IsDecision = all.IsDecision
	o.Metrics = all.Metrics
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
