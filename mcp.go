package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ddkwork/golibrary/std/mylog"
)

const DefaultX64dbgServer = "http://127.0.0.1:8888/"

var x64dbgServerURL string

// X64DbgRPC 结构体包含所有调试方法
type X64DbgRPC struct{}

var client = &http.Client{
	Timeout: 15 * time.Second,
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
}

func NewX64DbgRPC() *X64DbgRPC {
	return &X64DbgRPC{}
}

func request[T any](endpoint string, params map[string]string) T {
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

	// 尝试解析为 JSON
	var jsonData T
	if err := json.Unmarshal(body, &jsonData); err == nil {
		return jsonData
	}
	str := strings.TrimSpace(string(body))
	if strings.EqualFold(str, "true") {
		return any(true)
	}
	if strings.EqualFold(str, "false") {
		return any(false)
	}
	// 返回纯文本
	return any(str)
}

// =============================================================================
// UNIFIED COMMAND EXECUTION
// =============================================================================
// todo 接收多个参数，返回泛型
func (x *X64DbgRPC) ExecCommand(cmd string) bool {
	return request[bool]("ExecCommand", map[string]string{"cmd": cmd})
}

// =============================================================================
// DEBUGGING STATUS
// =============================================================================

func (x *X64DbgRPC) IsDebugActive() bool {
	return request[bool]("IsDebugActive", nil)
}

func (x *X64DbgRPC) IsDebugging() bool {
	return request[bool]("Is_Debugging", nil)
}

// =============================================================================
// REGISTER API
// =============================================================================

func (x *X64DbgRPC) RegisterGet(Register string) bool {
	return request[bool]("Register/Get", map[string]string{"register": Register})
}

func (x *X64DbgRPC) RegisterSet(Register, Value string) bool {
	return request[bool]("Register/Set", map[string]string{"register": Register, "value": Value})
}

// =============================================================================
// MEMORY API (Enhanced)
// =============================================================================

func (x *X64DbgRPC) MemoryRead(Addr, Size string) bool {
	return request[bool]("Memory/Read", map[string]string{"addr": Addr, "size": Size})
}

func (x *X64DbgRPC) MemoryWrite(Addr, Data string) bool {
	return request[bool]("Memory/Write", map[string]string{"addr": Addr, "data": Data})
}

func (x *X64DbgRPC) MemoryIsValidPtr(Addr string) bool {
	return request[bool]("Memory/IsValidPtr", map[string]string{"addr": Addr})
}

func (x *X64DbgRPC) MemoryGetProtect(Addr string) bool {
	return request[bool]("Memory/GetProtect", map[string]string{"addr": Addr})
}

// =============================================================================
// DEBUG API
// =============================================================================

func (x *X64DbgRPC) DebugRun() bool {
	return request[bool]("Debug/Run", nil)
}

func (x *X64DbgRPC) DebugPause() bool {
	return request[bool]("Debug/Pause", nil)
}

func (x *X64DbgRPC) DebugStop() bool {
	return request[bool]("Debug/Stop", nil)
}

func (x *X64DbgRPC) DebugStepIn() bool {
	return request[bool]("Debug/StepIn", nil)
}

func (x *X64DbgRPC) DebugStepOver() bool {
	return request[bool]("Debug/StepOver", nil)
}

func (x *X64DbgRPC) DebugStepOut() bool {
	return request[bool]("Debug/StepOut", nil)
}

func (x *X64DbgRPC) DebugSetBreakpoint(Addr string) bool {
	return request[bool]("Debug/SetBreakpoint", map[string]string{"addr": Addr})
}

func (x *X64DbgRPC) DebugDeleteBreakpoint(Addr string) bool {
	return request[bool]("Debug/DeleteBreakpoint", map[string]string{"addr": Addr})
}

// =============================================================================
// ASSEMBLER API
// =============================================================================

func (x *X64DbgRPC) AssemblerAssemble(Addr, Instruction string) bool {
	return request[bool]("Assembler/Assemble", map[string]string{
		"addr":        Addr,
		"instruction": Instruction,
	})
}

func (x *X64DbgRPC) AssemblerAssembleMem(Addr, Instruction string) bool {
	return request[bool]("Assembler/AssembleMem", map[string]string{
		"addr":        Addr,
		"instruction": Instruction,
	})
}

// =============================================================================
// STACK API
// =============================================================================

func (x *X64DbgRPC) StackPop() bool {
	return request[bool]("Stack/Pop", nil)
}

func (x *X64DbgRPC) StackPush(Value string) bool {
	return request[bool]("Stack/Push", map[string]string{"value": Value})
}

func (x *X64DbgRPC) StackPeek(Offset string) bool {
	return request[bool]("Stack/Peek", map[string]string{"offset": Offset})
}

// =============================================================================
// FLAG API
// =============================================================================

func (x *X64DbgRPC) FlagGet(Flag string) bool {
	return request[bool]("Flag/Get", map[string]string{"flag": Flag})
}

func (x *X64DbgRPC) FlagSet(Flag string, Value bool) bool {
	value := "false"
	if Value {
		value = "true"
	}
	return request[bool]("Flag/Set", map[string]string{"flag": Flag, "value": value})
}

// =============================================================================
// PATTERN API
// =============================================================================

func (x *X64DbgRPC) PatternFindMem(Start, Size, Pattern string) bool {
	return request[bool]("Pattern/FindMem", map[string]string{
		"start":   Start,
		"size":    Size,
		"pattern": Pattern,
	})
}

// =============================================================================
// MISC API
// =============================================================================

func (x *X64DbgRPC) MiscParseExpression(Expression string) bool {
	return request[bool]("Misc/ParseExpression", map[string]string{"expression": Expression})
}

func (x *X64DbgRPC) MiscRemoteGetProcAddress(Module, API string) bool {
	return request[bool]("Misc/RemoteGetProcAddress", map[string]string{"module": Module, "api": API})
}

// =============================================================================
// LEGACY COMPATIBILITY FUNCTIONS
// =============================================================================

func (x *X64DbgRPC) SetRegister(Name, Value string) bool {
	return x.ExecCommand("r " + Name + "=" + Value)
}

func (x *X64DbgRPC) MemRead(Addr, Size string) bool {
	return request[bool]("MemRead", map[string]string{"addr": Addr, "size": Size})
}

func (x *X64DbgRPC) MemWrite(Addr, Data string) bool {
	return request[bool]("MemWrite", map[string]string{"addr": Addr, "data": Data})
}

func (x *X64DbgRPC) SetBreakpoint(Addr string) bool {
	return x.ExecCommand("bp " + Addr)
}

func (x *X64DbgRPC) DeleteBreakpoint(Addr string) bool {
	return x.ExecCommand("bpc " + Addr)
}

func (x *X64DbgRPC) Run() bool {
	return x.ExecCommand("run")
}

func (x *X64DbgRPC) Pause() bool {
	return x.ExecCommand("pause")
}

func (x *X64DbgRPC) StepIn() bool {
	return x.ExecCommand("sti")
}

func (x *X64DbgRPC) StepOver() bool {
	return x.ExecCommand("sto")
}

func (x *X64DbgRPC) StepOut() bool {
	return x.ExecCommand("rtr")
}

func (x *X64DbgRPC) GetCallStack() {
	//var result string
	//if err := x.ExecCommand(struct{ Cmd string }{Cmd: "k"}, &result); err != nil {
	//	return err
	//}
	//
	//stack := []map[string]any{
	//	{"info": "Call stack information requested via command", "result": result},
	//}
	//
	//return nil
}

func (x *X64DbgRPC) Disassemble(Addr string) {
	//var result string
	//if err := x.ExecCommand(struct{ Cmd string }{Cmd: "dis " + Addr}, &result); err != nil {
	//	return err
	//}
	//
	//*res = map[string]any{
	//	"addr":           Addr,
	//	"command_result": result,
	//}
	//return nil
}

func (x *X64DbgRPC) DisasmGetInstruction(Addr string) bool {
	return request[bool]("Disasm/GetInstruction", map[string]string{"addr": Addr})
}

func (x *X64DbgRPC) DisasmGetInstructionRange(Addr string, Count int) bool {
	return request[bool]("Disasm/GetInstructionRange", map[string]string{
		"addr":  Addr,
		"count": strconv.Itoa(Count),
	})
}

func (x *X64DbgRPC) DisasmGetInstructionAtRIP() bool {
	return request[bool]("Disasm/GetInstructionAtRIP", nil)
}

func (x *X64DbgRPC) StepInWithDisasm() bool {
	return request[bool]("Disasm/StepInWithDisasm", nil)
}

func (x *X64DbgRPC) GetModuleList() bool {
	return request[bool]("GetModuleList", nil)
}

func (x *X64DbgRPC) MemoryBase(Addr string) bool {
	return request[bool]("MemoryBase", map[string]string{"addr": Addr})

	//switch v := result.(type) {
	//case map[string]any:
	//	*res = v
	//case string:
	//	// 尝试解析逗号分隔的响应
	//	if strings.Contains(v, ",") {
	//		parts := strings.Split(v, ",")
	//		if len(parts) == 2 {
	//			*res = map[string]any{
	//				"base_address": strings.TrimSpace(parts[0]),
	//				"size":         strings.TrimSpace(parts[1]),
	//			}
	//			return nil
	//		}
	//	}
	//
	//	// 尝试解析 JSON
	//	var data map[string]any
	//	if err := json.Unmarshal([]byte(v), &data); err == nil {
	//		*res = data
	//		return nil
	//	}
	//
	//default:
	//}
}

func main() {
	d := NewX64DbgRPC()
	d.ExecCommand("restartadmin")

	return
	// 解析命令行参数
	serverPtr := flag.String("server", "", "x64dbg server URL")
	portPtr := flag.String("port", "8889", "RPC server port")
	flag.Parse()

	// 设置x64dbg服务器URL
	if *serverPtr != "" {
		x64dbgServerURL = *serverPtr
	} else if len(os.Args) > 1 {
		x64dbgServerURL = os.Args[1]
	} else {
		x64dbgServerURL = DefaultX64dbgServer
	}

	// 确保URL以/结尾
	if !strings.HasSuffix(x64dbgServerURL, "/") {
		x64dbgServerURL += "/"
	}

	// 初始化RPC服务
	rpcServer := d
	mylog.Check(rpc.Register(rpcServer))
	rpc.HandleHTTP()

	// 启动服务
	port := *portPtr
	log.Printf("Starting RPC server on port %s for x64dbg server at %s\n", port, x64dbgServerURL)
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	mylog.Check(http.Serve(listener, nil))
}
