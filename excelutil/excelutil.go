package excelutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gogf/gf/util/gconv"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
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

// 读取EXCEL内容
//
//	filename：指定文件路径名称，例如："D:/aa.xlsx"
//	sheetName: sheet名称，不传默认第一个
func ReadFromExcel(filename string, sheetName string) ([][]string, error) {
	f, err := excelize.OpenFile(filename) //
	if err != nil {
		return nil, err
	}
	defer f.Close()
	firstSheet := ""
	if len(sheetName) > 0 {
		firstSheet = sheetName
	} else {
		firstSheet = f.GetSheetName(0)
	}
	rows, err := f.GetRows(firstSheet)
	f.Close()
	return rows, err
}

// 往EXCEL写入内容
//
//	filename：指定文件路径名称，例如："D:/aa.xlsx"
//	value：内容
//	sheetName: sheet名称，不传默认第一个
func WriteToExcel(filename string, value [][]string, sheetName string) error {
	f := excelize.NewFile()
	defer f.Close()
	//默认保存开始名称
	firstSheet := "Sheet1"
	if len(sheetName) > 0 {
		firstSheet = sheetName
	}

	index := f.NewSheet(firstSheet)
	for i := 0; i < len(value); i++ {
		for k, v := range value[i] { //列
			path, err := excelize.ColumnNumberToName(k + 1)
			if err != nil {
				return err
			}
			err = f.SetCellValue(firstSheet, path+gconv.String(i+1), v)
			if err != nil {
				return err
			}
		}
	}

	f.SetActiveSheet(index - 1)
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
