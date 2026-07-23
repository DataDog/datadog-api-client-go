// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3RepositorySpec The definition of Entity V3 Repository Spec object.
type EntityV3RepositorySpec struct {
	// A list of components the repository is a part of.
	ComponentOf []string `json:"componentOf,omitempty"`
	// A list of components that the repository is a dependency of.
	DependencyOf []string `json:"dependencyOf,omitempty"`
	// A list of components that the repository depends on.
	DependsOn []string `json:"dependsOn,omitempty"`
	// The lifecycle state of the repository.
	Lifecycle *string `json:"lifecycle,omitempty"`
	// The importance of the repository.
	Tier *string `json:"tier,omitempty"`
	// The type of repository.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3RepositorySpec instantiates a new EntityV3RepositorySpec object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3RepositorySpec() *EntityV3RepositorySpec {
	this := EntityV3RepositorySpec{}
	return &this
}

// NewEntityV3RepositorySpecWithDefaults instantiates a new EntityV3RepositorySpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3RepositorySpecWithDefaults() *EntityV3RepositorySpec {
	this := EntityV3RepositorySpec{}
	return &this
}

// GetComponentOf returns the ComponentOf field value if set, zero value otherwise.
func (o *EntityV3RepositorySpec) GetComponentOf() []string {
	if o == nil || o.ComponentOf == nil {
		var ret []string
		return ret
	}
	return o.ComponentOf
}

// GetComponentOfOk returns a tuple with the ComponentOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3RepositorySpec) GetComponentOfOk() (*[]string, bool) {
	if o == nil || o.ComponentOf == nil {
		return nil, false
	}
	return &o.ComponentOf, true
}

// HasComponentOf returns a boolean if a field has been set.
func (o *EntityV3RepositorySpec) HasComponentOf() bool {
	return o != nil && o.ComponentOf != nil
}

// SetComponentOf gets a reference to the given []string and assigns it to the ComponentOf field.
func (o *EntityV3RepositorySpec) SetComponentOf(v []string) {
	o.ComponentOf = v
}

// GetDependencyOf returns the DependencyOf field value if set, zero value otherwise.
func (o *EntityV3RepositorySpec) GetDependencyOf() []string {
	if o == nil || o.DependencyOf == nil {
		var ret []string
		return ret
	}
	return o.DependencyOf
}

// GetDependencyOfOk returns a tuple with the DependencyOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3RepositorySpec) GetDependencyOfOk() (*[]string, bool) {
	if o == nil || o.DependencyOf == nil {
		return nil, false
	}
	return &o.DependencyOf, true
}

// HasDependencyOf returns a boolean if a field has been set.
func (o *EntityV3RepositorySpec) HasDependencyOf() bool {
	return o != nil && o.DependencyOf != nil
}

// SetDependencyOf gets a reference to the given []string and assigns it to the DependencyOf field.
func (o *EntityV3RepositorySpec) SetDependencyOf(v []string) {
	o.DependencyOf = v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise.
func (o *EntityV3RepositorySpec) GetDependsOn() []string {
	if o == nil || o.DependsOn == nil {
		var ret []string
		return ret
	}
	return o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3RepositorySpec) GetDependsOnOk() (*[]string, bool) {
	if o == nil || o.DependsOn == nil {
		return nil, false
	}
	return &o.DependsOn, true
}

// HasDependsOn returns a boolean if a field has been set.
func (o *EntityV3RepositorySpec) HasDependsOn() bool {
	return o != nil && o.DependsOn != nil
}

// SetDependsOn gets a reference to the given []string and assigns it to the DependsOn field.
func (o *EntityV3RepositorySpec) SetDependsOn(v []string) {
	o.DependsOn = v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *EntityV3RepositorySpec) GetLifecycle() string {
	if o == nil || o.Lifecycle == nil {
		var ret string
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3RepositorySpec) GetLifecycleOk() (*string, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *EntityV3RepositorySpec) HasLifecycle() bool {
	return o != nil && o.Lifecycle != nil
}

// SetLifecycle gets a reference to the given string and assigns it to the Lifecycle field.
func (o *EntityV3RepositorySpec) SetLifecycle(v string) {
	o.Lifecycle = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *EntityV3RepositorySpec) GetTier() string {
	if o == nil || o.Tier == nil {
		var ret string
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3RepositorySpec) GetTierOk() (*string, bool) {
	if o == nil || o.Tier == nil {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *EntityV3RepositorySpec) HasTier() bool {
	return o != nil && o.Tier != nil
}

// SetTier gets a reference to the given string and assigns it to the Tier field.
func (o *EntityV3RepositorySpec) SetTier(v string) {
	o.Tier = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EntityV3RepositorySpec) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3RepositorySpec) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EntityV3RepositorySpec) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EntityV3RepositorySpec) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3RepositorySpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ComponentOf != nil {
		toSerialize["componentOf"] = o.ComponentOf
	}
	if o.DependencyOf != nil {
		toSerialize["dependencyOf"] = o.DependencyOf
	}
	if o.DependsOn != nil {
		toSerialize["dependsOn"] = o.DependsOn
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.Tier != nil {
		toSerialize["tier"] = o.Tier
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3RepositorySpec) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ComponentOf  []string `json:"componentOf,omitempty"`
		DependencyOf []string `json:"dependencyOf,omitempty"`
		DependsOn    []string `json:"dependsOn,omitempty"`
		Lifecycle    *string  `json:"lifecycle,omitempty"`
		Tier         *string  `json:"tier,omitempty"`
		Type         *string  `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.ComponentOf = all.ComponentOf
	o.DependencyOf = all.DependencyOf
	o.DependsOn = all.DependsOn
	o.Lifecycle = all.Lifecycle
	o.Tier = all.Tier
	o.Type = all.Type

	return nil
}
