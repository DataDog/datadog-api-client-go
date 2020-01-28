# SyntheticsSslCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cipher** | Pointer to **string** |  | [optional] 
**Exponent** | Pointer to **float64** |  | [optional] 
**ExtKeyUsage** | Pointer to **[]string** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Fingerprint256** | Pointer to **string** |  | [optional] 
**Issuer** | Pointer to [**SyntheticsSslCertificateIssuer**](SyntheticsSSLCertificate_issuer.md) |  | [optional] 
**Modulus** | Pointer to **string** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to [**SyntheticsSslCertificateSubject**](SyntheticsSSLCertificate_subject.md) |  | [optional] 
**ValidFrom** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ValidTo** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 

## Methods

### GetCipher

`func (o *SyntheticsSslCertificate) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *SyntheticsSslCertificate) GetCipherOk() (string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCipher

`func (o *SyntheticsSslCertificate) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### SetCipher

`func (o *SyntheticsSslCertificate) SetCipher(v string)`

SetCipher gets a reference to the given string and assigns it to the Cipher field.

### GetExponent

`func (o *SyntheticsSslCertificate) GetExponent() float64`

GetExponent returns the Exponent field if non-nil, zero value otherwise.

### GetExponentOk

`func (o *SyntheticsSslCertificate) GetExponentOk() (float64, bool)`

GetExponentOk returns a tuple with the Exponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExponent

`func (o *SyntheticsSslCertificate) HasExponent() bool`

HasExponent returns a boolean if a field has been set.

### SetExponent

`func (o *SyntheticsSslCertificate) SetExponent(v float64)`

SetExponent gets a reference to the given float64 and assigns it to the Exponent field.

### GetExtKeyUsage

`func (o *SyntheticsSslCertificate) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *SyntheticsSslCertificate) GetExtKeyUsageOk() ([]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExtKeyUsage

`func (o *SyntheticsSslCertificate) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.

### SetExtKeyUsage

`func (o *SyntheticsSslCertificate) SetExtKeyUsage(v []string)`

SetExtKeyUsage gets a reference to the given []string and assigns it to the ExtKeyUsage field.

### GetFingerprint

`func (o *SyntheticsSslCertificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SyntheticsSslCertificate) GetFingerprintOk() (string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprint

`func (o *SyntheticsSslCertificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprint

`func (o *SyntheticsSslCertificate) SetFingerprint(v string)`

SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.

### GetFingerprint256

`func (o *SyntheticsSslCertificate) GetFingerprint256() string`

GetFingerprint256 returns the Fingerprint256 field if non-nil, zero value otherwise.

### GetFingerprint256Ok

`func (o *SyntheticsSslCertificate) GetFingerprint256Ok() (string, bool)`

GetFingerprint256Ok returns a tuple with the Fingerprint256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFingerprint256

`func (o *SyntheticsSslCertificate) HasFingerprint256() bool`

HasFingerprint256 returns a boolean if a field has been set.

### SetFingerprint256

`func (o *SyntheticsSslCertificate) SetFingerprint256(v string)`

SetFingerprint256 gets a reference to the given string and assigns it to the Fingerprint256 field.

### GetIssuer

`func (o *SyntheticsSslCertificate) GetIssuer() SyntheticsSslCertificateIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SyntheticsSslCertificate) GetIssuerOk() (SyntheticsSslCertificateIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIssuer

`func (o *SyntheticsSslCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuer

`func (o *SyntheticsSslCertificate) SetIssuer(v SyntheticsSslCertificateIssuer)`

SetIssuer gets a reference to the given SyntheticsSslCertificateIssuer and assigns it to the Issuer field.

### GetModulus

`func (o *SyntheticsSslCertificate) GetModulus() string`

GetModulus returns the Modulus field if non-nil, zero value otherwise.

### GetModulusOk

`func (o *SyntheticsSslCertificate) GetModulusOk() (string, bool)`

GetModulusOk returns a tuple with the Modulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModulus

`func (o *SyntheticsSslCertificate) HasModulus() bool`

HasModulus returns a boolean if a field has been set.

### SetModulus

`func (o *SyntheticsSslCertificate) SetModulus(v string)`

SetModulus gets a reference to the given string and assigns it to the Modulus field.

### GetProtocol

`func (o *SyntheticsSslCertificate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyntheticsSslCertificate) GetProtocolOk() (string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProtocol

`func (o *SyntheticsSslCertificate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocol

`func (o *SyntheticsSslCertificate) SetProtocol(v string)`

SetProtocol gets a reference to the given string and assigns it to the Protocol field.

### GetSerialNumber

`func (o *SyntheticsSslCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SyntheticsSslCertificate) GetSerialNumberOk() (string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSerialNumber

`func (o *SyntheticsSslCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumber

`func (o *SyntheticsSslCertificate) SetSerialNumber(v string)`

SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.

### GetSubject

`func (o *SyntheticsSslCertificate) GetSubject() SyntheticsSslCertificateSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SyntheticsSslCertificate) GetSubjectOk() (SyntheticsSslCertificateSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubject

`func (o *SyntheticsSslCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubject

`func (o *SyntheticsSslCertificate) SetSubject(v SyntheticsSslCertificateSubject)`

SetSubject gets a reference to the given SyntheticsSslCertificateSubject and assigns it to the Subject field.

### GetValidFrom

`func (o *SyntheticsSslCertificate) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *SyntheticsSslCertificate) GetValidFromOk() (time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidFrom

`func (o *SyntheticsSslCertificate) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### SetValidFrom

`func (o *SyntheticsSslCertificate) SetValidFrom(v time.Time)`

SetValidFrom gets a reference to the given time.Time and assigns it to the ValidFrom field.

### GetValidTo

`func (o *SyntheticsSslCertificate) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *SyntheticsSslCertificate) GetValidToOk() (time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidTo

`func (o *SyntheticsSslCertificate) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### SetValidTo

`func (o *SyntheticsSslCertificate) SetValidTo(v time.Time)`

SetValidTo gets a reference to the given time.Time and assigns it to the ValidTo field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


