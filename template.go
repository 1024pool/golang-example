package main
          
import (
    "text/template"
    "time"
    "os"
)
          
type User struct {
    Username, Password string
    RegTime time.Time
}
          
func ShowTime(t time.Time, format string) string {
    return t.Format(format)
}
          
func main() {
    u := User{"dotcoo", "dotcoopwd", time.Now()}
    t, err := template.New("text").Funcs(template.FuncMap{"showtime":ShowTime}).
        Parse(`<p>{{.Username}}|{{.Password}}|{{.RegTime.Format "2006-01-02 15:04:05"}}</p>
<p>{{.Username}}|{{.Password}}|{{showtime .RegTime "2006-01-02 15:04:05"}}</p>
`)
    if err != nil {
        panic(err)
    }
    t.Execute(os.Stdout, u)
}
