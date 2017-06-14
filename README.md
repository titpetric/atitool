# atitool
Atitool is a tool for reading ATI Radeon RX 400 and 500 series VBIOS ROM files. 

It's inspired by PolarisBiosEditor. PBE is written in C# with Winforms GUI and has a terrible cross-platform experience. Atitool is designed to overcome those limitations. It outputs the following information.

* ROM data.
* Powerplay data.
* Powertune data.
* Fan data.
* GPU voltage data.
* Memory voltage data.
* GPU clock data.
* Memory clock data.

# TODO
* Modifying values and saving to a new ROM file.

# Compiling
```
go build
```

# Help
```
usage: atitool [<flags>] <command> [<args> ...]

A command-line tool for dealing with Radeon GPU bios files.

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  show <file>
    Show values from the specified bios file.
    
```

# Example
```
./atitool show ~/stock.rom

----------------------------------------
ROM
----------------------------------------
VendorID: 0x1002
DeviceID: 0x67df
SubID: 0x3417
SubVendorID: 0x1462
Firmware signature: 0x4d4f5441

----------------------------------------
Powerplay
----------------------------------------
Max GPU freq (Mhz): 2000
Max memory freq (Mhz): 2250
Power control limit (%%): 50

----------------------------------------
Powertune
----------------------------------------
TDP (W): 145
TDC (A): 132
Max Power Limit (W): 165
Max Temp. (C): 90
Shutdown Temp. (C): 94
Hotspot Temp. (C): 105

----------------------------------------
Fan
----------------------------------------
Temp. Hysteresis: 3
Min Temp. (C): 40
Med Temp. (C): 65
High Temp. (C): 85
Max Temp. (C): 109
Legacy or Fuzzy Fan Mode: 1
Min PWM (%): 20
Med PWM (%): 40
High PWM (%): 60
Max PWM (%): 1
Max RPM: 2400
Sensitivity: 4836
Acoustic Limit (MHz): 927

----------------------------------------
GPU
----------------------------------------
300 mV: 750 Mhz
600 mV: 65282 Mhz
927 mV: 65283 Mhz
1179 mV: 65284 Mhz
1251 mV: 65285 Mhz
1294 mV: 65286 Mhz
1339 mV: 65287 Mhz
1380 mV: 65288 Mhz
```