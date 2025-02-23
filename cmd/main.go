package main 

import (
	"io"
	"html/template"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"

)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("../views/*.html")),
	}
}

type Count struct {
 Count int
}
type Contact struct {
    Name  string
    Email string
}

func newContact(name, email string) Contact {  
    return Contact{
        Name:  name,
        Email: email,
    }
}

type Contacts = []Contact

type Data struct {
    Contacts Contacts
}

func newData() Data {
    return Data{
        Contacts: []Contact{
            newContact("John", "john@gamil.com"),
            newContact("Jane", "jane@example.com"),
            newContact("Bob", "bob@example.com"),
        },
    }
}

func main(){

	e := echo.New()
	e.Use(middleware.Logger())


	data := newData ()  
	e.Renderer = newTemplate()


	e.GET("/", func(c echo.Context) error {
        return c.Render(200, "index", data)
    })

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")	
		data.Contacts = append(data.Contacts, newContact(name, email))
        return c.Render(200, "display", data )
    })

	e.Logger.Fatal(e.Start(":26087"))


} 