// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityAgentInvestigationKeyEntities The entities the investigation found, by role.
type SecurityAgentInvestigationKeyEntities struct {
	// Entities that performed the activity.
	Actors []SecurityAgentInvestigationEntity `json:"actors"`
	// Entities the activity was aimed at.
	Targets []SecurityAgentInvestigationEntity `json:"targets"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityAgentInvestigationKeyEntities instantiates a new SecurityAgentInvestigationKeyEntities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityAgentInvestigationKeyEntities(actors []SecurityAgentInvestigationEntity, targets []SecurityAgentInvestigationEntity) *SecurityAgentInvestigationKeyEntities {
	this := SecurityAgentInvestigationKeyEntities{}
	this.Actors = actors
	this.Targets = targets
	return &this
}

// NewSecurityAgentInvestigationKeyEntitiesWithDefaults instantiates a new SecurityAgentInvestigationKeyEntities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityAgentInvestigationKeyEntitiesWithDefaults() *SecurityAgentInvestigationKeyEntities {
	this := SecurityAgentInvestigationKeyEntities{}
	return &this
}

// GetActors returns the Actors field value.
func (o *SecurityAgentInvestigationKeyEntities) GetActors() []SecurityAgentInvestigationEntity {
	if o == nil {
		var ret []SecurityAgentInvestigationEntity
		return ret
	}
	return o.Actors
}

// GetActorsOk returns a tuple with the Actors field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationKeyEntities) GetActorsOk() (*[]SecurityAgentInvestigationEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actors, true
}

// SetActors sets field value.
func (o *SecurityAgentInvestigationKeyEntities) SetActors(v []SecurityAgentInvestigationEntity) {
	o.Actors = v
}

// GetTargets returns the Targets field value.
func (o *SecurityAgentInvestigationKeyEntities) GetTargets() []SecurityAgentInvestigationEntity {
	if o == nil {
		var ret []SecurityAgentInvestigationEntity
		return ret
	}
	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *SecurityAgentInvestigationKeyEntities) GetTargetsOk() (*[]SecurityAgentInvestigationEntity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Targets, true
}

// SetTargets sets field value.
func (o *SecurityAgentInvestigationKeyEntities) SetTargets(v []SecurityAgentInvestigationEntity) {
	o.Targets = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityAgentInvestigationKeyEntities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actors"] = o.Actors
	toSerialize["targets"] = o.Targets

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityAgentInvestigationKeyEntities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Actors  *[]SecurityAgentInvestigationEntity `json:"actors"`
		Targets *[]SecurityAgentInvestigationEntity `json:"targets"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Actors == nil {
		return fmt.Errorf("required field actors missing")
	}
	if all.Targets == nil {
		return fmt.Errorf("required field targets missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actors", "targets"})
	} else {
		return err
	}
	o.Actors = *all.Actors
	o.Targets = *all.Targets

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
