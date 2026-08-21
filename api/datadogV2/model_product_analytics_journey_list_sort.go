// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsJourneyListSort Sort configuration for the returned rows. The sort is applied only when `facet`
// is one of the returned columns; otherwise it is ignored.
type ProductAnalyticsJourneyListSort struct {
	// Column to sort on.
	Facet *string `json:"facet,omitempty"`
	// Direction of sort.
	Order *QuerySortOrder `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsJourneyListSort instantiates a new ProductAnalyticsJourneyListSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsJourneyListSort() *ProductAnalyticsJourneyListSort {
	this := ProductAnalyticsJourneyListSort{}
	var order QuerySortOrder = QUERYSORTORDER_DESC
	this.Order = &order
	return &this
}

// NewProductAnalyticsJourneyListSortWithDefaults instantiates a new ProductAnalyticsJourneyListSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsJourneyListSortWithDefaults() *ProductAnalyticsJourneyListSort {
	this := ProductAnalyticsJourneyListSort{}
	var order QuerySortOrder = QUERYSORTORDER_DESC
	this.Order = &order
	return &this
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListSort) GetFacet() string {
	if o == nil || o.Facet == nil {
		var ret string
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListSort) GetFacetOk() (*string, bool) {
	if o == nil || o.Facet == nil {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListSort) HasFacet() bool {
	return o != nil && o.Facet != nil
}

// SetFacet gets a reference to the given string and assigns it to the Facet field.
func (o *ProductAnalyticsJourneyListSort) SetFacet(v string) {
	o.Facet = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ProductAnalyticsJourneyListSort) GetOrder() QuerySortOrder {
	if o == nil || o.Order == nil {
		var ret QuerySortOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsJourneyListSort) GetOrderOk() (*QuerySortOrder, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ProductAnalyticsJourneyListSort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given QuerySortOrder and assigns it to the Order field.
func (o *ProductAnalyticsJourneyListSort) SetOrder(v QuerySortOrder) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsJourneyListSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Facet != nil {
		toSerialize["facet"] = o.Facet
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsJourneyListSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Facet *string         `json:"facet,omitempty"`
		Order *QuerySortOrder `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet", "order"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Facet = all.Facet
	if all.Order != nil && !all.Order.IsValid() {
		hasInvalidField = true
	} else {
		o.Order = all.Order
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
