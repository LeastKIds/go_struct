package mapper

type AdapterMapper struct {
	AdapterDummyMapper IAdapterDummyMapper
}

func NewAdapterMapper(adapterDummyMapper IAdapterDummyMapper) *AdapterMapper {
	return &AdapterMapper{AdapterDummyMapper: adapterDummyMapper}
}
