package mapper

type InfrastructureMapper struct {
	iInfrastructureDummyMapper IInfrastructureDummyMapper
}

func NewInfrastructureMapper(iInfrastructureDummyMapper IInfrastructureDummyMapper) *InfrastructureMapper {
	return &InfrastructureMapper{
		iInfrastructureDummyMapper: iInfrastructureDummyMapper,
	}
}
