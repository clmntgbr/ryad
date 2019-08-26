package main

import (
	"os"

	"github.com/extrame/xls"
)

func main() {
	if xlsFile, err := xls.Open("cables.xls", "utf-8"); err == nil {
		for i := 0; i <= xlsFile.NumSheets(); i++ { // i is the sheet number
			if sheet := xlsFile.GetSheet(i); sheet != nil {
				for j := 0; j <= (int(sheet.MaxRow)); j++ { // j is the row number
					if row := sheet.Row(j); row != nil {
						for h := 0; h <= row.LastCol(); h++ { // h is the column number
							row := sheet.Row(j)
							column := row.Col(h)
							if column == "CHB_02337_AS_X31" {
								os.Mkdir(column, 0777)
							}
						}
					}
				}
			}
		}
	}
}
