// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureFallbackDestination The Azure archive destination.
type AzureFallbackDestination struct {
	// The container where the archive will be stored.
	Container string `json:"container"`
	// The Azure archive's integration destination.
	Integration AzureFallbackDestinationIntegration `json:"integration"`
	// The archive path.
	Path *string `json:"path,omitempty"`
	// The region where the archive will be stored.
	Region *string `json:"region,omitempty"`
	// The associated storage account.
	StorageAccount string `json:"storageAccount"`
	// Type of the Azure archive destination.
	Type AzureFallbackDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAzureFallbackDestination instantiates a new AzureFallbackDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureFallbackDestination(container string, integration AzureFallbackDestinationIntegration, storageAccount string, typeVar AzureFallbackDestinationType) *AzureFallbackDestination {
	this := AzureFallbackDestination{}
	this.Container = container
	this.Integration = integration
	this.StorageAccount = storageAccount
	this.Type = typeVar
	return &this
}

// NewAzureFallbackDestinationWithDefaults instantiates a new AzureFallbackDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureFallbackDestinationWithDefaults() *AzureFallbackDestination {
	this := AzureFallbackDestination{}
	var typeVar AzureFallbackDestinationType = AZUREFALLBACKDESTINATIONTYPE_AZURE
	this.Type = typeVar
	return &this
}

// GetContainer returns the Container field value.
func (o *AzureFallbackDestination) GetContainer() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Container
}

// GetContainerOk returns a tuple with the Container field value
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestination) GetContainerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Container, true
}

// SetContainer sets field value.
func (o *AzureFallbackDestination) SetContainer(v string) {
	o.Container = v
}

// GetIntegration returns the Integration field value.
func (o *AzureFallbackDestination) GetIntegration() AzureFallbackDestinationIntegration {
	if o == nil {
		var ret AzureFallbackDestinationIntegration
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestination) GetIntegrationOk() (*AzureFallbackDestinationIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *AzureFallbackDestination) SetIntegration(v AzureFallbackDestinationIntegration) {
	o.Integration = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *AzureFallbackDestination) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestination) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *AzureFallbackDestination) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *AzureFallbackDestination) SetPath(v string) {
	o.Path = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AzureFallbackDestination) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestination) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AzureFallbackDestination) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AzureFallbackDestination) SetRegion(v string) {
	o.Region = &v
}

// GetStorageAccount returns the StorageAccount field value.
func (o *AzureFallbackDestination) GetStorageAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StorageAccount
}

// GetStorageAccountOk returns a tuple with the StorageAccount field value
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestination) GetStorageAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageAccount, true
}

// SetStorageAccount sets field value.
func (o *AzureFallbackDestination) SetStorageAccount(v string) {
	o.StorageAccount = v
}

// GetType returns the Type field value.
func (o *AzureFallbackDestination) GetType() AzureFallbackDestinationType {
	if o == nil {
		var ret AzureFallbackDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureFallbackDestination) GetTypeOk() (*AzureFallbackDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AzureFallbackDestination) SetType(v AzureFallbackDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureFallbackDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["container"] = o.Container
	toSerialize["integration"] = o.Integration
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	toSerialize["storageAccount"] = o.StorageAccount
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureFallbackDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Container      *string                              `json:"container"`
		Integration    *AzureFallbackDestinationIntegration `json:"integration"`
		Path           *string                              `json:"path,omitempty"`
		Region         *string                              `json:"region,omitempty"`
		StorageAccount *string                              `json:"storageAccount"`
		Type           *AzureFallbackDestinationType        `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Container == nil {
		return fmt.Errorf("required field container missing")
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	if all.StorageAccount == nil {
		return fmt.Errorf("required field storageAccount missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"container", "integration", "path", "region", "storageAccount", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Container = *all.Container
	if all.Integration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integration = *all.Integration
	o.Path = all.Path
	o.Region = all.Region
	o.StorageAccount = *all.StorageAccount
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
