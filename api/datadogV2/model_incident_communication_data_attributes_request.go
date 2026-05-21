// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCommunicationDataAttributesRequest Attributes for creating or updating a communication.
type IncidentCommunicationDataAttributesRequest struct {
	// The kind of communication.
	CommunicationType IncidentCommunicationKind `json:"communication_type"`
	// The content of a communication.
	Content IncidentCommunicationContent `json:"content"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCommunicationDataAttributesRequest instantiates a new IncidentCommunicationDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCommunicationDataAttributesRequest(communicationType IncidentCommunicationKind, content IncidentCommunicationContent) *IncidentCommunicationDataAttributesRequest {
	this := IncidentCommunicationDataAttributesRequest{}
	this.CommunicationType = communicationType
	this.Content = content
	return &this
}

// NewIncidentCommunicationDataAttributesRequestWithDefaults instantiates a new IncidentCommunicationDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCommunicationDataAttributesRequestWithDefaults() *IncidentCommunicationDataAttributesRequest {
	this := IncidentCommunicationDataAttributesRequest{}
	return &this
}

// GetCommunicationType returns the CommunicationType field value.
func (o *IncidentCommunicationDataAttributesRequest) GetCommunicationType() IncidentCommunicationKind {
	if o == nil {
		var ret IncidentCommunicationKind
		return ret
	}
	return o.CommunicationType
}

// GetCommunicationTypeOk returns a tuple with the CommunicationType field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationDataAttributesRequest) GetCommunicationTypeOk() (*IncidentCommunicationKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunicationType, true
}

// SetCommunicationType sets field value.
func (o *IncidentCommunicationDataAttributesRequest) SetCommunicationType(v IncidentCommunicationKind) {
	o.CommunicationType = v
}

// GetContent returns the Content field value.
func (o *IncidentCommunicationDataAttributesRequest) GetContent() IncidentCommunicationContent {
	if o == nil {
		var ret IncidentCommunicationContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationDataAttributesRequest) GetContentOk() (*IncidentCommunicationContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentCommunicationDataAttributesRequest) SetContent(v IncidentCommunicationContent) {
	o.Content = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCommunicationDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["communication_type"] = o.CommunicationType
	toSerialize["content"] = o.Content

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCommunicationDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommunicationType *IncidentCommunicationKind    `json:"communication_type"`
		Content           *IncidentCommunicationContent `json:"content"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommunicationType == nil {
		return fmt.Errorf("required field communication_type missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"communication_type", "content"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.CommunicationType.IsValid() {
		hasInvalidField = true
	} else {
		o.CommunicationType = *all.CommunicationType
	}
	if all.Content.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Content = *all.Content

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
