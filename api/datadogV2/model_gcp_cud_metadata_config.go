// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPCudMetadataConfig The BigQuery table the committed use discount (CUD) metadata is read from. Only returned when CUD metadata is configured for the account.
type GCPCudMetadataConfig struct {
	// The BigQuery dataset ID of the CUD metadata export.
	DatasetId string `json:"dataset_id"`
	// The Google Cloud project ID of the CUD metadata export.
	ProjectId string `json:"project_id"`
	// The BigQuery table ID of the CUD metadata export.
	TableId string `json:"table_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPCudMetadataConfig instantiates a new GCPCudMetadataConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPCudMetadataConfig(datasetId string, projectId string, tableId string) *GCPCudMetadataConfig {
	this := GCPCudMetadataConfig{}
	this.DatasetId = datasetId
	this.ProjectId = projectId
	this.TableId = tableId
	return &this
}

// NewGCPCudMetadataConfigWithDefaults instantiates a new GCPCudMetadataConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPCudMetadataConfigWithDefaults() *GCPCudMetadataConfig {
	this := GCPCudMetadataConfig{}
	return &this
}

// GetDatasetId returns the DatasetId field value.
func (o *GCPCudMetadataConfig) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *GCPCudMetadataConfig) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *GCPCudMetadataConfig) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetProjectId returns the ProjectId field value.
func (o *GCPCudMetadataConfig) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GCPCudMetadataConfig) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *GCPCudMetadataConfig) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTableId returns the TableId field value.
func (o *GCPCudMetadataConfig) GetTableId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value
// and a boolean to check if the value has been set.
func (o *GCPCudMetadataConfig) GetTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableId, true
}

// SetTableId sets field value.
func (o *GCPCudMetadataConfig) SetTableId(v string) {
	o.TableId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPCudMetadataConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["project_id"] = o.ProjectId
	toSerialize["table_id"] = o.TableId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPCudMetadataConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetId *string `json:"dataset_id"`
		ProjectId *string `json:"project_id"`
		TableId   *string `json:"table_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.TableId == nil {
		return fmt.Errorf("required field table_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset_id", "project_id", "table_id"})
	} else {
		return err
	}
	o.DatasetId = *all.DatasetId
	o.ProjectId = *all.ProjectId
	o.TableId = *all.TableId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
