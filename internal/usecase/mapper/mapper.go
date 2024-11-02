package mapper

type UsecaseMapper struct {
	UsecaseDummyMapper IUsecaseDummyMapper
}

func NewUsecaseMapper(usecaseDummyMapper IUsecaseDummyMapper) *UsecaseMapper {
	return &UsecaseMapper{UsecaseDummyMapper: usecaseDummyMapper}
}
