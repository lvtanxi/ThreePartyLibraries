/**  
* Date: 2017-10-09 
* Time: 16:10 
* Description:
*/
package main

import (
	"text/template"
	"os"
)

type Per struct {
	UserName string
	Emails []string
	Friends []Friend
}


type Friend struct {
	Fname string
}

func main() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}")
	p := Per{UserName:"测试"}
	t.Execute(os.Stdout, p)


	p =Per{UserName:"测试2",Emails:[]string{"163","outlook"},Friends:[]Friend{{"minux.ma"},{"xushiwei"}}}

	t = template.New("fieldname example")
	t,_= t.Parse(`hello {{.UserName}}
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}
	`)
	t.Execute(os.Stdout, p)


	t =template.New("Empty  test")
	t = template.Must(t.Parse("空 pipeline if demo :{{if ``}} 不会出书 {{end}}"))
	t.Execute(os.Stdout, nil)

}
