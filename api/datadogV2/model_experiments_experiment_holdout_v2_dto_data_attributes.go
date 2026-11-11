// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsExperimentHoldoutV2DTODataAttributes Attributes of a holdout that the experiment is attached to.
type ExperimentsExperimentHoldoutV2DTODataAttributes struct {
	// Zero-based position of this MANAGED holdout in the experiment's evaluation order. The GLOBAL holdout is evaluated first. Null for EXTERNAL holdouts.
	EvaluationOrder datadog.NullableInt64 `json:"evaluation_order,omitempty"`
	// How Datadog serves a MANAGED holdout. GLOBAL holdouts apply to every eligible experiment of a subject type. SELECTIVE holdouts apply only to the experiments attached to them. Null for EXTERNAL holdouts.
	HoldoutType NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType `json:"holdout_type,omitempty"`
	// Key of the holdout.
	Key *string `json:"key,omitempty"`
	// Whether Datadog serves the holdout. Datadog serves MANAGED holdouts. EXTERNAL holdouts record what your own feature flagging system holds back, and Datadog does not serve them.
	Kind *ExperimentsExperimentHoldoutV2DTODataAttributesKind `json:"kind,omitempty"`
	// Display name of the holdout.
	Name *string `json:"name,omitempty"`
	// Lifecycle phase of the holdout. PENDING holdouts have not started including experiments. INCLUDING holdouts are adding experiments. WAITING_FOR_ANALYSIS holdouts have stopped including experiments but have not started analysis. ANALYSING holdouts are in their analysis window. COMPLETE applies only to EXTERNAL holdouts that have ended their analysis window but are not archived.
	State *ExperimentsExperimentHoldoutV2DTODataAttributesState `json:"state,omitempty"`
	// Share of subjects the holdout holds back, from 0 to 1.
	TrafficAllocation *float64 `json:"traffic_allocation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsExperimentHoldoutV2DTODataAttributes instantiates a new ExperimentsExperimentHoldoutV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsExperimentHoldoutV2DTODataAttributes() *ExperimentsExperimentHoldoutV2DTODataAttributes {
	this := ExperimentsExperimentHoldoutV2DTODataAttributes{}
	return &this
}

// NewExperimentsExperimentHoldoutV2DTODataAttributesWithDefaults instantiates a new ExperimentsExperimentHoldoutV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsExperimentHoldoutV2DTODataAttributesWithDefaults() *ExperimentsExperimentHoldoutV2DTODataAttributes {
	this := ExperimentsExperimentHoldoutV2DTODataAttributes{}
	return &this
}

// GetEvaluationOrder returns the EvaluationOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetEvaluationOrder() int64 {
	if o == nil || o.EvaluationOrder.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EvaluationOrder.Get()
}

// GetEvaluationOrderOk returns a tuple with the EvaluationOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetEvaluationOrderOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvaluationOrder.Get(), o.EvaluationOrder.IsSet()
}

// HasEvaluationOrder returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) HasEvaluationOrder() bool {
	return o != nil && o.EvaluationOrder.IsSet()
}

// SetEvaluationOrder gets a reference to the given datadog.NullableInt64 and assigns it to the EvaluationOrder field.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetEvaluationOrder(v int64) {
	o.EvaluationOrder.Set(&v)
}

// SetEvaluationOrderNil sets the value for EvaluationOrder to be an explicit nil.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetEvaluationOrderNil() {
	o.EvaluationOrder.Set(nil)
}

// UnsetEvaluationOrder ensures that no value is present for EvaluationOrder, not even an explicit nil.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) UnsetEvaluationOrder() {
	o.EvaluationOrder.Unset()
}

// GetHoldoutType returns the HoldoutType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetHoldoutType() ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType {
	if o == nil || o.HoldoutType.Get() == nil {
		var ret ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType
		return ret
	}
	return *o.HoldoutType.Get()
}

// GetHoldoutTypeOk returns a tuple with the HoldoutType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetHoldoutTypeOk() (*ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType, bool) {
	if o == nil {
		return nil, false
	}
	return o.HoldoutType.Get(), o.HoldoutType.IsSet()
}

// HasHoldoutType returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) HasHoldoutType() bool {
	return o != nil && o.HoldoutType.IsSet()
}

// SetHoldoutType gets a reference to the given NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType and assigns it to the HoldoutType field.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetHoldoutType(v ExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType) {
	o.HoldoutType.Set(&v)
}

// SetHoldoutTypeNil sets the value for HoldoutType to be an explicit nil.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetHoldoutTypeNil() {
	o.HoldoutType.Set(nil)
}

// UnsetHoldoutType ensures that no value is present for HoldoutType, not even an explicit nil.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) UnsetHoldoutType() {
	o.HoldoutType.Unset()
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetKey(v string) {
	o.Key = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetKind() ExperimentsExperimentHoldoutV2DTODataAttributesKind {
	if o == nil || o.Kind == nil {
		var ret ExperimentsExperimentHoldoutV2DTODataAttributesKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetKindOk() (*ExperimentsExperimentHoldoutV2DTODataAttributesKind, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given ExperimentsExperimentHoldoutV2DTODataAttributesKind and assigns it to the Kind field.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetKind(v ExperimentsExperimentHoldoutV2DTODataAttributesKind) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetState() ExperimentsExperimentHoldoutV2DTODataAttributesState {
	if o == nil || o.State == nil {
		var ret ExperimentsExperimentHoldoutV2DTODataAttributesState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetStateOk() (*ExperimentsExperimentHoldoutV2DTODataAttributesState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given ExperimentsExperimentHoldoutV2DTODataAttributesState and assigns it to the State field.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetState(v ExperimentsExperimentHoldoutV2DTODataAttributesState) {
	o.State = &v
}

// GetTrafficAllocation returns the TrafficAllocation field value if set, zero value otherwise.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetTrafficAllocation() float64 {
	if o == nil || o.TrafficAllocation == nil {
		var ret float64
		return ret
	}
	return *o.TrafficAllocation
}

// GetTrafficAllocationOk returns a tuple with the TrafficAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) GetTrafficAllocationOk() (*float64, bool) {
	if o == nil || o.TrafficAllocation == nil {
		return nil, false
	}
	return o.TrafficAllocation, true
}

// HasTrafficAllocation returns a boolean if a field has been set.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) HasTrafficAllocation() bool {
	return o != nil && o.TrafficAllocation != nil
}

// SetTrafficAllocation gets a reference to the given float64 and assigns it to the TrafficAllocation field.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) SetTrafficAllocation(v float64) {
	o.TrafficAllocation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsExperimentHoldoutV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EvaluationOrder.IsSet() {
		toSerialize["evaluation_order"] = o.EvaluationOrder.Get()
	}
	if o.HoldoutType.IsSet() {
		toSerialize["holdout_type"] = o.HoldoutType.Get()
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.TrafficAllocation != nil {
		toSerialize["traffic_allocation"] = o.TrafficAllocation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsExperimentHoldoutV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EvaluationOrder   datadog.NullableInt64                                              `json:"evaluation_order,omitempty"`
		HoldoutType       NullableExperimentsExperimentHoldoutV2DTODataAttributesHoldoutType `json:"holdout_type,omitempty"`
		Key               *string                                                            `json:"key,omitempty"`
		Kind              *ExperimentsExperimentHoldoutV2DTODataAttributesKind               `json:"kind,omitempty"`
		Name              *string                                                            `json:"name,omitempty"`
		State             *ExperimentsExperimentHoldoutV2DTODataAttributesState              `json:"state,omitempty"`
		TrafficAllocation *float64                                                           `json:"traffic_allocation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"evaluation_order", "holdout_type", "key", "kind", "name", "state", "traffic_allocation"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EvaluationOrder = all.EvaluationOrder
	if all.HoldoutType.Get() != nil && !all.HoldoutType.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.HoldoutType = all.HoldoutType
	}
	o.Key = all.Key
	if all.Kind != nil && !all.Kind.IsValid() {
		hasInvalidField = true
	} else {
		o.Kind = all.Kind
	}
	o.Name = all.Name
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}
	o.TrafficAllocation = all.TrafficAllocation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
