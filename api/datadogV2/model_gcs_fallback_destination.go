// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCSFallbackDestination The GCS archive destination.
type GCSFallbackDestination struct {
	// The bucket where the archive will be stored.
	Bucket string `json:"bucket"`
	// The GCS archive's integration destination.
	Integration GCSFallbackDestinationIntegration `json:"integration"`
	// The archive path.
	Path *string `json:"path,omitempty"`
	// Type of the GCS archive destination.
	Type GCSFallbackDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGCSFallbackDestination instantiates a new GCSFallbackDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCSFallbackDestination(bucket string, integration GCSFallbackDestinationIntegration, typeVar GCSFallbackDestinationType) *GCSFallbackDestination {
	this := GCSFallbackDestination{}
	this.Bucket = bucket
	this.Integration = integration
	this.Type = typeVar
	return &this
}

// NewGCSFallbackDestinationWithDefaults instantiates a new GCSFallbackDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCSFallbackDestinationWithDefaults() *GCSFallbackDestination {
	this := GCSFallbackDestination{}
	var typeVar GCSFallbackDestinationType = GCSFALLBACKDESTINATIONTYPE_GCS
	this.Type = typeVar
	return &this
}

// GetBucket returns the Bucket field value.
func (o *GCSFallbackDestination) GetBucket() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value
// and a boolean to check if the value has been set.
func (o *GCSFallbackDestination) GetBucketOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bucket, true
}

// SetBucket sets field value.
func (o *GCSFallbackDestination) SetBucket(v string) {
	o.Bucket = v
}

// GetIntegration returns the Integration field value.
func (o *GCSFallbackDestination) GetIntegration() GCSFallbackDestinationIntegration {
	if o == nil {
		var ret GCSFallbackDestinationIntegration
		return ret
	}
	return o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value
// and a boolean to check if the value has been set.
func (o *GCSFallbackDestination) GetIntegrationOk() (*GCSFallbackDestinationIntegration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Integration, true
}

// SetIntegration sets field value.
func (o *GCSFallbackDestination) SetIntegration(v GCSFallbackDestinationIntegration) {
	o.Integration = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GCSFallbackDestination) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCSFallbackDestination) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GCSFallbackDestination) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GCSFallbackDestination) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value.
func (o *GCSFallbackDestination) GetType() GCSFallbackDestinationType {
	if o == nil {
		var ret GCSFallbackDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GCSFallbackDestination) GetTypeOk() (*GCSFallbackDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *GCSFallbackDestination) SetType(v GCSFallbackDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCSFallbackDestination) MarshalJSON() ([]byte, error) {
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
func (o *GCSFallbackDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Bucket      *string                            `json:"bucket"`
		Integration *GCSFallbackDestinationIntegration `json:"integration"`
		Path        *string                            `json:"path,omitempty"`
		Type        *GCSFallbackDestinationType        `json:"type"`
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
