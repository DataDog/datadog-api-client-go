// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetadataOrigin Where the metric comes from. Each value is the display label shown on the Metrics Summary page. The set of labels can grow as Datadog adds products. Each field is `null` when it isn't known.
type MetricMetadataOrigin struct {
	// Product that reports the metric, such as `Agent`.
	Product datadog.NullableString `json:"product"`
	// Finer-grained origin within the sub-product, such as `Cpu`.
	ProductDetail datadog.NullableString `json:"product_detail"`
	// Part of the product that reports the metric, such as `Integration`.
	SubProduct datadog.NullableString `json:"sub_product"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricMetadataOrigin instantiates a new MetricMetadataOrigin object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricMetadataOrigin(product datadog.NullableString, productDetail datadog.NullableString, subProduct datadog.NullableString) *MetricMetadataOrigin {
	this := MetricMetadataOrigin{}
	this.Product = product
	this.ProductDetail = productDetail
	this.SubProduct = subProduct
	return &this
}

// NewMetricMetadataOriginWithDefaults instantiates a new MetricMetadataOrigin object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricMetadataOriginWithDefaults() *MetricMetadataOrigin {
	this := MetricMetadataOrigin{}
	return &this
}

// GetProduct returns the Product field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataOrigin) GetProduct() string {
	if o == nil || o.Product.Get() == nil {
		var ret string
		return ret
	}
	return *o.Product.Get()
}

// GetProductOk returns a tuple with the Product field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataOrigin) GetProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Product.Get(), o.Product.IsSet()
}

// SetProduct sets field value.
func (o *MetricMetadataOrigin) SetProduct(v string) {
	o.Product.Set(&v)
}

// GetProductDetail returns the ProductDetail field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataOrigin) GetProductDetail() string {
	if o == nil || o.ProductDetail.Get() == nil {
		var ret string
		return ret
	}
	return *o.ProductDetail.Get()
}

// GetProductDetailOk returns a tuple with the ProductDetail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataOrigin) GetProductDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductDetail.Get(), o.ProductDetail.IsSet()
}

// SetProductDetail sets field value.
func (o *MetricMetadataOrigin) SetProductDetail(v string) {
	o.ProductDetail.Set(&v)
}

// GetSubProduct returns the SubProduct field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *MetricMetadataOrigin) GetSubProduct() string {
	if o == nil || o.SubProduct.Get() == nil {
		var ret string
		return ret
	}
	return *o.SubProduct.Get()
}

// GetSubProductOk returns a tuple with the SubProduct field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MetricMetadataOrigin) GetSubProductOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubProduct.Get(), o.SubProduct.IsSet()
}

// SetSubProduct sets field value.
func (o *MetricMetadataOrigin) SetSubProduct(v string) {
	o.SubProduct.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricMetadataOrigin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["product"] = o.Product.Get()
	toSerialize["product_detail"] = o.ProductDetail.Get()
	toSerialize["sub_product"] = o.SubProduct.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricMetadataOrigin) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Product       datadog.NullableString `json:"product"`
		ProductDetail datadog.NullableString `json:"product_detail"`
		SubProduct    datadog.NullableString `json:"sub_product"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Product.IsSet() {
		return fmt.Errorf("required field product missing")
	}
	if !all.ProductDetail.IsSet() {
		return fmt.Errorf("required field product_detail missing")
	}
	if !all.SubProduct.IsSet() {
		return fmt.Errorf("required field sub_product missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"product", "product_detail", "sub_product"})
	} else {
		return err
	}
	o.Product = all.Product
	o.ProductDetail = all.ProductDetail
	o.SubProduct = all.SubProduct

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
