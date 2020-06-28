# SyntheticsSslCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cipher** | **string** | Cipher used for the connection. | [optional] 
**Exponent** | **float64** | Exponent associated to the certificate. | [optional] 
**ExtKeyUsage** | **[]string** | Array of extensions and details used for the certificate. | [optional] 
**Fingerprint** | **string** | MD5 digest of the DER-encoded Certificate information. | [optional] 
**Fingerprint256** | **string** | SHA-1 digest of the DER-encoded Certificate information. | [optional] 
**Issuer** | [**SyntheticsSslCertificateIssuer**](SyntheticsSSLCertificate_issuer.md) |  | [optional] 
**Modulus** | **string** | Modulus associated to the SSL certificate private key. | [optional] 
**Protocol** | **string** | TLS protocol used for the test. | [optional] 
**SerialNumber** | **string** | Serial Number assigned by Symantec to the SSL certificate. | [optional] 
**Subject** | [**SyntheticsSslCertificateSubject**](SyntheticsSSLCertificate_subject.md) |  | [optional] 
**ValidFrom** | [**time.Time**](time.Time.md) | Date from which the SSL certificate is valid. | [optional] 
**ValidTo** | [**time.Time**](time.Time.md) | Date until which the SSL certificate is valid. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


