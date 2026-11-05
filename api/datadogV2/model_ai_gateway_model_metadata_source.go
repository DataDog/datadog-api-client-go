// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelMetadataSource Provenance for model metadata enrichment.
type AiGatewayModelMetadataSource struct {
	// Identifier strategy used to match the gateway catalog entry. Date-suffix and shared canonical-model fallback matches produce partial metadata.
	MatchedBy AiGatewayModelMetadataSourceMatchedBy `json:"matched_by"`
	// Model identifier in the metadata source provider catalog.
	ModelId string `json:"model_id"`
	// Metadata source name.
	Name string `json:"name"`
	// Last successful observation of this valid source record. Retained invalid entries keep their previous timestamp.
	ObservedAt *time.Time `json:"observed_at,omitempty"`
	// Provider namespace in the metadata source, or canonical for a shared model record.
	Provider string `json:"provider"`
	// SHA-256 content revision of the source snapshot that supplied this record.
	Revision *string `json:"revision,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelMetadataSource instantiates a new AiGatewayModelMetadataSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelMetadataSource(matchedBy AiGatewayModelMetadataSourceMatchedBy, modelId string, name string, provider string) *AiGatewayModelMetadataSource {
	this := AiGatewayModelMetadataSource{}
	this.MatchedBy = matchedBy
	this.ModelId = modelId
	this.Name = name
	this.Provider = provider
	return &this
}

// NewAiGatewayModelMetadataSourceWithDefaults instantiates a new AiGatewayModelMetadataSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelMetadataSourceWithDefaults() *AiGatewayModelMetadataSource {
	this := AiGatewayModelMetadataSource{}
	return &this
}

// GetMatchedBy returns the MatchedBy field value.
func (o *AiGatewayModelMetadataSource) GetMatchedBy() AiGatewayModelMetadataSourceMatchedBy {
	if o == nil {
		var ret AiGatewayModelMetadataSourceMatchedBy
		return ret
	}
	return o.MatchedBy
}

// GetMatchedByOk returns a tuple with the MatchedBy field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadataSource) GetMatchedByOk() (*AiGatewayModelMetadataSourceMatchedBy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchedBy, true
}

// SetMatchedBy sets field value.
func (o *AiGatewayModelMetadataSource) SetMatchedBy(v AiGatewayModelMetadataSourceMatchedBy) {
	o.MatchedBy = v
}

// GetModelId returns the ModelId field value.
func (o *AiGatewayModelMetadataSource) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadataSource) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value.
func (o *AiGatewayModelMetadataSource) SetModelId(v string) {
	o.ModelId = v
}

// GetName returns the Name field value.
func (o *AiGatewayModelMetadataSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadataSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayModelMetadataSource) SetName(v string) {
	o.Name = v
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *AiGatewayModelMetadataSource) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadataSource) GetObservedAtOk() (*time.Time, bool) {
	if o == nil || o.ObservedAt == nil {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *AiGatewayModelMetadataSource) HasObservedAt() bool {
	return o != nil && o.ObservedAt != nil
}

// SetObservedAt gets a reference to the given time.Time and assigns it to the ObservedAt field.
func (o *AiGatewayModelMetadataSource) SetObservedAt(v time.Time) {
	o.ObservedAt = &v
}

// GetProvider returns the Provider field value.
func (o *AiGatewayModelMetadataSource) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadataSource) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *AiGatewayModelMetadataSource) SetProvider(v string) {
	o.Provider = v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *AiGatewayModelMetadataSource) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelMetadataSource) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *AiGatewayModelMetadataSource) HasRevision() bool {
	return o != nil && o.Revision != nil
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *AiGatewayModelMetadataSource) SetRevision(v string) {
	o.Revision = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelMetadataSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["matched_by"] = o.MatchedBy
	toSerialize["model_id"] = o.ModelId
	toSerialize["name"] = o.Name
	if o.ObservedAt != nil {
		if o.ObservedAt.Nanosecond() == 0 {
			toSerialize["observed_at"] = o.ObservedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["observed_at"] = o.ObservedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	toSerialize["provider"] = o.Provider
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelMetadataSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchedBy  *AiGatewayModelMetadataSourceMatchedBy `json:"matched_by"`
		ModelId    *string                                `json:"model_id"`
		Name       *string                                `json:"name"`
		ObservedAt *time.Time                             `json:"observed_at,omitempty"`
		Provider   *string                                `json:"provider"`
		Revision   *string                                `json:"revision,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MatchedBy == nil {
		return fmt.Errorf("required field matched_by missing")
	}
	if all.ModelId == nil {
		return fmt.Errorf("required field model_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"matched_by", "model_id", "name", "observed_at", "provider", "revision"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.MatchedBy.IsValid() {
		hasInvalidField = true
	} else {
		o.MatchedBy = *all.MatchedBy
	}
	o.ModelId = *all.ModelId
	o.Name = *all.Name
	o.ObservedAt = all.ObservedAt
	o.Provider = *all.Provider
	o.Revision = all.Revision

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
