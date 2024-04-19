package main

import (
	"fmt"
	"github.com/kairos-io/kairos-sdk/signatures"
	"os"
	"strings"
)

func main() {
	certs, err := signatures.GetAllFullCerts()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for index, cert := range certs.PK {
		commonName := strings.Replace(cert.Subject.CommonName, " ", "_", -1)
		filename := fmt.Sprintf("PK-%s-%d.der", commonName, index)
		err := os.WriteFile(filename, cert.Raw, 777)
		if err != nil {
			fmt.Printf("Could not write file at %s: %s\n", filename, err)
			continue
		}
		fmt.Printf("Wrote file at %s\n", filename)
	}

	for index, cert := range certs.KEK {
		commonName := strings.Replace(cert.Subject.CommonName, " ", "_", -1)
		filename := fmt.Sprintf("KEK-%s-%d.der", commonName, index)
		err := os.WriteFile(filename, cert.Raw, 777)
		if err != nil {
			fmt.Printf("Could not write file at %s: %s\n", filename, err)
			continue
		}
		fmt.Printf("Wrote file at %s\n", filename)
	}

	for index, cert := range certs.DB {
		commonName := strings.Replace(cert.Subject.CommonName, " ", "_", -1)
		filename := fmt.Sprintf("DB-%s-%d.der", commonName, index)
		err := os.WriteFile(filename, cert.Raw, 777)
		if err != nil {
			fmt.Printf("Could not write file at %s: %s\n", filename, err)
			continue
		}
		fmt.Printf("Wrote file at %s\n", filename)
	}
}
