// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogMessage The message delivery event details.
type TransportWebhookLogMessage struct {
	// The message authentication details.
	Auth *TransportWebhookLogMessageAuth `json:"auth,omitempty"`
	// Custom arguments passed through the email transport provider for tracking.
	CustomArgs *TransportWebhookLogMessageCustomArgs `json:"custom_args,omitempty"`
	// The message identifiers.
	Id *TransportWebhookLogMessageId `json:"id,omitempty"`
	// The delivery event type emitted by the transport provider (for example, "delivered", "dropped", "bounced").
	Name *string `json:"name,omitempty"`
	// The SMTP response information.
	Response *TransportWebhookLogMessageResponse `json:"response,omitempty"`
	// The IP address of the sending server.
	SenderIp *string `json:"sender_ip,omitempty"`
	// The message delivery timing information.
	Timestamp *TransportWebhookLogMessageTimestamp `json:"timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogMessage instantiates a new TransportWebhookLogMessage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogMessage() *TransportWebhookLogMessage {
	this := TransportWebhookLogMessage{}
	return &this
}

// NewTransportWebhookLogMessageWithDefaults instantiates a new TransportWebhookLogMessage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogMessageWithDefaults() *TransportWebhookLogMessage {
	this := TransportWebhookLogMessage{}
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *TransportWebhookLogMessage) GetAuth() TransportWebhookLogMessageAuth {
	if o == nil || o.Auth == nil {
		var ret TransportWebhookLogMessageAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessage) GetAuthOk() (*TransportWebhookLogMessageAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *TransportWebhookLogMessage) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given TransportWebhookLogMessageAuth and assigns it to the Auth field.
func (o *TransportWebhookLogMessage) SetAuth(v TransportWebhookLogMessageAuth) {
	o.Auth = &v
}

// GetCustomArgs returns the CustomArgs field value if set, zero value otherwise.
func (o *TransportWebhookLogMessage) GetCustomArgs() TransportWebhookLogMessageCustomArgs {
	if o == nil || o.CustomArgs == nil {
		var ret TransportWebhookLogMessageCustomArgs
		return ret
	}
	return *o.CustomArgs
}

// GetCustomArgsOk returns a tuple with the CustomArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessage) GetCustomArgsOk() (*TransportWebhookLogMessageCustomArgs, bool) {
	if o == nil || o.CustomArgs == nil {
		return nil, false
	}
	return o.CustomArgs, true
}

// HasCustomArgs returns a boolean if a field has been set.
func (o *TransportWebhookLogMessage) HasCustomArgs() bool {
	return o != nil && o.CustomArgs != nil
}

// SetCustomArgs gets a reference to the given TransportWebhookLogMessageCustomArgs and assigns it to the CustomArgs field.
func (o *TransportWebhookLogMessage) SetCustomArgs(v TransportWebhookLogMessageCustomArgs) {
	o.CustomArgs = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TransportWebhookLogMessage) GetId() TransportWebhookLogMessageId {
	if o == nil || o.Id == nil {
		var ret TransportWebhookLogMessageId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessage) GetIdOk() (*TransportWebhookLogMessageId, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TransportWebhookLogMessage) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given TransportWebhookLogMessageId and assigns it to the Id field.
func (o *TransportWebhookLogMessage) SetId(v TransportWebhookLogMessageId) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransportWebhookLogMessage) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessage) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransportWebhookLogMessage) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransportWebhookLogMessage) SetName(v string) {
	o.Name = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *TransportWebhookLogMessage) GetResponse() TransportWebhookLogMessageResponse {
	if o == nil || o.Response == nil {
		var ret TransportWebhookLogMessageResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessage) GetResponseOk() (*TransportWebhookLogMessageResponse, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *TransportWebhookLogMessage) HasResponse() bool {
	return o != nil && o.Response != nil
}

// SetResponse gets a reference to the given TransportWebhookLogMessageResponse and assigns it to the Response field.
func (o *TransportWebhookLogMessage) SetResponse(v TransportWebhookLogMessageResponse) {
	o.Response = &v
}

// GetSenderIp returns the SenderIp field value if set, zero value otherwise.
func (o *TransportWebhookLogMessage) GetSenderIp() string {
	if o == nil || o.SenderIp == nil {
		var ret string
		return ret
	}
	return *o.SenderIp
}

// GetSenderIpOk returns a tuple with the SenderIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessage) GetSenderIpOk() (*string, bool) {
	if o == nil || o.SenderIp == nil {
		return nil, false
	}
	return o.SenderIp, true
}

// HasSenderIp returns a boolean if a field has been set.
func (o *TransportWebhookLogMessage) HasSenderIp() bool {
	return o != nil && o.SenderIp != nil
}

// SetSenderIp gets a reference to the given string and assigns it to the SenderIp field.
func (o *TransportWebhookLogMessage) SetSenderIp(v string) {
	o.SenderIp = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *TransportWebhookLogMessage) GetTimestamp() TransportWebhookLogMessageTimestamp {
	if o == nil || o.Timestamp == nil {
		var ret TransportWebhookLogMessageTimestamp
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogMessage) GetTimestampOk() (*TransportWebhookLogMessageTimestamp, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *TransportWebhookLogMessage) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given TransportWebhookLogMessageTimestamp and assigns it to the Timestamp field.
func (o *TransportWebhookLogMessage) SetTimestamp(v TransportWebhookLogMessageTimestamp) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.CustomArgs != nil {
		toSerialize["custom_args"] = o.CustomArgs
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.SenderIp != nil {
		toSerialize["sender_ip"] = o.SenderIp
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogMessage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth       *TransportWebhookLogMessageAuth       `json:"auth,omitempty"`
		CustomArgs *TransportWebhookLogMessageCustomArgs `json:"custom_args,omitempty"`
		Id         *TransportWebhookLogMessageId         `json:"id,omitempty"`
		Name       *string                               `json:"name,omitempty"`
		Response   *TransportWebhookLogMessageResponse   `json:"response,omitempty"`
		SenderIp   *string                               `json:"sender_ip,omitempty"`
		Timestamp  *TransportWebhookLogMessageTimestamp  `json:"timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "custom_args", "id", "name", "response", "sender_ip", "timestamp"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	if all.CustomArgs != nil && all.CustomArgs.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CustomArgs = all.CustomArgs
	if all.Id != nil && all.Id.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Id = all.Id
	o.Name = all.Name
	if all.Response != nil && all.Response.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Response = all.Response
	o.SenderIp = all.SenderIp
	if all.Timestamp != nil && all.Timestamp.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
