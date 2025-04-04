// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PipelineQuotaProcessorLimit The maximum amount of data or number of events allowed before the quota is enforced. Can be specified in bytes or events.
type PipelineQuotaProcessorLimit struct {
	// Unit for quota enforcement in bytes for data size or events for count.
	Enforce PipelineQuotaProcessorLimitEnforceType `json:"enforce"`
	// The limit for quota enforcement.
	Limit int64 `json:"limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPipelineQuotaProcessorLimit instantiates a new PipelineQuotaProcessorLimit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPipelineQuotaProcessorLimit(enforce PipelineQuotaProcessorLimitEnforceType, limit int64) *PipelineQuotaProcessorLimit {
	this := PipelineQuotaProcessorLimit{}
	this.Enforce = enforce
	this.Limit = limit
	return &this
}

// NewPipelineQuotaProcessorLimitWithDefaults instantiates a new PipelineQuotaProcessorLimit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPipelineQuotaProcessorLimitWithDefaults() *PipelineQuotaProcessorLimit {
	this := PipelineQuotaProcessorLimit{}
	return &this
}

// GetEnforce returns the Enforce field value.
func (o *PipelineQuotaProcessorLimit) GetEnforce() PipelineQuotaProcessorLimitEnforceType {
	if o == nil {
		var ret PipelineQuotaProcessorLimitEnforceType
		return ret
	}
	return o.Enforce
}

// GetEnforceOk returns a tuple with the Enforce field value
// and a boolean to check if the value has been set.
func (o *PipelineQuotaProcessorLimit) GetEnforceOk() (*PipelineQuotaProcessorLimitEnforceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enforce, true
}

// SetEnforce sets field value.
func (o *PipelineQuotaProcessorLimit) SetEnforce(v PipelineQuotaProcessorLimitEnforceType) {
	o.Enforce = v
}

// GetLimit returns the Limit field value.
func (o *PipelineQuotaProcessorLimit) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *PipelineQuotaProcessorLimit) GetLimitOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *PipelineQuotaProcessorLimit) SetLimit(v int64) {
	o.Limit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PipelineQuotaProcessorLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enforce"] = o.Enforce
	toSerialize["limit"] = o.Limit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PipelineQuotaProcessorLimit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enforce *PipelineQuotaProcessorLimitEnforceType `json:"enforce"`
		Limit   *int64                                  `json:"limit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enforce == nil {
		return fmt.Errorf("required field enforce missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enforce", "limit"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Enforce.IsValid() {
		hasInvalidField = true
	} else {
		o.Enforce = *all.Enforce
	}
	o.Limit = *all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
