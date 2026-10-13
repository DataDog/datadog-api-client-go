// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli Authenticate using the Azure CLI credentials available in the environment.
type ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli struct {
	// The Azure credential kind. The value should always be `azure_cli`.
	AzureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind `json:"azure_credential_kind"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli(azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind) *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli{}
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// NewObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliWithDefaults instantiates a new ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliWithDefaults() *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli {
	this := ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli{}
	var azureCredentialKind ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind = OBSERVABILITYPIPELINEAZUREDATAEXPLORERDESTINATIONAUTHAZURECLIKIND_AZURE_CLI
	this.AzureCredentialKind = azureCredentialKind
	return &this
}

// GetAzureCredentialKind returns the AzureCredentialKind field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli) GetAzureCredentialKind() ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind {
	if o == nil {
		var ret ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind
		return ret
	}
	return o.AzureCredentialKind
}

// GetAzureCredentialKindOk returns a tuple with the AzureCredentialKind field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli) GetAzureCredentialKindOk() (*ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AzureCredentialKind, true
}

// SetAzureCredentialKind sets field value.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli) SetAzureCredentialKind(v ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind) {
	o.AzureCredentialKind = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["azure_credential_kind"] = o.AzureCredentialKind

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCli) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AzureCredentialKind *ObservabilityPipelineAzureDataExplorerDestinationAuthAzureCliKind `json:"azure_credential_kind"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AzureCredentialKind == nil {
		return fmt.Errorf("required field azure_credential_kind missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"azure_credential_kind"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AzureCredentialKind.IsValid() {
		hasInvalidField = true
	} else {
		o.AzureCredentialKind = *all.AzureCredentialKind
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
