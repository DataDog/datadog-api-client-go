// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueImpactView A RUM view where the issue occurred.
type IssueImpactView struct {
	// Distinct users, accounts, sessions, and view instances on a view. `affected` counts only those with the issue's
	// errors. `baseline` counts all users, accounts, sessions, and view instances on the view in the time window. It
	// is absent when RUM has no data for the view.
	Affected IssueImpactViewCounts `json:"affected"`
	// Distinct users, accounts, sessions, and view instances on a view. `affected` counts only those with the issue's
	// errors. `baseline` counts all users, accounts, sessions, and view instances on the view in the time window. It
	// is absent when RUM has no data for the view.
	Baseline *IssueImpactViewCounts `json:"baseline,omitempty"`
	// Number of errors reported by each source. Several sources can report the same error, such as an APM span and
	// a log, so the counts overlap: don't add them together.
	ErrorCounts IssueImpactErrorCounts `json:"error_counts"`
	// View name.
	Name string `json:"name"`
	// Whether the entry stands out as a likely driver of the issue, because it accounts for most of the issue's
	// errors or for far more errors than the other listed entries.
	Significant bool `json:"significant"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueImpactView instantiates a new IssueImpactView object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueImpactView(affected IssueImpactViewCounts, errorCounts IssueImpactErrorCounts, name string, significant bool) *IssueImpactView {
	this := IssueImpactView{}
	this.Affected = affected
	this.ErrorCounts = errorCounts
	this.Name = name
	this.Significant = significant
	return &this
}

// NewIssueImpactViewWithDefaults instantiates a new IssueImpactView object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueImpactViewWithDefaults() *IssueImpactView {
	this := IssueImpactView{}
	return &this
}

// GetAffected returns the Affected field value.
func (o *IssueImpactView) GetAffected() IssueImpactViewCounts {
	if o == nil {
		var ret IssueImpactViewCounts
		return ret
	}
	return o.Affected
}

// GetAffectedOk returns a tuple with the Affected field value
// and a boolean to check if the value has been set.
func (o *IssueImpactView) GetAffectedOk() (*IssueImpactViewCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Affected, true
}

// SetAffected sets field value.
func (o *IssueImpactView) SetAffected(v IssueImpactViewCounts) {
	o.Affected = v
}

// GetBaseline returns the Baseline field value if set, zero value otherwise.
func (o *IssueImpactView) GetBaseline() IssueImpactViewCounts {
	if o == nil || o.Baseline == nil {
		var ret IssueImpactViewCounts
		return ret
	}
	return *o.Baseline
}

// GetBaselineOk returns a tuple with the Baseline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueImpactView) GetBaselineOk() (*IssueImpactViewCounts, bool) {
	if o == nil || o.Baseline == nil {
		return nil, false
	}
	return o.Baseline, true
}

// HasBaseline returns a boolean if a field has been set.
func (o *IssueImpactView) HasBaseline() bool {
	return o != nil && o.Baseline != nil
}

// SetBaseline gets a reference to the given IssueImpactViewCounts and assigns it to the Baseline field.
func (o *IssueImpactView) SetBaseline(v IssueImpactViewCounts) {
	o.Baseline = &v
}

// GetErrorCounts returns the ErrorCounts field value.
func (o *IssueImpactView) GetErrorCounts() IssueImpactErrorCounts {
	if o == nil {
		var ret IssueImpactErrorCounts
		return ret
	}
	return o.ErrorCounts
}

// GetErrorCountsOk returns a tuple with the ErrorCounts field value
// and a boolean to check if the value has been set.
func (o *IssueImpactView) GetErrorCountsOk() (*IssueImpactErrorCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCounts, true
}

// SetErrorCounts sets field value.
func (o *IssueImpactView) SetErrorCounts(v IssueImpactErrorCounts) {
	o.ErrorCounts = v
}

// GetName returns the Name field value.
func (o *IssueImpactView) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IssueImpactView) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IssueImpactView) SetName(v string) {
	o.Name = v
}

// GetSignificant returns the Significant field value.
func (o *IssueImpactView) GetSignificant() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Significant
}

// GetSignificantOk returns a tuple with the Significant field value
// and a boolean to check if the value has been set.
func (o *IssueImpactView) GetSignificantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Significant, true
}

// SetSignificant sets field value.
func (o *IssueImpactView) SetSignificant(v bool) {
	o.Significant = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueImpactView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["affected"] = o.Affected
	if o.Baseline != nil {
		toSerialize["baseline"] = o.Baseline
	}
	toSerialize["error_counts"] = o.ErrorCounts
	toSerialize["name"] = o.Name
	toSerialize["significant"] = o.Significant

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueImpactView) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Affected    *IssueImpactViewCounts  `json:"affected"`
		Baseline    *IssueImpactViewCounts  `json:"baseline,omitempty"`
		ErrorCounts *IssueImpactErrorCounts `json:"error_counts"`
		Name        *string                 `json:"name"`
		Significant *bool                   `json:"significant"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Affected == nil {
		return fmt.Errorf("required field affected missing")
	}
	if all.ErrorCounts == nil {
		return fmt.Errorf("required field error_counts missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Significant == nil {
		return fmt.Errorf("required field significant missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"affected", "baseline", "error_counts", "name", "significant"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Affected.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Affected = *all.Affected
	if all.Baseline != nil && all.Baseline.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Baseline = all.Baseline
	if all.ErrorCounts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ErrorCounts = *all.ErrorCounts
	o.Name = *all.Name
	o.Significant = *all.Significant

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
