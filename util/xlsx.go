package util

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func ReadExcelData(file string, sheetIndex int) (rs [][]string, err error) {
	xlFile, err := xlsx.OpenFile(file)
	if err != nil {
		return
	}

	if sheetIndex >= len(xlFile.Sheets) {
		err = fmt.Errorf("sheet index out of range")
		return
	}

	sheet := xlFile.Sheets[sheetIndex]
	for _, row := range sheet.Rows {
		var rowStr []string
		for _, cell := range row.Cells {
			rowStr = append(rowStr, cell.String())
		}
		rs = append(rs, rowStr)
	}
	return
}
