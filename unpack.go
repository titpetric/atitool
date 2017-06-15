package main

import (
	"os"
	"github.com/ttacon/chalk"
	"reflect"
	"fmt"
	"encoding/binary"
	"gopkg.in/restruct.v1"
)

func unpackData(buffer []byte) Bios {
	bios := Bios{}

	// Unpack header.
	headerOffset := getValueAtPosition(buffer,16, ROM_HEADER_PTR)
	header := AtomRomHeader{}
	unpack(buffer, uint16(headerOffset), &header)
	bios.AtomRomHeader = header

	// Unpack data table.
	dataTable := AtomDataTables{}
	unpack(buffer, header.MasterDataTableOffset, &dataTable)
	bios.AtomDataTables = dataTable

	// Unpack powerplay table.
	powerplayTable := AtomPowerplayTable{}
	unpack(buffer, dataTable.PowerPlayInfo, &powerplayTable)
	bios.AtomPowerplayTable = powerplayTable

	// Unpack powertune table.
	powertuneTable := AtomPowertuneTable{}
	powertuneOffset := dataTable.PowerPlayInfo + powerplayTable.PowerTuneTableOffset
	unpack(buffer, powertuneOffset, &powertuneTable)
	bios.AtomPowertuneTable = powertuneTable

	// Unpack fan table.
	fanTable := AtomFanTable{}
	fanTableOffset := dataTable.PowerPlayInfo + powerplayTable.FanTableOffset
	unpack(buffer, fanTableOffset, &fanTable)
	bios.AtomFanTable = fanTable

	// Unpack mclk table.
	mclkTable := AtomMClkTable{}
	mclkOffset := dataTable.PowerPlayInfo + powerplayTable.VddcLookupTableOffset
	unpack(buffer, mclkOffset, &mclkTable)
	bios.AtomMClkTable = mclkTable

	// Unpack sclk table.
	sclkTable := AtomSClkTable{}
	sclkOffset := dataTable.PowerPlayInfo + powerplayTable.SclkDependencyTableOffset
	unpack(buffer, sclkOffset, &sclkTable)
	bios.AtomSClkTable = sclkTable

	// Unpack voltage table.
	voltageTable := AtomVoltageTable{}
	voltageOffset := dataTable.PowerPlayInfo + powerplayTable.VddcLookupTableOffset
	unpack(buffer, voltageOffset, &voltageTable)
	bios.AtomVoltageTable = voltageTable

	// Unpack VRAM info.
	vramInfoOffset := dataTable.VRAMInfo
	vramInfo := AtomVRAMInfo{}
	err := restruct.Unpack(buffer[vramInfoOffset:], binary.LittleEndian, &vramInfo)
	if err != nil {
		fmt.Println(chalk.Red, "Error unpacking VRAM info: ", err, chalk.Reset)
		os.Exit(1)

	}
	bios.AtomVRAMInfo = vramInfo

	// HACK: determine sizeof VRAM info.
	// See restruct issue #5.
	vramInfoData, err := restruct.Pack(binary.LittleEndian, vramInfo)
	if err != nil {
		fmt.Println(chalk.Red, "Error sizing VRAM info: ", err, chalk.Reset)
		os.Exit(1)
	}

	numberOfVRAMModule := int(vramInfo.NumOfVRAMModule)
	vramEntryOffset := int(vramInfoOffset) + len(vramInfoData)
	vramEntries := make([]AtomVRAMEntry, numberOfVRAMModule)
	for i := 0; i < numberOfVRAMModule; i++ {
		err := restruct.Unpack(buffer[vramEntryOffset:], binary.LittleEndian, &vramEntries[i])
		if err != nil {
			fmt.Println(chalk.Red, "Error unpacking VRAM entry: ", err, chalk.Reset)
		}
		vramEntryOffset += int(vramEntries[i].ModuleSize)
	}
	bios.AtomVRAMEntry = vramEntries

	// Loop over VRAM timing entries.
	vramTimingOffset := int(dataTable.VRAMInfo) + len(vramInfoData)
	vramTimingEntries := make([]AtomVRAMTimingEntry, numberOfVRAMModule)
	for i := 0; i < AtomMaxVRAMEntries; i++ {
		vramTimingEntry := AtomVRAMTimingEntry{}
		err := restruct.Unpack(buffer[vramTimingOffset:], binary.LittleEndian, &vramTimingEntry)
		if err != nil {
			fmt.Println(chalk.Red, "Error unpacking timing entry: ", err, chalk.Reset)
			os.Exit(1)
		}
		if vramTimingEntry.ClkRange == 0 {
			break
		}
		vramTimingEntries = append(vramTimingEntries, vramTimingEntry)
		vramTimingOffset += 0x34
	}
	bios.AtomVRAMTimingEntry = vramTimingEntries
	return bios
}

func unpack(buffer []byte, offset uint16, object interface{}) {
	err := restruct.Unpack(buffer[offset:], binary.LittleEndian, object)
	if err != nil {
		fmt.Println(chalk.Red, "Error unpacking ", reflect.TypeOf(object), err, chalk.Reset)
		os.Exit(1)
	}
}

func getValueAtPosition(buffer []byte, bits int32, position int32) int32 {
	if (position <= int32(len(buffer)) - 4) {
		switch bits {
		default:
		case 8:
			return int32(buffer[position])
		case 24:
			return int32(buffer[position + 2] << 16) | int32(buffer[position + 1] << 8) | int32(buffer[position])
		case 16:
			return int32(binary.LittleEndian.Uint16(buffer[position:]))
		case 32:
			return int32(binary.LittleEndian.Uint32(buffer[position:]))
		}
	}
	return 0
}

func setValueAtPosition(buffer []byte, value int32, bits int32, position int32) bool {
	if (position > int32(len(buffer)) - 4) {
		return false
	}
	switch bits {
	default:
	case 8:
		buffer[position] = byte(value)
		return true
	case 24:
		buffer[position] = byte(value)
		buffer[position + 1] = byte(value >> 8)
		buffer[position + 2] = byte(value >> 16)
		return true
	case 16:
		binary.LittleEndian.PutUint16(buffer[position:], uint16(value))
		return true
	case 32:
		binary.LittleEndian.PutUint32(buffer[position:], uint32(value))
		return true
	}

	return false
}