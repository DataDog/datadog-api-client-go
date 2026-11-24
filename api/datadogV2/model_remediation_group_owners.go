// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RemediationGroupOwners Who should fix the group: the code owners of the group's findings, from their `@git.codeowners` attribute,
// and their Datadog teams, from their `team` tag. A group with no owners has empty `items` and a `total` of 0.
type RemediationGroupOwners struct {
	// Owners of the group's open findings. `items` lists at most 5 owners, the owners of the most findings first,
	// then by name. `total` counts every distinct owner, so the list is cut when `total` is larger than the number
	// of items.
	Codeowners RemediationGroupOwnerList `json:"codeowners"`
	// Owners of the group's open findings. `items` lists at most 5 owners, the owners of the most findings first,
	// then by name. `total` counts every distinct owner, so the list is cut when `total` is larger than the number
	// of items.
	Teams RemediationGroupOwnerList `json:"teams"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediationGroupOwners instantiates a new RemediationGroupOwners object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediationGroupOwners(codeowners RemediationGroupOwnerList, teams RemediationGroupOwnerList) *RemediationGroupOwners {
	this := RemediationGroupOwners{}
	this.Codeowners = codeowners
	this.Teams = teams
	return &this
}

// NewRemediationGroupOwnersWithDefaults instantiates a new RemediationGroupOwners object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationGroupOwnersWithDefaults() *RemediationGroupOwners {
	this := RemediationGroupOwners{}
	return &this
}

// GetCodeowners returns the Codeowners field value.
func (o *RemediationGroupOwners) GetCodeowners() RemediationGroupOwnerList {
	if o == nil {
		var ret RemediationGroupOwnerList
		return ret
	}
	return o.Codeowners
}

// GetCodeownersOk returns a tuple with the Codeowners field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupOwners) GetCodeownersOk() (*RemediationGroupOwnerList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codeowners, true
}

// SetCodeowners sets field value.
func (o *RemediationGroupOwners) SetCodeowners(v RemediationGroupOwnerList) {
	o.Codeowners = v
}

// GetTeams returns the Teams field value.
func (o *RemediationGroupOwners) GetTeams() RemediationGroupOwnerList {
	if o == nil {
		var ret RemediationGroupOwnerList
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *RemediationGroupOwners) GetTeamsOk() (*RemediationGroupOwnerList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Teams, true
}

// SetTeams sets field value.
func (o *RemediationGroupOwners) SetTeams(v RemediationGroupOwnerList) {
	o.Teams = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RemediationGroupOwners) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["codeowners"] = o.Codeowners
	toSerialize["teams"] = o.Teams

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RemediationGroupOwners) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Codeowners *RemediationGroupOwnerList `json:"codeowners"`
		Teams      *RemediationGroupOwnerList `json:"teams"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Codeowners == nil {
		return fmt.Errorf("required field codeowners missing")
	}
	if all.Teams == nil {
		return fmt.Errorf("required field teams missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"codeowners", "teams"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Codeowners.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Codeowners = *all.Codeowners
	if all.Teams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Teams = *all.Teams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
