// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TransportWebhookLogAttributes Top-level attributes for the webhook log event, including delivery status, recipient details, and provider metadata.
type TransportWebhookLogAttributes struct {
	// The event categories.
	Category []string `json:"category,omitempty"`
	// The email address details.
	Email *TransportWebhookLogEmail `json:"email,omitempty"`
	// The unique email identifier.
	EmailId *string `json:"email_id,omitempty"`
	// The human-readable email type name.
	EmailTypeDisplayName *string `json:"email_type_display_name,omitempty"`
	// The message delivery event details.
	Message *TransportWebhookLogMessage `json:"message,omitempty"`
	// The network information for the event.
	Network *TransportWebhookLogNetwork `json:"network,omitempty"`
	// The numeric organization identifier.
	Org *int64 `json:"org,omitempty"`
	// Metadata about the organization that sent the email.
	OrgMetadata *TransportWebhookLogOrgMetadata `json:"org_metadata,omitempty"`
	// The organization UUID.
	OrgUuid *string `json:"org_uuid,omitempty"`
	// The timestamp when the email was queued.
	QueueTime *string `json:"queue_time,omitempty"`
	// Indicates whether the open event was triggered by automated machine activity rather than a human recipient (SendGrid-specific).
	SgMachineOpen *bool `json:"sg_machine_open,omitempty"`
	// The email subject line.
	Subject *string `json:"subject,omitempty"`
	// The user agent string for open events.
	Useragent *string `json:"useragent,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTransportWebhookLogAttributes instantiates a new TransportWebhookLogAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTransportWebhookLogAttributes() *TransportWebhookLogAttributes {
	this := TransportWebhookLogAttributes{}
	return &this
}

// NewTransportWebhookLogAttributesWithDefaults instantiates a new TransportWebhookLogAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTransportWebhookLogAttributesWithDefaults() *TransportWebhookLogAttributes {
	this := TransportWebhookLogAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetCategory() []string {
	if o == nil || o.Category == nil {
		var ret []string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetCategoryOk() (*[]string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return &o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given []string and assigns it to the Category field.
func (o *TransportWebhookLogAttributes) SetCategory(v []string) {
	o.Category = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetEmail() TransportWebhookLogEmail {
	if o == nil || o.Email == nil {
		var ret TransportWebhookLogEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetEmailOk() (*TransportWebhookLogEmail, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given TransportWebhookLogEmail and assigns it to the Email field.
func (o *TransportWebhookLogAttributes) SetEmail(v TransportWebhookLogEmail) {
	o.Email = &v
}

// GetEmailId returns the EmailId field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetEmailId() string {
	if o == nil || o.EmailId == nil {
		var ret string
		return ret
	}
	return *o.EmailId
}

// GetEmailIdOk returns a tuple with the EmailId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetEmailIdOk() (*string, bool) {
	if o == nil || o.EmailId == nil {
		return nil, false
	}
	return o.EmailId, true
}

// HasEmailId returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasEmailId() bool {
	return o != nil && o.EmailId != nil
}

// SetEmailId gets a reference to the given string and assigns it to the EmailId field.
func (o *TransportWebhookLogAttributes) SetEmailId(v string) {
	o.EmailId = &v
}

// GetEmailTypeDisplayName returns the EmailTypeDisplayName field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetEmailTypeDisplayName() string {
	if o == nil || o.EmailTypeDisplayName == nil {
		var ret string
		return ret
	}
	return *o.EmailTypeDisplayName
}

// GetEmailTypeDisplayNameOk returns a tuple with the EmailTypeDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetEmailTypeDisplayNameOk() (*string, bool) {
	if o == nil || o.EmailTypeDisplayName == nil {
		return nil, false
	}
	return o.EmailTypeDisplayName, true
}

// HasEmailTypeDisplayName returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasEmailTypeDisplayName() bool {
	return o != nil && o.EmailTypeDisplayName != nil
}

// SetEmailTypeDisplayName gets a reference to the given string and assigns it to the EmailTypeDisplayName field.
func (o *TransportWebhookLogAttributes) SetEmailTypeDisplayName(v string) {
	o.EmailTypeDisplayName = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetMessage() TransportWebhookLogMessage {
	if o == nil || o.Message == nil {
		var ret TransportWebhookLogMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetMessageOk() (*TransportWebhookLogMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given TransportWebhookLogMessage and assigns it to the Message field.
func (o *TransportWebhookLogAttributes) SetMessage(v TransportWebhookLogMessage) {
	o.Message = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetNetwork() TransportWebhookLogNetwork {
	if o == nil || o.Network == nil {
		var ret TransportWebhookLogNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetNetworkOk() (*TransportWebhookLogNetwork, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasNetwork() bool {
	return o != nil && o.Network != nil
}

// SetNetwork gets a reference to the given TransportWebhookLogNetwork and assigns it to the Network field.
func (o *TransportWebhookLogAttributes) SetNetwork(v TransportWebhookLogNetwork) {
	o.Network = &v
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetOrg() int64 {
	if o == nil || o.Org == nil {
		var ret int64
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetOrgOk() (*int64, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasOrg() bool {
	return o != nil && o.Org != nil
}

// SetOrg gets a reference to the given int64 and assigns it to the Org field.
func (o *TransportWebhookLogAttributes) SetOrg(v int64) {
	o.Org = &v
}

// GetOrgMetadata returns the OrgMetadata field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetOrgMetadata() TransportWebhookLogOrgMetadata {
	if o == nil || o.OrgMetadata == nil {
		var ret TransportWebhookLogOrgMetadata
		return ret
	}
	return *o.OrgMetadata
}

// GetOrgMetadataOk returns a tuple with the OrgMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetOrgMetadataOk() (*TransportWebhookLogOrgMetadata, bool) {
	if o == nil || o.OrgMetadata == nil {
		return nil, false
	}
	return o.OrgMetadata, true
}

// HasOrgMetadata returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasOrgMetadata() bool {
	return o != nil && o.OrgMetadata != nil
}

// SetOrgMetadata gets a reference to the given TransportWebhookLogOrgMetadata and assigns it to the OrgMetadata field.
func (o *TransportWebhookLogAttributes) SetOrgMetadata(v TransportWebhookLogOrgMetadata) {
	o.OrgMetadata = &v
}

// GetOrgUuid returns the OrgUuid field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetOrgUuid() string {
	if o == nil || o.OrgUuid == nil {
		var ret string
		return ret
	}
	return *o.OrgUuid
}

// GetOrgUuidOk returns a tuple with the OrgUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetOrgUuidOk() (*string, bool) {
	if o == nil || o.OrgUuid == nil {
		return nil, false
	}
	return o.OrgUuid, true
}

// HasOrgUuid returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasOrgUuid() bool {
	return o != nil && o.OrgUuid != nil
}

// SetOrgUuid gets a reference to the given string and assigns it to the OrgUuid field.
func (o *TransportWebhookLogAttributes) SetOrgUuid(v string) {
	o.OrgUuid = &v
}

// GetQueueTime returns the QueueTime field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetQueueTime() string {
	if o == nil || o.QueueTime == nil {
		var ret string
		return ret
	}
	return *o.QueueTime
}

// GetQueueTimeOk returns a tuple with the QueueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetQueueTimeOk() (*string, bool) {
	if o == nil || o.QueueTime == nil {
		return nil, false
	}
	return o.QueueTime, true
}

// HasQueueTime returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasQueueTime() bool {
	return o != nil && o.QueueTime != nil
}

// SetQueueTime gets a reference to the given string and assigns it to the QueueTime field.
func (o *TransportWebhookLogAttributes) SetQueueTime(v string) {
	o.QueueTime = &v
}

// GetSgMachineOpen returns the SgMachineOpen field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetSgMachineOpen() bool {
	if o == nil || o.SgMachineOpen == nil {
		var ret bool
		return ret
	}
	return *o.SgMachineOpen
}

// GetSgMachineOpenOk returns a tuple with the SgMachineOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetSgMachineOpenOk() (*bool, bool) {
	if o == nil || o.SgMachineOpen == nil {
		return nil, false
	}
	return o.SgMachineOpen, true
}

// HasSgMachineOpen returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasSgMachineOpen() bool {
	return o != nil && o.SgMachineOpen != nil
}

// SetSgMachineOpen gets a reference to the given bool and assigns it to the SgMachineOpen field.
func (o *TransportWebhookLogAttributes) SetSgMachineOpen(v bool) {
	o.SgMachineOpen = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasSubject() bool {
	return o != nil && o.Subject != nil
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *TransportWebhookLogAttributes) SetSubject(v string) {
	o.Subject = &v
}

// GetUseragent returns the Useragent field value if set, zero value otherwise.
func (o *TransportWebhookLogAttributes) GetUseragent() string {
	if o == nil || o.Useragent == nil {
		var ret string
		return ret
	}
	return *o.Useragent
}

// GetUseragentOk returns a tuple with the Useragent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportWebhookLogAttributes) GetUseragentOk() (*string, bool) {
	if o == nil || o.Useragent == nil {
		return nil, false
	}
	return o.Useragent, true
}

// HasUseragent returns a boolean if a field has been set.
func (o *TransportWebhookLogAttributes) HasUseragent() bool {
	return o != nil && o.Useragent != nil
}

// SetUseragent gets a reference to the given string and assigns it to the Useragent field.
func (o *TransportWebhookLogAttributes) SetUseragent(v string) {
	o.Useragent = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TransportWebhookLogAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.EmailId != nil {
		toSerialize["email_id"] = o.EmailId
	}
	if o.EmailTypeDisplayName != nil {
		toSerialize["email_type_display_name"] = o.EmailTypeDisplayName
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if o.OrgMetadata != nil {
		toSerialize["org_metadata"] = o.OrgMetadata
	}
	if o.OrgUuid != nil {
		toSerialize["org_uuid"] = o.OrgUuid
	}
	if o.QueueTime != nil {
		toSerialize["queue_time"] = o.QueueTime
	}
	if o.SgMachineOpen != nil {
		toSerialize["sg_machine_open"] = o.SgMachineOpen
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Useragent != nil {
		toSerialize["useragent"] = o.Useragent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TransportWebhookLogAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category             []string                        `json:"category,omitempty"`
		Email                *TransportWebhookLogEmail       `json:"email,omitempty"`
		EmailId              *string                         `json:"email_id,omitempty"`
		EmailTypeDisplayName *string                         `json:"email_type_display_name,omitempty"`
		Message              *TransportWebhookLogMessage     `json:"message,omitempty"`
		Network              *TransportWebhookLogNetwork     `json:"network,omitempty"`
		Org                  *int64                          `json:"org,omitempty"`
		OrgMetadata          *TransportWebhookLogOrgMetadata `json:"org_metadata,omitempty"`
		OrgUuid              *string                         `json:"org_uuid,omitempty"`
		QueueTime            *string                         `json:"queue_time,omitempty"`
		SgMachineOpen        *bool                           `json:"sg_machine_open,omitempty"`
		Subject              *string                         `json:"subject,omitempty"`
		Useragent            *string                         `json:"useragent,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "email", "email_id", "email_type_display_name", "message", "network", "org", "org_metadata", "org_uuid", "queue_time", "sg_machine_open", "subject", "useragent"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Category = all.Category
	if all.Email != nil && all.Email.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Email = all.Email
	o.EmailId = all.EmailId
	o.EmailTypeDisplayName = all.EmailTypeDisplayName
	if all.Message != nil && all.Message.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Message = all.Message
	if all.Network != nil && all.Network.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Network = all.Network
	o.Org = all.Org
	if all.OrgMetadata != nil && all.OrgMetadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OrgMetadata = all.OrgMetadata
	o.OrgUuid = all.OrgUuid
	o.QueueTime = all.QueueTime
	o.SgMachineOpen = all.SgMachineOpen
	o.Subject = all.Subject
	o.Useragent = all.Useragent

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
