package clientgenv2

import (

	// nolint:golint, nolintlint
	_ "embed"
	"fmt"

	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	gqlgencConfig "github.com/pettinz/gqlgenc/config"
)

//go:embed template.gotpl
var template string

func RenderTemplate(cfg *config.Config, query *Query, mutation *Mutation, fragments []*Fragment, operations []*Operation, operationResponses []*OperationResponse, structSources []*StructSource, generateCfg *gqlgencConfig.GenerateConfig, client config.PackageConfig) error {
	if err := templates.Render(templates.Options{
		PackageName: client.Package,
		Filename:    client.Filename,
		Template:    template,
		Data: map[string]interface{}{
			"Query":               query,
			"Mutation":            mutation,
			"Fragment":            fragments,
			"Operation":           operations,
			"OperationResponse":   operationResponses,
			"GenerateClient":      generateCfg.ShouldGenerateClient(),
			"StructSources":       structSources,
			"ClientInterfaceName": generateCfg.GetClientInterfaceName(),
		},
		Packages:   cfg.Packages,
		PackageDoc: "// Code generated by github.com/pettinz/gqlgenc, DO NOT EDIT.\n",
	}); err != nil {
		return fmt.Errorf("%s generating failed: %w", client.Filename, err)
	}

	return nil
}
