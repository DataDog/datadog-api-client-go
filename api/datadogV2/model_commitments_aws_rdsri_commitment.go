// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsAwsRDSRICommitment AWS RDS Reserved Instance commitment details.
type CommitmentsAwsRDSRICommitment struct {
	// The unique identifier of the Reserved Instance.
	CommitmentId string `json:"commitment_id"`
	// The database engine of the Reserved Instance.
	DatabaseEngine string `json:"database_engine"`
	// The expiration date of the commitment.
	ExpirationDate *string `json:"expiration_date,omitempty"`
	// The RDS instance type.
	InstanceType string `json:"instance_type"`
	// Whether the Reserved Instance is Multi-AZ.
	IsMultiAz *bool `json:"is_multi_az,omitempty"`
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

// NewCommitmentsAwsRDSRICommitment instantiates a new CommitmentsAwsRDSRICommitment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsAwsRDSRICommitment(commitmentId string, databaseEngine string, instanceType string, purchaseOption string, region string) *CommitmentsAwsRDSRICommitment {
	this := CommitmentsAwsRDSRICommitment{}
	this.CommitmentId = commitmentId
	this.DatabaseEngine = databaseEngine
	this.InstanceType = instanceType
	this.PurchaseOption = purchaseOption
	this.Region = region
	return &this
}

// NewCommitmentsAwsRDSRICommitmentWithDefaults instantiates a new CommitmentsAwsRDSRICommitment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsAwsRDSRICommitmentWithDefaults() *CommitmentsAwsRDSRICommitment {
	this := CommitmentsAwsRDSRICommitment{}
	return &this
}

// GetCommitmentId returns the CommitmentId field value.
func (o *CommitmentsAwsRDSRICommitment) GetCommitmentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitmentId
}

// GetCommitmentIdOk returns a tuple with the CommitmentId field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetCommitmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitmentId, true
}

// SetCommitmentId sets field value.
func (o *CommitmentsAwsRDSRICommitment) SetCommitmentId(v string) {
	o.CommitmentId = v
}

// GetDatabaseEngine returns the DatabaseEngine field value.
func (o *CommitmentsAwsRDSRICommitment) GetDatabaseEngine() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatabaseEngine
}

// GetDatabaseEngineOk returns a tuple with the DatabaseEngine field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetDatabaseEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseEngine, true
}

// SetDatabaseEngine sets field value.
func (o *CommitmentsAwsRDSRICommitment) SetDatabaseEngine(v string) {
	o.DatabaseEngine = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CommitmentsAwsRDSRICommitment) GetExpirationDate() string {
	if o == nil || o.ExpirationDate == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetExpirationDateOk() (*string, bool) {
	if o == nil || o.ExpirationDate == nil {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CommitmentsAwsRDSRICommitment) HasExpirationDate() bool {
	return o != nil && o.ExpirationDate != nil
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CommitmentsAwsRDSRICommitment) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetInstanceType returns the InstanceType field value.
func (o *CommitmentsAwsRDSRICommitment) GetInstanceType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetInstanceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceType, true
}

// SetInstanceType sets field value.
func (o *CommitmentsAwsRDSRICommitment) SetInstanceType(v string) {
	o.InstanceType = v
}

// GetIsMultiAz returns the IsMultiAz field value if set, zero value otherwise.
func (o *CommitmentsAwsRDSRICommitment) GetIsMultiAz() bool {
	if o == nil || o.IsMultiAz == nil {
		var ret bool
		return ret
	}
	return *o.IsMultiAz
}

// GetIsMultiAzOk returns a tuple with the IsMultiAz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetIsMultiAzOk() (*bool, bool) {
	if o == nil || o.IsMultiAz == nil {
		return nil, false
	}
	return o.IsMultiAz, true
}

// HasIsMultiAz returns a boolean if a field has been set.
func (o *CommitmentsAwsRDSRICommitment) HasIsMultiAz() bool {
	return o != nil && o.IsMultiAz != nil
}

// SetIsMultiAz gets a reference to the given bool and assigns it to the IsMultiAz field.
func (o *CommitmentsAwsRDSRICommitment) SetIsMultiAz(v bool) {
	o.IsMultiAz = &v
}

// GetNumberOfNfus returns the NumberOfNfus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitmentsAwsRDSRICommitment) GetNumberOfNfus() float64 {
	if o == nil || o.NumberOfNfus.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NumberOfNfus.Get()
}

// GetNumberOfNfusOk returns a tuple with the NumberOfNfus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CommitmentsAwsRDSRICommitment) GetNumberOfNfusOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberOfNfus.Get(), o.NumberOfNfus.IsSet()
}

// HasNumberOfNfus returns a boolean if a field has been set.
func (o *CommitmentsAwsRDSRICommitment) HasNumberOfNfus() bool {
	return o != nil && o.NumberOfNfus.IsSet()
}

// SetNumberOfNfus gets a reference to the given datadog.NullableFloat64 and assigns it to the NumberOfNfus field.
func (o *CommitmentsAwsRDSRICommitment) SetNumberOfNfus(v float64) {
	o.NumberOfNfus.Set(&v)
}

// SetNumberOfNfusNil sets the value for NumberOfNfus to be an explicit nil.
func (o *CommitmentsAwsRDSRICommitment) SetNumberOfNfusNil() {
	o.NumberOfNfus.Set(nil)
}

// UnsetNumberOfNfus ensures that no value is present for NumberOfNfus, not even an explicit nil.
func (o *CommitmentsAwsRDSRICommitment) UnsetNumberOfNfus() {
	o.NumberOfNfus.Unset()
}

// GetNumberOfReservations returns the NumberOfReservations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommitmentsAwsRDSRICommitment) GetNumberOfReservations() float64 {
	if o == nil || o.NumberOfReservations.Get() == nil {
		var ret float64
		return ret
	}
	return *o.NumberOfReservations.Get()
}

// GetNumberOfReservationsOk returns a tuple with the NumberOfReservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CommitmentsAwsRDSRICommitment) GetNumberOfReservationsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumberOfReservations.Get(), o.NumberOfReservations.IsSet()
}

// HasNumberOfReservations returns a boolean if a field has been set.
func (o *CommitmentsAwsRDSRICommitment) HasNumberOfReservations() bool {
	return o != nil && o.NumberOfReservations.IsSet()
}

// SetNumberOfReservations gets a reference to the given datadog.NullableFloat64 and assigns it to the NumberOfReservations field.
func (o *CommitmentsAwsRDSRICommitment) SetNumberOfReservations(v float64) {
	o.NumberOfReservations.Set(&v)
}

// SetNumberOfReservationsNil sets the value for NumberOfReservations to be an explicit nil.
func (o *CommitmentsAwsRDSRICommitment) SetNumberOfReservationsNil() {
	o.NumberOfReservations.Set(nil)
}

// UnsetNumberOfReservations ensures that no value is present for NumberOfReservations, not even an explicit nil.
func (o *CommitmentsAwsRDSRICommitment) UnsetNumberOfReservations() {
	o.NumberOfReservations.Unset()
}

// GetPurchaseOption returns the PurchaseOption field value.
func (o *CommitmentsAwsRDSRICommitment) GetPurchaseOption() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PurchaseOption
}

// GetPurchaseOptionOk returns a tuple with the PurchaseOption field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetPurchaseOptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchaseOption, true
}

// SetPurchaseOption sets field value.
func (o *CommitmentsAwsRDSRICommitment) SetPurchaseOption(v string) {
	o.PurchaseOption = v
}

// GetRegion returns the Region field value.
func (o *CommitmentsAwsRDSRICommitment) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *CommitmentsAwsRDSRICommitment) SetRegion(v string) {
	o.Region = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CommitmentsAwsRDSRICommitment) GetStartDate() string {
	if o == nil || o.StartDate == nil {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetStartDateOk() (*string, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CommitmentsAwsRDSRICommitment) HasStartDate() bool {
	return o != nil && o.StartDate != nil
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CommitmentsAwsRDSRICommitment) SetStartDate(v string) {
	o.StartDate = &v
}

// GetTermLengthMonths returns the TermLengthMonths field value if set, zero value otherwise.
func (o *CommitmentsAwsRDSRICommitment) GetTermLengthMonths() int64 {
	if o == nil || o.TermLengthMonths == nil {
		var ret int64
		return ret
	}
	return *o.TermLengthMonths
}

// GetTermLengthMonthsOk returns a tuple with the TermLengthMonths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetTermLengthMonthsOk() (*int64, bool) {
	if o == nil || o.TermLengthMonths == nil {
		return nil, false
	}
	return o.TermLengthMonths, true
}

// HasTermLengthMonths returns a boolean if a field has been set.
func (o *CommitmentsAwsRDSRICommitment) HasTermLengthMonths() bool {
	return o != nil && o.TermLengthMonths != nil
}

// SetTermLengthMonths gets a reference to the given int64 and assigns it to the TermLengthMonths field.
func (o *CommitmentsAwsRDSRICommitment) SetTermLengthMonths(v int64) {
	o.TermLengthMonths = &v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *CommitmentsAwsRDSRICommitment) GetUtilization() float64 {
	if o == nil || o.Utilization == nil {
		var ret float64
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitmentsAwsRDSRICommitment) GetUtilizationOk() (*float64, bool) {
	if o == nil || o.Utilization == nil {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *CommitmentsAwsRDSRICommitment) HasUtilization() bool {
	return o != nil && o.Utilization != nil
}

// SetUtilization gets a reference to the given float64 and assigns it to the Utilization field.
func (o *CommitmentsAwsRDSRICommitment) SetUtilization(v float64) {
	o.Utilization = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsAwsRDSRICommitment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commitment_id"] = o.CommitmentId
	toSerialize["database_engine"] = o.DatabaseEngine
	if o.ExpirationDate != nil {
		toSerialize["expiration_date"] = o.ExpirationDate
	}
	toSerialize["instance_type"] = o.InstanceType
	if o.IsMultiAz != nil {
		toSerialize["is_multi_az"] = o.IsMultiAz
	}
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
func (o *CommitmentsAwsRDSRICommitment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitmentId         *string                 `json:"commitment_id"`
		DatabaseEngine       *string                 `json:"database_engine"`
		ExpirationDate       *string                 `json:"expiration_date,omitempty"`
		InstanceType         *string                 `json:"instance_type"`
		IsMultiAz            *bool                   `json:"is_multi_az,omitempty"`
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
	if all.CommitmentId == nil {
		return fmt.Errorf("required field commitment_id missing")
	}
	if all.DatabaseEngine == nil {
		return fmt.Errorf("required field database_engine missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"commitment_id", "database_engine", "expiration_date", "instance_type", "is_multi_az", "number_of_nfus", "number_of_reservations", "purchase_option", "region", "start_date", "term_length_months", "utilization"})
	} else {
		return err
	}
	o.CommitmentId = *all.CommitmentId
	o.DatabaseEngine = *all.DatabaseEngine
	o.ExpirationDate = all.ExpirationDate
	o.InstanceType = *all.InstanceType
	o.IsMultiAz = all.IsMultiAz
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
