package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	advent18 "github.com/rcliao/advent-2018"
	"github.com/rcliao/advent-2018/day1"
)

var problems = make(map[string]advent18.SolutionFn)
var problemNumber = flag.String("id", "", "problem id (1,2,3) ...etc.")
var cookie = flag.String("cookie", "", "cookie session value for grabbing input from adventofcode")

func init() {
	problems[getInputURL("1")] = day1.Solution{}
}

func getInputURL(id string) string {
	return fmt.Sprintf("https://adventofcode.com/2018/day/%s/input", id)
}

func main() {
	flag.Parse()
	url := getInputURL(*problemNumber)
	cookie := http.Cookie{Name: "session", Value: *cookie}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.AddCookie(&cookie)
	if err != nil {
		fmt.Println("Failed to retrieve input from adventofcode.com")
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to retrieve input from adventofcode.com")
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to parse body content")
		panic(err)
	}
	fmt.Println("Hello advent 2018")
	fmt.Println(problems[url].Solve(string(body)))
}
