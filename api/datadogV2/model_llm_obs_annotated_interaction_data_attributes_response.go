// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotatedInteractionDataAttributesResponse Attributes containing the annotated interaction.
type LLMObsAnnotatedInteractionDataAttributesResponse struct {
	// An interaction with its associated annotations.
	AnnotatedInteraction LLMObsAnnotatedInteractionItem `json:"annotated_interaction"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotatedInteractionDataAttributesResponse instantiates a new LLMObsAnnotatedInteractionDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotatedInteractionDataAttributesResponse(annotatedInteraction LLMObsAnnotatedInteractionItem) *LLMObsAnnotatedInteractionDataAttributesResponse {
	this := LLMObsAnnotatedInteractionDataAttributesResponse{}
	this.AnnotatedInteraction = annotatedInteraction
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

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotatedInteractionDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotated_interaction"] = o.AnnotatedInteraction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotatedInteractionDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnnotatedInteraction *LLMObsAnnotatedInteractionItem `json:"annotated_interaction"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnnotatedInteraction == nil {
		return fmt.Errorf("required field annotated_interaction missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotated_interaction"})
	} else {
		return err
	}
	o.AnnotatedInteraction = *all.AnnotatedInteraction

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
