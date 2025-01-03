package util

import "fmt"

type TableColumn struct {
	Key   string
	Title string
}

func PrintTable(columns []TableColumn, rows []map[string]string, boldHeaders bool) {
	maxWidths := make(map[string]int)
	for _, col := range columns {
		maxWidths[col.Key] = len(col.Title)
		for _, row := range rows {
			maxWidths[col.Key] = max(maxWidths[col.Key], len(row[col.Key]))
		}
	}

	for _, col := range columns {
		field := formatField(col.Title, maxWidths[col.Key])
		if boldHeaders {
			field = Bold.Render(field)
		}
		fmt.Print(field + " ")
	}
	fmt.Println()

	for _, row := range rows {
		for _, col := range columns {
			fmt.Print(formatField(row[col.Key], maxWidths[col.Key]) + " ")
		}
		fmt.Println()
	}
}

func formatField(value string, width int) string {
	if value == "" {
		value = "-"
	}

	return PadLeft(value, width, " ")
}
