// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsServerSideEventItem A Product Analytics server-side event.
type ProductAnalyticsServerSideEventItem struct {
	// The account linked to your event.
	Account *ProductAnalyticsAccount `json:"account,omitempty"`
	// The application in which you want to send your events.
	Application ProductAnalyticsApplication `json:"application"`
	// Fields used for the event.
	Event ProductAnalyticsEvent `json:"event"`
	// The session linked to your event.
	Session *ProductAnalyticsSession `json:"session,omitempty"`
	// The type of Product Analytics event. Must be `server` for server-side events.
	Type ProductAnalyticsEventType `json:"type"`
	// The user linked to your event.
	Usr *ProductAnalyticsUser `json:"usr,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsServerSideEventItem instantiates a new ProductAnalyticsServerSideEventItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsServerSideEventItem(application ProductAnalyticsApplication, event ProductAnalyticsEvent, typeVar ProductAnalyticsEventType) *ProductAnalyticsServerSideEventItem {
	this := ProductAnalyticsServerSideEventItem{}
	this.Application = application
	this.Event = event
	this.Type = typeVar
	return &this
}

// NewProductAnalyticsServerSideEventItemWithDefaults instantiates a new ProductAnalyticsServerSideEventItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsServerSideEventItemWithDefaults() *ProductAnalyticsServerSideEventItem {
	this := ProductAnalyticsServerSideEventItem{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ProductAnalyticsServerSideEventItem) GetAccount() ProductAnalyticsAccount {
	if o == nil || o.Account == nil {
		var ret ProductAnalyticsAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsServerSideEventItem) GetAccountOk() (*ProductAnalyticsAccount, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ProductAnalyticsServerSideEventItem) HasAccount() bool {
	return o != nil && o.Account != nil
}

// SetAccount gets a reference to the given ProductAnalyticsAccount and assigns it to the Account field.
func (o *ProductAnalyticsServerSideEventItem) SetAccount(v ProductAnalyticsAccount) {
	o.Account = &v
}

// GetApplication returns the Application field value.
func (o *ProductAnalyticsServerSideEventItem) GetApplication() ProductAnalyticsApplication {
	if o == nil {
		var ret ProductAnalyticsApplication
		return ret
	}
	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsServerSideEventItem) GetApplicationOk() (*ProductAnalyticsApplication, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value.
func (o *ProductAnalyticsServerSideEventItem) SetApplication(v ProductAnalyticsApplication) {
	o.Application = v
}

// GetEvent returns the Event field value.
func (o *ProductAnalyticsServerSideEventItem) GetEvent() ProductAnalyticsEvent {
	if o == nil {
		var ret ProductAnalyticsEvent
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsServerSideEventItem) GetEventOk() (*ProductAnalyticsEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value.
func (o *ProductAnalyticsServerSideEventItem) SetEvent(v ProductAnalyticsEvent) {
	o.Event = v
}

// GetSession returns the Session field value if set, zero value otherwise.
func (o *ProductAnalyticsServerSideEventItem) GetSession() ProductAnalyticsSession {
	if o == nil || o.Session == nil {
		var ret ProductAnalyticsSession
		return ret
	}
	return *o.Session
}

// GetSessionOk returns a tuple with the Session field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsServerSideEventItem) GetSessionOk() (*ProductAnalyticsSession, bool) {
	if o == nil || o.Session == nil {
		return nil, false
	}
	return o.Session, true
}

// HasSession returns a boolean if a field has been set.
func (o *ProductAnalyticsServerSideEventItem) HasSession() bool {
	return o != nil && o.Session != nil
}

// SetSession gets a reference to the given ProductAnalyticsSession and assigns it to the Session field.
func (o *ProductAnalyticsServerSideEventItem) SetSession(v ProductAnalyticsSession) {
	o.Session = &v
}

// GetType returns the Type field value.
func (o *ProductAnalyticsServerSideEventItem) GetType() ProductAnalyticsEventType {
	if o == nil {
		var ret ProductAnalyticsEventType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsServerSideEventItem) GetTypeOk() (*ProductAnalyticsEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ProductAnalyticsServerSideEventItem) SetType(v ProductAnalyticsEventType) {
	o.Type = v
}

// GetUsr returns the Usr field value if set, zero value otherwise.
func (o *ProductAnalyticsServerSideEventItem) GetUsr() ProductAnalyticsUser {
	if o == nil || o.Usr == nil {
		var ret ProductAnalyticsUser
		return ret
	}
	return *o.Usr
}

// GetUsrOk returns a tuple with the Usr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsServerSideEventItem) GetUsrOk() (*ProductAnalyticsUser, bool) {
	if o == nil || o.Usr == nil {
		return nil, false
	}
	return o.Usr, true
}

// HasUsr returns a boolean if a field has been set.
func (o *ProductAnalyticsServerSideEventItem) HasUsr() bool {
	return o != nil && o.Usr != nil
}

// SetUsr gets a reference to the given ProductAnalyticsUser and assigns it to the Usr field.
func (o *ProductAnalyticsServerSideEventItem) SetUsr(v ProductAnalyticsUser) {
	o.Usr = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsServerSideEventItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Account != nil {
		toSerialize["account"] = o.Account
	}
	toSerialize["application"] = o.Application
	toSerialize["event"] = o.Event
	if o.Session != nil {
		toSerialize["session"] = o.Session
	}
	toSerialize["type"] = o.Type
	if o.Usr != nil {
		toSerialize["usr"] = o.Usr
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsServerSideEventItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Account     *ProductAnalyticsAccount     `json:"account,omitempty"`
		Application *ProductAnalyticsApplication `json:"application"`
		Event       *ProductAnalyticsEvent       `json:"event"`
		Session     *ProductAnalyticsSession     `json:"session,omitempty"`
		Type        *ProductAnalyticsEventType   `json:"type"`
		Usr         *ProductAnalyticsUser        `json:"usr,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Application == nil {
		return fmt.Errorf("required field application missing")
	}
	if all.Event == nil {
		return fmt.Errorf("required field event missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account", "application", "event", "session", "type", "usr"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Account != nil && all.Account.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Account = all.Account
	if all.Application.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Application = *all.Application
	if all.Event.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Event = *all.Event
	if all.Session != nil && all.Session.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Session = all.Session
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.Usr != nil && all.Usr.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Usr = all.Usr

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
