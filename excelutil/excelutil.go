package excelutil

import (
	"fmt"
	"math"
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

// filename：指定文件路径名称，例如："D:/aa.xlsx"
// tablename: （可变参数）指定Excel中的某个表格名，不传默认第一个表格
func ReadFromExcel(filename string, tablename string) ([][]string, error) {
	f, err := excelize.OpenFile(filename) //
	if err != nil {
		return nil, err
	}
	firstSheet := ""
	if len(tablename) > 0 {
		firstSheet = tablename
	} else {
		firstSheet = f.GetSheetName(0)
	}
	rows, err := f.GetRows(firstSheet)
	return rows, err
}

// filename：指定文件路径名称，例如："D:/aa.xlsx"
// value：内容
// tableName: （可变参数）指定Excel中的某个表格名，不传默认第一个表格
func WriteToExcel(filename string, value [][]string, tableName string) error {
	f := excelize.NewFile()
	//默认Excel每个页面保存500条数据，超出500条就新建一个页面保存，page为每页最多保存条数
	page := 500
	//默认保存开始名称
	firstSheet := "Sheet"
	if len(tableName) > 0 {
		firstSheet = tableName
	}
	// Create a new sheet.
	sheetRow := int(math.Floor(float64(len(value)/page)) + 1)
	for j := 0; j < sheetRow; j++ {
		index := f.NewSheet(firstSheet + gconv.String(j+1))
		if j > 0 {
			for k, v := range value[0] { //列
				path, err := excelize.ColumnNumberToName(k + 1)
				if err != nil {
					return err
				}
				err = f.SetCellValue(firstSheet+gconv.String(j+1), path+gconv.String(1), v)
				if err != nil {
					return err
				}
			}
			for i := 0 + j*page; i < (j+1)*page; i++ { //行
				if len(value) < i+1 {
					break
				}
				for k, v := range value[i] { //列
					path, err := excelize.ColumnNumberToName(k + 1)
					if err != nil {
						return err
					}
					err = f.SetCellValue(firstSheet+gconv.String(j+1), path+gconv.String(i+2-j*page), v)
					if err != nil {
						return err
					}
				}
			}
		} else {
			for i := 0 + j*page; i < (j+1)*page; i++ { //行
				if len(value) < i+1 {
					break
				}
				for k, v := range value[i] { //列
					path, err := excelize.ColumnNumberToName(k + 1)
					if err != nil {
						return err
					}
					err = f.SetCellValue(firstSheet+gconv.String(j+1), path+gconv.String(i+1-j*page), v)
					if err != nil {
						return err
					}
				}
			}
		}
		f.SetActiveSheet(index - 1)
	}
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
