package main

import (
	"cmp"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ddkwork/golibrary/std/mylog"
)

type RegisterEnum int
type RegisterManager struct{}

func (m RegisterManager) GetDR0() uint {
	return request[uint]("Register/Get", map[string]string{"register": "DR0"})
}
func (m RegisterManager) SetDR0(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DR0", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDR1() uint {
	return request[uint]("Register/Get", map[string]string{"register": "DR1"})
}
func (m RegisterManager) SetDR1(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DR1", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDR2() uint {
	return request[uint]("Register/Get", map[string]string{"register": "DR2"})
}
func (m RegisterManager) SetDR2(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DR2", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDR3() uint {
	return request[uint]("Register/Get", map[string]string{"register": "DR3"})
}
func (m RegisterManager) SetDR3(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DR3", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDR6() uint {
	return request[uint]("Register/Get", map[string]string{"register": "DR6"})
}
func (m RegisterManager) SetDR6(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DR6", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDR7() uint {
	return request[uint]("Register/Get", map[string]string{"register": "DR7"})
}
func (m RegisterManager) SetDR7(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DR7", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetEAX() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "EAX"})
}
func (m RegisterManager) SetEAX(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "EAX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetAX() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "AX"})
}
func (m RegisterManager) SetAX(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "AX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetAH() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "AH"})
}
func (m RegisterManager) SetAH(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "AH", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetAL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "AL"})
}
func (m RegisterManager) SetAL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "AL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetEBX() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "EBX"})
}
func (m RegisterManager) SetEBX(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "EBX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetBX() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "BX"})
}
func (m RegisterManager) SetBX(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "BX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetBH() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "BH"})
}
func (m RegisterManager) SetBH(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "BH", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetBL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "BL"})
}
func (m RegisterManager) SetBL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "BL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetECX() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "ECX"})
}
func (m RegisterManager) SetECX(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "ECX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCX() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "CX"})
}
func (m RegisterManager) SetCX(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCH() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "CH"})
}
func (m RegisterManager) SetCH(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CH", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "CL"})
}
func (m RegisterManager) SetCL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetEDX() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "EDX"})
}
func (m RegisterManager) SetEDX(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "EDX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDX() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "DX"})
}
func (m RegisterManager) SetDX(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDH() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "DH"})
}
func (m RegisterManager) SetDH(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DH", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "DL"})
}
func (m RegisterManager) SetDL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetEDI() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "EDI"})
}
func (m RegisterManager) SetEDI(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "EDI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDI() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "DI"})
}
func (m RegisterManager) SetDI(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetESI() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "ESI"})
}
func (m RegisterManager) SetESI(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "ESI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetSI() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "SI"})
}
func (m RegisterManager) SetSI(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "SI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetEBP() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "EBP"})
}
func (m RegisterManager) SetEBP(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "EBP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetBP() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "BP"})
}
func (m RegisterManager) SetBP(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "BP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetESP() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "ESP"})
}
func (m RegisterManager) SetESP(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "ESP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetSP() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "SP"})
}
func (m RegisterManager) SetSP(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "SP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetEIP() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "EIP"})
}
func (m RegisterManager) SetEIP(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "EIP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRAX() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RAX"})
}
func (m RegisterManager) SetRAX(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RAX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRBX() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RBX"})
}
func (m RegisterManager) SetRBX(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RBX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRCX() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RCX"})
}
func (m RegisterManager) SetRCX(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RCX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRDX() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RDX"})
}
func (m RegisterManager) SetRDX(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RDX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRSI() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RSI"})
}
func (m RegisterManager) SetRSI(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RSI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetSIL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "SIL"})
}
func (m RegisterManager) SetSIL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "SIL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRDI() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RDI"})
}
func (m RegisterManager) SetRDI(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RDI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetDIL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "DIL"})
}
func (m RegisterManager) SetDIL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "DIL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRBP() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RBP"})
}
func (m RegisterManager) SetRBP(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RBP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetBPL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "BPL"})
}
func (m RegisterManager) SetBPL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "BPL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRSP() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RSP"})
}
func (m RegisterManager) SetRSP(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RSP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetSPL() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "SPL"})
}
func (m RegisterManager) SetSPL(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "SPL", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetRIP() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "RIP"})
}
func (m RegisterManager) SetRIP(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "RIP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR8() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R8"})
}
func (m RegisterManager) SetR8(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R8", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR8D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R8D"})
}
func (m RegisterManager) SetR8D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R8D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR8W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R8W"})
}
func (m RegisterManager) SetR8W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R8W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR8B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R8B"})
}
func (m RegisterManager) SetR8B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R8B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR9() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R9"})
}
func (m RegisterManager) SetR9(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R9", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR9D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R9D"})
}
func (m RegisterManager) SetR9D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R9D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR9W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R9W"})
}
func (m RegisterManager) SetR9W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R9W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR9B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R9B"})
}
func (m RegisterManager) SetR9B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R9B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR10() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R10"})
}
func (m RegisterManager) SetR10(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R10", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR10D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R10D"})
}
func (m RegisterManager) SetR10D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R10D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR10W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R10W"})
}
func (m RegisterManager) SetR10W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R10W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR10B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R10B"})
}
func (m RegisterManager) SetR10B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R10B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR11() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R11"})
}
func (m RegisterManager) SetR11(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R11", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR11D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R11D"})
}
func (m RegisterManager) SetR11D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R11D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR11W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R11W"})
}
func (m RegisterManager) SetR11W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R11W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR11B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R11B"})
}
func (m RegisterManager) SetR11B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R11B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR12() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R12"})
}
func (m RegisterManager) SetR12(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R12", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR12D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R12D"})
}
func (m RegisterManager) SetR12D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R12D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR12W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R12W"})
}
func (m RegisterManager) SetR12W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R12W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR12B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R12B"})
}
func (m RegisterManager) SetR12B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R12B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR13() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R13"})
}
func (m RegisterManager) SetR13(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R13", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR13D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R13D"})
}
func (m RegisterManager) SetR13D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R13D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR13W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R13W"})
}
func (m RegisterManager) SetR13W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R13W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR13B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R13B"})
}
func (m RegisterManager) SetR13B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R13B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR14() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R14"})
}
func (m RegisterManager) SetR14(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R14", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR14D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R14D"})
}
func (m RegisterManager) SetR14D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R14D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR14W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R14W"})
}
func (m RegisterManager) SetR14W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R14W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR14B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R14B"})
}
func (m RegisterManager) SetR14B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R14B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR15() uint64 {
	return request[uint64]("Register/Get", map[string]string{"register": "R15"})
}
func (m RegisterManager) SetR15(v uint64) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R15", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR15D() uint32 {
	return request[uint32]("Register/Get", map[string]string{"register": "R15D"})
}
func (m RegisterManager) SetR15D(v uint32) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R15D", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR15W() uint16 {
	return request[uint16]("Register/Get", map[string]string{"register": "R15W"})
}
func (m RegisterManager) SetR15W(v uint16) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R15W", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetR15B() uint8 {
	return request[uint8]("Register/Get", map[string]string{"register": "R15B"})
}
func (m RegisterManager) SetR15B(v uint8) bool {
	return request[bool]("Register/Set", map[string]string{"register": "R15B", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCIP() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CIP"})
}
func (m RegisterManager) SetCIP(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CIP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCSP() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CSP"})
}
func (m RegisterManager) SetCSP(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CSP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCAX() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CAX"})
}
func (m RegisterManager) SetCAX(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CAX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCBX() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CBX"})
}
func (m RegisterManager) SetCBX(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CBX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCCX() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CCX"})
}
func (m RegisterManager) SetCCX(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CCX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCDX() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CDX"})
}
func (m RegisterManager) SetCDX(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CDX", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCDI() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CDI"})
}
func (m RegisterManager) SetCDI(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CDI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCSI() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CSI"})
}
func (m RegisterManager) SetCSI(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CSI", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCBP() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CBP"})
}
func (m RegisterManager) SetCBP(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CBP", "value": strconv.FormatUint(uint64(v), 16)})
}
func (m RegisterManager) GetCFLAGS() uint {
	return request[uint]("Register/Get", map[string]string{"register": "CFLAGS"})
}
func (m RegisterManager) SetCFLAGS(v uint) bool {
	return request[bool]("Register/Set", map[string]string{"register": "CFLAGS", "value": strconv.FormatUint(uint64(v), 16)})
}

func (m RegisterManager) Get(reg RegisterEnum) uint {
	switch reg {
	case DR0:
		return uint(m.GetDR0())
	case DR1:
		return uint(m.GetDR1())
	case DR2:
		return uint(m.GetDR2())
	case DR3:
		return uint(m.GetDR3())
	case DR6:
		return uint(m.GetDR6())
	case DR7:
		return uint(m.GetDR7())
	case EAX:
		return uint(m.GetEAX())
	case AX:
		return uint(m.GetAX())
	case AH:
		return uint(m.GetAH())
	case AL:
		return uint(m.GetAL())
	case EBX:
		return uint(m.GetEBX())
	case BX:
		return uint(m.GetBX())
	case BH:
		return uint(m.GetBH())
	case BL:
		return uint(m.GetBL())
	case ECX:
		return uint(m.GetECX())
	case CX:
		return uint(m.GetCX())
	case CH:
		return uint(m.GetCH())
	case CL:
		return uint(m.GetCL())
	case EDX:
		return uint(m.GetEDX())
	case DX:
		return uint(m.GetDX())
	case DH:
		return uint(m.GetDH())
	case DL:
		return uint(m.GetDL())
	case EDI:
		return uint(m.GetEDI())
	case DI:
		return uint(m.GetDI())
	case ESI:
		return uint(m.GetESI())
	case SI:
		return uint(m.GetSI())
	case EBP:
		return uint(m.GetEBP())
	case BP:
		return uint(m.GetBP())
	case ESP:
		return uint(m.GetESP())
	case SP:
		return uint(m.GetSP())
	case EIP:
		return uint(m.GetEIP())
	case RAX:
		return uint(m.GetRAX())
	case RBX:
		return uint(m.GetRBX())
	case RCX:
		return uint(m.GetRCX())
	case RDX:
		return uint(m.GetRDX())
	case RSI:
		return uint(m.GetRSI())
	case SIL:
		return uint(m.GetSIL())
	case RDI:
		return uint(m.GetRDI())
	case DIL:
		return uint(m.GetDIL())
	case RBP:
		return uint(m.GetRBP())
	case BPL:
		return uint(m.GetBPL())
	case RSP:
		return uint(m.GetRSP())
	case SPL:
		return uint(m.GetSPL())
	case RIP:
		return uint(m.GetRIP())
	case R8:
		return uint(m.GetR8())
	case R8D:
		return uint(m.GetR8D())
	case R8W:
		return uint(m.GetR8W())
	case R8B:
		return uint(m.GetR8B())
	case R9:
		return uint(m.GetR9())
	case R9D:
		return uint(m.GetR9D())
	case R9W:
		return uint(m.GetR9W())
	case R9B:
		return uint(m.GetR9B())
	case R10:
		return uint(m.GetR10())
	case R10D:
		return uint(m.GetR10D())
	case R10W:
		return uint(m.GetR10W())
	case R10B:
		return uint(m.GetR10B())
	case R11:
		return uint(m.GetR11())
	case R11D:
		return uint(m.GetR11D())
	case R11W:
		return uint(m.GetR11W())
	case R11B:
		return uint(m.GetR11B())
	case R12:
		return uint(m.GetR12())
	case R12D:
		return uint(m.GetR12D())
	case R12W:
		return uint(m.GetR12W())
	case R12B:
		return uint(m.GetR12B())
	case R13:
		return uint(m.GetR13())
	case R13D:
		return uint(m.GetR13D())
	case R13W:
		return uint(m.GetR13W())
	case R13B:
		return uint(m.GetR13B())
	case R14:
		return uint(m.GetR14())
	case R14D:
		return uint(m.GetR14D())
	case R14W:
		return uint(m.GetR14W())
	case R14B:
		return uint(m.GetR14B())
	case R15:
		return uint(m.GetR15())
	case R15D:
		return uint(m.GetR15D())
	case R15W:
		return uint(m.GetR15W())
	case R15B:
		return uint(m.GetR15B())
	case CIP:
		return uint(m.GetCIP())
	case CSP:
		return uint(m.GetCSP())
	case CAX:
		return uint(m.GetCAX())
	case CBX:
		return uint(m.GetCBX())
	case CCX:
		return uint(m.GetCCX())
	case CDX:
		return uint(m.GetCDX())
	case CDI:
		return uint(m.GetCDI())
	case CSI:
		return uint(m.GetCSI())
	case CBP:
		return uint(m.GetCBP())
	case CFLAGS:
		return uint(m.GetCFLAGS())
	default:
		panic("Invalid register enum")
	}
}

func (m RegisterManager) Set(reg RegisterEnum, value uint) bool {
	switch reg {
	case DR0:
		return m.SetDR0(uint(value))
	case DR1:
		return m.SetDR1(uint(value))
	case DR2:
		return m.SetDR2(uint(value))
	case DR3:
		return m.SetDR3(uint(value))
	case DR6:
		return m.SetDR6(uint((value)))
	case DR7:
		return m.SetDR7(uint(value))
	case EAX:
		return m.SetEAX(uint32(value))
	case AX:
		return m.SetAX(uint16(value))
	case AH:
		return m.SetAH(uint8(value))
	case AL:
		return m.SetAL(uint8(value))
	case EBX:
		return m.SetEBX(uint32(value))
	case BX:
		return m.SetBX(uint16(value))
	case BH:
		return m.SetBH(uint8(value))
	case BL:
		return m.SetBL(uint8(value))
	case ECX:
		return m.SetECX(uint32(value))
	case CX:
		return m.SetCX(uint16(value))
	case CH:
		return m.SetCH(uint8(value))
	case CL:
		return m.SetCL(uint8(value))
	case EDX:
		return m.SetEDX(uint32(value))
	case DX:
		return m.SetDX(uint16(value))
	case DH:
		return m.SetDH(uint8(value))
	case DL:
		return m.SetDL(uint8(value))
	case EDI:
		return m.SetEDI(uint32(value))
	case DI:
		return m.SetDI(uint16(value))
	case ESI:
		return m.SetESI(uint32(value))
	case SI:
		return m.SetSI(uint16(value))
	case EBP:
		return m.SetEBP(uint32(value))
	case BP:
		return m.SetBP(uint16(value))
	case ESP:
		return m.SetESP(uint32(value))
	case SP:
		return m.SetSP(uint16(value))
	case EIP:
		return m.SetEIP(uint32(value))
	case RAX:
		return m.SetRAX(uint64(value))
	case RBX:
		return m.SetRBX(uint64(value))
	case RCX:
		return m.SetRCX(uint64(value))
	case RDX:
		return m.SetRDX(uint64(value))
	case RSI:
		return m.SetRSI(uint64(value))
	case SIL:
		return m.SetSIL(uint8(value))
	case RDI:
		return m.SetRDI(uint64(value))
	case DIL:
		return m.SetDIL(uint8(value))
	case RBP:
		return m.SetRBP(uint64(value))
	case BPL:
		return m.SetBPL(uint8(value))
	case RSP:
		return m.SetRSP(uint64(value))
	case SPL:
		return m.SetSPL(uint8(value))
	case RIP:
		return m.SetRIP(uint64(value))
	case R8:
		return m.SetR8(uint64(value))
	case R8D:
		return m.SetR8D(uint32(value))
	case R8W:
		return m.SetR8W(uint16(value))
	case R8B:
		return m.SetR8B(uint8(value))
	case R9:
		return m.SetR9(uint64(value))
	case R9D:
		return m.SetR9D(uint32(value))
	case R9W:
		return m.SetR9W(uint16(value))
	case R9B:
		return m.SetR9B(uint8(value))
	case R10:
		return m.SetR10(uint64(value))
	case R10D:
		return m.SetR10D(uint32(value))
	case R10W:
		return m.SetR10W(uint16(value))
	case R10B:
		return m.SetR10B(uint8(value))
	case R11:
		return m.SetR11(uint64(value))
	case R11D:
		return m.SetR11D(uint32(value))
	case R11W:
		return m.SetR11W(uint16(value))
	case R11B:
		return m.SetR11B(uint8(value))
	case R12:
		return m.SetR12(uint64(value))
	case R12D:
		return m.SetR12D(uint32(value))
	case R12W:
		return m.SetR12W(uint16(value))
	case R12B:
		return m.SetR12B(uint8(value))
	case R13:
		return m.SetR13(uint64(value))
	case R13D:
		return m.SetR13D(uint32(value))
	case R13W:
		return m.SetR13W(uint16(value))
	case R13B:
		return m.SetR13B(uint8(value))
	case R14:
		return m.SetR14(uint64(value))
	case R14D:
		return m.SetR14D(uint32(value))
	case R14W:
		return m.SetR14W(uint16(value))
	case R14B:
		return m.SetR14B(uint8(value))
	case R15:
		return m.SetR15(uint64(value))
	case R15D:
		return m.SetR15D(uint32(value))
	case R15W:
		return m.SetR15W(uint16(value))
	case R15B:
		return m.SetR15B(uint8(value))
	case CIP:
		return m.SetCIP(uint(value))
	case CSP:
		return m.SetCSP(uint(value))
	case CAX:
		return m.SetCAX(uint(value))
	case CBX:
		return m.SetCBX(uint(value))
	case CCX:
		return m.SetCCX(uint(value))
	case CDX:
		return m.SetCDX(uint(value))
	case CDI:
		return m.SetCDI(uint(value))
	case CSI:
		return m.SetCSI(uint(value))
	case CBP:
		return m.SetCBP(uint(value))
	case CFLAGS:
		return m.SetCFLAGS(uint(value))
	default:
		panic("Invalid register enum")
	}
}

const DefaultX64dbgServer = "http://127.0.0.1:8888/"

var client = &http.Client{
	Timeout: 15 * time.Second,
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

//	Type type Ordered interface {
//		~int | ~int8 | ~int16 | ~int32 | ~int64 |
//			~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
//			~float32 | ~float64 |
//			~string
//	}
type Type interface {
	cmp.Ordered |
	bool |
	[]byte |
	moduleInfo |
	[]moduleInfo |
	moduleSectionInfo |
	[]moduleSectionInfo |
	moduleExport |
	[]moduleExport |
	moduleImport |
	[]moduleImport |
	memoryBase |
	disassemblerAddress |
	disassembleRip |
	disassembleRipWithSetupIn |
	assemblerResult |
	void
}

func request[T Type](endpoint string, params map[string]string) T {
	x64dbgServerURL := DefaultX64dbgServer
	url := x64dbgServerURL + endpoint

	// 添加查询参数
	if len(params) > 0 {
		query := ""
		for key, value := range params {
			query += key + "=" + value + "&"
		}
		url += "?" + strings.TrimSuffix(query, "&")
	}

	resp := mylog.Check2(client.Get(url))
	defer func() {
		mylog.Check(resp.Body.Close())
	}()

	body := mylog.Check2(io.ReadAll(resp.Body))

	if resp.StatusCode != http.StatusOK {
		mylog.Check(fmt.Sprintf("Error %d: %s", resp.StatusCode, string(body)))
	}

	str := strings.TrimSpace(string(body))
	str = strings.TrimPrefix(str, "0x")
	var jsonData T
	switch v := any(jsonData).(type) {
	case bool:
		if strings.EqualFold(str, "true") {
			return any(true).(T)
		}
		if strings.EqualFold(str, "false") {
			return any(false).(T)
		}
	case []byte:
		b := mylog.Check2(hex.DecodeString(str))
		return any(b).(T)
	case int:

		value := mylog.Check2(strconv.ParseInt(str, 16, 64))
		return any(value).(T)
	case int8:

		value := mylog.Check2(strconv.ParseInt(str, 16, 8))
		return any(value).(T)
	case int16:

		value := mylog.Check2(strconv.ParseInt(str, 16, 16))
		return any(value).(T)
	case int32:

		value := mylog.Check2(strconv.ParseInt(str, 16, 32))
		return any(value).(T)
	case int64:

		value := mylog.Check2(strconv.ParseInt(str, 16, 64))
		return any(value).(T)
	case uint:

		value := mylog.Check2(strconv.ParseUint(str, 16, 64))
		return any(value).(T)
	case uint8:

		value := mylog.Check2(strconv.ParseUint(str, 16, 8))
		return any(value).(T)
	case uint16:

		value := mylog.Check2(strconv.ParseUint(str, 16, 16))
		return any(value).(T)
	case uint32:

		value := mylog.Check2(strconv.ParseUint(str, 16, 32))
		return any(value).(T)
	case uint64:

		value := mylog.Check2(strconv.ParseUint(str, 16, 64))
		return any(value).(T)
	case uintptr:

		value := mylog.Check2(strconv.ParseUint(str, 16, 64))
		return any(value).(T)
	case float32:
		value := mylog.Check2(strconv.ParseFloat(str, 32))
		return any(value).(T)
	case float64:
		value := mylog.Check2(strconv.ParseFloat(str, 64))
		return any(value).(T)
	case string:
		return any(str).(T)

		//todo 处理cpp服务端的字段返回 0x12345678 这种格式，我估计json会解码失败
	case moduleInfo:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case []moduleInfo:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case moduleSectionInfo:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case []moduleSectionInfo:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case moduleExport:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case []moduleExport:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case moduleImport:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case []moduleImport:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case memoryBase:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case disassemblerAddress:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case disassembleRip:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case disassembleRipWithSetupIn:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case assemblerResult:
		mylog.Check(json.Unmarshal(body, &v))
		return any(v).(T)
	case void:
		return any(nil).(T)

	}
	panic("not support type")
}

const (
	DR0 RegisterEnum = iota
	DR1
	DR2
	DR3
	DR6
	DR7

	EAX
	AX
	AH
	AL
	EBX
	BX
	BH
	BL
	ECX
	CX
	CH
	CL
	EDX
	DX
	DH
	DL
	EDI
	DI
	ESI
	SI
	EBP
	BP
	ESP
	SP
	EIP

	RAX
	RBX
	RCX
	RDX
	RSI
	SIL
	RDI
	DIL
	RBP
	BPL
	RSP
	SPL
	RIP
	R8
	R8D
	R8W
	R8B
	R9
	R9D
	R9W
	R9B
	R10
	R10D
	R10W
	R10B
	R11
	R11D
	R11W
	R11B
	R12
	R12D
	R12W
	R12B
	R13
	R13D
	R13W
	R13B
	R14
	R14D
	R14W
	R14B
	R15
	R15D
	R15W
	R15B

	CIP
	CSP
	CAX
	CBX
	CCX
	CDX
	CDI
	CSI
	CBP
	CFLAGS
)
