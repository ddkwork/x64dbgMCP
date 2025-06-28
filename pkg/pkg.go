package pkg

import (
	"encoding/hex"
	"fmt"
)

type (
	command   struct{}
	register  struct{}
	memory    struct{}
	debug     struct{}
	assembler struct{}
	stack     struct{}
	flag      struct{}
	pattern   struct{}
	misc      struct{}
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

func (memory) Read(address int, size uint) []byte {
	return request[[]byte]("Memory/Read", map[string]string{"addr": fmt.Sprintf("0x%x", address), "size": fmt.Sprintf("%d", size)})
}

func (memory) Write(address int, data []byte) bool {
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
	Status string `json:"success"`
	Size   int    `json:"size"`
	Data   []byte `json:"bytes"`
}

func (assembler) Assemble(address int, instruction string) assemblerResult {
	return request[assemblerResult]("Assembler/Assemble", map[string]string{"addr": fmt.Sprintf("0x%x", address), "instruction": instruction})
}
func (assembler) AssembleMem(address int, instructionOpcodes []byte) bool {
	return request[bool]("Assembler/AssembleMem", map[string]string{"addr": fmt.Sprintf("0x%x", address), "instruction": hex.EncodeToString(instructionOpcodes)})
}

func (stack) Pop() uint { //todo 改成泛型
	return request[uint]("Stack/Pop", nil)
}
func (stack) Push(value uint) uint {
	return request[uint]("Stack/Push", map[string]string{"value": fmt.Sprintf("0x%x", value)})
}
func (stack) Peek(offset int) uint {
	return request[uint]("Stack/Peek", map[string]string{"offset": fmt.Sprintf("0x%x", offset)})
}

type disassembler struct{}

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
	Address     int    `json:"address"`
	Instruction string `json:"instruction"`
	Size        string `json:"size"`
}
type disassembleRip struct {
	Rip         int    `json:"rip"`
	Instruction string `json:"instruction"`
	Size        string `json:"size"`
}
type disassembleRipWithSetupIn struct {
	StepResult  string `json:"step_result"`
	Rip         int    `json:"rip"`
	Instruction string `json:"instruction"`
	Size        string `json:"size"`
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
func (pattern) FindMemory(start int, size int, pattern string) (address uint) {
	return request[uint]("Pattern/FindMem", map[string]string{"start": fmt.Sprintf("0x%x", start), "size": fmt.Sprintf("0x%x", size), "pattern": pattern})
}

func (misc) ParseExpression(expression string) (value uint) {
	return request[uint]("Misc/ParseExpression", map[string]string{"expression": expression})
}

func (misc) GetApiAddressFromModule(module string, api string) (address uint) {
	return request[uint]("Misc/RemoteGetProcAddress", map[string]string{"module": module, "api": api})
}

type memoryBase struct {
	BaseAddress uint `json:"base_address"`
	Size        uint `json:"size"`
}

func (misc) FindMemoryBaseByAddress(address int) memoryBase {
	return request[memoryBase]("MemoryBase", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
}

func (misc) GetModuleList() {}
