// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// NebiusCloudCostCollectionIntegrationDataflowSettingsRequest Settings of the Cloud Cost Management dataflow. Only the fields provided are changed.
type NebiusCloudCostCollectionIntegrationDataflowSettingsRequest struct {
	// ID of the Nebius Object Storage (S3-compatible) access key used to read the cost export bucket. Required when this dataflow is enabled.
	AccessKeyId *string `json:"access_key_id,omitempty"`
	// Secret access key for the access key ID. Required when this dataflow is enabled. Omit it on update to keep the stored one; an empty string is rejected with a `422` response.
	AccessKeySecret *string `json:"access_key_secret,omitempty"`
	// Name of the Nebius Object Storage bucket holding your cost export. Required when this dataflow is enabled; a request that enables it without the bucket fields and access keys is rejected with a `422` response.
	BucketName *string `json:"bucket_name,omitempty"`
	// Prefix path within the bucket that locates the cost export, combining the export name and path prefix.
	BucketPrefixPath *string `json:"bucket_prefix_path,omitempty"`
	// Region of the Nebius Object Storage bucket, such as `eu-north1`. Required when this dataflow is enabled.
	BucketRegion *string `json:"bucket_region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewNebiusCloudCostCollectionIntegrationDataflowSettingsRequest instantiates a new NebiusCloudCostCollectionIntegrationDataflowSettingsRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewNebiusCloudCostCollectionIntegrationDataflowSettingsRequest() *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest {
	this := NebiusCloudCostCollectionIntegrationDataflowSettingsRequest{}
	return &this
}

// NewNebiusCloudCostCollectionIntegrationDataflowSettingsRequestWithDefaults instantiates a new NebiusCloudCostCollectionIntegrationDataflowSettingsRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewNebiusCloudCostCollectionIntegrationDataflowSettingsRequestWithDefaults() *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest {
	this := NebiusCloudCostCollectionIntegrationDataflowSettingsRequest{}
	return &this
}

// GetAccessKeyId returns the AccessKeyId field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetAccessKeyId() string {
	if o == nil || o.AccessKeyId == nil {
		var ret string
		return ret
	}
	return *o.AccessKeyId
}

// GetAccessKeyIdOk returns a tuple with the AccessKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetAccessKeyIdOk() (*string, bool) {
	if o == nil || o.AccessKeyId == nil {
		return nil, false
	}
	return o.AccessKeyId, true
}

// HasAccessKeyId returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) HasAccessKeyId() bool {
	return o != nil && o.AccessKeyId != nil
}

// SetAccessKeyId gets a reference to the given string and assigns it to the AccessKeyId field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) SetAccessKeyId(v string) {
	o.AccessKeyId = &v
}

// GetAccessKeySecret returns the AccessKeySecret field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetAccessKeySecret() string {
	if o == nil || o.AccessKeySecret == nil {
		var ret string
		return ret
	}
	return *o.AccessKeySecret
}

// GetAccessKeySecretOk returns a tuple with the AccessKeySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetAccessKeySecretOk() (*string, bool) {
	if o == nil || o.AccessKeySecret == nil {
		return nil, false
	}
	return o.AccessKeySecret, true
}

// HasAccessKeySecret returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) HasAccessKeySecret() bool {
	return o != nil && o.AccessKeySecret != nil
}

// SetAccessKeySecret gets a reference to the given string and assigns it to the AccessKeySecret field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) SetAccessKeySecret(v string) {
	o.AccessKeySecret = &v
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) HasBucketName() bool {
	return o != nil && o.BucketName != nil
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) SetBucketName(v string) {
	o.BucketName = &v
}

// GetBucketPrefixPath returns the BucketPrefixPath field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetBucketPrefixPath() string {
	if o == nil || o.BucketPrefixPath == nil {
		var ret string
		return ret
	}
	return *o.BucketPrefixPath
}

// GetBucketPrefixPathOk returns a tuple with the BucketPrefixPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetBucketPrefixPathOk() (*string, bool) {
	if o == nil || o.BucketPrefixPath == nil {
		return nil, false
	}
	return o.BucketPrefixPath, true
}

// HasBucketPrefixPath returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) HasBucketPrefixPath() bool {
	return o != nil && o.BucketPrefixPath != nil
}

// SetBucketPrefixPath gets a reference to the given string and assigns it to the BucketPrefixPath field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) SetBucketPrefixPath(v string) {
	o.BucketPrefixPath = &v
}

// GetBucketRegion returns the BucketRegion field value if set, zero value otherwise.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetBucketRegion() string {
	if o == nil || o.BucketRegion == nil {
		var ret string
		return ret
	}
	return *o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) GetBucketRegionOk() (*string, bool) {
	if o == nil || o.BucketRegion == nil {
		return nil, false
	}
	return o.BucketRegion, true
}

// HasBucketRegion returns a boolean if a field has been set.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) HasBucketRegion() bool {
	return o != nil && o.BucketRegion != nil
}

// SetBucketRegion gets a reference to the given string and assigns it to the BucketRegion field.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) SetBucketRegion(v string) {
	o.BucketRegion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessKeyId != nil {
		toSerialize["access_key_id"] = o.AccessKeyId
	}
	if o.AccessKeySecret != nil {
		toSerialize["access_key_secret"] = o.AccessKeySecret
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
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *NebiusCloudCostCollectionIntegrationDataflowSettingsRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessKeyId      *string `json:"access_key_id,omitempty"`
		AccessKeySecret  *string `json:"access_key_secret,omitempty"`
		BucketName       *string `json:"bucket_name,omitempty"`
		BucketPrefixPath *string `json:"bucket_prefix_path,omitempty"`
		BucketRegion     *string `json:"bucket_region,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.AccessKeyId = all.AccessKeyId
	o.AccessKeySecret = all.AccessKeySecret
	o.BucketName = all.BucketName
	o.BucketPrefixPath = all.BucketPrefixPath
	o.BucketRegion = all.BucketRegion

	return nil
}
