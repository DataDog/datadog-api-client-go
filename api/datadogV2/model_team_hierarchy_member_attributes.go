// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamHierarchyMemberAttributes Membership attributes relative to the requested team.
type TeamHierarchyMemberAttributes struct {
	// Role on the requested team: admin or member. Null for inherited members.
	DirectRole datadog.NullableString `json:"direct_role"`
	// Whether the user belongs directly to the requested team.
	IsDirect bool `json:"is_direct"`
	// Membership provisioner. Null for inherited members.
	ProvisionedBy datadog.NullableString `json:"provisioned_by"`
	// Identifier supplied by the membership provisioner. Null when unavailable or for inherited members.
	ProvisionedById datadog.NullableString `json:"provisioned_by_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamHierarchyMemberAttributes instantiates a new TeamHierarchyMemberAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamHierarchyMemberAttributes(directRole datadog.NullableString, isDirect bool, provisionedBy datadog.NullableString, provisionedById datadog.NullableString) *TeamHierarchyMemberAttributes {
	this := TeamHierarchyMemberAttributes{}
	this.DirectRole = directRole
	this.IsDirect = isDirect
	this.ProvisionedBy = provisionedBy
	this.ProvisionedById = provisionedById
	return &this
}

// NewTeamHierarchyMemberAttributesWithDefaults instantiates a new TeamHierarchyMemberAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamHierarchyMemberAttributesWithDefaults() *TeamHierarchyMemberAttributes {
	this := TeamHierarchyMemberAttributes{}
	return &this
}

// GetDirectRole returns the DirectRole field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *TeamHierarchyMemberAttributes) GetDirectRole() string {
	if o == nil || o.DirectRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.DirectRole.Get()
}

// GetDirectRoleOk returns a tuple with the DirectRole field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMemberAttributes) GetDirectRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectRole.Get(), o.DirectRole.IsSet()
}

// SetDirectRole sets field value.
func (o *TeamHierarchyMemberAttributes) SetDirectRole(v string) {
	o.DirectRole.Set(&v)
}

// GetIsDirect returns the IsDirect field value.
func (o *TeamHierarchyMemberAttributes) GetIsDirect() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDirect
}

// GetIsDirectOk returns a tuple with the IsDirect field value
// and a boolean to check if the value has been set.
func (o *TeamHierarchyMemberAttributes) GetIsDirectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDirect, true
}

// SetIsDirect sets field value.
func (o *TeamHierarchyMemberAttributes) SetIsDirect(v bool) {
	o.IsDirect = v
}

// GetProvisionedBy returns the ProvisionedBy field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *TeamHierarchyMemberAttributes) GetProvisionedBy() string {
	if o == nil || o.ProvisionedBy.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProvisionedBy.Get()
}

// GetProvisionedByOk returns a tuple with the ProvisionedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMemberAttributes) GetProvisionedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisionedBy.Get(), o.ProvisionedBy.IsSet()
}

// SetProvisionedBy sets field value.
func (o *TeamHierarchyMemberAttributes) SetProvisionedBy(v string) {
	o.ProvisionedBy.Set(&v)
}

// GetProvisionedById returns the ProvisionedById field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *TeamHierarchyMemberAttributes) GetProvisionedById() string {
	if o == nil || o.ProvisionedById.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProvisionedById.Get()
}

// GetProvisionedByIdOk returns a tuple with the ProvisionedById field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *TeamHierarchyMemberAttributes) GetProvisionedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProvisionedById.Get(), o.ProvisionedById.IsSet()
}

// SetProvisionedById sets field value.
func (o *TeamHierarchyMemberAttributes) SetProvisionedById(v string) {
	o.ProvisionedById.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamHierarchyMemberAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["direct_role"] = o.DirectRole.Get()
	toSerialize["is_direct"] = o.IsDirect
	toSerialize["provisioned_by"] = o.ProvisionedBy.Get()
	toSerialize["provisioned_by_id"] = o.ProvisionedById.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamHierarchyMemberAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DirectRole      datadog.NullableString `json:"direct_role"`
		IsDirect        *bool                  `json:"is_direct"`
		ProvisionedBy   datadog.NullableString `json:"provisioned_by"`
		ProvisionedById datadog.NullableString `json:"provisioned_by_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.DirectRole.IsSet() {
		return fmt.Errorf("required field direct_role missing")
	}
	if all.IsDirect == nil {
		return fmt.Errorf("required field is_direct missing")
	}
	if !all.ProvisionedBy.IsSet() {
		return fmt.Errorf("required field provisioned_by missing")
	}
	if !all.ProvisionedById.IsSet() {
		return fmt.Errorf("required field provisioned_by_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"direct_role", "is_direct", "provisioned_by", "provisioned_by_id"})
	} else {
		return err
	}
	o.DirectRole = all.DirectRole
	o.IsDirect = *all.IsDirect
	o.ProvisionedBy = all.ProvisionedBy
	o.ProvisionedById = all.ProvisionedById

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
