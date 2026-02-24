// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GenerateCampaignTeamReportsRequestAttributes Attributes for generating team campaign reports.
type GenerateCampaignTeamReportsRequestAttributes struct {
	// List of entity owners and their Slack destinations.
	EntityOwners []EntityOwnerDestination `json:"entity_owners"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGenerateCampaignTeamReportsRequestAttributes instantiates a new GenerateCampaignTeamReportsRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGenerateCampaignTeamReportsRequestAttributes(entityOwners []EntityOwnerDestination) *GenerateCampaignTeamReportsRequestAttributes {
	this := GenerateCampaignTeamReportsRequestAttributes{}
	this.EntityOwners = entityOwners
	return &this
}

// NewGenerateCampaignTeamReportsRequestAttributesWithDefaults instantiates a new GenerateCampaignTeamReportsRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGenerateCampaignTeamReportsRequestAttributesWithDefaults() *GenerateCampaignTeamReportsRequestAttributes {
	this := GenerateCampaignTeamReportsRequestAttributes{}
	return &this
}

// GetEntityOwners returns the EntityOwners field value.
func (o *GenerateCampaignTeamReportsRequestAttributes) GetEntityOwners() []EntityOwnerDestination {
	if o == nil {
		var ret []EntityOwnerDestination
		return ret
	}
	return o.EntityOwners
}

// GetEntityOwnersOk returns a tuple with the EntityOwners field value
// and a boolean to check if the value has been set.
func (o *GenerateCampaignTeamReportsRequestAttributes) GetEntityOwnersOk() (*[]EntityOwnerDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityOwners, true
}

// SetEntityOwners sets field value.
func (o *GenerateCampaignTeamReportsRequestAttributes) SetEntityOwners(v []EntityOwnerDestination) {
	o.EntityOwners = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GenerateCampaignTeamReportsRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["entity_owners"] = o.EntityOwners

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GenerateCampaignTeamReportsRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntityOwners *[]EntityOwnerDestination `json:"entity_owners"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EntityOwners == nil {
		return fmt.Errorf("required field entity_owners missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entity_owners"})
	} else {
		return err
	}
	o.EntityOwners = *all.EntityOwners

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
