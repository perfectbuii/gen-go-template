{{define "list"}}
package features

import (
       "context"
       "github.com/cucumber/godog"
)

func (s*Suite) userList{{.PascalCase}}(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}

func (s*Suite) ourSystemMustReturn{{.PascalCase}}Correctly(ctx context.Context) (context.Context, error) {
       stepState := StepStateFromContext(ctx)
       return StepStateToContext(ctx, stepState), godog.ErrPending
}
{{end}}