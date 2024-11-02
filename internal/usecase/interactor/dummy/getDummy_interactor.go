package dummy

import (
	"github.com/LeastKIds/go_struct/internal/infrastructure/database/mysql/scope"
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"gorm.io/gorm"
)

func (i *InteractorDummy) GetDummy(db *gorm.DB, id int64) (*dto.Dummy, error) {
	dummy, err := i.repo.InfrastructureDummyRepo.FindBy(db, scope.ByID(id))
	if err != nil {
		return nil, err
	}

	return i.mapper.UsecaseDummyMapper.ToDTO(dummy), nil
}
