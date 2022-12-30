package cmdutil

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"

	"github.com/zhouweigithub/goutil/fileutil"
	"golang.org/x/text/encoding/simplifiedchinese"
)

// 执行cmd命令
func ExecuteCmd(command string, params ...string) (string, error) {
	cmd := exec.Command(command, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	cmd.Wait()
	bytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(out.Bytes())
	return string(bytes), err
}

// 在特定目录下执行cmd命令
//
// 如果 folder 为空，则在当前进程目录执行命令
func ExecuteCmdInFolder(folder string, command string, params ...string) (string, error) {
	cmd := exec.Command(command, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if folder != "" {
		cmd.Dir = folder
	}
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	cmd.Wait()
	bytes, err := simplifiedchinese.GBK.NewDecoder().Bytes(out.Bytes())
	return string(bytes), err
}

// 在特定目录下执行cmd命令，并使用 fmt.Println 实时输出结果
//
// 如果 folder 为空，则在当前进程目录执行命令
func ExecuteCmdInFolderToShow(folder string, command string, params ...string) error {
	cmd := exec.Command(command, params...)
	fmt.Println("ExecuteCmdInFolderToFile", folder, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		//fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	if folder != "" {
		cmd.Dir = folder
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		//line, err2 := reader.ReadString('\n')
		line, _, err2 := reader.ReadLine()
		if err2 != nil || io.EOF == err2 {
			break
		}

		var str, err = simplifiedchinese.GBK.NewDecoder().Bytes(line)
		fmt.Println(string(str), err)
	}
	err = cmd.Wait()
	return err
}

// 在特定目录下执行cmd命令，并将结果写入文件中
//
// 如果 folder 为空，则在当前进程目录执行命令
//
// fileName: 保存内容的文件路径
func ExecuteCmdInFolderToFile(folder, fileName, command string, params ...string) error {
	var fileFolder = path.Dir(fileName)
	fileutil.CreateFolderIfNotExists(fileFolder)
	var f *os.File
	var err error
	if !fileutil.IsExists(fileName) {
		f, _ = os.Create(fileName)
	} else {
		f, _ = os.OpenFile(fileName, os.O_APPEND, os.ModeAppend)
	}
	if err != nil {
		return err
	}
	defer f.Close()
	cmd := exec.Command(command, params...)
	fmt.Println("ExecuteCmdInFolderToFile", folder, cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe: ", err)
		return err
	}
	cmd.Stderr = os.Stderr
	if folder != "" {
		cmd.Dir = folder
	}
	err = cmd.Start()
	if err != nil {
		return err
	}
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)
	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		f.WriteString(line) //写入文件(字节数组)
		f.Sync()
	}
	f.WriteString("================= over =================") //写入文件(字节数组)
	f.Sync()
	err = cmd.Wait()
	return err
}
