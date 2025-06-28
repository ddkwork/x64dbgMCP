package main

import (
	"encoding/hex"
	"strconv"
	"strings"
)

type HexInt uint

func (h *HexInt) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)  // 去除 JSON 字符串的引号
	s = strings.TrimPrefix(s, "0x")       // 去掉 "0x" 前缀
	v, err := strconv.ParseInt(s, 16, 64) // 按十六进制解析
	*h = HexInt(v)
	return err
}

type HexBytes []byte

func (h *HexBytes) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	s = strings.TrimPrefix(s, "0x")
	decoded, err := hex.DecodeString(s)
	*h = decoded
	return err
}

type HexString string

func (h *HexString) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), `"`)
	s = strings.TrimPrefix(s, "0x")
	decoded, err := hex.DecodeString(s)
	*h = HexString(decoded)
	return err
}
