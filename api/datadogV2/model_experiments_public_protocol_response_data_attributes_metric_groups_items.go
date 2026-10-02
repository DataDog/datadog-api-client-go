// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems A group of metrics supplied by the protocol.
type ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems struct {
	// Whether the group contains decision metrics.
	IsDecision *bool `json:"is_decision,omitempty"`
	// Metrics included in this group.
	Metrics []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
	// Display name of the metric group.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems instantiates a new ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems() *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems {
	this := ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems{}
	return &this
}

// NewExperimentsPublicProtocolResponseDataAttributesMetricGroupsItemsWithDefaults instantiates a new ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPublicProtocolResponseDataAttributesMetricGroupsItemsWithDefaults() *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems {
	this := ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems{}
	return &this
}

// GetIsDecision returns the IsDecision field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) GetIsDecision() bool {
	if o == nil || o.IsDecision == nil {
		var ret bool
		return ret
	}
	return *o.IsDecision
}

// GetIsDecisionOk returns a tuple with the IsDecision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) GetIsDecisionOk() (*bool, bool) {
	if o == nil || o.IsDecision == nil {
		return nil, false
	}
	return o.IsDecision, true
}

// HasIsDecision returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) HasIsDecision() bool {
	return o != nil && o.IsDecision != nil
}

// SetIsDecision gets a reference to the given bool and assigns it to the IsDecision field.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) SetIsDecision(v bool) {
	o.IsDecision = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) GetMetrics() []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) GetMetricsOk() (*[]ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems and assigns it to the Metrics field.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) SetMetrics(v []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) MarshalJSON() ([]byte, error) {
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
func (o *ExperimentsPublicProtocolResponseDataAttributesMetricGroupsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsDecision *bool                                                                  `json:"is_decision,omitempty"`
		Metrics    []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
		Name       *string                                                                `json:"name,omitempty"`
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
