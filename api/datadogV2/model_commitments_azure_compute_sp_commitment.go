// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsAzureComputeSPCommitment Azure Compute Savings Plan commitment details.
type CommitmentsAzureComputeSPCommitment struct {
	// The display name of the Azure Savings Plan.
	BenefitName string `json:"benefit_name"`
	// The unique identifier of the Savings Plan.
	CommitmentId string `json:"commitment_id"`
	// The hourly committed spend for the Savings Plan.
	CommittedSpendPerHour *float64 `json:"committed_spend_per_hour,omitempty"`
	// The expiration date of the commitment.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The Azure Savings Plan type, `Compute` or `Database`, based on the product order name. Omitted when
	// the type can't be classified.
	SavingsPlanType *string `json:"savings_plan_type,omitempty"`
	// The start date of the commitment.
	StartDate *string `json:"start_date,omitempty"`
	// The commitment term in whole months. Omitted when the term is unknown.
	TermLengthMonths *int64 `json:"term_length_months,omitempty"`
	// The utilization percentage of the commitment, from 0 to 100.
	Utilization *float64 `json:"utilization,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsAzureComputeSPCommitment instantiates a new CommitmentsAzureComputeSPCommitment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsAzureComputeSPCommitment(benefitName string, commitmentId string) *CommitmentsAzureComputeSPCommitment {
	this := CommitmentsAzureComputeSPCommitment{}
	this.BenefitName = benefitName
	this.CommitmentId = commitmentId
	return &this
}

// NewCommitmentsAzureComputeSPCommitmentWithDefaults instantiates a new CommitmentsAzureComputeSPCommitment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsAzureComputeSPCommitmentWithDefaults() *CommitmentsAzureComputeSPCommitment {
	this := CommitmentsAzureComputeSPCommitment{}
	return &this
}

// GetBenefitName returns the BenefitName field value.
func (o *CommitmentsAzureComputeSPCommitment) GetBenefitName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BenefitName
}

// GetBenefitNameOk returns a tuple with the BenefitName field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetBenefitNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BenefitName, true
}

// SetBenefitName sets field value.
func (o *CommitmentsAzureComputeSPCommitment) SetBenefitName(v string) {
	o.BenefitName = v
}

// GetCommitmentId returns the CommitmentId field value.
func (o *CommitmentsAzureComputeSPCommitment) GetCommitmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitmentId
}

// GetCommitmentIdOk returns a tuple with the CommitmentId field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetCommitmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentId, true
}

// SetCommitmentId sets field value.
func (o *CommitmentsAzureComputeSPCommitment) SetCommitmentId(v string) {
	o.CommitmentId = v
}

// GetCommittedSpendPerHour returns the CommittedSpendPerHour field value if set, zero value otherwise.
func (o *CommitmentsAzureComputeSPCommitment) GetCommittedSpendPerHour() float64 {
	if o == nil || o.CommittedSpendPerHour == nil {
		var ret float64
		return ret
	}
	return *o.CommittedSpendPerHour
}

// GetCommittedSpendPerHourOk returns a tuple with the CommittedSpendPerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetCommittedSpendPerHourOk() (*float64, bool) {
	if o == nil || o.CommittedSpendPerHour == nil {
		return nil, false
	}
	return o.CommittedSpendPerHour, true
}

// HasCommittedSpendPerHour returns a boolean if a field has been set.
func (o *CommitmentsAzureComputeSPCommitment) HasCommittedSpendPerHour() bool {
	return o != nil && o.CommittedSpendPerHour != nil
}

// SetCommittedSpendPerHour gets a reference to the given float64 and assigns it to the CommittedSpendPerHour field.
func (o *CommitmentsAzureComputeSPCommitment) SetCommittedSpendPerHour(v float64) {
	o.CommittedSpendPerHour = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CommitmentsAzureComputeSPCommitment) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CommitmentsAzureComputeSPCommitment) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CommitmentsAzureComputeSPCommitment) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetSavingsPlanType returns the SavingsPlanType field value if set, zero value otherwise.
func (o *CommitmentsAzureComputeSPCommitment) GetSavingsPlanType() string {
	if o == nil || o.SavingsPlanType == nil {
		var ret string
		return ret
	}
	return *o.SavingsPlanType
}

// GetSavingsPlanTypeOk returns a tuple with the SavingsPlanType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetSavingsPlanTypeOk() (*string, bool) {
	if o == nil || o.SavingsPlanType == nil {
		return nil, false
	}
	return o.SavingsPlanType, true
}

// HasSavingsPlanType returns a boolean if a field has been set.
func (o *CommitmentsAzureComputeSPCommitment) HasSavingsPlanType() bool {
	return o != nil && o.SavingsPlanType != nil
}

// SetSavingsPlanType gets a reference to the given string and assigns it to the SavingsPlanType field.
func (o *CommitmentsAzureComputeSPCommitment) SetSavingsPlanType(v string) {
	o.SavingsPlanType = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CommitmentsAzureComputeSPCommitment) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CommitmentsAzureComputeSPCommitment) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CommitmentsAzureComputeSPCommitment) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTermLengthMonths returns the TermLengthMonths field value if set, zero value otherwise.
func (o *CommitmentsAzureComputeSPCommitment) GetTermLengthMonths() int64 {
	if o == nil || o.TermLengthMonths == nil {
		var ret int64
		return ret
	}
	return *o.TermLengthMonths
}

// GetTermLengthMonthsOk returns a tuple with the TermLengthMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetTermLengthMonthsOk() (*int64, bool) {
	if o == nil || o.TermLengthMonths == nil {
		return nil, false
	}
	return o.TermLengthMonths, true
}

// HasTermLengthMonths returns a boolean if a field has been set.
func (o *CommitmentsAzureComputeSPCommitment) HasTermLengthMonths() bool {
	return o != nil && o.TermLengthMonths != nil
}

// SetTermLengthMonths gets a reference to the given int64 and assigns it to the TermLengthMonths field.
func (o *CommitmentsAzureComputeSPCommitment) SetTermLengthMonths(v int64) {
	o.TermLengthMonths = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *CommitmentsAzureComputeSPCommitment) GetUtilization() float64 {
	if o == nil || o.Utilization == nil {
		var ret float64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAzureComputeSPCommitment) GetUtilizationOk() (*float64, bool) {
	if o == nil || o.Utilization == nil {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *CommitmentsAzureComputeSPCommitment) HasUtilization() bool {
	return o != nil && o.Utilization != nil
}

// SetUtilization gets a reference to the given float64 and assigns it to the Utilization field.
func (o *CommitmentsAzureComputeSPCommitment) SetUtilization(v float64) {
	o.Utilization = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsAzureComputeSPCommitment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["benefit_name"] = o.BenefitName
	toSerialize["commitment_id"] = o.CommitmentId
	if o.CommittedSpendPerHour != nil {
		toSerialize["committed_spend_per_hour"] = o.CommittedSpendPerHour
	}
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	if o.SavingsPlanType != nil {
		toSerialize["savings_plan_type"] = o.SavingsPlanType
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.TermLengthMonths != nil {
		toSerialize["term_length_months"] = o.TermLengthMonths
	}
	if o.Utilization != nil {
		toSerialize["utilization"] = o.Utilization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsAzureComputeSPCommitment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BenefitName           *string  `json:"benefit_name"`
		CommitmentId          *string  `json:"commitment_id"`
		CommittedSpendPerHour *float64 `json:"committed_spend_per_hour,omitempty"`
		ExpirationDate        *string  `json:"expiration_date,omitempty"`
		SavingsPlanType       *string  `json:"savings_plan_type,omitempty"`
		StartDate             *string  `json:"start_date,omitempty"`
		TermLengthMonths      *int64   `json:"term_length_months,omitempty"`
		Utilization           *float64 `json:"utilization,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BenefitName == nil {
		return fmt.Errorf("required field benefit_name missing")
	}
	if all.CommitmentId == nil {
		return fmt.Errorf("required field commitment_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"benefit_name", "commitment_id", "committed_spend_per_hour", "expiration_date", "savings_plan_type", "start_date", "term_length_months", "utilization"})
	} else {
		return err
	}
	o.BenefitName = *all.BenefitName
	o.CommitmentId = *all.CommitmentId
	o.CommittedSpendPerHour = all.CommittedSpendPerHour
	o.ExpirationDate = all.ExpirationDate
	o.SavingsPlanType = all.SavingsPlanType
	o.StartDate = all.StartDate
	o.TermLengthMonths = all.TermLengthMonths
	o.Utilization = all.Utilization

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
