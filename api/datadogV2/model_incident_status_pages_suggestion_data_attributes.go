// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatusPagesSuggestionDataAttributes Attributes of a status pages suggestion.
type IncidentStatusPagesSuggestionDataAttributes struct {
	// The suggested status for the status page.
	Status string `json:"status"`
	// The suggested update text for the status page notice.
	UpdateText string `json:"update_text"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatusPagesSuggestionDataAttributes instantiates a new IncidentStatusPagesSuggestionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatusPagesSuggestionDataAttributes(status string, updateText string) *IncidentStatusPagesSuggestionDataAttributes {
	this := IncidentStatusPagesSuggestionDataAttributes{}
	this.Status = status
	this.UpdateText = updateText
	return &this
}

// NewIncidentStatusPagesSuggestionDataAttributesWithDefaults instantiates a new IncidentStatusPagesSuggestionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatusPagesSuggestionDataAttributesWithDefaults() *IncidentStatusPagesSuggestionDataAttributes {
	this := IncidentStatusPagesSuggestionDataAttributes{}
	return &this
}

// GetStatus returns the Status field value.
func (o *IncidentStatusPagesSuggestionDataAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IncidentStatusPagesSuggestionDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *IncidentStatusPagesSuggestionDataAttributes) SetStatus(v string) {
	o.Status = v
}

// GetUpdateText returns the UpdateText field value.
func (o *IncidentStatusPagesSuggestionDataAttributes) GetUpdateText() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdateText
}

// GetUpdateTextOk returns a tuple with the UpdateText field value
// and a boolean to check if the value has been set.
func (o *IncidentStatusPagesSuggestionDataAttributes) GetUpdateTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateText, true
}

// SetUpdateText sets field value.
func (o *IncidentStatusPagesSuggestionDataAttributes) SetUpdateText(v string) {
	o.UpdateText = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatusPagesSuggestionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status
	toSerialize["update_text"] = o.UpdateText

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatusPagesSuggestionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status     *string `json:"status"`
		UpdateText *string `json:"update_text"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.UpdateText == nil {
		return fmt.Errorf("required field update_text missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status", "update_text"})
	} else {
		return err
	}
	o.Status = *all.Status
	o.UpdateText = *all.UpdateText

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
