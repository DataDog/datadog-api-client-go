// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleDocsPostmortemSettings Settings for a postmortem template stored in Google Docs. Required when `location` is `google_docs`.
type GoogleDocsPostmortemSettings struct {
	// The ID of the Google Drive integration account.
	AccountId string `json:"account_id"`
	// The ID of the Google Drive folder where postmortems are created.
	ParentFolderId string `json:"parent_folder_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleDocsPostmortemSettings instantiates a new GoogleDocsPostmortemSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleDocsPostmortemSettings(accountId string, parentFolderId string) *GoogleDocsPostmortemSettings {
	this := GoogleDocsPostmortemSettings{}
	this.AccountId = accountId
	this.ParentFolderId = parentFolderId
	return &this
}

// NewGoogleDocsPostmortemSettingsWithDefaults instantiates a new GoogleDocsPostmortemSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleDocsPostmortemSettingsWithDefaults() *GoogleDocsPostmortemSettings {
	this := GoogleDocsPostmortemSettings{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *GoogleDocsPostmortemSettings) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *GoogleDocsPostmortemSettings) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *GoogleDocsPostmortemSettings) SetAccountId(v string) {
	o.AccountId = v
}

// GetParentFolderId returns the ParentFolderId field value.
func (o *GoogleDocsPostmortemSettings) GetParentFolderId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentFolderId
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value
// and a boolean to check if the value has been set.
func (o *GoogleDocsPostmortemSettings) GetParentFolderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentFolderId, true
}

// SetParentFolderId sets field value.
func (o *GoogleDocsPostmortemSettings) SetParentFolderId(v string) {
	o.ParentFolderId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleDocsPostmortemSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["parent_folder_id"] = o.ParentFolderId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleDocsPostmortemSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId      *string `json:"account_id"`
		ParentFolderId *string `json:"parent_folder_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.ParentFolderId == nil {
		return fmt.Errorf("required field parent_folder_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "parent_folder_id"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.ParentFolderId = *all.ParentFolderId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
