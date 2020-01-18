package utils

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	WorkDir     string
	ResourceDir string
)

func init() {
	WorkDir = PathAppRunning()
	ResourceDir = filepath.Join(WorkDir, "resource")
}

// PathAppRunning 当前程序运行物理路径
func PathAppRunning() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// PathIsDir 判断所给路径是否为文件夹
func PathIsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// PathIsFile 判断所给路径是否为文件
func PathIsFile(path string) bool {
	return !PathIsDir(path)
}
