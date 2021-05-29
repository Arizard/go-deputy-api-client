package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/arizard/go-deputy-api-client/pkg/deputy"
	"github.com/arizard/go-deputy-api-client/pkg/deputy/codeupdate"
)

var subdomain string
var bearer string
var coffeePath string

func main() {
	subdomain = os.Getenv("EXAMPLE_SUBDOMAIN")
	bearer = os.Getenv("EXAMPLE_BEARER")
	coffeePath = os.Getenv("EXAMPLE_DECAF")

	// Create a new Deputy API V1Client
	var dc deputy.Client = deputy.NewV1Client(
		subdomain,
		deputy.NewBearerTokenRequestAuthoriser(bearer),
	)

	var coffee []byte
	if out, err := ioutil.ReadFile(coffeePath); err != nil {
		fmt.Println(err)
		return
	} else {
		coffee = out
	}

	decafBase64 := base64.StdEncoding.EncodeToString(coffee)

	scriptId := "arie_test_codeupdate"
	codeupdateOptions := codeupdate.ScriptOptions{
		DecafBase64: decafBase64,
		Label:       "Arie Test Codeupdate Go Client",
		ScriptType:  codeupdate.ExecuteViaApi,
	}
	var codeWasUpdated = false
	if err := dc.CodeUpdateScript(scriptId, codeupdateOptions, &codeWasUpdated); err != nil {
		fmt.Println(err)
	}

	var scriptResponse struct {
		Message   string
		Timestamp uint64
	}

	if err := dc.ExecDeXML(scriptId, []byte{}, &scriptResponse); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", scriptResponse)
	fmt.Printf("codeWasUpdated: %t\n", codeWasUpdated)
}
