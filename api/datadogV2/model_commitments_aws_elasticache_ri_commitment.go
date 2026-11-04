// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsAwsElasticacheRICommitment AWS ElastiCache Reserved Instance commitment details.
type CommitmentsAwsElasticacheRICommitment struct {
	// The cache engine type of the Reserved Instance.
	CacheEngine string `json:"cache_engine"`
	// The unique identifier of the Reserved Instance.
	CommitmentId string `json:"commitment_id"`
	// The expiration date of the commitment.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The ElastiCache instance type.
	InstanceType string `json:"instance_type"`
	// The number of Normalized Capacity Units.
	NumberOfNfus datadog.NullableFloat64 `json:"number_of_nfus,omitempty"`
	// The number of reserved instances.
	NumberOfReservations datadog.NullableFloat64 `json:"number_of_reservations,omitempty"`
	// The payment option for the Reserved Instance.
	PurchaseOption string `json:"purchase_option"`
	// The AWS region of the Reserved Instance.
	Region string `json:"region"`
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

// NewCommitmentsAwsElasticacheRICommitment instantiates a new CommitmentsAwsElasticacheRICommitment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsAwsElasticacheRICommitment(cacheEngine string, commitmentId string, instanceType string, purchaseOption string, region string) *CommitmentsAwsElasticacheRICommitment {
	this := CommitmentsAwsElasticacheRICommitment{}
	this.CacheEngine = cacheEngine
	this.CommitmentId = commitmentId
	this.InstanceType = instanceType
	this.PurchaseOption = purchaseOption
	this.Region = region
	return &this
}

// NewCommitmentsAwsElasticacheRICommitmentWithDefaults instantiates a new CommitmentsAwsElasticacheRICommitment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsAwsElasticacheRICommitmentWithDefaults() *CommitmentsAwsElasticacheRICommitment {
	this := CommitmentsAwsElasticacheRICommitment{}
	return &this
}

// GetCacheEngine returns the CacheEngine field value.
func (o *CommitmentsAwsElasticacheRICommitment) GetCacheEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CacheEngine
}

// GetCacheEngineOk returns a tuple with the CacheEngine field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetCacheEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CacheEngine, true
}

// SetCacheEngine sets field value.
func (o *CommitmentsAwsElasticacheRICommitment) SetCacheEngine(v string) {
	o.CacheEngine = v
}

// GetCommitmentId returns the CommitmentId field value.
func (o *CommitmentsAwsElasticacheRICommitment) GetCommitmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitmentId
}

// GetCommitmentIdOk returns a tuple with the CommitmentId field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetCommitmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentId, true
}

// SetCommitmentId sets field value.
func (o *CommitmentsAwsElasticacheRICommitment) SetCommitmentId(v string) {
	o.CommitmentId = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CommitmentsAwsElasticacheRICommitment) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CommitmentsAwsElasticacheRICommitment) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CommitmentsAwsElasticacheRICommitment) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetInstanceType returns the InstanceType field value.
func (o *CommitmentsAwsElasticacheRICommitment) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value.
func (o *CommitmentsAwsElasticacheRICommitment) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetNumberOfNfus returns the NumberOfNfus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitmentsAwsElasticacheRICommitment) GetNumberOfNfus() float64 {
	if o == nil || o.NumberOfNfus.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NumberOfNfus.Get()
}

// GetNumberOfNfusOk returns a tuple with the NumberOfNfus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CommitmentsAwsElasticacheRICommitment) GetNumberOfNfusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberOfNfus.Get(), o.NumberOfNfus.IsSet()
}

// HasNumberOfNfus returns a boolean if a field has been set.
func (o *CommitmentsAwsElasticacheRICommitment) HasNumberOfNfus() bool {
	return o != nil && o.NumberOfNfus.IsSet()
}

// SetNumberOfNfus gets a reference to the given datadog.NullableFloat64 and assigns it to the NumberOfNfus field.
func (o *CommitmentsAwsElasticacheRICommitment) SetNumberOfNfus(v float64) {
	o.NumberOfNfus.Set(&v)
}

// SetNumberOfNfusNil sets the value for NumberOfNfus to be an explicit nil.
func (o *CommitmentsAwsElasticacheRICommitment) SetNumberOfNfusNil() {
	o.NumberOfNfus.Set(nil)
}

// UnsetNumberOfNfus ensures that no value is present for NumberOfNfus, not even an explicit nil.
func (o *CommitmentsAwsElasticacheRICommitment) UnsetNumberOfNfus() {
	o.NumberOfNfus.Unset()
}

// GetNumberOfReservations returns the NumberOfReservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitmentsAwsElasticacheRICommitment) GetNumberOfReservations() float64 {
	if o == nil || o.NumberOfReservations.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NumberOfReservations.Get()
}

// GetNumberOfReservationsOk returns a tuple with the NumberOfReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CommitmentsAwsElasticacheRICommitment) GetNumberOfReservationsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberOfReservations.Get(), o.NumberOfReservations.IsSet()
}

// HasNumberOfReservations returns a boolean if a field has been set.
func (o *CommitmentsAwsElasticacheRICommitment) HasNumberOfReservations() bool {
	return o != nil && o.NumberOfReservations.IsSet()
}

// SetNumberOfReservations gets a reference to the given datadog.NullableFloat64 and assigns it to the NumberOfReservations field.
func (o *CommitmentsAwsElasticacheRICommitment) SetNumberOfReservations(v float64) {
	o.NumberOfReservations.Set(&v)
}

// SetNumberOfReservationsNil sets the value for NumberOfReservations to be an explicit nil.
func (o *CommitmentsAwsElasticacheRICommitment) SetNumberOfReservationsNil() {
	o.NumberOfReservations.Set(nil)
}

// UnsetNumberOfReservations ensures that no value is present for NumberOfReservations, not even an explicit nil.
func (o *CommitmentsAwsElasticacheRICommitment) UnsetNumberOfReservations() {
	o.NumberOfReservations.Unset()
}

// GetPurchaseOption returns the PurchaseOption field value.
func (o *CommitmentsAwsElasticacheRICommitment) GetPurchaseOption() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PurchaseOption
}

// GetPurchaseOptionOk returns a tuple with the PurchaseOption field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetPurchaseOptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOption, true
}

// SetPurchaseOption sets field value.
func (o *CommitmentsAwsElasticacheRICommitment) SetPurchaseOption(v string) {
	o.PurchaseOption = v
}

// GetRegion returns the Region field value.
func (o *CommitmentsAwsElasticacheRICommitment) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *CommitmentsAwsElasticacheRICommitment) SetRegion(v string) {
	o.Region = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CommitmentsAwsElasticacheRICommitment) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CommitmentsAwsElasticacheRICommitment) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CommitmentsAwsElasticacheRICommitment) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTermLengthMonths returns the TermLengthMonths field value if set, zero value otherwise.
func (o *CommitmentsAwsElasticacheRICommitment) GetTermLengthMonths() int64 {
	if o == nil || o.TermLengthMonths == nil {
		var ret int64
		return ret
	}
	return *o.TermLengthMonths
}

// GetTermLengthMonthsOk returns a tuple with the TermLengthMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetTermLengthMonthsOk() (*int64, bool) {
	if o == nil || o.TermLengthMonths == nil {
		return nil, false
	}
	return o.TermLengthMonths, true
}

// HasTermLengthMonths returns a boolean if a field has been set.
func (o *CommitmentsAwsElasticacheRICommitment) HasTermLengthMonths() bool {
	return o != nil && o.TermLengthMonths != nil
}

// SetTermLengthMonths gets a reference to the given int64 and assigns it to the TermLengthMonths field.
func (o *CommitmentsAwsElasticacheRICommitment) SetTermLengthMonths(v int64) {
	o.TermLengthMonths = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *CommitmentsAwsElasticacheRICommitment) GetUtilization() float64 {
	if o == nil || o.Utilization == nil {
		var ret float64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsElasticacheRICommitment) GetUtilizationOk() (*float64, bool) {
	if o == nil || o.Utilization == nil {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *CommitmentsAwsElasticacheRICommitment) HasUtilization() bool {
	return o != nil && o.Utilization != nil
}

// SetUtilization gets a reference to the given float64 and assigns it to the Utilization field.
func (o *CommitmentsAwsElasticacheRICommitment) SetUtilization(v float64) {
	o.Utilization = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsAwsElasticacheRICommitment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["cache_engine"] = o.CacheEngine
	toSerialize["commitment_id"] = o.CommitmentId
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	toSerialize["instance_type"] = o.InstanceType
	if o.NumberOfNfus.IsSet() {
		toSerialize["number_of_nfus"] = o.NumberOfNfus.Get()
	}
	if o.NumberOfReservations.IsSet() {
		toSerialize["number_of_reservations"] = o.NumberOfReservations.Get()
	}
	toSerialize["purchase_option"] = o.PurchaseOption
	toSerialize["region"] = o.Region
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
func (o *CommitmentsAwsElasticacheRICommitment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CacheEngine          *string                 `json:"cache_engine"`
		CommitmentId         *string                 `json:"commitment_id"`
		ExpirationDate       *string                 `json:"expiration_date,omitempty"`
		InstanceType         *string                 `json:"instance_type"`
		NumberOfNfus         datadog.NullableFloat64 `json:"number_of_nfus,omitempty"`
		NumberOfReservations datadog.NullableFloat64 `json:"number_of_reservations,omitempty"`
		PurchaseOption       *string                 `json:"purchase_option"`
		Region               *string                 `json:"region"`
		StartDate            *string                 `json:"start_date,omitempty"`
		TermLengthMonths     *int64                  `json:"term_length_months,omitempty"`
		Utilization          *float64                `json:"utilization,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CacheEngine == nil {
		return fmt.Errorf("required field cache_engine missing")
	}
	if all.CommitmentId == nil {
		return fmt.Errorf("required field commitment_id missing")
	}
	if all.InstanceType == nil {
		return fmt.Errorf("required field instance_type missing")
	}
	if all.PurchaseOption == nil {
		return fmt.Errorf("required field purchase_option missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cache_engine", "commitment_id", "expiration_date", "instance_type", "number_of_nfus", "number_of_reservations", "purchase_option", "region", "start_date", "term_length_months", "utilization"})
	} else {
		return err
	}
	o.CacheEngine = *all.CacheEngine
	o.CommitmentId = *all.CommitmentId
	o.ExpirationDate = all.ExpirationDate
	o.InstanceType = *all.InstanceType
	o.NumberOfNfus = all.NumberOfNfus
	o.NumberOfReservations = all.NumberOfReservations
	o.PurchaseOption = *all.PurchaseOption
	o.Region = *all.Region
	o.StartDate = all.StartDate
	o.TermLengthMonths = all.TermLengthMonths
	o.Utilization = all.Utilization

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
