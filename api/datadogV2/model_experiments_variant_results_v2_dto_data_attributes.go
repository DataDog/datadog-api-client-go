// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsVariantResultsV2DTODataAttributes Details of the variant result.
type ExperimentsVariantResultsV2DTODataAttributes struct {
	// Number of subjects assigned to this variant.
	AssignmentCount *int64 `json:"assignment_count,omitempty"`
	// ID of the experiment associated with this result.
	ExperimentId *string `json:"experiment_id,omitempty"`
	// Whether this variant is the experiment's control.
	IsControl *bool `json:"is_control,omitempty"`
	// Metrics reported for this variant.
	Metrics []ExperimentsVariantResultsV2DTODataAttributesMetricsItems `json:"metrics,omitempty"`
	// Key that identifies the experiment variant.
	VariantKey *string `json:"variant_key,omitempty"`
	// Display name of the experiment variant.
	VariantName *string `json:"variant_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsVariantResultsV2DTODataAttributes instantiates a new ExperimentsVariantResultsV2DTODataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsVariantResultsV2DTODataAttributes() *ExperimentsVariantResultsV2DTODataAttributes {
	this := ExperimentsVariantResultsV2DTODataAttributes{}
	return &this
}

// NewExperimentsVariantResultsV2DTODataAttributesWithDefaults instantiates a new ExperimentsVariantResultsV2DTODataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsVariantResultsV2DTODataAttributesWithDefaults() *ExperimentsVariantResultsV2DTODataAttributes {
	this := ExperimentsVariantResultsV2DTODataAttributes{}
	return &this
}

// GetAssignmentCount returns the AssignmentCount field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetAssignmentCount() int64 {
	if o == nil || o.AssignmentCount == nil {
		var ret int64
		return ret
	}
	return *o.AssignmentCount
}

// GetAssignmentCountOk returns a tuple with the AssignmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetAssignmentCountOk() (*int64, bool) {
	if o == nil || o.AssignmentCount == nil {
		return nil, false
	}
	return o.AssignmentCount, true
}

// HasAssignmentCount returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) HasAssignmentCount() bool {
	return o != nil && o.AssignmentCount != nil
}

// SetAssignmentCount gets a reference to the given int64 and assigns it to the AssignmentCount field.
func (o *ExperimentsVariantResultsV2DTODataAttributes) SetAssignmentCount(v int64) {
	o.AssignmentCount = &v
}

// GetExperimentId returns the ExperimentId field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetExperimentId() string {
	if o == nil || o.ExperimentId == nil {
		var ret string
		return ret
	}
	return *o.ExperimentId
}

// GetExperimentIdOk returns a tuple with the ExperimentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetExperimentIdOk() (*string, bool) {
	if o == nil || o.ExperimentId == nil {
		return nil, false
	}
	return o.ExperimentId, true
}

// HasExperimentId returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) HasExperimentId() bool {
	return o != nil && o.ExperimentId != nil
}

// SetExperimentId gets a reference to the given string and assigns it to the ExperimentId field.
func (o *ExperimentsVariantResultsV2DTODataAttributes) SetExperimentId(v string) {
	o.ExperimentId = &v
}

// GetIsControl returns the IsControl field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetIsControl() bool {
	if o == nil || o.IsControl == nil {
		var ret bool
		return ret
	}
	return *o.IsControl
}

// GetIsControlOk returns a tuple with the IsControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetIsControlOk() (*bool, bool) {
	if o == nil || o.IsControl == nil {
		return nil, false
	}
	return o.IsControl, true
}

// HasIsControl returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) HasIsControl() bool {
	return o != nil && o.IsControl != nil
}

// SetIsControl gets a reference to the given bool and assigns it to the IsControl field.
func (o *ExperimentsVariantResultsV2DTODataAttributes) SetIsControl(v bool) {
	o.IsControl = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetMetrics() []ExperimentsVariantResultsV2DTODataAttributesMetricsItems {
	if o == nil || o.Metrics == nil {
		var ret []ExperimentsVariantResultsV2DTODataAttributesMetricsItems
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetMetricsOk() (*[]ExperimentsVariantResultsV2DTODataAttributesMetricsItems, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []ExperimentsVariantResultsV2DTODataAttributesMetricsItems and assigns it to the Metrics field.
func (o *ExperimentsVariantResultsV2DTODataAttributes) SetMetrics(v []ExperimentsVariantResultsV2DTODataAttributesMetricsItems) {
	o.Metrics = v
}

// GetVariantKey returns the VariantKey field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetVariantKey() string {
	if o == nil || o.VariantKey == nil {
		var ret string
		return ret
	}
	return *o.VariantKey
}

// GetVariantKeyOk returns a tuple with the VariantKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetVariantKeyOk() (*string, bool) {
	if o == nil || o.VariantKey == nil {
		return nil, false
	}
	return o.VariantKey, true
}

// HasVariantKey returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) HasVariantKey() bool {
	return o != nil && o.VariantKey != nil
}

// SetVariantKey gets a reference to the given string and assigns it to the VariantKey field.
func (o *ExperimentsVariantResultsV2DTODataAttributes) SetVariantKey(v string) {
	o.VariantKey = &v
}

// GetVariantName returns the VariantName field value if set, zero value otherwise.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetVariantName() string {
	if o == nil || o.VariantName == nil {
		var ret string
		return ret
	}
	return *o.VariantName
}

// GetVariantNameOk returns a tuple with the VariantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) GetVariantNameOk() (*string, bool) {
	if o == nil || o.VariantName == nil {
		return nil, false
	}
	return o.VariantName, true
}

// HasVariantName returns a boolean if a field has been set.
func (o *ExperimentsVariantResultsV2DTODataAttributes) HasVariantName() bool {
	return o != nil && o.VariantName != nil
}

// SetVariantName gets a reference to the given string and assigns it to the VariantName field.
func (o *ExperimentsVariantResultsV2DTODataAttributes) SetVariantName(v string) {
	o.VariantName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsVariantResultsV2DTODataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignmentCount != nil {
		toSerialize["assignment_count"] = o.AssignmentCount
	}
	if o.ExperimentId != nil {
		toSerialize["experiment_id"] = o.ExperimentId
	}
	if o.IsControl != nil {
		toSerialize["is_control"] = o.IsControl
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.VariantKey != nil {
		toSerialize["variant_key"] = o.VariantKey
	}
	if o.VariantName != nil {
		toSerialize["variant_name"] = o.VariantName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsVariantResultsV2DTODataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentCount *int64                                                     `json:"assignment_count,omitempty"`
		ExperimentId    *string                                                    `json:"experiment_id,omitempty"`
		IsControl       *bool                                                      `json:"is_control,omitempty"`
		Metrics         []ExperimentsVariantResultsV2DTODataAttributesMetricsItems `json:"metrics,omitempty"`
		VariantKey      *string                                                    `json:"variant_key,omitempty"`
		VariantName     *string                                                    `json:"variant_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_count", "experiment_id", "is_control", "metrics", "variant_key", "variant_name"})
	} else {
		return err
	}
	o.AssignmentCount = all.AssignmentCount
	o.ExperimentId = all.ExperimentId
	o.IsControl = all.IsControl
	o.Metrics = all.Metrics
	o.VariantKey = all.VariantKey
	o.VariantName = all.VariantName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
