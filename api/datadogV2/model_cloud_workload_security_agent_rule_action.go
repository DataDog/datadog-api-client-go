// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleAction The action the rule can perform if triggered
type CloudWorkloadSecurityAgentRuleAction struct {
	// The core dump action applied on the process matching the rule.
	Coredump *CloudWorkloadSecurityAgentRuleActionCoreDump `json:"coredump,omitempty"`
	// Whether the action is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// SECL expression used to target the container to apply the action on
	Filter *string `json:"filter,omitempty"`
	// Hash file specified by the field attribute
	Hash *CloudWorkloadSecurityAgentRuleActionHash `json:"hash,omitempty"`
	// Kill system call applied on the container matching the rule
	Kill *CloudWorkloadSecurityAgentRuleKill `json:"kill,omitempty"`
	// The log action applied when the rule is triggered.
	Log *CloudWorkloadSecurityAgentRuleActionLog `json:"log,omitempty"`
	// The metadata action applied on the scope matching the rule
	Metadata *CloudWorkloadSecurityAgentRuleActionMetadata `json:"metadata,omitempty"`
	// The network filter action applied on the network traffic matching the rule.
	NetworkFilter *CloudWorkloadSecurityAgentRuleActionNetworkFilter `json:"network_filter,omitempty"`
	// The set action applied on the scope matching the rule
	Set *CloudWorkloadSecurityAgentRuleActionSet `json:"set,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleAction instantiates a new CloudWorkloadSecurityAgentRuleAction object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleAction() *CloudWorkloadSecurityAgentRuleAction {
	this := CloudWorkloadSecurityAgentRuleAction{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleActionWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleAction object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleActionWithDefaults() *CloudWorkloadSecurityAgentRuleAction {
	this := CloudWorkloadSecurityAgentRuleAction{}
	return &this
}

// GetCoredump returns the Coredump field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetCoredump() CloudWorkloadSecurityAgentRuleActionCoreDump {
	if o == nil || o.Coredump == nil {
		var ret CloudWorkloadSecurityAgentRuleActionCoreDump
		return ret
	}
	return *o.Coredump
}

// GetCoredumpOk returns a tuple with the Coredump field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetCoredumpOk() (*CloudWorkloadSecurityAgentRuleActionCoreDump, bool) {
	if o == nil || o.Coredump == nil {
		return nil, false
	}
	return o.Coredump, true
}

// HasCoredump returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasCoredump() bool {
	return o != nil && o.Coredump != nil
}

// SetCoredump gets a reference to the given CloudWorkloadSecurityAgentRuleActionCoreDump and assigns it to the Coredump field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetCoredump(v CloudWorkloadSecurityAgentRuleActionCoreDump) {
	o.Coredump = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetDisabled() bool {
	if o == nil || o.Disabled == nil {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetDisabledOk() (*bool, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasDisabled() bool {
	return o != nil && o.Disabled != nil
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasFilter() bool {
	return o != nil && o.Filter != nil
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetFilter(v string) {
	o.Filter = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetHash() CloudWorkloadSecurityAgentRuleActionHash {
	if o == nil || o.Hash == nil {
		var ret CloudWorkloadSecurityAgentRuleActionHash
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetHashOk() (*CloudWorkloadSecurityAgentRuleActionHash, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasHash() bool {
	return o != nil && o.Hash != nil
}

// SetHash gets a reference to the given CloudWorkloadSecurityAgentRuleActionHash and assigns it to the Hash field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetHash(v CloudWorkloadSecurityAgentRuleActionHash) {
	o.Hash = &v
}

// GetKill returns the Kill field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetKill() CloudWorkloadSecurityAgentRuleKill {
	if o == nil || o.Kill == nil {
		var ret CloudWorkloadSecurityAgentRuleKill
		return ret
	}
	return *o.Kill
}

// GetKillOk returns a tuple with the Kill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetKillOk() (*CloudWorkloadSecurityAgentRuleKill, bool) {
	if o == nil || o.Kill == nil {
		return nil, false
	}
	return o.Kill, true
}

// HasKill returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasKill() bool {
	return o != nil && o.Kill != nil
}

// SetKill gets a reference to the given CloudWorkloadSecurityAgentRuleKill and assigns it to the Kill field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetKill(v CloudWorkloadSecurityAgentRuleKill) {
	o.Kill = &v
}

// GetLog returns the Log field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetLog() CloudWorkloadSecurityAgentRuleActionLog {
	if o == nil || o.Log == nil {
		var ret CloudWorkloadSecurityAgentRuleActionLog
		return ret
	}
	return *o.Log
}

// GetLogOk returns a tuple with the Log field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetLogOk() (*CloudWorkloadSecurityAgentRuleActionLog, bool) {
	if o == nil || o.Log == nil {
		return nil, false
	}
	return o.Log, true
}

// HasLog returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasLog() bool {
	return o != nil && o.Log != nil
}

// SetLog gets a reference to the given CloudWorkloadSecurityAgentRuleActionLog and assigns it to the Log field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetLog(v CloudWorkloadSecurityAgentRuleActionLog) {
	o.Log = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetMetadata() CloudWorkloadSecurityAgentRuleActionMetadata {
	if o == nil || o.Metadata == nil {
		var ret CloudWorkloadSecurityAgentRuleActionMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetMetadataOk() (*CloudWorkloadSecurityAgentRuleActionMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given CloudWorkloadSecurityAgentRuleActionMetadata and assigns it to the Metadata field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetMetadata(v CloudWorkloadSecurityAgentRuleActionMetadata) {
	o.Metadata = &v
}

// GetNetworkFilter returns the NetworkFilter field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetNetworkFilter() CloudWorkloadSecurityAgentRuleActionNetworkFilter {
	if o == nil || o.NetworkFilter == nil {
		var ret CloudWorkloadSecurityAgentRuleActionNetworkFilter
		return ret
	}
	return *o.NetworkFilter
}

// GetNetworkFilterOk returns a tuple with the NetworkFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetNetworkFilterOk() (*CloudWorkloadSecurityAgentRuleActionNetworkFilter, bool) {
	if o == nil || o.NetworkFilter == nil {
		return nil, false
	}
	return o.NetworkFilter, true
}

// HasNetworkFilter returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasNetworkFilter() bool {
	return o != nil && o.NetworkFilter != nil
}

// SetNetworkFilter gets a reference to the given CloudWorkloadSecurityAgentRuleActionNetworkFilter and assigns it to the NetworkFilter field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetNetworkFilter(v CloudWorkloadSecurityAgentRuleActionNetworkFilter) {
	o.NetworkFilter = &v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleAction) GetSet() CloudWorkloadSecurityAgentRuleActionSet {
	if o == nil || o.Set == nil {
		var ret CloudWorkloadSecurityAgentRuleActionSet
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) GetSetOk() (*CloudWorkloadSecurityAgentRuleActionSet, bool) {
	if o == nil || o.Set == nil {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleAction) HasSet() bool {
	return o != nil && o.Set != nil
}

// SetSet gets a reference to the given CloudWorkloadSecurityAgentRuleActionSet and assigns it to the Set field.
func (o *CloudWorkloadSecurityAgentRuleAction) SetSet(v CloudWorkloadSecurityAgentRuleActionSet) {
	o.Set = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Coredump != nil {
		toSerialize["coredump"] = o.Coredump
	}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Kill != nil {
		toSerialize["kill"] = o.Kill
	}
	if o.Log != nil {
		toSerialize["log"] = o.Log
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.NetworkFilter != nil {
		toSerialize["network_filter"] = o.NetworkFilter
	}
	if o.Set != nil {
		toSerialize["set"] = o.Set
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentRuleAction) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Coredump      *CloudWorkloadSecurityAgentRuleActionCoreDump      `json:"coredump,omitempty"`
		Disabled      *bool                                              `json:"disabled,omitempty"`
		Filter        *string                                            `json:"filter,omitempty"`
		Hash          *CloudWorkloadSecurityAgentRuleActionHash          `json:"hash,omitempty"`
		Kill          *CloudWorkloadSecurityAgentRuleKill                `json:"kill,omitempty"`
		Log           *CloudWorkloadSecurityAgentRuleActionLog           `json:"log,omitempty"`
		Metadata      *CloudWorkloadSecurityAgentRuleActionMetadata      `json:"metadata,omitempty"`
		NetworkFilter *CloudWorkloadSecurityAgentRuleActionNetworkFilter `json:"network_filter,omitempty"`
		Set           *CloudWorkloadSecurityAgentRuleActionSet           `json:"set,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"coredump", "disabled", "filter", "hash", "kill", "log", "metadata", "network_filter", "set"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Coredump != nil && all.Coredump.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Coredump = all.Coredump
	o.Disabled = all.Disabled
	o.Filter = all.Filter
	if all.Hash != nil && all.Hash.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Hash = all.Hash
	if all.Kill != nil && all.Kill.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Kill = all.Kill
	if all.Log != nil && all.Log.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Log = all.Log
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	if all.NetworkFilter != nil && all.NetworkFilter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NetworkFilter = all.NetworkFilter
	if all.Set != nil && all.Set.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Set = all.Set

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
