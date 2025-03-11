// Package main provides a command-line interface (CLI) for generating new modules in the application.
// It uses Cobra to manage commands and templates to generate module files (service, controller, and module definition).
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

// svcTemplate defines the template for the service layer of the generated module.
const svcTemplate = `package svc

import (
	"log/slog"
)

type {{.ModuleName}}Svc struct {
	logger *slog.Logger
}

func New{{.ModuleName}}Service(logger *slog.Logger) *{{.ModuleName}}Svc {
	return &{{.ModuleName}}Svc{
		logger: logger,
	}
}

func (svc *{{.ModuleName}}Svc) Ping(_ context.Context) (bool, error) {
	return true, nil
}
`

// ctrlTemplate defines the template for the controller layer of the generated module.
const ctrlTemplate = `package ctrl

import (
	"{{.RootModuleName}}/internal/pkg/{{.moduleName}}/svc"
	"log/slog"
)

type {{.ModuleName}}Ctrl struct {
	logger     *slog.Logger
	{{.moduleName}}Svc *svc.{{.ModuleName}}Svc
}

func New{{.ModuleName}}Controller(logger *slog.Logger, svc *svc.{{.ModuleName}}Svc) *{{.ModuleName}}Ctrl {
	return &{{.ModuleName}}Ctrl{
		logger:     logger,
		{{.moduleName}}Svc: svc,
	}
}

func (c *{{.ModuleName}}Ctrl) Ping(ctx *fiber.Ctx) error {
	res, err := c.{{.moduleName}}Svc.Ping(ctx.Context())
	if err != nil {
		return http.NewResponse(ctx, http.BadRequest, nil, http.CodeInternalError, "Something went wrong!")
	}
	return http.NewResponse(ctx, http.OK, res, 0, "")
}

`

// moduleTemplate defines the template for the module definition file of the generated module.
const moduleTemplate = `package {{.moduleName}}

import (
	"{{.RootModuleName}}/internal/pkg/{{.moduleName}}/ctrl"
	"{{.RootModuleName}}/internal/pkg/{{.moduleName}}/svc"
	"github.com/gofiber/fiber/v2"
	"log/slog"
)

type Module struct {
	Name    string
	Version string

	logger            *slog.Logger
	{{.ModuleName}}Controller *ctrl.{{.ModuleName}}Ctrl
}

func New(logger *slog.Logger) *Module {
	mod := &Module{
		Name:    "{{.moduleName}}-module",
		Version: "v1",
		logger:  logger,
		{{.ModuleName}}Controller: ctrl.New{{.ModuleName}}Controller(
			logger,
			svc.New{{.ModuleName}}Service(logger),
		),
	}
	return mod
}

func (m *Module) InitHTTPRoutes(r fiber.Router) {
	root := r.Group("/" + m.Name).
		Group("/api").
		Group("/" + m.Version)

	modGroup := root.Group("/{{.moduleName}}")
	modGroup.Get("/ping", m.{{.ModuleName}}Controller.Ping)
}
`

// newGenerateCommand creates a new Cobra command for generating modules.
//
// It defines the flag for the module name and handles the execution of the module generation.
func newGenerateCommand() *cobra.Command {
	var moduleName string

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Генерация модулей",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Ensure the module name is provided.
			if moduleName == "" {
				return fmt.Errorf("module name is empty")
			}

			// Generate the module based on the provided name.
			return generateModule(moduleName)
		},
	}

	// Add the module name flag.
	cmd.Flags().StringVarP(&moduleName, "module", "m", "", "module name (required)")
	return cmd
}

// getModuleName reads the go.mod file in the current directory and extracts the module name.
func getModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", fmt.Errorf("failed to open go.mod file: %w", err)
	}
	defer file.Close()

	// Read the first line of the go.mod file to extract the module name.
	var moduleName string
	_, err = fmt.Fscanf(file, "module %s", &moduleName)
	if err != nil {
		return "", fmt.Errorf("failed to read module name from go.mod: %w", err)
	}

	return moduleName, nil
}

// generateModule creates the necessary files for a new module based on the provided module name.
// It writes the service, controller, and module definition files.
func generateModule(moduleName string) error {
	// Get the root module name from go.mod.
	rootModuleName, err := getModuleName()
	if err != nil {
		return fmt.Errorf("could not get module name: %w", err)
	}

	// Define the path for the new module.
	modulePath := fmt.Sprintf("internal/pkg/%s", strings.ToLower(moduleName))

	// Define the files to generate and their corresponding templates.
	files := []struct {
		path     string
		template string
	}{
		{fmt.Sprintf("%s/svc/%s_svc.go", modulePath, strings.ToLower(moduleName)), svcTemplate},
		{fmt.Sprintf("%s/ctrl/%s_ctrl.go", modulePath, strings.ToLower(moduleName)), ctrlTemplate},
		{fmt.Sprintf("%s/%s_module.go", modulePath, strings.ToLower(moduleName)), moduleTemplate},
	}

	// Write the templates to files.
	for _, file := range files {
		if err := writeTemplate(file.path, file.template, moduleName, rootModuleName); err != nil {
			return err
		}
	}

	// Output a success message.
	fmt.Println("Module generated and router updated!")
	return nil
}

// writeTemplate writes a template to a file, creating the necessary directories if they do not exist.
func writeTemplate(path, tmpl, moduleName, rootModuleName string) error {
	// Create the directories for the file.
	if err := os.MkdirAll(getDir(path), os.ModePerm); err != nil {
		return err
	}

	// Create the file.
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse and execute the template.
	t := template.Must(template.New("file").Parse(tmpl))
	return t.Execute(file, map[string]string{
		"ModuleName":     strings.Title(moduleName),
		"moduleName":     strings.ToLower(moduleName),
		"RootModuleName": rootModuleName,
	})
}

// getDir returns the directory path from a file path.
func getDir(path string) string {
	parts := strings.Split(path, "/")
	return strings.Join(parts[:len(parts)-1], "/")
}

// main is the entry point of the application that sets up the CLI commands and executes them.
func main() {
	// Create the root command for the CLI application.
	rootCmd := &cobra.Command{
		Use:   "modulegen",
		Short: "CLI for new modules gen",
	}

	// Add the generate command for generating new modules.
	rootCmd.AddCommand(newGenerateCommand())

	// Execute the root command and handle errors.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
