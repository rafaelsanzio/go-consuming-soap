package country

type CountryConsumer interface {
	List() ([]Country, error)
	Info(countryCode string) (*CountryInfo, error)
}

type CountryInfo struct {
	Capital   string
	Currency  Currency
	Flag      string
	PhoneCode int
}
