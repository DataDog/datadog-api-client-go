// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionMatch
type ActionMatch struct {
	// The fully qualified name of the action.
	ActionFqn string `json:"actionFqn"`
	// The description of the action.
	Description string `json:"description"`
	// The relevance score of the match.
	Score float64 `json:"score"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionMatch instantiates a new ActionMatch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionMatch(actionFqn string, description string, score float64) *ActionMatch {
	this := ActionMatch{}
	this.ActionFqn = actionFqn
	this.Description = description
	this.Score = score
	return &this
}

// NewActionMatchWithDefaults instantiates a new ActionMatch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionMatchWithDefaults() *ActionMatch {
	this := ActionMatch{}
	return &this
}

// GetActionFqn returns the ActionFqn field value.
func (o *ActionMatch) GetActionFqn() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActionFqn
}

// GetActionFqnOk returns a tuple with the ActionFqn field value
// and a boolean to check if the value has been set.
func (o *ActionMatch) GetActionFqnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionFqn, true
}

// SetActionFqn sets field value.
func (o *ActionMatch) SetActionFqn(v string) {
	o.ActionFqn = v
}

// GetDescription returns the Description field value.
func (o *ActionMatch) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ActionMatch) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ActionMatch) SetDescription(v string) {
	o.Description = v
}

// GetScore returns the Score field value.
func (o *ActionMatch) GetScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *ActionMatch) GetScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value.
func (o *ActionMatch) SetScore(v float64) {
	o.Score = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionMatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["actionFqn"] = o.ActionFqn
	toSerialize["description"] = o.Description
	toSerialize["score"] = o.Score

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionMatch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ActionFqn   *string  `json:"actionFqn"`
		Description *string  `json:"description"`
		Score       *float64 `json:"score"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ActionFqn == nil {
		return fmt.Errorf("required field actionFqn missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Score == nil {
		return fmt.Errorf("required field score missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"actionFqn", "description", "score"})
	} else {
		return err
	}
	o.ActionFqn = *all.ActionFqn
	o.Description = *all.Description
	o.Score = *all.Score

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
