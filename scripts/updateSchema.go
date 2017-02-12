package main

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
	"qiushibaike.com/test-vue/data"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	result := graphql.Do(graphql.Params{
		Schema: data.Schema,
		RequestString: testutil.IntrospectionQuery,
	})
	if result.HasErrors() {
		log.Fatalf("ERROR introspecting schema: %v", result.Errors)
		return
	} else {
		b, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}

		err = ioutil.WriteFile("../data/schema.json", b, os.ModePerm)
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		}
	}
}
