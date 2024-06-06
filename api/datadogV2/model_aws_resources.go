// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSResources AWS Resources config
type AWSResources struct {
	// Whether Datadog collects cloud security posture management resources from your AWS account.
	CloudSecurityPostureManagementCollection *bool `json:"cloud_security_posture_management_collection,omitempty"`
	// Whether Datadog collects additional attributes and configuration information about the resources in your AWS account. Required for `cspm_resource_collection`.
	ExtendedCollection *bool `json:"extended_collection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSResources instantiates a new AWSResources object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSResources() *AWSResources {
	this := AWSResources{}
	return &this
}

// NewAWSResourcesWithDefaults instantiates a new AWSResources object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSResourcesWithDefaults() *AWSResources {
	this := AWSResources{}
	return &this
}

// GetCloudSecurityPostureManagementCollection returns the CloudSecurityPostureManagementCollection field value if set, zero value otherwise.
func (o *AWSResources) GetCloudSecurityPostureManagementCollection() bool {
	if o == nil || o.CloudSecurityPostureManagementCollection == nil {
		var ret bool
		return ret
	}
	return *o.CloudSecurityPostureManagementCollection
}

// GetCloudSecurityPostureManagementCollectionOk returns a tuple with the CloudSecurityPostureManagementCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSResources) GetCloudSecurityPostureManagementCollectionOk() (*bool, bool) {
	if o == nil || o.CloudSecurityPostureManagementCollection == nil {
		return nil, false
	}
	return o.CloudSecurityPostureManagementCollection, true
}

// HasCloudSecurityPostureManagementCollection returns a boolean if a field has been set.
func (o *AWSResources) HasCloudSecurityPostureManagementCollection() bool {
	return o != nil && o.CloudSecurityPostureManagementCollection != nil
}

// SetCloudSecurityPostureManagementCollection gets a reference to the given bool and assigns it to the CloudSecurityPostureManagementCollection field.
func (o *AWSResources) SetCloudSecurityPostureManagementCollection(v bool) {
	o.CloudSecurityPostureManagementCollection = &v
}

// GetExtendedCollection returns the ExtendedCollection field value if set, zero value otherwise.
func (o *AWSResources) GetExtendedCollection() bool {
	if o == nil || o.ExtendedCollection == nil {
		var ret bool
		return ret
	}
	return *o.ExtendedCollection
}

// GetExtendedCollectionOk returns a tuple with the ExtendedCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSResources) GetExtendedCollectionOk() (*bool, bool) {
	if o == nil || o.ExtendedCollection == nil {
		return nil, false
	}
	return o.ExtendedCollection, true
}

// HasExtendedCollection returns a boolean if a field has been set.
func (o *AWSResources) HasExtendedCollection() bool {
	return o != nil && o.ExtendedCollection != nil
}

// SetExtendedCollection gets a reference to the given bool and assigns it to the ExtendedCollection field.
func (o *AWSResources) SetExtendedCollection(v bool) {
	o.ExtendedCollection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CloudSecurityPostureManagementCollection != nil {
		toSerialize["cloud_security_posture_management_collection"] = o.CloudSecurityPostureManagementCollection
	}
	if o.ExtendedCollection != nil {
		toSerialize["extended_collection"] = o.ExtendedCollection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSResources) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudSecurityPostureManagementCollection *bool `json:"cloud_security_posture_management_collection,omitempty"`
		ExtendedCollection                       *bool `json:"extended_collection,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud_security_posture_management_collection", "extended_collection"})
	} else {
		return err
	}
	o.CloudSecurityPostureManagementCollection = all.CloudSecurityPostureManagementCollection
	o.ExtendedCollection = all.ExtendedCollection

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
