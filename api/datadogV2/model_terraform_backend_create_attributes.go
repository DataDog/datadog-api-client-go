// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformBackendCreateAttributes Settings for a new Terraform backend sync configuration.
type TerraformBackendCreateAttributes struct {
	// AWS account ID that owns the S3 buckets.
	AccountId string `json:"account_id"`
	// Backend type to synchronize.
	BackendType TerraformBackendKind `json:"backend_type"`
	// Complete set of S3 bucket names to synchronize. Names must be nonempty and unique.
	BucketNames []string `json:"bucket_names"`
	// AWS region containing the S3 buckets.
	Region string `json:"region"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTerraformBackendCreateAttributes instantiates a new TerraformBackendCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformBackendCreateAttributes(accountId string, backendType TerraformBackendKind, bucketNames []string, region string) *TerraformBackendCreateAttributes {
	this := TerraformBackendCreateAttributes{}
	this.AccountId = accountId
	this.BackendType = backendType
	this.BucketNames = bucketNames
	this.Region = region
	return &this
}

// NewTerraformBackendCreateAttributesWithDefaults instantiates a new TerraformBackendCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformBackendCreateAttributesWithDefaults() *TerraformBackendCreateAttributes {
	this := TerraformBackendCreateAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *TerraformBackendCreateAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendCreateAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *TerraformBackendCreateAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetBackendType returns the BackendType field value.
func (o *TerraformBackendCreateAttributes) GetBackendType() TerraformBackendKind {
	if o == nil {
		var ret TerraformBackendKind
		return ret
	}
	return o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendCreateAttributes) GetBackendTypeOk() (*TerraformBackendKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendType, true
}

// SetBackendType sets field value.
func (o *TerraformBackendCreateAttributes) SetBackendType(v TerraformBackendKind) {
	o.BackendType = v
}

// GetBucketNames returns the BucketNames field value.
func (o *TerraformBackendCreateAttributes) GetBucketNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BucketNames
}

// GetBucketNamesOk returns a tuple with the BucketNames field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendCreateAttributes) GetBucketNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketNames, true
}

// SetBucketNames sets field value.
func (o *TerraformBackendCreateAttributes) SetBucketNames(v []string) {
	o.BucketNames = v
}

// GetRegion returns the Region field value.
func (o *TerraformBackendCreateAttributes) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendCreateAttributes) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *TerraformBackendCreateAttributes) SetRegion(v string) {
	o.Region = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformBackendCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["backend_type"] = o.BackendType
	toSerialize["bucket_names"] = o.BucketNames
	toSerialize["region"] = o.Region

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformBackendCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId   *string               `json:"account_id"`
		BackendType *TerraformBackendKind `json:"backend_type"`
		BucketNames *[]string             `json:"bucket_names"`
		Region      *string               `json:"region"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.BackendType == nil {
		return fmt.Errorf("required field backend_type missing")
	}
	if all.BucketNames == nil {
		return fmt.Errorf("required field bucket_names missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "backend_type", "bucket_names", "region"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountId = *all.AccountId
	if !all.BackendType.IsValid() {
		hasInvalidField = true
	} else {
		o.BackendType = *all.BackendType
	}
	o.BucketNames = *all.BucketNames
	o.Region = *all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
