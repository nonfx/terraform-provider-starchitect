package main

import (
	"context"
	"log"

	"terraform-provider-starchitect/starchitect"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary.
	version string = "dev"

	// goreleaser can pass other information to the main package, such as the specific commit
	// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "hashicorp.com/edu/starchitect",
	}
	if err := providerserver.Serve(context.Background(), starchitect.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}
