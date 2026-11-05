// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelGroupModel Model catalog entry embedded in a model group.
type AiGatewayModelGroupModel struct {
	// Time when the resource was created.
	CreatedAt time.Time `json:"created_at"`
	// Whether the resource is enabled.
	Enabled bool `json:"enabled"`
	// Model catalog entry identifier.
	Id string `json:"id"`
	// Optional, typed enrichment for a model catalog entry. Missing properties remain unknown.
	Metadata *AiGatewayModelMetadata `json:"metadata,omitempty"`
	// Provider-independent model identifier.
	ModelIdCanonical string `json:"model_id_canonical"`
	// Model identifier sent to the provider.
	ModelIdWire string `json:"model_id_wire"`
	// Provider identifier.
	ProviderId int64 `json:"provider_id"`
	// How a model catalog entry is managed.
	Source AiGatewayModelSource `json:"source"`
	// Operations supported by the model.
	SupportedOperations []AiGatewaySupportedOperation `json:"supported_operations"`
	// Time when the resource was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelGroupModel instantiates a new AiGatewayModelGroupModel object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelGroupModel(createdAt time.Time, enabled bool, id string, modelIdCanonical string, modelIdWire string, providerId int64, source AiGatewayModelSource, supportedOperations []AiGatewaySupportedOperation, updatedAt time.Time) *AiGatewayModelGroupModel {
	this := AiGatewayModelGroupModel{}
	this.CreatedAt = createdAt
	this.Enabled = enabled
	this.Id = id
	this.ModelIdCanonical = modelIdCanonical
	this.ModelIdWire = modelIdWire
	this.ProviderId = providerId
	this.Source = source
	this.SupportedOperations = supportedOperations
	this.UpdatedAt = updatedAt
	return &this
}

// NewAiGatewayModelGroupModelWithDefaults instantiates a new AiGatewayModelGroupModel object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelGroupModelWithDefaults() *AiGatewayModelGroupModel {
	this := AiGatewayModelGroupModel{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AiGatewayModelGroupModel) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AiGatewayModelGroupModel) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEnabled returns the Enabled field value.
func (o *AiGatewayModelGroupModel) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *AiGatewayModelGroupModel) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value.
func (o *AiGatewayModelGroupModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AiGatewayModelGroupModel) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AiGatewayModelGroupModel) GetMetadata() AiGatewayModelMetadata {
	if o == nil || o.Metadata == nil {
		var ret AiGatewayModelMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetMetadataOk() (*AiGatewayModelMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AiGatewayModelGroupModel) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given AiGatewayModelMetadata and assigns it to the Metadata field.
func (o *AiGatewayModelGroupModel) SetMetadata(v AiGatewayModelMetadata) {
	o.Metadata = &v
}

// GetModelIdCanonical returns the ModelIdCanonical field value.
func (o *AiGatewayModelGroupModel) GetModelIdCanonical() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelIdCanonical
}

// GetModelIdCanonicalOk returns a tuple with the ModelIdCanonical field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetModelIdCanonicalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelIdCanonical, true
}

// SetModelIdCanonical sets field value.
func (o *AiGatewayModelGroupModel) SetModelIdCanonical(v string) {
	o.ModelIdCanonical = v
}

// GetModelIdWire returns the ModelIdWire field value.
func (o *AiGatewayModelGroupModel) GetModelIdWire() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelIdWire
}

// GetModelIdWireOk returns a tuple with the ModelIdWire field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetModelIdWireOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelIdWire, true
}

// SetModelIdWire sets field value.
func (o *AiGatewayModelGroupModel) SetModelIdWire(v string) {
	o.ModelIdWire = v
}

// GetProviderId returns the ProviderId field value.
func (o *AiGatewayModelGroupModel) GetProviderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetProviderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value.
func (o *AiGatewayModelGroupModel) SetProviderId(v int64) {
	o.ProviderId = v
}

// GetSource returns the Source field value.
func (o *AiGatewayModelGroupModel) GetSource() AiGatewayModelSource {
	if o == nil {
		var ret AiGatewayModelSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetSourceOk() (*AiGatewayModelSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *AiGatewayModelGroupModel) SetSource(v AiGatewayModelSource) {
	o.Source = v
}

// GetSupportedOperations returns the SupportedOperations field value.
func (o *AiGatewayModelGroupModel) GetSupportedOperations() []AiGatewaySupportedOperation {
	if o == nil {
		var ret []AiGatewaySupportedOperation
		return ret
	}
	return o.SupportedOperations
}

// GetSupportedOperationsOk returns a tuple with the SupportedOperations field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetSupportedOperationsOk() (*[]AiGatewaySupportedOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportedOperations, true
}

// SetSupportedOperations sets field value.
func (o *AiGatewayModelGroupModel) SetSupportedOperations(v []AiGatewaySupportedOperation) {
	o.SupportedOperations = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *AiGatewayModelGroupModel) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelGroupModel) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *AiGatewayModelGroupModel) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelGroupModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["id"] = o.Id
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["model_id_canonical"] = o.ModelIdCanonical
	toSerialize["model_id_wire"] = o.ModelIdWire
	toSerialize["provider_id"] = o.ProviderId
	toSerialize["source"] = o.Source
	toSerialize["supported_operations"] = o.SupportedOperations
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelGroupModel) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt           *time.Time                     `json:"created_at"`
		Enabled             *bool                          `json:"enabled"`
		Id                  *string                        `json:"id"`
		Metadata            *AiGatewayModelMetadata        `json:"metadata,omitempty"`
		ModelIdCanonical    *string                        `json:"model_id_canonical"`
		ModelIdWire         *string                        `json:"model_id_wire"`
		ProviderId          *int64                         `json:"provider_id"`
		Source              *AiGatewayModelSource          `json:"source"`
		SupportedOperations *[]AiGatewaySupportedOperation `json:"supported_operations"`
		UpdatedAt           *time.Time                     `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.ModelIdCanonical == nil {
		return fmt.Errorf("required field model_id_canonical missing")
	}
	if all.ModelIdWire == nil {
		return fmt.Errorf("required field model_id_wire missing")
	}
	if all.ProviderId == nil {
		return fmt.Errorf("required field provider_id missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.SupportedOperations == nil {
		return fmt.Errorf("required field supported_operations missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "enabled", "id", "metadata", "model_id_canonical", "model_id_wire", "provider_id", "source", "supported_operations", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	o.Enabled = *all.Enabled
	o.Id = *all.Id
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.ModelIdCanonical = *all.ModelIdCanonical
	o.ModelIdWire = *all.ModelIdWire
	o.ProviderId = *all.ProviderId
	if !all.Source.IsValid() {
		hasInvalidField = true
	} else {
		o.Source = *all.Source
	}
	o.SupportedOperations = *all.SupportedOperations
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
