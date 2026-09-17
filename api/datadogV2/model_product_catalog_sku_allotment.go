// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductCatalogSKUAllotment A quantity of one SKU that is included with, and consumed before, the billable usage of
// another SKU.
type ProductCatalogSKUAllotment struct {
	// The code of the SKU that receives the allotment.
	ChildSkuCode string `json:"child_sku_code"`
	// The quantity allotted per hour. Fractional for some allotments, and equal to
	// `monthly_quantity` for others, depending on how the child SKU meters usage.
	HourlyQuantity float64 `json:"hourly_quantity"`
	// The quantity allotted per month.
	MonthlyQuantity int64 `json:"monthly_quantity"`
	// The code of the SKU that provides the allotment. Always the code of the SKU the
	// allotment is returned under.
	ParentSkuCode string `json:"parent_sku_code"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductCatalogSKUAllotment instantiates a new ProductCatalogSKUAllotment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductCatalogSKUAllotment(childSkuCode string, hourlyQuantity float64, monthlyQuantity int64, parentSkuCode string) *ProductCatalogSKUAllotment {
	this := ProductCatalogSKUAllotment{}
	this.ChildSkuCode = childSkuCode
	this.HourlyQuantity = hourlyQuantity
	this.MonthlyQuantity = monthlyQuantity
	this.ParentSkuCode = parentSkuCode
	return &this
}

// NewProductCatalogSKUAllotmentWithDefaults instantiates a new ProductCatalogSKUAllotment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductCatalogSKUAllotmentWithDefaults() *ProductCatalogSKUAllotment {
	this := ProductCatalogSKUAllotment{}
	return &this
}

// GetChildSkuCode returns the ChildSkuCode field value.
func (o *ProductCatalogSKUAllotment) GetChildSkuCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ChildSkuCode
}

// GetChildSkuCodeOk returns a tuple with the ChildSkuCode field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUAllotment) GetChildSkuCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildSkuCode, true
}

// SetChildSkuCode sets field value.
func (o *ProductCatalogSKUAllotment) SetChildSkuCode(v string) {
	o.ChildSkuCode = v
}

// GetHourlyQuantity returns the HourlyQuantity field value.
func (o *ProductCatalogSKUAllotment) GetHourlyQuantity() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.HourlyQuantity
}

// GetHourlyQuantityOk returns a tuple with the HourlyQuantity field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUAllotment) GetHourlyQuantityOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HourlyQuantity, true
}

// SetHourlyQuantity sets field value.
func (o *ProductCatalogSKUAllotment) SetHourlyQuantity(v float64) {
	o.HourlyQuantity = v
}

// GetMonthlyQuantity returns the MonthlyQuantity field value.
func (o *ProductCatalogSKUAllotment) GetMonthlyQuantity() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MonthlyQuantity
}

// GetMonthlyQuantityOk returns a tuple with the MonthlyQuantity field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUAllotment) GetMonthlyQuantityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonthlyQuantity, true
}

// SetMonthlyQuantity sets field value.
func (o *ProductCatalogSKUAllotment) SetMonthlyQuantity(v int64) {
	o.MonthlyQuantity = v
}

// GetParentSkuCode returns the ParentSkuCode field value.
func (o *ProductCatalogSKUAllotment) GetParentSkuCode() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentSkuCode
}

// GetParentSkuCodeOk returns a tuple with the ParentSkuCode field value
// and a boolean to check if the value has been set.
func (o *ProductCatalogSKUAllotment) GetParentSkuCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentSkuCode, true
}

// SetParentSkuCode sets field value.
func (o *ProductCatalogSKUAllotment) SetParentSkuCode(v string) {
	o.ParentSkuCode = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductCatalogSKUAllotment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["child_sku_code"] = o.ChildSkuCode
	toSerialize["hourly_quantity"] = o.HourlyQuantity
	toSerialize["monthly_quantity"] = o.MonthlyQuantity
	toSerialize["parent_sku_code"] = o.ParentSkuCode

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductCatalogSKUAllotment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChildSkuCode    *string  `json:"child_sku_code"`
		HourlyQuantity  *float64 `json:"hourly_quantity"`
		MonthlyQuantity *int64   `json:"monthly_quantity"`
		ParentSkuCode   *string  `json:"parent_sku_code"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChildSkuCode == nil {
		return fmt.Errorf("required field child_sku_code missing")
	}
	if all.HourlyQuantity == nil {
		return fmt.Errorf("required field hourly_quantity missing")
	}
	if all.MonthlyQuantity == nil {
		return fmt.Errorf("required field monthly_quantity missing")
	}
	if all.ParentSkuCode == nil {
		return fmt.Errorf("required field parent_sku_code missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"child_sku_code", "hourly_quantity", "monthly_quantity", "parent_sku_code"})
	} else {
		return err
	}
	o.ChildSkuCode = *all.ChildSkuCode
	o.HourlyQuantity = *all.HourlyQuantity
	o.MonthlyQuantity = *all.MonthlyQuantity
	o.ParentSkuCode = *all.ParentSkuCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
