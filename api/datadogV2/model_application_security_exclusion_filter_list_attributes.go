// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterListAttributes The attributes of the Application Security exclusion filter.
type ApplicationSecurityExclusionFilterListAttributes struct {
	// A description for the exclusion filter.
	Description *string `json:"description,omitempty"`
	// Indicates whether the exclusion filter is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The IPs list for the exclusion filter.
	IpList []string `json:"ip_list,omitempty"`
	// Metadata about the exclusion filter.
	Metadata *ApplicationSecurityExclusionFilterMetadata `json:"metadata,omitempty"`
	// A list of parameters for the exclusion filter.
	Parameters []string `json:"parameters,omitempty"`
	// The path glob for the exclusion filter.
	PathGlob *string `json:"path_glob,omitempty"`
	// A list of rules targeted by the exclusion filter.
	RulesTarget []ApplicationSecurityExclusionFilterListRulesTarget `json:"rules_target,omitempty"`
	// The scope of the exclusion filter.
	Scope []ApplicationSecurityExclusionFilterScope `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterListAttributes instantiates a new ApplicationSecurityExclusionFilterListAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterListAttributes() *ApplicationSecurityExclusionFilterListAttributes {
	this := ApplicationSecurityExclusionFilterListAttributes{}
	return &this
}

// NewApplicationSecurityExclusionFilterListAttributesWithDefaults instantiates a new ApplicationSecurityExclusionFilterListAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterListAttributesWithDefaults() *ApplicationSecurityExclusionFilterListAttributes {
	this := ApplicationSecurityExclusionFilterListAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetIpList() []string {
	if o == nil || o.IpList == nil {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetIpListOk() (*[]string, bool) {
	if o == nil || o.IpList == nil {
		return nil, false
	}
	return &o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasIpList() bool {
	return o != nil && o.IpList != nil
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetIpList(v []string) {
	o.IpList = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetMetadata() ApplicationSecurityExclusionFilterMetadata {
	if o == nil || o.Metadata == nil {
		var ret ApplicationSecurityExclusionFilterMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetMetadataOk() (*ApplicationSecurityExclusionFilterMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given ApplicationSecurityExclusionFilterMetadata and assigns it to the Metadata field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetMetadata(v ApplicationSecurityExclusionFilterMetadata) {
	o.Metadata = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetParameters() []string {
	if o == nil || o.Parameters == nil {
		var ret []string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetParametersOk() (*[]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []string and assigns it to the Parameters field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetParameters(v []string) {
	o.Parameters = v
}

// GetPathGlob returns the PathGlob field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetPathGlob() string {
	if o == nil || o.PathGlob == nil {
		var ret string
		return ret
	}
	return *o.PathGlob
}

// GetPathGlobOk returns a tuple with the PathGlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetPathGlobOk() (*string, bool) {
	if o == nil || o.PathGlob == nil {
		return nil, false
	}
	return o.PathGlob, true
}

// HasPathGlob returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasPathGlob() bool {
	return o != nil && o.PathGlob != nil
}

// SetPathGlob gets a reference to the given string and assigns it to the PathGlob field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetPathGlob(v string) {
	o.PathGlob = &v
}

// GetRulesTarget returns the RulesTarget field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetRulesTarget() []ApplicationSecurityExclusionFilterListRulesTarget {
	if o == nil || o.RulesTarget == nil {
		var ret []ApplicationSecurityExclusionFilterListRulesTarget
		return ret
	}
	return o.RulesTarget
}

// GetRulesTargetOk returns a tuple with the RulesTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetRulesTargetOk() (*[]ApplicationSecurityExclusionFilterListRulesTarget, bool) {
	if o == nil || o.RulesTarget == nil {
		return nil, false
	}
	return &o.RulesTarget, true
}

// HasRulesTarget returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasRulesTarget() bool {
	return o != nil && o.RulesTarget != nil
}

// SetRulesTarget gets a reference to the given []ApplicationSecurityExclusionFilterListRulesTarget and assigns it to the RulesTarget field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetRulesTarget(v []ApplicationSecurityExclusionFilterListRulesTarget) {
	o.RulesTarget = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetScope() []ApplicationSecurityExclusionFilterScope {
	if o == nil || o.Scope == nil {
		var ret []ApplicationSecurityExclusionFilterScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) GetScopeOk() (*[]ApplicationSecurityExclusionFilterScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterListAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []ApplicationSecurityExclusionFilterScope and assigns it to the Scope field.
func (o *ApplicationSecurityExclusionFilterListAttributes) SetScope(v []ApplicationSecurityExclusionFilterScope) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterListAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.IpList != nil {
		toSerialize["ip_list"] = o.IpList
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.PathGlob != nil {
		toSerialize["path_glob"] = o.PathGlob
	}
	if o.RulesTarget != nil {
		toSerialize["rules_target"] = o.RulesTarget
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityExclusionFilterListAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                             `json:"description,omitempty"`
		Enabled     *bool                                               `json:"enabled,omitempty"`
		IpList      []string                                            `json:"ip_list,omitempty"`
		Metadata    *ApplicationSecurityExclusionFilterMetadata         `json:"metadata,omitempty"`
		Parameters  []string                                            `json:"parameters,omitempty"`
		PathGlob    *string                                             `json:"path_glob,omitempty"`
		RulesTarget []ApplicationSecurityExclusionFilterListRulesTarget `json:"rules_target,omitempty"`
		Scope       []ApplicationSecurityExclusionFilterScope           `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "ip_list", "metadata", "parameters", "path_glob", "rules_target", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.IpList = all.IpList
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.Parameters = all.Parameters
	o.PathGlob = all.PathGlob
	o.RulesTarget = all.RulesTarget
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
