// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPUsageCostConfigPostRequestAttributes Attributes for GCP Usage Cost config post request.
type GCPUsageCostConfigPostRequestAttributes struct {
	// The GCP account ID.
	AccountId string `json:"account_id"`
	// The GCP bucket name used to store the Usage Cost export.
	BucketName string `json:"bucket_name"`
	// The export dataset name used for the GCP Usage Cost report.
	Dataset string `json:"dataset"`
	// The export prefix used for the GCP Usage Cost report.
	ExportPrefix string `json:"export_prefix"`
	// The name of the GCP Usage Cost report.
	ExportProjectName string `json:"export_project_name"`
	// Whether or not the Cloud Cost Management account is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`
	// The month of the report.
	Months *int32 `json:"months,omitempty"`
	// The unique GCP service account email.
	ServiceAccount string `json:"service_account"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPUsageCostConfigPostRequestAttributes instantiates a new GCPUsageCostConfigPostRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPUsageCostConfigPostRequestAttributes(accountId string, bucketName string, dataset string, exportPrefix string, exportProjectName string, serviceAccount string) *GCPUsageCostConfigPostRequestAttributes {
	this := GCPUsageCostConfigPostRequestAttributes{}
	this.AccountId = accountId
	this.BucketName = bucketName
	this.Dataset = dataset
	this.ExportPrefix = exportPrefix
	this.ExportProjectName = exportProjectName
	this.ServiceAccount = serviceAccount
	return &this
}

// NewGCPUsageCostConfigPostRequestAttributesWithDefaults instantiates a new GCPUsageCostConfigPostRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPUsageCostConfigPostRequestAttributesWithDefaults() *GCPUsageCostConfigPostRequestAttributes {
	this := GCPUsageCostConfigPostRequestAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetBucketName returns the BucketName field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetBucketName(v string) {
	o.BucketName = v
}

// GetDataset returns the Dataset field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetDataset() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetDatasetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dataset, true
}

// SetDataset sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetDataset(v string) {
	o.Dataset = v
}

// GetExportPrefix returns the ExportPrefix field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportPrefix
}

// GetExportPrefixOk returns a tuple with the ExportPrefix field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportPrefix, true
}

// SetExportPrefix sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetExportPrefix(v string) {
	o.ExportPrefix = v
}

// GetExportProjectName returns the ExportProjectName field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportProjectName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ExportProjectName
}

// GetExportProjectNameOk returns a tuple with the ExportProjectName field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetExportProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportProjectName, true
}

// SetExportProjectName sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetExportProjectName(v string) {
	o.ExportProjectName = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *GCPUsageCostConfigPostRequestAttributes) GetIsEnabled() bool {
	if o == nil || o.IsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil || o.IsEnabled == nil {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) HasIsEnabled() bool {
	return o != nil && o.IsEnabled != nil
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *GCPUsageCostConfigPostRequestAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetMonths returns the Months field value if set, zero value otherwise.
func (o *GCPUsageCostConfigPostRequestAttributes) GetMonths() int32 {
	if o == nil || o.Months == nil {
		var ret int32
		return ret
	}
	return *o.Months
}

// GetMonthsOk returns a tuple with the Months field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetMonthsOk() (*int32, bool) {
	if o == nil || o.Months == nil {
		return nil, false
	}
	return o.Months, true
}

// HasMonths returns a boolean if a field has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) HasMonths() bool {
	return o != nil && o.Months != nil
}

// SetMonths gets a reference to the given int32 and assigns it to the Months field.
func (o *GCPUsageCostConfigPostRequestAttributes) SetMonths(v int32) {
	o.Months = &v
}

// GetServiceAccount returns the ServiceAccount field value.
func (o *GCPUsageCostConfigPostRequestAttributes) GetServiceAccount() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *GCPUsageCostConfigPostRequestAttributes) GetServiceAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value.
func (o *GCPUsageCostConfigPostRequestAttributes) SetServiceAccount(v string) {
	o.ServiceAccount = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPUsageCostConfigPostRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["dataset"] = o.Dataset
	toSerialize["export_prefix"] = o.ExportPrefix
	toSerialize["export_project_name"] = o.ExportProjectName
	if o.IsEnabled != nil {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if o.Months != nil {
		toSerialize["months"] = o.Months
	}
	toSerialize["service_account"] = o.ServiceAccount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPUsageCostConfigPostRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId         *string `json:"account_id"`
		BucketName        *string `json:"bucket_name"`
		Dataset           *string `json:"dataset"`
		ExportPrefix      *string `json:"export_prefix"`
		ExportProjectName *string `json:"export_project_name"`
		IsEnabled         *bool   `json:"is_enabled,omitempty"`
		Months            *int32  `json:"months,omitempty"`
		ServiceAccount    *string `json:"service_account"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.BucketName == nil {
		return fmt.Errorf("required field bucket_name missing")
	}
	if all.Dataset == nil {
		return fmt.Errorf("required field dataset missing")
	}
	if all.ExportPrefix == nil {
		return fmt.Errorf("required field export_prefix missing")
	}
	if all.ExportProjectName == nil {
		return fmt.Errorf("required field export_project_name missing")
	}
	if all.ServiceAccount == nil {
		return fmt.Errorf("required field service_account missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "bucket_name", "dataset", "export_prefix", "export_project_name", "is_enabled", "months", "service_account"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.BucketName = *all.BucketName
	o.Dataset = *all.Dataset
	o.ExportPrefix = *all.ExportPrefix
	o.ExportProjectName = *all.ExportProjectName
	o.IsEnabled = all.IsEnabled
	o.Months = all.Months
	o.ServiceAccount = *all.ServiceAccount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
