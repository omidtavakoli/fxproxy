package proxy

func (s service) Company() (string, error) {
	return s.app.CompanyAPI()
}
