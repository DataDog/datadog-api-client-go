// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetAccessScope A fleet access scope. Use `fleet-scope:<id>` as the restriction policy resource identifier.
type FleetAccessScope struct {
	// Fleet access scope attributes and audit information.
	Attributes FleetAccessScopeResponseAttributes `json:"attributes"`
	// Unique identifier of the scope, also used in the restriction policy resource identifier `fleet-scope:<id>`.
	Id string `json:"id"`
	// The JSON:API resource type for a fleet access scope.
	Type FleetAccessScopeResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetAccessScope instantiates a new FleetAccessScope object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetAccessScope(attributes FleetAccessScopeResponseAttributes, id string, typeVar FleetAccessScopeResourceType) *FleetAccessScope {
	this := FleetAccessScope{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewFleetAccessScopeWithDefaults instantiates a new FleetAccessScope object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetAccessScopeWithDefaults() *FleetAccessScope {
	this := FleetAccessScope{}
	var typeVar FleetAccessScopeResourceType = FLEETACCESSSCOPERESOURCETYPE_ACCESS_SCOPE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *FleetAccessScope) GetAttributes() FleetAccessScopeResponseAttributes {
	if o == nil {
		var ret FleetAccessScopeResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *FleetAccessScope) GetAttributesOk() (*FleetAccessScopeResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *FleetAccessScope) SetAttributes(v FleetAccessScopeResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *FleetAccessScope) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FleetAccessScope) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *FleetAccessScope) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *FleetAccessScope) GetType() FleetAccessScopeResourceType {
	if o == nil {
		var ret FleetAccessScopeResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FleetAccessScope) GetTypeOk() (*FleetAccessScopeResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FleetAccessScope) SetType(v FleetAccessScopeResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetAccessScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetAccessScope) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *FleetAccessScopeResponseAttributes `json:"attributes"`
		Id         *string                             `json:"id"`
		Type       *FleetAccessScopeResourceType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
