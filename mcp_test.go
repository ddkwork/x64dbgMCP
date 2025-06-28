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
			//func (m *RegisterManager) SetDR0(v HexInt) bool {
			//	return request[bool](tagSetRegister, map[string]string{"register": "DR0", "value": strconv.FormatHexInt(HexInt64(v), 16)})
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
						": strconv.FormatHexInt(HexInt64(v), 16)})")

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
//Get(reg RegisterEnum) HexInt
//Set(reg RegisterEnum, value HexInt) bool
//Size() HexInt
GetDR0() HexInt
SetDR0(v HexInt) bool
GetDR1() HexInt
SetDR1(v HexInt) bool
GetDR2() HexInt
SetDR2(v HexInt) bool
GetDR3() HexInt
SetDR3(v HexInt) bool
GetDR6() HexInt
SetDR6(v HexInt) bool
GetDR7() HexInt
SetDR7(v HexInt) bool
GetEAX() HexInt32
SetEAX(v HexInt32) bool
GetAX() HexInt16
SetAX(v HexInt16) bool
GetAH() HexInt8
SetAH(v HexInt8) bool
GetAL() HexInt8
SetAL(v HexInt8) bool
GetEBX() HexInt32
SetEBX(v HexInt32) bool
GetBX() HexInt16
SetBX(v HexInt16) bool
GetBH() HexInt8
SetBH(v HexInt8) bool
GetBL() HexInt8
SetBL(v HexInt8) bool
GetECX() HexInt32
SetECX(v HexInt32) bool
GetCX() HexInt16
SetCX(v HexInt16) bool
GetCH() HexInt8
SetCH(v HexInt8) bool
GetCL() HexInt8
SetCL(v HexInt8) bool
GetEDX() HexInt32
SetEDX(v HexInt32) bool
GetDX() HexInt16
SetDX(v HexInt16) bool
GetDH() HexInt8
SetDH(v HexInt8) bool
GetDL() HexInt8
SetDL(v HexInt8) bool
GetEDI() HexInt32
SetEDI(v HexInt32) bool
GetDI() HexInt16
SetDI(v HexInt16) bool
GetESI() HexInt32
SetESI(v HexInt32) bool
GetSI() HexInt16
SetSI(v HexInt16) bool
GetEBP() HexInt32
SetEBP(v HexInt32) bool
GetBP() HexInt16
SetBP(v HexInt16) bool
GetESP() HexInt32
SetESP(v HexInt32) bool
GetSP() HexInt16
SetSP(v HexInt16) bool
GetEIP() HexInt32
SetEIP(v HexInt32) bool
GetRAX() HexInt64
SetRAX(v HexInt64) bool
GetRBX() HexInt64
SetRBX(v HexInt64) bool
GetRCX() HexInt64
SetRCX(v HexInt64) bool
GetRDX() HexInt64
SetRDX(v HexInt64) bool
GetRSI() HexInt64
SetRSI(v HexInt64) bool
GetSIL() HexInt8
SetSIL(v HexInt8) bool
GetRDI() HexInt64
SetRDI(v HexInt64) bool
GetDIL() HexInt8
SetDIL(v HexInt8) bool
GetRBP() HexInt64
SetRBP(v HexInt64) bool
GetBPL() HexInt8
SetBPL(v HexInt8) bool
GetRSP() HexInt64
SetRSP(v HexInt64) bool
GetSPL() HexInt8
SetSPL(v HexInt8) bool
GetRIP() HexInt64
SetRIP(v HexInt64) bool
GetR8() HexInt64
SetR8(v HexInt64) bool
GetR8D() HexInt32
SetR8D(v HexInt32) bool
GetR8W() HexInt16
SetR8W(v HexInt16) bool
GetR8B() HexInt8
SetR8B(v HexInt8) bool
GetR9() HexInt64
SetR9(v HexInt64) bool
GetR9D() HexInt32
SetR9D(v HexInt32) bool
GetR9W() HexInt16
SetR9W(v HexInt16) bool
GetR9B() HexInt8
SetR9B(v HexInt8) bool
GetR10() HexInt64
SetR10(v HexInt64) bool
GetR10D() HexInt32
SetR10D(v HexInt32) bool
GetR10W() HexInt16
SetR10W(v HexInt16) bool
GetR10B() HexInt8
SetR10B(v HexInt8) bool
GetR11() HexInt64
SetR11(v HexInt64) bool
GetR11D() HexInt32
SetR11D(v HexInt32) bool
GetR11W() HexInt16
SetR11W(v HexInt16) bool
GetR11B() HexInt8
SetR11B(v HexInt8) bool
GetR12() HexInt64
SetR12(v HexInt64) bool
GetR12D() HexInt32
SetR12D(v HexInt32) bool
GetR12W() HexInt16
SetR12W(v HexInt16) bool
GetR12B() HexInt8
SetR12B(v HexInt8) bool
GetR13() HexInt64
SetR13(v HexInt64) bool
GetR13D() HexInt32
SetR13D(v HexInt32) bool
GetR13W() HexInt16
SetR13W(v HexInt16) bool
GetR13B() HexInt8
SetR13B(v HexInt8) bool
GetR14() HexInt64
SetR14(v HexInt64) bool
GetR14D() HexInt32
SetR14D(v HexInt32) bool
GetR14W() HexInt16
SetR14W(v HexInt16) bool
GetR14B() HexInt8
SetR14B(v HexInt8) bool
GetR15() HexInt64
SetR15(v HexInt64) bool
GetR15D() HexInt32
SetR15D(v HexInt32) bool
GetR15W() HexInt16
SetR15W(v HexInt16) bool
GetR15B() HexInt8
SetR15B(v HexInt8) bool
GetCIP() HexInt
SetCIP(v HexInt) bool
GetCSP() HexInt
SetCSP(v HexInt) bool
GetCAX() HexInt
SetCAX(v HexInt) bool
GetCBX() HexInt
SetCBX(v HexInt) bool
GetCCX() HexInt
SetCCX(v HexInt) bool
GetCDX() HexInt
SetCDX(v HexInt) bool
GetCDI() HexInt
SetCDI(v HexInt) bool
GetCSI() HexInt
SetCSI(v HexInt) bool
GetCBP() HexInt
SetCBP(v HexInt) bool
GetCFLAGS() HexInt
SetCFLAGS(v HexInt) bool 
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
//			~HexInt | ~HexInt8 | ~HexInt16 | ~HexInt32 | ~HexInt64 | ~HexIntptr |
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
		void|
HexInt|
HexInt8|
HexInt16|
HexInt32|
HexInt64|
HexIntptr|
float32|	
float64|
HexUint|
HexUint8|
HexUint16|
HexUint32|
HexUint64|
HexUintptr|
string|
HexBytes|
HexString

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
	case HexInt:

		value := mylog.Check2(strconv.ParseHexInt(str, base, 64))
		return any(value).(T)
	case HexInt8:

		value := mylog.Check2(strconv.ParseHexInt(str, base, 8))
		return any(value).(T)
	case HexInt16:

		value := mylog.Check2(strconv.ParseHexInt(str, base, 16))
		return any(value).(T)
	case HexInt32:

		value := mylog.Check2(strconv.ParseHexInt(str, base, 32))
		return any(value).(T)
	case HexInt64:

		value := mylog.Check2(strconv.ParseHexInt(str, base, 64))
		return any(value).(T)
	case HexIntptr:

		value := mylog.Check2(strconv.ParseHexInt(str, base, 64))
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
func (m RegisterManager) Get(reg RegisterEnum) HexInt {
	switch reg {
	case DR0:
		return HexInt(m.GetDR0())
	case DR1:
		return HexInt(m.GetDR1())
	case DR2:
		return HexInt(m.GetDR2())
	case DR3:
		return HexInt(m.GetDR3())
	case DR6:
		return HexInt(m.GetDR6())
	case DR7:
		return HexInt(m.GetDR7())
	case EAX:
		return HexInt(m.GetEAX())
	case AX:
		return HexInt(m.GetAX())
	case AH:
		return HexInt(m.GetAH())
	case AL:
		return HexInt(m.GetAL())
	case EBX:
		return HexInt(m.GetEBX())
	case BX:
		return HexInt(m.GetBX())
	case BH:
		return HexInt(m.GetBH())
	case BL:
		return HexInt(m.GetBL())
	case ECX:
		return HexInt(m.GetECX())
	case CX:
		return HexInt(m.GetCX())
	case CH:
		return HexInt(m.GetCH())
	case CL:
		return HexInt(m.GetCL())
	case EDX:
		return HexInt(m.GetEDX())
	case DX:
		return HexInt(m.GetDX())
	case DH:
		return HexInt(m.GetDH())
	case DL:
		return HexInt(m.GetDL())
	case EDI:
		return HexInt(m.GetEDI())
	case DI:
		return HexInt(m.GetDI())
	case ESI:
		return HexInt(m.GetESI())
	case SI:
		return HexInt(m.GetSI())
	case EBP:
		return HexInt(m.GetEBP())
	case BP:
		return HexInt(m.GetBP())
	case ESP:
		return HexInt(m.GetESP())
	case SP:
		return HexInt(m.GetSP())
	case EIP:
		return HexInt(m.GetEIP())
	case RAX:
		return HexInt(m.GetRAX())
	case RBX:
		return HexInt(m.GetRBX())
	case RCX:
		return HexInt(m.GetRCX())
	case RDX:
		return HexInt(m.GetRDX())
	case RSI:
		return HexInt(m.GetRSI())
	case SIL:
		return HexInt(m.GetSIL())
	case RDI:
		return HexInt(m.GetRDI())
	case DIL:
		return HexInt(m.GetDIL())
	case RBP:
		return HexInt(m.GetRBP())
	case BPL:
		return HexInt(m.GetBPL())
	case RSP:
		return HexInt(m.GetRSP())
	case SPL:
		return HexInt(m.GetSPL())
	case RIP:
		return HexInt(m.GetRIP())
	case R8:
		return HexInt(m.GetR8())
	case R8D:
		return HexInt(m.GetR8D())
	case R8W:
		return HexInt(m.GetR8W())
	case R8B:
		return HexInt(m.GetR8B())
	case R9:
		return HexInt(m.GetR9())
	case R9D:
		return HexInt(m.GetR9D())
	case R9W:
		return HexInt(m.GetR9W())
	case R9B:
		return HexInt(m.GetR9B())
	case R10:
		return HexInt(m.GetR10())
	case R10D:
		return HexInt(m.GetR10D())
	case R10W:
		return HexInt(m.GetR10W())
	case R10B:
		return HexInt(m.GetR10B())
	case R11:
		return HexInt(m.GetR11())
	case R11D:
		return HexInt(m.GetR11D())
	case R11W:
		return HexInt(m.GetR11W())
	case R11B:
		return HexInt(m.GetR11B())
	case R12:
		return HexInt(m.GetR12())
	case R12D:
		return HexInt(m.GetR12D())
	case R12W:
		return HexInt(m.GetR12W())
	case R12B:
		return HexInt(m.GetR12B())
	case R13:
		return HexInt(m.GetR13())
	case R13D:
		return HexInt(m.GetR13D())
	case R13W:
		return HexInt(m.GetR13W())
	case R13B:
		return HexInt(m.GetR13B())
	case R14:
		return HexInt(m.GetR14())
	case R14D:
		return HexInt(m.GetR14D())
	case R14W:
		return HexInt(m.GetR14W())
	case R14B:
		return HexInt(m.GetR14B())
	case R15:
		return HexInt(m.GetR15())
	case R15D:
		return HexInt(m.GetR15D())
	case R15W:
		return HexInt(m.GetR15W())
	case R15B:
		return HexInt(m.GetR15B())
	case CIP:
		return HexInt(m.GetCIP())
	case CSP:
		return HexInt(m.GetCSP())
	case CAX:
		return HexInt(m.GetCAX())
	case CBX:
		return HexInt(m.GetCBX())
	case CCX:
		return HexInt(m.GetCCX())
	case CDX:
		return HexInt(m.GetCDX())
	case CDI:
		return HexInt(m.GetCDI())
	case CSI:
		return HexInt(m.GetCSI())
	case CBP:
		return HexInt(m.GetCBP())
	case CFLAGS:
		return HexInt(m.GetCFLAGS())
	default:
		panic("Invalid register enum")
	}
}

func (m RegisterManager) Set(reg RegisterEnum, value HexInt) bool {
	switch reg {
	case DR0:
		return m.SetDR0(HexInt(value))
	case DR1:
		return m.SetDR1(HexInt(value))
	case DR2:
		return m.SetDR2(HexInt(value))
	case DR3:
		return m.SetDR3(HexInt(value))
	case DR6:
		return m.SetDR6(HexInt((value)))
	case DR7:
		return m.SetDR7(HexInt(value))
	case EAX:
		return m.SetEAX(HexInt32(value))
	case AX:
		return m.SetAX(HexInt16(value))
	case AH:
		return m.SetAH(HexInt8(value))
	case AL:
		return m.SetAL(HexInt8(value))
	case EBX:
		return m.SetEBX(HexInt32(value))
	case BX:
		return m.SetBX(HexInt16(value))
	case BH:
		return m.SetBH(HexInt8(value))
	case BL:
		return m.SetBL(HexInt8(value))
	case ECX:
		return m.SetECX(HexInt32(value))
	case CX:
		return m.SetCX(HexInt16(value))
	case CH:
		return m.SetCH(HexInt8(value))
	case CL:
		return m.SetCL(HexInt8(value))
	case EDX:
		return m.SetEDX(HexInt32(value))
	case DX:
		return m.SetDX(HexInt16(value))
	case DH:
		return m.SetDH(HexInt8(value))
	case DL:
		return m.SetDL(HexInt8(value))
	case EDI:
		return m.SetEDI(HexInt32(value))
	case DI:
		return m.SetDI(HexInt16(value))
	case ESI:
		return m.SetESI(HexInt32(value))
	case SI:
		return m.SetSI(HexInt16(value))
	case EBP:
		return m.SetEBP(HexInt32(value))
	case BP:
		return m.SetBP(HexInt16(value))
	case ESP:
		return m.SetESP(HexInt32(value))
	case SP:
		return m.SetSP(HexInt16(value))
	case EIP:
		return m.SetEIP(HexInt32(value))
	case RAX:
		return m.SetRAX(HexInt64(value))
	case RBX:
		return m.SetRBX(HexInt64(value))
	case RCX:
		return m.SetRCX(HexInt64(value))
	case RDX:
		return m.SetRDX(HexInt64(value))
	case RSI:
		return m.SetRSI(HexInt64(value))
	case SIL:
		return m.SetSIL(HexInt8(value))
	case RDI:
		return m.SetRDI(HexInt64(value))
	case DIL:
		return m.SetDIL(HexInt8(value))
	case RBP:
		return m.SetRBP(HexInt64(value))
	case BPL:
		return m.SetBPL(HexInt8(value))
	case RSP:
		return m.SetRSP(HexInt64(value))
	case SPL:
		return m.SetSPL(HexInt8(value))
	case RIP:
		return m.SetRIP(HexInt64(value))
	case R8:
		return m.SetR8(HexInt64(value))
	case R8D:
		return m.SetR8D(HexInt32(value))
	case R8W:
		return m.SetR8W(HexInt16(value))
	case R8B:
		return m.SetR8B(HexInt8(value))
	case R9:
		return m.SetR9(HexInt64(value))
	case R9D:
		return m.SetR9D(HexInt32(value))
	case R9W:
		return m.SetR9W(HexInt16(value))
	case R9B:
		return m.SetR9B(HexInt8(value))
	case R10:
		return m.SetR10(HexInt64(value))
	case R10D:
		return m.SetR10D(HexInt32(value))
	case R10W:
		return m.SetR10W(HexInt16(value))
	case R10B:
		return m.SetR10B(HexInt8(value))
	case R11:
		return m.SetR11(HexInt64(value))
	case R11D:
		return m.SetR11D(HexInt32(value))
	case R11W:
		return m.SetR11W(HexInt16(value))
	case R11B:
		return m.SetR11B(HexInt8(value))
	case R12:
		return m.SetR12(HexInt64(value))
	case R12D:
		return m.SetR12D(HexInt32(value))
	case R12W:
		return m.SetR12W(HexInt16(value))
	case R12B:
		return m.SetR12B(HexInt8(value))
	case R13:
		return m.SetR13(HexInt64(value))
	case R13D:
		return m.SetR13D(HexInt32(value))
	case R13W:
		return m.SetR13W(HexInt16(value))
	case R13B:
		return m.SetR13B(HexInt8(value))
	case R14:
		return m.SetR14(HexInt64(value))
	case R14D:
		return m.SetR14D(HexInt32(value))
	case R14W:
		return m.SetR14W(HexInt16(value))
	case R14B:
		return m.SetR14B(HexInt8(value))
	case R15:
		return m.SetR15(HexInt64(value))
	case R15D:
		return m.SetR15D(HexInt32(value))
	case R15W:
		return m.SetR15W(HexInt16(value))
	case R15B:
		return m.SetR15B(HexInt8(value))
	case CIP:
		return m.SetCIP(HexInt(value))
	case CSP:
		return m.SetCSP(HexInt(value))
	case CAX:
		return m.SetCAX(HexInt(value))
	case CBX:
		return m.SetCBX(HexInt(value))
	case CCX:
		return m.SetCCX(HexInt(value))
	case CDX:
		return m.SetCDX(HexInt(value))
	case CDI:
		return m.SetCDI(HexInt(value))
	case CSI:
		return m.SetCSI(HexInt(value))
	case CBP:
		return m.SetCBP(HexInt(value))
	case CFLAGS:
		return m.SetCFLAGS(HexInt(value))
	default:
		panic("Invalid register enum")
	}
}

`
)
