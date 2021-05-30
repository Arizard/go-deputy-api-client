package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/arizard/go-deputy-api-client/pkg/deputy"
	"github.com/arizard/go-deputy-api-client/pkg/deputy/codeupdate"
)

func readAndEncode(path string) string {
	var contents []byte
	if out, err := ioutil.ReadFile(path); err != nil {
		fmt.Println(err)
		return ""
	} else {
		contents = out
	}

	return base64.StdEncoding.EncodeToString(contents)
}

var subdomain string
var bearer string

func main() {
	subdomain = os.Getenv("EXAMPLE_SUBDOMAIN")
	bearer = os.Getenv("EXAMPLE_BEARER")

	// Create a new Deputy API V1Client
	var dc deputy.Client = deputy.NewV1Client(
		subdomain,
		deputy.NewBearerTokenRequestAuthoriser(bearer),
	)

	appId := "arie_test_codeupdate2"
	codeupdateOptions := codeupdate.CustomAppOptions{
		Name:             "Arie Test Codeupdate Application",
		HTMLBase64:       readAndEncode(os.Getenv("EXAMPLE_REPORT_HTML")),
		JavascriptBase64: readAndEncode(os.Getenv("EXAMPLE_REPORT_JS")),
		DecafBase64:      readAndEncode(os.Getenv("EXAMPLE_REPORT_DECAF")),
		CSSBase64:        readAndEncode(os.Getenv("EXAMPLE_REPORT_CSS")),
	}
	var codeWasUpdated = false
	if err := dc.CodeUpdateCustomApp(appId, codeupdateOptions, &codeWasUpdated); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("codeWasUpdated: %t\n", codeWasUpdated)
}
