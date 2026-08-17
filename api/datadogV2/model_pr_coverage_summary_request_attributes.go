// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PRCoverageSummaryRequestAttributes Attributes for requesting code coverage summary for a pull request.
type PRCoverageSummaryRequestAttributes struct {
	// The pull request number. Must be a positive integer.
	PrNumber int64 `json:"pr_number"`
	// The repository URL. Accepts a full URL with or without a scheme (for example, `https://github.com/org/repo` or `github.com/org/repo`).
	RepositoryUrl string `json:"repository_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPRCoverageSummaryRequestAttributes instantiates a new PRCoverageSummaryRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPRCoverageSummaryRequestAttributes(prNumber int64, repositoryUrl string) *PRCoverageSummaryRequestAttributes {
	this := PRCoverageSummaryRequestAttributes{}
	this.PrNumber = prNumber
	this.RepositoryUrl = repositoryUrl
	return &this
}

// NewPRCoverageSummaryRequestAttributesWithDefaults instantiates a new PRCoverageSummaryRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPRCoverageSummaryRequestAttributesWithDefaults() *PRCoverageSummaryRequestAttributes {
	this := PRCoverageSummaryRequestAttributes{}
	return &this
}

// GetPrNumber returns the PrNumber field value.
func (o *PRCoverageSummaryRequestAttributes) GetPrNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PrNumber
}

// GetPrNumberOk returns a tuple with the PrNumber field value
// and a boolean to check if the value has been set.
func (o *PRCoverageSummaryRequestAttributes) GetPrNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrNumber, true
}

// SetPrNumber sets field value.
func (o *PRCoverageSummaryRequestAttributes) SetPrNumber(v int64) {
	o.PrNumber = v
}

// GetRepositoryUrl returns the RepositoryUrl field value.
func (o *PRCoverageSummaryRequestAttributes) GetRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *PRCoverageSummaryRequestAttributes) GetRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryUrl, true
}

// SetRepositoryUrl sets field value.
func (o *PRCoverageSummaryRequestAttributes) SetRepositoryUrl(v string) {
	o.RepositoryUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PRCoverageSummaryRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["pr_number"] = o.PrNumber
	toSerialize["repository_url"] = o.RepositoryUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PRCoverageSummaryRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PrNumber      *int64  `json:"pr_number"`
		RepositoryUrl *string `json:"repository_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PrNumber == nil {
		return fmt.Errorf("required field pr_number missing")
	}
	if all.RepositoryUrl == nil {
		return fmt.Errorf("required field repository_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pr_number", "repository_url"})
	} else {
		return err
	}
	o.PrNumber = *all.PrNumber
	o.RepositoryUrl = *all.RepositoryUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
