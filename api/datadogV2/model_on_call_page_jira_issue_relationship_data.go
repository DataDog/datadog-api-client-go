// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallPageJiraIssueRelationshipData Reference to a Jira issue.
type OnCallPageJiraIssueRelationshipData struct {
	// The ID of the Jira issue, in the form `<account_id>:<issue_id>`.
	Id string `json:"id"`
	// Jira issue resource type.
	Type OnCallPageJiraIssueType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallPageJiraIssueRelationshipData instantiates a new OnCallPageJiraIssueRelationshipData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallPageJiraIssueRelationshipData(id string, typeVar OnCallPageJiraIssueType) *OnCallPageJiraIssueRelationshipData {
	this := OnCallPageJiraIssueRelationshipData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOnCallPageJiraIssueRelationshipDataWithDefaults instantiates a new OnCallPageJiraIssueRelationshipData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallPageJiraIssueRelationshipDataWithDefaults() *OnCallPageJiraIssueRelationshipData {
	this := OnCallPageJiraIssueRelationshipData{}
	var typeVar OnCallPageJiraIssueType = ONCALLPAGEJIRAISSUETYPE_JIRA_ISSUES
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *OnCallPageJiraIssueRelationshipData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OnCallPageJiraIssueRelationshipData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OnCallPageJiraIssueRelationshipData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *OnCallPageJiraIssueRelationshipData) GetType() OnCallPageJiraIssueType {
	if o == nil {
		var ret OnCallPageJiraIssueType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnCallPageJiraIssueRelationshipData) GetTypeOk() (*OnCallPageJiraIssueType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OnCallPageJiraIssueRelationshipData) SetType(v OnCallPageJiraIssueType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallPageJiraIssueRelationshipData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallPageJiraIssueRelationshipData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                  `json:"id"`
		Type *OnCallPageJiraIssueType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
