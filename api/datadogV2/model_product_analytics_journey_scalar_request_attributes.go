// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyScalarRequestAttributes Attributes of a journey scalar request.
type ProductAnalyticsJourneyScalarRequestAttributes struct {
	// Start of the query window, in epoch milliseconds.
	From int64 `json:"from"`
	// Query definition for a journey scalar request.
	Query ProductAnalyticsJourneyScalarQuery `json:"query"`
	// End of the query window, in epoch milliseconds.
	To int64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyScalarRequestAttributes instantiates a new ProductAnalyticsJourneyScalarRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyScalarRequestAttributes(from int64, query ProductAnalyticsJourneyScalarQuery, to int64) *ProductAnalyticsJourneyScalarRequestAttributes {
	this := ProductAnalyticsJourneyScalarRequestAttributes{}
	this.From = from
	this.Query = query
	this.To = to
	return &this
}

// NewProductAnalyticsJourneyScalarRequestAttributesWithDefaults instantiates a new ProductAnalyticsJourneyScalarRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyScalarRequestAttributesWithDefaults() *ProductAnalyticsJourneyScalarRequestAttributes {
	this := ProductAnalyticsJourneyScalarRequestAttributes{}
	return &this
}

// GetFrom returns the From field value.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) GetFrom() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) GetFromOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) SetFrom(v int64) {
	o.From = v
}

// GetQuery returns the Query field value.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) GetQuery() ProductAnalyticsJourneyScalarQuery {
	if o == nil {
		var ret ProductAnalyticsJourneyScalarQuery
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) GetQueryOk() (*ProductAnalyticsJourneyScalarQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) SetQuery(v ProductAnalyticsJourneyScalarQuery) {
	o.Query = v
}

// GetTo returns the To field value.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) GetTo() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) GetToOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *ProductAnalyticsJourneyScalarRequestAttributes) SetTo(v int64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyScalarRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *ProductAnalyticsJourneyScalarRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From  *int64                              `json:"from"`
		Query *ProductAnalyticsJourneyScalarQuery `json:"query"`
		To    *int64                              `json:"to"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"from", "query", "to"})
	} else {
		return err
	}

	hasInvalidField := false
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
