// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionGridRequestAttributes Attributes of a retention grid request.
type ProductAnalyticsRetentionGridRequestAttributes struct {
	// Whether to exclude sessions that are not tied to an identified user.
	ExcludeAnonymousTraffic *bool `json:"exclude_anonymous_traffic,omitempty"`
	// Start of the query window, in epoch milliseconds.
	From int64 `json:"from"`
	// Query definition for a retention grid or retention metadata request.
	Query ProductAnalyticsRetentionGridQuery `json:"query"`
	// End of the query window, in epoch milliseconds.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionGridRequestAttributes instantiates a new ProductAnalyticsRetentionGridRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionGridRequestAttributes(from int64, query ProductAnalyticsRetentionGridQuery, to int64) *ProductAnalyticsRetentionGridRequestAttributes {
	this := ProductAnalyticsRetentionGridRequestAttributes{}
	var excludeAnonymousTraffic bool = false
	this.ExcludeAnonymousTraffic = &excludeAnonymousTraffic
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewProductAnalyticsRetentionGridRequestAttributesWithDefaults instantiates a new ProductAnalyticsRetentionGridRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionGridRequestAttributesWithDefaults() *ProductAnalyticsRetentionGridRequestAttributes {
	this := ProductAnalyticsRetentionGridRequestAttributes{}
	var excludeAnonymousTraffic bool = false
	this.ExcludeAnonymousTraffic = &excludeAnonymousTraffic
	return &this
}

// GetExcludeAnonymousTraffic returns the ExcludeAnonymousTraffic field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetExcludeAnonymousTraffic() bool {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		var ret bool
		return ret
	}
	return *o.ExcludeAnonymousTraffic
}

// GetExcludeAnonymousTrafficOk returns a tuple with the ExcludeAnonymousTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetExcludeAnonymousTrafficOk() (*bool, bool) {
	if o == nil || o.ExcludeAnonymousTraffic == nil {
		return nil, false
	}
	return o.ExcludeAnonymousTraffic, true
}

// HasExcludeAnonymousTraffic returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionGridRequestAttributes) HasExcludeAnonymousTraffic() bool {
	return o != nil && o.ExcludeAnonymousTraffic != nil
}

// SetExcludeAnonymousTraffic gets a reference to the given bool and assigns it to the ExcludeAnonymousTraffic field.
func (o *ProductAnalyticsRetentionGridRequestAttributes) SetExcludeAnonymousTraffic(v bool) {
	o.ExcludeAnonymousTraffic = &v
}

// GetFrom returns the From field value.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ProductAnalyticsRetentionGridRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetQuery() ProductAnalyticsRetentionGridQuery {
	if o == nil {
		var ret ProductAnalyticsRetentionGridQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetQueryOk() (*ProductAnalyticsRetentionGridQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsRetentionGridRequestAttributes) SetQuery(v ProductAnalyticsRetentionGridQuery) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionGridRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *ProductAnalyticsRetentionGridRequestAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionGridRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExcludeAnonymousTraffic != nil {
		toSerialize["exclude_anonymous_traffic"] = o.ExcludeAnonymousTraffic
	}
	toSerialize["from"] = o.From
	toSerialize["query"] = o.Query
	toSerialize["to"] = o.To

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionGridRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExcludeAnonymousTraffic *bool                               `json:"exclude_anonymous_traffic,omitempty"`
		From                    *int64                              `json:"from"`
		Query                   *ProductAnalyticsRetentionGridQuery `json:"query"`
		To                      *int64                              `json:"to"`
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
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclude_anonymous_traffic", "from", "query", "to"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExcludeAnonymousTraffic = all.ExcludeAnonymousTraffic
	o.From = *all.From
	if all.Query.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Query = *all.Query
	o.To = *all.To

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
