package csvutil

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/zhouweigithub/goutil/fileutil"
)

// 读取CSV文件到二维字符数组
//
//	path: 文件路径
//	lineDvidStr: 行分隔符，若为空则默认为\r\n
//	wordDvidStr: 列分隔符，若为空则默认为,
func ReadCsv(path, lineDvidStr, wordDvidStr string) [][]string {
	if lineDvidStr == "" {
		lineDvidStr = "\r\n"
	}
	if wordDvidStr == "" {
		wordDvidStr = ","
	}
	res, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("打开文件失败：", path, err.Error())
	}
	if len(res) == 0 {
		log.Println("该文件无内容：", path)
		return nil
	}
	var content = string(res)
	// 文件每行的内容
	var datas = strings.Split(content, lineDvidStr)
	// 首行列数
	var firstRowFieldCount = 0
	// 所有行数据矩阵
	var result [][]string
	for i := 0; i < len(datas); i++ {
		var rowData []string
		rowData = append(rowData, strings.Split(datas[i], wordDvidStr)...)
		if i == 0 {
			firstRowFieldCount = len(rowData)
		} else {
			// 如果某行的字段数少于首行字段数量，则丢弃此行
			if len(rowData) < firstRowFieldCount {
				continue
			}
		}
		result = append(result, rowData)
	}
	return result
}

// 将内容保存为CSV文件，（覆盖已有文件，文件不存在则创建，目录需要提前创建）
//
//	path: 文件路径
//	lineDvidStr: 行分隔符，若为空则默认为\r\n
//	wordDvidStr: 列分隔符，若为空则默认为,
//	datas: CSV文件的内容
func WriterCoverCSV(path, lineDvidStr, wordDvidStr string, datas [][]string) {
	if lineDvidStr == "" {
		lineDvidStr = "\r\n"
	}
	if wordDvidStr == "" {
		wordDvidStr = ","
	}
	var sb strings.Builder
	for i := 0; i < len(datas); i++ {
		for j := 0; j < len(datas[i]); j++ {
			sb.WriteString(datas[i][j])
			if j != len(datas[i])-1 {
				sb.WriteString(wordDvidStr)
			}
		}
		sb.WriteString(lineDvidStr)
	}
	fileutil.WriteTextFile(path, sb.String())
}
