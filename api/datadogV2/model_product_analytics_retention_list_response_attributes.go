// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsRetentionListResponseAttributes Attributes of a retention list response, containing the matching entity rows.
type ProductAnalyticsRetentionListResponseAttributes struct {
	// The matching entity rows.
	Records []map[string]interface{} `json:"records,omitempty"`
	// The entity whose retention was measured.
	RetentionEntity *string `json:"retention_entity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsRetentionListResponseAttributes instantiates a new ProductAnalyticsRetentionListResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsRetentionListResponseAttributes() *ProductAnalyticsRetentionListResponseAttributes {
	this := ProductAnalyticsRetentionListResponseAttributes{}
	return &this
}

// NewProductAnalyticsRetentionListResponseAttributesWithDefaults instantiates a new ProductAnalyticsRetentionListResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsRetentionListResponseAttributesWithDefaults() *ProductAnalyticsRetentionListResponseAttributes {
	this := ProductAnalyticsRetentionListResponseAttributes{}
	return &this
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionListResponseAttributes) GetRecords() []map[string]interface{} {
	if o == nil || o.Records == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionListResponseAttributes) GetRecordsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return &o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionListResponseAttributes) HasRecords() bool {
	return o != nil && o.Records != nil
}

// SetRecords gets a reference to the given []map[string]interface{} and assigns it to the Records field.
func (o *ProductAnalyticsRetentionListResponseAttributes) SetRecords(v []map[string]interface{}) {
	o.Records = v
}

// GetRetentionEntity returns the RetentionEntity field value if set, zero value otherwise.
func (o *ProductAnalyticsRetentionListResponseAttributes) GetRetentionEntity() string {
	if o == nil || o.RetentionEntity == nil {
		var ret string
		return ret
	}
	return *o.RetentionEntity
}

// GetRetentionEntityOk returns a tuple with the RetentionEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsRetentionListResponseAttributes) GetRetentionEntityOk() (*string, bool) {
	if o == nil || o.RetentionEntity == nil {
		return nil, false
	}
	return o.RetentionEntity, true
}

// HasRetentionEntity returns a boolean if a field has been set.
func (o *ProductAnalyticsRetentionListResponseAttributes) HasRetentionEntity() bool {
	return o != nil && o.RetentionEntity != nil
}

// SetRetentionEntity gets a reference to the given string and assigns it to the RetentionEntity field.
func (o *ProductAnalyticsRetentionListResponseAttributes) SetRetentionEntity(v string) {
	o.RetentionEntity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsRetentionListResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	if o.RetentionEntity != nil {
		toSerialize["retention_entity"] = o.RetentionEntity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsRetentionListResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Records         []map[string]interface{} `json:"records,omitempty"`
		RetentionEntity *string                  `json:"retention_entity,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"records", "retention_entity"})
	} else {
		return err
	}
	o.Records = all.Records
	o.RetentionEntity = all.RetentionEntity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
