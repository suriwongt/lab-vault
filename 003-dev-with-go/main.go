package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
)

func main() {
	config := api.DefaultConfig()
	config.Address = "http://xxxx:8200"
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Print(err.Error())
	}

	username := "read"
	password := "xxxxxxxxx"

	// Authenticate using the userpass method
	resp, err := client.Logical().Write("auth/userpass/login/"+username, map[string]interface{}{
		"password": password,
	})

	if err != nil {
		panic(err)
	}
	// Set the Vault address
	client.SetToken(resp.Auth.ClientToken)
	secret, err := client.KVv2("xxxxxx").Get(context.Background(), "xxxxxx")
	if err != nil {
		panic(err)
	}

	if secret != nil {
		data := secret.Data
		for key, value := range data {
			envVarName := fmt.Sprintf("%s", strings.ToUpper(key))
			os.Setenv(envVarName, fmt.Sprintf("%v", value))
			fmt.Printf("Variable: %s=%v\n", envVarName, os.Getenv(envVarName))
		}
	}

}
