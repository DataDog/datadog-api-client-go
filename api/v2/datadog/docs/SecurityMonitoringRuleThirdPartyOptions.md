# SecurityMonitoringRuleThirdPartyOptions

## Properties

| Name                     | Type                                                                               | Description                                                                                      | Notes      |
| ------------------------ | ---------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------ | ---------- |
| **DefaultNotifications** | Pointer to **[]string**                                                            | Notification targets for the root query.                                                         | [optional] |
| **DefaultStatus**        | Pointer to [**SecurityMonitoringRuleSeverity**](SecurityMonitoringRuleSeverity.md) |                                                                                                  | [optional] |
| **FirstSeenOverride**    | Pointer to **string**                                                              | (Optional): the name of an attribute to override the first seen value of the third party signal. | [optional] |
| **LastSeenOverride**     | Pointer to **string**                                                              | (Optional): the name of an attribute to override the last seen value of the third party signal.  | [optional] |
| **RootQuery**            | Pointer to **string**                                                              | Root query of the rule.                                                                          | [optional] |
| **SignalId**             | Pointer to **string**                                                              | Optional mapping of the third-party signal ID.                                                   | [optional] |

## Methods

### NewSecurityMonitoringRuleThirdPartyOptions

`func NewSecurityMonitoringRuleThirdPartyOptions() *SecurityMonitoringRuleThirdPartyOptions`

NewSecurityMonitoringRuleThirdPartyOptions instantiates a new SecurityMonitoringRuleThirdPartyOptions object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSecurityMonitoringRuleThirdPartyOptionsWithDefaults

`func NewSecurityMonitoringRuleThirdPartyOptionsWithDefaults() *SecurityMonitoringRuleThirdPartyOptions`

NewSecurityMonitoringRuleThirdPartyOptionsWithDefaults instantiates a new SecurityMonitoringRuleThirdPartyOptions object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDefaultNotifications

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultNotifications() []string`

GetDefaultNotifications returns the DefaultNotifications field if non-nil, zero value otherwise.

### GetDefaultNotificationsOk

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultNotificationsOk() (*[]string, bool)`

GetDefaultNotificationsOk returns a tuple with the DefaultNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNotifications

`func (o *SecurityMonitoringRuleThirdPartyOptions) SetDefaultNotifications(v []string)`

SetDefaultNotifications sets DefaultNotifications field to given value.

### HasDefaultNotifications

`func (o *SecurityMonitoringRuleThirdPartyOptions) HasDefaultNotifications() bool`

HasDefaultNotifications returns a boolean if a field has been set.

### GetDefaultStatus

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultStatus() SecurityMonitoringRuleSeverity`

GetDefaultStatus returns the DefaultStatus field if non-nil, zero value otherwise.

### GetDefaultStatusOk

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetDefaultStatusOk() (*SecurityMonitoringRuleSeverity, bool)`

GetDefaultStatusOk returns a tuple with the DefaultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStatus

`func (o *SecurityMonitoringRuleThirdPartyOptions) SetDefaultStatus(v SecurityMonitoringRuleSeverity)`

SetDefaultStatus sets DefaultStatus field to given value.

### HasDefaultStatus

`func (o *SecurityMonitoringRuleThirdPartyOptions) HasDefaultStatus() bool`

HasDefaultStatus returns a boolean if a field has been set.

### GetFirstSeenOverride

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetFirstSeenOverride() string`

GetFirstSeenOverride returns the FirstSeenOverride field if non-nil, zero value otherwise.

### GetFirstSeenOverrideOk

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetFirstSeenOverrideOk() (*string, bool)`

GetFirstSeenOverrideOk returns a tuple with the FirstSeenOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenOverride

`func (o *SecurityMonitoringRuleThirdPartyOptions) SetFirstSeenOverride(v string)`

SetFirstSeenOverride sets FirstSeenOverride field to given value.

### HasFirstSeenOverride

`func (o *SecurityMonitoringRuleThirdPartyOptions) HasFirstSeenOverride() bool`

HasFirstSeenOverride returns a boolean if a field has been set.

### GetLastSeenOverride

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetLastSeenOverride() string`

GetLastSeenOverride returns the LastSeenOverride field if non-nil, zero value otherwise.

### GetLastSeenOverrideOk

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetLastSeenOverrideOk() (*string, bool)`

GetLastSeenOverrideOk returns a tuple with the LastSeenOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenOverride

`func (o *SecurityMonitoringRuleThirdPartyOptions) SetLastSeenOverride(v string)`

SetLastSeenOverride sets LastSeenOverride field to given value.

### HasLastSeenOverride

`func (o *SecurityMonitoringRuleThirdPartyOptions) HasLastSeenOverride() bool`

HasLastSeenOverride returns a boolean if a field has been set.

### GetRootQuery

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetRootQuery() string`

GetRootQuery returns the RootQuery field if non-nil, zero value otherwise.

### GetRootQueryOk

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetRootQueryOk() (*string, bool)`

GetRootQueryOk returns a tuple with the RootQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootQuery

`func (o *SecurityMonitoringRuleThirdPartyOptions) SetRootQuery(v string)`

SetRootQuery sets RootQuery field to given value.

### HasRootQuery

`func (o *SecurityMonitoringRuleThirdPartyOptions) HasRootQuery() bool`

HasRootQuery returns a boolean if a field has been set.

### GetSignalId

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetSignalId() string`

GetSignalId returns the SignalId field if non-nil, zero value otherwise.

### GetSignalIdOk

`func (o *SecurityMonitoringRuleThirdPartyOptions) GetSignalIdOk() (*string, bool)`

GetSignalIdOk returns a tuple with the SignalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalId

`func (o *SecurityMonitoringRuleThirdPartyOptions) SetSignalId(v string)`

SetSignalId sets SignalId field to given value.

### HasSignalId

`func (o *SecurityMonitoringRuleThirdPartyOptions) HasSignalId() bool`

HasSignalId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
