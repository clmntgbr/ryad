package main

import (
	"fmt"
	"github.com/extrame/xls"
)

func main() {
	if xlFile, err := xls.Open("cables.xls", "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(3); sheet1 != nil {
			fmt.Print("Total Lines ", sheet1.MaxRow, sheet1.Name)
			col1 := sheet1.Row(1).Col()
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				col1 = row1.Col(1)
				fmt.Print("\n", col1)
			}
		}
	}
}