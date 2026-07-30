// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ElasticCloudCcmSettingsUpdate Partial Elastic Cloud CCM interface settings for updates.
type ElasticCloudCcmSettingsUpdate struct {
	// Your Elastic Cloud organization ID, found in your organization settings.
	ElasticOrgId *string `json:"elastic_org_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticCloudCcmSettingsUpdate instantiates a new ElasticCloudCcmSettingsUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticCloudCcmSettingsUpdate() *ElasticCloudCcmSettingsUpdate {
	this := ElasticCloudCcmSettingsUpdate{}
	return &this
}

// NewElasticCloudCcmSettingsUpdateWithDefaults instantiates a new ElasticCloudCcmSettingsUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticCloudCcmSettingsUpdateWithDefaults() *ElasticCloudCcmSettingsUpdate {
	this := ElasticCloudCcmSettingsUpdate{}
	return &this
}

// GetElasticOrgId returns the ElasticOrgId field value if set, zero value otherwise.
func (o *ElasticCloudCcmSettingsUpdate) GetElasticOrgId() string {
	if o == nil || o.ElasticOrgId == nil {
		var ret string
		return ret
	}
	return *o.ElasticOrgId
}

// GetElasticOrgIdOk returns a tuple with the ElasticOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticCloudCcmSettingsUpdate) GetElasticOrgIdOk() (*string, bool) {
	if o == nil || o.ElasticOrgId == nil {
		return nil, false
	}
	return o.ElasticOrgId, true
}

// HasElasticOrgId returns a boolean if a field has been set.
func (o *ElasticCloudCcmSettingsUpdate) HasElasticOrgId() bool {
	return o != nil && o.ElasticOrgId != nil
}

// SetElasticOrgId gets a reference to the given string and assigns it to the ElasticOrgId field.
func (o *ElasticCloudCcmSettingsUpdate) SetElasticOrgId(v string) {
	o.ElasticOrgId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticCloudCcmSettingsUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ElasticOrgId != nil {
		toSerialize["elastic_org_id"] = o.ElasticOrgId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticCloudCcmSettingsUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ElasticOrgId *string `json:"elastic_org_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"elastic_org_id"})
	} else {
		return err
	}
	o.ElasticOrgId = all.ElasticOrgId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
