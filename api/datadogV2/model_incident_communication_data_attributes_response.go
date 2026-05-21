// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentCommunicationDataAttributesResponse Attributes of an incident communication response.
type IncidentCommunicationDataAttributesResponse struct {
	// The kind of communication.
	CommunicationType IncidentCommunicationKind `json:"communication_type"`
	// The content of a communication.
	Content IncidentCommunicationContent `json:"content"`
	// Timestamp when the communication was created.
	Created time.Time `json:"created"`
	// The incident identifier.
	IncidentId uuid.UUID `json:"incident_id"`
	// Timestamp when the communication was last modified.
	Modified time.Time `json:"modified"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentCommunicationDataAttributesResponse instantiates a new IncidentCommunicationDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentCommunicationDataAttributesResponse(communicationType IncidentCommunicationKind, content IncidentCommunicationContent, created time.Time, incidentId uuid.UUID, modified time.Time) *IncidentCommunicationDataAttributesResponse {
	this := IncidentCommunicationDataAttributesResponse{}
	this.CommunicationType = communicationType
	this.Content = content
	this.Created = created
	this.IncidentId = incidentId
	this.Modified = modified
	return &this
}

// NewIncidentCommunicationDataAttributesResponseWithDefaults instantiates a new IncidentCommunicationDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentCommunicationDataAttributesResponseWithDefaults() *IncidentCommunicationDataAttributesResponse {
	this := IncidentCommunicationDataAttributesResponse{}
	return &this
}

// GetCommunicationType returns the CommunicationType field value.
func (o *IncidentCommunicationDataAttributesResponse) GetCommunicationType() IncidentCommunicationKind {
	if o == nil {
		var ret IncidentCommunicationKind
		return ret
	}
	return o.CommunicationType
}

// GetCommunicationTypeOk returns a tuple with the CommunicationType field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationDataAttributesResponse) GetCommunicationTypeOk() (*IncidentCommunicationKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommunicationType, true
}

// SetCommunicationType sets field value.
func (o *IncidentCommunicationDataAttributesResponse) SetCommunicationType(v IncidentCommunicationKind) {
	o.CommunicationType = v
}

// GetContent returns the Content field value.
func (o *IncidentCommunicationDataAttributesResponse) GetContent() IncidentCommunicationContent {
	if o == nil {
		var ret IncidentCommunicationContent
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationDataAttributesResponse) GetContentOk() (*IncidentCommunicationContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentCommunicationDataAttributesResponse) SetContent(v IncidentCommunicationContent) {
	o.Content = v
}

// GetCreated returns the Created field value.
func (o *IncidentCommunicationDataAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationDataAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentCommunicationDataAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetIncidentId returns the IncidentId field value.
func (o *IncidentCommunicationDataAttributesResponse) GetIncidentId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.IncidentId
}

// GetIncidentIdOk returns a tuple with the IncidentId field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationDataAttributesResponse) GetIncidentIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncidentId, true
}

// SetIncidentId sets field value.
func (o *IncidentCommunicationDataAttributesResponse) SetIncidentId(v uuid.UUID) {
	o.IncidentId = v
}

// GetModified returns the Modified field value.
func (o *IncidentCommunicationDataAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentCommunicationDataAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentCommunicationDataAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentCommunicationDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["communication_type"] = o.CommunicationType
	toSerialize["content"] = o.Content
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["incident_id"] = o.IncidentId
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentCommunicationDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommunicationType *IncidentCommunicationKind    `json:"communication_type"`
		Content           *IncidentCommunicationContent `json:"content"`
		Created           *time.Time                    `json:"created"`
		IncidentId        *uuid.UUID                    `json:"incident_id"`
		Modified          *time.Time                    `json:"modified"`
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
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.IncidentId == nil {
		return fmt.Errorf("required field incident_id missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"communication_type", "content", "created", "incident_id", "modified"})
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
	o.Created = *all.Created
	o.IncidentId = *all.IncidentId
	o.Modified = *all.Modified

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
