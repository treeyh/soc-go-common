package excel

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tealeg/xlsx/v2"
	"github.com/treeyh/soc-go-common/core/consts"
	"github.com/treeyh/soc-go-common/core/logger"
	"net/http"
)

/**
* excelFileName 本地excel路径名
* orginFileUrl  远程excel url
* sheetIndex    第几份excel
* ignoreRows    忽略行数
* timeFormatIndex 需要转换时间格式的列
 */
func GenerateCSVFromXLSXFile(excelFileName string, orginFileUrl string, sheetIndex int, ignoreRows int, timeFormatIndex []int64) ([][]string, error) {
	var error error
	var xlFile *xlsx.File
	if orginFileUrl != "" {
		resp, err := http.Get(orginFileUrl)
		if err != nil {
			logger.Logger().Error(err)
			return nil, err
		}
		resBody := resp.Body
		buf := new(bytes.Buffer)
		buf.ReadFrom(resBody)

		xlFile, error = xlsx.OpenBinary(buf.Bytes())
	} else {
		xlFile, error = xlsx.OpenFile(excelFileName)

	}
	if error != nil {
		logger.Logger().Error(error)
		return nil, error
	}
	sheetLen := len(xlFile.Sheets)
	switch {
	case sheetLen == 0:
		return nil, errors.New("This XLSX file contains no sheets.")
	case sheetIndex >= sheetLen:
		return nil, fmt.Errorf("No sheet %d available, please select a sheet between 0 and %d\n", sheetIndex, sheetLen-1)
	}
	sheet := xlFile.Sheets[sheetIndex]

	result := assemblyResult(sheet, ignoreRows, timeFormatIndex)

	return result, nil
}

func assemblyResult(sheet *xlsx.Sheet, ignoreRows int, timeFormatIndex []int64) [][]string {

	result := [][]string{}
	for k, row := range sheet.Rows {
		if k <= ignoreRows-1 {
			continue
		}
		var vals []string
		if row != nil {
			for k, cell := range row.Cells {
				str, err := cell.FormattedValue()
				if err != nil {
					vals = append(vals, err.Error())
					continue
				}
				if isInIndex(timeFormatIndex, int64(k)) {
					time, err := cell.GetTime(false)
					if err != nil {
						logger.Logger().Error(err)
					} else {
						str = time.Format(consts.AppTimeFormat)
					}
				}
				//vals = append(vals, fmt.Sprintf("%q", str))
				vals = append(vals, str)
			}
		}
		result = append(result, vals)
	}
	return result
}

func isInIndex(param []int64, index int64) bool {
	for _, v := range param {
		if v == index {
			return true
		}
	}
	return false
}
