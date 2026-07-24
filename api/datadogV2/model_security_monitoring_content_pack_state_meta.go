// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringContentPackStateMeta Metadata for content pack states.
type SecurityMonitoringContentPackStateMeta struct {
	// Whether the Cloud SIEM index configuration is incorrect for the organization.
	CloudSiemIndexIncorrect bool `json:"cloud_siem_index_incorrect"`
	// The number of months that standard logs are retained for organizations on the standalone_indexed` pricing model. This field is omitted for other pricing models.
	RetentionMonths *int32 `json:"retention_months,omitempty"`
	// The Cloud SIEM pricing model (SKU) for the organization.
	Sku SecurityMonitoringSKU `json:"sku"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringContentPackStateMeta instantiates a new SecurityMonitoringContentPackStateMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringContentPackStateMeta(cloudSiemIndexIncorrect bool, sku SecurityMonitoringSKU) *SecurityMonitoringContentPackStateMeta {
	this := SecurityMonitoringContentPackStateMeta{}
	this.CloudSiemIndexIncorrect = cloudSiemIndexIncorrect
	this.Sku = sku
	return &this
}

// NewSecurityMonitoringContentPackStateMetaWithDefaults instantiates a new SecurityMonitoringContentPackStateMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringContentPackStateMetaWithDefaults() *SecurityMonitoringContentPackStateMeta {
	this := SecurityMonitoringContentPackStateMeta{}
	return &this
}

// GetCloudSiemIndexIncorrect returns the CloudSiemIndexIncorrect field value.
func (o *SecurityMonitoringContentPackStateMeta) GetCloudSiemIndexIncorrect() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.CloudSiemIndexIncorrect
}

// GetCloudSiemIndexIncorrectOk returns a tuple with the CloudSiemIndexIncorrect field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateMeta) GetCloudSiemIndexIncorrectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudSiemIndexIncorrect, true
}

// SetCloudSiemIndexIncorrect sets field value.
func (o *SecurityMonitoringContentPackStateMeta) SetCloudSiemIndexIncorrect(v bool) {
	o.CloudSiemIndexIncorrect = v
}

// GetRetentionMonths returns the RetentionMonths field value if set, zero value otherwise.
func (o *SecurityMonitoringContentPackStateMeta) GetRetentionMonths() int32 {
	if o == nil || o.RetentionMonths == nil {
		var ret int32
		return ret
	}
	return *o.RetentionMonths
}

// GetRetentionMonthsOk returns a tuple with the RetentionMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateMeta) GetRetentionMonthsOk() (*int32, bool) {
	if o == nil || o.RetentionMonths == nil {
		return nil, false
	}
	return o.RetentionMonths, true
}

// HasRetentionMonths returns a boolean if a field has been set.
func (o *SecurityMonitoringContentPackStateMeta) HasRetentionMonths() bool {
	return o != nil && o.RetentionMonths != nil
}

// SetRetentionMonths gets a reference to the given int32 and assigns it to the RetentionMonths field.
func (o *SecurityMonitoringContentPackStateMeta) SetRetentionMonths(v int32) {
	o.RetentionMonths = &v
}

// GetSku returns the Sku field value.
func (o *SecurityMonitoringContentPackStateMeta) GetSku() SecurityMonitoringSKU {
	if o == nil {
		var ret SecurityMonitoringSKU
		return ret
	}
	return o.Sku
}

// GetSkuOk returns a tuple with the Sku field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringContentPackStateMeta) GetSkuOk() (*SecurityMonitoringSKU, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sku, true
}

// SetSku sets field value.
func (o *SecurityMonitoringContentPackStateMeta) SetSku(v SecurityMonitoringSKU) {
	o.Sku = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringContentPackStateMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cloud_siem_index_incorrect"] = o.CloudSiemIndexIncorrect
	if o.RetentionMonths != nil {
		toSerialize["retention_months"] = o.RetentionMonths
	}
	toSerialize["sku"] = o.Sku

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringContentPackStateMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudSiemIndexIncorrect *bool                  `json:"cloud_siem_index_incorrect"`
		RetentionMonths         *int32                 `json:"retention_months,omitempty"`
		Sku                     *SecurityMonitoringSKU `json:"sku"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CloudSiemIndexIncorrect == nil {
		return fmt.Errorf("required field cloud_siem_index_incorrect missing")
	}
	if all.Sku == nil {
		return fmt.Errorf("required field sku missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud_siem_index_incorrect", "retention_months", "sku"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CloudSiemIndexIncorrect = *all.CloudSiemIndexIncorrect
	o.RetentionMonths = all.RetentionMonths
	if !all.Sku.IsValid() {
		hasInvalidField = true
	} else {
		o.Sku = *all.Sku
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
