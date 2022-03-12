package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	renderTable()
}

func renderTable() {
	headers := []string{"NAME", "DESCRIPTION", "SCOPE", "DISCOVERY", "VERSION", "STATUS"}
	data := [][]string{
		{"cluster", "Kubernetes cluster operations", "Context", "default-codegold-mar09", "v0.19.0-dev", "installed"},
		{"feature", "Operate on features and featuregates", "Context", "default-codegold-mar09", "v0.19.0-dev", "installed"},
		{"kubernetes-release", "Kubernetes release operations", "Context", "default-codegold-mar09", "v0.19.0-dev", "installed"},
		{"kubernetes-release", "Kubernetes release operations space Kubernetes release space operations Kubernetes release space operations Kubernetes release operations Kubernetes release operations Kubernetes release operations Kubernetes release operations", "Context", "default-codegold-mar09", "v0.19.0-dev", "installed"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetColWidth(80)
	table.SetTablePadding("\t\t")
	table.SetHeader(headers)
	table.AppendBulk(data)
	table.Render()
}
