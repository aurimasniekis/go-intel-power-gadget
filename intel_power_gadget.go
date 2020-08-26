package intel_power_gadget

/*
#cgo CFLAGS: -g -Wall -I/Library/Frameworks/IntelPowerGadget.framework/Headers
#cgo LDFLAGS: -framework IntelPowerGadget -F/Library/Frameworks
#include "PowerGadgetLib.h"
*/
import "C"
import (
	"math"
	"time"
)

type Frequency float32
type FrequencyStat struct {
	Mean Frequency `json:"mean"`
	Min  Frequency `json:"min"`
	Max  Frequency `json:"max"`
}

type Watts float32
type Joules float32

type Power struct {
	Watts  Watts  `json:"watts"`
	Joules Joules `json:"joules"`
}

type Temperature float32
type TemperatureStat struct {
	Mean Temperature `json:"mean"`
	Min  Temperature `json:"min"`
	Max  Temperature `json:"max"`
}

type Utilization float32

type IntelPowerGadgetPackage struct {
	PackageNo    int `json:"package_no"`
	PackageCores int `json:"package_cores"`

	IaBaseFrequency Frequency `json:"ia_base_frequency"`
	IaMaxFrequency  Frequency `json:"ia_max_frequency"`
	GtMaxFrequency  Frequency `json:"gt_max_frequency"`

	PackageTDP     Watts       `json:"package_tdp"`
	MaxTemperature Temperature `json:"max_temperature"`

	GtAvailable             bool `json:"gt_available"`
	IaEnergyAvailable       bool `json:"ia_energy_available"`
	DramEnergyAvailable     bool `json:"dram_energy_available"`
	PlatformEnergyAvailable bool `json:"platform_energy_available"`
}

type IntelPowerGadgetSample struct {
	Pkg *IntelPowerGadgetPackage `json:"pkg"`

	TimestampStart time.Time     `json:"timestamp_start"`
	TimestampEnd   time.Time     `json:"timestamp_end"`
	Interval       time.Duration `json:"interval"`

	IaFrequency        FrequencyStat   `json:"ia_frequency"`
	IaFrequencyRequest FrequencyStat   `json:"ia_frequency_request"`
	IaPower            Power           `json:"ia_power"`
	IaTemperature      TemperatureStat `json:"ia_temperature"`
	IaUtilization      Utilization     `json:"ia_utilization"`

	IaCoreFrequency        map[int]FrequencyStat   `json:"ia_core_frequency"`
	IaCoreFrequencyRequest map[int]FrequencyStat   `json:"ia_core_frequency_request"`
	IaCoreTemperature      map[int]TemperatureStat `json:"ia_core_temperature"`
	IaCoreUtilization      map[int]Utilization     `json:"ia_core_utilization"`

	GtFrequency        Frequency   `json:"gt_frequency"`
	GtFrequencyRequest Frequency   `json:"gt_frequency_request"`
	GtUtilization      Utilization `json:"gt_utilization"`

	PackagePower Power `json:"package_power"`

	PlatformPower Power `json:"platform_power"`
	DramPower     Power `json:"dram_power"`

	PackageTemperature Temperature `json:"package_temperature"`
	Tdp                Watts       `json:"tdp"`
}

func Initialize() bool {
	return bool(C.PG_Initialize())
}

func Shutdown() bool {
	return bool(C.PG_Shutdown())
}

func NumPackages() int {
	var num C.int

	C.PG_GetNumPackages(&num)

	return int(num)
}

func NumCores(iPackage int) int {
	var result C.int

	C.PG_GetNumCores(C.int(iPackage), &result)

	return int(result)
}

func GetPackage(index int) *IntelPowerGadgetPackage {
	result := new(IntelPowerGadgetPackage)

	var numCores C.int
	var iaBaseFrequency C.double
	var iaMaxFrequency C.double
	var gtMaxFrequency C.double
	var packageTDP C.double
	var maxTemperature C.uchar
	var gtAvailable C.bool
	var iaEnergyAvailable C.bool
	var dramEnergyAvailable C.bool
	var platformEnergyAvailable C.bool

	result.PackageNo = index

	C.PG_GetNumCores(C.int(index), &numCores)
	result.PackageCores = int(numCores)

	C.PG_IsGTAvailable(C.int(index), &gtAvailable)
	result.GtAvailable = bool(gtAvailable)

	C.PG_IsIAEnergyAvailable(C.int(index), &iaEnergyAvailable)
	result.IaEnergyAvailable = bool(iaEnergyAvailable)

	C.PG_IsDRAMEnergyAvailable(C.int(index), &dramEnergyAvailable)
	result.DramEnergyAvailable = bool(dramEnergyAvailable)

	C.PG_IsPlatformEnergyAvailable(C.int(index), &platformEnergyAvailable)
	result.PlatformEnergyAvailable = bool(platformEnergyAvailable)

	C.PG_GetMaxTemperature(C.int(index), &maxTemperature)
	result.MaxTemperature = Temperature(maxTemperature)

	C.PG_GetIABaseFrequency(C.int(index), &iaBaseFrequency)
	result.IaBaseFrequency = Frequency(iaBaseFrequency)

	C.PG_GetIAMaxFrequency(C.int(index), &iaMaxFrequency)
	result.IaMaxFrequency = Frequency(iaMaxFrequency)

	if result.GtAvailable {
		C.PG_GetGTMaxFrequency(C.int(index), &gtMaxFrequency)
		result.GtMaxFrequency = Frequency(gtMaxFrequency)
	}

	C.PG_GetTDP(C.int(index), &packageTDP)
	result.PackageTDP = Watts(packageTDP)

	return result
}

func GetPackages() map[int]*IntelPowerGadgetPackage {
	result := make(map[int]*IntelPowerGadgetPackage)

	for i := 0; i < NumPackages(); i++ {
		result[i] = GetPackage(i)
	}

	return result
}

type SampleId uint64

func StartSampling(pkg *IntelPowerGadgetPackage) SampleId {
	var sampleId C.ulonglong

	C.PG_ReadSample(C.int(pkg.PackageNo), &sampleId)

	return SampleId(sampleId)
}

func FinishSampling(sampleId SampleId, pkg *IntelPowerGadgetPackage) *IntelPowerGadgetSample {
	var start, end C.ulonglong

	start = C.ulonglong(sampleId)
	C.PG_ReadSample(C.int(pkg.PackageNo), &end)

	result := new(IntelPowerGadgetSample)
	result.Pkg = pkg

	var timeStartSeconds C.uint
	var timeStartNanoSeconds C.uint
	var timeEndSeconds C.uint
	var timeEndNanoSeconds C.uint
	var interval C.double
	var mean C.double
	var min C.double
	var max C.double

	C.PGSample_GetTime(start, &timeStartSeconds, &timeStartNanoSeconds)
	result.TimestampStart = time.Unix(int64(timeStartSeconds), int64(timeStartNanoSeconds))

	C.PGSample_GetTime(end, &timeEndSeconds, &timeEndNanoSeconds)
	result.TimestampEnd = time.Unix(int64(timeEndSeconds), int64(timeEndNanoSeconds))

	C.PGSample_GetTimeInterval(start, end, &interval)
	result.Interval = time.Duration(math.Round(float64(interval))) * time.Second

	C.PGSample_GetIAFrequency(start, end, &mean, &min, &max)
	result.IaFrequency.Mean = Frequency(mean)
	result.IaFrequency.Min = Frequency(min)
	result.IaFrequency.Max = Frequency(max)

	C.PGSample_GetIAFrequencyRequest(end, &mean, &min, &max)
	result.IaFrequencyRequest.Mean = Frequency(mean)
	result.IaFrequencyRequest.Min = Frequency(min)
	result.IaFrequencyRequest.Max = Frequency(max)

	C.PGSample_GetIAPower(start, end, &mean, &min)
	result.IaPower = Power{
		Watts:  Watts(mean),
		Joules: Joules(min),
	}

	C.PGSample_GetIATemperature(end, &mean, &min, &max)
	result.IaTemperature = TemperatureStat{
		Mean: Temperature(mean),
		Min:  Temperature(min),
		Max:  Temperature(max),
	}

	C.PGSample_GetIAUtilization(start, end, &mean)
	result.IaUtilization = Utilization(mean)

	if pkg.GtAvailable {
		C.PGSample_GetGTFrequency(end, &mean)
		result.GtFrequency = Frequency(mean)

		C.PGSample_GetGTFrequencyRequest(end, &mean)
		result.GtFrequencyRequest = Frequency(mean)

		C.PGSample_GetGTUtilization(end, &mean)
		result.GtUtilization = Utilization(mean)

	}

	C.PGSample_GetPackagePower(start, end, &mean, &min)
	result.PackagePower = Power{
		Watts:  Watts(mean),
		Joules: Joules(min),
	}

	C.PGSample_GetPlatformPower(start, end, &mean, &min)
	result.PlatformPower = Power{
		Watts:  Watts(mean),
		Joules: Joules(min),
	}

	C.PGSample_GetDRAMPower(start, end, &mean, &min)
	result.DramPower = Power{
		Watts:  Watts(mean),
		Joules: Joules(min),
	}

	C.PGSample_GetPackageTemperature(end, &mean)
	result.PackageTemperature = Temperature(mean)

	C.PGSample_GetTDP(end, &mean)
	result.Tdp = Watts(mean)

	result.IaCoreFrequency = make(map[int]FrequencyStat)
	result.IaCoreFrequencyRequest = make(map[int]FrequencyStat)
	result.IaCoreTemperature = make(map[int]TemperatureStat)
	result.IaCoreUtilization = make(map[int]Utilization)
	for core := 0; core < pkg.PackageCores; core++ {
		var mean C.double
		var min C.double
		var max C.double

		coreIaFrequency := FrequencyStat{}
		coreIaFrequencyRequest := FrequencyStat{}
		coreTemperature := TemperatureStat{}

		C.PGSample_GetIACoreFrequency(start, end, C.int(core), &mean, &min, &max)
		coreIaFrequency.Mean = Frequency(mean)
		coreIaFrequency.Min = Frequency(min)
		coreIaFrequency.Max = Frequency(max)

		C.PGSample_GetIACoreFrequencyRequest(end, C.int(core), &mean, &min, &max)
		coreIaFrequencyRequest.Mean = Frequency(mean)
		coreIaFrequencyRequest.Min = Frequency(min)
		coreIaFrequencyRequest.Max = Frequency(max)

		C.PGSample_GetIACoreTemperature(end, C.int(core), &mean, &min, &max)
		coreTemperature.Mean = Temperature(mean)
		coreTemperature.Min = Temperature(min)
		coreTemperature.Max = Temperature(max)

		C.PGSample_GetIACoreUtilization(start, end, C.int(core), &mean)
		result.IaCoreUtilization[core] = Utilization(mean)

		result.IaCoreFrequency[core] = coreIaFrequency
		result.IaCoreFrequencyRequest[core] = coreIaFrequencyRequest
		result.IaCoreTemperature[core] = coreTemperature
	}

	C.PGSample_Release(start)
	C.PGSample_Release(end)

	return result
}
