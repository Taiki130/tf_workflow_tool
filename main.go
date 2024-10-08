package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hclparse"
)

type provider struct {
	Provider []config `hcl:"provider,block"`
}

type config struct {
	RegistryURL string   `hcl:"registryURL,label"`
	Version     string   `hcl:"version"`
	Constraints string   `hcl:"constraints,optional"`
	Hashes      []string `hcl:"hashes"`
}

func main() {
	parser := hclparse.NewParser()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	lockFilePath := dir + "/.terraform.lock.hcl"
	f, parseDiags := parser.ParseHCLFile(lockFilePath)
	if parseDiags.HasErrors() {
		log.Fatal(parseDiags.Error())
	}

	var provider provider
	decodeDiags := gohcl.DecodeBody(f.Body, nil, &provider)
	if decodeDiags.HasErrors() {
		log.Fatal(decodeDiags.Error())
	}

	registryURL := provider.Provider[0].RegistryURL
	if strings.Contains(registryURL, "terraform") {
		fmt.Print("terraform")
	} else if strings.Contains(registryURL, "tofu") {
		fmt.Print("tofu")
	} else {
		log.Fatal("Failed to determine tool")
	}
}
