// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestination The `azure_data_explorer` destination sends log events to an Azure Data Explorer table.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineAzureDataExplorerDestination struct {
	// Authentication configuration for Azure Data Explorer. The `azure_credential_kind` field selects the credential type.
	Auth ObservabilityPipelineAzureDataExplorerDestinationAuth `json:"auth"`
	// Event batching settings for Azure Data Explorer ingestion.
	Batch *ObservabilityPipelineAzureDataExplorerDestinationBatch `json:"batch,omitempty"`
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// Gzip compression.
	Compression *ObservabilityPipelineAzureStorageDestinationCompressionGzip `json:"compression,omitempty"`
	// The name of the Azure Data Explorer database to ingest into. Supports template syntax.
	Database string `json:"database"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// Name of the environment variable or secret that holds the Azure Data Explorer ingestion endpoint URL.
	// Defaults to `DESTINATION_AZURE_DATA_EXPLORER_INGESTION_ENDPOINT` (prefixed with `DD_OP_` at runtime).
	IngestionEndpointKey *string `json:"ingestion_endpoint_key,omitempty"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The name of a pre-created ingestion mapping on the table used to map incoming events to columns. Supports template syntax.
	MappingReference datadog.NullableString `json:"mapping_reference,omitempty"`
	// The name of the Azure Data Explorer table to ingest into. Supports template syntax.
	Table string `json:"table"`
	// The OAuth scope requested when acquiring an access token for Azure Data Explorer.
	// Defaults to `https://kusto.kusto.windows.net/.default`.
	TokenScope *string `json:"token_scope,omitempty"`
	// The destination type. The value should always be `azure_data_explorer`.
	Type ObservabilityPipelineAzureDataExplorerDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerDestination instantiates a new ObservabilityPipelineAzureDataExplorerDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerDestination(auth ObservabilityPipelineAzureDataExplorerDestinationAuth, database string, id string, inputs []string, table string, typeVar ObservabilityPipelineAzureDataExplorerDestinationType) *ObservabilityPipelineAzureDataExplorerDestination {
	this := ObservabilityPipelineAzureDataExplorerDestination{}
	this.Auth = auth
	this.Database = database
	this.Id = id
	this.Inputs = inputs
	this.Table = table
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineAzureDataExplorerDestinationWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerDestinationWithDefaults() *ObservabilityPipelineAzureDataExplorerDestination {
	this := ObservabilityPipelineAzureDataExplorerDestination{}
	var typeVar ObservabilityPipelineAzureDataExplorerDestinationType = OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONTYPE_AZURE_DATA_EXPLORER
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetAuth() ObservabilityPipelineAzureDataExplorerDestinationAuth {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetAuthOk() (*ObservabilityPipelineAzureDataExplorerDestinationAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetAuth(v ObservabilityPipelineAzureDataExplorerDestinationAuth) {
	o.Auth = v
}

// GetBatch returns the Batch field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetBatch() ObservabilityPipelineAzureDataExplorerDestinationBatch {
	if o == nil || o.Batch == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationBatch
		return ret
	}
	return *o.Batch
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetBatchOk() (*ObservabilityPipelineAzureDataExplorerDestinationBatch, bool) {
	if o == nil || o.Batch == nil {
		return nil, false
	}
	return o.Batch, true
}

// HasBatch returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) HasBatch() bool {
	return o != nil && o.Batch != nil
}

// SetBatch gets a reference to the given ObservabilityPipelineAzureDataExplorerDestinationBatch and assigns it to the Batch field.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetBatch(v ObservabilityPipelineAzureDataExplorerDestinationBatch) {
	o.Batch = &v
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetCompression() ObservabilityPipelineAzureStorageDestinationCompressionGzip {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineAzureStorageDestinationCompressionGzip
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetCompressionOk() (*ObservabilityPipelineAzureStorageDestinationCompressionGzip, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineAzureStorageDestinationCompressionGzip and assigns it to the Compression field.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetCompression(v ObservabilityPipelineAzureStorageDestinationCompressionGzip) {
	o.Compression = &v
}

// GetDatabase returns the Database field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Database, true
}

// SetDatabase sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetDatabase(v string) {
	o.Database = v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetId(v string) {
	o.Id = v
}

// GetIngestionEndpointKey returns the IngestionEndpointKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetIngestionEndpointKey() string {
	if o == nil || o.IngestionEndpointKey == nil {
		var ret string
		return ret
	}
	return *o.IngestionEndpointKey
}

// GetIngestionEndpointKeyOk returns a tuple with the IngestionEndpointKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetIngestionEndpointKeyOk() (*string, bool) {
	if o == nil || o.IngestionEndpointKey == nil {
		return nil, false
	}
	return o.IngestionEndpointKey, true
}

// HasIngestionEndpointKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) HasIngestionEndpointKey() bool {
	return o != nil && o.IngestionEndpointKey != nil
}

// SetIngestionEndpointKey gets a reference to the given string and assigns it to the IngestionEndpointKey field.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetIngestionEndpointKey(v string) {
	o.IngestionEndpointKey = &v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetMappingReference returns the MappingReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetMappingReference() string {
	if o == nil || o.MappingReference.Get() == nil {
		var ret string
		return ret
	}
	return *o.MappingReference.Get()
}

// GetMappingReferenceOk returns a tuple with the MappingReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetMappingReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MappingReference.Get(), o.MappingReference.IsSet()
}

// HasMappingReference returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) HasMappingReference() bool {
	return o != nil && o.MappingReference.IsSet()
}

// SetMappingReference gets a reference to the given datadog.NullableString and assigns it to the MappingReference field.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetMappingReference(v string) {
	o.MappingReference.Set(&v)
}

// SetMappingReferenceNil sets the value for MappingReference to be an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetMappingReferenceNil() {
	o.MappingReference.Set(nil)
}

// UnsetMappingReference ensures that no value is present for MappingReference, not even an explicit nil.
func (o *ObservabilityPipelineAzureDataExplorerDestination) UnsetMappingReference() {
	o.MappingReference.Unset()
}

// GetTable returns the Table field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetTable(v string) {
	o.Table = v
}

// GetTokenScope returns the TokenScope field value if set, zero value otherwise.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetTokenScope() string {
	if o == nil || o.TokenScope == nil {
		var ret string
		return ret
	}
	return *o.TokenScope
}

// GetTokenScopeOk returns a tuple with the TokenScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetTokenScopeOk() (*string, bool) {
	if o == nil || o.TokenScope == nil {
		return nil, false
	}
	return o.TokenScope, true
}

// HasTokenScope returns a boolean if a field has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) HasTokenScope() bool {
	return o != nil && o.TokenScope != nil
}

// SetTokenScope gets a reference to the given string and assigns it to the TokenScope field.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetTokenScope(v string) {
	o.TokenScope = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetType() ObservabilityPipelineAzureDataExplorerDestinationType {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestination) GetTypeOk() (*ObservabilityPipelineAzureDataExplorerDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestination) SetType(v ObservabilityPipelineAzureDataExplorerDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	if o.Batch != nil {
		toSerialize["batch"] = o.Batch
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	toSerialize["database"] = o.Database
	toSerialize["id"] = o.Id
	if o.IngestionEndpointKey != nil {
		toSerialize["ingestion_endpoint_key"] = o.IngestionEndpointKey
	}
	toSerialize["inputs"] = o.Inputs
	if o.MappingReference.IsSet() {
		toSerialize["mapping_reference"] = o.MappingReference.Get()
	}
	toSerialize["table"] = o.Table
	if o.TokenScope != nil {
		toSerialize["token_scope"] = o.TokenScope
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAzureDataExplorerDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth                 *ObservabilityPipelineAzureDataExplorerDestinationAuth       `json:"auth"`
		Batch                *ObservabilityPipelineAzureDataExplorerDestinationBatch      `json:"batch,omitempty"`
		Buffer               *ObservabilityPipelineBufferOptions                          `json:"buffer,omitempty"`
		Compression          *ObservabilityPipelineAzureStorageDestinationCompressionGzip `json:"compression,omitempty"`
		Database             *string                                                      `json:"database"`
		Id                   *string                                                      `json:"id"`
		IngestionEndpointKey *string                                                      `json:"ingestion_endpoint_key,omitempty"`
		Inputs               *[]string                                                    `json:"inputs"`
		MappingReference     datadog.NullableString                                       `json:"mapping_reference,omitempty"`
		Table                *string                                                      `json:"table"`
		TokenScope           *string                                                      `json:"token_scope,omitempty"`
		Type                 *ObservabilityPipelineAzureDataExplorerDestinationType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.Database == nil {
		return fmt.Errorf("required field database missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Table == nil {
		return fmt.Errorf("required field table missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "batch", "buffer", "compression", "database", "id", "ingestion_endpoint_key", "inputs", "mapping_reference", "table", "token_scope", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Auth = *all.Auth
	if all.Batch != nil && all.Batch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Batch = all.Batch
	o.Buffer = all.Buffer
	if all.Compression != nil && all.Compression.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Compression = all.Compression
	o.Database = *all.Database
	o.Id = *all.Id
	o.IngestionEndpointKey = all.IngestionEndpointKey
	o.Inputs = *all.Inputs
	o.MappingReference = all.MappingReference
	o.Table = *all.Table
	o.TokenScope = all.TokenScope
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
