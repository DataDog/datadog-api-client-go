// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TerraformBackendUpdateAttributes Replacement bucket set for a Terraform backend sync configuration.
type TerraformBackendUpdateAttributes struct {
	// Complete set of S3 bucket names to synchronize. Names must be nonempty and unique.
	BucketNames []string `json:"bucket_names"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTerraformBackendUpdateAttributes instantiates a new TerraformBackendUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTerraformBackendUpdateAttributes(bucketNames []string) *TerraformBackendUpdateAttributes {
	this := TerraformBackendUpdateAttributes{}
	this.BucketNames = bucketNames
	return &this
}

// NewTerraformBackendUpdateAttributesWithDefaults instantiates a new TerraformBackendUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTerraformBackendUpdateAttributesWithDefaults() *TerraformBackendUpdateAttributes {
	this := TerraformBackendUpdateAttributes{}
	return &this
}

// GetBucketNames returns the BucketNames field value.
func (o *TerraformBackendUpdateAttributes) GetBucketNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BucketNames
}

// GetBucketNamesOk returns a tuple with the BucketNames field value
// and a boolean to check if the value has been set.
func (o *TerraformBackendUpdateAttributes) GetBucketNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketNames, true
}

// SetBucketNames sets field value.
func (o *TerraformBackendUpdateAttributes) SetBucketNames(v []string) {
	o.BucketNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TerraformBackendUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bucket_names"] = o.BucketNames

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TerraformBackendUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketNames *[]string `json:"bucket_names"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BucketNames == nil {
		return fmt.Errorf("required field bucket_names missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket_names"})
	} else {
		return err
	}
	o.BucketNames = *all.BucketNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
