package service

type CompositeStreamManager struct {
	source      StreamSource
	distributor StreamDistributor
}

func NewCompositeStreamManager(source StreamSource, distributor StreamDistributor) *CompositeStreamManager {
	return &CompositeStreamManager{
		source:      source,
		distributor: distributor,
	}
}
