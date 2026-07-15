// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AccountFiltersAttributes Attributes for the account filters of a cloud account.
type AccountFiltersAttributes struct {
	// The account filtering configuration.
	AccountFilters *AccountFilteringConfig `json:"account_filters,omitempty"`
	// The cloud account ID.
	AccountId *string `json:"account_id,omitempty"`
	// The cloud provider of the account, for example `aws`, `aws_cur2`, or `oci`.
	Cloud *string `json:"cloud,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAccountFiltersAttributes instantiates a new AccountFiltersAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAccountFiltersAttributes() *AccountFiltersAttributes {
	this := AccountFiltersAttributes{}
	return &this
}

// NewAccountFiltersAttributesWithDefaults instantiates a new AccountFiltersAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAccountFiltersAttributesWithDefaults() *AccountFiltersAttributes {
	this := AccountFiltersAttributes{}
	return &this
}

// GetAccountFilters returns the AccountFilters field value if set, zero value otherwise.
func (o *AccountFiltersAttributes) GetAccountFilters() AccountFilteringConfig {
	if o == nil || o.AccountFilters == nil {
		var ret AccountFilteringConfig
		return ret
	}
	return *o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFiltersAttributes) GetAccountFiltersOk() (*AccountFilteringConfig, bool) {
	if o == nil || o.AccountFilters == nil {
		return nil, false
	}
	return o.AccountFilters, true
}

// HasAccountFilters returns a boolean if a field has been set.
func (o *AccountFiltersAttributes) HasAccountFilters() bool {
	return o != nil && o.AccountFilters != nil
}

// SetAccountFilters gets a reference to the given AccountFilteringConfig and assigns it to the AccountFilters field.
func (o *AccountFiltersAttributes) SetAccountFilters(v AccountFilteringConfig) {
	o.AccountFilters = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AccountFiltersAttributes) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFiltersAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AccountFiltersAttributes) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AccountFiltersAttributes) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *AccountFiltersAttributes) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFiltersAttributes) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *AccountFiltersAttributes) HasCloud() bool {
	return o != nil && o.Cloud != nil
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *AccountFiltersAttributes) SetCloud(v string) {
	o.Cloud = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AccountFiltersAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountFilters != nil {
		toSerialize["account_filters"] = o.AccountFilters
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AccountFiltersAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountFilters *AccountFilteringConfig `json:"account_filters,omitempty"`
		AccountId      *string                 `json:"account_id,omitempty"`
		Cloud          *string                 `json:"cloud,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_filters", "account_id", "cloud"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AccountFilters != nil && all.AccountFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AccountFilters = all.AccountFilters
	o.AccountId = all.AccountId
	o.Cloud = all.Cloud

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
