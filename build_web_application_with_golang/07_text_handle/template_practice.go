package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person1 struct {
	UserName string
	email    string
}

func SimpleTemplate1() {
	t := template.New("fieldname example")
	t, _ = t.Parse("hello {{.UserName}}! {{.email}}")
	p := Person1{UserName: "Ass"}
	t.Execute(os.Stdout, p)
}

type Person2 struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}
type Friend struct {
	Fname string
}

func SimpleTemplate2() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`
        hello {{.UserName}}!
        {{range .Emails}}
            an email {{.}}
        {{end}}

        {{with .Friends}}
        {{range .}}
            my friend name is {{.Fname}}
        {{end}}
        {{end}}
    `)

	p := Person2{
		UserName: "asd",
		Emails:   []string{"sdf@sf.com", "sdkfsdf@s.com"},
		Friends:  []*Friend{&f1, &f2},
	}
	t.Execute(os.Stdout, p)
}

func ConditionInTemplate1() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出. {{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分 {{else}} else部分.{{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	// find the @ symbol
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}
	// replace the @ by " at "
	return (substrs[0] + " at " + substrs[1])
}

func CustomPipe() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.UserName}}!
                 {{range .Emails}}
                     an emails {{.|emailDeal}}
                 {{end}}
                 {{with .Friends}}
                 {{range .}}
                     my friend name is {{.Fname}}
                 {{end}}
                 {{end}}
                 `)
	p := Person2{UserName: "Astaxie",
		Emails:  []string{"astaxie@beego.me", "astaxie@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}

func MustOperation() {
	tOk := template.New("first")
	template.Must(tOk.Parse(" some static text /* and a comment */"))
	fmt.Println("The first one parsed OK.")

	template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	fmt.Println("The second one parsed OK.")

	fmt.Println("The next one ought to fail.")
	tErr := template.New("check parse error with Must")
	template.Must(tErr.Parse(" some static text {{ .Name }}"))
}

func EmbeddedTemplate() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()
	s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()
	s1.Execute(os.Stdout, nil)

}

func main() {
	SimpleTemplate1()
	SimpleTemplate2()
	ConditionInTemplate1()
	CustomPipe()
	MustOperation()
	EmbeddedTemplate()
}
