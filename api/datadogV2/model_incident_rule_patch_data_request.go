// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentRulePatchDataRequest Incident rule data in a patch request.
type IncidentRulePatchDataRequest struct {
	// Attributes for patching an incident rule. All fields are optional.
	Attributes *IncidentRulePatchDataAttributesRequest `json:"attributes,omitempty"`
	// The rule identifier.
	Id uuid.UUID `json:"id"`
	// Incident rule resource type.
	Type IncidentRuleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentRulePatchDataRequest instantiates a new IncidentRulePatchDataRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentRulePatchDataRequest(id uuid.UUID, typeVar IncidentRuleType) *IncidentRulePatchDataRequest {
	this := IncidentRulePatchDataRequest{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentRulePatchDataRequestWithDefaults instantiates a new IncidentRulePatchDataRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentRulePatchDataRequestWithDefaults() *IncidentRulePatchDataRequest {
	this := IncidentRulePatchDataRequest{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentRulePatchDataRequest) GetAttributes() IncidentRulePatchDataAttributesRequest {
	if o == nil || o.Attributes == nil {
		var ret IncidentRulePatchDataAttributesRequest
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataRequest) GetAttributesOk() (*IncidentRulePatchDataAttributesRequest, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentRulePatchDataRequest) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentRulePatchDataAttributesRequest and assigns it to the Attributes field.
func (o *IncidentRulePatchDataRequest) SetAttributes(v IncidentRulePatchDataAttributesRequest) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *IncidentRulePatchDataRequest) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataRequest) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentRulePatchDataRequest) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *IncidentRulePatchDataRequest) GetType() IncidentRuleType {
	if o == nil {
		var ret IncidentRuleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentRulePatchDataRequest) GetTypeOk() (*IncidentRuleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentRulePatchDataRequest) SetType(v IncidentRuleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentRulePatchDataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentRulePatchDataRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *IncidentRulePatchDataAttributesRequest `json:"attributes,omitempty"`
		Id         *uuid.UUID                              `json:"id"`
		Type       *IncidentRuleType                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
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
