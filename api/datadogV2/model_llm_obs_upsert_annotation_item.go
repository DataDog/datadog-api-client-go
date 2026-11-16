// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsUpsertAnnotationItem A single annotation to create or update. The annotation is matched by
// `interaction_id` and the requesting user's identity.
type LLMObsUpsertAnnotationItem struct {
	// Caller-supplied name or identifier for the external person or automated system that produced the labels.
	// The value is an opaque string with no required format. Use a stable, recognizable value for the same
	// producer across annotations.
	// The value is stored and returned as provided. It is not a Datadog user ID or login, and it is not
	// resolved against Datadog users.
	// Omit this field or set it to `null` to preserve existing attribution. Set it to an empty string to clear
	// attribution.
	Annotator datadog.NullableString `json:"annotator,omitempty"`
	// ID of the interaction to annotate.
	InteractionId string `json:"interaction_id"`
	// Label values for this annotation. Each entry references a label schema by ID
	// and provides the corresponding value validated against the schema type constraints.
	LabelValues []LLMObsAnnotationLabelValue `json:"label_values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsUpsertAnnotationItem instantiates a new LLMObsUpsertAnnotationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsUpsertAnnotationItem(interactionId string, labelValues []LLMObsAnnotationLabelValue) *LLMObsUpsertAnnotationItem {
	this := LLMObsUpsertAnnotationItem{}
	this.InteractionId = interactionId
	this.LabelValues = labelValues
	return &this
}

// NewLLMObsUpsertAnnotationItemWithDefaults instantiates a new LLMObsUpsertAnnotationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsUpsertAnnotationItemWithDefaults() *LLMObsUpsertAnnotationItem {
	this := LLMObsUpsertAnnotationItem{}
	return &this
}

// GetAnnotator returns the Annotator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LLMObsUpsertAnnotationItem) GetAnnotator() string {
	if o == nil || o.Annotator.Get() == nil {
		var ret string
		return ret
	}
	return *o.Annotator.Get()
}

// GetAnnotatorOk returns a tuple with the Annotator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsUpsertAnnotationItem) GetAnnotatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Annotator.Get(), o.Annotator.IsSet()
}

// HasAnnotator returns a boolean if a field has been set.
func (o *LLMObsUpsertAnnotationItem) HasAnnotator() bool {
	return o != nil && o.Annotator.IsSet()
}

// SetAnnotator gets a reference to the given datadog.NullableString and assigns it to the Annotator field.
func (o *LLMObsUpsertAnnotationItem) SetAnnotator(v string) {
	o.Annotator.Set(&v)
}

// SetAnnotatorNil sets the value for Annotator to be an explicit nil.
func (o *LLMObsUpsertAnnotationItem) SetAnnotatorNil() {
	o.Annotator.Set(nil)
}

// UnsetAnnotator ensures that no value is present for Annotator, not even an explicit nil.
func (o *LLMObsUpsertAnnotationItem) UnsetAnnotator() {
	o.Annotator.Unset()
}

// GetInteractionId returns the InteractionId field value.
func (o *LLMObsUpsertAnnotationItem) GetInteractionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InteractionId
}

// GetInteractionIdOk returns a tuple with the InteractionId field value
// and a boolean to check if the value has been set.
func (o *LLMObsUpsertAnnotationItem) GetInteractionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionId, true
}

// SetInteractionId sets field value.
func (o *LLMObsUpsertAnnotationItem) SetInteractionId(v string) {
	o.InteractionId = v
}

// GetLabelValues returns the LabelValues field value.
func (o *LLMObsUpsertAnnotationItem) GetLabelValues() []LLMObsAnnotationLabelValue {
	if o == nil {
		var ret []LLMObsAnnotationLabelValue
		return ret
	}
	return o.LabelValues
}

// GetLabelValuesOk returns a tuple with the LabelValues field value
// and a boolean to check if the value has been set.
func (o *LLMObsUpsertAnnotationItem) GetLabelValuesOk() (*[]LLMObsAnnotationLabelValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelValues, true
}

// SetLabelValues sets field value.
func (o *LLMObsUpsertAnnotationItem) SetLabelValues(v []LLMObsAnnotationLabelValue) {
	o.LabelValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsUpsertAnnotationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Annotator.IsSet() {
		toSerialize["annotator"] = o.Annotator.Get()
	}
	toSerialize["interaction_id"] = o.InteractionId
	toSerialize["label_values"] = o.LabelValues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsUpsertAnnotationItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotator     datadog.NullableString        `json:"annotator,omitempty"`
		InteractionId *string                       `json:"interaction_id"`
		LabelValues   *[]LLMObsAnnotationLabelValue `json:"label_values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InteractionId == nil {
		return fmt.Errorf("required field interaction_id missing")
	}
	if all.LabelValues == nil {
		return fmt.Errorf("required field label_values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotator", "interaction_id", "label_values"})
	} else {
		return err
	}
	o.Annotator = all.Annotator
	o.InteractionId = *all.InteractionId
	o.LabelValues = *all.LabelValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
