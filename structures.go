package main

const (
	AtomROMChecksumOffset 	= 0x21
	AtomROMHeaderPtr      	= 0x48
	AtomMaxVRAMEntries 	= 24
)

type Bios struct {
	AtomRomHeader AtomRomHeader
	AtomDataTables AtomDataTables
	AtomPowerplayTable AtomPowerplayTable
	AtomPowertuneTable AtomPowertuneTable
	AtomFanTable AtomFanTable
	AtomMClkTable AtomMClkTable
	AtomSClkTable AtomSClkTable
	AtomVoltageTable AtomVoltageTable
	AtomVRAMInfo AtomVRAMInfo
	AtomVRAMTimingEntry [AtomMaxVRAMEntries]AtomVRAMTimingEntry
}

type AtomCommonTableHeader struct {
	StructureSize        int16
	TableFormatRevision  byte
	TableContentRevision byte
}

type AtomRomHeader struct {
	Header                    AtomCommonTableHeader
	FirmWareSignature         uint32
	BiosRuntimeSegmentAddress uint16
	ProtectedModeInfoOffset   uint16
	ConfigFilenameOffset      uint16
	CRCBlockOffset            uint16
	BIOSBootupMessageOffset   uint16
	Int10Offset               uint16
	PciBusDevInitCode         uint16
	IoBaseAddress             uint16
	SubsystemVendorID         uint16
	SubsystemID               uint16
	PCIInfoOffset             uint16
	MasterCommandTableOffset  uint16
	MasterDataTableOffset     uint16
	ExtendedFunctionCode      byte
	_                         byte
	PSPDirTableOffset         uint32
	VendorID                  uint16
	DeviceID                  uint16
}

type AtomDataTables struct {
	Header                   AtomCommonTableHeader
	UtilityPipeLine          uint16
	MultimediaCapabilityInfo uint16
	MultimediaConfigInfo     uint16
	StandardVESATiming       uint16
	FirmwareInfo             uint16
	PaletteData              uint16
	LCDInfo                  uint16
	DIGTransmitterInfo       uint16
	SMUInfo                  uint16
	SupportedDevicesInfo     uint16
	GPIOI2CInfo              uint16
	VRAMUsageByFirmware      uint16
	GPIOPinLUT               uint16
	VESAToInternalModeLUT    uint16
	GFXInfo                  uint16
	PowerPlayInfo            uint16
	GPUVirtualizationInfo    uint16
	SaveRestoreInfo          uint16
	PPLLSSInfo               uint16
	OemInfo                  uint16
	XTMDSInfo                uint16
	MclkSSInfo               uint16
	ObjectHeader             uint16
	IndirectIOAccess         uint16
	MCInitParameter          uint16
	ASICVDDCInfo             uint16
	ASICInternalSSInfo       uint16
	TVVideoMode              uint16
	VRAMInfo                 uint16
	MemoryTrainingInfo       uint16
	IntegratedSystemInfo     uint16
	ASICProfilingInfo        uint16
	VoltageObjectInfo        uint16
	PowerSourceInfo          uint16
	ServiceInfo              uint16
}

type AtomPowerplayTable struct {
	Header                    AtomCommonTableHeader
	TableRevision             byte
	TableSize                 uint16
	GoldenPPID                uint32
	GoldenRevision            uint32
	FormatID                  uint16
	VoltageTime               uint16
	PlatformCaps              uint32
	MaxODEngineClock          uint32
	MaxODMemoryClock          uint32
	PowerControlLimit         uint16
	UlvVoltageOffset          uint16
	StateArrayOffset          uint16
	FanTableOffset            uint16
	ThermalControllerOffset   uint16
	_                         uint16
	MclkDependencyTableOffset uint16
	SclkDependencyTableOffset uint16
	VddcLookupTableOffset     uint16
	VddgfxLookupTableOffset   uint16
	MMDependencyTableOffset   uint16
	VCEStateTableOffset       uint16
	PPMTableOffset            uint16
	PowerTuneTableOffset      uint16
	HardLimitTableOffset      uint16
	PCIETableOffset           uint16
	GPIOTableOffset           uint16
	_                         [6]uint16
}

type AtomMClkEntry struct {
	VddcInd      byte
	Vddci        uint16
	VddgfxOffset uint16
	Mvdd         uint16
	Mclk         uint32
	_            uint16
}

type AtomMClkTable struct {
	RevID      byte
	NumEntries byte
	Entries    []AtomMClkEntry
}

type AtomSClkEntry struct {
	VddInd                 byte
	VddcOffset             uint16
	Sclk                   uint32
	EdcCurrent             uint16
	ReliabilityTemperature byte
	CKSVOffsetandDisable   byte
	SclkOffset             uint32
	// Polaris Only, remove for compatibility with Fiji
}

type AtomSClkTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomSClkEntry
}

type AtomVoltageEntry struct {
	Vdd     uint16
	CACLow  uint16
	CACMid  uint16
	CACHigh uint16
}

type AtomVoltageTable struct {
	RevID      byte
	NumEntries byte `struct:"sizeof=Entries"`
	Entries    []AtomVoltageEntry
}

type AtomFanTable struct {
	RevID                   byte
	THyst                   byte
	TMin                    uint16
	TMed                    uint16
	THigh                   uint16
	PWMMin                  uint16
	PWMMed                  uint16
	PWMHigh                 uint16
	TMax                    uint16
	FanControlMode          byte
	FanPWMMax               uint16
	FanOutputSensitivity    uint16
	FanRPMMax               uint16
	MinFanSCLKAcousticLimit uint32
	TargetTemperature       byte
	MinimumPWMLimit         byte
	FanGainEdge             uint16
	FanGainHotspot          uint16
	FanGainLiquid           uint16
	FanGainVrVddc           uint16
	FanGainVrMvdd           uint16
	FanGainPlx              uint16
	FanGainHbm              uint16
	_                       uint16
}

type AtomPowertuneTable struct {
	RevID                     byte
	TDP                       uint16
	ConfigurableTDP           uint16
	TDC                       uint16
	BatteryPowerLimit         uint16
	SmallPowerLimit           uint16
	LowCACLeakage             uint16
	HighCACLeakage            uint16
	MaximumPowerDeliveryLimit uint16
	TjMax                     uint16
	PowerTuneDataSetID        uint16
	EDCLimit                  uint16
	SoftwareShutdownTemp      uint16
	ClockStretchAmount        uint16
	TemperatureLimitHotspot   uint16
	TemperatureLimitLiquid1   uint16
	TemperatureLimitLiquid2   uint16
	TemperatureLimitVrVddc    uint16
	TemperatureLimitVrMvdd    uint16
	TemperatureLimitPlx       uint16
	Liquid1I2CAddress         byte
	Liquid2I2CAddress         byte
	LiquidI2CLine             byte
	VrI2CAddress              byte
	VrI2CLine                 byte
	PlxI2CAddress             byte
	PlxI2CLine                byte
	_                         uint16
}

type AtomVRAMTimingEntry struct {
	ClkRange uint32
	Latency  [0x30]byte
}

type AtomVRAMEntry struct {
	ChannelMapCfg     uint32
	ModuleSize        uint16
	McRAMCfg          uint16
	EnableChannels    uint16
	ExtMemoryID       byte
	MemoryType        byte
	ChannelNum        byte
	ChannelWidth      byte
	Density           byte
	BankCol           byte
	Misc              byte
	VREFI             byte
	_                 uint16
	MemorySize        uint16
	McTunningSetID    byte
	RowNum            byte
	EMRS2Value        uint16
	EMRS3Value        uint16
	MemoryVenderID    byte
	RefreshRateFactor byte
	FIFODepth         byte
	CDRBandwidth      byte
	ChannelMapCfg1    uint32
	BankMapCfg        uint32
	_                 uint32
	MemPNString       string `struct:"[20]byte"`
}

type AtomVRAMInfo struct {
	Header                   AtomCommonTableHeader
	MemAdjustTblOffset       uint16
	MemClkPatchTblOffset     uint16
	McAdjustPerTileTblOffset uint16
	McPhyInitTableOffset     uint16
	DramDataRemapTblOffset   uint16
	_                        uint16
	NumOfVRAMModule          byte `struct:"sizeof=VramInfo"`
	MemoryClkPatchTblVer     byte
	VramModuleVer            byte
	McPhyTileNum             byte
	VramInfo                 []AtomVRAMEntry
}
