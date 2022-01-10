# HostMeta

## Properties

Name | Type | Description | Notes
---- | ---- | ----------- | ------
**AgentChecks** | Pointer to **[][]interface{}** | A list of Agent checks running on the host. | [optional] 
**AgentVersion** | Pointer to **string** | The Datadog Agent version. | [optional] 
**CpuCores** | Pointer to **int64** | The number of cores. | [optional] 
**FbsdV** | Pointer to **[]string** | An array of Mac versions. | [optional] 
**Gohai** | Pointer to **string** | JSON string containing system information. | [optional] 
**InstallMethod** | Pointer to [**HostMetaInstallMethod**](HostMetaInstallMethod.md) |  | [optional] 
**MacV** | Pointer to **[]string** | An array of Mac versions. | [optional] 
**Machine** | Pointer to **string** | The machine architecture. | [optional] 
**NixV** | Pointer to **[]string** | Array of Unix versions. | [optional] 
**Platform** | Pointer to **string** | The OS platform. | [optional] 
**Processor** | Pointer to **string** | The processor. | [optional] 
**PythonV** | Pointer to **string** | The Python version. | [optional] 
**SocketFqdn** | Pointer to **string** | The socket fqdn. | [optional] 
**SocketHostname** | Pointer to **string** | The socket hostname. | [optional] 
**WinV** | Pointer to **[]string** | An array of Windows versions. | [optional] 

## Methods

### NewHostMeta

`func NewHostMeta() *HostMeta`

NewHostMeta instantiates a new HostMeta object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewHostMetaWithDefaults

`func NewHostMetaWithDefaults() *HostMeta`

NewHostMetaWithDefaults instantiates a new HostMeta object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAgentChecks

`func (o *HostMeta) GetAgentChecks() [][]interface{}`

GetAgentChecks returns the AgentChecks field if non-nil, zero value otherwise.

### GetAgentChecksOk

`func (o *HostMeta) GetAgentChecksOk() (*[][]interface{}, bool)`

GetAgentChecksOk returns a tuple with the AgentChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentChecks

`func (o *HostMeta) SetAgentChecks(v [][]interface{})`

SetAgentChecks sets AgentChecks field to given value.

### HasAgentChecks

`func (o *HostMeta) HasAgentChecks() bool`

HasAgentChecks returns a boolean if a field has been set.

### GetAgentVersion

`func (o *HostMeta) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *HostMeta) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *HostMeta) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *HostMeta) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetCpuCores

`func (o *HostMeta) GetCpuCores() int64`

GetCpuCores returns the CpuCores field if non-nil, zero value otherwise.

### GetCpuCoresOk

`func (o *HostMeta) GetCpuCoresOk() (*int64, bool)`

GetCpuCoresOk returns a tuple with the CpuCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCores

`func (o *HostMeta) SetCpuCores(v int64)`

SetCpuCores sets CpuCores field to given value.

### HasCpuCores

`func (o *HostMeta) HasCpuCores() bool`

HasCpuCores returns a boolean if a field has been set.

### GetFbsdV

`func (o *HostMeta) GetFbsdV() []string`

GetFbsdV returns the FbsdV field if non-nil, zero value otherwise.

### GetFbsdVOk

`func (o *HostMeta) GetFbsdVOk() (*[]string, bool)`

GetFbsdVOk returns a tuple with the FbsdV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbsdV

`func (o *HostMeta) SetFbsdV(v []string)`

SetFbsdV sets FbsdV field to given value.

### HasFbsdV

`func (o *HostMeta) HasFbsdV() bool`

HasFbsdV returns a boolean if a field has been set.

### GetGohai

`func (o *HostMeta) GetGohai() string`

GetGohai returns the Gohai field if non-nil, zero value otherwise.

### GetGohaiOk

`func (o *HostMeta) GetGohaiOk() (*string, bool)`

GetGohaiOk returns a tuple with the Gohai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGohai

`func (o *HostMeta) SetGohai(v string)`

SetGohai sets Gohai field to given value.

### HasGohai

`func (o *HostMeta) HasGohai() bool`

HasGohai returns a boolean if a field has been set.

### GetInstallMethod

`func (o *HostMeta) GetInstallMethod() HostMetaInstallMethod`

GetInstallMethod returns the InstallMethod field if non-nil, zero value otherwise.

### GetInstallMethodOk

`func (o *HostMeta) GetInstallMethodOk() (*HostMetaInstallMethod, bool)`

GetInstallMethodOk returns a tuple with the InstallMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallMethod

`func (o *HostMeta) SetInstallMethod(v HostMetaInstallMethod)`

SetInstallMethod sets InstallMethod field to given value.

### HasInstallMethod

`func (o *HostMeta) HasInstallMethod() bool`

HasInstallMethod returns a boolean if a field has been set.

### GetMacV

`func (o *HostMeta) GetMacV() []string`

GetMacV returns the MacV field if non-nil, zero value otherwise.

### GetMacVOk

`func (o *HostMeta) GetMacVOk() (*[]string, bool)`

GetMacVOk returns a tuple with the MacV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacV

`func (o *HostMeta) SetMacV(v []string)`

SetMacV sets MacV field to given value.

### HasMacV

`func (o *HostMeta) HasMacV() bool`

HasMacV returns a boolean if a field has been set.

### GetMachine

`func (o *HostMeta) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *HostMeta) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *HostMeta) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *HostMeta) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetNixV

`func (o *HostMeta) GetNixV() []string`

GetNixV returns the NixV field if non-nil, zero value otherwise.

### GetNixVOk

`func (o *HostMeta) GetNixVOk() (*[]string, bool)`

GetNixVOk returns a tuple with the NixV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNixV

`func (o *HostMeta) SetNixV(v []string)`

SetNixV sets NixV field to given value.

### HasNixV

`func (o *HostMeta) HasNixV() bool`

HasNixV returns a boolean if a field has been set.

### GetPlatform

`func (o *HostMeta) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *HostMeta) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *HostMeta) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *HostMeta) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProcessor

`func (o *HostMeta) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *HostMeta) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *HostMeta) SetProcessor(v string)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *HostMeta) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetPythonV

`func (o *HostMeta) GetPythonV() string`

GetPythonV returns the PythonV field if non-nil, zero value otherwise.

### GetPythonVOk

`func (o *HostMeta) GetPythonVOk() (*string, bool)`

GetPythonVOk returns a tuple with the PythonV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonV

`func (o *HostMeta) SetPythonV(v string)`

SetPythonV sets PythonV field to given value.

### HasPythonV

`func (o *HostMeta) HasPythonV() bool`

HasPythonV returns a boolean if a field has been set.

### GetSocketFqdn

`func (o *HostMeta) GetSocketFqdn() string`

GetSocketFqdn returns the SocketFqdn field if non-nil, zero value otherwise.

### GetSocketFqdnOk

`func (o *HostMeta) GetSocketFqdnOk() (*string, bool)`

GetSocketFqdnOk returns a tuple with the SocketFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketFqdn

`func (o *HostMeta) SetSocketFqdn(v string)`

SetSocketFqdn sets SocketFqdn field to given value.

### HasSocketFqdn

`func (o *HostMeta) HasSocketFqdn() bool`

HasSocketFqdn returns a boolean if a field has been set.

### GetSocketHostname

`func (o *HostMeta) GetSocketHostname() string`

GetSocketHostname returns the SocketHostname field if non-nil, zero value otherwise.

### GetSocketHostnameOk

`func (o *HostMeta) GetSocketHostnameOk() (*string, bool)`

GetSocketHostnameOk returns a tuple with the SocketHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocketHostname

`func (o *HostMeta) SetSocketHostname(v string)`

SetSocketHostname sets SocketHostname field to given value.

### HasSocketHostname

`func (o *HostMeta) HasSocketHostname() bool`

HasSocketHostname returns a boolean if a field has been set.

### GetWinV

`func (o *HostMeta) GetWinV() []string`

GetWinV returns the WinV field if non-nil, zero value otherwise.

### GetWinVOk

`func (o *HostMeta) GetWinVOk() (*[]string, bool)`

GetWinVOk returns a tuple with the WinV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinV

`func (o *HostMeta) SetWinV(v []string)`

SetWinV sets WinV field to given value.

### HasWinV

`func (o *HostMeta) HasWinV() bool`

HasWinV returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


