// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentJiraTemplateDataRequest Incident Jira template data for a create or update request.
type IncidentJiraTemplateDataRequest struct {
	// Attributes for creating or updating an incident Jira template.
	Attributes IncidentJiraTemplateDataAttributesRequest `json:"attributes"`
	// Relationships for an incident Jira template.
	Relationships *IncidentJiraTemplateRelationships `json:"relationships,omitempty"`
	// Incident Jira template resource type.
	Type IncidentJiraTemplateType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentJiraTemplateDataRequest instantiates a new IncidentJiraTemplateDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentJiraTemplateDataRequest(attributes IncidentJiraTemplateDataAttributesRequest, typeVar IncidentJiraTemplateType) *IncidentJiraTemplateDataRequest {
	this := IncidentJiraTemplateDataRequest{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewIncidentJiraTemplateDataRequestWithDefaults instantiates a new IncidentJiraTemplateDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentJiraTemplateDataRequestWithDefaults() *IncidentJiraTemplateDataRequest {
	this := IncidentJiraTemplateDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentJiraTemplateDataRequest) GetAttributes() IncidentJiraTemplateDataAttributesRequest {
	if o == nil {
		var ret IncidentJiraTemplateDataAttributesRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataRequest) GetAttributesOk() (*IncidentJiraTemplateDataAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentJiraTemplateDataRequest) SetAttributes(v IncidentJiraTemplateDataAttributesRequest) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentJiraTemplateDataRequest) GetRelationships() IncidentJiraTemplateRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentJiraTemplateRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataRequest) GetRelationshipsOk() (*IncidentJiraTemplateRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentJiraTemplateDataRequest) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentJiraTemplateRelationships and assigns it to the Relationships field.
func (o *IncidentJiraTemplateDataRequest) SetRelationships(v IncidentJiraTemplateRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentJiraTemplateDataRequest) GetType() IncidentJiraTemplateType {
	if o == nil {
		var ret IncidentJiraTemplateType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentJiraTemplateDataRequest) GetTypeOk() (*IncidentJiraTemplateType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentJiraTemplateDataRequest) SetType(v IncidentJiraTemplateType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentJiraTemplateDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentJiraTemplateDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentJiraTemplateDataAttributesRequest `json:"attributes"`
		Relationships *IncidentJiraTemplateRelationships         `json:"relationships,omitempty"`
		Type          *IncidentJiraTemplateType                  `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
