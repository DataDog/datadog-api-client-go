// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConnectionGroupDataAttributesRequest Attributes for updating a connection group.
type ConnectionGroupDataAttributesRequest struct {
	// List of connection IDs associated with the connection group.
	Connections []string `json:"connections,omitempty"`
	// The description of the connection group.
	Description *string `json:"description,omitempty"`
	// The integration type of the connection group.
	IntegrationType *string `json:"integration_type,omitempty"`
	// The name of the connection group.
	Name *string `json:"name,omitempty"`
	// Policy attributes for the connection group.
	PolicyAttributes *string `json:"policy_attributes,omitempty"`
	// Tag keys associated with the connection group.
	TagKeys []string `json:"tag_keys,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConnectionGroupDataAttributesRequest instantiates a new ConnectionGroupDataAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConnectionGroupDataAttributesRequest() *ConnectionGroupDataAttributesRequest {
	this := ConnectionGroupDataAttributesRequest{}
	return &this
}

// NewConnectionGroupDataAttributesRequestWithDefaults instantiates a new ConnectionGroupDataAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConnectionGroupDataAttributesRequestWithDefaults() *ConnectionGroupDataAttributesRequest {
	this := ConnectionGroupDataAttributesRequest{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesRequest) GetConnections() []string {
	if o == nil || o.Connections == nil {
		var ret []string
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesRequest) GetConnectionsOk() (*[]string, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return &o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesRequest) HasConnections() bool {
	return o != nil && o.Connections != nil
}

// SetConnections gets a reference to the given []string and assigns it to the Connections field.
func (o *ConnectionGroupDataAttributesRequest) SetConnections(v []string) {
	o.Connections = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectionGroupDataAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesRequest) GetIntegrationType() string {
	if o == nil || o.IntegrationType == nil {
		var ret string
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesRequest) GetIntegrationTypeOk() (*string, bool) {
	if o == nil || o.IntegrationType == nil {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesRequest) HasIntegrationType() bool {
	return o != nil && o.IntegrationType != nil
}

// SetIntegrationType gets a reference to the given string and assigns it to the IntegrationType field.
func (o *ConnectionGroupDataAttributesRequest) SetIntegrationType(v string) {
	o.IntegrationType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectionGroupDataAttributesRequest) SetName(v string) {
	o.Name = &v
}

// GetPolicyAttributes returns the PolicyAttributes field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesRequest) GetPolicyAttributes() string {
	if o == nil || o.PolicyAttributes == nil {
		var ret string
		return ret
	}
	return *o.PolicyAttributes
}

// GetPolicyAttributesOk returns a tuple with the PolicyAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesRequest) GetPolicyAttributesOk() (*string, bool) {
	if o == nil || o.PolicyAttributes == nil {
		return nil, false
	}
	return o.PolicyAttributes, true
}

// HasPolicyAttributes returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesRequest) HasPolicyAttributes() bool {
	return o != nil && o.PolicyAttributes != nil
}

// SetPolicyAttributes gets a reference to the given string and assigns it to the PolicyAttributes field.
func (o *ConnectionGroupDataAttributesRequest) SetPolicyAttributes(v string) {
	o.PolicyAttributes = &v
}

// GetTagKeys returns the TagKeys field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesRequest) GetTagKeys() []string {
	if o == nil || o.TagKeys == nil {
		var ret []string
		return ret
	}
	return o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesRequest) GetTagKeysOk() (*[]string, bool) {
	if o == nil || o.TagKeys == nil {
		return nil, false
	}
	return &o.TagKeys, true
}

// HasTagKeys returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesRequest) HasTagKeys() bool {
	return o != nil && o.TagKeys != nil
}

// SetTagKeys gets a reference to the given []string and assigns it to the TagKeys field.
func (o *ConnectionGroupDataAttributesRequest) SetTagKeys(v []string) {
	o.TagKeys = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConnectionGroupDataAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IntegrationType != nil {
		toSerialize["integration_type"] = o.IntegrationType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PolicyAttributes != nil {
		toSerialize["policy_attributes"] = o.PolicyAttributes
	}
	if o.TagKeys != nil {
		toSerialize["tag_keys"] = o.TagKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConnectionGroupDataAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Connections      []string `json:"connections,omitempty"`
		Description      *string  `json:"description,omitempty"`
		IntegrationType  *string  `json:"integration_type,omitempty"`
		Name             *string  `json:"name,omitempty"`
		PolicyAttributes *string  `json:"policy_attributes,omitempty"`
		TagKeys          []string `json:"tag_keys,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connections", "description", "integration_type", "name", "policy_attributes", "tag_keys"})
	} else {
		return err
	}
	o.Connections = all.Connections
	o.Description = all.Description
	o.IntegrationType = all.IntegrationType
	o.Name = all.Name
	o.PolicyAttributes = all.PolicyAttributes
	o.TagKeys = all.TagKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
