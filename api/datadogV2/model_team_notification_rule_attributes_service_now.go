// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamNotificationRuleAttributesServiceNow ServiceNow notification settings for the team.
type TeamNotificationRuleAttributesServiceNow struct {
	// ServiceNow template handle names to use for notifications.
	Templates []string `json:"templates,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamNotificationRuleAttributesServiceNow instantiates a new TeamNotificationRuleAttributesServiceNow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamNotificationRuleAttributesServiceNow() *TeamNotificationRuleAttributesServiceNow {
	this := TeamNotificationRuleAttributesServiceNow{}
	return &this
}

// NewTeamNotificationRuleAttributesServiceNowWithDefaults instantiates a new TeamNotificationRuleAttributesServiceNow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamNotificationRuleAttributesServiceNowWithDefaults() *TeamNotificationRuleAttributesServiceNow {
	this := TeamNotificationRuleAttributesServiceNow{}
	return &this
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributesServiceNow) GetTemplates() []string {
	if o == nil || o.Templates == nil {
		var ret []string
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributesServiceNow) GetTemplatesOk() (*[]string, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return &o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributesServiceNow) HasTemplates() bool {
	return o != nil && o.Templates != nil
}

// SetTemplates gets a reference to the given []string and assigns it to the Templates field.
func (o *TeamNotificationRuleAttributesServiceNow) SetTemplates(v []string) {
	o.Templates = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamNotificationRuleAttributesServiceNow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamNotificationRuleAttributesServiceNow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Templates []string `json:"templates,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"templates"})
	} else {
		return err
	}
	o.Templates = all.Templates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
