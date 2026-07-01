package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fehmicorp/agent/windows/debug/logger"
	"github.com/fehmicorp/agent/windows/services/defender"
	"github.com/fehmicorp/agent/windows/services/firewall"
	"github.com/fehmicorp/agent/windows/services/metric"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// var (
// 	prevNetState      map[string]net.IOCountersStat
// 	lastNetworkSample time.Time
// )

type UsageStats struct {
	CPU     string `json:"cpu"`
	RAM     string `json:"ram"`
	Disk    string `json:"disk"`
	Network string `json:"network"`
}

func (a *App) fetchAndEmit() {
	cpuStr := metric.GetCPUUsage()
	ramStr := metric.GetRAMUsage()
	diskStr := metric.GetDiskUsage("C:")
	_, rxStr, err := metric.GetNetwork()
	netStr := "↓ 0B"
	if err == nil {
		netStr = fmt.Sprintf("↓ %s", rxStr)
	} else {
		logger.Logger.Log("ERROR", "Network monitoring metrics update dropped: "+err.Error())
	}

	runtime.EventsEmit(a.ctx, "metrics_update", UsageStats{
		CPU:     cpuStr,
		RAM:     ramStr,
		Disk:    diskStr,
		Network: netStr,
	})
}

func (a *App) StartMetric(intervalMs int) {
	a.fetchAndEmit()
	ticker := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-a.ctx.Done():
			return
		case <-ticker.C:
			a.fetchAndEmit()
		}
	}
}

type SecurityInfo struct {
	WindowsDefender bool `json:"windowsDefender"`
	Firewall        bool `json:"firewall"`
	TPM             bool `json:"tpm"`
	BitLocker       bool `json:"bitLocker"`
}

func GetSecurityStatus() (SecurityInfo, error) {
	var errSummary []string
	fw, err := firewall.GetFirewallStatus()
	fwStatus := false
	if err != nil {
		errSummary = append(errSummary, "firewall: "+err.Error())
	} else if fw != nil {
		fwStatus = fw.Enabled
	}
	def, err := defender.GetDefenderStatus()
	defStatus := false
	tpmStatus := false
	if err != nil {
		errSummary = append(errSummary, "defender: "+err.Error())
	} else if def != nil {
		defStatus = def.AntivirusEnabled &&
			def.RealTimeProtection &&
			def.DefenderServiceState == "Running"
		tpmStatus = def.TamperProtection
	}
	bitlockerStatus := false
	blVolumes, err := GetBitLockerStatus()
	if err != nil {
		errSummary = append(errSummary, "bitlocker: "+err.Error())
	} else {
		for _, vol := range blVolumes {
			if strings.EqualFold(vol.Drive, "C:") &&
				vol.Encrypted &&
				vol.ProtectionEnabled {

				bitlockerStatus = true
				break
			}
		}
	}
	return SecurityInfo{
		WindowsDefender: defStatus,
		Firewall:        fwStatus,
		TPM:             tpmStatus,
		BitLocker:       bitlockerStatus,
	}, nil
}

type DeviceInfo struct {
	Hostname string `json:"hostname"`
	Domain   string `json:"domain"`
	User     string `json:"user"`
	OS       string `json:"os"`
}

func GetSystemDeviceInfo() (DeviceInfo, error) {
	var errSummary []string
	hostname, err := metric.GetHostname()
	if err != nil {
		errSummary = append(errSummary, "hostname: "+err.Error())
	}
	username, err := metric.GetCurrentUser()
	if err != nil {
		errSummary = append(errSummary, "username: "+err.Error())
	}
	osDisplay, err := metric.GetExactOSName()
	if err != nil {
		errSummary = append(errSummary, "os_name: "+err.Error())
	}
	domain, err := metric.GetWmiDomain()
	if err != nil {
		errSummary = append(errSummary, "domain: "+err.Error())
	}
	data := DeviceInfo{
		Hostname: hostname,
		Domain:   domain,
		User:     username,
		OS:       osDisplay,
	}
	if len(errSummary) > 0 {
		return data, fmt.Errorf("device info metrics completed with errors: [%s]", strings.Join(errSummary, "; "))
	}
	return data, nil
}

type AppInfo struct {
	Id          string `json:"id"`
	DeviceToken string `json:"deviceToken"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Build       string `json:"build"`
	Tag         string `json:"tag"`
	BuildType   string `json:"buildType"`
	Company     string `json:"company"`
	Website     string `json:"website"`
	Endpoint    string `json:"endpoint"`
	Description string `json:"description"`
}

const (
	Id          = "fehmi-endpoint-agent"
	Name        = "Fehmi Endpoint Agent"
	Version     = "v1.0.1"
	Build       = "20260630"
	Tag         = "win-agent"
	BuildType   = "stable"
	Company     = "Fehmi Corporation"
	Website     = "https://fehmicorp.in"
	Endpoint    = "https://edp.fehmicorp.in"
	Description = "Unified Endpoint Management Agent"
)

func GenerateAppSignature() string {
	appIdOnce.Do(func() {
		payload := fmt.Sprintf("name=%s|version=%s|build=%s|tag=%s|buildtype=%s",
			Name, Version, Build, Tag, BuildType)
		signatureSalt := "fehmi-endpoint-windows-agent"
		mac := hmac.New(sha256.New, []byte(signatureSalt))
		mac.Write([]byte(payload))
		fullSignature := hex.EncodeToString(mac.Sum(nil))
		shortSignature, _ := ShortSignature(fullSignature, Tag, "windows", time.Now())
		RegisterSignaturePair(fullSignature, shortSignature)
		cachedAppId = shortSignature
	})
	return cachedAppId
}

var (
	prevNetState      map[string]net.IOCountersStat
	lastNetworkSample time.Time
	sigMutex          sync.RWMutex
	fullToShortMap    = make(map[string]string)
	shortToFullMap    = make(map[string]string)
	cachedAppId       string
	cachedDeviceToken string
	deviceTokenOnce   sync.Once
	appIdOnce         sync.Once
)

func ShortSignature(fullSig string, tag string, osName string, timestamp time.Time) (string, bool) {
	currentYear := timestamp.Format("2006")
	prefix := "agnt"
	if strings.Contains(strings.ToLower(tag), "server") {
		prefix = "srvr"
	} else if strings.Contains(strings.ToLower(tag), "distribution") {
		prefix = "dstr"
	}
	osLower := strings.ToLower(osName)
	osCode := "u"
	if strings.Contains(osLower, "win") {
		osCode = "w"
	} else if strings.Contains(osLower, "mac") || strings.Contains(osLower, "darwin") {
		osCode = "m"
	} else if strings.Contains(osLower, "linux") {
		osCode = "l"
	}
	if len(fullSig) < 10 {
		return "", false
	}
	shortSig := fmt.Sprintf("%s_%s_%s%s", prefix, fullSig[:10], osCode, currentYear)
	return shortSig, true
}

func RegisterSignaturePair(full, short string) {
	sigMutex.Lock()
	defer sigMutex.Unlock()
	fullToShortMap[full] = short
	shortToFullMap[short] = full
}

func GetShortSignature(fullSig string) (string, bool) {
	sigMutex.RLock()
	defer sigMutex.RUnlock()
	short, exists := fullToShortMap[fullSig]
	return short, exists
}

func GetFullSignature(shortSig string) (string, bool) {
	sigMutex.RLock()
	defer sigMutex.RUnlock()
	full, exists := shortToFullMap[shortSig]
	return full, exists
}

func (a *App) GenerateDeviceSignature(tag, hostname, domain, osName, macAddress string) string {
	payload := fmt.Sprintf("tag=%s|hostname=%s|domain=%s|os=%s|mac=%s",
		tag, hostname, domain, osName, macAddress)
	signatureSalt := "fehmi-endpoint-windows-agent-device-" + tag
	mac := hmac.New(sha256.New, []byte(signatureSalt))
	mac.Write([]byte(payload))
	fullSignature := hex.EncodeToString(mac.Sum(nil))
	shortSignature, _ := ShortSignature(fullSignature, tag, osName, time.Now())
	RegisterSignaturePair(fullSignature, shortSignature)
	return shortSignature
}

func (a *App) InitializeFirstTimeSetup() string {
	deviceTokenOnce.Do(func() {
		device, err := GetSystemDeviceInfo()
		if err != nil {
			logger.Logger.Log("ERROR", "Failed to collect core system details for first time provisioning: "+err.Error())
		}
		_, macStr, err := metric.GetNetwork()
		if err != nil {
			macStr = "00:00:00:00:00:00"
		}
		deviceToken := a.GenerateDeviceSignature(Tag, device.Hostname, device.Domain, device.OS, macStr)
		logger.Logger.Log("INFO", "Application initialized successfully. Created Hardware Token: "+deviceToken)
		cachedDeviceToken = deviceToken
	})
	return cachedDeviceToken
}

func (a *App) AppInfoUpdate() AppInfo {
	return AppInfo{
		Id:          GenerateAppSignature(),
		DeviceToken: cachedDeviceToken,
		Name:        Name,
		Version:     Version,
		Build:       Build,
		Tag:         Tag,
		BuildType:   BuildType,
		Company:     Company,
		Website:     Website,
		Endpoint:    Endpoint,
		Description: Description,
	}
}

func (a *App) SystemInfoUpdate() DeviceInfo {
	data, err := GetSystemDeviceInfo()
	if err != nil {
		errStr := err.Error()
		if start := strings.Index(errStr, "["); start != -1 {
			if end := strings.LastIndex(errStr, "]"); end > start {
				rawErrors := errStr[start+1 : end]
				individualErrors := strings.Split(rawErrors, "; ")
				for _, singleErr := range individualErrors {
					if singleErr != "" {
						logger.Logger.Log("WARN", "Subsystem Failure Details -> "+singleErr)
					}
				}
			}
		}
	}
	return data
}

func (a *App) SecurityInfoUpdate() SecurityInfo {
	data, err := GetSecurityStatus()
	if err != nil {
		errStr := err.Error()
		if start := strings.Index(errStr, "["); start != -1 {
			if end := strings.LastIndex(errStr, "]"); end > start {
				rawErrors := errStr[start+1 : end]
				individualErrors := strings.Split(rawErrors, "; ")
				for _, singleErr := range individualErrors {
					if singleErr != "" {
						logger.Logger.Log("WARN", "Subsystem Failure Details -> "+singleErr)
					}
				}
			}
		}
	}
	return data
}
