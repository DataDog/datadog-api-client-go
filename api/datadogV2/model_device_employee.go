// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceEmployee Directory information about the employee associated with the device,
// sourced from a customer reference table.
// Absent when no reference table is configured or the device has no match.
type DeviceEmployee struct {
	// City where the employee is based.
	City *string `json:"city,omitempty"`
	// Country where the employee is based.
	Country *string `json:"country,omitempty"`
	// Department the employee belongs to.
	Department *string `json:"department,omitempty"`
	// Email address of the employee.
	Email *string `json:"email,omitempty"`
	// Full name of the employee.
	Name *string `json:"name,omitempty"`
	// Office the employee is assigned to.
	Office *string `json:"office,omitempty"`
	// Team the employee belongs to.
	Team *string `json:"team,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeviceEmployee instantiates a new DeviceEmployee object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeviceEmployee() *DeviceEmployee {
	this := DeviceEmployee{}
	return &this
}

// NewDeviceEmployeeWithDefaults instantiates a new DeviceEmployee object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeviceEmployeeWithDefaults() *DeviceEmployee {
	this := DeviceEmployee{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *DeviceEmployee) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEmployee) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *DeviceEmployee) HasCity() bool {
	return o != nil && o.City != nil
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *DeviceEmployee) SetCity(v string) {
	o.City = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DeviceEmployee) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEmployee) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DeviceEmployee) HasCountry() bool {
	return o != nil && o.Country != nil
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *DeviceEmployee) SetCountry(v string) {
	o.Country = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *DeviceEmployee) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEmployee) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *DeviceEmployee) HasDepartment() bool {
	return o != nil && o.Department != nil
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *DeviceEmployee) SetDepartment(v string) {
	o.Department = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DeviceEmployee) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEmployee) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DeviceEmployee) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *DeviceEmployee) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceEmployee) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEmployee) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceEmployee) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceEmployee) SetName(v string) {
	o.Name = &v
}

// GetOffice returns the Office field value if set, zero value otherwise.
func (o *DeviceEmployee) GetOffice() string {
	if o == nil || o.Office == nil {
		var ret string
		return ret
	}
	return *o.Office
}

// GetOfficeOk returns a tuple with the Office field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEmployee) GetOfficeOk() (*string, bool) {
	if o == nil || o.Office == nil {
		return nil, false
	}
	return o.Office, true
}

// HasOffice returns a boolean if a field has been set.
func (o *DeviceEmployee) HasOffice() bool {
	return o != nil && o.Office != nil
}

// SetOffice gets a reference to the given string and assigns it to the Office field.
func (o *DeviceEmployee) SetOffice(v string) {
	o.Office = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *DeviceEmployee) GetTeam() string {
	if o == nil || o.Team == nil {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceEmployee) GetTeamOk() (*string, bool) {
	if o == nil || o.Team == nil {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *DeviceEmployee) HasTeam() bool {
	return o != nil && o.Team != nil
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *DeviceEmployee) SetTeam(v string) {
	o.Team = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeviceEmployee) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Office != nil {
		toSerialize["office"] = o.Office
	}
	if o.Team != nil {
		toSerialize["team"] = o.Team
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeviceEmployee) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		City       *string `json:"city,omitempty"`
		Country    *string `json:"country,omitempty"`
		Department *string `json:"department,omitempty"`
		Email      *string `json:"email,omitempty"`
		Name       *string `json:"name,omitempty"`
		Office     *string `json:"office,omitempty"`
		Team       *string `json:"team,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"city", "country", "department", "email", "name", "office", "team"})
	} else {
		return err
	}
	o.City = all.City
	o.Country = all.Country
	o.Department = all.Department
	o.Email = all.Email
	o.Name = all.Name
	o.Office = all.Office
	o.Team = all.Team

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
