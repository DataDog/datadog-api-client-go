// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveIntegrationS3AccessKey The S3 Archive's integration destination using an access key.
type LogsArchiveIntegrationS3AccessKey struct {
	// The access key ID for the integration.
	AccessKeyId string `json:"access_key_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveIntegrationS3AccessKey instantiates a new LogsArchiveIntegrationS3AccessKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveIntegrationS3AccessKey(accessKeyId string) *LogsArchiveIntegrationS3AccessKey {
	this := LogsArchiveIntegrationS3AccessKey{}
	this.AccessKeyId = accessKeyId
	return &this
}

// NewLogsArchiveIntegrationS3AccessKeyWithDefaults instantiates a new LogsArchiveIntegrationS3AccessKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveIntegrationS3AccessKeyWithDefaults() *LogsArchiveIntegrationS3AccessKey {
	this := LogsArchiveIntegrationS3AccessKey{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value.
func (o *LogsArchiveIntegrationS3AccessKey) GetAccessKeyId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveIntegrationS3AccessKey) GetAccessKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessKeyId, true
}

// SetAccessKeyId sets field value.
func (o *LogsArchiveIntegrationS3AccessKey) SetAccessKeyId(v string) {
	o.AccessKeyId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveIntegrationS3AccessKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_key_id"] = o.AccessKeyId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveIntegrationS3AccessKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKeyId *string `json:"access_key_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessKeyId == nil {
		return fmt.Errorf("required field access_key_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_key_id"})
	} else {
		return err
	}
	o.AccessKeyId = *all.AccessKeyId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
