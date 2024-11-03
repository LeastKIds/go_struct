package scope

import "github.com/LeastKIds/go_struct/internal/domain/repository"

type Scope struct{}

func NewScope() repository.IScopes {
	return &Scope{}
}
