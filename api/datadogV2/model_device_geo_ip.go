// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceGeoIP Location derived from the device's public IP address.
// Absent when no location data is available.
type DeviceGeoIP struct {
	// City resolved from the device's public IP address.
	GeoCity *string `json:"geo_city,omitempty"`
	// ISO 3166-1 alpha-2 country code resolved from the device's public IP address.
	GeoCountryIsoCode *string `json:"geo_country_iso_code,omitempty"`
	// Latitude resolved from the device's public IP address.
	GeoLatitude *float64 `json:"geo_latitude,omitempty"`
	// Longitude resolved from the device's public IP address.
	GeoLongitude *float64 `json:"geo_longitude,omitempty"`
	// ISO 3166-2 subdivision code (for example, a state or province)
	// resolved from the device's public IP address.
	GeoSubdivisionIsoCode *string `json:"geo_subdivision_iso_code,omitempty"`
	// Public IPv4 or IPv6 address of the device.
	PublicIpAddress *string `json:"public_ip_address,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeviceGeoIP instantiates a new DeviceGeoIP object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeviceGeoIP() *DeviceGeoIP {
	this := DeviceGeoIP{}
	return &this
}

// NewDeviceGeoIPWithDefaults instantiates a new DeviceGeoIP object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeviceGeoIPWithDefaults() *DeviceGeoIP {
	this := DeviceGeoIP{}
	return &this
}

// GetGeoCity returns the GeoCity field value if set, zero value otherwise.
func (o *DeviceGeoIP) GetGeoCity() string {
	if o == nil || o.GeoCity == nil {
		var ret string
		return ret
	}
	return *o.GeoCity
}

// GetGeoCityOk returns a tuple with the GeoCity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGeoIP) GetGeoCityOk() (*string, bool) {
	if o == nil || o.GeoCity == nil {
		return nil, false
	}
	return o.GeoCity, true
}

// HasGeoCity returns a boolean if a field has been set.
func (o *DeviceGeoIP) HasGeoCity() bool {
	return o != nil && o.GeoCity != nil
}

// SetGeoCity gets a reference to the given string and assigns it to the GeoCity field.
func (o *DeviceGeoIP) SetGeoCity(v string) {
	o.GeoCity = &v
}

// GetGeoCountryIsoCode returns the GeoCountryIsoCode field value if set, zero value otherwise.
func (o *DeviceGeoIP) GetGeoCountryIsoCode() string {
	if o == nil || o.GeoCountryIsoCode == nil {
		var ret string
		return ret
	}
	return *o.GeoCountryIsoCode
}

// GetGeoCountryIsoCodeOk returns a tuple with the GeoCountryIsoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGeoIP) GetGeoCountryIsoCodeOk() (*string, bool) {
	if o == nil || o.GeoCountryIsoCode == nil {
		return nil, false
	}
	return o.GeoCountryIsoCode, true
}

// HasGeoCountryIsoCode returns a boolean if a field has been set.
func (o *DeviceGeoIP) HasGeoCountryIsoCode() bool {
	return o != nil && o.GeoCountryIsoCode != nil
}

// SetGeoCountryIsoCode gets a reference to the given string and assigns it to the GeoCountryIsoCode field.
func (o *DeviceGeoIP) SetGeoCountryIsoCode(v string) {
	o.GeoCountryIsoCode = &v
}

// GetGeoLatitude returns the GeoLatitude field value if set, zero value otherwise.
func (o *DeviceGeoIP) GetGeoLatitude() float64 {
	if o == nil || o.GeoLatitude == nil {
		var ret float64
		return ret
	}
	return *o.GeoLatitude
}

// GetGeoLatitudeOk returns a tuple with the GeoLatitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGeoIP) GetGeoLatitudeOk() (*float64, bool) {
	if o == nil || o.GeoLatitude == nil {
		return nil, false
	}
	return o.GeoLatitude, true
}

// HasGeoLatitude returns a boolean if a field has been set.
func (o *DeviceGeoIP) HasGeoLatitude() bool {
	return o != nil && o.GeoLatitude != nil
}

// SetGeoLatitude gets a reference to the given float64 and assigns it to the GeoLatitude field.
func (o *DeviceGeoIP) SetGeoLatitude(v float64) {
	o.GeoLatitude = &v
}

// GetGeoLongitude returns the GeoLongitude field value if set, zero value otherwise.
func (o *DeviceGeoIP) GetGeoLongitude() float64 {
	if o == nil || o.GeoLongitude == nil {
		var ret float64
		return ret
	}
	return *o.GeoLongitude
}

// GetGeoLongitudeOk returns a tuple with the GeoLongitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGeoIP) GetGeoLongitudeOk() (*float64, bool) {
	if o == nil || o.GeoLongitude == nil {
		return nil, false
	}
	return o.GeoLongitude, true
}

// HasGeoLongitude returns a boolean if a field has been set.
func (o *DeviceGeoIP) HasGeoLongitude() bool {
	return o != nil && o.GeoLongitude != nil
}

// SetGeoLongitude gets a reference to the given float64 and assigns it to the GeoLongitude field.
func (o *DeviceGeoIP) SetGeoLongitude(v float64) {
	o.GeoLongitude = &v
}

// GetGeoSubdivisionIsoCode returns the GeoSubdivisionIsoCode field value if set, zero value otherwise.
func (o *DeviceGeoIP) GetGeoSubdivisionIsoCode() string {
	if o == nil || o.GeoSubdivisionIsoCode == nil {
		var ret string
		return ret
	}
	return *o.GeoSubdivisionIsoCode
}

// GetGeoSubdivisionIsoCodeOk returns a tuple with the GeoSubdivisionIsoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGeoIP) GetGeoSubdivisionIsoCodeOk() (*string, bool) {
	if o == nil || o.GeoSubdivisionIsoCode == nil {
		return nil, false
	}
	return o.GeoSubdivisionIsoCode, true
}

// HasGeoSubdivisionIsoCode returns a boolean if a field has been set.
func (o *DeviceGeoIP) HasGeoSubdivisionIsoCode() bool {
	return o != nil && o.GeoSubdivisionIsoCode != nil
}

// SetGeoSubdivisionIsoCode gets a reference to the given string and assigns it to the GeoSubdivisionIsoCode field.
func (o *DeviceGeoIP) SetGeoSubdivisionIsoCode(v string) {
	o.GeoSubdivisionIsoCode = &v
}

// GetPublicIpAddress returns the PublicIpAddress field value if set, zero value otherwise.
func (o *DeviceGeoIP) GetPublicIpAddress() string {
	if o == nil || o.PublicIpAddress == nil {
		var ret string
		return ret
	}
	return *o.PublicIpAddress
}

// GetPublicIpAddressOk returns a tuple with the PublicIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceGeoIP) GetPublicIpAddressOk() (*string, bool) {
	if o == nil || o.PublicIpAddress == nil {
		return nil, false
	}
	return o.PublicIpAddress, true
}

// HasPublicIpAddress returns a boolean if a field has been set.
func (o *DeviceGeoIP) HasPublicIpAddress() bool {
	return o != nil && o.PublicIpAddress != nil
}

// SetPublicIpAddress gets a reference to the given string and assigns it to the PublicIpAddress field.
func (o *DeviceGeoIP) SetPublicIpAddress(v string) {
	o.PublicIpAddress = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeviceGeoIP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GeoCity != nil {
		toSerialize["geo_city"] = o.GeoCity
	}
	if o.GeoCountryIsoCode != nil {
		toSerialize["geo_country_iso_code"] = o.GeoCountryIsoCode
	}
	if o.GeoLatitude != nil {
		toSerialize["geo_latitude"] = o.GeoLatitude
	}
	if o.GeoLongitude != nil {
		toSerialize["geo_longitude"] = o.GeoLongitude
	}
	if o.GeoSubdivisionIsoCode != nil {
		toSerialize["geo_subdivision_iso_code"] = o.GeoSubdivisionIsoCode
	}
	if o.PublicIpAddress != nil {
		toSerialize["public_ip_address"] = o.PublicIpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeviceGeoIP) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GeoCity               *string  `json:"geo_city,omitempty"`
		GeoCountryIsoCode     *string  `json:"geo_country_iso_code,omitempty"`
		GeoLatitude           *float64 `json:"geo_latitude,omitempty"`
		GeoLongitude          *float64 `json:"geo_longitude,omitempty"`
		GeoSubdivisionIsoCode *string  `json:"geo_subdivision_iso_code,omitempty"`
		PublicIpAddress       *string  `json:"public_ip_address,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"geo_city", "geo_country_iso_code", "geo_latitude", "geo_longitude", "geo_subdivision_iso_code", "public_ip_address"})
	} else {
		return err
	}
	o.GeoCity = all.GeoCity
	o.GeoCountryIsoCode = all.GeoCountryIsoCode
	o.GeoLatitude = all.GeoLatitude
	o.GeoLongitude = all.GeoLongitude
	o.GeoSubdivisionIsoCode = all.GeoSubdivisionIsoCode
	o.PublicIpAddress = all.PublicIpAddress

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
