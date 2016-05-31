package gofile

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/bufio.v1"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

//将Json内容写入文件
func Json2File(content, path string, perm os.FileMode) error {
	if FileExists(path) {
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	bytes := []byte(content)
	return IoUtilWriteFile(path, bytes, perm)
}

//将Json内容从文件读取并解析
func File2Json(path string, v interface{}) error {
	if FileExists(path) {
		if bytes, err := IoUtilReadFile(path); err != nil {
			return err
		} else {
			if err := json.Unmarshal(bytes, v); err != nil {
				return err
			} else {
				return nil
			}
		}
	} else {
		return fmt.Errorf("File not exists: %s\n", path)
	}
}

//通过IOUTIL读文件
func IoUtilReadFile(path string) ([]byte, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

//通过IOUtil写文件
func IoUtilWriteFile(path string, content []byte, perm os.FileMode) error {
	return ioutil.WriteFile(path, content, perm)
}

//通过OS写文件
func OsWriteFile(path string, content []byte, perm os.FileMode, append bool) error {
	var f *os.File
	var err error
	if append {
		f, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, perm)
	} else {
		f, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, perm)
	}
	defer f.Close()
	if err != nil {
		return err
	} else {
		_, err := f.Write(content)
		if err != nil {
			return err
		} else {
			f.Sync()
			f.Chmod(perm)
			return nil
		}
	}
}

//通过Bufio写文件
func BufWriteFile(path string, content []byte, perm os.FileMode) error {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	} else {
		w := bufio.NewWriter(f)
		w.Write(content)
		w.Flush()
		f.Chmod(perm)
		return nil
	}
}

//判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//获取文件大小
func FileSize(path string) int64 {
	if FileExists(path) {
		if info, err := os.Stat(path); err == nil {
			return info.Size()
		}
	}
	return int64(0)
}

//获取用户家目录
func GetUserHome() (string, error) {
	user, err := user.Current()
	if err != nil {
		switch runtime.GOOS {
		case "windows":
			return homeWindows()
		case "darwin", "freebsd", "netbsd", "openbsd", "solaris", "dragonfly":
			return homeUnix()
		default:
			return "", err
		}
	} else {
		return user.HomeDir, nil
	}
}

//获取Unix系统家目录
func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

//获取windows系统家目录
func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", fmt.Errorf("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

func WorkDir() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	p, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	return path.Dir(strings.Replace(p, "\\", "/", -1)), err
}
