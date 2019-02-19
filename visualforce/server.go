package visualforce

import (
	"fmt"

	"net/http"

	"encoding/base64"
	"encoding/json"

	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
	"github.com/tzmfreedom/land/interpreter"
)

func handleRequest(i *interpreter.Interpreter, r *http.Request, w http.ResponseWriter) {
	// initialize page to specify controller
	pagePath := r.URL.Path[1:]
	n, err := createNode(pagePath)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	attrs := n.attributeValues()
	r.ParseForm()
	method := r.Form["__action"][0]
	if method == "" {
		panic("method not blank")
	}
	b64viewstate := r.Form["__viewstate"][0]
	state := map[string]interface{}{}
	viewstate, err := base64.StdEncoding.DecodeString(b64viewstate)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(viewstate, &state)
	if err != nil {
		panic(err)
	}
	// run action method
	controller := attrs.Get("controller")
	c, retValue, err := i.BindAndRun(controller, method, r.Form, state)
	if err != nil {
		panic(err)
	}
	// evaluate return value
	if retValue == builtin.Null {
		pagePath = r.URL.Path[1:]
	} else {
		pagePath = retValue.Extra["url"].(*ast.Object).StringValue()
		if pagePath != r.Referer() {
			http.Redirect(w, r, pagePath, http.StatusFound)
			return
		}
	}
	// render page if return value is same page
	for k, v := range c.InstanceFields.All() {
		state[k] = v
	}
	n, err = createNode(pagePath)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	attrs = n.attributeValues()
	body := renderNodes(n.Nodes, c)
	templateStore["page"].Execute(w, PageParameter{
		Body:       body,
		ShowHeader: attrs.Get("showHeader") != "false",
	})
}

func Server(i *interpreter.Interpreter) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			body, err := render(r.URL.Path[1:], i)
			if err != nil {
				w.WriteHeader(404)
				return
			}
			fmt.Fprint(w, body)
		case http.MethodPost:
			handleRequest(i, r, w)
		}
	})
	fmt.Println("listening to 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}
