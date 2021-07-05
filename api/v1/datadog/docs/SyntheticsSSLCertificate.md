# SyntheticsSSLCertificate

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**Cipher** | Pointer to **string** | Cipher used for the connection. | [optional] 
**Exponent** | Pointer to **float64** | Exponent associated to the certificate. | [optional] 
**ExtKeyUsage** | Pointer to **[]string** | Array of extensions and details used for the certificate. | [optional] 
**Fingerprint** | Pointer to **string** | MD5 digest of the DER-encoded Certificate information. | [optional] 
**Fingerprint256** | Pointer to **string** | SHA-1 digest of the DER-encoded Certificate information. | [optional] 
**Issuer** | Pointer to [**SyntheticsSSLCertificateIssuer**](SyntheticsSSLCertificateIssuer.md) |  | [optional] 
**Modulus** | Pointer to **string** | Modulus associated to the SSL certificate private key. | [optional] 
**Protocol** | Pointer to **string** | TLS protocol used for the test. | [optional] 
**SerialNumber** | Pointer to **string** | Serial Number assigned by Symantec to the SSL certificate. | [optional] 
**Subject** | Pointer to [**SyntheticsSSLCertificateSubject**](SyntheticsSSLCertificateSubject.md) |  | [optional] 
**ValidFrom** | Pointer to **time.Time** | Date from which the SSL certificate is valid. | [optional] 
**ValidTo** | Pointer to **time.Time** | Date until which the SSL certificate is valid. | [optional] 

## Methods

### NewSyntheticsSSLCertificate

`func NewSyntheticsSSLCertificate() *SyntheticsSSLCertificate`

NewSyntheticsSSLCertificate instantiates a new SyntheticsSSLCertificate object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsSSLCertificateWithDefaults

`func NewSyntheticsSSLCertificateWithDefaults() *SyntheticsSSLCertificate`

NewSyntheticsSSLCertificateWithDefaults instantiates a new SyntheticsSSLCertificate object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCipher

`func (o *SyntheticsSSLCertificate) GetCipher() string`

GetCipher returns the Cipher field if non-nil, zero value otherwise.

### GetCipherOk

`func (o *SyntheticsSSLCertificate) GetCipherOk() (*string, bool)`

GetCipherOk returns a tuple with the Cipher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipher

`func (o *SyntheticsSSLCertificate) SetCipher(v string)`

SetCipher sets Cipher field to given value.

### HasCipher

`func (o *SyntheticsSSLCertificate) HasCipher() bool`

HasCipher returns a boolean if a field has been set.

### GetExponent

`func (o *SyntheticsSSLCertificate) GetExponent() float64`

GetExponent returns the Exponent field if non-nil, zero value otherwise.

### GetExponentOk

`func (o *SyntheticsSSLCertificate) GetExponentOk() (*float64, bool)`

GetExponentOk returns a tuple with the Exponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExponent

`func (o *SyntheticsSSLCertificate) SetExponent(v float64)`

SetExponent sets Exponent field to given value.

### HasExponent

`func (o *SyntheticsSSLCertificate) HasExponent() bool`

HasExponent returns a boolean if a field has been set.

### GetExtKeyUsage

`func (o *SyntheticsSSLCertificate) GetExtKeyUsage() []string`

GetExtKeyUsage returns the ExtKeyUsage field if non-nil, zero value otherwise.

### GetExtKeyUsageOk

`func (o *SyntheticsSSLCertificate) GetExtKeyUsageOk() (*[]string, bool)`

GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtKeyUsage

`func (o *SyntheticsSSLCertificate) SetExtKeyUsage(v []string)`

SetExtKeyUsage sets ExtKeyUsage field to given value.

### HasExtKeyUsage

`func (o *SyntheticsSSLCertificate) HasExtKeyUsage() bool`

HasExtKeyUsage returns a boolean if a field has been set.

### GetFingerprint

`func (o *SyntheticsSSLCertificate) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *SyntheticsSSLCertificate) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *SyntheticsSSLCertificate) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *SyntheticsSSLCertificate) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetFingerprint256

`func (o *SyntheticsSSLCertificate) GetFingerprint256() string`

GetFingerprint256 returns the Fingerprint256 field if non-nil, zero value otherwise.

### GetFingerprint256Ok

`func (o *SyntheticsSSLCertificate) GetFingerprint256Ok() (*string, bool)`

GetFingerprint256Ok returns a tuple with the Fingerprint256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint256

`func (o *SyntheticsSSLCertificate) SetFingerprint256(v string)`

SetFingerprint256 sets Fingerprint256 field to given value.

### HasFingerprint256

`func (o *SyntheticsSSLCertificate) HasFingerprint256() bool`

HasFingerprint256 returns a boolean if a field has been set.

### GetIssuer

`func (o *SyntheticsSSLCertificate) GetIssuer() SyntheticsSSLCertificateIssuer`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SyntheticsSSLCertificate) GetIssuerOk() (*SyntheticsSSLCertificateIssuer, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SyntheticsSSLCertificate) SetIssuer(v SyntheticsSSLCertificateIssuer)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *SyntheticsSSLCertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetModulus

`func (o *SyntheticsSSLCertificate) GetModulus() string`

GetModulus returns the Modulus field if non-nil, zero value otherwise.

### GetModulusOk

`func (o *SyntheticsSSLCertificate) GetModulusOk() (*string, bool)`

GetModulusOk returns a tuple with the Modulus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulus

`func (o *SyntheticsSSLCertificate) SetModulus(v string)`

SetModulus sets Modulus field to given value.

### HasModulus

`func (o *SyntheticsSSLCertificate) HasModulus() bool`

HasModulus returns a boolean if a field has been set.

### GetProtocol

`func (o *SyntheticsSSLCertificate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SyntheticsSSLCertificate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SyntheticsSSLCertificate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SyntheticsSSLCertificate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSerialNumber

`func (o *SyntheticsSSLCertificate) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *SyntheticsSSLCertificate) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *SyntheticsSSLCertificate) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *SyntheticsSSLCertificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSubject

`func (o *SyntheticsSSLCertificate) GetSubject() SyntheticsSSLCertificateSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *SyntheticsSSLCertificate) GetSubjectOk() (*SyntheticsSSLCertificateSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *SyntheticsSSLCertificate) SetSubject(v SyntheticsSSLCertificateSubject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *SyntheticsSSLCertificate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetValidFrom

`func (o *SyntheticsSSLCertificate) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *SyntheticsSSLCertificate) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *SyntheticsSSLCertificate) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *SyntheticsSSLCertificate) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *SyntheticsSSLCertificate) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *SyntheticsSSLCertificate) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *SyntheticsSSLCertificate) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *SyntheticsSSLCertificate) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


