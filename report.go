package main

import (
	"fmt"
	"os"
	"text/template"
)

type Tmpl_param struct {
	Date          string
	Channel_stats ChannelStats
	Video_list    []Video
}

func render_report(date string, channel_stats ChannelStats, video_list []Video) {

	param := Tmpl_param{
		Date:          date,
		Channel_stats: channel_stats,
		Video_list:    video_list,
	}

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

	err = tmpl.Execute(f, param)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
