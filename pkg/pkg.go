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

func (memory) Read(address int) []byte {
	return request[[]byte]("Memory/Read", map[string]string{"addr": fmt.Sprintf("0x%x", address)})
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

// "Disasm/GetInstruction", {"addr": addr}
// "Disasm/GetInstructionRange", {"addr": addr, "count": str(count)}
// "Disasm/GetInstructionAtRIP"
// "Disasm/StepInWithDisasm"

func (misc) DisasmGetInstruction()      {}
func (misc) DisasmGetInstructionRange() {}
func (misc) DisasmGetInstructionAtRIP() {}
func (misc) DisasmStepInWithDisasm()    {}

// Get flag: Flag name (ZF, OF, CF, PF, SF, TF, AF, DF, IF)
// todo gen enum flag
func (flag) Get(name string) bool {
	return request[bool]("Flag/Get", map[string]string{"flag": name})
}

func (flag) Set(name string, value bool) {
	request[bool]("Flag/Set", map[string]string{"flag": name, "value": fmt.Sprintf("%v", value)})
}

// # PATTERN API
// "Pattern/FindMem", {"start": start, "size": size, "pattern": pattern}
func (pattern) FindMem(startpos int, size int, pattern string) []byte {
	return request[[]byte]("Pattern/FindMem", map[string]string{"start": fmt.Sprintf("0x%x", startpos), "size": fmt.Sprintf("0x%x", size), "pattern": pattern})
}

// # MISC API
// "Misc/ParseExpression", {"expression": expression}
// "Misc/RemoteGetProcAddress", {"module": module, "api": api}
func (misc) ParseExpression() {

}

func (misc) RemoteGetProcAddress() {}

// # LEGACY COMPATIBILITY FUNCTIONS
//
// # Construct command to set register
// "MemRead", {"addr": addr, "size": size}
// "MemWrite", {"addr": addr, "data": data}

// "GetModuleList"
func (memory) MemRead()     {}
func (memory) MemWrite()    {}
func (misc) GetModuleList() {}
