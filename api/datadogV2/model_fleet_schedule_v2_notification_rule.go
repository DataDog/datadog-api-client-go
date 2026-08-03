// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetScheduleV2NotificationRule Notification configuration attached to a schedule.
//
// Included when available. If the notification rule cannot be retrieved, this field is
// omitted and the schedule is still returned. If the notification rule is retrieved but its
// handles cannot be resolved, it is still included with an empty `handles` array.
type FleetScheduleV2NotificationRule struct {
	// Notification handles (for example, Slack channels or PagerDuty integrations).
	Handles []string `json:"handles,omitempty"`
	// Tags associated with the notification rule.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetScheduleV2NotificationRule instantiates a new FleetScheduleV2NotificationRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetScheduleV2NotificationRule() *FleetScheduleV2NotificationRule {
	this := FleetScheduleV2NotificationRule{}
	return &this
}

// NewFleetScheduleV2NotificationRuleWithDefaults instantiates a new FleetScheduleV2NotificationRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetScheduleV2NotificationRuleWithDefaults() *FleetScheduleV2NotificationRule {
	this := FleetScheduleV2NotificationRule{}
	return &this
}

// GetHandles returns the Handles field value if set, zero value otherwise.
func (o *FleetScheduleV2NotificationRule) GetHandles() []string {
	if o == nil || o.Handles == nil {
		var ret []string
		return ret
	}
	return o.Handles
}

// GetHandlesOk returns a tuple with the Handles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2NotificationRule) GetHandlesOk() (*[]string, bool) {
	if o == nil || o.Handles == nil {
		return nil, false
	}
	return &o.Handles, true
}

// HasHandles returns a boolean if a field has been set.
func (o *FleetScheduleV2NotificationRule) HasHandles() bool {
	return o != nil && o.Handles != nil
}

// SetHandles gets a reference to the given []string and assigns it to the Handles field.
func (o *FleetScheduleV2NotificationRule) SetHandles(v []string) {
	o.Handles = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FleetScheduleV2NotificationRule) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetScheduleV2NotificationRule) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FleetScheduleV2NotificationRule) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *FleetScheduleV2NotificationRule) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetScheduleV2NotificationRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Handles != nil {
		toSerialize["handles"] = o.Handles
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetScheduleV2NotificationRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Handles []string `json:"handles,omitempty"`
		Tags    []string `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"handles", "tags"})
	} else {
		return err
	}
	o.Handles = all.Handles
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
