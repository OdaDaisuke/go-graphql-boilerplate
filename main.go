package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/OdaDaisuke/go-graphql-sample/graphql_util"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		result := executeQuery(fmt.Sprintf("%s", body), graphql_util.Schema)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}