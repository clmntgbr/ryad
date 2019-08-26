package main

import (
	"fmt"
	"regexp"

	"github.com/extrame/xls"
)

func main() {
	if xlsFile, err := xls.Open("cables.xls", "utf-8"); err == nil {
		for i := 0; i <= xlsFile.NumSheets(); i++ { //i is the sheet number
			if sheet := xlsFile.GetSheet(i); sheet != nil {
				for j := 0; j <= (int(sheet.MaxRow)); j++ { //j is the row number
					if row := sheet.Row(j); row != nil {
						for h := 0; h <= row.LastCol(); h++ { //h is the column number
							row := sheet.Row(j)
							column := row.Col(h)
							find(column)
						}
					}
				}
			}
		}
	}
}

func find(value string) {
	if matched, error := regexp.Match(`([A-Z]{3})_([0-9]{5})\w+`, []byte(value)); error == nil {
		if matched == true {
			fmt.Println(value)
			//os.Mkdir(value, 0777)
		}
	}
	if matched, error := regexp.Match(`(SITE)\w+`, []byte(value)); error == nil {
		if matched == true {
			fmt.Println(value)
			//os.Mkdir(value, 0777)
		}
	}
}
