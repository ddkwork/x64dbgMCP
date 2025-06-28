package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/std/stream"
)

func TestGenRegister(t *testing.T) {
	g := stream.NewGeneratedFile()
	g.P("type RegisterEnum int")
	g.P("type RegisterManager struct{}")
	for api := range strings.Lines(apis) {
		api = strings.TrimSpace(api)
		if api == "" {
			continue
		}
		if strings.HasPrefix(api, "//") {
			continue
		}
		g.P(`func (m RegisterManager) `, api, "{")

		const (
			getUrlPath = "Register/Get"
			setUrlPath = "Register/Set"
			paramName  = "register"
		)

		g.AddImport("strconv")
		if strings.HasPrefix(api, "Get") {
			split := strings.Split(api, " ")
			reg := split[0]
			reg = strings.TrimPrefix(reg, "Get")
			reg = strings.TrimSpace(reg)
			reg = strings.TrimSuffix(reg, "()")
			retType := split[1]
			g.P("return request[",
				retType,
				"](",
				strconv.Quote(getUrlPath),
				", map[string]string{",
				strconv.Quote(paramName),
				": ",
				strconv.Quote(reg),
				"})")
		}

		if strings.HasPrefix(api, "Set") {
			//func (m *RegisterManager) SetDR0(v uint) bool {
			//	return request[bool](tagSetRegister, map[string]string{"register": "DR0", "value": strconv.FormatUint(uint64(v), 16)})
			//}
			api = strings.TrimSpace(api)
			api = strings.TrimPrefix(api, "Set")
			reg, after, found := strings.Cut(api, "(")
			if found {
				_, retType, f := strings.Cut(after, ")")
				if f {
					retType = strings.TrimSpace(retType) //must be bool
					//split := strings.Split(before, " ")
					//name := split[0]
					//paramType := strings.TrimSpace(split[1])

					g.P("return request[bool](",
						strconv.Quote(setUrlPath),
						", map[string]string{",
						strconv.Quote(paramName),
						": ",
						strconv.Quote(reg),
						", ",
						strconv.Quote("value"),
						": strconv.FormatUint(uint64(v), 16)})")

				}
			}

		}

		g.P("}")
	}
	g.P()
	g.P(getSet)

	g.AddImport("encoding/json")
	g.AddImport("fmt")
	g.AddImport("io")
	g.AddImport("net/http")
	g.AddImport("strings")
	g.AddImport("time")
	g.AddImport("github.com/ddkwork/golibrary/std/mylog")
	g.P(common)
	g.P(enum)
	g.InsertPackageWithImports("main")
	stream.WriteGoFile("register.go", g.String())

}

const (
	apis = ` 
//Get(reg RegisterEnum) uint
//Set(reg RegisterEnum, value uint) bool
//Size() uint
GetDR0() uint
SetDR0(v uint) bool
GetDR1() uint
SetDR1(v uint) bool
GetDR2() uint
SetDR2(v uint) bool
GetDR3() uint
SetDR3(v uint) bool
GetDR6() uint
SetDR6(v uint) bool
GetDR7() uint
SetDR7(v uint) bool
GetEAX() uint32
SetEAX(v uint32) bool
GetAX() uint16
SetAX(v uint16) bool
GetAH() uint8
SetAH(v uint8) bool
GetAL() uint8
SetAL(v uint8) bool
GetEBX() uint32
SetEBX(v uint32) bool
GetBX() uint16
SetBX(v uint16) bool
GetBH() uint8
SetBH(v uint8) bool
GetBL() uint8
SetBL(v uint8) bool
GetECX() uint32
SetECX(v uint32) bool
GetCX() uint16
SetCX(v uint16) bool
GetCH() uint8
SetCH(v uint8) bool
GetCL() uint8
SetCL(v uint8) bool
GetEDX() uint32
SetEDX(v uint32) bool
GetDX() uint16
SetDX(v uint16) bool
GetDH() uint8
SetDH(v uint8) bool
GetDL() uint8
SetDL(v uint8) bool
GetEDI() uint32
SetEDI(v uint32) bool
GetDI() uint16
SetDI(v uint16) bool
GetESI() uint32
SetESI(v uint32) bool
GetSI() uint16
SetSI(v uint16) bool
GetEBP() uint32
SetEBP(v uint32) bool
GetBP() uint16
SetBP(v uint16) bool
GetESP() uint32
SetESP(v uint32) bool
GetSP() uint16
SetSP(v uint16) bool
GetEIP() uint32
SetEIP(v uint32) bool
GetRAX() uint64
SetRAX(v uint64) bool
GetRBX() uint64
SetRBX(v uint64) bool
GetRCX() uint64
SetRCX(v uint64) bool
GetRDX() uint64
SetRDX(v uint64) bool
GetRSI() uint64
SetRSI(v uint64) bool
GetSIL() uint8
SetSIL(v uint8) bool
GetRDI() uint64
SetRDI(v uint64) bool
GetDIL() uint8
SetDIL(v uint8) bool
GetRBP() uint64
SetRBP(v uint64) bool
GetBPL() uint8
SetBPL(v uint8) bool
GetRSP() uint64
SetRSP(v uint64) bool
GetSPL() uint8
SetSPL(v uint8) bool
GetRIP() uint64
SetRIP(v uint64) bool
GetR8() uint64
SetR8(v uint64) bool
GetR8D() uint32
SetR8D(v uint32) bool
GetR8W() uint16
SetR8W(v uint16) bool
GetR8B() uint8
SetR8B(v uint8) bool
GetR9() uint64
SetR9(v uint64) bool
GetR9D() uint32
SetR9D(v uint32) bool
GetR9W() uint16
SetR9W(v uint16) bool
GetR9B() uint8
SetR9B(v uint8) bool
GetR10() uint64
SetR10(v uint64) bool
GetR10D() uint32
SetR10D(v uint32) bool
GetR10W() uint16
SetR10W(v uint16) bool
GetR10B() uint8
SetR10B(v uint8) bool
GetR11() uint64
SetR11(v uint64) bool
GetR11D() uint32
SetR11D(v uint32) bool
GetR11W() uint16
SetR11W(v uint16) bool
GetR11B() uint8
SetR11B(v uint8) bool
GetR12() uint64
SetR12(v uint64) bool
GetR12D() uint32
SetR12D(v uint32) bool
GetR12W() uint16
SetR12W(v uint16) bool
GetR12B() uint8
SetR12B(v uint8) bool
GetR13() uint64
SetR13(v uint64) bool
GetR13D() uint32
SetR13D(v uint32) bool
GetR13W() uint16
SetR13W(v uint16) bool
GetR13B() uint8
SetR13B(v uint8) bool
GetR14() uint64
SetR14(v uint64) bool
GetR14D() uint32
SetR14D(v uint32) bool
GetR14W() uint16
SetR14W(v uint16) bool
GetR14B() uint8
SetR14B(v uint8) bool
GetR15() uint64
SetR15(v uint64) bool
GetR15D() uint32
SetR15D(v uint32) bool
GetR15W() uint16
SetR15W(v uint16) bool
GetR15B() uint8
SetR15B(v uint8) bool
GetCIP() uint
SetCIP(v uint) bool
GetCSP() uint
SetCSP(v uint) bool
GetCAX() uint
SetCAX(v uint) bool
GetCBX() uint
SetCBX(v uint) bool
GetCCX() uint
SetCCX(v uint) bool
GetCDX() uint
SetCDX(v uint) bool
GetCDI() uint
SetCDI(v uint) bool
GetCSI() uint
SetCSI(v uint) bool
GetCBP() uint
SetCBP(v uint) bool
GetCFLAGS() uint
SetCFLAGS(v uint) bool 
`

	common = `

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
	base := 10
	if strings.HasPrefix(str, "0x") {
		base = 16
	}
	str = strings.TrimPrefix(str, "0x")
	var zero T
	switch v := any(zero).(type) {
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

		value := mylog.Check2(strconv.ParseInt(str, base, 64))
		return any(value).(T)
	case int8:

		value := mylog.Check2(strconv.ParseInt(str, base, 8))
		return any(value).(T)
	case int16:

		value := mylog.Check2(strconv.ParseInt(str, base, 16))
		return any(value).(T)
	case int32:

		value := mylog.Check2(strconv.ParseInt(str, base, 32))
		return any(value).(T)
	case int64:

		value := mylog.Check2(strconv.ParseInt(str, base, 64))
		return any(value).(T)
	case uint:

		value := mylog.Check2(strconv.ParseUint(str, base, 64))
		return any(value).(T)
	case uint8:

		value := mylog.Check2(strconv.ParseUint(str, base, 8))
		return any(value).(T)
	case uint16:

		value := mylog.Check2(strconv.ParseUint(str, base, 16))
		return any(value).(T)
	case uint32:

		value := mylog.Check2(strconv.ParseUint(str, base, 32))
		return any(value).(T)
	case uint64:

		value := mylog.Check2(strconv.ParseUint(str, base, 64))
		return any(value).(T)
	case uintptr:

		value := mylog.Check2(strconv.ParseUint(str, base, 64))
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


`

	enum = `
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
`

	getSet = `
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

`
)
