package dummy

import (
	"github.com/LeastKIds/go_struct/internal/usecase/dto"
	"gorm.io/gorm"
)

func (i *DummyInteractor) FindDummyById(db *gorm.DB, id int64) (*dto.Dummy, error) {
	return i.repo.Dummy.FindBy(db, scope.ByID(id))
}
