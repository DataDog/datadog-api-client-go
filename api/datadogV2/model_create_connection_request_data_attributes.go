// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateConnectionRequestDataAttributes Attributes defining the data source connection, including join configuration and custom fields.
type CreateConnectionRequestDataAttributes struct {
	// Type-specific configuration of the connection. Only the field matching the connection type is set.
	Details *ListConnectionsResponseDataAttributesConnectionsItemsDetails `json:"details,omitempty"`
	// List of custom attribute fields to import from the data source.
	Fields []CreateConnectionRequestDataAttributesFieldsItems `json:"fields,omitempty"`
	// The attribute in the data source used to join records with the entity.
	JoinAttribute string `json:"join_attribute"`
	// The type of join key used to link the data source to the entity (for example, email or user_id).
	JoinType string `json:"join_type"`
	// Additional key-value metadata associated with the connection.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Display name of the connection.
	Name *string `json:"name,omitempty"`
	// The type of data source connection (for example, ref_table).
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateConnectionRequestDataAttributes instantiates a new CreateConnectionRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateConnectionRequestDataAttributes(joinAttribute string, joinType string, typeVar string) *CreateConnectionRequestDataAttributes {
	this := CreateConnectionRequestDataAttributes{}
	this.JoinAttribute = joinAttribute
	this.JoinType = joinType
	this.Type = typeVar
	return &this
}

// NewCreateConnectionRequestDataAttributesWithDefaults instantiates a new CreateConnectionRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateConnectionRequestDataAttributesWithDefaults() *CreateConnectionRequestDataAttributes {
	this := CreateConnectionRequestDataAttributes{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributes) GetDetails() ListConnectionsResponseDataAttributesConnectionsItemsDetails {
	if o == nil || o.Details == nil {
		var ret ListConnectionsResponseDataAttributesConnectionsItemsDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetDetailsOk() (*ListConnectionsResponseDataAttributesConnectionsItemsDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributes) HasDetails() bool {
	return o != nil && o.Details != nil
}

// SetDetails gets a reference to the given ListConnectionsResponseDataAttributesConnectionsItemsDetails and assigns it to the Details field.
func (o *CreateConnectionRequestDataAttributes) SetDetails(v ListConnectionsResponseDataAttributesConnectionsItemsDetails) {
	o.Details = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributes) GetFields() []CreateConnectionRequestDataAttributesFieldsItems {
	if o == nil || o.Fields == nil {
		var ret []CreateConnectionRequestDataAttributesFieldsItems
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetFieldsOk() (*[]CreateConnectionRequestDataAttributesFieldsItems, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given []CreateConnectionRequestDataAttributesFieldsItems and assigns it to the Fields field.
func (o *CreateConnectionRequestDataAttributes) SetFields(v []CreateConnectionRequestDataAttributesFieldsItems) {
	o.Fields = v
}

// GetJoinAttribute returns the JoinAttribute field value.
func (o *CreateConnectionRequestDataAttributes) GetJoinAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JoinAttribute
}

// GetJoinAttributeOk returns a tuple with the JoinAttribute field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetJoinAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinAttribute, true
}

// SetJoinAttribute sets field value.
func (o *CreateConnectionRequestDataAttributes) SetJoinAttribute(v string) {
	o.JoinAttribute = v
}

// GetJoinType returns the JoinType field value.
func (o *CreateConnectionRequestDataAttributes) GetJoinType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetJoinTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinType, true
}

// SetJoinType sets field value.
func (o *CreateConnectionRequestDataAttributes) SetJoinType(v string) {
	o.JoinType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributes) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateConnectionRequestDataAttributes) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateConnectionRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value.
func (o *CreateConnectionRequestDataAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateConnectionRequestDataAttributes) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateConnectionRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["join_attribute"] = o.JoinAttribute
	toSerialize["join_type"] = o.JoinType
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateConnectionRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Details       *ListConnectionsResponseDataAttributesConnectionsItemsDetails `json:"details,omitempty"`
		Fields        []CreateConnectionRequestDataAttributesFieldsItems            `json:"fields,omitempty"`
		JoinAttribute *string                                                       `json:"join_attribute"`
		JoinType      *string                                                       `json:"join_type"`
		Metadata      map[string]string                                             `json:"metadata,omitempty"`
		Name          *string                                                       `json:"name,omitempty"`
		Type          *string                                                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JoinAttribute == nil {
		return fmt.Errorf("required field join_attribute missing")
	}
	if all.JoinType == nil {
		return fmt.Errorf("required field join_type missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"details", "fields", "join_attribute", "join_type", "metadata", "name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Details != nil && all.Details.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Details = all.Details
	o.Fields = all.Fields
	o.JoinAttribute = *all.JoinAttribute
	o.JoinType = *all.JoinType
	o.Metadata = all.Metadata
	o.Name = all.Name
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
