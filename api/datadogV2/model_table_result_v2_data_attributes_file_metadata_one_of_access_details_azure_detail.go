// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail The definition of `TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail` object.
type TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail struct {
	// The Azure client ID.
	AzureClientId *string `json:"azure_client_id,omitempty"`
	// The name of the Azure container.
	AzureContainerName *string `json:"azure_container_name,omitempty"`
	// The name of the Azure storage account.
	AzureStorageAccountName *string `json:"azure_storage_account_name,omitempty"`
	// The ID of the Azure tenant.
	AzureTenantId *string `json:"azure_tenant_id,omitempty"`
	// The relative file path from the Azure container root to the CSV file.
	FilePath *string `json:"file_path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail instantiates a new TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail() *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail {
	this := TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail{}
	return &this
}

// NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetailWithDefaults instantiates a new TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetailWithDefaults() *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail {
	this := TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail{}
	return &this
}

// GetAzureClientId returns the AzureClientId field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureClientId() string {
	if o == nil || o.AzureClientId == nil {
		var ret string
		return ret
	}
	return *o.AzureClientId
}

// GetAzureClientIdOk returns a tuple with the AzureClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureClientIdOk() (*string, bool) {
	if o == nil || o.AzureClientId == nil {
		return nil, false
	}
	return o.AzureClientId, true
}

// HasAzureClientId returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) HasAzureClientId() bool {
	return o != nil && o.AzureClientId != nil
}

// SetAzureClientId gets a reference to the given string and assigns it to the AzureClientId field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureClientId(v string) {
	o.AzureClientId = &v
}

// GetAzureContainerName returns the AzureContainerName field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureContainerName() string {
	if o == nil || o.AzureContainerName == nil {
		var ret string
		return ret
	}
	return *o.AzureContainerName
}

// GetAzureContainerNameOk returns a tuple with the AzureContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureContainerNameOk() (*string, bool) {
	if o == nil || o.AzureContainerName == nil {
		return nil, false
	}
	return o.AzureContainerName, true
}

// HasAzureContainerName returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) HasAzureContainerName() bool {
	return o != nil && o.AzureContainerName != nil
}

// SetAzureContainerName gets a reference to the given string and assigns it to the AzureContainerName field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureContainerName(v string) {
	o.AzureContainerName = &v
}

// GetAzureStorageAccountName returns the AzureStorageAccountName field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureStorageAccountName() string {
	if o == nil || o.AzureStorageAccountName == nil {
		var ret string
		return ret
	}
	return *o.AzureStorageAccountName
}

// GetAzureStorageAccountNameOk returns a tuple with the AzureStorageAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureStorageAccountNameOk() (*string, bool) {
	if o == nil || o.AzureStorageAccountName == nil {
		return nil, false
	}
	return o.AzureStorageAccountName, true
}

// HasAzureStorageAccountName returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) HasAzureStorageAccountName() bool {
	return o != nil && o.AzureStorageAccountName != nil
}

// SetAzureStorageAccountName gets a reference to the given string and assigns it to the AzureStorageAccountName field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureStorageAccountName(v string) {
	o.AzureStorageAccountName = &v
}

// GetAzureTenantId returns the AzureTenantId field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureTenantId() string {
	if o == nil || o.AzureTenantId == nil {
		var ret string
		return ret
	}
	return *o.AzureTenantId
}

// GetAzureTenantIdOk returns a tuple with the AzureTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetAzureTenantIdOk() (*string, bool) {
	if o == nil || o.AzureTenantId == nil {
		return nil, false
	}
	return o.AzureTenantId, true
}

// HasAzureTenantId returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) HasAzureTenantId() bool {
	return o != nil && o.AzureTenantId != nil
}

// SetAzureTenantId gets a reference to the given string and assigns it to the AzureTenantId field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetAzureTenantId(v string) {
	o.AzureTenantId = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) HasFilePath() bool {
	return o != nil && o.FilePath != nil
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) SetFilePath(v string) {
	o.FilePath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AzureClientId != nil {
		toSerialize["azure_client_id"] = o.AzureClientId
	}
	if o.AzureContainerName != nil {
		toSerialize["azure_container_name"] = o.AzureContainerName
	}
	if o.AzureStorageAccountName != nil {
		toSerialize["azure_storage_account_name"] = o.AzureStorageAccountName
	}
	if o.AzureTenantId != nil {
		toSerialize["azure_tenant_id"] = o.AzureTenantId
	}
	if o.FilePath != nil {
		toSerialize["file_path"] = o.FilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureClientId           *string `json:"azure_client_id,omitempty"`
		AzureContainerName      *string `json:"azure_container_name,omitempty"`
		AzureStorageAccountName *string `json:"azure_storage_account_name,omitempty"`
		AzureTenantId           *string `json:"azure_tenant_id,omitempty"`
		FilePath                *string `json:"file_path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_client_id", "azure_container_name", "azure_storage_account_name", "azure_tenant_id", "file_path"})
	} else {
		return err
	}
	o.AzureClientId = all.AzureClientId
	o.AzureContainerName = all.AzureContainerName
	o.AzureStorageAccountName = all.AzureStorageAccountName
	o.AzureTenantId = all.AzureTenantId
	o.FilePath = all.FilePath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
