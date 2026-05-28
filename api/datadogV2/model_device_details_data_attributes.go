// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceDetailsDataAttributes Extended set of attributes for a single End User Device Monitoring device,
// including detailed network and battery metrics.
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
	// Number of full charge cycles the battery has gone through.
	BatteryCycleCount *int64 `json:"battery_cycle_count,omitempty"`
	// Maximum battery capacity expressed as a percentage of the device's design capacity.
	BatteryMaxCapacityPct *int64 `json:"battery_max_capacity_pct,omitempty"`
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
	// Last observed IPv4 or IPv6 address of the device.
	IpAddress *string `json:"ip_address,omitempty"`
	// List of issue identifiers currently affecting the device.
	// References entries returned by the issues endpoint.
	Issues []string `json:"issues,omitempty"`
	// Name of the operating system kernel running on the device.
	KernelName *string `json:"kernel_name,omitempty"`
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
	// Operating system family running on the device (for example, `mac`, `windows`, or `linux`).
	Os *string `json:"os,omitempty"`
	// Operating system version running on the device.
	OsVersion *string `json:"os_version,omitempty"`
	// Average rate of dropped inbound network packets, in packets per second.
	PacketsInDrop *float64 `json:"packets_in_drop,omitempty"`
	// Average rate of inbound network packets received with errors, in packets per second.
	PacketsInError *float64 `json:"packets_in_error,omitempty"`
	// Average rate of dropped outbound network packets, in packets per second.
	PacketsOutDrop *float64 `json:"packets_out_drop,omitempty"`
	// Average rate of outbound network packets sent with errors, in packets per second.
	PacketsOutError *float64 `json:"packets_out_error,omitempty"`
	// Datadog resource identifier for the device.
	ResourceId *string `json:"resource_id,omitempty"`
	// Serial number assigned to the device by its manufacturer.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Health status of the device computed from its issues and recent telemetry.
	Status *string `json:"status,omitempty"`
	// Average rate of TCP segments sent by the device, in segments per second.
	TcpOutSegs *float64 `json:"tcp_out_segs,omitempty"`
	// Average rate of TCP segments retransmitted by the device, in segments per second.
	TcpRetransSegs *float64 `json:"tcp_retrans_segs,omitempty"`
	// Hardware type of the device (for example, `laptop`, `desktop`, or `mobile`).
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

// GetPacketsInDrop returns the PacketsInDrop field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetPacketsInDrop() float64 {
	if o == nil || o.PacketsInDrop == nil {
		var ret float64
		return ret
	}
	return *o.PacketsInDrop
}

// GetPacketsInDropOk returns a tuple with the PacketsInDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetPacketsInDropOk() (*float64, bool) {
	if o == nil || o.PacketsInDrop == nil {
		return nil, false
	}
	return o.PacketsInDrop, true
}

// HasPacketsInDrop returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasPacketsInDrop() bool {
	return o != nil && o.PacketsInDrop != nil
}

// SetPacketsInDrop gets a reference to the given float64 and assigns it to the PacketsInDrop field.
func (o *DeviceDetailsDataAttributes) SetPacketsInDrop(v float64) {
	o.PacketsInDrop = &v
}

// GetPacketsInError returns the PacketsInError field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetPacketsInError() float64 {
	if o == nil || o.PacketsInError == nil {
		var ret float64
		return ret
	}
	return *o.PacketsInError
}

// GetPacketsInErrorOk returns a tuple with the PacketsInError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetPacketsInErrorOk() (*float64, bool) {
	if o == nil || o.PacketsInError == nil {
		return nil, false
	}
	return o.PacketsInError, true
}

// HasPacketsInError returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasPacketsInError() bool {
	return o != nil && o.PacketsInError != nil
}

// SetPacketsInError gets a reference to the given float64 and assigns it to the PacketsInError field.
func (o *DeviceDetailsDataAttributes) SetPacketsInError(v float64) {
	o.PacketsInError = &v
}

// GetPacketsOutDrop returns the PacketsOutDrop field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetPacketsOutDrop() float64 {
	if o == nil || o.PacketsOutDrop == nil {
		var ret float64
		return ret
	}
	return *o.PacketsOutDrop
}

// GetPacketsOutDropOk returns a tuple with the PacketsOutDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetPacketsOutDropOk() (*float64, bool) {
	if o == nil || o.PacketsOutDrop == nil {
		return nil, false
	}
	return o.PacketsOutDrop, true
}

// HasPacketsOutDrop returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasPacketsOutDrop() bool {
	return o != nil && o.PacketsOutDrop != nil
}

// SetPacketsOutDrop gets a reference to the given float64 and assigns it to the PacketsOutDrop field.
func (o *DeviceDetailsDataAttributes) SetPacketsOutDrop(v float64) {
	o.PacketsOutDrop = &v
}

// GetPacketsOutError returns the PacketsOutError field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetPacketsOutError() float64 {
	if o == nil || o.PacketsOutError == nil {
		var ret float64
		return ret
	}
	return *o.PacketsOutError
}

// GetPacketsOutErrorOk returns a tuple with the PacketsOutError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetPacketsOutErrorOk() (*float64, bool) {
	if o == nil || o.PacketsOutError == nil {
		return nil, false
	}
	return o.PacketsOutError, true
}

// HasPacketsOutError returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasPacketsOutError() bool {
	return o != nil && o.PacketsOutError != nil
}

// SetPacketsOutError gets a reference to the given float64 and assigns it to the PacketsOutError field.
func (o *DeviceDetailsDataAttributes) SetPacketsOutError(v float64) {
	o.PacketsOutError = &v
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
func (o *DeviceDetailsDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceDetailsDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetTcpOutSegs returns the TcpOutSegs field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpOutSegs() float64 {
	if o == nil || o.TcpOutSegs == nil {
		var ret float64
		return ret
	}
	return *o.TcpOutSegs
}

// GetTcpOutSegsOk returns a tuple with the TcpOutSegs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpOutSegsOk() (*float64, bool) {
	if o == nil || o.TcpOutSegs == nil {
		return nil, false
	}
	return o.TcpOutSegs, true
}

// HasTcpOutSegs returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpOutSegs() bool {
	return o != nil && o.TcpOutSegs != nil
}

// SetTcpOutSegs gets a reference to the given float64 and assigns it to the TcpOutSegs field.
func (o *DeviceDetailsDataAttributes) SetTcpOutSegs(v float64) {
	o.TcpOutSegs = &v
}

// GetTcpRetransSegs returns the TcpRetransSegs field value if set, zero value otherwise.
func (o *DeviceDetailsDataAttributes) GetTcpRetransSegs() float64 {
	if o == nil || o.TcpRetransSegs == nil {
		var ret float64
		return ret
	}
	return *o.TcpRetransSegs
}

// GetTcpRetransSegsOk returns a tuple with the TcpRetransSegs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDetailsDataAttributes) GetTcpRetransSegsOk() (*float64, bool) {
	if o == nil || o.TcpRetransSegs == nil {
		return nil, false
	}
	return o.TcpRetransSegs, true
}

// HasTcpRetransSegs returns a boolean if a field has been set.
func (o *DeviceDetailsDataAttributes) HasTcpRetransSegs() bool {
	return o != nil && o.TcpRetransSegs != nil
}

// SetTcpRetransSegs gets a reference to the given float64 and assigns it to the TcpRetransSegs field.
func (o *DeviceDetailsDataAttributes) SetTcpRetransSegs(v float64) {
	o.TcpRetransSegs = &v
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
	if o.BatteryCycleCount != nil {
		toSerialize["battery_cycle_count"] = o.BatteryCycleCount
	}
	if o.BatteryMaxCapacityPct != nil {
		toSerialize["battery_max_capacity_pct"] = o.BatteryMaxCapacityPct
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
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}
	if o.KernelName != nil {
		toSerialize["kernel_name"] = o.KernelName
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
	if o.PacketsInDrop != nil {
		toSerialize["packets_in_drop"] = o.PacketsInDrop
	}
	if o.PacketsInError != nil {
		toSerialize["packets_in_error"] = o.PacketsInError
	}
	if o.PacketsOutDrop != nil {
		toSerialize["packets_out_drop"] = o.PacketsOutDrop
	}
	if o.PacketsOutError != nil {
		toSerialize["packets_out_error"] = o.PacketsOutError
	}
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TcpOutSegs != nil {
		toSerialize["tcp_out_segs"] = o.TcpOutSegs
	}
	if o.TcpRetransSegs != nil {
		toSerialize["tcp_retrans_segs"] = o.TcpRetransSegs
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
		AgentKey              *string  `json:"agent_key,omitempty"`
		AgentVersion          *string  `json:"agent_version,omitempty"`
		BatteryChargePct      *int64   `json:"battery_charge_pct,omitempty"`
		BatteryChargeRate     *int64   `json:"battery_charge_rate,omitempty"`
		BatteryCycleCount     *int64   `json:"battery_cycle_count,omitempty"`
		BatteryMaxCapacityPct *int64   `json:"battery_max_capacity_pct,omitempty"`
		CpuCores              *int64   `json:"cpu_cores,omitempty"`
		CpuLogicalProcessors  *int64   `json:"cpu_logical_processors,omitempty"`
		CpuModel              *string  `json:"cpu_model,omitempty"`
		CpuUsage              *float64 `json:"cpu_usage,omitempty"`
		DiskUsage             *float64 `json:"disk_usage,omitempty"`
		IpAddress             *string  `json:"ip_address,omitempty"`
		Issues                []string `json:"issues,omitempty"`
		KernelName            *string  `json:"kernel_name,omitempty"`
		LastSeen              *string  `json:"last_seen,omitempty"`
		Manufacturer          *string  `json:"manufacturer,omitempty"`
		MemUsage              *float64 `json:"mem_usage,omitempty"`
		MemoryTotalKb         *int64   `json:"memory_total_kb,omitempty"`
		ModelName             *string  `json:"model_name,omitempty"`
		ModelNumber           *string  `json:"model_number,omitempty"`
		Os                    *string  `json:"os,omitempty"`
		OsVersion             *string  `json:"os_version,omitempty"`
		PacketsInDrop         *float64 `json:"packets_in_drop,omitempty"`
		PacketsInError        *float64 `json:"packets_in_error,omitempty"`
		PacketsOutDrop        *float64 `json:"packets_out_drop,omitempty"`
		PacketsOutError       *float64 `json:"packets_out_error,omitempty"`
		ResourceId            *string  `json:"resource_id,omitempty"`
		SerialNumber          *string  `json:"serial_number,omitempty"`
		Status                *string  `json:"status,omitempty"`
		TcpOutSegs            *float64 `json:"tcp_out_segs,omitempty"`
		TcpRetransSegs        *float64 `json:"tcp_retrans_segs,omitempty"`
		Type                  *string  `json:"type,omitempty"`
		Uptime                *float64 `json:"uptime,omitempty"`
		WlanBssid             *string  `json:"wlan_bssid,omitempty"`
		WlanRssi              *float64 `json:"wlan_rssi,omitempty"`
		WlanSsid              *string  `json:"wlan_ssid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_key", "agent_version", "battery_charge_pct", "battery_charge_rate", "battery_cycle_count", "battery_max_capacity_pct", "cpu_cores", "cpu_logical_processors", "cpu_model", "cpu_usage", "disk_usage", "ip_address", "issues", "kernel_name", "last_seen", "manufacturer", "mem_usage", "memory_total_kb", "model_name", "model_number", "os", "os_version", "packets_in_drop", "packets_in_error", "packets_out_drop", "packets_out_error", "resource_id", "serial_number", "status", "tcp_out_segs", "tcp_retrans_segs", "type", "uptime", "wlan_bssid", "wlan_rssi", "wlan_ssid"})
	} else {
		return err
	}
	o.AgentKey = all.AgentKey
	o.AgentVersion = all.AgentVersion
	o.BatteryChargePct = all.BatteryChargePct
	o.BatteryChargeRate = all.BatteryChargeRate
	o.BatteryCycleCount = all.BatteryCycleCount
	o.BatteryMaxCapacityPct = all.BatteryMaxCapacityPct
	o.CpuCores = all.CpuCores
	o.CpuLogicalProcessors = all.CpuLogicalProcessors
	o.CpuModel = all.CpuModel
	o.CpuUsage = all.CpuUsage
	o.DiskUsage = all.DiskUsage
	o.IpAddress = all.IpAddress
	o.Issues = all.Issues
	o.KernelName = all.KernelName
	o.LastSeen = all.LastSeen
	o.Manufacturer = all.Manufacturer
	o.MemUsage = all.MemUsage
	o.MemoryTotalKb = all.MemoryTotalKb
	o.ModelName = all.ModelName
	o.ModelNumber = all.ModelNumber
	o.Os = all.Os
	o.OsVersion = all.OsVersion
	o.PacketsInDrop = all.PacketsInDrop
	o.PacketsInError = all.PacketsInError
	o.PacketsOutDrop = all.PacketsOutDrop
	o.PacketsOutError = all.PacketsOutError
	o.ResourceId = all.ResourceId
	o.SerialNumber = all.SerialNumber
	o.Status = all.Status
	o.TcpOutSegs = all.TcpOutSegs
	o.TcpRetransSegs = all.TcpRetransSegs
	o.Type = all.Type
	o.Uptime = all.Uptime
	o.WlanBssid = all.WlanBssid
	o.WlanRssi = all.WlanRssi
	o.WlanSsid = all.WlanSsid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
