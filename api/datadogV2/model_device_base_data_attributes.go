// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceBaseDataAttributes Common health and identity attributes shared by every End User Device Monitoring device record.
type DeviceBaseDataAttributes struct {
	// Public key of the Datadog Agent installed on the device.
	AgentKey *string `json:"agent_key,omitempty"`
	// Version of the Datadog Agent installed on the device.
	AgentVersion *string `json:"agent_version,omitempty"`
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
	// Datadog resource identifier for the device.
	ResourceId *string `json:"resource_id,omitempty"`
	// Serial number assigned to the device by its manufacturer.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Health status of the device computed from its issues and recent telemetry.
	Status *string `json:"status,omitempty"`
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

// NewDeviceBaseDataAttributes instantiates a new DeviceBaseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeviceBaseDataAttributes() *DeviceBaseDataAttributes {
	this := DeviceBaseDataAttributes{}
	return &this
}

// NewDeviceBaseDataAttributesWithDefaults instantiates a new DeviceBaseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeviceBaseDataAttributesWithDefaults() *DeviceBaseDataAttributes {
	this := DeviceBaseDataAttributes{}
	return &this
}

// GetAgentKey returns the AgentKey field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetAgentKey() string {
	if o == nil || o.AgentKey == nil {
		var ret string
		return ret
	}
	return *o.AgentKey
}

// GetAgentKeyOk returns a tuple with the AgentKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetAgentKeyOk() (*string, bool) {
	if o == nil || o.AgentKey == nil {
		return nil, false
	}
	return o.AgentKey, true
}

// HasAgentKey returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasAgentKey() bool {
	return o != nil && o.AgentKey != nil
}

// SetAgentKey gets a reference to the given string and assigns it to the AgentKey field.
func (o *DeviceBaseDataAttributes) SetAgentKey(v string) {
	o.AgentKey = &v
}

// GetAgentVersion returns the AgentVersion field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetAgentVersion() string {
	if o == nil || o.AgentVersion == nil {
		var ret string
		return ret
	}
	return *o.AgentVersion
}

// GetAgentVersionOk returns a tuple with the AgentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetAgentVersionOk() (*string, bool) {
	if o == nil || o.AgentVersion == nil {
		return nil, false
	}
	return o.AgentVersion, true
}

// HasAgentVersion returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasAgentVersion() bool {
	return o != nil && o.AgentVersion != nil
}

// SetAgentVersion gets a reference to the given string and assigns it to the AgentVersion field.
func (o *DeviceBaseDataAttributes) SetAgentVersion(v string) {
	o.AgentVersion = &v
}

// GetBatteryMaxCapacityPct returns the BatteryMaxCapacityPct field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetBatteryMaxCapacityPct() int64 {
	if o == nil || o.BatteryMaxCapacityPct == nil {
		var ret int64
		return ret
	}
	return *o.BatteryMaxCapacityPct
}

// GetBatteryMaxCapacityPctOk returns a tuple with the BatteryMaxCapacityPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetBatteryMaxCapacityPctOk() (*int64, bool) {
	if o == nil || o.BatteryMaxCapacityPct == nil {
		return nil, false
	}
	return o.BatteryMaxCapacityPct, true
}

// HasBatteryMaxCapacityPct returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasBatteryMaxCapacityPct() bool {
	return o != nil && o.BatteryMaxCapacityPct != nil
}

// SetBatteryMaxCapacityPct gets a reference to the given int64 and assigns it to the BatteryMaxCapacityPct field.
func (o *DeviceBaseDataAttributes) SetBatteryMaxCapacityPct(v int64) {
	o.BatteryMaxCapacityPct = &v
}

// GetCpuCores returns the CpuCores field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetCpuCores() int64 {
	if o == nil || o.CpuCores == nil {
		var ret int64
		return ret
	}
	return *o.CpuCores
}

// GetCpuCoresOk returns a tuple with the CpuCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetCpuCoresOk() (*int64, bool) {
	if o == nil || o.CpuCores == nil {
		return nil, false
	}
	return o.CpuCores, true
}

// HasCpuCores returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasCpuCores() bool {
	return o != nil && o.CpuCores != nil
}

// SetCpuCores gets a reference to the given int64 and assigns it to the CpuCores field.
func (o *DeviceBaseDataAttributes) SetCpuCores(v int64) {
	o.CpuCores = &v
}

// GetCpuLogicalProcessors returns the CpuLogicalProcessors field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetCpuLogicalProcessors() int64 {
	if o == nil || o.CpuLogicalProcessors == nil {
		var ret int64
		return ret
	}
	return *o.CpuLogicalProcessors
}

// GetCpuLogicalProcessorsOk returns a tuple with the CpuLogicalProcessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetCpuLogicalProcessorsOk() (*int64, bool) {
	if o == nil || o.CpuLogicalProcessors == nil {
		return nil, false
	}
	return o.CpuLogicalProcessors, true
}

// HasCpuLogicalProcessors returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasCpuLogicalProcessors() bool {
	return o != nil && o.CpuLogicalProcessors != nil
}

// SetCpuLogicalProcessors gets a reference to the given int64 and assigns it to the CpuLogicalProcessors field.
func (o *DeviceBaseDataAttributes) SetCpuLogicalProcessors(v int64) {
	o.CpuLogicalProcessors = &v
}

// GetCpuModel returns the CpuModel field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetCpuModel() string {
	if o == nil || o.CpuModel == nil {
		var ret string
		return ret
	}
	return *o.CpuModel
}

// GetCpuModelOk returns a tuple with the CpuModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetCpuModelOk() (*string, bool) {
	if o == nil || o.CpuModel == nil {
		return nil, false
	}
	return o.CpuModel, true
}

// HasCpuModel returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasCpuModel() bool {
	return o != nil && o.CpuModel != nil
}

// SetCpuModel gets a reference to the given string and assigns it to the CpuModel field.
func (o *DeviceBaseDataAttributes) SetCpuModel(v string) {
	o.CpuModel = &v
}

// GetCpuUsage returns the CpuUsage field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetCpuUsage() float64 {
	if o == nil || o.CpuUsage == nil {
		var ret float64
		return ret
	}
	return *o.CpuUsage
}

// GetCpuUsageOk returns a tuple with the CpuUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetCpuUsageOk() (*float64, bool) {
	if o == nil || o.CpuUsage == nil {
		return nil, false
	}
	return o.CpuUsage, true
}

// HasCpuUsage returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasCpuUsage() bool {
	return o != nil && o.CpuUsage != nil
}

// SetCpuUsage gets a reference to the given float64 and assigns it to the CpuUsage field.
func (o *DeviceBaseDataAttributes) SetCpuUsage(v float64) {
	o.CpuUsage = &v
}

// GetDiskUsage returns the DiskUsage field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetDiskUsage() float64 {
	if o == nil || o.DiskUsage == nil {
		var ret float64
		return ret
	}
	return *o.DiskUsage
}

// GetDiskUsageOk returns a tuple with the DiskUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetDiskUsageOk() (*float64, bool) {
	if o == nil || o.DiskUsage == nil {
		return nil, false
	}
	return o.DiskUsage, true
}

// HasDiskUsage returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasDiskUsage() bool {
	return o != nil && o.DiskUsage != nil
}

// SetDiskUsage gets a reference to the given float64 and assigns it to the DiskUsage field.
func (o *DeviceBaseDataAttributes) SetDiskUsage(v float64) {
	o.DiskUsage = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasIpAddress() bool {
	return o != nil && o.IpAddress != nil
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *DeviceBaseDataAttributes) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetIssues() []string {
	if o == nil || o.Issues == nil {
		var ret []string
		return ret
	}
	return o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetIssuesOk() (*[]string, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return &o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasIssues() bool {
	return o != nil && o.Issues != nil
}

// SetIssues gets a reference to the given []string and assigns it to the Issues field.
func (o *DeviceBaseDataAttributes) SetIssues(v []string) {
	o.Issues = v
}

// GetKernelName returns the KernelName field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetKernelName() string {
	if o == nil || o.KernelName == nil {
		var ret string
		return ret
	}
	return *o.KernelName
}

// GetKernelNameOk returns a tuple with the KernelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetKernelNameOk() (*string, bool) {
	if o == nil || o.KernelName == nil {
		return nil, false
	}
	return o.KernelName, true
}

// HasKernelName returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasKernelName() bool {
	return o != nil && o.KernelName != nil
}

// SetKernelName gets a reference to the given string and assigns it to the KernelName field.
func (o *DeviceBaseDataAttributes) SetKernelName(v string) {
	o.KernelName = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetLastSeen() string {
	if o == nil || o.LastSeen == nil {
		var ret string
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetLastSeenOk() (*string, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given string and assigns it to the LastSeen field.
func (o *DeviceBaseDataAttributes) SetLastSeen(v string) {
	o.LastSeen = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetManufacturer() string {
	if o == nil || o.Manufacturer == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetManufacturerOk() (*string, bool) {
	if o == nil || o.Manufacturer == nil {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasManufacturer() bool {
	return o != nil && o.Manufacturer != nil
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *DeviceBaseDataAttributes) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetMemUsage returns the MemUsage field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetMemUsage() float64 {
	if o == nil || o.MemUsage == nil {
		var ret float64
		return ret
	}
	return *o.MemUsage
}

// GetMemUsageOk returns a tuple with the MemUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetMemUsageOk() (*float64, bool) {
	if o == nil || o.MemUsage == nil {
		return nil, false
	}
	return o.MemUsage, true
}

// HasMemUsage returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasMemUsage() bool {
	return o != nil && o.MemUsage != nil
}

// SetMemUsage gets a reference to the given float64 and assigns it to the MemUsage field.
func (o *DeviceBaseDataAttributes) SetMemUsage(v float64) {
	o.MemUsage = &v
}

// GetMemoryTotalKb returns the MemoryTotalKb field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetMemoryTotalKb() int64 {
	if o == nil || o.MemoryTotalKb == nil {
		var ret int64
		return ret
	}
	return *o.MemoryTotalKb
}

// GetMemoryTotalKbOk returns a tuple with the MemoryTotalKb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetMemoryTotalKbOk() (*int64, bool) {
	if o == nil || o.MemoryTotalKb == nil {
		return nil, false
	}
	return o.MemoryTotalKb, true
}

// HasMemoryTotalKb returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasMemoryTotalKb() bool {
	return o != nil && o.MemoryTotalKb != nil
}

// SetMemoryTotalKb gets a reference to the given int64 and assigns it to the MemoryTotalKb field.
func (o *DeviceBaseDataAttributes) SetMemoryTotalKb(v int64) {
	o.MemoryTotalKb = &v
}

// GetModelName returns the ModelName field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetModelName() string {
	if o == nil || o.ModelName == nil {
		var ret string
		return ret
	}
	return *o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetModelNameOk() (*string, bool) {
	if o == nil || o.ModelName == nil {
		return nil, false
	}
	return o.ModelName, true
}

// HasModelName returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasModelName() bool {
	return o != nil && o.ModelName != nil
}

// SetModelName gets a reference to the given string and assigns it to the ModelName field.
func (o *DeviceBaseDataAttributes) SetModelName(v string) {
	o.ModelName = &v
}

// GetModelNumber returns the ModelNumber field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetModelNumber() string {
	if o == nil || o.ModelNumber == nil {
		var ret string
		return ret
	}
	return *o.ModelNumber
}

// GetModelNumberOk returns a tuple with the ModelNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetModelNumberOk() (*string, bool) {
	if o == nil || o.ModelNumber == nil {
		return nil, false
	}
	return o.ModelNumber, true
}

// HasModelNumber returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasModelNumber() bool {
	return o != nil && o.ModelNumber != nil
}

// SetModelNumber gets a reference to the given string and assigns it to the ModelNumber field.
func (o *DeviceBaseDataAttributes) SetModelNumber(v string) {
	o.ModelNumber = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasOs() bool {
	return o != nil && o.Os != nil
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *DeviceBaseDataAttributes) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasOsVersion() bool {
	return o != nil && o.OsVersion != nil
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *DeviceBaseDataAttributes) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetResourceId() string {
	if o == nil || o.ResourceId == nil {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil || o.ResourceId == nil {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasResourceId() bool {
	return o != nil && o.ResourceId != nil
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *DeviceBaseDataAttributes) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasSerialNumber() bool {
	return o != nil && o.SerialNumber != nil
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *DeviceBaseDataAttributes) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceBaseDataAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeviceBaseDataAttributes) SetType(v string) {
	o.Type = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetUptime() float64 {
	if o == nil || o.Uptime == nil {
		var ret float64
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetUptimeOk() (*float64, bool) {
	if o == nil || o.Uptime == nil {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasUptime() bool {
	return o != nil && o.Uptime != nil
}

// SetUptime gets a reference to the given float64 and assigns it to the Uptime field.
func (o *DeviceBaseDataAttributes) SetUptime(v float64) {
	o.Uptime = &v
}

// GetWlanBssid returns the WlanBssid field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetWlanBssid() string {
	if o == nil || o.WlanBssid == nil {
		var ret string
		return ret
	}
	return *o.WlanBssid
}

// GetWlanBssidOk returns a tuple with the WlanBssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetWlanBssidOk() (*string, bool) {
	if o == nil || o.WlanBssid == nil {
		return nil, false
	}
	return o.WlanBssid, true
}

// HasWlanBssid returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasWlanBssid() bool {
	return o != nil && o.WlanBssid != nil
}

// SetWlanBssid gets a reference to the given string and assigns it to the WlanBssid field.
func (o *DeviceBaseDataAttributes) SetWlanBssid(v string) {
	o.WlanBssid = &v
}

// GetWlanRssi returns the WlanRssi field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetWlanRssi() float64 {
	if o == nil || o.WlanRssi == nil {
		var ret float64
		return ret
	}
	return *o.WlanRssi
}

// GetWlanRssiOk returns a tuple with the WlanRssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetWlanRssiOk() (*float64, bool) {
	if o == nil || o.WlanRssi == nil {
		return nil, false
	}
	return o.WlanRssi, true
}

// HasWlanRssi returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasWlanRssi() bool {
	return o != nil && o.WlanRssi != nil
}

// SetWlanRssi gets a reference to the given float64 and assigns it to the WlanRssi field.
func (o *DeviceBaseDataAttributes) SetWlanRssi(v float64) {
	o.WlanRssi = &v
}

// GetWlanSsid returns the WlanSsid field value if set, zero value otherwise.
func (o *DeviceBaseDataAttributes) GetWlanSsid() string {
	if o == nil || o.WlanSsid == nil {
		var ret string
		return ret
	}
	return *o.WlanSsid
}

// GetWlanSsidOk returns a tuple with the WlanSsid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceBaseDataAttributes) GetWlanSsidOk() (*string, bool) {
	if o == nil || o.WlanSsid == nil {
		return nil, false
	}
	return o.WlanSsid, true
}

// HasWlanSsid returns a boolean if a field has been set.
func (o *DeviceBaseDataAttributes) HasWlanSsid() bool {
	return o != nil && o.WlanSsid != nil
}

// SetWlanSsid gets a reference to the given string and assigns it to the WlanSsid field.
func (o *DeviceBaseDataAttributes) SetWlanSsid(v string) {
	o.WlanSsid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeviceBaseDataAttributes) MarshalJSON() ([]byte, error) {
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
	if o.ResourceId != nil {
		toSerialize["resource_id"] = o.ResourceId
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
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
func (o *DeviceBaseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentKey              *string  `json:"agent_key,omitempty"`
		AgentVersion          *string  `json:"agent_version,omitempty"`
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
		ResourceId            *string  `json:"resource_id,omitempty"`
		SerialNumber          *string  `json:"serial_number,omitempty"`
		Status                *string  `json:"status,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_key", "agent_version", "battery_max_capacity_pct", "cpu_cores", "cpu_logical_processors", "cpu_model", "cpu_usage", "disk_usage", "ip_address", "issues", "kernel_name", "last_seen", "manufacturer", "mem_usage", "memory_total_kb", "model_name", "model_number", "os", "os_version", "resource_id", "serial_number", "status", "type", "uptime", "wlan_bssid", "wlan_rssi", "wlan_ssid"})
	} else {
		return err
	}
	o.AgentKey = all.AgentKey
	o.AgentVersion = all.AgentVersion
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
	o.ResourceId = all.ResourceId
	o.SerialNumber = all.SerialNumber
	o.Status = all.Status
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
