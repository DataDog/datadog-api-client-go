// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudCostCollectionIntegrationDataflowSettingsResponse Settings of the Cloud Cost Management dataflow.
type NebiusCloudCostCollectionIntegrationDataflowSettingsResponse struct {
	// ID of the Nebius Object Storage access key used to read the cost export bucket.
	AccessKeyId *string `json:"access_key_id,omitempty"`
	// Name of the Nebius Object Storage bucket holding the cost export.
	BucketName *string `json:"bucket_name,omitempty"`
	// Prefix path within the bucket that locates the cost export.
	BucketPrefixPath *string `json:"bucket_prefix_path,omitempty"`
	// Region of the Nebius Object Storage bucket, such as `eu-north1`.
	BucketRegion *string `json:"bucket_region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewNebiusCloudCostCollectionIntegrationDataflowSettingsResponse instantiates a new NebiusCloudCostCollectionIntegrationDataflowSettingsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudCostCollectionIntegrationDataflowSettingsResponse() *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse {
	this := NebiusCloudCostCollectionIntegrationDataflowSettingsResponse{}
	return &this
}

// NewNebiusCloudCostCollectionIntegrationDataflowSettingsResponseWithDefaults instantiates a new NebiusCloudCostCollectionIntegrationDataflowSettingsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudCostCollectionIntegrationDataflowSettingsResponseWithDefaults() *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse {
	this := NebiusCloudCostCollectionIntegrationDataflowSettingsResponse{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) HasAccessKeyId() bool {
	return o != nil && o.AccessKeyId != nil
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) HasBucketName() bool {
	return o != nil && o.BucketName != nil
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) SetBucketName(v string) {
	o.BucketName = &v
}

// GetBucketPrefixPath returns the BucketPrefixPath field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetBucketPrefixPath() string {
	if o == nil || o.BucketPrefixPath == nil {
		var ret string
		return ret
	}
	return *o.BucketPrefixPath
}

// GetBucketPrefixPathOk returns a tuple with the BucketPrefixPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetBucketPrefixPathOk() (*string, bool) {
	if o == nil || o.BucketPrefixPath == nil {
		return nil, false
	}
	return o.BucketPrefixPath, true
}

// HasBucketPrefixPath returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) HasBucketPrefixPath() bool {
	return o != nil && o.BucketPrefixPath != nil
}

// SetBucketPrefixPath gets a reference to the given string and assigns it to the BucketPrefixPath field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) SetBucketPrefixPath(v string) {
	o.BucketPrefixPath = &v
}

// GetBucketRegion returns the BucketRegion field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetBucketRegion() string {
	if o == nil || o.BucketRegion == nil {
		var ret string
		return ret
	}
	return *o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) GetBucketRegionOk() (*string, bool) {
	if o == nil || o.BucketRegion == nil {
		return nil, false
	}
	return o.BucketRegion, true
}

// HasBucketRegion returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) HasBucketRegion() bool {
	return o != nil && o.BucketRegion != nil
}

// SetBucketRegion gets a reference to the given string and assigns it to the BucketRegion field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) SetBucketRegion(v string) {
	o.BucketRegion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessKeyId != nil {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.BucketPrefixPath != nil {
		toSerialize["bucket_prefix_path"] = o.BucketPrefixPath
	}
	if o.BucketRegion != nil {
		toSerialize["bucket_region"] = o.BucketRegion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKeyId      *string `json:"access_key_id,omitempty"`
		BucketName       *string `json:"bucket_name,omitempty"`
		BucketPrefixPath *string `json:"bucket_prefix_path,omitempty"`
		BucketRegion     *string `json:"bucket_region,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_key_id", "bucket_name", "bucket_prefix_path", "bucket_region"})
	} else {
		return err
	}
	o.AccessKeyId = all.AccessKeyId
	o.BucketName = all.BucketName
	o.BucketPrefixPath = all.BucketPrefixPath
	o.BucketRegion = all.BucketRegion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
