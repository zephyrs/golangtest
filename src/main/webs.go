package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	webBase = iota
	webStruct
	webHead
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

var (
	strSay   = "/fun1/"
	strRoar  = "/fun2/"
	strForm  = "/form/"
	strRec   = "/rec/"
	localURL = "localhost:9999"
)

type localHandler struct {
}

func (h *localHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ss := req.URL.Path
	switch {
	case strings.HasPrefix(ss, strSay):
		say(w, req)
	case strings.HasPrefix(ss, strRoar):
		roar(w, req)
	case strings.HasPrefix(ss, strForm):
		w.Header().Set("Content-Type", "text/html")
		switch req.Method {
		case "GET":
			io.WriteString(w, form)
		case "POST":
			io.WriteString(w, req.FormValue("in"))
		}
	case strings.HasPrefix(ss, strRec):
		rec(w, req)
	}
}

func runTestWebs() {
	defer un(trace("runTestWebs"))

	t := webStruct
	switch t {
	case webBase:
		http.HandleFunc(strSay, say)
		http.HandleFunc(strRoar, roar)
		http.ListenAndServe(localURL, nil)
	case webStruct:
		var h localHandler
		http.ListenAndServe(localURL, &h)
	case webHead:
		url := "http://www.baidu.com/"
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}

}

func say(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[len(strSay):]
	fmt.Fprintf(w, "bonjour %s", s)
}

func roar(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[len(strRoar):]
	fmt.Fprintf(w, "BONJOUR %s!", strings.ToUpper(s))
}

func rec(w http.ResponseWriter, r *http.Request) {
	fmt.Println("**********************************************")
	switch r.Method {
	case "GET":
		fmt.Println("Method: GET")
	case "POST":
		fmt.Println("Method: POST")
	}

	fmt.Println("-------")

	fmt.Println("HEADER:")
	for k, v := range r.Header {
		fmt.Printf("%s : %s\n", k, v)
	}

	fmt.Println("-------")

	fmt.Println("BODY:")
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	fmt.Println(buf.String())
	fmt.Println("**********************************************")

	fmt.Fprintf(w, `
	{"msg": "success", 
	"code": 0}`)
}
