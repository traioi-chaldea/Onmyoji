package tpl

import (
	"fmt"
	"text/template"
	"os"
)

type DAILY_ENTRY struct {
	Greeting string
	Time string
	Content interface{}
}

func TPLDaily(cTime string, content interface{}) {
	var td DAILY_ENTRY
	tplName := fmt.Sprintf("daily%sh.tpl", cTime)

	td.Greeting = getGreeting()
	td.Time = cTime
	td.Content = content

	t := template.Must(template.New(tplName).ParseFiles("template/" + tplName))
	err := t.Execute(os.Stdout, td)
	if err != nil {
		panic(err)
	}
}


