package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strings"
	"time"
)

func CreateExcelEventList(fileName string, sheetName string, eventList []model.EventDetail) (filePath string, err error) {
	ind := strings.LastIndex(fileName, ".xlsx")
	if ind == 0 {
		return "", errors.New("provide valid file name. EX. abc.xlsx")
	} else if ind < 0 {
		fileName = fileName + ".xlsx"
	}

	f := excelize.NewFile()
	//index := f.NewSheet(sheetName)
	//f.SetActiveSheet(index)

	existingSheetName := f.GetSheetName(0)
	f.SetSheetName(existingSheetName, sheetName)
	index := f.GetSheetIndex(sheetName)
	f.SetActiveSheet(index)

	f.SetCellValue(sheetName, "A1", "Event Date")
	f.SetCellValue(sheetName, "B1", "Title")
	f.SetCellValue(sheetName, "C1", "Actual Event Date")
	f.SetCellValue(sheetName, "D1", "Description")

	f.SetColWidth(sheetName, "A", "A", 15)
	f.SetColWidth(sheetName, "B", "B", 40)
	f.SetColWidth(sheetName, "C", "C", 20)
	f.SetColWidth(sheetName, "D", "D", 80)

	// wrap text in cell
	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})

	for ind, event := range eventList {
		dateTime := time.Date(time.Now().Year(), time.Month(event.Month), event.Day, 0, 0, 0, 0, time.Local)

		cellAddressEventDate := fmt.Sprintf("A%d", (ind + 2))
		cellAddressTitle := fmt.Sprintf("B%d", (ind + 2))
		cellAddressActualEventDate := fmt.Sprintf("C%d", (ind + 2))
		cellAddressDescription := fmt.Sprintf("D%d", +(ind + 2))

		f.SetCellValue(sheetName, cellAddressEventDate, dateTime.Format("Mon Jan 2 2006"))

		f.SetCellValue(sheetName, cellAddressTitle, event.Title)
		f.SetCellValue(sheetName, cellAddressActualEventDate, event.EventDate.Format(constants.DATE_FORMAT_EVENT_DATE_DISPLAY))
		f.SetCellValue(sheetName, cellAddressDescription, event.Info)

		if err := f.SetCellStyle(sheetName, cellAddressDescription, cellAddressDescription, style); err != nil {
			logrus.Errorf("error setting style for text wrapping")
		}

		// Highlight today's events
		//if time.Now().Day() == event.Day {
		if ind == 0 {
			color := "2354e8" // blue
			highlightCell(f, sheetName, cellAddressEventDate, dateTime.Format("Mon Jan 2 2006"), color)
			highlightCell(f, sheetName, cellAddressTitle, event.Title, color)
			highlightCell(f, sheetName, cellAddressActualEventDate, event.EventDate.Format(constants.DATE_FORMAT_EVENT_DATE_DISPLAY), color)
			highlightCell(f, sheetName, cellAddressDescription, event.Info, color)
		}
	}

	if err := f.SaveAs(fileName); err != nil {
		return "", err
	}

	path, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filePath = path + "/" + fileName

	return filePath, nil
}

func highlightCell(f *excelize.File, sheetName, cellAddress, cellValue, color string) {
	err := f.SetCellRichText(sheetName, cellAddress, []excelize.RichTextRun{
		{
			Text: cellValue,
			Font: &excelize.Font{
				Bold:   true,
				Color:  color,
				Family: "Times New Roman",
			},
		},
	})
	if err != nil {
		logrus.Errorf("error in highlighting cellAddress=%v, sheetName=%v, cellValue=%v. error=%v", cellAddress, sheetName, cellValue, err)
	}
}
