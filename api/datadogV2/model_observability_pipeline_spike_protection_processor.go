// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSpikeProtectionProcessor The `spike_protection` processor learns the normal ingestion volume for the logs that match a specified filter and takes action when the volume spikes unexpectedly.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineSpikeProtectionProcessor struct {
	// How to handle logs that are missing one of the `partition_fields`. Options:
	// - `ignore`: Forward the event without applying spike protection.
	// - `default_bucket`: Assign the event to a default partition.
	ActionWhenMissingPartitions *ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions `json:"action_when_missing_partitions,omitempty"`
	// The display name for a component.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the processor is enabled.
	Enabled bool `json:"enabled"`
	// The unique identifier for this component. Used in other parts of the pipeline to reference this component (for example, as the `input` to downstream components).
	Id string `json:"id"`
	// A Datadog search query used to determine which logs this processor targets.
	Include string `json:"include"`
	// The minimum volume, in bytes, that must be observed within 24 hours before spike protection can activate. Defaults to 100 GiB.
	MinimalThreshold *int64 `json:"minimal_threshold,omitempty"`
	// A list of per-partition overrides for the minimal threshold, identified by matching field values. Requires `partition_fields` to be set.
	MinimalThresholdOverrides []ObservabilityPipelineSpikeProtectionProcessorThresholdOverride `json:"minimal_threshold_overrides,omitempty"`
	// Name of the spike protection rule.
	Name string `json:"name"`
	// A list of fields used to segment log traffic. Volume is learned and evaluated independently for each unique combination of these field values. A maximum of five fields is supported.
	PartitionFields []string `json:"partition_fields,omitempty"`
	// The action to take with matching logs while spike protection is active. Options:
	// - `drop`: Drop the event.
	// - `no_action`: Let the event pass through.
	// - `overflow_routing`: Route to an overflow destination.
	//
	// When `overflow_routing` is used, there must be a destination whose `inputs` reference this processor with the `<processor-id>.overflow_events` suffix to route the overflowing events.
	SpikeAction ObservabilityPipelineSpikeProtectionProcessorSpikeAction `json:"spike_action"`
	// The number of standard deviations above the learned mean volume that must be exceeded before a spike is detected. Defaults to `5`.
	SpikeThreshold *float64 `json:"spike_threshold,omitempty"`
	// The processor type. The value should always be `spike_protection`.
	Type ObservabilityPipelineSpikeProtectionProcessorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSpikeProtectionProcessor instantiates a new ObservabilityPipelineSpikeProtectionProcessor object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSpikeProtectionProcessor(enabled bool, id string, include string, name string, spikeAction ObservabilityPipelineSpikeProtectionProcessorSpikeAction, typeVar ObservabilityPipelineSpikeProtectionProcessorType) *ObservabilityPipelineSpikeProtectionProcessor {
	this := ObservabilityPipelineSpikeProtectionProcessor{}
	var actionWhenMissingPartitions ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions = OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORACTIONWHENMISSINGPARTITIONS_IGNORE
	this.ActionWhenMissingPartitions = &actionWhenMissingPartitions
	this.Enabled = enabled
	this.Id = id
	this.Include = include
	this.Name = name
	this.SpikeAction = spikeAction
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSpikeProtectionProcessorWithDefaults instantiates a new ObservabilityPipelineSpikeProtectionProcessor object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSpikeProtectionProcessorWithDefaults() *ObservabilityPipelineSpikeProtectionProcessor {
	this := ObservabilityPipelineSpikeProtectionProcessor{}
	var actionWhenMissingPartitions ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions = OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORACTIONWHENMISSINGPARTITIONS_IGNORE
	this.ActionWhenMissingPartitions = &actionWhenMissingPartitions
	var typeVar ObservabilityPipelineSpikeProtectionProcessorType = OBSERVABILITYPIPELINESPIKEPROTECTIONPROCESSORTYPE_SPIKE_PROTECTION
	this.Type = typeVar
	return &this
}

// GetActionWhenMissingPartitions returns the ActionWhenMissingPartitions field value if set, zero value otherwise.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetActionWhenMissingPartitions() ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions {
	if o == nil || o.ActionWhenMissingPartitions == nil {
		var ret ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions
		return ret
	}
	return *o.ActionWhenMissingPartitions
}

// GetActionWhenMissingPartitionsOk returns a tuple with the ActionWhenMissingPartitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetActionWhenMissingPartitionsOk() (*ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions, bool) {
	if o == nil || o.ActionWhenMissingPartitions == nil {
		return nil, false
	}
	return o.ActionWhenMissingPartitions, true
}

// HasActionWhenMissingPartitions returns a boolean if a field has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) HasActionWhenMissingPartitions() bool {
	return o != nil && o.ActionWhenMissingPartitions != nil
}

// SetActionWhenMissingPartitions gets a reference to the given ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions and assigns it to the ActionWhenMissingPartitions field.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetActionWhenMissingPartitions(v ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions) {
	o.ActionWhenMissingPartitions = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetId(v string) {
	o.Id = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetInclude(v string) {
	o.Include = v
}

// GetMinimalThreshold returns the MinimalThreshold field value if set, zero value otherwise.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetMinimalThreshold() int64 {
	if o == nil || o.MinimalThreshold == nil {
		var ret int64
		return ret
	}
	return *o.MinimalThreshold
}

// GetMinimalThresholdOk returns a tuple with the MinimalThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetMinimalThresholdOk() (*int64, bool) {
	if o == nil || o.MinimalThreshold == nil {
		return nil, false
	}
	return o.MinimalThreshold, true
}

// HasMinimalThreshold returns a boolean if a field has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) HasMinimalThreshold() bool {
	return o != nil && o.MinimalThreshold != nil
}

// SetMinimalThreshold gets a reference to the given int64 and assigns it to the MinimalThreshold field.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetMinimalThreshold(v int64) {
	o.MinimalThreshold = &v
}

// GetMinimalThresholdOverrides returns the MinimalThresholdOverrides field value if set, zero value otherwise.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetMinimalThresholdOverrides() []ObservabilityPipelineSpikeProtectionProcessorThresholdOverride {
	if o == nil || o.MinimalThresholdOverrides == nil {
		var ret []ObservabilityPipelineSpikeProtectionProcessorThresholdOverride
		return ret
	}
	return o.MinimalThresholdOverrides
}

// GetMinimalThresholdOverridesOk returns a tuple with the MinimalThresholdOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetMinimalThresholdOverridesOk() (*[]ObservabilityPipelineSpikeProtectionProcessorThresholdOverride, bool) {
	if o == nil || o.MinimalThresholdOverrides == nil {
		return nil, false
	}
	return &o.MinimalThresholdOverrides, true
}

// HasMinimalThresholdOverrides returns a boolean if a field has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) HasMinimalThresholdOverrides() bool {
	return o != nil && o.MinimalThresholdOverrides != nil
}

// SetMinimalThresholdOverrides gets a reference to the given []ObservabilityPipelineSpikeProtectionProcessorThresholdOverride and assigns it to the MinimalThresholdOverrides field.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetMinimalThresholdOverrides(v []ObservabilityPipelineSpikeProtectionProcessorThresholdOverride) {
	o.MinimalThresholdOverrides = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetName(v string) {
	o.Name = v
}

// GetPartitionFields returns the PartitionFields field value if set, zero value otherwise.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetPartitionFields() []string {
	if o == nil || o.PartitionFields == nil {
		var ret []string
		return ret
	}
	return o.PartitionFields
}

// GetPartitionFieldsOk returns a tuple with the PartitionFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetPartitionFieldsOk() (*[]string, bool) {
	if o == nil || o.PartitionFields == nil {
		return nil, false
	}
	return &o.PartitionFields, true
}

// HasPartitionFields returns a boolean if a field has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) HasPartitionFields() bool {
	return o != nil && o.PartitionFields != nil
}

// SetPartitionFields gets a reference to the given []string and assigns it to the PartitionFields field.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetPartitionFields(v []string) {
	o.PartitionFields = v
}

// GetSpikeAction returns the SpikeAction field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetSpikeAction() ObservabilityPipelineSpikeProtectionProcessorSpikeAction {
	if o == nil {
		var ret ObservabilityPipelineSpikeProtectionProcessorSpikeAction
		return ret
	}
	return o.SpikeAction
}

// GetSpikeActionOk returns a tuple with the SpikeAction field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetSpikeActionOk() (*ObservabilityPipelineSpikeProtectionProcessorSpikeAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpikeAction, true
}

// SetSpikeAction sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetSpikeAction(v ObservabilityPipelineSpikeProtectionProcessorSpikeAction) {
	o.SpikeAction = v
}

// GetSpikeThreshold returns the SpikeThreshold field value if set, zero value otherwise.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetSpikeThreshold() float64 {
	if o == nil || o.SpikeThreshold == nil {
		var ret float64
		return ret
	}
	return *o.SpikeThreshold
}

// GetSpikeThresholdOk returns a tuple with the SpikeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetSpikeThresholdOk() (*float64, bool) {
	if o == nil || o.SpikeThreshold == nil {
		return nil, false
	}
	return o.SpikeThreshold, true
}

// HasSpikeThreshold returns a boolean if a field has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) HasSpikeThreshold() bool {
	return o != nil && o.SpikeThreshold != nil
}

// SetSpikeThreshold gets a reference to the given float64 and assigns it to the SpikeThreshold field.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetSpikeThreshold(v float64) {
	o.SpikeThreshold = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetType() ObservabilityPipelineSpikeProtectionProcessorType {
	if o == nil {
		var ret ObservabilityPipelineSpikeProtectionProcessorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSpikeProtectionProcessor) GetTypeOk() (*ObservabilityPipelineSpikeProtectionProcessorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSpikeProtectionProcessor) SetType(v ObservabilityPipelineSpikeProtectionProcessorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSpikeProtectionProcessor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ActionWhenMissingPartitions != nil {
		toSerialize["action_when_missing_partitions"] = o.ActionWhenMissingPartitions
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	toSerialize["include"] = o.Include
	if o.MinimalThreshold != nil {
		toSerialize["minimal_threshold"] = o.MinimalThreshold
	}
	if o.MinimalThresholdOverrides != nil {
		toSerialize["minimal_threshold_overrides"] = o.MinimalThresholdOverrides
	}
	toSerialize["name"] = o.Name
	if o.PartitionFields != nil {
		toSerialize["partition_fields"] = o.PartitionFields
	}
	toSerialize["spike_action"] = o.SpikeAction
	if o.SpikeThreshold != nil {
		toSerialize["spike_threshold"] = o.SpikeThreshold
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSpikeProtectionProcessor) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionWhenMissingPartitions *ObservabilityPipelineSpikeProtectionProcessorActionWhenMissingPartitions `json:"action_when_missing_partitions,omitempty"`
		DisplayName                 *string                                                                   `json:"display_name,omitempty"`
		Enabled                     *bool                                                                     `json:"enabled"`
		Id                          *string                                                                   `json:"id"`
		Include                     *string                                                                   `json:"include"`
		MinimalThreshold            *int64                                                                    `json:"minimal_threshold,omitempty"`
		MinimalThresholdOverrides   []ObservabilityPipelineSpikeProtectionProcessorThresholdOverride          `json:"minimal_threshold_overrides,omitempty"`
		Name                        *string                                                                   `json:"name"`
		PartitionFields             []string                                                                  `json:"partition_fields,omitempty"`
		SpikeAction                 *ObservabilityPipelineSpikeProtectionProcessorSpikeAction                 `json:"spike_action"`
		SpikeThreshold              *float64                                                                  `json:"spike_threshold,omitempty"`
		Type                        *ObservabilityPipelineSpikeProtectionProcessorType                        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.SpikeAction == nil {
		return fmt.Errorf("required field spike_action missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action_when_missing_partitions", "display_name", "enabled", "id", "include", "minimal_threshold", "minimal_threshold_overrides", "name", "partition_fields", "spike_action", "spike_threshold", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ActionWhenMissingPartitions != nil && !all.ActionWhenMissingPartitions.IsValid() {
		hasInvalidField = true
	} else {
		o.ActionWhenMissingPartitions = all.ActionWhenMissingPartitions
	}
	o.DisplayName = all.DisplayName
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	o.Include = *all.Include
	o.MinimalThreshold = all.MinimalThreshold
	o.MinimalThresholdOverrides = all.MinimalThresholdOverrides
	o.Name = *all.Name
	o.PartitionFields = all.PartitionFields
	if !all.SpikeAction.IsValid() {
		hasInvalidField = true
	} else {
		o.SpikeAction = *all.SpikeAction
	}
	o.SpikeThreshold = all.SpikeThreshold
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
