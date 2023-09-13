// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// S3FallbackDestination The S3 archive destination.
type S3FallbackDestination struct {
	// The bucket where the archive will be stored.
	Bucket string `json:"bucket"`
	// The S3 Archive's integration destination.
	Integration S3FallbackDestinationIntegration `json:"integration"`
	// The archive path.
	Path *string `json:"path,omitempty"`
	// Type of the S3 archive destination.
	Type S3FallbackDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewS3FallbackDestination instantiates a new S3FallbackDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewS3FallbackDestination(bucket string, integration S3FallbackDestinationIntegration, typeVar S3FallbackDestinationType) *S3FallbackDestination {
	this := S3FallbackDestination{}
	this.Bucket = bucket
	this.Integration = integration
	this.Type = typeVar
	return &this
}

// NewS3FallbackDestinationWithDefaults instantiates a new S3FallbackDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewS3FallbackDestinationWithDefaults() *S3FallbackDestination {
	this := S3FallbackDestination{}
	var typeVar S3FallbackDestinationType = S3FALLBACKDESTINATIONTYPE_S3
	this.Type = typeVar
	return &this
}

// GetBucket returns the Bucket field value.
func (o *S3FallbackDestination) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *S3FallbackDestination) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *S3FallbackDestination) SetBucket(v string) {
	o.Bucket = v
}

// GetIntegration returns the Integration field value.
func (o *S3FallbackDestination) GetIntegration() S3FallbackDestinationIntegration {
	if o == nil {
		var ret S3FallbackDestinationIntegration
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *S3FallbackDestination) GetIntegrationOk() (*S3FallbackDestinationIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *S3FallbackDestination) SetIntegration(v S3FallbackDestinationIntegration) {
	o.Integration = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *S3FallbackDestination) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3FallbackDestination) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *S3FallbackDestination) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *S3FallbackDestination) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value.
func (o *S3FallbackDestination) GetType() S3FallbackDestinationType {
	if o == nil {
		var ret S3FallbackDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *S3FallbackDestination) GetTypeOk() (*S3FallbackDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *S3FallbackDestination) SetType(v S3FallbackDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o S3FallbackDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["bucket"] = o.Bucket
	toSerialize["integration"] = o.Integration
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *S3FallbackDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bucket      *string                           `json:"bucket"`
		Integration *S3FallbackDestinationIntegration `json:"integration"`
		Path        *string                           `json:"path,omitempty"`
		Type        *S3FallbackDestinationType        `json:"type"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Bucket == nil {
		return fmt.Errorf("required field bucket missing")
	}
	if all.Integration == nil {
		return fmt.Errorf("required field integration missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket", "integration", "path", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Bucket = *all.Bucket
	if all.Integration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Integration = *all.Integration
	o.Path = all.Path
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
