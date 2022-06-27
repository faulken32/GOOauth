package realms

type EndPointsService interface {
	Save() (*Endpoint, error)
	TruncateTable()
}
