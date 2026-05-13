// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PROutput A pull request created by the dependency upgrade automation.
type PROutput struct {
	// The aggregate CI check status for the pull request.
	CiStatus *PROutputCiStatus `json:"ci_status,omitempty"`
	// The pull request number.
	PrNumber *int64 `json:"pr_number,omitempty"`
	// The URL of the pull request.
	PrUrl string `json:"pr_url"`
	// The repository name in owner/repo format.
	Repository *string `json:"repository,omitempty"`
	// The current status of the pull request.
	Status *PROutputStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPROutput instantiates a new PROutput object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPROutput(prUrl string) *PROutput {
	this := PROutput{}
	this.PrUrl = prUrl
	return &this
}

// NewPROutputWithDefaults instantiates a new PROutput object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPROutputWithDefaults() *PROutput {
	this := PROutput{}
	return &this
}

// GetCiStatus returns the CiStatus field value if set, zero value otherwise.
func (o *PROutput) GetCiStatus() PROutputCiStatus {
	if o == nil || o.CiStatus == nil {
		var ret PROutputCiStatus
		return ret
	}
	return *o.CiStatus
}

// GetCiStatusOk returns a tuple with the CiStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PROutput) GetCiStatusOk() (*PROutputCiStatus, bool) {
	if o == nil || o.CiStatus == nil {
		return nil, false
	}
	return o.CiStatus, true
}

// HasCiStatus returns a boolean if a field has been set.
func (o *PROutput) HasCiStatus() bool {
	return o != nil && o.CiStatus != nil
}

// SetCiStatus gets a reference to the given PROutputCiStatus and assigns it to the CiStatus field.
func (o *PROutput) SetCiStatus(v PROutputCiStatus) {
	o.CiStatus = &v
}

// GetPrNumber returns the PrNumber field value if set, zero value otherwise.
func (o *PROutput) GetPrNumber() int64 {
	if o == nil || o.PrNumber == nil {
		var ret int64
		return ret
	}
	return *o.PrNumber
}

// GetPrNumberOk returns a tuple with the PrNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PROutput) GetPrNumberOk() (*int64, bool) {
	if o == nil || o.PrNumber == nil {
		return nil, false
	}
	return o.PrNumber, true
}

// HasPrNumber returns a boolean if a field has been set.
func (o *PROutput) HasPrNumber() bool {
	return o != nil && o.PrNumber != nil
}

// SetPrNumber gets a reference to the given int64 and assigns it to the PrNumber field.
func (o *PROutput) SetPrNumber(v int64) {
	o.PrNumber = &v
}

// GetPrUrl returns the PrUrl field value.
func (o *PROutput) GetPrUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PrUrl
}

// GetPrUrlOk returns a tuple with the PrUrl field value
// and a boolean to check if the value has been set.
func (o *PROutput) GetPrUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrUrl, true
}

// SetPrUrl sets field value.
func (o *PROutput) SetPrUrl(v string) {
	o.PrUrl = v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *PROutput) GetRepository() string {
	if o == nil || o.Repository == nil {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PROutput) GetRepositoryOk() (*string, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *PROutput) HasRepository() bool {
	return o != nil && o.Repository != nil
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *PROutput) SetRepository(v string) {
	o.Repository = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PROutput) GetStatus() PROutputStatus {
	if o == nil || o.Status == nil {
		var ret PROutputStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PROutput) GetStatusOk() (*PROutputStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PROutput) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given PROutputStatus and assigns it to the Status field.
func (o *PROutput) SetStatus(v PROutputStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PROutput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CiStatus != nil {
		toSerialize["ci_status"] = o.CiStatus
	}
	if o.PrNumber != nil {
		toSerialize["pr_number"] = o.PrNumber
	}
	toSerialize["pr_url"] = o.PrUrl
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PROutput) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CiStatus   *PROutputCiStatus `json:"ci_status,omitempty"`
		PrNumber   *int64            `json:"pr_number,omitempty"`
		PrUrl      *string           `json:"pr_url"`
		Repository *string           `json:"repository,omitempty"`
		Status     *PROutputStatus   `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PrUrl == nil {
		return fmt.Errorf("required field pr_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ci_status", "pr_number", "pr_url", "repository", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CiStatus != nil && !all.CiStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.CiStatus = all.CiStatus
	}
	o.PrNumber = all.PrNumber
	o.PrUrl = *all.PrUrl
	o.Repository = all.Repository
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
