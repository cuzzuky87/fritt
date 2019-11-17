package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/urfave/cli"
)

func main() {
	var tplType string
	var strDate string
	var filename string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "type,t",
			Usage:       "select tmplate type",
			Value:       "giji",
			Destination: &tplType,
		},
		cli.StringFlag{
			Name:        "date,d",
			Usage:       "Change dsignated date",
			Value:       time.Now().Format("2006-0-02"),
			Destination: &strDate,
		},
		cli.StringFlag{
			Name:        "name,n",
			Usage:       "file name",
			Value:       "input_name",
			Destination: &filename,
		},
	}

	app.Action = func(c *cli.Context) error {
		Render(filename, tplType, strDate)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Render(title string, tpltype string, date string) error {

	m := map[string]string{
		"title": title,
		"date":  date,
	}

	p, err := os.Executable()
	if err != nil {
		log.Fatal(err)
		return err
	}

	f, err := os.Create(filepath.Join(filepath.Dir(p), (title + ".md")))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer f.Close()

	switch tpltype {
	case "giji":
		tpl := template.Must(template.ParseFiles("templates//giji.tpl"))
		tpl.Execute(f, m)
	case "repo":
		tpl := template.Must(template.ParseFiles("templates//repo.tpl"))
		tpl.Execute(f, m)
	default:
		tpl := template.Must(template.ParseFiles("templates/empt.tpl"))
		tpl.Execute(f, m)
	}
	return nil
}
