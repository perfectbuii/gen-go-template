{{define "create"}}
package features

import (
       "context"
       "github.com/cucumber/godog"
)

func (s*Suite) {{.CamelCase}}MustBeCreated(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}

func (s*Suite) userCreate{{.PascalCase}}(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}
{{end}}