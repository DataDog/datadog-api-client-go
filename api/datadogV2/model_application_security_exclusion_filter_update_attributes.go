// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterUpdateAttributes Update an existing Application Security Exclusion filter.
type ApplicationSecurityExclusionFilterUpdateAttributes struct {
	// The description of the Exclusion filter.
	Description string `json:"description"`
	// Whether the Exclusion filter is enabled.
	Enabled bool `json:"enabled"`
	// The IPs list for the exclusion filter.
	IpList []string `json:"ip_list,omitempty"`
	// A list of parameters for the exclusion filter.
	Parameters []string `json:"parameters,omitempty"`
	// The path glob for the exclusion filter.
	PathGlob string `json:"path_glob"`
	// A list of rules targeted by the exclusion filter.
	RulesTarget []ApplicationSecurityExclusionFilterRulesTarget `json:"rules_target,omitempty"`
	// The scope of the exclusion filter.
	Scope []ApplicationSecurityExclusionFilterScope `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterUpdateAttributes instantiates a new ApplicationSecurityExclusionFilterUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterUpdateAttributes(description string, enabled bool, pathGlob string) *ApplicationSecurityExclusionFilterUpdateAttributes {
	this := ApplicationSecurityExclusionFilterUpdateAttributes{}
	this.Description = description
	this.Enabled = enabled
	this.PathGlob = pathGlob
	return &this
}

// NewApplicationSecurityExclusionFilterUpdateAttributesWithDefaults instantiates a new ApplicationSecurityExclusionFilterUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterUpdateAttributesWithDefaults() *ApplicationSecurityExclusionFilterUpdateAttributes {
	this := ApplicationSecurityExclusionFilterUpdateAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetIpList() []string {
	if o == nil || o.IpList == nil {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetIpListOk() (*[]string, bool) {
	if o == nil || o.IpList == nil {
		return nil, false
	}
	return &o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) HasIpList() bool {
	return o != nil && o.IpList != nil
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) SetIpList(v []string) {
	o.IpList = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetParameters() []string {
	if o == nil || o.Parameters == nil {
		var ret []string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetParametersOk() (*[]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []string and assigns it to the Parameters field.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) SetParameters(v []string) {
	o.Parameters = v
}

// GetPathGlob returns the PathGlob field value.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetPathGlob() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PathGlob
}

// GetPathGlobOk returns a tuple with the PathGlob field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetPathGlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PathGlob, true
}

// SetPathGlob sets field value.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) SetPathGlob(v string) {
	o.PathGlob = v
}

// GetRulesTarget returns the RulesTarget field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetRulesTarget() []ApplicationSecurityExclusionFilterRulesTarget {
	if o == nil || o.RulesTarget == nil {
		var ret []ApplicationSecurityExclusionFilterRulesTarget
		return ret
	}
	return o.RulesTarget
}

// GetRulesTargetOk returns a tuple with the RulesTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetRulesTargetOk() (*[]ApplicationSecurityExclusionFilterRulesTarget, bool) {
	if o == nil || o.RulesTarget == nil {
		return nil, false
	}
	return &o.RulesTarget, true
}

// HasRulesTarget returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) HasRulesTarget() bool {
	return o != nil && o.RulesTarget != nil
}

// SetRulesTarget gets a reference to the given []ApplicationSecurityExclusionFilterRulesTarget and assigns it to the RulesTarget field.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) SetRulesTarget(v []ApplicationSecurityExclusionFilterRulesTarget) {
	o.RulesTarget = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetScope() []ApplicationSecurityExclusionFilterScope {
	if o == nil || o.Scope == nil {
		var ret []ApplicationSecurityExclusionFilterScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) GetScopeOk() (*[]ApplicationSecurityExclusionFilterScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []ApplicationSecurityExclusionFilterScope and assigns it to the Scope field.
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) SetScope(v []ApplicationSecurityExclusionFilterScope) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["enabled"] = o.Enabled
	if o.IpList != nil {
		toSerialize["ip_list"] = o.IpList
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	toSerialize["path_glob"] = o.PathGlob
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
func (o *ApplicationSecurityExclusionFilterUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                         `json:"description"`
		Enabled     *bool                                           `json:"enabled"`
		IpList      []string                                        `json:"ip_list,omitempty"`
		Parameters  []string                                        `json:"parameters,omitempty"`
		PathGlob    *string                                         `json:"path_glob"`
		RulesTarget []ApplicationSecurityExclusionFilterRulesTarget `json:"rules_target,omitempty"`
		Scope       []ApplicationSecurityExclusionFilterScope       `json:"scope,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.PathGlob == nil {
		return fmt.Errorf("required field path_glob missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "ip_list", "parameters", "path_glob", "rules_target", "scope"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Enabled = *all.Enabled
	o.IpList = all.IpList
	o.Parameters = all.Parameters
	o.PathGlob = *all.PathGlob
	o.RulesTarget = all.RulesTarget
	o.Scope = all.Scope

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
