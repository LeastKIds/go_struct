package mapper

type InfrastructureMapper struct {
	InfrastructureDummyMapper IInfrastructureDummyMapper
}

func NewInfrastructureMapper(iInfrastructureDummyMapper IInfrastructureDummyMapper) *InfrastructureMapper {
	return &InfrastructureMapper{
		InfrastructureDummyMapper: iInfrastructureDummyMapper,
	}
}
