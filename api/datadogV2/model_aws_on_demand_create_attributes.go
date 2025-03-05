// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsOnDemandCreateAttributes Attributes for the AWS on demand task.
type AwsOnDemandCreateAttributes struct {
	// The arn of the resource to scan. Agentless supports the scan of EC2 instances, lambda functions, AMI, ECR, RDS and S3 buckets.
	Arn *string `json:"arn,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsOnDemandCreateAttributes instantiates a new AwsOnDemandCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsOnDemandCreateAttributes() *AwsOnDemandCreateAttributes {
	this := AwsOnDemandCreateAttributes{}
	return &this
}

// NewAwsOnDemandCreateAttributesWithDefaults instantiates a new AwsOnDemandCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsOnDemandCreateAttributesWithDefaults() *AwsOnDemandCreateAttributes {
	this := AwsOnDemandCreateAttributes{}
	return &this
}

// GetArn returns the Arn field value if set, zero value otherwise.
func (o *AwsOnDemandCreateAttributes) GetArn() string {
	if o == nil || o.Arn == nil {
		var ret string
		return ret
	}
	return *o.Arn
}

// GetArnOk returns a tuple with the Arn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsOnDemandCreateAttributes) GetArnOk() (*string, bool) {
	if o == nil || o.Arn == nil {
		return nil, false
	}
	return o.Arn, true
}

// HasArn returns a boolean if a field has been set.
func (o *AwsOnDemandCreateAttributes) HasArn() bool {
	return o != nil && o.Arn != nil
}

// SetArn gets a reference to the given string and assigns it to the Arn field.
func (o *AwsOnDemandCreateAttributes) SetArn(v string) {
	o.Arn = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsOnDemandCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arn != nil {
		toSerialize["arn"] = o.Arn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AwsOnDemandCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arn *string `json:"arn,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arn"})
	} else {
		return err
	}
	o.Arn = all.Arn

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
