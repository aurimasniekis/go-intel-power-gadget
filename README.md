# go-intel-power-gadget
A Golang macOS IntelPowerGadget framework binding module. Which I have written about on my blog [TechProwd.com](https://www.techprowd.com/go-intel-r-power-gadget/).

## Install

A simple way using `go get`. Requires the IntelPowerGadget installed in the system.

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
  "pkg": {
    "package_no": 0,
    "package_cores": 16,
    "ia_base_frequency": 3600,
    "ia_max_frequency": 5000,
    "gt_max_frequency": 0,
    "package_tdp": 95,
    "max_temperature": 100,
    "gt_available": true,
    "ia_energy_available": true,
    "dram_energy_available": false,
    "platform_energy_available": false
  },
  "timestamp_start": "2020-09-05T18:10:10.745816626+09:00",
  "timestamp_end": "2020-09-05T18:10:11.750083978+09:00",
  "interval": 1000000000,
  "ia_frequency": {
    "mean": 3344.1914,
    "min": 1200,
    "max": 5000
  },
  "ia_frequency_request": {
    "mean": 3416.7263,
    "min": 1300,
    "max": 5000
  },
  "ia_power": {
    "watts": 10.599621,
    "joules": 10.644775
  },
  "ia_temperature": {
    "mean": 40.278374,
    "min": 37,
    "max": 46
  },
  "ia_utilization": 0.7346676,
  "ia_core_frequency": {
    "0": {
      "mean": 3113.8245,
      "min": 1200,
      "max": 5000
    },
    "1": {
      "mean": 4127.8115,
      "min": 3300,
      "max": 5000
    },
    "10": {
      "mean": 3492.069,
      "min": 1700,
      "max": 5000
    },
    "11": {
      "mean": 3501.7676,
      "min": 3000,
      "max": 4300
    },
    "12": {
      "mean": 3534.1304,
      "min": 1600,
      "max": 5000
    },
    "13": {
      "mean": 3500.1667,
      "min": 3000,
      "max": 4000
    },
    "14": {
      "mean": 3112.9294,
      "min": 1800,
      "max": 5000
    },
    "15": {
      "mean": 3394.508,
      "min": 3300,
      "max": 3400
    },
    "2": {
      "mean": 3107.1162,
      "min": 1200,
      "max": 5000
    },
    "3": {
      "mean": 3324.2092,
      "min": 2300,
      "max": 4500
    },
    "4": {
      "mean": 3390.9626,
      "min": 1600,
      "max": 5000
    },
    "5": {
      "mean": 3353.0151,
      "min": 2300,
      "max": 4700
    },
    "6": {
      "mean": 3604.3132,
      "min": 2000,
      "max": 5000
    },
    "7": {
      "mean": 3293.2722,
      "min": 2300,
      "max": 4400
    },
    "8": {
      "mean": 3556.1345,
      "min": 1600,
      "max": 5000
    },
    "9": {
      "mean": 3914.441,
      "min": 3200,
      "max": 4500
    }
  },
  "ia_core_frequency_request": {
    "0": {
      "mean": 3191.0054,
      "min": 1300,
      "max": 5000
    },
    "1": {
      "mean": 4133.3335,
      "min": 2400,
      "max": 5000
    },
    "10": {
      "mean": 3622.4138,
      "min": 1300,
      "max": 5000
    },
    "11": {
      "mean": 3866.6667,
      "min": 2400,
      "max": 4600
    },
    "12": {
      "mean": 3637.5,
      "min": 1300,
      "max": 5000
    },
    "13": {
      "mean": 3500,
      "min": 2400,
      "max": 4600
    },
    "14": {
      "mean": 3132.1428,
      "min": 1300,
      "max": 5000
    },
    "15": {
      "mean": 5000,
      "min": 5000,
      "max": 5000
    },
    "2": {
      "mean": 3168.2927,
      "min": 1300,
      "max": 5000
    },
    "3": {
      "mean": 3142.8572,
      "min": 2400,
      "max": 4600
    },
    "4": {
      "mean": 3444.8276,
      "min": 1300,
      "max": 5000
    },
    "5": {
      "mean": 3171.4285,
      "min": 2400,
      "max": 4600
    },
    "6": {
      "mean": 3651.6128,
      "min": 2400,
      "max": 5000
    },
    "7": {
      "mean": 3500,
      "min": 2400,
      "max": 4600
    },
    "8": {
      "mean": 3657.3914,
      "min": 1300,
      "max": 5000
    },
    "9": {
      "mean": 3866.6667,
      "min": 2400,
      "max": 4600
    }
  },
  "ia_core_temperature": {
    "0": {
      "mean": 40.544975,
      "min": 40,
      "max": 44
    },
    "1": {
      "mean": 42,
      "min": 40,
      "max": 43
    },
    "10": {
      "mean": 39.75862,
      "min": 39,
      "max": 44
    },
    "11": {
      "mean": 39.333332,
      "min": 39,
      "max": 40
    },
    "12": {
      "mean": 40.94643,
      "min": 40,
      "max": 43
    },
    "13": {
      "mean": 40.5,
      "min": 40,
      "max": 41
    },
    "14": {
      "mean": 37.07143,
      "min": 37,
      "max": 38
    },
    "15": {
      "mean": 38,
      "min": 38,
      "max": 38
    },
    "2": {
      "mean": 39.715446,
      "min": 39,
      "max": 42
    },
    "3": {
      "mean": 39.714287,
      "min": 39,
      "max": 40
    },
    "4": {
      "mean": 42.41379,
      "min": 41,
      "max": 46
    },
    "5": {
      "mean": 42.142857,
      "min": 41,
      "max": 43
    },
    "6": {
      "mean": 39.55914,
      "min": 38,
      "max": 46
    },
    "7": {
      "mean": 39.5,
      "min": 38,
      "max": 42
    },
    "8": {
      "mean": 39.06087,
      "min": 37,
      "max": 46
    },
    "9": {
      "mean": 37.666668,
      "min": 37,
      "max": 39
    }
  },
  "ia_core_utilization": {
    "0": 2.6572914,
    "1": 0.042078994,
    "10": 0.8163582,
    "11": 0.042066548,
    "12": 0.78319126,
    "13": 0.028445434,
    "14": 0.39599562,
    "15": 0.014052608,
    "2": 1.7276866,
    "3": 0.098214746,
    "4": 2.0361304,
    "5": 0.09804464,
    "6": 1.3010633,
    "7": 0.056032028,
    "8": 1.6158384,
    "9": 0.042191017
  },
  "gt_frequency": 0,
  "gt_frequency_request": 1199.988,
  "gt_utilization": 0,
  "package_power": {
    "watts": 12.453844,
    "joules": 12.506897
  },
  "platform_power": {
    "watts": 0,
    "joules": 0
  },
  "dram_power": {
    "watts": 0,
    "joules": 0
  },
  "package_temperature": 43,
  "tdp": 95
}
```


## License

Please see [License File](LICENSE) for more information.