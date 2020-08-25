# go-intel-power-gadget
A Golang macOS IntelPowerGadget framework binding

## Install

A simplay way using `go get`. Requires the IntelPowerGadget installed in system

```sh
$ go get github.com/aurimasniekis/go-intel-power-gadget
```

## Usage

A way to start sampling CPU information

```go
package main

import (
	"encoding/json"
	"fmt"
	ipg "github.com/aurimasniekis/go-intel-power-gadget"
	"time"
)

func main() {
	ipg.Initialize()

	var pkg *ipg.IntelPowerGadgetPackage

	// Get CPU 0
	pkg = ipg.GetPackage(0)

	sampleId := ipg.StartSampling(pkg)

	// Sleep for 1 second
	time.Sleep(time.Second)

	result := ipg.FinishSampling(sampleId, pkg)

	e, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))

	ipg.Shutdown()
}
```

Will produce

```json
{
  "Pkg": {
    "PackageNo": 0,
    "PackageCores": 16,
    "IaBaseFrequency": 3601,
    "IaMaxFrequency": 5000,
    "GtMaxFrequency": 0,
    "PackageTDP": 95,
    "MaxTemperature": 100,
    "GtAvailable": true,
    "IaEnergyAvailable": true,
    "DramEnergyAvailable": false,
    "PlatformEnergyAvailable": false
  },
  "TimestampStart": "2020-08-25T15:46:00.179948337+09:00",
  "TimestampEnd": "2020-08-25T15:46:01.184624025+09:00",
  "Interval": 1000000000,
  "IaFrequency": {
    "Mean": 4701.283,
    "Min": 1200,
    "Max": 5000
  },
  "IaFrequencyRequest": {
    "Mean": 4837.6885,
    "Min": 1300,
    "Max": 5000
  },
  "IaPower": {
    "Watts": 11.6643505,
    "Joules": 11.718811
  },
  "IaTemperature": {
    "Mean": 44.298996,
    "Min": 38,
    "Max": 49
  },
  "IaUtilization": 0.34665084,
  "IaCoreFrequency": {
    "0": {
      "Mean": 4540.4326,
      "Min": 1500,
      "Max": 5000
    },
    "1": {
      "Mean": 4513.7964,
      "Min": 4000,
      "Max": 4800
    },
    "10": {
      "Mean": 4922.033,
      "Min": 4600,
      "Max": 5000
    },
    "11": {
      "Mean": 4900.3325,
      "Min": 4900,
      "Max": 5000
    },
    "12": {
      "Mean": 4991.89,
      "Min": 4900,
      "Max": 5000
    },
    "13": {
      "Mean": 4905.8906,
      "Min": 4900,
      "Max": 5000
    },
    "14": {
      "Mean": 4931.931,
      "Min": 4900,
      "Max": 5000
    },
    "15": {
      "Mean": 4956.719,
      "Min": 4900,
      "Max": 5000
    },
    "2": {
      "Mean": 4807.717,
      "Min": 2500,
      "Max": 5000
    },
    "3": {
      "Mean": 4558.1123,
      "Min": 4200,
      "Max": 5000
    },
    "4": {
      "Mean": 4752.1895,
      "Min": 1600,
      "Max": 5000
    },
    "5": {
      "Mean": 4412.544,
      "Min": 3800,
      "Max": 5000
    },
    "6": {
      "Mean": 4690.9683,
      "Min": 1200,
      "Max": 5000
    },
    "7": {
      "Mean": 4761.6226,
      "Min": 3700,
      "Max": 5000
    },
    "8": {
      "Mean": 4840.3027,
      "Min": 3900,
      "Max": 5000
    },
    "9": {
      "Mean": 3865.2747,
      "Min": 3800,
      "Max": 3900
    }
  },
  "IaCoreFrequencyRequest": {
    "0": {
      "Mean": 4705.4546,
      "Min": 3400,
      "Max": 5000
    },
    "1": {
      "Mean": 4840,
      "Min": 4600,
      "Max": 5000
    },
    "10": {
      "Mean": 4933.3335,
      "Min": 4600,
      "Max": 5000
    },
    "11": {
      "Mean": 5000,
      "Min": 5000,
      "Max": 5000
    },
    "12": {
      "Mean": 5000,
      "Min": 5000,
      "Max": 5000
    },
    "13": {
      "Mean": 4600,
      "Min": 4600,
      "Max": 4600
    },
    "14": {
      "Mean": 5000,
      "Min": 5000,
      "Max": 5000
    },
    "15": {
      "Mean": 5000,
      "Min": 5000,
      "Max": 5000
    },
    "2": {
      "Mean": 4917.391,
      "Min": 3400,
      "Max": 5000
    },
    "3": {
      "Mean": 4920,
      "Min": 4600,
      "Max": 5000
    },
    "4": {
      "Mean": 4833.696,
      "Min": 1300,
      "Max": 5000
    },
    "5": {
      "Mean": 4920,
      "Min": 4600,
      "Max": 5000
    },
    "6": {
      "Mean": 4882.222,
      "Min": 1300,
      "Max": 5000
    },
    "7": {
      "Mean": 5000,
      "Min": 5000,
      "Max": 5000
    },
    "8": {
      "Mean": 4905.8823,
      "Min": 4600,
      "Max": 5000
    },
    "9": {
      "Mean": 5000,
      "Min": 5000,
      "Max": 5000
    }
  },
  "IaCoreTemperature": {
    "0": {
      "Mean": 44.945454,
      "Min": 43,
      "Max": 48
    },
    "1": {
      "Mean": 45.2,
      "Min": 44,
      "Max": 46
    },
    "10": {
      "Mean": 40.333332,
      "Min": 40,
      "Max": 41
    },
    "11": {
      "Mean": 40,
      "Min": 40,
      "Max": 40
    },
    "12": {
      "Mean": 41.77778,
      "Min": 41,
      "Max": 43
    },
    "13": {
      "Mean": 41,
      "Min": 41,
      "Max": 41
    },
    "14": {
      "Mean": 39,
      "Min": 39,
      "Max": 39
    },
    "15": {
      "Mean": 39,
      "Min": 39,
      "Max": 39
    },
    "2": {
      "Mean": 44.934784,
      "Min": 41,
      "Max": 48
    },
    "3": {
      "Mean": 44.6,
      "Min": 43,
      "Max": 47
    },
    "4": {
      "Mean": 45.880436,
      "Min": 43,
      "Max": 49
    },
    "5": {
      "Mean": 46.2,
      "Min": 45,
      "Max": 47
    },
    "6": {
      "Mean": 41.822224,
      "Min": 41,
      "Max": 43
    },
    "7": {
      "Mean": 41.857143,
      "Min": 41,
      "Max": 42
    },
    "8": {
      "Mean": 38.882355,
      "Min": 38,
      "Max": 40
    },
    "9": {
      "Mean": 39,
      "Min": 39,
      "Max": 39
    }
  },
  "IaCoreUtilization": {
    "0": 1.5328028,
    "1": 0.07353569,
    "10": 0.08407396,
    "11": 0.013976383,
    "12": 0.12534368,
    "13": 0.014092508,
    "14": 0.013951499,
    "15": 0.0139639415,
    "2": 1.2775866,
    "3": 0.07005195,
    "4": 1.2790008,
    "5": 0.07203022,
    "6": 0.627917,
    "7": 0.097826384,
    "8": 0.23627552,
    "9": 0.013984677
  },
  "GtFrequency": 0,
  "GtFrequencyRequest": 1199.988,
  "GtUtilization": 0,
  "PackagePower": {
    "Watts": 13.607669,
    "Joules": 13.671204
  },
  "PlatformPower": {
    "Watts": 0,
    "Joules": 0
  },
  "DramPower": {
    "Watts": 0,
    "Joules": 0
  },
  "PackageTemperature": 43,
  "Tdp": 95
}

```


## License

Please see [License File](LICENSE) for more information.