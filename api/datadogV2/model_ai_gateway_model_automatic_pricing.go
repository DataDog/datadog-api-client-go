// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelAutomaticPricing The automatically synced published price. Present only when it is the effective price.
type AiGatewayModelAutomaticPricing struct {
	// Time the price was last revalidated against the source.
	LastValidatedAt time.Time `json:"last_validated_at"`
	// The reviewed endpoint mapping used to identify a custom provider. Omitted when the provider and model identity matched exactly.
	Mapping *AiGatewayModelAutomaticPricingMapping `json:"mapping,omitempty"`
	// Request fields that select each priced mode, such as `speed: fast` for `anthropic` models. Omitted when the model has a single mode.
	ModeTriggers []AiGatewayModelAutomaticPricingModeTrigger `json:"mode_triggers,omitempty"`
	// Time the price was observed in the source.
	ObservedAt time.Time `json:"observed_at"`
	// Published rates.
	Rates []AiGatewayModelAutomaticPricingRate `json:"rates"`
	// The pricing source entry the price was read from.
	Source AiGatewayModelAutomaticPricingSource `json:"source"`
	// True when the price has not been revalidated recently and is being retained as last-known-good.
	Stale bool `json:"stale"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelAutomaticPricing instantiates a new AiGatewayModelAutomaticPricing object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelAutomaticPricing(lastValidatedAt time.Time, observedAt time.Time, rates []AiGatewayModelAutomaticPricingRate, source AiGatewayModelAutomaticPricingSource, stale bool) *AiGatewayModelAutomaticPricing {
	this := AiGatewayModelAutomaticPricing{}
	this.LastValidatedAt = lastValidatedAt
	this.ObservedAt = observedAt
	this.Rates = rates
	this.Source = source
	this.Stale = stale
	return &this
}

// NewAiGatewayModelAutomaticPricingWithDefaults instantiates a new AiGatewayModelAutomaticPricing object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelAutomaticPricingWithDefaults() *AiGatewayModelAutomaticPricing {
	this := AiGatewayModelAutomaticPricing{}
	return &this
}

// GetLastValidatedAt returns the LastValidatedAt field value.
func (o *AiGatewayModelAutomaticPricing) GetLastValidatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.LastValidatedAt
}

// GetLastValidatedAtOk returns a tuple with the LastValidatedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricing) GetLastValidatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastValidatedAt, true
}

// SetLastValidatedAt sets field value.
func (o *AiGatewayModelAutomaticPricing) SetLastValidatedAt(v time.Time) {
	o.LastValidatedAt = v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *AiGatewayModelAutomaticPricing) GetMapping() AiGatewayModelAutomaticPricingMapping {
	if o == nil || o.Mapping == nil {
		var ret AiGatewayModelAutomaticPricingMapping
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricing) GetMappingOk() (*AiGatewayModelAutomaticPricingMapping, bool) {
	if o == nil || o.Mapping == nil {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *AiGatewayModelAutomaticPricing) HasMapping() bool {
	return o != nil && o.Mapping != nil
}

// SetMapping gets a reference to the given AiGatewayModelAutomaticPricingMapping and assigns it to the Mapping field.
func (o *AiGatewayModelAutomaticPricing) SetMapping(v AiGatewayModelAutomaticPricingMapping) {
	o.Mapping = &v
}

// GetModeTriggers returns the ModeTriggers field value if set, zero value otherwise.
func (o *AiGatewayModelAutomaticPricing) GetModeTriggers() []AiGatewayModelAutomaticPricingModeTrigger {
	if o == nil || o.ModeTriggers == nil {
		var ret []AiGatewayModelAutomaticPricingModeTrigger
		return ret
	}
	return o.ModeTriggers
}

// GetModeTriggersOk returns a tuple with the ModeTriggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricing) GetModeTriggersOk() (*[]AiGatewayModelAutomaticPricingModeTrigger, bool) {
	if o == nil || o.ModeTriggers == nil {
		return nil, false
	}
	return &o.ModeTriggers, true
}

// HasModeTriggers returns a boolean if a field has been set.
func (o *AiGatewayModelAutomaticPricing) HasModeTriggers() bool {
	return o != nil && o.ModeTriggers != nil
}

// SetModeTriggers gets a reference to the given []AiGatewayModelAutomaticPricingModeTrigger and assigns it to the ModeTriggers field.
func (o *AiGatewayModelAutomaticPricing) SetModeTriggers(v []AiGatewayModelAutomaticPricingModeTrigger) {
	o.ModeTriggers = v
}

// GetObservedAt returns the ObservedAt field value.
func (o *AiGatewayModelAutomaticPricing) GetObservedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricing) GetObservedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObservedAt, true
}

// SetObservedAt sets field value.
func (o *AiGatewayModelAutomaticPricing) SetObservedAt(v time.Time) {
	o.ObservedAt = v
}

// GetRates returns the Rates field value.
func (o *AiGatewayModelAutomaticPricing) GetRates() []AiGatewayModelAutomaticPricingRate {
	if o == nil {
		var ret []AiGatewayModelAutomaticPricingRate
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricing) GetRatesOk() (*[]AiGatewayModelAutomaticPricingRate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rates, true
}

// SetRates sets field value.
func (o *AiGatewayModelAutomaticPricing) SetRates(v []AiGatewayModelAutomaticPricingRate) {
	o.Rates = v
}

// GetSource returns the Source field value.
func (o *AiGatewayModelAutomaticPricing) GetSource() AiGatewayModelAutomaticPricingSource {
	if o == nil {
		var ret AiGatewayModelAutomaticPricingSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricing) GetSourceOk() (*AiGatewayModelAutomaticPricingSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *AiGatewayModelAutomaticPricing) SetSource(v AiGatewayModelAutomaticPricingSource) {
	o.Source = v
}

// GetStale returns the Stale field value.
func (o *AiGatewayModelAutomaticPricing) GetStale() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Stale
}

// GetStaleOk returns a tuple with the Stale field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricing) GetStaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stale, true
}

// SetStale sets field value.
func (o *AiGatewayModelAutomaticPricing) SetStale(v bool) {
	o.Stale = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelAutomaticPricing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LastValidatedAt.Nanosecond() == 0 {
		toSerialize["last_validated_at"] = o.LastValidatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["last_validated_at"] = o.LastValidatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Mapping != nil {
		toSerialize["mapping"] = o.Mapping
	}
	if o.ModeTriggers != nil {
		toSerialize["mode_triggers"] = o.ModeTriggers
	}
	if o.ObservedAt.Nanosecond() == 0 {
		toSerialize["observed_at"] = o.ObservedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["observed_at"] = o.ObservedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["rates"] = o.Rates
	toSerialize["source"] = o.Source
	toSerialize["stale"] = o.Stale

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelAutomaticPricing) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LastValidatedAt *time.Time                                  `json:"last_validated_at"`
		Mapping         *AiGatewayModelAutomaticPricingMapping      `json:"mapping,omitempty"`
		ModeTriggers    []AiGatewayModelAutomaticPricingModeTrigger `json:"mode_triggers,omitempty"`
		ObservedAt      *time.Time                                  `json:"observed_at"`
		Rates           *[]AiGatewayModelAutomaticPricingRate       `json:"rates"`
		Source          *AiGatewayModelAutomaticPricingSource       `json:"source"`
		Stale           *bool                                       `json:"stale"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LastValidatedAt == nil {
		return fmt.Errorf("required field last_validated_at missing")
	}
	if all.ObservedAt == nil {
		return fmt.Errorf("required field observed_at missing")
	}
	if all.Rates == nil {
		return fmt.Errorf("required field rates missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Stale == nil {
		return fmt.Errorf("required field stale missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"last_validated_at", "mapping", "mode_triggers", "observed_at", "rates", "source", "stale"})
	} else {
		return err
	}

	hasInvalidField := false
	o.LastValidatedAt = *all.LastValidatedAt
	if all.Mapping != nil && all.Mapping.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Mapping = all.Mapping
	o.ModeTriggers = all.ModeTriggers
	o.ObservedAt = *all.ObservedAt
	o.Rates = *all.Rates
	if all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = *all.Source
	o.Stale = *all.Stale

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
