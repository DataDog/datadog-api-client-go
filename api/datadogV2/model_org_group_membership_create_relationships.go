// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipCreateRelationships Relationships for adding organizations to an org group.
type OrgGroupMembershipCreateRelationships struct {
	// Relationship to a single org group.
	OrgGroup OrgGroupRelationshipToOne `json:"org_group"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupMembershipCreateRelationships instantiates a new OrgGroupMembershipCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupMembershipCreateRelationships(orgGroup OrgGroupRelationshipToOne) *OrgGroupMembershipCreateRelationships {
	this := OrgGroupMembershipCreateRelationships{}
	this.OrgGroup = orgGroup
	return &this
}

// NewOrgGroupMembershipCreateRelationshipsWithDefaults instantiates a new OrgGroupMembershipCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupMembershipCreateRelationshipsWithDefaults() *OrgGroupMembershipCreateRelationships {
	this := OrgGroupMembershipCreateRelationships{}
	return &this
}

// GetOrgGroup returns the OrgGroup field value.
func (o *OrgGroupMembershipCreateRelationships) GetOrgGroup() OrgGroupRelationshipToOne {
	if o == nil {
		var ret OrgGroupRelationshipToOne
		return ret
	}
	return o.OrgGroup
}

// GetOrgGroupOk returns a tuple with the OrgGroup field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipCreateRelationships) GetOrgGroupOk() (*OrgGroupRelationshipToOne, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgGroup, true
}

// SetOrgGroup sets field value.
func (o *OrgGroupMembershipCreateRelationships) SetOrgGroup(v OrgGroupRelationshipToOne) {
	o.OrgGroup = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupMembershipCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["org_group"] = o.OrgGroup

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupMembershipCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgGroup *OrgGroupRelationshipToOne `json:"org_group"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.OrgGroup == nil {
		return fmt.Errorf("required field org_group missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org_group"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.OrgGroup.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OrgGroup = *all.OrgGroup

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
