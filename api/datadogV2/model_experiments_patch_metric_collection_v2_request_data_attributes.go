// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsPatchMetricCollectionV2RequestDataAttributes Fields supplied to update the metric collection.
type ExperimentsPatchMetricCollectionV2RequestDataAttributes struct {
	// Text that explains the metric collection.
	Description datadog.NullableString `json:"description,omitempty"`
	// Whether the collection is used for guardrail metrics.
	IsGuardrail *bool `json:"is_guardrail,omitempty"`
	// Metrics included in this collection.
	Metrics []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
	// Display name of the metric collection.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsPatchMetricCollectionV2RequestDataAttributes instantiates a new ExperimentsPatchMetricCollectionV2RequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsPatchMetricCollectionV2RequestDataAttributes() *ExperimentsPatchMetricCollectionV2RequestDataAttributes {
	this := ExperimentsPatchMetricCollectionV2RequestDataAttributes{}
	return &this
}

// NewExperimentsPatchMetricCollectionV2RequestDataAttributesWithDefaults instantiates a new ExperimentsPatchMetricCollectionV2RequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsPatchMetricCollectionV2RequestDataAttributesWithDefaults() *ExperimentsPatchMetricCollectionV2RequestDataAttributes {
	this := ExperimentsPatchMetricCollectionV2RequestDataAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description.IsSet()
}

// SetDescription gets a reference to the given datadog.NullableString and assigns it to the Description field.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) SetDescription(v string) {
	o.Description.Set(&v)
}

// SetDescriptionNil sets the value for Description to be an explicit nil.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) UnsetDescription() {
	o.Description.Unset()
}

// GetIsGuardrail returns the IsGuardrail field value if set, zero value otherwise.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetIsGuardrail() bool {
	if o == nil || o.IsGuardrail == nil {
		var ret bool
		return ret
	}
	return *o.IsGuardrail
}

// GetIsGuardrailOk returns a tuple with the IsGuardrail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetIsGuardrailOk() (*bool, bool) {
	if o == nil || o.IsGuardrail == nil {
		return nil, false
	}
	return o.IsGuardrail, true
}

// HasIsGuardrail returns a boolean if a field has been set.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) HasIsGuardrail() bool {
	return o != nil && o.IsGuardrail != nil
}

// SetIsGuardrail gets a reference to the given bool and assigns it to the IsGuardrail field.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) SetIsGuardrail(v bool) {
	o.IsGuardrail = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetMetrics() []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetMetricsOk() (*[]ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems and assigns it to the Metrics field.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) SetMetrics(v []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsPatchMetricCollectionV2RequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.IsGuardrail != nil {
		toSerialize["is_guardrail"] = o.IsGuardrail
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
func (o *ExperimentsPatchMetricCollectionV2RequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description datadog.NullableString                                                 `json:"description,omitempty"`
		IsGuardrail *bool                                                                  `json:"is_guardrail,omitempty"`
		Metrics     []ExperimentsCreateMetricCollectionV2RequestDataAttributesMetricsItems `json:"metrics,omitempty"`
		Name        *string                                                                `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "is_guardrail", "metrics", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.IsGuardrail = all.IsGuardrail
	o.Metrics = all.Metrics
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
