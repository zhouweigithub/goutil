package excelutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tealeg/xlsx"
)

func ReadExcel(filePath string, isFirstRowTitle bool) []map[string]string {
	var result = []map[string]string{}
	//获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return result
	}
	xlsxPath := dir + "\\" + filePath
	//打开文件路径
	xlsxFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		fmt.Println(err)
		return result
	}
	//读取每一个sheet
	//for x := range xlsxFile.Sheets {
	var sheet = xlsxFile.Sheets[0]
	if len(sheet.Rows) == 0 {
		return result
	}

	var titles = getExcelTitle(sheet.Rows, isFirstRowTitle)

	//读取每个sheet下面的行数据
	for y := range sheet.Rows {
		if isFirstRowTitle && y == 0 {
			continue
		}
		var row = sheet.Rows[y]
		//读取每个cell的内容
		var rowMap = make(map[string]string)
		for i := range row.Cells {
			rowMap[titles[i]] = row.Cells[i].Value
		}
		result = append(result, rowMap)
	}
	//}
	return result
}

func getExcelTitle(rows []*xlsx.Row, isFirstRowTitle bool) map[int]string {
	var rowMap = make(map[int]string)
	if len(rows) == 0 {
		return rowMap
	} else {
		for i := range rows[0].Cells {
			if isFirstRowTitle {
				rowMap[i] = rows[0].Cells[i].Value
			} else {
				rowMap[i] = strconv.Itoa(i)
			}
		}
		return rowMap
	}
}
