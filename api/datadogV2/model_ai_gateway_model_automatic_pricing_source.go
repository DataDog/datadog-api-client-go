// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelAutomaticPricingSource The pricing source entry the price was read from.
type AiGatewayModelAutomaticPricingSource struct {
	// How the gateway matched the model to the source entry, such as `wire_id`.
	MatchMethod string `json:"match_method"`
	// Model identifier in the source.
	ModelId string `json:"model_id"`
	// Pricing source name.
	Name string `json:"name"`
	// Provider identifier in the source.
	ProviderId string `json:"provider_id"`
	// Source revision the price was read from.
	Revision string `json:"revision"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelAutomaticPricingSource instantiates a new AiGatewayModelAutomaticPricingSource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelAutomaticPricingSource(matchMethod string, modelId string, name string, providerId string, revision string) *AiGatewayModelAutomaticPricingSource {
	this := AiGatewayModelAutomaticPricingSource{}
	this.MatchMethod = matchMethod
	this.ModelId = modelId
	this.Name = name
	this.ProviderId = providerId
	this.Revision = revision
	return &this
}

// NewAiGatewayModelAutomaticPricingSourceWithDefaults instantiates a new AiGatewayModelAutomaticPricingSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelAutomaticPricingSourceWithDefaults() *AiGatewayModelAutomaticPricingSource {
	this := AiGatewayModelAutomaticPricingSource{}
	return &this
}

// GetMatchMethod returns the MatchMethod field value.
func (o *AiGatewayModelAutomaticPricingSource) GetMatchMethod() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.MatchMethod
}

// GetMatchMethodOk returns a tuple with the MatchMethod field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingSource) GetMatchMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchMethod, true
}

// SetMatchMethod sets field value.
func (o *AiGatewayModelAutomaticPricingSource) SetMatchMethod(v string) {
	o.MatchMethod = v
}

// GetModelId returns the ModelId field value.
func (o *AiGatewayModelAutomaticPricingSource) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingSource) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value.
func (o *AiGatewayModelAutomaticPricingSource) SetModelId(v string) {
	o.ModelId = v
}

// GetName returns the Name field value.
func (o *AiGatewayModelAutomaticPricingSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiGatewayModelAutomaticPricingSource) SetName(v string) {
	o.Name = v
}

// GetProviderId returns the ProviderId field value.
func (o *AiGatewayModelAutomaticPricingSource) GetProviderId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingSource) GetProviderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderId, true
}

// SetProviderId sets field value.
func (o *AiGatewayModelAutomaticPricingSource) SetProviderId(v string) {
	o.ProviderId = v
}

// GetRevision returns the Revision field value.
func (o *AiGatewayModelAutomaticPricingSource) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingSource) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value.
func (o *AiGatewayModelAutomaticPricingSource) SetRevision(v string) {
	o.Revision = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelAutomaticPricingSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["match_method"] = o.MatchMethod
	toSerialize["model_id"] = o.ModelId
	toSerialize["name"] = o.Name
	toSerialize["provider_id"] = o.ProviderId
	toSerialize["revision"] = o.Revision

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelAutomaticPricingSource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MatchMethod *string `json:"match_method"`
		ModelId     *string `json:"model_id"`
		Name        *string `json:"name"`
		ProviderId  *string `json:"provider_id"`
		Revision    *string `json:"revision"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MatchMethod == nil {
		return fmt.Errorf("required field match_method missing")
	}
	if all.ModelId == nil {
		return fmt.Errorf("required field model_id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProviderId == nil {
		return fmt.Errorf("required field provider_id missing")
	}
	if all.Revision == nil {
		return fmt.Errorf("required field revision missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"match_method", "model_id", "name", "provider_id", "revision"})
	} else {
		return err
	}
	o.MatchMethod = *all.MatchMethod
	o.ModelId = *all.ModelId
	o.Name = *all.Name
	o.ProviderId = *all.ProviderId
	o.Revision = *all.Revision

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
