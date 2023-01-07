{{define "delete"}}
package features

import "github.com/cucumber/godog"

func {{.CamelCase}}HasBeenDeletedCorrectly(ctx context.Context) (context.Context, error) {
       return StepStateToContext(ctx, stepState), godog.ErrPending
}

func userDelete{{.PascalCase}}(ctx context.Context) (context.Context, error) {
       return StepStateToContext(ctx, stepState), godog.ErrPending
}

func userDelete{{.PascalCase}}Again(ctx context.Context) (context.Context, error) {
       return StepStateToContext(ctx, stepState), godog.ErrPending
}
{{end}}