// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerTargetGroupAttributes Attributes of a scanning group of a target.
type SensitiveDataScannerTargetGroupAttributes struct {
	// Description of the group.
	Description string `json:"description"`
	// Filter for the Scanning Group.
	Filter SensitiveDataScannerFilter `json:"filter"`
	// Whether or not the group is enabled.
	IsEnabled bool `json:"is_enabled"`
	// Name of the group.
	Name string `json:"name"`
	// Products the scanning group applies to.
	ProductList []SensitiveDataScannerGroupTarget `json:"product_list"`
	// Sampling rates per product type. Always empty for these targets.
	Samplings []SensitiveDataScannerSamplings `json:"samplings"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerTargetGroupAttributes instantiates a new SensitiveDataScannerTargetGroupAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerTargetGroupAttributes(description string, filter SensitiveDataScannerFilter, isEnabled bool, name string, productList []SensitiveDataScannerGroupTarget, samplings []SensitiveDataScannerSamplings) *SensitiveDataScannerTargetGroupAttributes {
	this := SensitiveDataScannerTargetGroupAttributes{}
	this.Description = description
	this.Filter = filter
	this.IsEnabled = isEnabled
	this.Name = name
	this.ProductList = productList
	this.Samplings = samplings
	return &this
}

// NewSensitiveDataScannerTargetGroupAttributesWithDefaults instantiates a new SensitiveDataScannerTargetGroupAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerTargetGroupAttributesWithDefaults() *SensitiveDataScannerTargetGroupAttributes {
	this := SensitiveDataScannerTargetGroupAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetDescription(v string) {
	o.Description = v
}

// GetFilter returns the Filter field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetFilter() SensitiveDataScannerFilter {
	if o == nil {
		var ret SensitiveDataScannerFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetFilterOk() (*SensitiveDataScannerFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetFilter(v SensitiveDataScannerFilter) {
	o.Filter = v
}

// GetIsEnabled returns the IsEnabled field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetName returns the Name field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetName(v string) {
	o.Name = v
}

// GetProductList returns the ProductList field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetProductList() []SensitiveDataScannerGroupTarget {
	if o == nil {
		var ret []SensitiveDataScannerGroupTarget
		return ret
	}
	return o.ProductList
}

// GetProductListOk returns a tuple with the ProductList field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetProductListOk() (*[]SensitiveDataScannerGroupTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductList, true
}

// SetProductList sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetProductList(v []SensitiveDataScannerGroupTarget) {
	o.ProductList = v
}

// GetSamplings returns the Samplings field value.
func (o *SensitiveDataScannerTargetGroupAttributes) GetSamplings() []SensitiveDataScannerSamplings {
	if o == nil {
		var ret []SensitiveDataScannerSamplings
		return ret
	}
	return o.Samplings
}

// GetSamplingsOk returns a tuple with the Samplings field value
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTargetGroupAttributes) GetSamplingsOk() (*[]SensitiveDataScannerSamplings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Samplings, true
}

// SetSamplings sets field value.
func (o *SensitiveDataScannerTargetGroupAttributes) SetSamplings(v []SensitiveDataScannerSamplings) {
	o.Samplings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerTargetGroupAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["filter"] = o.Filter
	toSerialize["is_enabled"] = o.IsEnabled
	toSerialize["name"] = o.Name
	toSerialize["product_list"] = o.ProductList
	toSerialize["samplings"] = o.Samplings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerTargetGroupAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                            `json:"description"`
		Filter      *SensitiveDataScannerFilter        `json:"filter"`
		IsEnabled   *bool                              `json:"is_enabled"`
		Name        *string                            `json:"name"`
		ProductList *[]SensitiveDataScannerGroupTarget `json:"product_list"`
		Samplings   *[]SensitiveDataScannerSamplings   `json:"samplings"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	if all.IsEnabled == nil {
		return fmt.Errorf("required field is_enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ProductList == nil {
		return fmt.Errorf("required field product_list missing")
	}
	if all.Samplings == nil {
		return fmt.Errorf("required field samplings missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "filter", "is_enabled", "name", "product_list", "samplings"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = *all.Description
	if all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = *all.Filter
	o.IsEnabled = *all.IsEnabled
	o.Name = *all.Name
	o.ProductList = *all.ProductList
	o.Samplings = *all.Samplings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
