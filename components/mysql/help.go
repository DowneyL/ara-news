package mysql

import "ara-news/helper"

func GetCols(model interface{}, exceptedCols ...string) []string {
	var cols []string
	fields := helper.GetFieldName(model)
	if len(exceptedCols) > 0 {
		for _, filed := range fields {
			for _, col := range exceptedCols {
				if col == filed {
					continue
				}
				cols = append(cols, filed)
			}
		}
	}

	return cols
}
