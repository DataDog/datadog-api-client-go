// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachLinearIssueRequestDataAttributes Attributes of the Linear issue to attach security findings to.
type AttachLinearIssueRequestDataAttributes struct {
	// URL of the Linear issue to attach security findings to.
	LinearIssueUrl string `json:"linear_issue_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAttachLinearIssueRequestDataAttributes instantiates a new AttachLinearIssueRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttachLinearIssueRequestDataAttributes(linearIssueUrl string) *AttachLinearIssueRequestDataAttributes {
	this := AttachLinearIssueRequestDataAttributes{}
	this.LinearIssueUrl = linearIssueUrl
	return &this
}

// NewAttachLinearIssueRequestDataAttributesWithDefaults instantiates a new AttachLinearIssueRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttachLinearIssueRequestDataAttributesWithDefaults() *AttachLinearIssueRequestDataAttributes {
	this := AttachLinearIssueRequestDataAttributes{}
	return &this
}

// GetLinearIssueUrl returns the LinearIssueUrl field value.
func (o *AttachLinearIssueRequestDataAttributes) GetLinearIssueUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LinearIssueUrl
}

// GetLinearIssueUrlOk returns a tuple with the LinearIssueUrl field value
// and a boolean to check if the value has been set.
func (o *AttachLinearIssueRequestDataAttributes) GetLinearIssueUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinearIssueUrl, true
}

// SetLinearIssueUrl sets field value.
func (o *AttachLinearIssueRequestDataAttributes) SetLinearIssueUrl(v string) {
	o.LinearIssueUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttachLinearIssueRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["linear_issue_url"] = o.LinearIssueUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AttachLinearIssueRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LinearIssueUrl *string `json:"linear_issue_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LinearIssueUrl == nil {
		return fmt.Errorf("required field linear_issue_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"linear_issue_url"})
	} else {
		return err
	}
	o.LinearIssueUrl = *all.LinearIssueUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
