// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntakeAPIKeyAttributes Attributes of an intake API key returned after successful authentication.
type IntakeAPIKeyAttributes struct {
	// The Datadog API key the workload can use to send telemetry.
	ApiKey string `json:"api_key"`
	// The numeric ID of the Datadog organization the API key belongs to.
	OrgId int64 `json:"org_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntakeAPIKeyAttributes instantiates a new IntakeAPIKeyAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntakeAPIKeyAttributes(apiKey string, orgId int64) *IntakeAPIKeyAttributes {
	this := IntakeAPIKeyAttributes{}
	this.ApiKey = apiKey
	this.OrgId = orgId
	return &this
}

// NewIntakeAPIKeyAttributesWithDefaults instantiates a new IntakeAPIKeyAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntakeAPIKeyAttributesWithDefaults() *IntakeAPIKeyAttributes {
	this := IntakeAPIKeyAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *IntakeAPIKeyAttributes) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *IntakeAPIKeyAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *IntakeAPIKeyAttributes) SetApiKey(v string) {
	o.ApiKey = v
}

// GetOrgId returns the OrgId field value.
func (o *IntakeAPIKeyAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *IntakeAPIKeyAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *IntakeAPIKeyAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntakeAPIKeyAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["org_id"] = o.OrgId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntakeAPIKeyAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey *string `json:"api_key"`
		OrgId  *int64  `json:"org_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "org_id"})
	} else {
		return err
	}
	o.ApiKey = *all.ApiKey
	o.OrgId = *all.OrgId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
