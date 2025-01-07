package mbta

type MBTAStreamSource struct {
	url         string
	apiKey      string
	distributor StreamDistributor
}

func NewMBTAStreamSource(url, apiKey string, distributor StreamDistributor) *MBTAStreamSource {
	return &MBTAStreamSource{
		url:         url,
		apiKey:      apiKey,
		distributor: distributor,
	}
}
