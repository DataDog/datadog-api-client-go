// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Entity An entity participating in the dependency upgrade workflow.
type Entity struct {
	// The kind of the entity (for example, service or package).
	EntityKind *string `json:"entity_kind,omitempty"`
	// The name of the entity.
	EntityName string `json:"entity_name"`
	// The namespace of the entity.
	EntityNamespace *string `json:"entity_namespace,omitempty"`
	// The team that owns the entity.
	EntityTeam *string `json:"entity_team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntity instantiates a new Entity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntity(entityName string) *Entity {
	this := Entity{}
	this.EntityName = entityName
	return &this
}

// NewEntityWithDefaults instantiates a new Entity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityWithDefaults() *Entity {
	this := Entity{}
	return &this
}

// GetEntityKind returns the EntityKind field value if set, zero value otherwise.
func (o *Entity) GetEntityKind() string {
	if o == nil || o.EntityKind == nil {
		var ret string
		return ret
	}
	return *o.EntityKind
}

// GetEntityKindOk returns a tuple with the EntityKind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entity) GetEntityKindOk() (*string, bool) {
	if o == nil || o.EntityKind == nil {
		return nil, false
	}
	return o.EntityKind, true
}

// HasEntityKind returns a boolean if a field has been set.
func (o *Entity) HasEntityKind() bool {
	return o != nil && o.EntityKind != nil
}

// SetEntityKind gets a reference to the given string and assigns it to the EntityKind field.
func (o *Entity) SetEntityKind(v string) {
	o.EntityKind = &v
}

// GetEntityName returns the EntityName field value.
func (o *Entity) GetEntityName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value
// and a boolean to check if the value has been set.
func (o *Entity) GetEntityNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityName, true
}

// SetEntityName sets field value.
func (o *Entity) SetEntityName(v string) {
	o.EntityName = v
}

// GetEntityNamespace returns the EntityNamespace field value if set, zero value otherwise.
func (o *Entity) GetEntityNamespace() string {
	if o == nil || o.EntityNamespace == nil {
		var ret string
		return ret
	}
	return *o.EntityNamespace
}

// GetEntityNamespaceOk returns a tuple with the EntityNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entity) GetEntityNamespaceOk() (*string, bool) {
	if o == nil || o.EntityNamespace == nil {
		return nil, false
	}
	return o.EntityNamespace, true
}

// HasEntityNamespace returns a boolean if a field has been set.
func (o *Entity) HasEntityNamespace() bool {
	return o != nil && o.EntityNamespace != nil
}

// SetEntityNamespace gets a reference to the given string and assigns it to the EntityNamespace field.
func (o *Entity) SetEntityNamespace(v string) {
	o.EntityNamespace = &v
}

// GetEntityTeam returns the EntityTeam field value if set, zero value otherwise.
func (o *Entity) GetEntityTeam() string {
	if o == nil || o.EntityTeam == nil {
		var ret string
		return ret
	}
	return *o.EntityTeam
}

// GetEntityTeamOk returns a tuple with the EntityTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entity) GetEntityTeamOk() (*string, bool) {
	if o == nil || o.EntityTeam == nil {
		return nil, false
	}
	return o.EntityTeam, true
}

// HasEntityTeam returns a boolean if a field has been set.
func (o *Entity) HasEntityTeam() bool {
	return o != nil && o.EntityTeam != nil
}

// SetEntityTeam gets a reference to the given string and assigns it to the EntityTeam field.
func (o *Entity) SetEntityTeam(v string) {
	o.EntityTeam = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Entity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EntityKind != nil {
		toSerialize["entity_kind"] = o.EntityKind
	}
	toSerialize["entity_name"] = o.EntityName
	if o.EntityNamespace != nil {
		toSerialize["entity_namespace"] = o.EntityNamespace
	}
	if o.EntityTeam != nil {
		toSerialize["entity_team"] = o.EntityTeam
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Entity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EntityKind      *string `json:"entity_kind,omitempty"`
		EntityName      *string `json:"entity_name"`
		EntityNamespace *string `json:"entity_namespace,omitempty"`
		EntityTeam      *string `json:"entity_team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.EntityName == nil {
		return fmt.Errorf("required field entity_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"entity_kind", "entity_name", "entity_namespace", "entity_team"})
	} else {
		return err
	}
	o.EntityKind = all.EntityKind
	o.EntityName = *all.EntityName
	o.EntityNamespace = all.EntityNamespace
	o.EntityTeam = all.EntityTeam

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
