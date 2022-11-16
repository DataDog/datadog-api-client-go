// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
)

// DashboardReportDestinationEmail Email destinations for a report.
type DashboardReportDestinationEmail struct {
	// List of email addresses to send reports to.
	RecipientAddresses []string `json:"recipient_addresses,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewDashboardReportDestinationEmail instantiates a new DashboardReportDestinationEmail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardReportDestinationEmail() *DashboardReportDestinationEmail {
	this := DashboardReportDestinationEmail{}
	return &this
}

// NewDashboardReportDestinationEmailWithDefaults instantiates a new DashboardReportDestinationEmail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardReportDestinationEmailWithDefaults() *DashboardReportDestinationEmail {
	this := DashboardReportDestinationEmail{}
	return &this
}

// GetRecipientAddresses returns the RecipientAddresses field value if set, zero value otherwise.
func (o *DashboardReportDestinationEmail) GetRecipientAddresses() []string {
	if o == nil || o.RecipientAddresses == nil {
		var ret []string
		return ret
	}
	return o.RecipientAddresses
}

// GetRecipientAddressesOk returns a tuple with the RecipientAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardReportDestinationEmail) GetRecipientAddressesOk() (*[]string, bool) {
	if o == nil || o.RecipientAddresses == nil {
		return nil, false
	}
	return &o.RecipientAddresses, true
}

// HasRecipientAddresses returns a boolean if a field has been set.
func (o *DashboardReportDestinationEmail) HasRecipientAddresses() bool {
	return o != nil && o.RecipientAddresses != nil
}

// SetRecipientAddresses gets a reference to the given []string and assigns it to the RecipientAddresses field.
func (o *DashboardReportDestinationEmail) SetRecipientAddresses(v []string) {
	o.RecipientAddresses = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardReportDestinationEmail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.RecipientAddresses != nil {
		toSerialize["recipient_addresses"] = o.RecipientAddresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardReportDestinationEmail) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		RecipientAddresses []string `json:"recipient_addresses,omitempty"`
	}{}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.RecipientAddresses = all.RecipientAddresses
	return nil
}
