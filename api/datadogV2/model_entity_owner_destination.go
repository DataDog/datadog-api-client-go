// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityOwnerDestination Entity owner and Slack destination for team reports.
type EntityOwnerDestination struct {
	// Slack routing options for report delivery.
	Slack SlackRoutingOptions `json:"slack"`
	// Team UUID.
	TeamId string `json:"team_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityOwnerDestination instantiates a new EntityOwnerDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityOwnerDestination(slack SlackRoutingOptions, teamId string) *EntityOwnerDestination {
	this := EntityOwnerDestination{}
	this.Slack = slack
	this.TeamId = teamId
	return &this
}

// NewEntityOwnerDestinationWithDefaults instantiates a new EntityOwnerDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityOwnerDestinationWithDefaults() *EntityOwnerDestination {
	this := EntityOwnerDestination{}
	return &this
}

// GetSlack returns the Slack field value.
func (o *EntityOwnerDestination) GetSlack() SlackRoutingOptions {
	if o == nil {
		var ret SlackRoutingOptions
		return ret
	}
	return o.Slack
}

// GetSlackOk returns a tuple with the Slack field value
// and a boolean to check if the value has been set.
func (o *EntityOwnerDestination) GetSlackOk() (*SlackRoutingOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slack, true
}

// SetSlack sets field value.
func (o *EntityOwnerDestination) SetSlack(v SlackRoutingOptions) {
	o.Slack = v
}

// GetTeamId returns the TeamId field value.
func (o *EntityOwnerDestination) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *EntityOwnerDestination) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value.
func (o *EntityOwnerDestination) SetTeamId(v string) {
	o.TeamId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityOwnerDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["slack"] = o.Slack
	toSerialize["team_id"] = o.TeamId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityOwnerDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Slack  *SlackRoutingOptions `json:"slack"`
		TeamId *string              `json:"team_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Slack == nil {
		return fmt.Errorf("required field slack missing")
	}
	if all.TeamId == nil {
		return fmt.Errorf("required field team_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"slack", "team_id"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Slack.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Slack = *all.Slack
	o.TeamId = *all.TeamId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
