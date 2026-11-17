// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayModelAutomaticPricingModeTrigger Request fields that select a priced mode.
type AiGatewayModelAutomaticPricingModeTrigger struct {
	// Request body fields that must match.
	Body []AiGatewayModelAutomaticPricingFieldMatch `json:"body,omitempty"`
	// Request headers that must match.
	Headers []AiGatewayModelAutomaticPricingFieldMatch `json:"headers,omitempty"`
	// Priced mode selected by the fields.
	Mode string `json:"mode"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayModelAutomaticPricingModeTrigger instantiates a new AiGatewayModelAutomaticPricingModeTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayModelAutomaticPricingModeTrigger(mode string) *AiGatewayModelAutomaticPricingModeTrigger {
	this := AiGatewayModelAutomaticPricingModeTrigger{}
	this.Mode = mode
	return &this
}

// NewAiGatewayModelAutomaticPricingModeTriggerWithDefaults instantiates a new AiGatewayModelAutomaticPricingModeTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayModelAutomaticPricingModeTriggerWithDefaults() *AiGatewayModelAutomaticPricingModeTrigger {
	this := AiGatewayModelAutomaticPricingModeTrigger{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *AiGatewayModelAutomaticPricingModeTrigger) GetBody() []AiGatewayModelAutomaticPricingFieldMatch {
	if o == nil || o.Body == nil {
		var ret []AiGatewayModelAutomaticPricingFieldMatch
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingModeTrigger) GetBodyOk() (*[]AiGatewayModelAutomaticPricingFieldMatch, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return &o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *AiGatewayModelAutomaticPricingModeTrigger) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given []AiGatewayModelAutomaticPricingFieldMatch and assigns it to the Body field.
func (o *AiGatewayModelAutomaticPricingModeTrigger) SetBody(v []AiGatewayModelAutomaticPricingFieldMatch) {
	o.Body = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *AiGatewayModelAutomaticPricingModeTrigger) GetHeaders() []AiGatewayModelAutomaticPricingFieldMatch {
	if o == nil || o.Headers == nil {
		var ret []AiGatewayModelAutomaticPricingFieldMatch
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingModeTrigger) GetHeadersOk() (*[]AiGatewayModelAutomaticPricingFieldMatch, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *AiGatewayModelAutomaticPricingModeTrigger) HasHeaders() bool {
	return o != nil && o.Headers != nil
}

// SetHeaders gets a reference to the given []AiGatewayModelAutomaticPricingFieldMatch and assigns it to the Headers field.
func (o *AiGatewayModelAutomaticPricingModeTrigger) SetHeaders(v []AiGatewayModelAutomaticPricingFieldMatch) {
	o.Headers = v
}

// GetMode returns the Mode field value.
func (o *AiGatewayModelAutomaticPricingModeTrigger) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *AiGatewayModelAutomaticPricingModeTrigger) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value.
func (o *AiGatewayModelAutomaticPricingModeTrigger) SetMode(v string) {
	o.Mode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayModelAutomaticPricingModeTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	toSerialize["mode"] = o.Mode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayModelAutomaticPricingModeTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Body    []AiGatewayModelAutomaticPricingFieldMatch `json:"body,omitempty"`
		Headers []AiGatewayModelAutomaticPricingFieldMatch `json:"headers,omitempty"`
		Mode    *string                                    `json:"mode"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Mode == nil {
		return fmt.Errorf("required field mode missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"body", "headers", "mode"})
	} else {
		return err
	}
	o.Body = all.Body
	o.Headers = all.Headers
	o.Mode = *all.Mode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
