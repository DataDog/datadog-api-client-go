// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRecoveryBatchCreateAttributes Recover sessions selected by a query. The range cannot exceed 180 days; from cannot be more than 4 years ago and to cannot be more than 1 hour in the future. Session availability depends on the configured archive retention.
type RumRecoveryBatchCreateAttributes struct {
	// Optional client estimate. Creation validates and stores a fresh server estimate instead of this value.
	EstimatedCount *int64 `json:"estimated_count,omitempty"`
	// Start of the range in Unix milliseconds. Must be positive and earlier than to.
	From int64 `json:"from"`
	// RUM search query. Not blank, at most 4096 bytes for batch creation and count.
	Query string `json:"query"`
	// Query-based batch recovery.
	RequestType RumRecoveryBatchRequestType `json:"request_type"`
	// End of the range in Unix milliseconds.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRecoveryBatchCreateAttributes instantiates a new RumRecoveryBatchCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRecoveryBatchCreateAttributes(from int64, query string, requestType RumRecoveryBatchRequestType, to int64) *RumRecoveryBatchCreateAttributes {
	this := RumRecoveryBatchCreateAttributes{}
	this.From = from
	this.Query = query
	this.RequestType = requestType
	this.To = to
	return &this
}

// NewRumRecoveryBatchCreateAttributesWithDefaults instantiates a new RumRecoveryBatchCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRecoveryBatchCreateAttributesWithDefaults() *RumRecoveryBatchCreateAttributes {
	this := RumRecoveryBatchCreateAttributes{}
	return &this
}

// GetEstimatedCount returns the EstimatedCount field value if set, zero value otherwise.
func (o *RumRecoveryBatchCreateAttributes) GetEstimatedCount() int64 {
	if o == nil || o.EstimatedCount == nil {
		var ret int64
		return ret
	}
	return *o.EstimatedCount
}

// GetEstimatedCountOk returns a tuple with the EstimatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchCreateAttributes) GetEstimatedCountOk() (*int64, bool) {
	if o == nil || o.EstimatedCount == nil {
		return nil, false
	}
	return o.EstimatedCount, true
}

// HasEstimatedCount returns a boolean if a field has been set.
func (o *RumRecoveryBatchCreateAttributes) HasEstimatedCount() bool {
	return o != nil && o.EstimatedCount != nil
}

// SetEstimatedCount gets a reference to the given int64 and assigns it to the EstimatedCount field.
func (o *RumRecoveryBatchCreateAttributes) SetEstimatedCount(v int64) {
	o.EstimatedCount = &v
}

// GetFrom returns the From field value.
func (o *RumRecoveryBatchCreateAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchCreateAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *RumRecoveryBatchCreateAttributes) SetFrom(v int64) {
	o.From = v
}

// GetQuery returns the Query field value.
func (o *RumRecoveryBatchCreateAttributes) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchCreateAttributes) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *RumRecoveryBatchCreateAttributes) SetQuery(v string) {
	o.Query = v
}

// GetRequestType returns the RequestType field value.
func (o *RumRecoveryBatchCreateAttributes) GetRequestType() RumRecoveryBatchRequestType {
	if o == nil {
		var ret RumRecoveryBatchRequestType
		return ret
	}
	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchCreateAttributes) GetRequestTypeOk() (*RumRecoveryBatchRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value.
func (o *RumRecoveryBatchCreateAttributes) SetRequestType(v RumRecoveryBatchRequestType) {
	o.RequestType = v
}

// GetTo returns the To field value.
func (o *RumRecoveryBatchCreateAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RumRecoveryBatchCreateAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *RumRecoveryBatchCreateAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRecoveryBatchCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EstimatedCount != nil {
		toSerialize["estimated_count"] = o.EstimatedCount
	}
	toSerialize["from"] = o.From
	toSerialize["query"] = o.Query
	toSerialize["request_type"] = o.RequestType
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumRecoveryBatchCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EstimatedCount *int64                       `json:"estimated_count,omitempty"`
		From           *int64                       `json:"from"`
		Query          *string                      `json:"query"`
		RequestType    *RumRecoveryBatchRequestType `json:"request_type"`
		To             *int64                       `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.RequestType == nil {
		return fmt.Errorf("required field request_type missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"estimated_count", "from", "query", "request_type", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EstimatedCount = all.EstimatedCount
	o.From = *all.From
	o.Query = *all.Query
	if !all.RequestType.IsValid() {
		hasInvalidField = true
	} else {
		o.RequestType = *all.RequestType
	}
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
