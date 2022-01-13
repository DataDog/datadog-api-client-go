# SyntheticsTestRequest

## Properties

| Name                     | Type                                                                                   | Description                                                                                                                                                                                          | Notes      |
| ------------------------ | -------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| **AllowInsecure**        | Pointer to **bool**                                                                    | Allows loading insecure content for an HTTP request in a multistep test step.                                                                                                                        | [optional] |
| **BasicAuth**            | Pointer to [**SyntheticsBasicAuth**](SyntheticsBasicAuth.md)                           |                                                                                                                                                                                                      | [optional] |
| **Body**                 | Pointer to **string**                                                                  | Body to include in the test.                                                                                                                                                                         | [optional] |
| **Certificate**          | Pointer to [**SyntheticsTestRequestCertificate**](SyntheticsTestRequestCertificate.md) |                                                                                                                                                                                                      | [optional] |
| **DnsServer**            | Pointer to **string**                                                                  | DNS server to use for DNS tests.                                                                                                                                                                     | [optional] |
| **DnsServerPort**        | Pointer to **int32**                                                                   | DNS server port to use for DNS tests.                                                                                                                                                                | [optional] |
| **FollowRedirects**      | Pointer to **bool**                                                                    | Specifies whether or not the request follows redirects.                                                                                                                                              | [optional] |
| **Headers**              | Pointer to **map[string]string**                                                       | Headers to include when performing the test.                                                                                                                                                         | [optional] |
| **Host**                 | Pointer to **string**                                                                  | Host name to perform the test with.                                                                                                                                                                  | [optional] |
| **Message**              | Pointer to **string**                                                                  | Message to send for UDP or WebSocket tests.                                                                                                                                                          | [optional] |
| **Method**               | Pointer to [**HTTPMethod**](HTTPMethod.md)                                             |                                                                                                                                                                                                      | [optional] |
| **NoSavingResponseBody** | Pointer to **bool**                                                                    | Determines whether or not to save the response body.                                                                                                                                                 | [optional] |
| **NumberOfPackets**      | Pointer to **int32**                                                                   | Number of pings to use per test.                                                                                                                                                                     | [optional] |
| **Port**                 | Pointer to **int64**                                                                   | Port to use when performing the test.                                                                                                                                                                | [optional] |
| **Proxy**                | Pointer to [**SyntheticsTestRequestProxy**](SyntheticsTestRequestProxy.md)             |                                                                                                                                                                                                      | [optional] |
| **Query**                | Pointer to **interface{}**                                                             | Query to use for the test.                                                                                                                                                                           | [optional] |
| **Servername**           | Pointer to **string**                                                                  | For SSL tests, it specifies on which server you want to initiate the TLS handshake, allowing the server to present one of multiple possible certificates on the same IP address and TCP port number. | [optional] |
| **ShouldTrackHops**      | Pointer to **bool**                                                                    | Turns on a traceroute probe to discover all gateways along the path to the host destination.                                                                                                         | [optional] |
| **Timeout**              | Pointer to **float64**                                                                 | Timeout in seconds for the test.                                                                                                                                                                     | [optional] |
| **Url**                  | Pointer to **string**                                                                  | URL to perform the test with.                                                                                                                                                                        | [optional] |

## Methods

### NewSyntheticsTestRequest

`func NewSyntheticsTestRequest() *SyntheticsTestRequest`

NewSyntheticsTestRequest instantiates a new SyntheticsTestRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsTestRequestWithDefaults

`func NewSyntheticsTestRequestWithDefaults() *SyntheticsTestRequest`

NewSyntheticsTestRequestWithDefaults instantiates a new SyntheticsTestRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAllowInsecure

`func (o *SyntheticsTestRequest) GetAllowInsecure() bool`

GetAllowInsecure returns the AllowInsecure field if non-nil, zero value otherwise.

### GetAllowInsecureOk

`func (o *SyntheticsTestRequest) GetAllowInsecureOk() (*bool, bool)`

GetAllowInsecureOk returns a tuple with the AllowInsecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInsecure

`func (o *SyntheticsTestRequest) SetAllowInsecure(v bool)`

SetAllowInsecure sets AllowInsecure field to given value.

### HasAllowInsecure

`func (o *SyntheticsTestRequest) HasAllowInsecure() bool`

HasAllowInsecure returns a boolean if a field has been set.

### GetBasicAuth

`func (o *SyntheticsTestRequest) GetBasicAuth() SyntheticsBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *SyntheticsTestRequest) GetBasicAuthOk() (*SyntheticsBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *SyntheticsTestRequest) SetBasicAuth(v SyntheticsBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *SyntheticsTestRequest) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBody

`func (o *SyntheticsTestRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SyntheticsTestRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SyntheticsTestRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SyntheticsTestRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCertificate

`func (o *SyntheticsTestRequest) GetCertificate() SyntheticsTestRequestCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SyntheticsTestRequest) GetCertificateOk() (*SyntheticsTestRequestCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SyntheticsTestRequest) SetCertificate(v SyntheticsTestRequestCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *SyntheticsTestRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDnsServer

`func (o *SyntheticsTestRequest) GetDnsServer() string`

GetDnsServer returns the DnsServer field if non-nil, zero value otherwise.

### GetDnsServerOk

`func (o *SyntheticsTestRequest) GetDnsServerOk() (*string, bool)`

GetDnsServerOk returns a tuple with the DnsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServer

`func (o *SyntheticsTestRequest) SetDnsServer(v string)`

SetDnsServer sets DnsServer field to given value.

### HasDnsServer

`func (o *SyntheticsTestRequest) HasDnsServer() bool`

HasDnsServer returns a boolean if a field has been set.

### GetDnsServerPort

`func (o *SyntheticsTestRequest) GetDnsServerPort() int32`

GetDnsServerPort returns the DnsServerPort field if non-nil, zero value otherwise.

### GetDnsServerPortOk

`func (o *SyntheticsTestRequest) GetDnsServerPortOk() (*int32, bool)`

GetDnsServerPortOk returns a tuple with the DnsServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServerPort

`func (o *SyntheticsTestRequest) SetDnsServerPort(v int32)`

SetDnsServerPort sets DnsServerPort field to given value.

### HasDnsServerPort

`func (o *SyntheticsTestRequest) HasDnsServerPort() bool`

HasDnsServerPort returns a boolean if a field has been set.

### GetFollowRedirects

`func (o *SyntheticsTestRequest) GetFollowRedirects() bool`

GetFollowRedirects returns the FollowRedirects field if non-nil, zero value otherwise.

### GetFollowRedirectsOk

`func (o *SyntheticsTestRequest) GetFollowRedirectsOk() (*bool, bool)`

GetFollowRedirectsOk returns a tuple with the FollowRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowRedirects

`func (o *SyntheticsTestRequest) SetFollowRedirects(v bool)`

SetFollowRedirects sets FollowRedirects field to given value.

### HasFollowRedirects

`func (o *SyntheticsTestRequest) HasFollowRedirects() bool`

HasFollowRedirects returns a boolean if a field has been set.

### GetHeaders

`func (o *SyntheticsTestRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *SyntheticsTestRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *SyntheticsTestRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *SyntheticsTestRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetHost

`func (o *SyntheticsTestRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SyntheticsTestRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SyntheticsTestRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SyntheticsTestRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMessage

`func (o *SyntheticsTestRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SyntheticsTestRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SyntheticsTestRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SyntheticsTestRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMethod

`func (o *SyntheticsTestRequest) GetMethod() HTTPMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *SyntheticsTestRequest) GetMethodOk() (*HTTPMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *SyntheticsTestRequest) SetMethod(v HTTPMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *SyntheticsTestRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNoSavingResponseBody

`func (o *SyntheticsTestRequest) GetNoSavingResponseBody() bool`

GetNoSavingResponseBody returns the NoSavingResponseBody field if non-nil, zero value otherwise.

### GetNoSavingResponseBodyOk

`func (o *SyntheticsTestRequest) GetNoSavingResponseBodyOk() (*bool, bool)`

GetNoSavingResponseBodyOk returns a tuple with the NoSavingResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoSavingResponseBody

`func (o *SyntheticsTestRequest) SetNoSavingResponseBody(v bool)`

SetNoSavingResponseBody sets NoSavingResponseBody field to given value.

### HasNoSavingResponseBody

`func (o *SyntheticsTestRequest) HasNoSavingResponseBody() bool`

HasNoSavingResponseBody returns a boolean if a field has been set.

### GetNumberOfPackets

`func (o *SyntheticsTestRequest) GetNumberOfPackets() int32`

GetNumberOfPackets returns the NumberOfPackets field if non-nil, zero value otherwise.

### GetNumberOfPacketsOk

`func (o *SyntheticsTestRequest) GetNumberOfPacketsOk() (*int32, bool)`

GetNumberOfPacketsOk returns a tuple with the NumberOfPackets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPackets

`func (o *SyntheticsTestRequest) SetNumberOfPackets(v int32)`

SetNumberOfPackets sets NumberOfPackets field to given value.

### HasNumberOfPackets

`func (o *SyntheticsTestRequest) HasNumberOfPackets() bool`

HasNumberOfPackets returns a boolean if a field has been set.

### GetPort

`func (o *SyntheticsTestRequest) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SyntheticsTestRequest) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SyntheticsTestRequest) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *SyntheticsTestRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProxy

`func (o *SyntheticsTestRequest) GetProxy() SyntheticsTestRequestProxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *SyntheticsTestRequest) GetProxyOk() (*SyntheticsTestRequestProxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *SyntheticsTestRequest) SetProxy(v SyntheticsTestRequestProxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *SyntheticsTestRequest) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetQuery

`func (o *SyntheticsTestRequest) GetQuery() interface{}`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SyntheticsTestRequest) GetQueryOk() (*interface{}, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SyntheticsTestRequest) SetQuery(v interface{})`

SetQuery sets Query field to given value.

### HasQuery

`func (o *SyntheticsTestRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetServername

`func (o *SyntheticsTestRequest) GetServername() string`

GetServername returns the Servername field if non-nil, zero value otherwise.

### GetServernameOk

`func (o *SyntheticsTestRequest) GetServernameOk() (*string, bool)`

GetServernameOk returns a tuple with the Servername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServername

`func (o *SyntheticsTestRequest) SetServername(v string)`

SetServername sets Servername field to given value.

### HasServername

`func (o *SyntheticsTestRequest) HasServername() bool`

HasServername returns a boolean if a field has been set.

### GetShouldTrackHops

`func (o *SyntheticsTestRequest) GetShouldTrackHops() bool`

GetShouldTrackHops returns the ShouldTrackHops field if non-nil, zero value otherwise.

### GetShouldTrackHopsOk

`func (o *SyntheticsTestRequest) GetShouldTrackHopsOk() (*bool, bool)`

GetShouldTrackHopsOk returns a tuple with the ShouldTrackHops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldTrackHops

`func (o *SyntheticsTestRequest) SetShouldTrackHops(v bool)`

SetShouldTrackHops sets ShouldTrackHops field to given value.

### HasShouldTrackHops

`func (o *SyntheticsTestRequest) HasShouldTrackHops() bool`

HasShouldTrackHops returns a boolean if a field has been set.

### GetTimeout

`func (o *SyntheticsTestRequest) GetTimeout() float64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *SyntheticsTestRequest) GetTimeoutOk() (*float64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *SyntheticsTestRequest) SetTimeout(v float64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *SyntheticsTestRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUrl

`func (o *SyntheticsTestRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SyntheticsTestRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SyntheticsTestRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SyntheticsTestRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
