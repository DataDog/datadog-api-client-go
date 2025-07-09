// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseGoogleSecurityOperationsDestinationAuth Google Security Operations destination authentication.
type CustomDestinationResponseGoogleSecurityOperationsDestinationAuth struct {
	// The Google Security Operations client email.
	ClientEmail string `json:"client_email"`
	// Google Security Operations project ID.
	ProjectId string `json:"project_id"`
	// Type of the Google Security Operations destination authentication.
	Type CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationResponseGoogleSecurityOperationsDestinationAuth instantiates a new CustomDestinationResponseGoogleSecurityOperationsDestinationAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationResponseGoogleSecurityOperationsDestinationAuth(clientEmail string, projectId string, typeVar CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType) *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth {
	this := CustomDestinationResponseGoogleSecurityOperationsDestinationAuth{}
	this.ClientEmail = clientEmail
	this.ProjectId = projectId
	this.Type = typeVar
	return &this
}

// NewCustomDestinationResponseGoogleSecurityOperationsDestinationAuthWithDefaults instantiates a new CustomDestinationResponseGoogleSecurityOperationsDestinationAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationResponseGoogleSecurityOperationsDestinationAuthWithDefaults() *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth {
	this := CustomDestinationResponseGoogleSecurityOperationsDestinationAuth{}
	var typeVar CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType = CUSTOMDESTINATIONRESPONSEGOOGLESECURITYOPERATIONSDESTINATIONAUTHTYPE_GCP_PRIVATE_KEY
	this.Type = typeVar
	return &this
}

// GetClientEmail returns the ClientEmail field value.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) GetClientEmail() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientEmail
}

// GetClientEmailOk returns a tuple with the ClientEmail field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) GetClientEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientEmail, true
}

// SetClientEmail sets field value.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) SetClientEmail(v string) {
	o.ClientEmail = v
}

// GetProjectId returns the ProjectId field value.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) SetProjectId(v string) {
	o.ProjectId = v
}

// GetType returns the Type field value.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) GetType() CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType {
	if o == nil {
		var ret CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) GetTypeOk() (*CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) SetType(v CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["client_email"] = o.ClientEmail
	toSerialize["project_id"] = o.ProjectId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationResponseGoogleSecurityOperationsDestinationAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ClientEmail *string                                                               `json:"client_email"`
		ProjectId   *string                                                               `json:"project_id"`
		Type        *CustomDestinationResponseGoogleSecurityOperationsDestinationAuthType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ClientEmail == nil {
		return fmt.Errorf("required field client_email missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"client_email", "project_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ClientEmail = *all.ClientEmail
	o.ProjectId = *all.ProjectId
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
