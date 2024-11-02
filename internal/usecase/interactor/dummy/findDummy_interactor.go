package dummy

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/repository/mysql/scope"
	usecase "github.com/LeastKIds/go_struct/internal/usecase/dto"
	"github.com/LeastKIds/go_struct/internal/usecase/mapper"
)

func (di *DummyInteractor) FindDummy(id int64) (*usecase.Dummy, error) {
	dummy, err := di.DummyRepo.FindBy(scope.ByID(id))
	if err != nil {
		return nil, err
	}

	return mapper.NewDummyMapper().ToUsecase(dummy), nil
}
