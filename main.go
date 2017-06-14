package main

import (
	"os"
	"fmt"
	"github.com/alecthomas/kingpin"
	"github.com/ttacon/chalk"
)

var (
	app 	= kingpin.New("atitool", "A command-line tool for dealing with Radeon GPU bios files.")
	show 	= app.Command("show", "Show values from the specified bios file.")
	file 	= show.Arg("file", "Bios file to open.").Required().String()

	VALID_BIOS_FILESIZE 	int64 	= 524288
	ROM_CHECKSUM_OFFSET 	int32 	= 0x21
	ROM_HEADER_PTR 		int32 	= 0x48
	VRAM_ENTRIES_COUNT	int	= 0
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case show.FullCommand():
		openFile(*file)
	}
}

func openFile(filename string) {
	file, err := os.Open( filename )
	if err != nil {
		fmt.Println(chalk.Red, err, chalk.Reset)
		os.Exit(1)
	}
	fi, err := file.Stat()
	if err != nil {
		fmt.Println(chalk.Red, err, chalk.Reset)
		os.Exit(1)
	}
	if fi.Size()  != VALID_BIOS_FILESIZE && fi.Size() != VALID_BIOS_FILESIZE / 2 {
		fmt.Println(chalk.Red, "This BIOS is non standard size. Flashing this BIOS may corrupt your graphics card.", chalk.Reset)
		os.Exit(1)
	}

	buffer := make([]byte, fi.Size())
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(chalk.Red, "Unable to read ", filename, chalk.Reset)
		os.Exit(1)
	}

	bios := unpackData(buffer)
	displayRom(bios)
	displayPowerplay(bios)
	displayPowertune(bios)
	displayFan(bios)
	displayGPU(bios)
	//displayMemory(bios) // Crashes with panic: runtime error: index out of range
	displayVRAM(bios)
}

func displayRom(bios Bios) {
	fmt.Printf("\n%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s\n", chalk.Blue, "ROM", chalk.Reset)
	fmt.Printf("%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s0x%x%s\n", chalk.Bold, "VendorID: ", chalk.White,
		bios.AtomRomHeader.VendorID, chalk.Reset)
	fmt.Printf("%s%s%s0x%x%s\n", chalk.Bold, "DeviceID: ", chalk.White,
		bios.AtomRomHeader.DeviceID, chalk.Reset)
	fmt.Printf("%s%s%s0x%x%s\n", chalk.Bold, "SubID: ", chalk.White,
		bios.AtomRomHeader.SubsystemID, chalk.Reset)
	fmt.Printf("%s%s%s0x%x%s\n", chalk.Bold, "SubVendorID: ", chalk.White,
		bios.AtomRomHeader.SubsystemVendorID, chalk.Reset)
	fmt.Printf("%s%s%s0x%x%s\n", chalk.Bold, "Firmware signature: ", chalk.White,
		bios.AtomRomHeader.FirmWareSignature, chalk.Reset)
}

func displayPowerplay(bios Bios) {
	fmt.Printf("\n%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s\n", chalk.Blue, "Powerplay",  chalk.Reset)
	fmt.Printf("%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Max GPU freq (Mhz): ", chalk.White,
		bios.AtomPowerplayTable.MaxODEngineClock / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Max memory freq (Mhz): ", chalk.White,
		bios.AtomPowerplayTable.MaxODMemoryClock / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Power control limit (%%): ", chalk.White,
		bios.AtomPowerplayTable.PowerControlLimit, chalk.Reset)
}

func displayPowertune(bios Bios) {
	fmt.Printf("\n%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s\n", chalk.Blue, "Powertune", chalk.Reset)
	fmt.Printf("%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "TDP (W): ", chalk.White,
		bios.AtomPowertuneTable.TDP, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "TDC (A): ", chalk.White,
		bios.AtomPowertuneTable.TDC, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Max Power Limit (W): ", chalk.White,
		bios.AtomPowertuneTable.MaximumPowerDeliveryLimit, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Max Temp. (C): ", chalk.White,
		bios.AtomPowertuneTable.TjMax, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Shutdown Temp. (C): ", chalk.White,
		bios.AtomPowertuneTable.SoftwareShutdownTemp, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Hotspot Temp. (C): ", chalk.White,
		bios.AtomPowertuneTable.TemperatureLimitHotspot, chalk.Reset)
}

func displayFan(bios Bios) {
	fmt.Printf("\n%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s\n", chalk.Blue, "Fan", chalk.Reset)
	fmt.Printf("%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Temp. Hysteresis: ", chalk.White,
		bios.AtomFanTable.THyst, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Min Temp. (C): ", chalk.White,
		bios.AtomFanTable.TMin / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Med Temp. (C): ", chalk.White,
		bios.AtomFanTable.TMed / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "High Temp. (C): ", chalk.White,
		bios.AtomFanTable.THigh / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Max Temp. (C): ", chalk.White,
		bios.AtomFanTable.TMax / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Legacy or Fuzzy Fan Mode: ", chalk.White,
		bios.AtomFanTable.FanControlMode, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Min PWM (%): ", chalk.White,
		bios.AtomFanTable.PWMMin / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Med PWM (%): ", chalk.White,
		bios.AtomFanTable.PWMMed / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "High PWM (%): ", chalk.White,
		bios.AtomFanTable.PWMHigh / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Max PWM (%): ", chalk.White,
		bios.AtomFanTable.FanPWMMax / 100, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Max RPM: ", chalk.White,
		bios.AtomFanTable.FanRPMMax, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Sensitivity: ", chalk.White,
		bios.AtomFanTable.FanOutputSensitivity, chalk.Reset)
	fmt.Printf("%s%s%s%d%s\n", chalk.Bold, "Acoustic Limit (MHz): ", chalk.White,
		bios.AtomFanTable.MinFanSCLKAcousticLimit / 100, chalk.Reset)
}

func displayGPU(bios Bios) {
	fmt.Printf("\n%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s\n", chalk.Blue, "GPU", chalk.Reset)
	fmt.Printf("%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)

	count := int(bios.AtomSClkTable.NumEntries)
	for i := 0; i < count; i++ {
		index := int(bios.AtomSClkTable.Entries[i].VddInd)
		fmt.Printf("%s%d %s: %s%d %s%s\n", chalk.Bold, bios.AtomSClkTable.Entries[i].Sclk / 100, "mV", chalk.White,
			bios.AtomVoltageTable.Entries[index].Vdd, "Mhz", chalk.Reset)
	}
}

func displayMemory(bios Bios) {
	fmt.Printf("\n%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s\n", chalk.Blue, "Memory", chalk.Reset)
	fmt.Printf("%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)

	count := int(bios.AtomMClkTable.NumEntries)
	fmt.Println(count)
	for i := 0; i < count; i++ {
		index := int(bios.AtomMClkTable.Entries[i].Mclk)
		fmt.Printf("%s%d %s: %s%d %s%s\n", chalk.Bold, bios.AtomMClkTable.Entries[i].VddcInd / 100, "mV", chalk.White,
			bios.AtomVoltageTable.Entries[index].Vdd, "Mhz", chalk.Reset)
	}
}

func displayVRAM(bios Bios) {
	fmt.Printf("\n%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)
	fmt.Printf("%s%s%s\n", chalk.Blue, "VRAM", chalk.Reset)
	fmt.Printf("%s----------------------------------------%s\n", chalk.Blue, chalk.Reset)

	//count := int(bios.AtomVRAMInfo.NumOfVRAMModule)
	//for i := 0; i < count; i++ {
	//	if bios.AtomVRAMInfo.VramInfo[i].MemPNString[0] != 0 {
	//		fmt.Printf("%s%s %s: %s%d %s%s\n", chalk.Bold, string(bios.AtomVRAMInfo.VramInfo[i].MemPNString[:10]), "Mhz", chalk.White,
	//			"", "Timing", chalk.Reset)
	//
	//	}
	//}

	count := int(bios.AtomVRAMInfo.NumOfVRAMModule)
	for i := 0; i < count; i++ {
		if bios.AtomVRAMInfo.VramInfo[i].MemPNString[0] != 0 {
			fmt.Printf("%s\n", string(bios.AtomVRAMInfo.VramInfo[i].MemPNString[:10]))
		}
	}
}
