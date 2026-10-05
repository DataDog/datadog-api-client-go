// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupMembershipCreateAttributes Attributes for adding organizations to an org group.
type OrgGroupMembershipCreateAttributes struct {
	// List of organizations to add. Between 1 and 100 per request. Each `org_uuid` and `org_site` pair must be unique.
	Orgs []GlobalOrgIdentifier `json:"orgs"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupMembershipCreateAttributes instantiates a new OrgGroupMembershipCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupMembershipCreateAttributes(orgs []GlobalOrgIdentifier) *OrgGroupMembershipCreateAttributes {
	this := OrgGroupMembershipCreateAttributes{}
	this.Orgs = orgs
	return &this
}

// NewOrgGroupMembershipCreateAttributesWithDefaults instantiates a new OrgGroupMembershipCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupMembershipCreateAttributesWithDefaults() *OrgGroupMembershipCreateAttributes {
	this := OrgGroupMembershipCreateAttributes{}
	return &this
}

// GetOrgs returns the Orgs field value.
func (o *OrgGroupMembershipCreateAttributes) GetOrgs() []GlobalOrgIdentifier {
	if o == nil {
		var ret []GlobalOrgIdentifier
		return ret
	}
	return o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value
// and a boolean to check if the value has been set.
func (o *OrgGroupMembershipCreateAttributes) GetOrgsOk() (*[]GlobalOrgIdentifier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Orgs, true
}

// SetOrgs sets field value.
func (o *OrgGroupMembershipCreateAttributes) SetOrgs(v []GlobalOrgIdentifier) {
	o.Orgs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupMembershipCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["orgs"] = o.Orgs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupMembershipCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Orgs *[]GlobalOrgIdentifier `json:"orgs"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Orgs == nil {
		return fmt.Errorf("required field orgs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"orgs"})
	} else {
		return err
	}
	o.Orgs = *all.Orgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
