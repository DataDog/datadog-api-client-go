// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformBackendAttributes Terraform backend sync configuration and bucket statuses.
type TerraformBackendAttributes struct {
	// AWS account ID that owns the S3 buckets.
	AccountId string `json:"account_id"`
	// Backend type to synchronize.
	BackendType TerraformBackendKind `json:"backend_type"`
	// Source buckets and their synchronization statuses.
	Buckets []TerraformBackendBucket `json:"buckets"`
	// Datadog organization ID.
	OrgId string `json:"org_id"`
	// AWS region containing the S3 buckets.
	Region string `json:"region"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTerraformBackendAttributes instantiates a new TerraformBackendAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformBackendAttributes(accountId string, backendType TerraformBackendKind, buckets []TerraformBackendBucket, orgId string, region string) *TerraformBackendAttributes {
	this := TerraformBackendAttributes{}
	this.AccountId = accountId
	this.BackendType = backendType
	this.Buckets = buckets
	this.OrgId = orgId
	this.Region = region
	return &this
}

// NewTerraformBackendAttributesWithDefaults instantiates a new TerraformBackendAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformBackendAttributesWithDefaults() *TerraformBackendAttributes {
	this := TerraformBackendAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *TerraformBackendAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *TerraformBackendAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetBackendType returns the BackendType field value.
func (o *TerraformBackendAttributes) GetBackendType() TerraformBackendKind {
	if o == nil {
		var ret TerraformBackendKind
		return ret
	}
	return o.BackendType
}

// GetBackendTypeOk returns a tuple with the BackendType field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendAttributes) GetBackendTypeOk() (*TerraformBackendKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendType, true
}

// SetBackendType sets field value.
func (o *TerraformBackendAttributes) SetBackendType(v TerraformBackendKind) {
	o.BackendType = v
}

// GetBuckets returns the Buckets field value.
func (o *TerraformBackendAttributes) GetBuckets() []TerraformBackendBucket {
	if o == nil {
		var ret []TerraformBackendBucket
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendAttributes) GetBucketsOk() (*[]TerraformBackendBucket, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Buckets, true
}

// SetBuckets sets field value.
func (o *TerraformBackendAttributes) SetBuckets(v []TerraformBackendBucket) {
	o.Buckets = v
}

// GetOrgId returns the OrgId field value.
func (o *TerraformBackendAttributes) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendAttributes) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *TerraformBackendAttributes) SetOrgId(v string) {
	o.OrgId = v
}

// GetRegion returns the Region field value.
func (o *TerraformBackendAttributes) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendAttributes) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *TerraformBackendAttributes) SetRegion(v string) {
	o.Region = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformBackendAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["backend_type"] = o.BackendType
	toSerialize["buckets"] = o.Buckets
	toSerialize["org_id"] = o.OrgId
	toSerialize["region"] = o.Region

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformBackendAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId   *string                   `json:"account_id"`
		BackendType *TerraformBackendKind     `json:"backend_type"`
		Buckets     *[]TerraformBackendBucket `json:"buckets"`
		OrgId       *string                   `json:"org_id"`
		Region      *string                   `json:"region"`
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
	if all.Buckets == nil {
		return fmt.Errorf("required field buckets missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "backend_type", "buckets", "org_id", "region"})
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
	o.Buckets = *all.Buckets
	o.OrgId = *all.OrgId
	o.Region = *all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
