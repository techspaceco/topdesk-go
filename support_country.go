package topdesk

import (
	"context"
	"path"
)

type CountryIterator struct {
	*ListIterator
}

func (i CountryIterator) Country() (*Country, error) {
	response := &Country{}
	err := i.decode(&response)
	return response, err
}

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (c Country) Ref() *Ref {
	return &Ref{ID: c.ID}
}

type ListCountriesRequest struct{}

func (rc RestClient) ListCountries(ctx context.Context, request *ListCountriesRequest) (*CountryIterator, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "countries")

	it, err := rc.list(ctx, &uri)
	return &CountryIterator{it}, err
}
