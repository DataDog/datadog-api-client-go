// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentStatusPageNoticeCreateDataAttributes Attributes for publishing a status page notice.
type IncidentStatusPageNoticeCreateDataAttributes struct {
	// Map of component identifiers to their status.
	Components map[string]string `json:"components,omitempty"`
	// The message body of the notice.
	Message *string `json:"message,omitempty"`
	// The status of the notice.
	Status *string `json:"status,omitempty"`
	// The title of the notice.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentStatusPageNoticeCreateDataAttributes instantiates a new IncidentStatusPageNoticeCreateDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentStatusPageNoticeCreateDataAttributes() *IncidentStatusPageNoticeCreateDataAttributes {
	this := IncidentStatusPageNoticeCreateDataAttributes{}
	return &this
}

// NewIncidentStatusPageNoticeCreateDataAttributesWithDefaults instantiates a new IncidentStatusPageNoticeCreateDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentStatusPageNoticeCreateDataAttributesWithDefaults() *IncidentStatusPageNoticeCreateDataAttributes {
	this := IncidentStatusPageNoticeCreateDataAttributes{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetComponents() map[string]string {
	if o == nil || o.Components == nil {
		var ret map[string]string
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetComponentsOk() (*map[string]string, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given map[string]string and assigns it to the Components field.
func (o *IncidentStatusPageNoticeCreateDataAttributes) SetComponents(v map[string]string) {
	o.Components = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IncidentStatusPageNoticeCreateDataAttributes) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IncidentStatusPageNoticeCreateDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IncidentStatusPageNoticeCreateDataAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IncidentStatusPageNoticeCreateDataAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentStatusPageNoticeCreateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentStatusPageNoticeCreateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Components map[string]string `json:"components,omitempty"`
		Message    *string           `json:"message,omitempty"`
		Status     *string           `json:"status,omitempty"`
		Title      *string           `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"components", "message", "status", "title"})
	} else {
		return err
	}
	o.Components = all.Components
	o.Message = all.Message
	o.Status = all.Status
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
