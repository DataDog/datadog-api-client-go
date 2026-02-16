// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConnectionGroupDataAttributesResponse Attributes of a connection group.
type ConnectionGroupDataAttributesResponse struct {
	// List of connection IDs associated with the connection group.
	Connections []string `json:"connections,omitempty"`
	// The creation timestamp of the connection group.
	CreatedAt time.Time `json:"created_at"`
	// The description of the connection group.
	Description *string `json:"description,omitempty"`
	// The integration type of the connection group.
	IntegrationType string `json:"integration_type"`
	// Indicates if the connection group is marked as favorite.
	IsFavorite bool `json:"is_favorite"`
	// The last updated timestamp of the connection group.
	LastUpdatedAt time.Time `json:"last_updated_at"`
	// The name of the connection group.
	Name string `json:"name"`
	// Tag keys associated with the connection group.
	TagKeys []string `json:"tag_keys"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConnectionGroupDataAttributesResponse instantiates a new ConnectionGroupDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConnectionGroupDataAttributesResponse(createdAt time.Time, integrationType string, isFavorite bool, lastUpdatedAt time.Time, name string, tagKeys []string) *ConnectionGroupDataAttributesResponse {
	this := ConnectionGroupDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.IntegrationType = integrationType
	this.IsFavorite = isFavorite
	this.LastUpdatedAt = lastUpdatedAt
	this.Name = name
	this.TagKeys = tagKeys
	return &this
}

// NewConnectionGroupDataAttributesResponseWithDefaults instantiates a new ConnectionGroupDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConnectionGroupDataAttributesResponseWithDefaults() *ConnectionGroupDataAttributesResponse {
	this := ConnectionGroupDataAttributesResponse{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesResponse) GetConnections() []string {
	if o == nil || o.Connections == nil {
		var ret []string
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetConnectionsOk() (*[]string, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return &o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesResponse) HasConnections() bool {
	return o != nil && o.Connections != nil
}

// SetConnections gets a reference to the given []string and assigns it to the Connections field.
func (o *ConnectionGroupDataAttributesResponse) SetConnections(v []string) {
	o.Connections = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ConnectionGroupDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ConnectionGroupDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectionGroupDataAttributesResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectionGroupDataAttributesResponse) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectionGroupDataAttributesResponse) SetDescription(v string) {
	o.Description = &v
}

// GetIntegrationType returns the IntegrationType field value.
func (o *ConnectionGroupDataAttributesResponse) GetIntegrationType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetIntegrationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationType, true
}

// SetIntegrationType sets field value.
func (o *ConnectionGroupDataAttributesResponse) SetIntegrationType(v string) {
	o.IntegrationType = v
}

// GetIsFavorite returns the IsFavorite field value.
func (o *ConnectionGroupDataAttributesResponse) GetIsFavorite() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsFavorite
}

// GetIsFavoriteOk returns a tuple with the IsFavorite field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetIsFavoriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFavorite, true
}

// SetIsFavorite sets field value.
func (o *ConnectionGroupDataAttributesResponse) SetIsFavorite(v bool) {
	o.IsFavorite = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value.
func (o *ConnectionGroupDataAttributesResponse) GetLastUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value.
func (o *ConnectionGroupDataAttributesResponse) SetLastUpdatedAt(v time.Time) {
	o.LastUpdatedAt = v
}

// GetName returns the Name field value.
func (o *ConnectionGroupDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ConnectionGroupDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetTagKeys returns the TagKeys field value.
func (o *ConnectionGroupDataAttributesResponse) GetTagKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TagKeys
}

// GetTagKeysOk returns a tuple with the TagKeys field value
// and a boolean to check if the value has been set.
func (o *ConnectionGroupDataAttributesResponse) GetTagKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagKeys, true
}

// SetTagKeys sets field value.
func (o *ConnectionGroupDataAttributesResponse) SetTagKeys(v []string) {
	o.TagKeys = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConnectionGroupDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	toSerialize["integration_type"] = o.IntegrationType
	toSerialize["is_favorite"] = o.IsFavorite
	if o.LastUpdatedAt.Nanosecond() == 0 {
		toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["last_updated_at"] = o.LastUpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["tag_keys"] = o.TagKeys

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConnectionGroupDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Connections     []string   `json:"connections,omitempty"`
		CreatedAt       *time.Time `json:"created_at"`
		Description     *string    `json:"description,omitempty"`
		IntegrationType *string    `json:"integration_type"`
		IsFavorite      *bool      `json:"is_favorite"`
		LastUpdatedAt   *time.Time `json:"last_updated_at"`
		Name            *string    `json:"name"`
		TagKeys         *[]string  `json:"tag_keys"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.IntegrationType == nil {
		return fmt.Errorf("required field integration_type missing")
	}
	if all.IsFavorite == nil {
		return fmt.Errorf("required field is_favorite missing")
	}
	if all.LastUpdatedAt == nil {
		return fmt.Errorf("required field last_updated_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.TagKeys == nil {
		return fmt.Errorf("required field tag_keys missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connections", "created_at", "description", "integration_type", "is_favorite", "last_updated_at", "name", "tag_keys"})
	} else {
		return err
	}
	o.Connections = all.Connections
	o.CreatedAt = *all.CreatedAt
	o.Description = all.Description
	o.IntegrationType = *all.IntegrationType
	o.IsFavorite = *all.IsFavorite
	o.LastUpdatedAt = *all.LastUpdatedAt
	o.Name = *all.Name
	o.TagKeys = *all.TagKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
