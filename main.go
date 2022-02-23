package chartmodify

import (
	"encoding/json"

	"github.com/xuri/excelize/v2"
)

func main() {
	name := "load.xlsx"
	building := &Builing{ID: "", Name: "建筑物逐时负荷", Systems: []Systems{
		{ID: "1", Name: "默认系统"},
		{ID: "2", Name: "默认系统"},
	},
		Regions: 7}
	f, err := excelize.OpenFile(name, excelize.Options{Password: "password"})
	if err != nil {
		c := 1
		cloumn, _ := excelize.ColumnNumberToName(building.Regions + 2) //从第3列开始
		for _, s := range building.Systems {
			chartSetting := &ChartSetting{
				Type:   "col",
				Series: []Series{{Name: "冷负荷", Categories: "$C$9229:$" + cloumn + "$9229", Values: "$C$9230:$" + cloumn + "$9230"}},
				Format: Format{},
				Legend: Legend{},
				Title:  Title{Name: "冷负荷"},
			}
			cStr, err := json.Marshal(chartSetting)
			if err == nil {

				sheetName := "系统报表-" + s.Name
				index := f.GetSheetIndex(sheetName)
				if index > -1 {
					if err := f.AddChart(sheetName, "B9226", string(cStr)); err != nil {
						println(err.Error())
					} else {
						c += 1
					}
				}

			}

		}

		// 根据指定路径保存文件
		if err := f.SaveAs("Book1.xlsx"); err != nil {
			println(err.Error())
		}
	}

}
