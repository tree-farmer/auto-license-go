package main

import "context"
import "fmt"
import "github.com/machinebox/graphql"

func main() {
	// create a client (safe to share across requests)
	client := graphql.NewClient("https://api.github.com/graphql")

	// make a request
	req := graphql.NewRequest(`
		query organization(login: "tree-farmer") {
		repositories(first: 10) {
		  totalCount
		  edges {
			node {
			  name
			  licenseInfo {
				name
			  }
			}
		  }
		}
	  }
	`)

	// set any variables
	// req.Var("key", "value")

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "tree-farmer")
	req.Header.Set("Authorization", "bearer 6514ab070a2d53e6ff1e9b98c6a3184faec2142f")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	// var respData ResponseStruct
	respData := make(map[string]string)
	if err := client.Run(ctx, req, &respData); err != nil {
		fmt.Println("Error! ", err)
	}

	fmt.Println(respData)
}