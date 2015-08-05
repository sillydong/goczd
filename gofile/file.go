package gofile
import (
	"os/user"
	"runtime"
	"errors"
	"os"
	"fmt"
	"bytes"
	"os/exec"
	"strings"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/bufio.v1"
)

//将Json内容写入文件
func Json2File(content, path string,perm os.FileMode) error {
	if FileExists(path){
		if err:=os.Remove(path);err!=nil{
			return err
		}
	}
	bytes := []byte(content)
	return IoUtilWriteFile(path,bytes,perm)
}

//将Json内容从文件读取并解析
func File2Json(path string, v interface{}) error{
	if FileExists(path) {
		if bytes, err := IoUtilReadFile(path); err!=nil {
			return err
		}else {
			if err := json.Unmarshal(bytes, v); err!=nil {
				return err
			}else {
				return nil
			}
		}
	}else {
		return fmt.Errorf("File not exists: %s\n", path)
	}
}

//通过IOUTIL读文件
func IoUtilReadFile(path string)([]byte,error){
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

//通过IOUtil写文件
func IoUtilWriteFile(path string,content []byte, perm os.FileMode)error{
	return ioutil.WriteFile(path,content,perm)
}

//通过OS写文件
func OsWriteFile(path string, content []byte, perm os.FileMode) error{
	f, err := os.Create(path)
	if err != nil {
		return err
	} else {
		defer f.Close()
		_,err:=f.Write(content)
		if err != nil {
			return err
		}else {
			f.Sync()
			f.Chmod(perm)
			return nil
		}
	}
}

//通过Bufio写文件
func BufWriteFile(path string, content []byte, perm os.FileMode) error{
	f,err:=os.Create(path)
	if err != nil {
		return err
	} else {
		defer f.Close()
		w:=bufio.NewWriter(f)
		w.Write(content)
		w.Flush()
		f.Chmod(perm)
		return nil
	}
}

//判断文件是否存在
func FileExists(path string)bool{
	_,err := os.Stat(path)
	return err==nil || os.IsExist(err)
}

//获取用户家目录
func GetUserHome() (string, error) {
	user, err := user.Current()
	if err != nil {
		switch runtime.GOOS {
		case "windows":
			return homeWindows()
		case "darwin","freebsd","netbsd","openbsd","solaris","dragonfly":
			return homeUnix()
		default:
			return "", err
		}
	} else {
		return user.HomeDir,nil
	}
}

//获取Unix系统家目录
func homeUnix() (string,error) {
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
func homeWindows() (string,error) {
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
