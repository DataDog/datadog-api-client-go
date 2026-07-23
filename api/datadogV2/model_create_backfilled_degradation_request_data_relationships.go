// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateBackfilledDegradationRequestDataRelationships The supported relationships for creating a backfilled degradation.
type CreateBackfilledDegradationRequestDataRelationships struct {
	// The template used to create the backfilled degradation.
	Template *CreateBackfilledDegradationRequestDataRelationshipsTemplate `json:"template,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateBackfilledDegradationRequestDataRelationships instantiates a new CreateBackfilledDegradationRequestDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateBackfilledDegradationRequestDataRelationships() *CreateBackfilledDegradationRequestDataRelationships {
	this := CreateBackfilledDegradationRequestDataRelationships{}
	return &this
}

// NewCreateBackfilledDegradationRequestDataRelationshipsWithDefaults instantiates a new CreateBackfilledDegradationRequestDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateBackfilledDegradationRequestDataRelationshipsWithDefaults() *CreateBackfilledDegradationRequestDataRelationships {
	this := CreateBackfilledDegradationRequestDataRelationships{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *CreateBackfilledDegradationRequestDataRelationships) GetTemplate() CreateBackfilledDegradationRequestDataRelationshipsTemplate {
	if o == nil || o.Template == nil {
		var ret CreateBackfilledDegradationRequestDataRelationshipsTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBackfilledDegradationRequestDataRelationships) GetTemplateOk() (*CreateBackfilledDegradationRequestDataRelationshipsTemplate, bool) {
	if o == nil || o.Template == nil {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *CreateBackfilledDegradationRequestDataRelationships) HasTemplate() bool {
	return o != nil && o.Template != nil
}

// SetTemplate gets a reference to the given CreateBackfilledDegradationRequestDataRelationshipsTemplate and assigns it to the Template field.
func (o *CreateBackfilledDegradationRequestDataRelationships) SetTemplate(v CreateBackfilledDegradationRequestDataRelationshipsTemplate) {
	o.Template = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateBackfilledDegradationRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Template != nil {
		toSerialize["template"] = o.Template
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateBackfilledDegradationRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Template *CreateBackfilledDegradationRequestDataRelationshipsTemplate `json:"template,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"template"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Template != nil && all.Template.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Template = all.Template

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
