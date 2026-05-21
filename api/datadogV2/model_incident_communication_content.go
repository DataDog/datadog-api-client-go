// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCommunicationContent The content of a communication.
type IncidentCommunicationContent struct {
	// A key used for grouping communications.
	GroupingKey *string `json:"grouping_key,omitempty"`
	// The list of handles the communication is sent to.
	Handles []IncidentCommunicationContentHandle `json:"handles"`
	// The message body of the communication.
	Message string `json:"message"`
	// The status code of the communication.
	Status *int32 `json:"status,omitempty"`
	// The subject line of the communication.
	Subject *string `json:"subject,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCommunicationContent instantiates a new IncidentCommunicationContent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCommunicationContent(handles []IncidentCommunicationContentHandle, message string) *IncidentCommunicationContent {
	this := IncidentCommunicationContent{}
	this.Handles = handles
	this.Message = message
	return &this
}

// NewIncidentCommunicationContentWithDefaults instantiates a new IncidentCommunicationContent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCommunicationContentWithDefaults() *IncidentCommunicationContent {
	this := IncidentCommunicationContent{}
	return &this
}

// GetGroupingKey returns the GroupingKey field value if set, zero value otherwise.
func (o *IncidentCommunicationContent) GetGroupingKey() string {
	if o == nil || o.GroupingKey == nil {
		var ret string
		return ret
	}
	return *o.GroupingKey
}

// GetGroupingKeyOk returns a tuple with the GroupingKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContent) GetGroupingKeyOk() (*string, bool) {
	if o == nil || o.GroupingKey == nil {
		return nil, false
	}
	return o.GroupingKey, true
}

// HasGroupingKey returns a boolean if a field has been set.
func (o *IncidentCommunicationContent) HasGroupingKey() bool {
	return o != nil && o.GroupingKey != nil
}

// SetGroupingKey gets a reference to the given string and assigns it to the GroupingKey field.
func (o *IncidentCommunicationContent) SetGroupingKey(v string) {
	o.GroupingKey = &v
}

// GetHandles returns the Handles field value.
func (o *IncidentCommunicationContent) GetHandles() []IncidentCommunicationContentHandle {
	if o == nil {
		var ret []IncidentCommunicationContentHandle
		return ret
	}
	return o.Handles
}

// GetHandlesOk returns a tuple with the Handles field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContent) GetHandlesOk() (*[]IncidentCommunicationContentHandle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handles, true
}

// SetHandles sets field value.
func (o *IncidentCommunicationContent) SetHandles(v []IncidentCommunicationContentHandle) {
	o.Handles = v
}

// GetMessage returns the Message field value.
func (o *IncidentCommunicationContent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value.
func (o *IncidentCommunicationContent) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IncidentCommunicationContent) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContent) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IncidentCommunicationContent) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *IncidentCommunicationContent) SetStatus(v int32) {
	o.Status = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IncidentCommunicationContent) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationContent) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IncidentCommunicationContent) HasSubject() bool {
	return o != nil && o.Subject != nil
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *IncidentCommunicationContent) SetSubject(v string) {
	o.Subject = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCommunicationContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GroupingKey != nil {
		toSerialize["grouping_key"] = o.GroupingKey
	}
	toSerialize["handles"] = o.Handles
	toSerialize["message"] = o.Message
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCommunicationContent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupingKey *string                               `json:"grouping_key,omitempty"`
		Handles     *[]IncidentCommunicationContentHandle `json:"handles"`
		Message     *string                               `json:"message"`
		Status      *int32                                `json:"status,omitempty"`
		Subject     *string                               `json:"subject,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Handles == nil {
		return fmt.Errorf("required field handles missing")
	}
	if all.Message == nil {
		return fmt.Errorf("required field message missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"grouping_key", "handles", "message", "status", "subject"})
	} else {
		return err
	}
	o.GroupingKey = all.GroupingKey
	o.Handles = *all.Handles
	o.Message = *all.Message
	o.Status = all.Status
	o.Subject = all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
