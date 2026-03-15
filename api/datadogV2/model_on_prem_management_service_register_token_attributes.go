// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnPremManagementServiceRegisterTokenAttributes Attributes for registering a token.
type OnPremManagementServiceRegisterTokenAttributes struct {
	// The application identifier.
	AppId *uuid.UUID `json:"app_id,omitempty"`
	// The application version.
	AppVersion *int64 `json:"app_version,omitempty"`
	// The connection identifier.
	ConnectionId uuid.UUID `json:"connection_id"`
	// The query identifier.
	QueryId *uuid.UUID `json:"query_id,omitempty"`
	// The on-prem runner identifier.
	RunnerId string `json:"runner_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnPremManagementServiceRegisterTokenAttributes instantiates a new OnPremManagementServiceRegisterTokenAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnPremManagementServiceRegisterTokenAttributes(connectionId uuid.UUID, runnerId string) *OnPremManagementServiceRegisterTokenAttributes {
	this := OnPremManagementServiceRegisterTokenAttributes{}
	this.ConnectionId = connectionId
	this.RunnerId = runnerId
	return &this
}

// NewOnPremManagementServiceRegisterTokenAttributesWithDefaults instantiates a new OnPremManagementServiceRegisterTokenAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnPremManagementServiceRegisterTokenAttributesWithDefaults() *OnPremManagementServiceRegisterTokenAttributes {
	this := OnPremManagementServiceRegisterTokenAttributes{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetAppId() uuid.UUID {
	if o == nil || o.AppId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetAppIdOk() (*uuid.UUID, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) HasAppId() bool {
	return o != nil && o.AppId != nil
}

// SetAppId gets a reference to the given uuid.UUID and assigns it to the AppId field.
func (o *OnPremManagementServiceRegisterTokenAttributes) SetAppId(v uuid.UUID) {
	o.AppId = &v
}

// GetAppVersion returns the AppVersion field value if set, zero value otherwise.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetAppVersion() int64 {
	if o == nil || o.AppVersion == nil {
		var ret int64
		return ret
	}
	return *o.AppVersion
}

// GetAppVersionOk returns a tuple with the AppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetAppVersionOk() (*int64, bool) {
	if o == nil || o.AppVersion == nil {
		return nil, false
	}
	return o.AppVersion, true
}

// HasAppVersion returns a boolean if a field has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) HasAppVersion() bool {
	return o != nil && o.AppVersion != nil
}

// SetAppVersion gets a reference to the given int64 and assigns it to the AppVersion field.
func (o *OnPremManagementServiceRegisterTokenAttributes) SetAppVersion(v int64) {
	o.AppVersion = &v
}

// GetConnectionId returns the ConnectionId field value.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetConnectionId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetConnectionIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value.
func (o *OnPremManagementServiceRegisterTokenAttributes) SetConnectionId(v uuid.UUID) {
	o.ConnectionId = v
}

// GetQueryId returns the QueryId field value if set, zero value otherwise.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetQueryId() uuid.UUID {
	if o == nil || o.QueryId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.QueryId
}

// GetQueryIdOk returns a tuple with the QueryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetQueryIdOk() (*uuid.UUID, bool) {
	if o == nil || o.QueryId == nil {
		return nil, false
	}
	return o.QueryId, true
}

// HasQueryId returns a boolean if a field has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) HasQueryId() bool {
	return o != nil && o.QueryId != nil
}

// SetQueryId gets a reference to the given uuid.UUID and assigns it to the QueryId field.
func (o *OnPremManagementServiceRegisterTokenAttributes) SetQueryId(v uuid.UUID) {
	o.QueryId = &v
}

// GetRunnerId returns the RunnerId field value.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetRunnerId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunnerId
}

// GetRunnerIdOk returns a tuple with the RunnerId field value
// and a boolean to check if the value has been set.
func (o *OnPremManagementServiceRegisterTokenAttributes) GetRunnerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunnerId, true
}

// SetRunnerId sets field value.
func (o *OnPremManagementServiceRegisterTokenAttributes) SetRunnerId(v string) {
	o.RunnerId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnPremManagementServiceRegisterTokenAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.AppVersion != nil {
		toSerialize["app_version"] = o.AppVersion
	}
	toSerialize["connection_id"] = o.ConnectionId
	if o.QueryId != nil {
		toSerialize["query_id"] = o.QueryId
	}
	toSerialize["runner_id"] = o.RunnerId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnPremManagementServiceRegisterTokenAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppId        *uuid.UUID `json:"app_id,omitempty"`
		AppVersion   *int64     `json:"app_version,omitempty"`
		ConnectionId *uuid.UUID `json:"connection_id"`
		QueryId      *uuid.UUID `json:"query_id,omitempty"`
		RunnerId     *string    `json:"runner_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConnectionId == nil {
		return fmt.Errorf("required field connection_id missing")
	}
	if all.RunnerId == nil {
		return fmt.Errorf("required field runner_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"app_id", "app_version", "connection_id", "query_id", "runner_id"})
	} else {
		return err
	}
	o.AppId = all.AppId
	o.AppVersion = all.AppVersion
	o.ConnectionId = *all.ConnectionId
	o.QueryId = all.QueryId
	o.RunnerId = *all.RunnerId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
