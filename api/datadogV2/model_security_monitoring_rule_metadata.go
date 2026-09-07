// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringRuleMetadata Metadata associated with the rule.
type SecurityMonitoringRuleMetadata struct {
	// Entities associated with the rule, or null when metadata is not requested.
	Entities []interface{} `json:"entities,omitempty"`
	// Sources associated with the rule, or null when metadata is not requested.
	Sources datadog.NullableList[string] `json:"sources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringRuleMetadata instantiates a new SecurityMonitoringRuleMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringRuleMetadata() *SecurityMonitoringRuleMetadata {
	this := SecurityMonitoringRuleMetadata{}
	return &this
}

// NewSecurityMonitoringRuleMetadataWithDefaults instantiates a new SecurityMonitoringRuleMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringRuleMetadataWithDefaults() *SecurityMonitoringRuleMetadata {
	this := SecurityMonitoringRuleMetadata{}
	return &this
}

// GetEntities returns the Entities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringRuleMetadata) GetEntities() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringRuleMetadata) GetEntitiesOk() (*[]interface{}, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return &o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleMetadata) HasEntities() bool {
	return o != nil && o.Entities != nil
}

// SetEntities gets a reference to the given []interface{} and assigns it to the Entities field.
func (o *SecurityMonitoringRuleMetadata) SetEntities(v []interface{}) {
	o.Entities = v
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SecurityMonitoringRuleMetadata) GetSources() []string {
	if o == nil || o.Sources.Get() == nil {
		var ret []string
		return ret
	}
	return *o.Sources.Get()
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SecurityMonitoringRuleMetadata) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sources.Get(), o.Sources.IsSet()
}

// HasSources returns a boolean if a field has been set.
func (o *SecurityMonitoringRuleMetadata) HasSources() bool {
	return o != nil && o.Sources.IsSet()
}

// SetSources gets a reference to the given datadog.NullableList[string] and assigns it to the Sources field.
func (o *SecurityMonitoringRuleMetadata) SetSources(v []string) {
	o.Sources.Set(&v)
}

// SetSourcesNil sets the value for Sources to be an explicit nil.
func (o *SecurityMonitoringRuleMetadata) SetSourcesNil() {
	o.Sources.Set(nil)
}

// UnsetSources ensures that no value is present for Sources, not even an explicit nil.
func (o *SecurityMonitoringRuleMetadata) UnsetSources() {
	o.Sources.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringRuleMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.Sources.IsSet() {
		toSerialize["sources"] = o.Sources.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringRuleMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Entities []interface{}                `json:"entities,omitempty"`
		Sources  datadog.NullableList[string] `json:"sources,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entities", "sources"})
	} else {
		return err
	}
	o.Entities = all.Entities
	o.Sources = all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
