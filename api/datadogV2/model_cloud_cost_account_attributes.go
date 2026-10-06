// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudCostAccountAttributes Read-only status and identifiers for a cloud cost account.
type CloudCostAccountAttributes struct {
	// The cloud provider account identifier, such as an OCI tenancy OCID or AWS account ID.
	AccountId string `json:"account_id"`
	// The cloud provider and cost report type. Currently supports `oci` and `aws_cur2`.
	Cloud string `json:"cloud"`
	// The timestamp when the cloud account was created.
	CreatedAt string `json:"created_at"`
	// Validation errors for the cloud account. Empty when there are no errors.
	ErrorMessages []string `json:"error_messages"`
	// The cloud account status, one of `active`, `warn`, `error`, or `disabled`.
	Status string `json:"status"`
	// The timestamp when the cloud account status was last updated.
	StatusUpdatedAt string `json:"status_updated_at"`
	// The timestamp when the cloud account was last updated.
	UpdatedAt string `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudCostAccountAttributes instantiates a new CloudCostAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudCostAccountAttributes(accountId string, cloud string, createdAt string, errorMessages []string, status string, statusUpdatedAt string, updatedAt string) *CloudCostAccountAttributes {
	this := CloudCostAccountAttributes{}
	this.AccountId = accountId
	this.Cloud = cloud
	this.CreatedAt = createdAt
	this.ErrorMessages = errorMessages
	this.Status = status
	this.StatusUpdatedAt = statusUpdatedAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewCloudCostAccountAttributesWithDefaults instantiates a new CloudCostAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudCostAccountAttributesWithDefaults() *CloudCostAccountAttributes {
	this := CloudCostAccountAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *CloudCostAccountAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *CloudCostAccountAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *CloudCostAccountAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetCloud returns the Cloud field value.
func (o *CloudCostAccountAttributes) GetCloud() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value
// and a boolean to check if the value has been set.
func (o *CloudCostAccountAttributes) GetCloudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cloud, true
}

// SetCloud sets field value.
func (o *CloudCostAccountAttributes) SetCloud(v string) {
	o.Cloud = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CloudCostAccountAttributes) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CloudCostAccountAttributes) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CloudCostAccountAttributes) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetErrorMessages returns the ErrorMessages field value.
func (o *CloudCostAccountAttributes) GetErrorMessages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value
// and a boolean to check if the value has been set.
func (o *CloudCostAccountAttributes) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessages, true
}

// SetErrorMessages sets field value.
func (o *CloudCostAccountAttributes) SetErrorMessages(v []string) {
	o.ErrorMessages = v
}

// GetStatus returns the Status field value.
func (o *CloudCostAccountAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudCostAccountAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CloudCostAccountAttributes) SetStatus(v string) {
	o.Status = v
}

// GetStatusUpdatedAt returns the StatusUpdatedAt field value.
func (o *CloudCostAccountAttributes) GetStatusUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.StatusUpdatedAt
}

// GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CloudCostAccountAttributes) GetStatusUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusUpdatedAt, true
}

// SetStatusUpdatedAt sets field value.
func (o *CloudCostAccountAttributes) SetStatusUpdatedAt(v string) {
	o.StatusUpdatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *CloudCostAccountAttributes) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CloudCostAccountAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *CloudCostAccountAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudCostAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["cloud"] = o.Cloud
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["error_messages"] = o.ErrorMessages
	toSerialize["status"] = o.Status
	toSerialize["status_updated_at"] = o.StatusUpdatedAt
	toSerialize["updated_at"] = o.UpdatedAt

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudCostAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId       *string   `json:"account_id"`
		Cloud           *string   `json:"cloud"`
		CreatedAt       *string   `json:"created_at"`
		ErrorMessages   *[]string `json:"error_messages"`
		Status          *string   `json:"status"`
		StatusUpdatedAt *string   `json:"status_updated_at"`
		UpdatedAt       *string   `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.Cloud == nil {
		return fmt.Errorf("required field cloud missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ErrorMessages == nil {
		return fmt.Errorf("required field error_messages missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.StatusUpdatedAt == nil {
		return fmt.Errorf("required field status_updated_at missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "cloud", "created_at", "error_messages", "status", "status_updated_at", "updated_at"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.Cloud = *all.Cloud
	o.CreatedAt = *all.CreatedAt
	o.ErrorMessages = *all.ErrorMessages
	o.Status = *all.Status
	o.StatusUpdatedAt = *all.StatusUpdatedAt
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
