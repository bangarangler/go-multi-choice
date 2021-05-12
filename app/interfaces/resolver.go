package interfaces

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	// "context"
	// "fmt"

	// "github.com/99designs/gqlgen/graphql/introspection"
	"github.com/bangarangler/go-multi-choice/app/domain/repository/answer"
	"github.com/bangarangler/go-multi-choice/app/domain/repository/question"
	"github.com/bangarangler/go-multi-choice/app/domain/repository/question_option"
	// "github.com/bangarangler/go-multi-choice/graph/generated"
)

type Resolver struct {
	AnsService            answer.AnsService
	QuestionService       question.QuesService
	QuestionOptionService question_option.OptService
	// Directives            DirectiveRoot
}

// func (r *__DirectiveResolver) IsRepeatable(ctx context.Context, obj *introspection.Directive) (bool, error) {
// 	panic(fmt.Errorf("not implemented"))
// }

// __Directive returns generated.__DirectiveResolver implementation.
// func (r *Resolver) __Directive() generated.__DirectiveResolver { return &__DirectiveResolver{r} }

// type __DirectiveResolver struct{ *Resolver }
//
// type DirectiveRoot struct{}
