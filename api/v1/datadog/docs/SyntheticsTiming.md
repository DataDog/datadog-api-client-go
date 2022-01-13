# SyntheticsTiming

## Properties

| Name          | Type                   | Description                                                       | Notes      |
| ------------- | ---------------------- | ----------------------------------------------------------------- | ---------- |
| **Dns**       | Pointer to **float64** | The duration in millisecond of the DNS lookup.                    | [optional] |
| **Download**  | Pointer to **float64** | The time in millisecond to download the response.                 | [optional] |
| **FirstByte** | Pointer to **float64** | The time in millisecond to first byte.                            | [optional] |
| **Handshake** | Pointer to **float64** | The duration in millisecond of the TLS handshake.                 | [optional] |
| **Redirect**  | Pointer to **float64** | The time in millisecond spent during redirections.                | [optional] |
| **Ssl**       | Pointer to **float64** | The duration in millisecond of the TLS handshake.                 | [optional] |
| **Tcp**       | Pointer to **float64** | Time in millisecond to establish the TCP connection.              | [optional] |
| **Total**     | Pointer to **float64** | The overall time in millisecond the request took to be processed. | [optional] |
| **Wait**      | Pointer to **float64** | Time spent in millisecond waiting for a response.                 | [optional] |

## Methods

### NewSyntheticsTiming

`func NewSyntheticsTiming() *SyntheticsTiming`

NewSyntheticsTiming instantiates a new SyntheticsTiming object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSyntheticsTimingWithDefaults

`func NewSyntheticsTimingWithDefaults() *SyntheticsTiming`

NewSyntheticsTimingWithDefaults instantiates a new SyntheticsTiming object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDns

`func (o *SyntheticsTiming) GetDns() float64`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *SyntheticsTiming) GetDnsOk() (*float64, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *SyntheticsTiming) SetDns(v float64)`

SetDns sets Dns field to given value.

### HasDns

`func (o *SyntheticsTiming) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetDownload

`func (o *SyntheticsTiming) GetDownload() float64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *SyntheticsTiming) GetDownloadOk() (*float64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *SyntheticsTiming) SetDownload(v float64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *SyntheticsTiming) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetFirstByte

`func (o *SyntheticsTiming) GetFirstByte() float64`

GetFirstByte returns the FirstByte field if non-nil, zero value otherwise.

### GetFirstByteOk

`func (o *SyntheticsTiming) GetFirstByteOk() (*float64, bool)`

GetFirstByteOk returns a tuple with the FirstByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstByte

`func (o *SyntheticsTiming) SetFirstByte(v float64)`

SetFirstByte sets FirstByte field to given value.

### HasFirstByte

`func (o *SyntheticsTiming) HasFirstByte() bool`

HasFirstByte returns a boolean if a field has been set.

### GetHandshake

`func (o *SyntheticsTiming) GetHandshake() float64`

GetHandshake returns the Handshake field if non-nil, zero value otherwise.

### GetHandshakeOk

`func (o *SyntheticsTiming) GetHandshakeOk() (*float64, bool)`

GetHandshakeOk returns a tuple with the Handshake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandshake

`func (o *SyntheticsTiming) SetHandshake(v float64)`

SetHandshake sets Handshake field to given value.

### HasHandshake

`func (o *SyntheticsTiming) HasHandshake() bool`

HasHandshake returns a boolean if a field has been set.

### GetRedirect

`func (o *SyntheticsTiming) GetRedirect() float64`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *SyntheticsTiming) GetRedirectOk() (*float64, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *SyntheticsTiming) SetRedirect(v float64)`

SetRedirect sets Redirect field to given value.

### HasRedirect

`func (o *SyntheticsTiming) HasRedirect() bool`

HasRedirect returns a boolean if a field has been set.

### GetSsl

`func (o *SyntheticsTiming) GetSsl() float64`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *SyntheticsTiming) GetSslOk() (*float64, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *SyntheticsTiming) SetSsl(v float64)`

SetSsl sets Ssl field to given value.

### HasSsl

`func (o *SyntheticsTiming) HasSsl() bool`

HasSsl returns a boolean if a field has been set.

### GetTcp

`func (o *SyntheticsTiming) GetTcp() float64`

GetTcp returns the Tcp field if non-nil, zero value otherwise.

### GetTcpOk

`func (o *SyntheticsTiming) GetTcpOk() (*float64, bool)`

GetTcpOk returns a tuple with the Tcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcp

`func (o *SyntheticsTiming) SetTcp(v float64)`

SetTcp sets Tcp field to given value.

### HasTcp

`func (o *SyntheticsTiming) HasTcp() bool`

HasTcp returns a boolean if a field has been set.

### GetTotal

`func (o *SyntheticsTiming) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SyntheticsTiming) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SyntheticsTiming) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SyntheticsTiming) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWait

`func (o *SyntheticsTiming) GetWait() float64`

GetWait returns the Wait field if non-nil, zero value otherwise.

### GetWaitOk

`func (o *SyntheticsTiming) GetWaitOk() (*float64, bool)`

GetWaitOk returns a tuple with the Wait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWait

`func (o *SyntheticsTiming) SetWait(v float64)`

SetWait sets Wait field to given value.

### HasWait

`func (o *SyntheticsTiming) HasWait() bool`

HasWait returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
