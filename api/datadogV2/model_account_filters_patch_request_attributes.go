// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AccountFiltersPatchRequestAttributes Attributes for an account filters patch request.
type AccountFiltersPatchRequestAttributes struct {
	// The account filtering configuration.
	AccountFilters AccountFilteringConfig `json:"account_filters"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAccountFiltersPatchRequestAttributes instantiates a new AccountFiltersPatchRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAccountFiltersPatchRequestAttributes(accountFilters AccountFilteringConfig) *AccountFiltersPatchRequestAttributes {
	this := AccountFiltersPatchRequestAttributes{}
	this.AccountFilters = accountFilters
	return &this
}

// NewAccountFiltersPatchRequestAttributesWithDefaults instantiates a new AccountFiltersPatchRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAccountFiltersPatchRequestAttributesWithDefaults() *AccountFiltersPatchRequestAttributes {
	this := AccountFiltersPatchRequestAttributes{}
	return &this
}

// GetAccountFilters returns the AccountFilters field value.
func (o *AccountFiltersPatchRequestAttributes) GetAccountFilters() AccountFilteringConfig {
	if o == nil {
		var ret AccountFilteringConfig
		return ret
	}
	return o.AccountFilters
}

// GetAccountFiltersOk returns a tuple with the AccountFilters field value
// and a boolean to check if the value has been set.
func (o *AccountFiltersPatchRequestAttributes) GetAccountFiltersOk() (*AccountFilteringConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountFilters, true
}

// SetAccountFilters sets field value.
func (o *AccountFiltersPatchRequestAttributes) SetAccountFilters(v AccountFilteringConfig) {
	o.AccountFilters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AccountFiltersPatchRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_filters"] = o.AccountFilters

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AccountFiltersPatchRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountFilters *AccountFilteringConfig `json:"account_filters"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountFilters == nil {
		return fmt.Errorf("required field account_filters missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_filters"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AccountFilters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AccountFilters = *all.AccountFilters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
