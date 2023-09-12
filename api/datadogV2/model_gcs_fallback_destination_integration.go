// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCSFallbackDestinationIntegration The GCS archive's integration destination.
type GCSFallbackDestinationIntegration struct {
	// A client email.
	ClientEmail string `json:"clientEmail"`
	// A project ID.
	ProjectId string `json:"projectId"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGCSFallbackDestinationIntegration instantiates a new GCSFallbackDestinationIntegration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCSFallbackDestinationIntegration(clientEmail string, projectId string) *GCSFallbackDestinationIntegration {
	this := GCSFallbackDestinationIntegration{}
	this.ClientEmail = clientEmail
	this.ProjectId = projectId
	return &this
}

// NewGCSFallbackDestinationIntegrationWithDefaults instantiates a new GCSFallbackDestinationIntegration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCSFallbackDestinationIntegrationWithDefaults() *GCSFallbackDestinationIntegration {
	this := GCSFallbackDestinationIntegration{}
	return &this
}

// GetClientEmail returns the ClientEmail field value.
func (o *GCSFallbackDestinationIntegration) GetClientEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value
// and a boolean to check if the value has been set.
func (o *GCSFallbackDestinationIntegration) GetClientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientEmail, true
}

// SetClientEmail sets field value.
func (o *GCSFallbackDestinationIntegration) SetClientEmail(v string) {
	o.ClientEmail = v
}

// GetProjectId returns the ProjectId field value.
func (o *GCSFallbackDestinationIntegration) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GCSFallbackDestinationIntegration) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *GCSFallbackDestinationIntegration) SetProjectId(v string) {
	o.ProjectId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCSFallbackDestinationIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["clientEmail"] = o.ClientEmail
	toSerialize["projectId"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCSFallbackDestinationIntegration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientEmail *string `json:"clientEmail"`
		ProjectId   *string `json:"projectId"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientEmail == nil {
		return fmt.Errorf("required field clientEmail missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field projectId missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"clientEmail", "projectId"})
	} else {
		return err
	}
	o.ClientEmail = *all.ClientEmail
	o.ProjectId = *all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
