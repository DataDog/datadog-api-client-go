// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumArchiveSessionAttributes Archived session metadata.
type RumArchiveSessionAttributes struct {
	// Session account id, or null when unavailable.
	AccountId datadog.NullableString `json:"account_id"`
	// Session anonymous id, or null when unavailable.
	AnonymousId datadog.NullableString `json:"anonymous_id"`
	// Application UUID.
	ApplicationId uuid.UUID `json:"application_id"`
	// Session country, or null when unavailable.
	Country datadog.NullableString `json:"country"`
	// Session recovery state. A missing recovery record can be represented as null. Optional request_id and recovered_at are omitted when unavailable.
	RecoveryStatus NullableRumArchiveSessionRecoveryStatus `json:"recovery_status"`
	// Service name.
	Service string `json:"service"`
	// Session identifier.
	SessionId string `json:"session_id"`
	// Session timestamp in Unix milliseconds.
	Timestamp int64 `json:"timestamp"`
	// Session user id, or null when unavailable.
	UserId datadog.NullableString `json:"user_id"`
	// Session version, or null when unavailable.
	Version datadog.NullableString `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumArchiveSessionAttributes instantiates a new RumArchiveSessionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumArchiveSessionAttributes(accountId datadog.NullableString, anonymousId datadog.NullableString, applicationId uuid.UUID, country datadog.NullableString, recoveryStatus NullableRumArchiveSessionRecoveryStatus, service string, sessionId string, timestamp int64, userId datadog.NullableString, version datadog.NullableString) *RumArchiveSessionAttributes {
	this := RumArchiveSessionAttributes{}
	this.AccountId = accountId
	this.AnonymousId = anonymousId
	this.ApplicationId = applicationId
	this.Country = country
	this.RecoveryStatus = recoveryStatus
	this.Service = service
	this.SessionId = sessionId
	this.Timestamp = timestamp
	this.UserId = userId
	this.Version = version
	return &this
}

// NewRumArchiveSessionAttributesWithDefaults instantiates a new RumArchiveSessionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumArchiveSessionAttributesWithDefaults() *RumArchiveSessionAttributes {
	this := RumArchiveSessionAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *RumArchiveSessionAttributes) GetAccountId() string {
	if o == nil || o.AccountId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountId.Get()
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RumArchiveSessionAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountId.Get(), o.AccountId.IsSet()
}

// SetAccountId sets field value.
func (o *RumArchiveSessionAttributes) SetAccountId(v string) {
	o.AccountId.Set(&v)
}

// GetAnonymousId returns the AnonymousId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *RumArchiveSessionAttributes) GetAnonymousId() string {
	if o == nil || o.AnonymousId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AnonymousId.Get()
}

// GetAnonymousIdOk returns a tuple with the AnonymousId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RumArchiveSessionAttributes) GetAnonymousIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AnonymousId.Get(), o.AnonymousId.IsSet()
}

// SetAnonymousId sets field value.
func (o *RumArchiveSessionAttributes) SetAnonymousId(v string) {
	o.AnonymousId.Set(&v)
}

// GetApplicationId returns the ApplicationId field value.
func (o *RumArchiveSessionAttributes) GetApplicationId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionAttributes) GetApplicationIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value.
func (o *RumArchiveSessionAttributes) SetApplicationId(v uuid.UUID) {
	o.ApplicationId = v
}

// GetCountry returns the Country field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *RumArchiveSessionAttributes) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RumArchiveSessionAttributes) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// SetCountry sets field value.
func (o *RumArchiveSessionAttributes) SetCountry(v string) {
	o.Country.Set(&v)
}

// GetRecoveryStatus returns the RecoveryStatus field value.
// If the value is explicit nil, the zero value for RumArchiveSessionRecoveryStatus will be returned.
func (o *RumArchiveSessionAttributes) GetRecoveryStatus() RumArchiveSessionRecoveryStatus {
	if o == nil || o.RecoveryStatus.Get() == nil {
		var ret RumArchiveSessionRecoveryStatus
		return ret
	}
	return *o.RecoveryStatus.Get()
}

// GetRecoveryStatusOk returns a tuple with the RecoveryStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RumArchiveSessionAttributes) GetRecoveryStatusOk() (*RumArchiveSessionRecoveryStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecoveryStatus.Get(), o.RecoveryStatus.IsSet()
}

// SetRecoveryStatus sets field value.
func (o *RumArchiveSessionAttributes) SetRecoveryStatus(v RumArchiveSessionRecoveryStatus) {
	o.RecoveryStatus.Set(&v)
}

// GetService returns the Service field value.
func (o *RumArchiveSessionAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *RumArchiveSessionAttributes) SetService(v string) {
	o.Service = v
}

// GetSessionId returns the SessionId field value.
func (o *RumArchiveSessionAttributes) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionAttributes) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value.
func (o *RumArchiveSessionAttributes) SetSessionId(v string) {
	o.SessionId = v
}

// GetTimestamp returns the Timestamp field value.
func (o *RumArchiveSessionAttributes) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *RumArchiveSessionAttributes) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value.
func (o *RumArchiveSessionAttributes) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetUserId returns the UserId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *RumArchiveSessionAttributes) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RumArchiveSessionAttributes) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// SetUserId sets field value.
func (o *RumArchiveSessionAttributes) SetUserId(v string) {
	o.UserId.Set(&v)
}

// GetVersion returns the Version field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *RumArchiveSessionAttributes) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *RumArchiveSessionAttributes) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// SetVersion sets field value.
func (o *RumArchiveSessionAttributes) SetVersion(v string) {
	o.Version.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o RumArchiveSessionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId.Get()
	toSerialize["anonymous_id"] = o.AnonymousId.Get()
	toSerialize["application_id"] = o.ApplicationId
	toSerialize["country"] = o.Country.Get()
	toSerialize["recovery_status"] = o.RecoveryStatus.Get()
	toSerialize["service"] = o.Service
	toSerialize["session_id"] = o.SessionId
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["user_id"] = o.UserId.Get()
	toSerialize["version"] = o.Version.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumArchiveSessionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId      datadog.NullableString                  `json:"account_id"`
		AnonymousId    datadog.NullableString                  `json:"anonymous_id"`
		ApplicationId  *uuid.UUID                              `json:"application_id"`
		Country        datadog.NullableString                  `json:"country"`
		RecoveryStatus NullableRumArchiveSessionRecoveryStatus `json:"recovery_status"`
		Service        *string                                 `json:"service"`
		SessionId      *string                                 `json:"session_id"`
		Timestamp      *int64                                  `json:"timestamp"`
		UserId         datadog.NullableString                  `json:"user_id"`
		Version        datadog.NullableString                  `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.AccountId.IsSet() {
		return fmt.Errorf("required field account_id missing")
	}
	if !all.AnonymousId.IsSet() {
		return fmt.Errorf("required field anonymous_id missing")
	}
	if all.ApplicationId == nil {
		return fmt.Errorf("required field application_id missing")
	}
	if !all.Country.IsSet() {
		return fmt.Errorf("required field country missing")
	}
	if !all.RecoveryStatus.IsSet() {
		return fmt.Errorf("required field recovery_status missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.SessionId == nil {
		return fmt.Errorf("required field session_id missing")
	}
	if all.Timestamp == nil {
		return fmt.Errorf("required field timestamp missing")
	}
	if !all.UserId.IsSet() {
		return fmt.Errorf("required field user_id missing")
	}
	if !all.Version.IsSet() {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "anonymous_id", "application_id", "country", "recovery_status", "service", "session_id", "timestamp", "user_id", "version"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.AnonymousId = all.AnonymousId
	o.ApplicationId = *all.ApplicationId
	o.Country = all.Country
	o.RecoveryStatus = all.RecoveryStatus
	o.Service = *all.Service
	o.SessionId = *all.SessionId
	o.Timestamp = *all.Timestamp
	o.UserId = all.UserId
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
