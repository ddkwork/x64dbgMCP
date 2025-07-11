package main

import (
	"encoding/hex"
	"fmt"
)

type (
	command      struct{}
	register     struct{}
	memory       struct{}
	debug        struct{}
	assembler    struct{}
	stack        struct{}
	flag         struct{}
	pattern      struct{}
	misc         struct{}
	module       struct{}
	disassembler struct{}

	x64dbg struct {
		Command      command
		Register     register
		Memory       memory
		Debug        debug
		Assembler    assembler
		Stack        stack
		Flag         flag
		Pattern      pattern
		Misc         misc
		Module       module
		Disassembler disassembler
	}
)

func (command) Exec(cmd string) string {
	return request[string]("ExecCommand", map[string]string{"cmd": cmd})
}

func (debug) Active() bool {
	return request[bool]("IsDebugActive", nil)
}

func (debug) Debugging() bool {
	return request[bool]("Is_Debugging", nil)
}

func (memory) Read(address int, size uint) HexBytes {
	return request[HexBytes]("Memory/Read", map[string]string{"addr": fmt.Sprintf("0x%x", address), "size": fmt.Sprintf("%d", size)})
}

func (memory) Write(address int, data HexBytes) bool {
	return request[bool]("Memory/Write", map[string]string{"addr": fmt.Sprintf("0x%x", address), "data": hex.EncodeToString(data)})
}

func (memory) IsValidPtr(address int) bool {
	return request[bool]("Memory/IsValidPtr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

func (memory) GetProtectFlag(address int) string { //todo gen enum flag
	return request[string]("Memory/GetProtect", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

type void any

func (debug) Run()      { request[void]("Debug/Run", nil) }
func (debug) Pause()    { request[void]("Debug/Pause", nil) }
func (debug) Stop()     { request[void]("Debug/Stop", nil) }
func (debug) StepIn()   { request[void]("Debug/StepIn", nil) }
func (debug) StepOver() { request[void]("Debug/StepOver", nil) }
func (debug) StepOut()  { request[void]("Debug/StepOut", nil) }
func (debug) SetBreakpoint(address int) bool { //todo 添加硬件断点
	return request[bool]("Debug/SetBreakpoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (debug) DeleteBreakpoint(address int) bool {
	return request[bool]("Debug/DeleteBreakpoint", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

type assemblerResult struct {
	Status string   `json:"success"`
	Size   int      `json:"size"`
	Data   HexBytes `json:"bytes"`
}

func (assembler) Assemble(address int, instruction string) assemblerResult {
	return request[assemblerResult]("Assembler/Assemble", map[string]string{"addr": fmt.Sprintf("0x%x", address), "instruction": instruction})
}
func (assembler) AssembleMem(address int, instructionOpcodes HexBytes) bool {
	return request[bool]("Assembler/AssembleMem", map[string]string{"addr": fmt.Sprintf("0x%x", address), "instruction": hex.EncodeToString(instructionOpcodes)})
}

func (stack) Pop() HexInt { //todo 改成泛型
	return request[HexInt]("Stack/Pop", nil)
}
func (stack) Push(value uint) HexInt {
	return request[HexInt]("Stack/Push", map[string]string{"value": fmt.Sprintf("0x%x", value)})
}
func (stack) Peek(offset int) HexInt {
	return request[HexInt]("Stack/Peek", map[string]string{"offset": fmt.Sprintf("0x%x", offset)})
}

func (disassembler) AtAddress(address int) disassemblerAddress {
	return request[disassemblerAddress]("Disasm/GetInstruction", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (disassembler) AtAddressWithSize(address int, size int) []disassemblerAddress {
	if size < 1 || size > 100 {
		panic("count should be between 1 and 100 bytes buffer")
	}
	return request[[]disassemblerAddress]("Disasm/GetInstructionRange", map[string]string{"addr": fmt.Sprintf("0x%x", address), "count": fmt.Sprintf("%d", size)})
}
func (disassembler) AtRip() disassembleRip {
	return request[disassembleRip]("Disasm/GetInstructionAtRIP", nil)
}
func (disassembler) AtRipFromStepIn() disassembleRipWithSetupIn {
	return request[disassembleRipWithSetupIn]("Disasm/StepInWithDisasm", nil)
}

type disassemblerAddress struct {
	Address     HexInt `json:"address"`
	Instruction string `json:"instruction"`
	Size        int    `json:"size"`
}
type disassembleRip struct {
	Rip         HexInt `json:"rip"`
	Instruction string `json:"instruction"`
	Size        int    `json:"size"`
}
type disassembleRipWithSetupIn struct {
	StepResult  string `json:"step_result"`
	Rip         HexInt `json:"rip"`
	Instruction string `json:"instruction"`
	Size        int    `json:"size"`
}

// Get flag: Flag name (ZF, OF, CF, PF, SF, TF, AF, DF, IF)
// todo gen enum flag
func (flag) Get(name string) bool {
	return request[bool]("Flag/Get", map[string]string{"flag": name})
}

func (flag) Set(name string, value bool) string {
	return request[string]("Flag/Set", map[string]string{"flag": name, "value": fmt.Sprintf("%v", value)})
}

// FindMemory todo 特征码支持字节切片类型
func (pattern) FindMemory(start int, size int, pattern string) (address HexInt) {
	return request[HexInt]("Pattern/FindMem", map[string]string{"start": fmt.Sprintf("0x%x", start), "size": fmt.Sprintf("%d", size), "pattern": pattern})
}

func (misc) ParseExpression(expression string) (value uint) {
	return request[uint]("Misc/ParseExpression", map[string]string{"expression": expression})
}

func (misc) GetApiAddressFromModule(module string, api string) (address HexInt) {
	return request[HexInt]("Misc/RemoteGetProcAddress", map[string]string{"module": module, "api": api})
}

type memoryBase struct {
	BaseAddress HexInt `json:"base_address"`
	Size        uint   `json:"size"`
}

func (memory) FindBaseByAddress(address int) memoryBase {
	return request[memoryBase]("MemoryBase", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

type moduleInfo struct {
	BaseAddress  HexInt `json:"base_address"`
	Size         uint   `json:"size"`
	Entry        HexInt `json:"entry"`
	SectionCount int    `json:"section_count"`
	Name         string
	Path         string
}

type moduleSectionInfo struct {
	Address HexInt `json:"address"`
	Size    uint   `json:"size"`
	Name    string
}

type moduleExport struct {
	Ordinal         HexInt `json:"ordinal"`
	Rva             HexInt `json:"rva"`
	Va              HexInt `json:"va"`
	Forwarded       bool   `json:"forwarded"`
	ForwardName     string
	Name            string
	UndecoratedName string
}

type moduleImport struct {
	IatRva          HexInt `json:"iat_rva"`
	IatVa           HexInt `json:"iat_va"`
	Ordinal         HexInt `json:"ordinal"`
	Name            string
	UndecoratedName string
}

// todo implement other method in cpp server

func (module) InfoFromAddr(address int) moduleInfo {
	return request[moduleInfo]("Module/InfoFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) InfoFromName(name string) moduleInfo {
	return request[moduleInfo]("Module/InfoFromName", map[string]string{"name": name})
}
func (module) BaseFromAddr(address int) HexInt {
	return request[HexInt]("Module/BaseFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) BaseFromName(name string) HexInt {
	return request[HexInt]("Module/BaseFromName", map[string]string{"name": name})
}
func (module) SizeFromAddr(address int) HexInt {
	return request[HexInt]("Module/SizeFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) SizeFromName(name string) HexInt {
	return request[HexInt]("Module/SizeFromName", map[string]string{"name": name})
}
func (module) NameFromAddr(address int) string {
	return request[string]("Module/NameFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) PathFromAddr(address int) string {
	return request[string]("Module/PathFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) PathFromName(name string) string {
	return request[string]("Module/PathFromName", map[string]string{"name": name})
}
func (module) EntryFromAddr(address int) HexInt {
	return request[HexInt]("Module/EntryFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) EntryFromName(name string) HexInt {
	return request[HexInt]("Module/EntryFromName", map[string]string{"name": name})
}
func (module) SectionCountFromAddr(address int) HexInt {
	return request[HexInt]("Module/SectionCountFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) SectionCountFromName(name string) HexInt {
	return request[HexInt]("Module/SectionCountFromName", map[string]string{"name": name})
}
func (module) SectionFromAddr(address int, number int) moduleSectionInfo {
	return request[moduleSectionInfo]("Module/SectionFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address), "number": fmt.Sprintf("%d", number)})
}
func (module) SectionFromName(name string, number int) moduleSectionInfo {
	return request[moduleSectionInfo]("Module/SectionFromName", map[string]string{"name": name, "number": fmt.Sprintf("%d", number)})
}
func (module) SectionListFromAddr(address int) []moduleSectionInfo {
	return request[[]moduleSectionInfo]("Module/SectionListFromAddr", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}
func (module) SectionListFromName(name string) []moduleSectionInfo {
	return request[[]moduleSectionInfo]("Module/SectionListFromName", map[string]string{"name": name})
}
func (module) GetMainModuleInfo() moduleInfo {
	return request[moduleInfo]("Module/GetMainModuleInfo", nil)
}
func (module) GetMainModuleBase() HexInt {
	return request[HexInt]("Module/GetMainModuleBase", nil)
}
func (module) GetMainModuleSize() HexInt {
	return request[HexInt]("Module/GetMainModuleSize", nil)
}
func (module) GetMainModuleEntry() HexInt {
	return request[HexInt]("Module/GetMainModuleEntry", nil)
}
func (module) GetMainModuleSectionCount() int {
	return request[int]("Module/GetMainModuleSectionCount", nil)
}
func (module) GetMainModuleName() string {
	return request[string]("Module/GetMainModuleName", nil)
}
func (module) GetMainModulePath() string {
	return request[string]("Module/GetMainModulePath", nil)
}
func (module) GetMainModuleSectionList() []moduleSectionInfo {
	return request[[]moduleSectionInfo]("Module/GetMainModuleSectionList", nil)
}
func (module) GetList() []moduleInfo {
	return request[[]moduleInfo]("Module/GetList", nil)
}
func (module) GetExports(mod moduleInfo) []moduleExport {
	return request[[]moduleExport]("Module/GetExports", map[string]string{"mod": fmt.Sprintf("%v", mod)})
}
func (module) GetImports(mod moduleInfo) []moduleImport {
	return request[[]moduleImport]("Module/GetImports", map[string]string{"mod": fmt.Sprintf("%v", mod)})
}
