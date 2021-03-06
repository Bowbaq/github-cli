package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	versionNumber = 1
)

type application struct {
	cli *cli.App
	gh  *github.Client
}

type tokenSource struct {
	token *oauth2.Token
}

func (t tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

func newApp() *application {
	app := application{
		cli: cli.NewApp(),
	}

	app.cli.Name = "github"
	app.cli.Usage = "github command line interface"
	app.cli.Version = "0.0." + strconv.Itoa(versionNumber)
	app.cli.HideVersion = true
	app.cli.HideHelp = true
	app.cli.Author = "Maxime Bury <maxime.bury@gmail.com>"

	var tc *http.Client
	if token := os.Getenv("GITHUB_API_TOKEN"); token != "" {
		tc = oauth2.NewClient(oauth2.NoContext, tokenSource{
			&oauth2.Token{AccessToken: token},
		})
	}

	app.gh = github.NewClient(tc)

	return &app
}

var app = newApp()

func main() {
	app.cli.Run(os.Args)
}

func fixHelp(c *cli.Context) {
	c.App.Author = app.cli.Author
	c.App.Email = app.cli.Email
	c.App.Version = app.cli.Version
	cli.ShowAppHelp(c)
}

func showHelp(c *cli.Context, methodName, usage string) {
	var out bytes.Buffer
	c.App.Writer = &out
	cli.ShowSubcommandHelp(c)

	re := regexp.MustCompile("command " + methodName + " [^\n]+")
	fmt.Print(re.ReplaceAllString(out.String(), fmt.Sprintf("%s %s [command options]", c.App.Name, usage)))
	os.Exit(1)
}

func fatalln(v ...interface{}) {
	fmt.Println(v...)
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkResponse(res *http.Response, err error) {
	check(err)
	check(github.CheckResponse(res))
}

func timePointer(t time.Time) *time.Time {
	p := new(time.Time)
	*p = t
	return p
}

func stringSlicePointer(s []string) *[]string {
	p := new([]string)
	*p = s
	return p
}
