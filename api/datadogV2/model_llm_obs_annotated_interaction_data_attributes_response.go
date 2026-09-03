// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotatedInteractionDataAttributesResponse Attributes containing an annotated interaction and its related events.
type LLMObsAnnotatedInteractionDataAttributesResponse struct {
	// An interaction with its associated annotations.
	AnnotatedInteraction LLMObsAnnotatedInteractionItem `json:"annotated_interaction"`
	// Page of events associated with the annotated interaction.
	Events []LLMObsAnnotatedInteractionEvent `json:"events"`
	// Type of an annotated interaction.
	InteractionType LLMObsAnyInteractionType `json:"interaction_type"`
	// Cursor to retrieve the next page of events. Absent when there are no more events.
	NextCursor *string `json:"next_cursor,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotatedInteractionDataAttributesResponse instantiates a new LLMObsAnnotatedInteractionDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotatedInteractionDataAttributesResponse(annotatedInteraction LLMObsAnnotatedInteractionItem, events []LLMObsAnnotatedInteractionEvent, interactionType LLMObsAnyInteractionType) *LLMObsAnnotatedInteractionDataAttributesResponse {
	this := LLMObsAnnotatedInteractionDataAttributesResponse{}
	this.AnnotatedInteraction = annotatedInteraction
	this.Events = events
	this.InteractionType = interactionType
	return &this
}

// NewLLMObsAnnotatedInteractionDataAttributesResponseWithDefaults instantiates a new LLMObsAnnotatedInteractionDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotatedInteractionDataAttributesResponseWithDefaults() *LLMObsAnnotatedInteractionDataAttributesResponse {
	this := LLMObsAnnotatedInteractionDataAttributesResponse{}
	return &this
}

// GetAnnotatedInteraction returns the AnnotatedInteraction field value.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetAnnotatedInteraction() LLMObsAnnotatedInteractionItem {
	if o == nil {
		var ret LLMObsAnnotatedInteractionItem
		return ret
	}
	return o.AnnotatedInteraction
}

// GetAnnotatedInteractionOk returns a tuple with the AnnotatedInteraction field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetAnnotatedInteractionOk() (*LLMObsAnnotatedInteractionItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnnotatedInteraction, true
}

// SetAnnotatedInteraction sets field value.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) SetAnnotatedInteraction(v LLMObsAnnotatedInteractionItem) {
	o.AnnotatedInteraction = v
}

// GetEvents returns the Events field value.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetEvents() []LLMObsAnnotatedInteractionEvent {
	if o == nil {
		var ret []LLMObsAnnotatedInteractionEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetEventsOk() (*[]LLMObsAnnotatedInteractionEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) SetEvents(v []LLMObsAnnotatedInteractionEvent) {
	o.Events = v
}

// GetInteractionType returns the InteractionType field value.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetInteractionType() LLMObsAnyInteractionType {
	if o == nil {
		var ret LLMObsAnyInteractionType
		return ret
	}
	return o.InteractionType
}

// GetInteractionTypeOk returns a tuple with the InteractionType field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetInteractionTypeOk() (*LLMObsAnyInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionType, true
}

// SetInteractionType sets field value.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) SetInteractionType(v LLMObsAnyInteractionType) {
	o.InteractionType = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) SetNextCursor(v string) {
	o.NextCursor = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotatedInteractionDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotated_interaction"] = o.AnnotatedInteraction
	toSerialize["events"] = o.Events
	toSerialize["interaction_type"] = o.InteractionType
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotatedInteraction *LLMObsAnnotatedInteractionItem    `json:"annotated_interaction"`
		Events               *[]LLMObsAnnotatedInteractionEvent `json:"events"`
		InteractionType      *LLMObsAnyInteractionType          `json:"interaction_type"`
		NextCursor           *string                            `json:"next_cursor,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnnotatedInteraction == nil {
		return fmt.Errorf("required field annotated_interaction missing")
	}
	if all.Events == nil {
		return fmt.Errorf("required field events missing")
	}
	if all.InteractionType == nil {
		return fmt.Errorf("required field interaction_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotated_interaction", "events", "interaction_type", "next_cursor"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AnnotatedInteraction = *all.AnnotatedInteraction
	o.Events = *all.Events
	if !all.InteractionType.IsValid() {
		hasInvalidField = true
	} else {
		o.InteractionType = *all.InteractionType
	}
	o.NextCursor = all.NextCursor

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
