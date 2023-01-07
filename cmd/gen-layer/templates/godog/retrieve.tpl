{{define "retrieve"}}
package features

import (
       "context"
       "github.com/cucumber/godog"
)

func (s*Suite) {{.CamelCase}}IsDeleted(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}

func (s*Suite) userRetrieve{{.PascalCase}}(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}
{{end}}