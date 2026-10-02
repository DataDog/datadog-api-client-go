// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems Exposure count and identity of one experiment variant.
type ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems struct {
	// Number of recorded exposures for this variant.
	ExposureCount *int64 `json:"exposure_count,omitempty"`
	// Key that identifies the experiment variant.
	VariantKey *string `json:"variant_key,omitempty"`
	// Display name of the experiment variant.
	VariantName *string `json:"variant_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewExperimentsTrafficSummaryV2DTODataAttributesVariantsItems instantiates a new ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewExperimentsTrafficSummaryV2DTODataAttributesVariantsItems() *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems {
	this := ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems{}
	return &this
}

// NewExperimentsTrafficSummaryV2DTODataAttributesVariantsItemsWithDefaults instantiates a new ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewExperimentsTrafficSummaryV2DTODataAttributesVariantsItemsWithDefaults() *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems {
	this := ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems{}
	return &this
}

// GetExposureCount returns the ExposureCount field value if set, zero value otherwise.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) GetExposureCount() int64 {
	if o == nil || o.ExposureCount == nil {
		var ret int64
		return ret
	}
	return *o.ExposureCount
}

// GetExposureCountOk returns a tuple with the ExposureCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) GetExposureCountOk() (*int64, bool) {
	if o == nil || o.ExposureCount == nil {
		return nil, false
	}
	return o.ExposureCount, true
}

// HasExposureCount returns a boolean if a field has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) HasExposureCount() bool {
	return o != nil && o.ExposureCount != nil
}

// SetExposureCount gets a reference to the given int64 and assigns it to the ExposureCount field.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) SetExposureCount(v int64) {
	o.ExposureCount = &v
}

// GetVariantKey returns the VariantKey field value if set, zero value otherwise.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) GetVariantKey() string {
	if o == nil || o.VariantKey == nil {
		var ret string
		return ret
	}
	return *o.VariantKey
}

// GetVariantKeyOk returns a tuple with the VariantKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) GetVariantKeyOk() (*string, bool) {
	if o == nil || o.VariantKey == nil {
		return nil, false
	}
	return o.VariantKey, true
}

// HasVariantKey returns a boolean if a field has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) HasVariantKey() bool {
	return o != nil && o.VariantKey != nil
}

// SetVariantKey gets a reference to the given string and assigns it to the VariantKey field.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) SetVariantKey(v string) {
	o.VariantKey = &v
}

// GetVariantName returns the VariantName field value if set, zero value otherwise.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) GetVariantName() string {
	if o == nil || o.VariantName == nil {
		var ret string
		return ret
	}
	return *o.VariantName
}

// GetVariantNameOk returns a tuple with the VariantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) GetVariantNameOk() (*string, bool) {
	if o == nil || o.VariantName == nil {
		return nil, false
	}
	return o.VariantName, true
}

// HasVariantName returns a boolean if a field has been set.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) HasVariantName() bool {
	return o != nil && o.VariantName != nil
}

// SetVariantName gets a reference to the given string and assigns it to the VariantName field.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) SetVariantName(v string) {
	o.VariantName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ExposureCount != nil {
		toSerialize["exposure_count"] = o.ExposureCount
	}
	if o.VariantKey != nil {
		toSerialize["variant_key"] = o.VariantKey
	}
	if o.VariantName != nil {
		toSerialize["variant_name"] = o.VariantName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ExperimentsTrafficSummaryV2DTODataAttributesVariantsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExposureCount *int64  `json:"exposure_count,omitempty"`
		VariantKey    *string `json:"variant_key,omitempty"`
		VariantName   *string `json:"variant_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exposure_count", "variant_key", "variant_name"})
	} else {
		return err
	}
	o.ExposureCount = all.ExposureCount
	o.VariantKey = all.VariantKey
	o.VariantName = all.VariantName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
