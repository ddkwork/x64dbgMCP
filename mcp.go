package main

import (
	"bytes"
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

// safeGet 发送 GET 请求并处理响应
func safeGet[T any](endpoint string, params map[string]string) T {
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

// safePost 发送 POST 请求并处理响应
func (x *X64DbgRPC) safePost(endpoint string, data any) any {
	url := x64dbgServerURL + endpoint

	var payload []byte
	switch d := data.(type) {
	case string:
		payload = []byte(d)
	default:
		var err error
		payload, err = json.Marshal(data)
		if err != nil {
			return fmt.Sprintf("Failed to marshal data: %v", err)
		}
	}

	resp, err := client.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Sprintf("Request failed: %v", err)
	}
	defer func() {
		mylog.Check(resp.Body.Close())
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Error %d: %s", resp.StatusCode, string(body))
	}

	// 尝试解析为 JSON
	var jsonData any
	if err := json.Unmarshal(body, &jsonData); err == nil {
		return jsonData
	}

	// 返回纯文本
	return strings.TrimSpace(string(body))
}

// =============================================================================
// UNIFIED COMMAND EXECUTION
// =============================================================================
// todo 接收多个参数，返回泛型
func (x *X64DbgRPC) ExecCommand(cmd string) bool {
	return safeGet[bool]("ExecCommand", map[string]string{"cmd": cmd})
}

// =============================================================================
// DEBUGGING STATUS
// =============================================================================

func (x *X64DbgRPC) IsDebugActive() bool {
	return safeGet[bool]("IsDebugActive", nil)
}

func (x *X64DbgRPC) IsDebugging() bool {
	return safeGet[bool]("Is_Debugging", nil)
}

// =============================================================================
// REGISTER API
// =============================================================================

func (x *X64DbgRPC) RegisterGet(Register string) bool {
	return safeGet[bool]("Register/Get", map[string]string{"register": Register})
}

func (x *X64DbgRPC) RegisterSet(Register, Value string) bool {
	return safeGet[bool]("Register/Set", map[string]string{"register": Register, "value": Value})
}

// =============================================================================
// MEMORY API (Enhanced)
// =============================================================================

func (x *X64DbgRPC) MemoryRead(Addr, Size string) bool {
	return safeGet[bool]("Memory/Read", map[string]string{"addr": Addr, "size": Size})
}

func (x *X64DbgRPC) MemoryWrite(Addr, Data string) bool {
	return safeGet[bool]("Memory/Write", map[string]string{"addr": Addr, "data": Data})
}

func (x *X64DbgRPC) MemoryIsValidPtr(Addr string) bool {
	return safeGet[bool]("Memory/IsValidPtr", map[string]string{"addr": Addr})
}

func (x *X64DbgRPC) MemoryGetProtect(Addr string) bool {
	return safeGet[bool]("Memory/GetProtect", map[string]string{"addr": Addr})
}

// =============================================================================
// DEBUG API
// =============================================================================

func (x *X64DbgRPC) DebugRun() bool {
	return safeGet[bool]("Debug/Run", nil)
}

func (x *X64DbgRPC) DebugPause() bool {
	return safeGet[bool]("Debug/Pause", nil)
}

func (x *X64DbgRPC) DebugStop() bool {
	return safeGet[bool]("Debug/Stop", nil)
}

func (x *X64DbgRPC) DebugStepIn() bool {
	return safeGet[bool]("Debug/StepIn", nil)
}

func (x *X64DbgRPC) DebugStepOver() bool {
	return safeGet[bool]("Debug/StepOver", nil)
}

func (x *X64DbgRPC) DebugStepOut() bool {
	return safeGet[bool]("Debug/StepOut", nil)
}

func (x *X64DbgRPC) DebugSetBreakpoint(Addr string) bool {
	return safeGet[bool]("Debug/SetBreakpoint", map[string]string{"addr": Addr})
}

func (x *X64DbgRPC) DebugDeleteBreakpoint(Addr string) bool {
	return safeGet[bool]("Debug/DeleteBreakpoint", map[string]string{"addr": Addr})
}

// =============================================================================
// ASSEMBLER API
// =============================================================================

func (x *X64DbgRPC) AssemblerAssemble(Addr, Instruction string) bool {
	return safeGet[bool]("Assembler/Assemble", map[string]string{
		"addr":        Addr,
		"instruction": Instruction,
	})
}

func (x *X64DbgRPC) AssemblerAssembleMem(Addr, Instruction string) bool {
	return safeGet[bool]("Assembler/AssembleMem", map[string]string{
		"addr":        Addr,
		"instruction": Instruction,
	})
}

// =============================================================================
// STACK API
// =============================================================================

func (x *X64DbgRPC) StackPop() bool {
	return safeGet[bool]("Stack/Pop", nil)
}

func (x *X64DbgRPC) StackPush(Value string) bool {
	return safeGet[bool]("Stack/Push", map[string]string{"value": Value})
}

func (x *X64DbgRPC) StackPeek(Offset string) bool {
	return safeGet[bool]("Stack/Peek", map[string]string{"offset": Offset})
}

// =============================================================================
// FLAG API
// =============================================================================

func (x *X64DbgRPC) FlagGet(Flag string) bool {
	return safeGet[bool]("Flag/Get", map[string]string{"flag": Flag})
}

func (x *X64DbgRPC) FlagSet(Flag string, Value bool) bool {
	value := "false"
	if Value {
		value = "true"
	}
	return safeGet[bool]("Flag/Set", map[string]string{"flag": Flag, "value": value})
}

// =============================================================================
// PATTERN API
// =============================================================================

func (x *X64DbgRPC) PatternFindMem(Start, Size, Pattern string) bool {
	return safeGet[bool]("Pattern/FindMem", map[string]string{
		"start":   Start,
		"size":    Size,
		"pattern": Pattern,
	})
}

// =============================================================================
// MISC API
// =============================================================================

func (x *X64DbgRPC) MiscParseExpression(Expression string) bool {
	return safeGet[bool]("Misc/ParseExpression", map[string]string{"expression": Expression})
}

func (x *X64DbgRPC) MiscRemoteGetProcAddress(Module, API string) bool {
	return safeGet[bool]("Misc/RemoteGetProcAddress", map[string]string{"module": Module, "api": API})
}

// =============================================================================
// LEGACY COMPATIBILITY FUNCTIONS
// =============================================================================

func (x *X64DbgRPC) SetRegister(Name, Value string) bool {
	return x.ExecCommand("r " + Name + "=" + Value)
}

func (x *X64DbgRPC) MemRead(Addr, Size string) bool {
	return safeGet[bool]("MemRead", map[string]string{"addr": Addr, "size": Size})
}

func (x *X64DbgRPC) MemWrite(Addr, Data string) bool {
	return safeGet[bool]("MemWrite", map[string]string{"addr": Addr, "data": Data})
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
	return safeGet[bool]("Disasm/GetInstruction", map[string]string{"addr": Addr})
}

func (x *X64DbgRPC) DisasmGetInstructionRange(Addr string, Count int) bool {
	return safeGet[bool]("Disasm/GetInstructionRange", map[string]string{
		"addr":  Addr,
		"count": strconv.Itoa(Count),
	})
}

func (x *X64DbgRPC) DisasmGetInstructionAtRIP() bool {
	return safeGet[bool]("Disasm/GetInstructionAtRIP", nil)
}

func (x *X64DbgRPC) StepInWithDisasm() bool {
	return safeGet[bool]("Disasm/StepInWithDisasm", nil)
}

func (x *X64DbgRPC) GetModuleList() bool {
	return safeGet[bool]("GetModuleList", nil)
}

func (x *X64DbgRPC) MemoryBase(Addr string) bool {
	return safeGet[bool]("MemoryBase", map[string]string{"addr": Addr})

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
