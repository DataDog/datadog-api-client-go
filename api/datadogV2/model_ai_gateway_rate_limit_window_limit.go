// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiGatewayRateLimitWindowLimit A limit on one metric over a rolling window longer than a minute.
type AiGatewayRateLimitWindowLimit struct {
	// Maximum amount of the metric over the window. For `cost`, the amount is priced spend in microcents of 1e-8 USD each, as in `cost_microcents_per_minute`. For the same metric, a longer window must have a larger amount.
	Amount int64 `json:"amount"`
	// The metric that a window limit counts. `cost` counts priced spend in microcents.
	Metric AiGatewayRateLimitWindowLimitMetric `json:"metric"`
	// The rolling window that a limit applies to. For per-minute limits, use the `*_per_minute` fields instead.
	Window AiGatewayRateLimitWindowLimitWindow `json:"window"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiGatewayRateLimitWindowLimit instantiates a new AiGatewayRateLimitWindowLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiGatewayRateLimitWindowLimit(amount int64, metric AiGatewayRateLimitWindowLimitMetric, window AiGatewayRateLimitWindowLimitWindow) *AiGatewayRateLimitWindowLimit {
	this := AiGatewayRateLimitWindowLimit{}
	this.Amount = amount
	this.Metric = metric
	this.Window = window
	return &this
}

// NewAiGatewayRateLimitWindowLimitWithDefaults instantiates a new AiGatewayRateLimitWindowLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiGatewayRateLimitWindowLimitWithDefaults() *AiGatewayRateLimitWindowLimit {
	this := AiGatewayRateLimitWindowLimit{}
	return &this
}

// GetAmount returns the Amount field value.
func (o *AiGatewayRateLimitWindowLimit) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitWindowLimit) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value.
func (o *AiGatewayRateLimitWindowLimit) SetAmount(v int64) {
	o.Amount = v
}

// GetMetric returns the Metric field value.
func (o *AiGatewayRateLimitWindowLimit) GetMetric() AiGatewayRateLimitWindowLimitMetric {
	if o == nil {
		var ret AiGatewayRateLimitWindowLimitMetric
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitWindowLimit) GetMetricOk() (*AiGatewayRateLimitWindowLimitMetric, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *AiGatewayRateLimitWindowLimit) SetMetric(v AiGatewayRateLimitWindowLimitMetric) {
	o.Metric = v
}

// GetWindow returns the Window field value.
func (o *AiGatewayRateLimitWindowLimit) GetWindow() AiGatewayRateLimitWindowLimitWindow {
	if o == nil {
		var ret AiGatewayRateLimitWindowLimitWindow
		return ret
	}
	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *AiGatewayRateLimitWindowLimit) GetWindowOk() (*AiGatewayRateLimitWindowLimitWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value.
func (o *AiGatewayRateLimitWindowLimit) SetWindow(v AiGatewayRateLimitWindowLimitWindow) {
	o.Window = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiGatewayRateLimitWindowLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["amount"] = o.Amount
	toSerialize["metric"] = o.Metric
	toSerialize["window"] = o.Window

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiGatewayRateLimitWindowLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Amount *int64                               `json:"amount"`
		Metric *AiGatewayRateLimitWindowLimitMetric `json:"metric"`
		Window *AiGatewayRateLimitWindowLimitWindow `json:"window"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Amount == nil {
		return fmt.Errorf("required field amount missing")
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	if all.Window == nil {
		return fmt.Errorf("required field window missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"amount", "metric", "window"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Amount = *all.Amount
	if !all.Metric.IsValid() {
		hasInvalidField = true
	} else {
		o.Metric = *all.Metric
	}
	if !all.Window.IsValid() {
		hasInvalidField = true
	} else {
		o.Window = *all.Window
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
