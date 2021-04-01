package chapter4

import "time"

//模板字符串
//声明玄幻action
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func Main46() {

}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
