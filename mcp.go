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

const (
	DefaultX64dbgServer = "http://127.0.0.1:8888/"
)

var x64dbgServerURL string

// X64DbgRPC 结构体包含所有调试方法
type X64DbgRPC struct {
	client *http.Client
}

func NewX64DbgRPC() *X64DbgRPC {
	return &X64DbgRPC{
		client: &http.Client{
			Timeout: 15 * time.Second,
			Transport: &http.Transport{
				DisableKeepAlives: true,
			},
		},
	}
}

// safeGet 发送 GET 请求并处理响应
func (x *X64DbgRPC) safeGet(endpoint string, params map[string]string) any {
	url := x64dbgServerURL + endpoint

	// 添加查询参数
	if len(params) > 0 {
		query := ""
		for key, value := range params {
			query += key + "=" + value + "&"
		}
		url += "?" + strings.TrimSuffix(query, "&")
	}

	resp, err := x.client.Get(url)
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

	resp, err := x.client.Post(url, "application/json", bytes.NewBuffer(payload))
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

func (x *X64DbgRPC) ExecCommand(req struct{ Cmd string }, res *string) error {
	result := x.safeGet("ExecCommand", map[string]string{"cmd": req.Cmd})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// DEBUGGING STATUS
// =============================================================================

func (x *X64DbgRPC) IsDebugActive(_ struct{}, res *bool) error {
	result := x.safeGet("IsDebugActive", nil)
	if str, ok := result.(string); ok {
		*res = strings.EqualFold(str, "true")
		return nil
	}
	*res = false
	return nil
}

func (x *X64DbgRPC) IsDebugging(_ struct{}, res *bool) error {
	result := x.safeGet("Is_Debugging", nil)
	if str, ok := result.(string); ok {
		*res = strings.EqualFold(str, "true")
		return nil
	}
	*res = false
	return nil
}

// =============================================================================
// REGISTER API
// =============================================================================

func (x *X64DbgRPC) RegisterGet(req struct{ Register string }, res *string) error {
	result := x.safeGet("Register/Get", map[string]string{"register": req.Register})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) RegisterSet(req struct{ Register, Value string }, res *string) error {
	result := x.safeGet("Register/Set", map[string]string{"register": req.Register, "value": req.Value})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// MEMORY API (Enhanced)
// =============================================================================

func (x *X64DbgRPC) MemoryRead(req struct{ Addr, Size string }, res *string) error {
	result := x.safeGet("Memory/Read", map[string]string{"addr": req.Addr, "size": req.Size})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) MemoryWrite(req struct{ Addr, Data string }, res *string) error {
	result := x.safeGet("Memory/Write", map[string]string{"addr": req.Addr, "data": req.Data})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) MemoryIsValidPtr(req struct{ Addr string }, res *bool) error {
	result := x.safeGet("Memory/IsValidPtr", map[string]string{"addr": req.Addr})
	if str, ok := result.(string); ok {
		*res = strings.EqualFold(str, "true")
		return nil
	}
	*res = false
	return nil
}

func (x *X64DbgRPC) MemoryGetProtect(req struct{ Addr string }, res *string) error {
	result := x.safeGet("Memory/GetProtect", map[string]string{"addr": req.Addr})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// DEBUG API
// =============================================================================

func (x *X64DbgRPC) DebugRun(_ struct{}, res *string) error {
	result := x.safeGet("Debug/Run", nil)
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) DebugPause(_ struct{}, res *string) error {
	result := x.safeGet("Debug/Pause", nil)
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) DebugStop(_ struct{}, res *string) error {
	result := x.safeGet("Debug/Stop", nil)
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) DebugStepIn(_ struct{}, res *string) error {
	result := x.safeGet("Debug/StepIn", nil)
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) DebugStepOver(_ struct{}, res *string) error {
	result := x.safeGet("Debug/StepOver", nil)
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) DebugStepOut(_ struct{}, res *string) error {
	result := x.safeGet("Debug/StepOut", nil)
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) DebugSetBreakpoint(req struct{ Addr string }, res *string) error {
	result := x.safeGet("Debug/SetBreakpoint", map[string]string{"addr": req.Addr})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) DebugDeleteBreakpoint(req struct{ Addr string }, res *string) error {
	result := x.safeGet("Debug/DeleteBreakpoint", map[string]string{"addr": req.Addr})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// ASSEMBLER API
// =============================================================================

func (x *X64DbgRPC) AssemblerAssemble(req struct{ Addr, Instruction string }, res *map[string]any) error {
	result := x.safeGet("Assembler/Assemble", map[string]string{
		"addr":        req.Addr,
		"instruction": req.Instruction,
	})

	switch v := result.(type) {
	case map[string]any:
		*res = v
	case string:
		var data map[string]any
		if err := json.Unmarshal([]byte(v), &data); err != nil {
			*res = map[string]any{"error": "Failed to parse assembly result", "raw": v}
			return nil
		}
		*res = data
	default:
		*res = map[string]any{"error": "Unexpected response format", "raw": result}
	}
	return nil
}

func (x *X64DbgRPC) AssemblerAssembleMem(req struct{ Addr, Instruction string }, res *string) error {
	result := x.safeGet("Assembler/AssembleMem", map[string]string{
		"addr":        req.Addr,
		"instruction": req.Instruction,
	})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// STACK API
// =============================================================================

func (x *X64DbgRPC) StackPop(_ struct{}, res *string) error {
	result := x.safeGet("Stack/Pop", nil)
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) StackPush(req struct{ Value string }, res *string) error {
	result := x.safeGet("Stack/Push", map[string]string{"value": req.Value})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) StackPeek(req struct{ Offset string }, res *string) error {
	result := x.safeGet("Stack/Peek", map[string]string{"offset": req.Offset})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// FLAG API
// =============================================================================

func (x *X64DbgRPC) FlagGet(req struct{ Flag string }, res *bool) error {
	result := x.safeGet("Flag/Get", map[string]string{"flag": req.Flag})
	if str, ok := result.(string); ok {
		*res = strings.EqualFold(str, "true")
		return nil
	}
	*res = false
	return nil
}

func (x *X64DbgRPC) FlagSet(req struct {
	Flag  string
	Value bool
}, res *string) error {
	value := "false"
	if req.Value {
		value = "true"
	}
	result := x.safeGet("Flag/Set", map[string]string{"flag": req.Flag, "value": value})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// PATTERN API
// =============================================================================

func (x *X64DbgRPC) PatternFindMem(req struct{ Start, Size, Pattern string }, res *string) error {
	result := x.safeGet("Pattern/FindMem", map[string]string{
		"start":   req.Start,
		"size":    req.Size,
		"pattern": req.Pattern,
	})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// MISC API
// =============================================================================

func (x *X64DbgRPC) MiscParseExpression(req struct{ Expression string }, res *string) error {
	result := x.safeGet("Misc/ParseExpression", map[string]string{"expression": req.Expression})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) MiscRemoteGetProcAddress(req struct{ Module, API string }, res *string) error {
	result := x.safeGet("Misc/RemoteGetProcAddress", map[string]string{"module": req.Module, "api": req.API})
	*res = fmt.Sprintf("%v", result)
	return nil
}

// =============================================================================
// LEGACY COMPATIBILITY FUNCTIONS
// =============================================================================

func (x *X64DbgRPC) SetRegister(req struct{ Name, Value string }, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "r " + req.Name + "=" + req.Value}, res)
	return result
}

func (x *X64DbgRPC) MemRead(req struct{ Addr, Size string }, res *string) error {
	result := x.safeGet("MemRead", map[string]string{"addr": req.Addr, "size": req.Size})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) MemWrite(req struct{ Addr, Data string }, res *string) error {
	result := x.safeGet("MemWrite", map[string]string{"addr": req.Addr, "data": req.Data})
	*res = fmt.Sprintf("%v", result)
	return nil
}

func (x *X64DbgRPC) SetBreakpoint(req struct{ Addr string }, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "bp " + req.Addr}, res)
	return result
}

func (x *X64DbgRPC) DeleteBreakpoint(req struct{ Addr string }, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "bpc " + req.Addr}, res)
	return result
}

func (x *X64DbgRPC) Run(_ struct{}, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "run"}, res)
	return result
}

func (x *X64DbgRPC) Pause(_ struct{}, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "pause"}, res)
	return result
}

func (x *X64DbgRPC) StepIn(_ struct{}, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "sti"}, res)
	return result
}

func (x *X64DbgRPC) StepOver(_ struct{}, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "sto"}, res)
	return result
}

func (x *X64DbgRPC) StepOut(_ struct{}, res *string) error {
	result := x.ExecCommand(struct{ Cmd string }{Cmd: "rtr"}, res)
	return result
}

func (x *X64DbgRPC) GetCallStack(_ struct{}, res *[]map[string]any) error {
	var result string
	if err := x.ExecCommand(struct{ Cmd string }{Cmd: "k"}, &result); err != nil {
		return err
	}

	stack := []map[string]any{
		{"info": "Call stack information requested via command", "result": result},
	}

	*res = stack
	return nil
}

func (x *X64DbgRPC) Disassemble(req struct{ Addr string }, res *map[string]any) error {
	var result string
	if err := x.ExecCommand(struct{ Cmd string }{Cmd: "dis " + req.Addr}, &result); err != nil {
		return err
	}

	*res = map[string]any{
		"addr":           req.Addr,
		"command_result": result,
	}
	return nil
}

func (x *X64DbgRPC) DisasmGetInstruction(req struct{ Addr string }, res *map[string]any) error {
	result := x.safeGet("Disasm/GetInstruction", map[string]string{"addr": req.Addr})

	switch v := result.(type) {
	case map[string]any:
		*res = v
	case string:
		var data map[string]any
		if err := json.Unmarshal([]byte(v), &data); err != nil {
			*res = map[string]any{"error": "Failed to parse disassembly result", "raw": v}
			return nil
		}
		*res = data
	default:
		*res = map[string]any{"error": "Unexpected response format", "raw": result}
	}
	return nil
}

func (x *X64DbgRPC) DisasmGetInstructionRange(req struct {
	Addr  string
	Count int
}, res *[]map[string]any) error {
	result := x.safeGet("Disasm/GetInstructionRange", map[string]string{
		"addr":  req.Addr,
		"count": strconv.Itoa(req.Count),
	})

	switch v := result.(type) {
	case []map[string]any:
		*res = v
	case string:
		var data []map[string]any
		if err := json.Unmarshal([]byte(v), &data); err != nil {
			*res = []map[string]any{{"error": "Failed to parse disassembly result", "raw": v}}
			return nil
		}
		*res = data
	default:
		*res = []map[string]any{{"error": "Unexpected response format"}}
	}
	return nil
}

func (x *X64DbgRPC) DisasmGetInstructionAtRIP(_ struct{}, res *map[string]any) error {
	result := x.safeGet("Disasm/GetInstructionAtRIP", nil)

	switch v := result.(type) {
	case map[string]any:
		*res = v
	case string:
		var data map[string]any
		if err := json.Unmarshal([]byte(v), &data); err != nil {
			*res = map[string]any{"error": "Failed to parse disassembly result", "raw": v}
			return nil
		}
		*res = data
	default:
		*res = map[string]any{"error": "Unexpected response format"}
	}
	return nil
}

func (x *X64DbgRPC) StepInWithDisasm(_ struct{}, res *map[string]any) error {
	result := x.safeGet("Disasm/StepInWithDisasm", nil)

	switch v := result.(type) {
	case map[string]any:
		*res = v
	case string:
		var data map[string]any
		if err := json.Unmarshal([]byte(v), &data); err != nil {
			*res = map[string]any{"error": "Failed to parse step result", "raw": v}
			return nil
		}
		*res = data
	default:
		*res = map[string]any{"error": "Unexpected response format"}
	}
	return nil
}

func (x *X64DbgRPC) GetModuleList(_ struct{}, res *[]map[string]any) error {
	result := x.safeGet("GetModuleList", nil)

	switch v := result.(type) {
	case []map[string]any:
		*res = v
	case string:
		var data []map[string]any
		if err := json.Unmarshal([]byte(v), &data); err != nil {
			*res = []map[string]any{{"error": "Failed to parse module list", "raw": v}}
			return nil
		}
		*res = data
	default:
		*res = []map[string]any{{"error": "Unexpected response format"}}
	}
	return nil
}

func (x *X64DbgRPC) MemoryBase(req struct{ Addr string }, res *map[string]any) error {
	result := x.safeGet("MemoryBase", map[string]string{"addr": req.Addr})

	switch v := result.(type) {
	case map[string]any:
		*res = v
	case string:
		// 尝试解析逗号分隔的响应
		if strings.Contains(v, ",") {
			parts := strings.Split(v, ",")
			if len(parts) == 2 {
				*res = map[string]any{
					"base_address": strings.TrimSpace(parts[0]),
					"size":         strings.TrimSpace(parts[1]),
				}
				return nil
			}
		}

		// 尝试解析 JSON
		var data map[string]any
		if err := json.Unmarshal([]byte(v), &data); err == nil {
			*res = data
			return nil
		}

		// 默认返回原始响应
		*res = map[string]any{"raw_response": v}
	default:
		*res = map[string]any{"error": "Unexpected response format"}
	}
	return nil
}

func main() {
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
	rpcServer := NewX64DbgRPC()
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
