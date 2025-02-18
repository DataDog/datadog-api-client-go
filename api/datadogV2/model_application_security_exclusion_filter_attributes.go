// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityExclusionFilterAttributes Attributes describing an Application Security exclusion filter.
type ApplicationSecurityExclusionFilterAttributes struct {
	// A description for the exclusion filter.
	Description string `json:"description"`
	// Indicates whether the exclusion filter is enabled.
	Enabled bool `json:"enabled"`
	// The event query matched by the legacy exclusion filter.
	EventQuery *string `json:"event_query,omitempty"`
	// The client IP addresses matched by the exclusion filter (CIDR notation is supported).
	IpList []string `json:"ip_list,omitempty"`
	// The action taken when the exclusion filter matches. When set to `monitor`, security traces are emitted but the requests are not blocked. By default, security traces are not emitted and the requests are not blocked.
	OnMatch *ApplicationSecurityExclusionFilterOnMatch `json:"on_match,omitempty"`
	// A list of parameters matched by the exclusion filter in the HTTP query string and HTTP request body. Nested parameters can be matched by joining fields with a dot character.
	Parameters []string `json:"parameters,omitempty"`
	// The HTTP path glob expression matched by the exclusion filter.
	PathGlob *string `json:"path_glob,omitempty"`
	// The WAF rules targeted by the exclusion filter.
	RulesTarget []ApplicationSecurityExclusionFilterRulesTarget `json:"rules_target,omitempty"`
	// The services where the exclusion filter is deployed.
	Scope []ApplicationSecurityExclusionFilterScope `json:"scope,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityExclusionFilterAttributes instantiates a new ApplicationSecurityExclusionFilterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityExclusionFilterAttributes(description string, enabled bool) *ApplicationSecurityExclusionFilterAttributes {
	this := ApplicationSecurityExclusionFilterAttributes{}
	this.Description = description
	this.Enabled = enabled
	return &this
}

// NewApplicationSecurityExclusionFilterAttributesWithDefaults instantiates a new ApplicationSecurityExclusionFilterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityExclusionFilterAttributesWithDefaults() *ApplicationSecurityExclusionFilterAttributes {
	this := ApplicationSecurityExclusionFilterAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *ApplicationSecurityExclusionFilterAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ApplicationSecurityExclusionFilterAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEnabled returns the Enabled field value.
func (o *ApplicationSecurityExclusionFilterAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ApplicationSecurityExclusionFilterAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterAttributes) GetEventQuery() string {
	if o == nil || o.EventQuery == nil {
		var ret string
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetEventQueryOk() (*string, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) HasEventQuery() bool {
	return o != nil && o.EventQuery != nil
}

// SetEventQuery gets a reference to the given string and assigns it to the EventQuery field.
func (o *ApplicationSecurityExclusionFilterAttributes) SetEventQuery(v string) {
	o.EventQuery = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterAttributes) GetIpList() []string {
	if o == nil || o.IpList == nil {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetIpListOk() (*[]string, bool) {
	if o == nil || o.IpList == nil {
		return nil, false
	}
	return &o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) HasIpList() bool {
	return o != nil && o.IpList != nil
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *ApplicationSecurityExclusionFilterAttributes) SetIpList(v []string) {
	o.IpList = v
}

// GetOnMatch returns the OnMatch field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterAttributes) GetOnMatch() ApplicationSecurityExclusionFilterOnMatch {
	if o == nil || o.OnMatch == nil {
		var ret ApplicationSecurityExclusionFilterOnMatch
		return ret
	}
	return *o.OnMatch
}

// GetOnMatchOk returns a tuple with the OnMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetOnMatchOk() (*ApplicationSecurityExclusionFilterOnMatch, bool) {
	if o == nil || o.OnMatch == nil {
		return nil, false
	}
	return o.OnMatch, true
}

// HasOnMatch returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) HasOnMatch() bool {
	return o != nil && o.OnMatch != nil
}

// SetOnMatch gets a reference to the given ApplicationSecurityExclusionFilterOnMatch and assigns it to the OnMatch field.
func (o *ApplicationSecurityExclusionFilterAttributes) SetOnMatch(v ApplicationSecurityExclusionFilterOnMatch) {
	o.OnMatch = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterAttributes) GetParameters() []string {
	if o == nil || o.Parameters == nil {
		var ret []string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetParametersOk() (*[]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []string and assigns it to the Parameters field.
func (o *ApplicationSecurityExclusionFilterAttributes) SetParameters(v []string) {
	o.Parameters = v
}

// GetPathGlob returns the PathGlob field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterAttributes) GetPathGlob() string {
	if o == nil || o.PathGlob == nil {
		var ret string
		return ret
	}
	return *o.PathGlob
}

// GetPathGlobOk returns a tuple with the PathGlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetPathGlobOk() (*string, bool) {
	if o == nil || o.PathGlob == nil {
		return nil, false
	}
	return o.PathGlob, true
}

// HasPathGlob returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) HasPathGlob() bool {
	return o != nil && o.PathGlob != nil
}

// SetPathGlob gets a reference to the given string and assigns it to the PathGlob field.
func (o *ApplicationSecurityExclusionFilterAttributes) SetPathGlob(v string) {
	o.PathGlob = &v
}

// GetRulesTarget returns the RulesTarget field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterAttributes) GetRulesTarget() []ApplicationSecurityExclusionFilterRulesTarget {
	if o == nil || o.RulesTarget == nil {
		var ret []ApplicationSecurityExclusionFilterRulesTarget
		return ret
	}
	return o.RulesTarget
}

// GetRulesTargetOk returns a tuple with the RulesTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetRulesTargetOk() (*[]ApplicationSecurityExclusionFilterRulesTarget, bool) {
	if o == nil || o.RulesTarget == nil {
		return nil, false
	}
	return &o.RulesTarget, true
}

// HasRulesTarget returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) HasRulesTarget() bool {
	return o != nil && o.RulesTarget != nil
}

// SetRulesTarget gets a reference to the given []ApplicationSecurityExclusionFilterRulesTarget and assigns it to the RulesTarget field.
func (o *ApplicationSecurityExclusionFilterAttributes) SetRulesTarget(v []ApplicationSecurityExclusionFilterRulesTarget) {
	o.RulesTarget = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationSecurityExclusionFilterAttributes) GetScope() []ApplicationSecurityExclusionFilterScope {
	if o == nil || o.Scope == nil {
		var ret []ApplicationSecurityExclusionFilterScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) GetScopeOk() (*[]ApplicationSecurityExclusionFilterScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationSecurityExclusionFilterAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []ApplicationSecurityExclusionFilterScope and assigns it to the Scope field.
func (o *ApplicationSecurityExclusionFilterAttributes) SetScope(v []ApplicationSecurityExclusionFilterScope) {
	o.Scope = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityExclusionFilterAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["enabled"] = o.Enabled
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.IpList != nil {
		toSerialize["ip_list"] = o.IpList
	}
	if o.OnMatch != nil {
		toSerialize["on_match"] = o.OnMatch
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
func (o *ApplicationSecurityExclusionFilterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                         `json:"description"`
		Enabled     *bool                                           `json:"enabled"`
		EventQuery  *string                                         `json:"event_query,omitempty"`
		IpList      []string                                        `json:"ip_list,omitempty"`
		OnMatch     *ApplicationSecurityExclusionFilterOnMatch      `json:"on_match,omitempty"`
		Parameters  []string                                        `json:"parameters,omitempty"`
		PathGlob    *string                                         `json:"path_glob,omitempty"`
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
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "event_query", "ip_list", "on_match", "parameters", "path_glob", "rules_target", "scope"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	o.Enabled = *all.Enabled
	o.EventQuery = all.EventQuery
	o.IpList = all.IpList
	if all.OnMatch != nil && !all.OnMatch.IsValid() {
		hasInvalidField = true
	} else {
		o.OnMatch = all.OnMatch
	}
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
