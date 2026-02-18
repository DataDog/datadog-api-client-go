// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTrigger Trigger a workflow from an Incident. For automatic triggering a handle must be configured and the workflow must be published.
type IncidentTrigger struct {
	// Defines a rate limit for a trigger.
	RateLimit *TriggerRateLimit `json:"rateLimit,omitempty"`
	// Version of the incident trigger.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTrigger instantiates a new IncidentTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTrigger() *IncidentTrigger {
	this := IncidentTrigger{}
	return &this
}

// NewIncidentTriggerWithDefaults instantiates a new IncidentTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTriggerWithDefaults() *IncidentTrigger {
	this := IncidentTrigger{}
	return &this
}

// GetRateLimit returns the RateLimit field value if set, zero value otherwise.
func (o *IncidentTrigger) GetRateLimit() TriggerRateLimit {
	if o == nil || o.RateLimit == nil {
		var ret TriggerRateLimit
		return ret
	}
	return *o.RateLimit
}

// GetRateLimitOk returns a tuple with the RateLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTrigger) GetRateLimitOk() (*TriggerRateLimit, bool) {
	if o == nil || o.RateLimit == nil {
		return nil, false
	}
	return o.RateLimit, true
}

// HasRateLimit returns a boolean if a field has been set.
func (o *IncidentTrigger) HasRateLimit() bool {
	return o != nil && o.RateLimit != nil
}

// SetRateLimit gets a reference to the given TriggerRateLimit and assigns it to the RateLimit field.
func (o *IncidentTrigger) SetRateLimit(v TriggerRateLimit) {
	o.RateLimit = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IncidentTrigger) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTrigger) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IncidentTrigger) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IncidentTrigger) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RateLimit != nil {
		toSerialize["rateLimit"] = o.RateLimit
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RateLimit *TriggerRateLimit `json:"rateLimit,omitempty"`
		Version   *string           `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rateLimit", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.RateLimit != nil && all.RateLimit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RateLimit = all.RateLimit
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
