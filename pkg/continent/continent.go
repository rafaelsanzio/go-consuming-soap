package continent

type ContinentConsumer interface {
	List() ([]Continent, error)
}
