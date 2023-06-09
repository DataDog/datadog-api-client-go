// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// DowntimeAttributeMonitorIdentifierTags Object of the monitor tags.
type DowntimeAttributeMonitorIdentifierTags struct {
	// A list of monitor tags. For example, tags that are applied directly to monitors,
	// not tags that are used in monitor queries (which are filtered by the scope parameter), to which the downtime applies.
	// The resulting downtime applies to monitors that match **all** provided monitor tags.
	MonitorTags []string `json:"monitor_tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewDowntimeAttributeMonitorIdentifierTags instantiates a new DowntimeAttributeMonitorIdentifierTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDowntimeAttributeMonitorIdentifierTags(monitorTags []string) *DowntimeAttributeMonitorIdentifierTags {
	this := DowntimeAttributeMonitorIdentifierTags{}
	this.MonitorTags = monitorTags
	return &this
}

// NewDowntimeAttributeMonitorIdentifierTagsWithDefaults instantiates a new DowntimeAttributeMonitorIdentifierTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDowntimeAttributeMonitorIdentifierTagsWithDefaults() *DowntimeAttributeMonitorIdentifierTags {
	this := DowntimeAttributeMonitorIdentifierTags{}
	return &this
}

// GetMonitorTags returns the MonitorTags field value.
func (o *DowntimeAttributeMonitorIdentifierTags) GetMonitorTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MonitorTags
}

// GetMonitorTagsOk returns a tuple with the MonitorTags field value
// and a boolean to check if the value has been set.
func (o *DowntimeAttributeMonitorIdentifierTags) GetMonitorTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorTags, true
}

// SetMonitorTags sets field value.
func (o *DowntimeAttributeMonitorIdentifierTags) SetMonitorTags(v []string) {
	o.MonitorTags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DowntimeAttributeMonitorIdentifierTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["monitor_tags"] = o.MonitorTags
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DowntimeAttributeMonitorIdentifierTags) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		MonitorTags *[]string `json:"monitor_tags"`
	}{}
	all := struct {
		MonitorTags []string `json:"monitor_tags"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.MonitorTags == nil {
		return fmt.Errorf("required field monitor_tags missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.MonitorTags = all.MonitorTags

	return nil
}
