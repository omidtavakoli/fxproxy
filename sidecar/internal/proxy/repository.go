package proxy

type AppRepository interface {
	CompanyAPI() (string, error)
}
