{{define "update"}}
package features

import (
       "context"
       "github.com/cucumber/godog"
)

func (s*Suite) updated{{.PascalCase}}SetAsExpected(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}

func (s*Suite) userUpdate{{.PascalCase}}(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}
{{end}}