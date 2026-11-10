// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceDetailsDataAttributes Extended set of attributes for a single End User Device Monitoring device,
// including Cloud Network Monitoring and battery metrics.
type DeviceDetailsDataAttributes struct {
	// Public key of the Datadog Agent installed on the device.
	AgentKey *string `json:"agent_key,omitempty"`
	// Version of the Datadog Agent installed on the device.
	AgentVersion *string `json:"agent_version,omitempty"`
	// Current battery charge level as a percentage between 0 and 100.
	BatteryChargePct *int64 `json:"battery_charge_pct,omitempty"`
	// Rate at which the battery is charging or discharging, in milliamperes.
	// Negative values indicate discharge.
	BatteryChargeRate *int64 `json:"battery_charge_rate,omitempty"`
	// Whether the device's battery health has degraded to a critical level.
	BatteryCritical *bool `json:"battery_critical,omitempty"`
	// Number of full charge cycles the battery has gone through.
	BatteryCycleCount *int64 `json:"battery_cycle_count,omitempty"`
	// Maximum battery capacity expressed as a percentage of the device's design capacity.
	BatteryMaxCapacityPct *int64 `json:"battery_max_capacity_pct,omitempty"`
	// Whether Cloud Network Monitoring is enabled on the device.
	CnmEnabled *bool `json:"cnm_enabled,omitempty"`
	// Number of physical CPU cores on the device.
	CpuCores *int64 `json:"cpu_cores,omitempty"`
	// Number of logical CPU processors (hardware threads) on the device.
	CpuLogicalProcessors *int64 `json:"cpu_logical_processors,omitempty"`
	// Human-readable name of the device's CPU model.
	CpuModel *string `json:"cpu_model,omitempty"`
	// Average CPU usage on the device, as a percentage between 0 and 100.
	CpuUsage *float64 `json:"cpu_usage,omitempty"`
	// Average disk usage on the device, as a percentage between 0 and 100.
	DiskUsage *float64 `json:"disk_usage,omitempty"`
	// Directory information about the employee associated with the device,
	// sourced from a customer reference table.
	// Absent when no reference table is configured or the device has no match.
	Employee *DeviceEmployee `json:"employee,omitempty"`
	// Location derived from the device's public IP address.
	// Absent when no location data is available.
	GeoIp *DeviceGeoIP `json:"geo_ip,omitempty"`
	// Last observed IPv4 or IPv6 address of the device.
	IpAddress *string `json:"ip_address,omitempty"`
	// List of issue identifiers currently affecting the device.
	Issues []string `json:"issues,omitempty"`
	// Name of the operating system kernel running on the device.
	KernelName *string `json:"kernel_name,omitempty"`
	// Timestamp of the most recent device reboot, in RFC 3339 format.
	LastReboot *string `json:"last_reboot,omitempty"`
	// Timestamp of the most recent telemetry received from the device, in RFC 3339 format.
	LastSeen *string `json:"last_seen,omitempty"`
	// Manufacturer of the device.
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Average memory usage on the device, as a percentage between 0 and 100.
	MemUsage *float64 `json:"mem_usage,omitempty"`
	// Total amount of physical memory available on the device, in kilobytes.
	MemoryTotalKb *int64 `json:"memory_total_kb,omitempty"`
	// Marketing or product name of the device model.
	ModelName *string `json:"model_name,omitempty"`
	// Manufacturer-assigned model number of the device.
	ModelNumber *string `json:"model_number,omitempty"`
	// Operating system family running on the device, as reported by the Datadog Agent
	// (for example, `macOS`, `Windows`, or `linux`). Empty when unknown.
	Os *string `json:"os,omitempty"`
	// Operating system version running on the device.
	OsVersion *string `json:"os_version,omitempty"`
	// Number of TCP packets sent by the device over the query window,
	// as reported by Cloud Network Monitoring.
	PacketSent *float64 `json:"packet_sent,omitempty"`
	// Datadog resource identifier for the device.
	ResourceId *string `json:"resource_id,omitempty"`
	// Number of TCP packets retransmitted by the device over the query window.
	Retransmit *float64 `json:"retransmit,omitempty"`
	// Percentage of sent TCP packets that were retransmitted, between 0 and 100.
	RetransmitPct *float64 `json:"retransmit_pct,omitempty"`
	// Serial number assigned to the device by its manufacturer.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Health status of the device, computed from its issues and recent telemetry.
	// `inactive` means the device stopped reporting. `unknown` means no telemetry is available yet.
	Status *DeviceStatus `json:"status,omitempty"`
	// Number of TCP connection attempts canceled by the device over the query window.
	TcpCanceled *float64 `json:"tcp_canceled,omitempty"`
	// Number of TCP connections established by the device over the query window.
	TcpEstablished *float64 `json:"tcp_established,omitempty"`
	// Percentage of TCP connection attempts that failed, between 0 and 100.
	TcpEstablishmentFailurePct *float64 `json:"tcp_establishment_failure_pct,omitempty"`
	// Number of TCP connection attempts that failed over the query window.
	TcpEstablishmentFailures *float64 `json:"tcp_establishment_failures,omitempty"`
	// Number of TCP connection attempts that failed because the host was unreachable.
	TcpHostUnreachable *float64 `json:"tcp_host_unreachable,omitempty"`
	// Number of TCP connection attempts that failed because the network was unreachable.
	TcpNetworkUnreachable *float64 `json:"tcp_network_unreachable,omitempty"`
	// Number of TCP connection attempts refused by the remote host over the query window.
	TcpRefusals *float64 `json:"tcp_refusals,omitempty"`
	// Number of TCP connection attempts that timed out over the query window.
	TcpTimeouts *float64 `json:"tcp_timeouts,omitempty"`
	// Hardware type of the device (for example, `laptop` or `desktop`). `unknown` when the type cannot be determined.
	Type *string `json:"type,omitempty"`
	// Time elapsed since the device last booted, in seconds.
	Uptime *float64 `json:"uptime,omitempty"`
	// BSSID (MAC address of the access point) of the wireless network the device is
	// currently connected to.
	WlanBssid *string `json:"wlan_bssid,omitempty"`
	// Received signal strength indicator of the device's current wireless connection, in dBm.
	WlanRssi *float64 `json:"wlan_rssi,omitempty"`
	// SSID of the wireless network the device is currently connected to.
	WlanSsid *string `json:"wlan_ssid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeviceDetailsDataAttributes instantiates a new DeviceDetailsDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeviceDetailsDataAttributes() *DeviceDetailsDataAttributes {
	this := DeviceDetailsDataAttributes{}
	return &this
}

// NewDeviceDetailsDataAttributesWithDefaults instantiates a new DeviceDetailsDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeviceDetailsDataAttributesWithDefaults() *DeviceDetailsDataAttributes {
	this := DeviceDetailsDataAttributes{}
	return &this
}

// GetAgentKey returns the AgentKey field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetAgentKey() string {
	if o == nil || o.AgentKey == nil {
		var ret string
		return ret
	}
	return *o.AgentKey
}

// GetAgentKeyOk returns a tuple with the AgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetAgentKeyOk() (*string, bool) {
	if o == nil || o.AgentKey == nil {
		return nil, false
	}
	return o.AgentKey, true
}

// HasAgentKey returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasAgentKey() bool {
	return o != nil && o.AgentKey != nil
}

// SetAgentKey gets a reference to the given string and assigns it to the AgentKey field.
func (o *DeviceDetailsDataAttributes) SetAgentKey(v string) {
	o.AgentKey = &v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *DeviceDetailsDataAttributes) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetBatteryChargePct returns the BatteryChargePct field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetBatteryChargePct() int64 {
	if o == nil || o.BatteryChargePct == nil {
		var ret int64
		return ret
	}
	return *o.BatteryChargePct
}

// GetBatteryChargePctOk returns a tuple with the BatteryChargePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetBatteryChargePctOk() (*int64, bool) {
	if o == nil || o.BatteryChargePct == nil {
		return nil, false
	}
	return o.BatteryChargePct, true
}

// HasBatteryChargePct returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasBatteryChargePct() bool {
	return o != nil && o.BatteryChargePct != nil
}

// SetBatteryChargePct gets a reference to the given int64 and assigns it to the BatteryChargePct field.
func (o *DeviceDetailsDataAttributes) SetBatteryChargePct(v int64) {
	o.BatteryChargePct = &v
}

// GetBatteryChargeRate returns the BatteryChargeRate field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetBatteryChargeRate() int64 {
	if o == nil || o.BatteryChargeRate == nil {
		var ret int64
		return ret
	}
	return *o.BatteryChargeRate
}

// GetBatteryChargeRateOk returns a tuple with the BatteryChargeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetBatteryChargeRateOk() (*int64, bool) {
	if o == nil || o.BatteryChargeRate == nil {
		return nil, false
	}
	return o.BatteryChargeRate, true
}

// HasBatteryChargeRate returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasBatteryChargeRate() bool {
	return o != nil && o.BatteryChargeRate != nil
}

// SetBatteryChargeRate gets a reference to the given int64 and assigns it to the BatteryChargeRate field.
func (o *DeviceDetailsDataAttributes) SetBatteryChargeRate(v int64) {
	o.BatteryChargeRate = &v
}

// GetBatteryCritical returns the BatteryCritical field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetBatteryCritical() bool {
	if o == nil || o.BatteryCritical == nil {
		var ret bool
		return ret
	}
	return *o.BatteryCritical
}

// GetBatteryCriticalOk returns a tuple with the BatteryCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetBatteryCriticalOk() (*bool, bool) {
	if o == nil || o.BatteryCritical == nil {
		return nil, false
	}
	return o.BatteryCritical, true
}

// HasBatteryCritical returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasBatteryCritical() bool {
	return o != nil && o.BatteryCritical != nil
}

// SetBatteryCritical gets a reference to the given bool and assigns it to the BatteryCritical field.
func (o *DeviceDetailsDataAttributes) SetBatteryCritical(v bool) {
	o.BatteryCritical = &v
}

// GetBatteryCycleCount returns the BatteryCycleCount field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetBatteryCycleCount() int64 {
	if o == nil || o.BatteryCycleCount == nil {
		var ret int64
		return ret
	}
	return *o.BatteryCycleCount
}

// GetBatteryCycleCountOk returns a tuple with the BatteryCycleCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetBatteryCycleCountOk() (*int64, bool) {
	if o == nil || o.BatteryCycleCount == nil {
		return nil, false
	}
	return o.BatteryCycleCount, true
}

// HasBatteryCycleCount returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasBatteryCycleCount() bool {
	return o != nil && o.BatteryCycleCount != nil
}

// SetBatteryCycleCount gets a reference to the given int64 and assigns it to the BatteryCycleCount field.
func (o *DeviceDetailsDataAttributes) SetBatteryCycleCount(v int64) {
	o.BatteryCycleCount = &v
}

// GetBatteryMaxCapacityPct returns the BatteryMaxCapacityPct field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetBatteryMaxCapacityPct() int64 {
	if o == nil || o.BatteryMaxCapacityPct == nil {
		var ret int64
		return ret
	}
	return *o.BatteryMaxCapacityPct
}

// GetBatteryMaxCapacityPctOk returns a tuple with the BatteryMaxCapacityPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetBatteryMaxCapacityPctOk() (*int64, bool) {
	if o == nil || o.BatteryMaxCapacityPct == nil {
		return nil, false
	}
	return o.BatteryMaxCapacityPct, true
}

// HasBatteryMaxCapacityPct returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasBatteryMaxCapacityPct() bool {
	return o != nil && o.BatteryMaxCapacityPct != nil
}

// SetBatteryMaxCapacityPct gets a reference to the given int64 and assigns it to the BatteryMaxCapacityPct field.
func (o *DeviceDetailsDataAttributes) SetBatteryMaxCapacityPct(v int64) {
	o.BatteryMaxCapacityPct = &v
}

// GetCnmEnabled returns the CnmEnabled field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetCnmEnabled() bool {
	if o == nil || o.CnmEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CnmEnabled
}

// GetCnmEnabledOk returns a tuple with the CnmEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetCnmEnabledOk() (*bool, bool) {
	if o == nil || o.CnmEnabled == nil {
		return nil, false
	}
	return o.CnmEnabled, true
}

// HasCnmEnabled returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasCnmEnabled() bool {
	return o != nil && o.CnmEnabled != nil
}

// SetCnmEnabled gets a reference to the given bool and assigns it to the CnmEnabled field.
func (o *DeviceDetailsDataAttributes) SetCnmEnabled(v bool) {
	o.CnmEnabled = &v
}

// GetCpuCores returns the CpuCores field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetCpuCores() int64 {
	if o == nil || o.CpuCores == nil {
		var ret int64
		return ret
	}
	return *o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetCpuCoresOk() (*int64, bool) {
	if o == nil || o.CpuCores == nil {
		return nil, false
	}
	return o.CpuCores, true
}

// HasCpuCores returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasCpuCores() bool {
	return o != nil && o.CpuCores != nil
}

// SetCpuCores gets a reference to the given int64 and assigns it to the CpuCores field.
func (o *DeviceDetailsDataAttributes) SetCpuCores(v int64) {
	o.CpuCores = &v
}

// GetCpuLogicalProcessors returns the CpuLogicalProcessors field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetCpuLogicalProcessors() int64 {
	if o == nil || o.CpuLogicalProcessors == nil {
		var ret int64
		return ret
	}
	return *o.CpuLogicalProcessors
}

// GetCpuLogicalProcessorsOk returns a tuple with the CpuLogicalProcessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetCpuLogicalProcessorsOk() (*int64, bool) {
	if o == nil || o.CpuLogicalProcessors == nil {
		return nil, false
	}
	return o.CpuLogicalProcessors, true
}

// HasCpuLogicalProcessors returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasCpuLogicalProcessors() bool {
	return o != nil && o.CpuLogicalProcessors != nil
}

// SetCpuLogicalProcessors gets a reference to the given int64 and assigns it to the CpuLogicalProcessors field.
func (o *DeviceDetailsDataAttributes) SetCpuLogicalProcessors(v int64) {
	o.CpuLogicalProcessors = &v
}

// GetCpuModel returns the CpuModel field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetCpuModel() string {
	if o == nil || o.CpuModel == nil {
		var ret string
		return ret
	}
	return *o.CpuModel
}

// GetCpuModelOk returns a tuple with the CpuModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetCpuModelOk() (*string, bool) {
	if o == nil || o.CpuModel == nil {
		return nil, false
	}
	return o.CpuModel, true
}

// HasCpuModel returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasCpuModel() bool {
	return o != nil && o.CpuModel != nil
}

// SetCpuModel gets a reference to the given string and assigns it to the CpuModel field.
func (o *DeviceDetailsDataAttributes) SetCpuModel(v string) {
	o.CpuModel = &v
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetCpuUsage() float64 {
	if o == nil || o.CpuUsage == nil {
		var ret float64
		return ret
	}
	return *o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetCpuUsageOk() (*float64, bool) {
	if o == nil || o.CpuUsage == nil {
		return nil, false
	}
	return o.CpuUsage, true
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasCpuUsage() bool {
	return o != nil && o.CpuUsage != nil
}

// SetCpuUsage gets a reference to the given float64 and assigns it to the CpuUsage field.
func (o *DeviceDetailsDataAttributes) SetCpuUsage(v float64) {
	o.CpuUsage = &v
}

// GetDiskUsage returns the DiskUsage field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetDiskUsage() float64 {
	if o == nil || o.DiskUsage == nil {
		var ret float64
		return ret
	}
	return *o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetDiskUsageOk() (*float64, bool) {
	if o == nil || o.DiskUsage == nil {
		return nil, false
	}
	return o.DiskUsage, true
}

// HasDiskUsage returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasDiskUsage() bool {
	return o != nil && o.DiskUsage != nil
}

// SetDiskUsage gets a reference to the given float64 and assigns it to the DiskUsage field.
func (o *DeviceDetailsDataAttributes) SetDiskUsage(v float64) {
	o.DiskUsage = &v
}

// GetEmployee returns the Employee field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetEmployee() DeviceEmployee {
	if o == nil || o.Employee == nil {
		var ret DeviceEmployee
		return ret
	}
	return *o.Employee
}

// GetEmployeeOk returns a tuple with the Employee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetEmployeeOk() (*DeviceEmployee, bool) {
	if o == nil || o.Employee == nil {
		return nil, false
	}
	return o.Employee, true
}

// HasEmployee returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasEmployee() bool {
	return o != nil && o.Employee != nil
}

// SetEmployee gets a reference to the given DeviceEmployee and assigns it to the Employee field.
func (o *DeviceDetailsDataAttributes) SetEmployee(v DeviceEmployee) {
	o.Employee = &v
}

// GetGeoIp returns the GeoIp field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetGeoIp() DeviceGeoIP {
	if o == nil || o.GeoIp == nil {
		var ret DeviceGeoIP
		return ret
	}
	return *o.GeoIp
}

// GetGeoIpOk returns a tuple with the GeoIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetGeoIpOk() (*DeviceGeoIP, bool) {
	if o == nil || o.GeoIp == nil {
		return nil, false
	}
	return o.GeoIp, true
}

// HasGeoIp returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasGeoIp() bool {
	return o != nil && o.GeoIp != nil
}

// SetGeoIp gets a reference to the given DeviceGeoIP and assigns it to the GeoIp field.
func (o *DeviceDetailsDataAttributes) SetGeoIp(v DeviceGeoIP) {
	o.GeoIp = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasIpAddress() bool {
	return o != nil && o.IpAddress != nil
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *DeviceDetailsDataAttributes) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetIssues() []string {
	if o == nil || o.Issues == nil {
		var ret []string
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetIssuesOk() (*[]string, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return &o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasIssues() bool {
	return o != nil && o.Issues != nil
}

// SetIssues gets a reference to the given []string and assigns it to the Issues field.
func (o *DeviceDetailsDataAttributes) SetIssues(v []string) {
	o.Issues = v
}

// GetKernelName returns the KernelName field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetKernelName() string {
	if o == nil || o.KernelName == nil {
		var ret string
		return ret
	}
	return *o.KernelName
}

// GetKernelNameOk returns a tuple with the KernelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetKernelNameOk() (*string, bool) {
	if o == nil || o.KernelName == nil {
		return nil, false
	}
	return o.KernelName, true
}

// HasKernelName returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasKernelName() bool {
	return o != nil && o.KernelName != nil
}

// SetKernelName gets a reference to the given string and assigns it to the KernelName field.
func (o *DeviceDetailsDataAttributes) SetKernelName(v string) {
	o.KernelName = &v
}

// GetLastReboot returns the LastReboot field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetLastReboot() string {
	if o == nil || o.LastReboot == nil {
		var ret string
		return ret
	}
	return *o.LastReboot
}

// GetLastRebootOk returns a tuple with the LastReboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetLastRebootOk() (*string, bool) {
	if o == nil || o.LastReboot == nil {
		return nil, false
	}
	return o.LastReboot, true
}

// HasLastReboot returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasLastReboot() bool {
	return o != nil && o.LastReboot != nil
}

// SetLastReboot gets a reference to the given string and assigns it to the LastReboot field.
func (o *DeviceDetailsDataAttributes) SetLastReboot(v string) {
	o.LastReboot = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *DeviceDetailsDataAttributes) SetLastSeen(v string) {
	o.LastSeen = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasManufacturer() bool {
	return o != nil && o.Manufacturer != nil
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *DeviceDetailsDataAttributes) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetMemUsage returns the MemUsage field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetMemUsage() float64 {
	if o == nil || o.MemUsage == nil {
		var ret float64
		return ret
	}
	return *o.MemUsage
}

// GetMemUsageOk returns a tuple with the MemUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetMemUsageOk() (*float64, bool) {
	if o == nil || o.MemUsage == nil {
		return nil, false
	}
	return o.MemUsage, true
}

// HasMemUsage returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasMemUsage() bool {
	return o != nil && o.MemUsage != nil
}

// SetMemUsage gets a reference to the given float64 and assigns it to the MemUsage field.
func (o *DeviceDetailsDataAttributes) SetMemUsage(v float64) {
	o.MemUsage = &v
}

// GetMemoryTotalKb returns the MemoryTotalKb field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetMemoryTotalKb() int64 {
	if o == nil || o.MemoryTotalKb == nil {
		var ret int64
		return ret
	}
	return *o.MemoryTotalKb
}

// GetMemoryTotalKbOk returns a tuple with the MemoryTotalKb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetMemoryTotalKbOk() (*int64, bool) {
	if o == nil || o.MemoryTotalKb == nil {
		return nil, false
	}
	return o.MemoryTotalKb, true
}

// HasMemoryTotalKb returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasMemoryTotalKb() bool {
	return o != nil && o.MemoryTotalKb != nil
}

// SetMemoryTotalKb gets a reference to the given int64 and assigns it to the MemoryTotalKb field.
func (o *DeviceDetailsDataAttributes) SetMemoryTotalKb(v int64) {
	o.MemoryTotalKb = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasModelName() bool {
	return o != nil && o.ModelName != nil
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *DeviceDetailsDataAttributes) SetModelName(v string) {
	o.ModelName = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasModelNumber() bool {
	return o != nil && o.ModelNumber != nil
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *DeviceDetailsDataAttributes) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasOs() bool {
	return o != nil && o.Os != nil
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *DeviceDetailsDataAttributes) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasOsVersion() bool {
	return o != nil && o.OsVersion != nil
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *DeviceDetailsDataAttributes) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPacketSent returns the PacketSent field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetPacketSent() float64 {
	if o == nil || o.PacketSent == nil {
		var ret float64
		return ret
	}
	return *o.PacketSent
}

// GetPacketSentOk returns a tuple with the PacketSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetPacketSentOk() (*float64, bool) {
	if o == nil || o.PacketSent == nil {
		return nil, false
	}
	return o.PacketSent, true
}

// HasPacketSent returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasPacketSent() bool {
	return o != nil && o.PacketSent != nil
}

// SetPacketSent gets a reference to the given float64 and assigns it to the PacketSent field.
func (o *DeviceDetailsDataAttributes) SetPacketSent(v float64) {
	o.PacketSent = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *DeviceDetailsDataAttributes) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetRetransmit returns the Retransmit field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetRetransmit() float64 {
	if o == nil || o.Retransmit == nil {
		var ret float64
		return ret
	}
	return *o.Retransmit
}

// GetRetransmitOk returns a tuple with the Retransmit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetRetransmitOk() (*float64, bool) {
	if o == nil || o.Retransmit == nil {
		return nil, false
	}
	return o.Retransmit, true
}

// HasRetransmit returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasRetransmit() bool {
	return o != nil && o.Retransmit != nil
}

// SetRetransmit gets a reference to the given float64 and assigns it to the Retransmit field.
func (o *DeviceDetailsDataAttributes) SetRetransmit(v float64) {
	o.Retransmit = &v
}

// GetRetransmitPct returns the RetransmitPct field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetRetransmitPct() float64 {
	if o == nil || o.RetransmitPct == nil {
		var ret float64
		return ret
	}
	return *o.RetransmitPct
}

// GetRetransmitPctOk returns a tuple with the RetransmitPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetRetransmitPctOk() (*float64, bool) {
	if o == nil || o.RetransmitPct == nil {
		return nil, false
	}
	return o.RetransmitPct, true
}

// HasRetransmitPct returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasRetransmitPct() bool {
	return o != nil && o.RetransmitPct != nil
}

// SetRetransmitPct gets a reference to the given float64 and assigns it to the RetransmitPct field.
func (o *DeviceDetailsDataAttributes) SetRetransmitPct(v float64) {
	o.RetransmitPct = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasSerialNumber() bool {
	return o != nil && o.SerialNumber != nil
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *DeviceDetailsDataAttributes) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetStatus() DeviceStatus {
	if o == nil || o.Status == nil {
		var ret DeviceStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetStatusOk() (*DeviceStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given DeviceStatus and assigns it to the Status field.
func (o *DeviceDetailsDataAttributes) SetStatus(v DeviceStatus) {
	o.Status = &v
}

// GetTcpCanceled returns the TcpCanceled field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpCanceled() float64 {
	if o == nil || o.TcpCanceled == nil {
		var ret float64
		return ret
	}
	return *o.TcpCanceled
}

// GetTcpCanceledOk returns a tuple with the TcpCanceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpCanceledOk() (*float64, bool) {
	if o == nil || o.TcpCanceled == nil {
		return nil, false
	}
	return o.TcpCanceled, true
}

// HasTcpCanceled returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpCanceled() bool {
	return o != nil && o.TcpCanceled != nil
}

// SetTcpCanceled gets a reference to the given float64 and assigns it to the TcpCanceled field.
func (o *DeviceDetailsDataAttributes) SetTcpCanceled(v float64) {
	o.TcpCanceled = &v
}

// GetTcpEstablished returns the TcpEstablished field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpEstablished() float64 {
	if o == nil || o.TcpEstablished == nil {
		var ret float64
		return ret
	}
	return *o.TcpEstablished
}

// GetTcpEstablishedOk returns a tuple with the TcpEstablished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpEstablishedOk() (*float64, bool) {
	if o == nil || o.TcpEstablished == nil {
		return nil, false
	}
	return o.TcpEstablished, true
}

// HasTcpEstablished returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpEstablished() bool {
	return o != nil && o.TcpEstablished != nil
}

// SetTcpEstablished gets a reference to the given float64 and assigns it to the TcpEstablished field.
func (o *DeviceDetailsDataAttributes) SetTcpEstablished(v float64) {
	o.TcpEstablished = &v
}

// GetTcpEstablishmentFailurePct returns the TcpEstablishmentFailurePct field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpEstablishmentFailurePct() float64 {
	if o == nil || o.TcpEstablishmentFailurePct == nil {
		var ret float64
		return ret
	}
	return *o.TcpEstablishmentFailurePct
}

// GetTcpEstablishmentFailurePctOk returns a tuple with the TcpEstablishmentFailurePct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpEstablishmentFailurePctOk() (*float64, bool) {
	if o == nil || o.TcpEstablishmentFailurePct == nil {
		return nil, false
	}
	return o.TcpEstablishmentFailurePct, true
}

// HasTcpEstablishmentFailurePct returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpEstablishmentFailurePct() bool {
	return o != nil && o.TcpEstablishmentFailurePct != nil
}

// SetTcpEstablishmentFailurePct gets a reference to the given float64 and assigns it to the TcpEstablishmentFailurePct field.
func (o *DeviceDetailsDataAttributes) SetTcpEstablishmentFailurePct(v float64) {
	o.TcpEstablishmentFailurePct = &v
}

// GetTcpEstablishmentFailures returns the TcpEstablishmentFailures field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpEstablishmentFailures() float64 {
	if o == nil || o.TcpEstablishmentFailures == nil {
		var ret float64
		return ret
	}
	return *o.TcpEstablishmentFailures
}

// GetTcpEstablishmentFailuresOk returns a tuple with the TcpEstablishmentFailures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpEstablishmentFailuresOk() (*float64, bool) {
	if o == nil || o.TcpEstablishmentFailures == nil {
		return nil, false
	}
	return o.TcpEstablishmentFailures, true
}

// HasTcpEstablishmentFailures returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpEstablishmentFailures() bool {
	return o != nil && o.TcpEstablishmentFailures != nil
}

// SetTcpEstablishmentFailures gets a reference to the given float64 and assigns it to the TcpEstablishmentFailures field.
func (o *DeviceDetailsDataAttributes) SetTcpEstablishmentFailures(v float64) {
	o.TcpEstablishmentFailures = &v
}

// GetTcpHostUnreachable returns the TcpHostUnreachable field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpHostUnreachable() float64 {
	if o == nil || o.TcpHostUnreachable == nil {
		var ret float64
		return ret
	}
	return *o.TcpHostUnreachable
}

// GetTcpHostUnreachableOk returns a tuple with the TcpHostUnreachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpHostUnreachableOk() (*float64, bool) {
	if o == nil || o.TcpHostUnreachable == nil {
		return nil, false
	}
	return o.TcpHostUnreachable, true
}

// HasTcpHostUnreachable returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpHostUnreachable() bool {
	return o != nil && o.TcpHostUnreachable != nil
}

// SetTcpHostUnreachable gets a reference to the given float64 and assigns it to the TcpHostUnreachable field.
func (o *DeviceDetailsDataAttributes) SetTcpHostUnreachable(v float64) {
	o.TcpHostUnreachable = &v
}

// GetTcpNetworkUnreachable returns the TcpNetworkUnreachable field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpNetworkUnreachable() float64 {
	if o == nil || o.TcpNetworkUnreachable == nil {
		var ret float64
		return ret
	}
	return *o.TcpNetworkUnreachable
}

// GetTcpNetworkUnreachableOk returns a tuple with the TcpNetworkUnreachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpNetworkUnreachableOk() (*float64, bool) {
	if o == nil || o.TcpNetworkUnreachable == nil {
		return nil, false
	}
	return o.TcpNetworkUnreachable, true
}

// HasTcpNetworkUnreachable returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpNetworkUnreachable() bool {
	return o != nil && o.TcpNetworkUnreachable != nil
}

// SetTcpNetworkUnreachable gets a reference to the given float64 and assigns it to the TcpNetworkUnreachable field.
func (o *DeviceDetailsDataAttributes) SetTcpNetworkUnreachable(v float64) {
	o.TcpNetworkUnreachable = &v
}

// GetTcpRefusals returns the TcpRefusals field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpRefusals() float64 {
	if o == nil || o.TcpRefusals == nil {
		var ret float64
		return ret
	}
	return *o.TcpRefusals
}

// GetTcpRefusalsOk returns a tuple with the TcpRefusals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpRefusalsOk() (*float64, bool) {
	if o == nil || o.TcpRefusals == nil {
		return nil, false
	}
	return o.TcpRefusals, true
}

// HasTcpRefusals returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpRefusals() bool {
	return o != nil && o.TcpRefusals != nil
}

// SetTcpRefusals gets a reference to the given float64 and assigns it to the TcpRefusals field.
func (o *DeviceDetailsDataAttributes) SetTcpRefusals(v float64) {
	o.TcpRefusals = &v
}

// GetTcpTimeouts returns the TcpTimeouts field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpTimeouts() float64 {
	if o == nil || o.TcpTimeouts == nil {
		var ret float64
		return ret
	}
	return *o.TcpTimeouts
}

// GetTcpTimeoutsOk returns a tuple with the TcpTimeouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpTimeoutsOk() (*float64, bool) {
	if o == nil || o.TcpTimeouts == nil {
		return nil, false
	}
	return o.TcpTimeouts, true
}

// HasTcpTimeouts returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpTimeouts() bool {
	return o != nil && o.TcpTimeouts != nil
}

// SetTcpTimeouts gets a reference to the given float64 and assigns it to the TcpTimeouts field.
func (o *DeviceDetailsDataAttributes) SetTcpTimeouts(v float64) {
	o.TcpTimeouts = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeviceDetailsDataAttributes) SetType(v string) {
	o.Type = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetUptime() float64 {
	if o == nil || o.Uptime == nil {
		var ret float64
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetUptimeOk() (*float64, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasUptime() bool {
	return o != nil && o.Uptime != nil
}

// SetUptime gets a reference to the given float64 and assigns it to the Uptime field.
func (o *DeviceDetailsDataAttributes) SetUptime(v float64) {
	o.Uptime = &v
}

// GetWlanBssid returns the WlanBssid field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetWlanBssid() string {
	if o == nil || o.WlanBssid == nil {
		var ret string
		return ret
	}
	return *o.WlanBssid
}

// GetWlanBssidOk returns a tuple with the WlanBssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetWlanBssidOk() (*string, bool) {
	if o == nil || o.WlanBssid == nil {
		return nil, false
	}
	return o.WlanBssid, true
}

// HasWlanBssid returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasWlanBssid() bool {
	return o != nil && o.WlanBssid != nil
}

// SetWlanBssid gets a reference to the given string and assigns it to the WlanBssid field.
func (o *DeviceDetailsDataAttributes) SetWlanBssid(v string) {
	o.WlanBssid = &v
}

// GetWlanRssi returns the WlanRssi field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetWlanRssi() float64 {
	if o == nil || o.WlanRssi == nil {
		var ret float64
		return ret
	}
	return *o.WlanRssi
}

// GetWlanRssiOk returns a tuple with the WlanRssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetWlanRssiOk() (*float64, bool) {
	if o == nil || o.WlanRssi == nil {
		return nil, false
	}
	return o.WlanRssi, true
}

// HasWlanRssi returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasWlanRssi() bool {
	return o != nil && o.WlanRssi != nil
}

// SetWlanRssi gets a reference to the given float64 and assigns it to the WlanRssi field.
func (o *DeviceDetailsDataAttributes) SetWlanRssi(v float64) {
	o.WlanRssi = &v
}

// GetWlanSsid returns the WlanSsid field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetWlanSsid() string {
	if o == nil || o.WlanSsid == nil {
		var ret string
		return ret
	}
	return *o.WlanSsid
}

// GetWlanSsidOk returns a tuple with the WlanSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetWlanSsidOk() (*string, bool) {
	if o == nil || o.WlanSsid == nil {
		return nil, false
	}
	return o.WlanSsid, true
}

// HasWlanSsid returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasWlanSsid() bool {
	return o != nil && o.WlanSsid != nil
}

// SetWlanSsid gets a reference to the given string and assigns it to the WlanSsid field.
func (o *DeviceDetailsDataAttributes) SetWlanSsid(v string) {
	o.WlanSsid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeviceDetailsDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AgentKey != nil {
		toSerialize["agent_key"] = o.AgentKey
	}
	if o.AgentVersion != nil {
		toSerialize["agent_version"] = o.AgentVersion
	}
	if o.BatteryChargePct != nil {
		toSerialize["battery_charge_pct"] = o.BatteryChargePct
	}
	if o.BatteryChargeRate != nil {
		toSerialize["battery_charge_rate"] = o.BatteryChargeRate
	}
	if o.BatteryCritical != nil {
		toSerialize["battery_critical"] = o.BatteryCritical
	}
	if o.BatteryCycleCount != nil {
		toSerialize["battery_cycle_count"] = o.BatteryCycleCount
	}
	if o.BatteryMaxCapacityPct != nil {
		toSerialize["battery_max_capacity_pct"] = o.BatteryMaxCapacityPct
	}
	if o.CnmEnabled != nil {
		toSerialize["cnm_enabled"] = o.CnmEnabled
	}
	if o.CpuCores != nil {
		toSerialize["cpu_cores"] = o.CpuCores
	}
	if o.CpuLogicalProcessors != nil {
		toSerialize["cpu_logical_processors"] = o.CpuLogicalProcessors
	}
	if o.CpuModel != nil {
		toSerialize["cpu_model"] = o.CpuModel
	}
	if o.CpuUsage != nil {
		toSerialize["cpu_usage"] = o.CpuUsage
	}
	if o.DiskUsage != nil {
		toSerialize["disk_usage"] = o.DiskUsage
	}
	if o.Employee != nil {
		toSerialize["employee"] = o.Employee
	}
	if o.GeoIp != nil {
		toSerialize["geo_ip"] = o.GeoIp
	}
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}
	if o.KernelName != nil {
		toSerialize["kernel_name"] = o.KernelName
	}
	if o.LastReboot != nil {
		toSerialize["last_reboot"] = o.LastReboot
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	if o.Manufacturer != nil {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if o.MemUsage != nil {
		toSerialize["mem_usage"] = o.MemUsage
	}
	if o.MemoryTotalKb != nil {
		toSerialize["memory_total_kb"] = o.MemoryTotalKb
	}
	if o.ModelName != nil {
		toSerialize["model_name"] = o.ModelName
	}
	if o.ModelNumber != nil {
		toSerialize["model_number"] = o.ModelNumber
	}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.PacketSent != nil {
		toSerialize["packet_sent"] = o.PacketSent
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.Retransmit != nil {
		toSerialize["retransmit"] = o.Retransmit
	}
	if o.RetransmitPct != nil {
		toSerialize["retransmit_pct"] = o.RetransmitPct
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TcpCanceled != nil {
		toSerialize["tcp_canceled"] = o.TcpCanceled
	}
	if o.TcpEstablished != nil {
		toSerialize["tcp_established"] = o.TcpEstablished
	}
	if o.TcpEstablishmentFailurePct != nil {
		toSerialize["tcp_establishment_failure_pct"] = o.TcpEstablishmentFailurePct
	}
	if o.TcpEstablishmentFailures != nil {
		toSerialize["tcp_establishment_failures"] = o.TcpEstablishmentFailures
	}
	if o.TcpHostUnreachable != nil {
		toSerialize["tcp_host_unreachable"] = o.TcpHostUnreachable
	}
	if o.TcpNetworkUnreachable != nil {
		toSerialize["tcp_network_unreachable"] = o.TcpNetworkUnreachable
	}
	if o.TcpRefusals != nil {
		toSerialize["tcp_refusals"] = o.TcpRefusals
	}
	if o.TcpTimeouts != nil {
		toSerialize["tcp_timeouts"] = o.TcpTimeouts
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Uptime != nil {
		toSerialize["uptime"] = o.Uptime
	}
	if o.WlanBssid != nil {
		toSerialize["wlan_bssid"] = o.WlanBssid
	}
	if o.WlanRssi != nil {
		toSerialize["wlan_rssi"] = o.WlanRssi
	}
	if o.WlanSsid != nil {
		toSerialize["wlan_ssid"] = o.WlanSsid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeviceDetailsDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentKey                   *string         `json:"agent_key,omitempty"`
		AgentVersion               *string         `json:"agent_version,omitempty"`
		BatteryChargePct           *int64          `json:"battery_charge_pct,omitempty"`
		BatteryChargeRate          *int64          `json:"battery_charge_rate,omitempty"`
		BatteryCritical            *bool           `json:"battery_critical,omitempty"`
		BatteryCycleCount          *int64          `json:"battery_cycle_count,omitempty"`
		BatteryMaxCapacityPct      *int64          `json:"battery_max_capacity_pct,omitempty"`
		CnmEnabled                 *bool           `json:"cnm_enabled,omitempty"`
		CpuCores                   *int64          `json:"cpu_cores,omitempty"`
		CpuLogicalProcessors       *int64          `json:"cpu_logical_processors,omitempty"`
		CpuModel                   *string         `json:"cpu_model,omitempty"`
		CpuUsage                   *float64        `json:"cpu_usage,omitempty"`
		DiskUsage                  *float64        `json:"disk_usage,omitempty"`
		Employee                   *DeviceEmployee `json:"employee,omitempty"`
		GeoIp                      *DeviceGeoIP    `json:"geo_ip,omitempty"`
		IpAddress                  *string         `json:"ip_address,omitempty"`
		Issues                     []string        `json:"issues,omitempty"`
		KernelName                 *string         `json:"kernel_name,omitempty"`
		LastReboot                 *string         `json:"last_reboot,omitempty"`
		LastSeen                   *string         `json:"last_seen,omitempty"`
		Manufacturer               *string         `json:"manufacturer,omitempty"`
		MemUsage                   *float64        `json:"mem_usage,omitempty"`
		MemoryTotalKb              *int64          `json:"memory_total_kb,omitempty"`
		ModelName                  *string         `json:"model_name,omitempty"`
		ModelNumber                *string         `json:"model_number,omitempty"`
		Os                         *string         `json:"os,omitempty"`
		OsVersion                  *string         `json:"os_version,omitempty"`
		PacketSent                 *float64        `json:"packet_sent,omitempty"`
		ResourceId                 *string         `json:"resource_id,omitempty"`
		Retransmit                 *float64        `json:"retransmit,omitempty"`
		RetransmitPct              *float64        `json:"retransmit_pct,omitempty"`
		SerialNumber               *string         `json:"serial_number,omitempty"`
		Status                     *DeviceStatus   `json:"status,omitempty"`
		TcpCanceled                *float64        `json:"tcp_canceled,omitempty"`
		TcpEstablished             *float64        `json:"tcp_established,omitempty"`
		TcpEstablishmentFailurePct *float64        `json:"tcp_establishment_failure_pct,omitempty"`
		TcpEstablishmentFailures   *float64        `json:"tcp_establishment_failures,omitempty"`
		TcpHostUnreachable         *float64        `json:"tcp_host_unreachable,omitempty"`
		TcpNetworkUnreachable      *float64        `json:"tcp_network_unreachable,omitempty"`
		TcpRefusals                *float64        `json:"tcp_refusals,omitempty"`
		TcpTimeouts                *float64        `json:"tcp_timeouts,omitempty"`
		Type                       *string         `json:"type,omitempty"`
		Uptime                     *float64        `json:"uptime,omitempty"`
		WlanBssid                  *string         `json:"wlan_bssid,omitempty"`
		WlanRssi                   *float64        `json:"wlan_rssi,omitempty"`
		WlanSsid                   *string         `json:"wlan_ssid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_key", "agent_version", "battery_charge_pct", "battery_charge_rate", "battery_critical", "battery_cycle_count", "battery_max_capacity_pct", "cnm_enabled", "cpu_cores", "cpu_logical_processors", "cpu_model", "cpu_usage", "disk_usage", "employee", "geo_ip", "ip_address", "issues", "kernel_name", "last_reboot", "last_seen", "manufacturer", "mem_usage", "memory_total_kb", "model_name", "model_number", "os", "os_version", "packet_sent", "resource_id", "retransmit", "retransmit_pct", "serial_number", "status", "tcp_canceled", "tcp_established", "tcp_establishment_failure_pct", "tcp_establishment_failures", "tcp_host_unreachable", "tcp_network_unreachable", "tcp_refusals", "tcp_timeouts", "type", "uptime", "wlan_bssid", "wlan_rssi", "wlan_ssid"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AgentKey = all.AgentKey
	o.AgentVersion = all.AgentVersion
	o.BatteryChargePct = all.BatteryChargePct
	o.BatteryChargeRate = all.BatteryChargeRate
	o.BatteryCritical = all.BatteryCritical
	o.BatteryCycleCount = all.BatteryCycleCount
	o.BatteryMaxCapacityPct = all.BatteryMaxCapacityPct
	o.CnmEnabled = all.CnmEnabled
	o.CpuCores = all.CpuCores
	o.CpuLogicalProcessors = all.CpuLogicalProcessors
	o.CpuModel = all.CpuModel
	o.CpuUsage = all.CpuUsage
	o.DiskUsage = all.DiskUsage
	if all.Employee != nil && all.Employee.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Employee = all.Employee
	if all.GeoIp != nil && all.GeoIp.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GeoIp = all.GeoIp
	o.IpAddress = all.IpAddress
	o.Issues = all.Issues
	o.KernelName = all.KernelName
	o.LastReboot = all.LastReboot
	o.LastSeen = all.LastSeen
	o.Manufacturer = all.Manufacturer
	o.MemUsage = all.MemUsage
	o.MemoryTotalKb = all.MemoryTotalKb
	o.ModelName = all.ModelName
	o.ModelNumber = all.ModelNumber
	o.Os = all.Os
	o.OsVersion = all.OsVersion
	o.PacketSent = all.PacketSent
	o.ResourceId = all.ResourceId
	o.Retransmit = all.Retransmit
	o.RetransmitPct = all.RetransmitPct
	o.SerialNumber = all.SerialNumber
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.TcpCanceled = all.TcpCanceled
	o.TcpEstablished = all.TcpEstablished
	o.TcpEstablishmentFailurePct = all.TcpEstablishmentFailurePct
	o.TcpEstablishmentFailures = all.TcpEstablishmentFailures
	o.TcpHostUnreachable = all.TcpHostUnreachable
	o.TcpNetworkUnreachable = all.TcpNetworkUnreachable
	o.TcpRefusals = all.TcpRefusals
	o.TcpTimeouts = all.TcpTimeouts
	o.Type = all.Type
	o.Uptime = all.Uptime
	o.WlanBssid = all.WlanBssid
	o.WlanRssi = all.WlanRssi
	o.WlanSsid = all.WlanSsid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
