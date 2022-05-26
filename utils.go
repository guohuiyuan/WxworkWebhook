package wxworkwebhook

import "strings"

// IsNetFile 是否为网络文件
func IsNetFile(path string) bool {
	b, _, flag := strings.Cut(path, ":")
	return flag && (b == "http" || b == "https")
}
