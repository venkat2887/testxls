package main

import (
    "fmt"
    "github.com/tealeg/xlsx"
)

func main() {
    excelFileName := "/Users/shanks/Documents/foo.xlsx"
    xlFile, err := xlsx.OpenFile(excelFileName)
    if err != nil {
        ...
    }
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            for _, cell := range row.Cells {
                fmt.Printf("%s\n", cell.String())
            }
        }
    }
}