// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDeletedPromptDataAttributes Attributes confirming that an LLM Observability prompt was deleted.
type LLMObsDeletedPromptDataAttributes struct {
	// Timestamp when the prompt was deleted.
	DeletedAt time.Time `json:"deleted_at"`
	// Customer-provided identifier of the deleted prompt.
	PromptId string `json:"prompt_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDeletedPromptDataAttributes instantiates a new LLMObsDeletedPromptDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDeletedPromptDataAttributes(deletedAt time.Time, promptId string) *LLMObsDeletedPromptDataAttributes {
	this := LLMObsDeletedPromptDataAttributes{}
	this.DeletedAt = deletedAt
	this.PromptId = promptId
	return &this
}

// NewLLMObsDeletedPromptDataAttributesWithDefaults instantiates a new LLMObsDeletedPromptDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDeletedPromptDataAttributesWithDefaults() *LLMObsDeletedPromptDataAttributes {
	this := LLMObsDeletedPromptDataAttributes{}
	return &this
}

// GetDeletedAt returns the DeletedAt field value.
func (o *LLMObsDeletedPromptDataAttributes) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeletedPromptDataAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value.
func (o *LLMObsDeletedPromptDataAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt = v
}

// GetPromptId returns the PromptId field value.
func (o *LLMObsDeletedPromptDataAttributes) GetPromptId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PromptId
}

// GetPromptIdOk returns a tuple with the PromptId field value
// and a boolean to check if the value has been set.
func (o *LLMObsDeletedPromptDataAttributes) GetPromptIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptId, true
}

// SetPromptId sets field value.
func (o *LLMObsDeletedPromptDataAttributes) SetPromptId(v string) {
	o.PromptId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDeletedPromptDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeletedAt.Nanosecond() == 0 {
		toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["prompt_id"] = o.PromptId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDeletedPromptDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeletedAt *time.Time `json:"deleted_at"`
		PromptId  *string    `json:"prompt_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DeletedAt == nil {
		return fmt.Errorf("required field deleted_at missing")
	}
	if all.PromptId == nil {
		return fmt.Errorf("required field prompt_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deleted_at", "prompt_id"})
	} else {
		return err
	}
	o.DeletedAt = *all.DeletedAt
	o.PromptId = *all.PromptId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
