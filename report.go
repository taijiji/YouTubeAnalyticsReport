package main

import (
	"fmt"
	"os"
	"text/template"
)

func render_report(video_list []Video, date string) {

	// custom logic to culuculate remainder
	//funcMap := template.FuncMap{
	//	"mod": func(a, b int) int {
	//		return a % b
	//	},
	//}
	tmpl, err := template.ParseFiles("templates/tmpl_report.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := tmpl.Execute(os.Stdout, video_list); err != nil {
		panic(err)
	}

	f, err := os.Create("reports/showint_report_" + date + ".md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, video_list)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
