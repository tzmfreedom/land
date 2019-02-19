package server

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
	"github.com/tzmfreedom/land/compiler"
	"github.com/tzmfreedom/land/interpreter"
)

type Server struct {
	ClassTypes []*ast.ClassType
}

func (s *Server) Run() {
	classMap := builtin.NewClassMapWithPrimivie(s.ClassTypes)
	interpreter := interpreter.NewInterpreter(classMap)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reg := regexp.MustCompile("^/([^/]+)/([^/]+)")
		match := reg.FindStringSubmatch(r.RequestURI)
		klass := match[1]
		method := match[2]
		params := []interface{}{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		json.Unmarshal(buf.Bytes(), &params)

		parameters := make([]ast.Node, len(params))
		for i, param := range params {
			var p ast.Node
			switch n := param.(type) {
			case string:
				p = &ast.StringLiteral{Value: n}
			case float64:
				p = &ast.DoubleLiteral{Value: n}
			case int:
				p = &ast.IntegerLiteral{Value: n}
			case bool:
				p = &ast.BooleanLiteral{Value: n}
			}
			parameters[i] = p
		}
		name := []string{klass, method}
		invoke := &ast.MethodInvocation{
			NameOrExpression: &ast.Name{
				Value: name,
			},
			Parameters: parameters,
		}
		//resolver := &compiler.TypeResolver{
		//	Context: &compiler.Context{ClassTypes: classMap},
		//}
		//if _, err := resolver.ResolveMethod(name, parameters); err != nil {
		//	panic(err)
		//}
		res, err := invoke.Accept(interpreter)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, builtin.String(res.(*ast.Object)))
		fmt.Fprintln(w)
	})
	port := getServerPort()
	fmt.Println("listening to 0.0.0.0:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

type EvalServer struct{}

type EvalRequest struct {
	String    string
	Method    string
	WithClass bool
}

type EvalResult struct {
	String string
	Result bool
	Error  string
}

type FormatRequest struct {
	String string
}

type FormatResponse struct {
	String string
}

func (s *EvalServer) Run() {
	dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	http.HandleFunc("/code/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/code/")
		switch r.Method {
		case http.MethodPost:
			req := map[string]string{}
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			json.Unmarshal(buf.Bytes(), &req)
			_, err := db.Exec("UPDATE snippets SET code = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", req["code"], id)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			body, err := json.Marshal(map[string]string{
				"id": id,
			})
			fmt.Fprintf(w, string(body))
		case http.MethodGet:
			rows, err := db.Query("SELECT code FROM snippets WHERE id = $1", id)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			defer rows.Close()
			var code string
			rows.Next()
			rows.Scan(&code)
			body, err := json.Marshal(map[string]string{
				"code": code,
			})
			fmt.Fprintf(w, string(body))
		}
	})
	http.HandleFunc("/code", func(w http.ResponseWriter, r *http.Request) {
		req := map[string]string{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		json.Unmarshal(buf.Bytes(), &req)
		id := generateId()
		_, err := db.Exec("INSERT INTO snippets(id, code) VALUES ($1, $2)", id, req["code"])
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		body, err := json.Marshal(map[string]string{
			"id": id,
		})
		fmt.Fprintf(w, string(body))
	})
	http.HandleFunc("/eval", func(w http.ResponseWriter, r *http.Request) {
		eval(w, r)
	})
	http.HandleFunc("/format", func(w http.ResponseWriter, r *http.Request) {
		// parse request
		req := &FormatRequest{}
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		json.Unmarshal(buf.Bytes(), &req)
		b, err := base64.StdEncoding.DecodeString(req.String)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}

		// parse source
		root, err := ast.ParseString(string(b))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		// format
		visitor := &ast.TosVisitor{}
		str, err := root.Accept(visitor)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}

		// response
		b64body := base64.StdEncoding.EncodeToString([]byte(str.(string)))
		body, err := json.Marshal(&FormatResponse{
			String: b64body,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		fmt.Fprintf(w, string(body))
	})
	http.Handle("/", http.FileServer(http.Dir("eval-server")))

	port := getServerPort()
	fmt.Println("listening to 0.0.0.0:" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func eval(w http.ResponseWriter, r *http.Request) {
	classMap := builtin.NewClassMapWithPrimivie([]*ast.ClassType{})

	req := &EvalRequest{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	json.Unmarshal(buf.Bytes(), &req)
	b, err := base64.StdEncoding.DecodeString(req.String)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}

	classBody := string(b)
	if !req.WithClass {
		classBody = fmt.Sprintf(`public class Land { public static void action() { %s; } }`, classBody)
	}
	// compile
	root, err := ast.ParseString(classBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}
	register := &compiler.ClassRegisterVisitor{}
	t, err := root.Accept(register)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}
	classType := t.(*ast.ClassType)
	classMap.Set(classType.Name, classType)
	typeChecker := compiler.NewTypeChecker()
	typeChecker.Context.ClassTypes = classMap
	_, err = typeChecker.VisitClassType(classType)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		body, err := json.Marshal(&EvalResult{
			Error:  err.Error(),
			Result: false,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error")
		}
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(body))
		return
	}
	if len(typeChecker.Errors) != 0 {
		for _, e := range typeChecker.Errors {
			fmt.Fprintf(os.Stderr, "%s\n", e.Message)
		}
	}
	// interpreter
	invoke := &ast.MethodInvocation{
		NameOrExpression: &ast.Name{
			Value: []string{classType.Name, req.Method},
		},
	}

	interpreter := interpreter.NewInterpreter(classMap)
	interpreter.LoadStaticField()
	stdout := new(bytes.Buffer)
	interpreter.Extra["stdout"] = stdout
	interpreter.Extra["stderr"] = new(bytes.Buffer)
	_, err = invoke.Accept(interpreter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}

	b64body := base64.StdEncoding.EncodeToString(stdout.Bytes())
	body, err := json.Marshal(&EvalResult{
		String: b64body,
		Result: true,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error")
	}
	fmt.Fprint(w, string(body))
}

func getServerPort() string {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return port
}

func generateId() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func Run(classTypes []*ast.ClassType) {
	server := &Server{ClassTypes: classTypes}
	server.Run()
}
