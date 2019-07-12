package topdesk

import (
	"context"
	"path"
)

type LocationIterator struct {
	*ListIterator
}

func (i LocationIterator) Location() (*Location, error) {
	response := &Ref{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetLocation(i.ctx, response.ID)
}

type Location struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name"`
	RoomNumber    string `json:"roomNumber"`
	FunctionalUse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"functionalUse"`
	Type struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Capacity      int    `json:"capacity"`
	Specification string `json:"specification"`
	BudgetHolder  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"budgetHolder"`
	Branch struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"branch"`
}

func (l Location) Ref() *Ref {
	return &Ref{ID: l.ID}
}

type ListLocationsRequest struct{}

func (rc RestClient) ListLocations(ctx context.Context, request *ListLocationsRequest) (*LocationIterator, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "locations")

	it, err := rc.list(ctx, &uri)
	return &LocationIterator{it}, err
}

func (rc RestClient) GetLocation(ctx context.Context, id string) (*Location, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "locations", "id", id)

	response := &Location{}
	err := rc.get(ctx, &uri, response)
	return response, err
}
