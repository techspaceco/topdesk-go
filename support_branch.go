package topdesk

// Supporting files.
//
// https://developers.topdesk.com/explorer/?page=supporting-files

import (
	"context"
	"path"
)

type BranchIterator struct {
	*ListIterator
}

func (i BranchIterator) Branch() (*Branch, error) {
	response := &Ref{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetBranch(i.ctx, response.ID)
}

// Branch structure.
//
// https://developers.topdesk.com/explorer/?page=supporting-files#/Branches/retrieveBranches
type Branch struct {
	ID                    string `json:"id,omitempty"`
	Name                  string `json:"name"`
	Specification         string `json:"specification"`
	ClientReferenceNumber string `json:"clientReferenceNumber"`
	Phone                 string `json:"phone"`
	Fax                   string `json:"fax"`
	Email                 string `json:"email"`
	Website               string `json:"website"`
	BranchType            string `json:"branchType"`
	HeadBranch            *Ref   `json:"headBranch"`
	Address               struct {
		Country struct {
			ID string `json:"id"`
		} `json:"country"`
		Street      string `json:"street"`
		Number      string `json:"number"`
		County      string `json:"county"`
		City        string `json:"city"`
		Postcode    string `json:"postcode"`
		AddressMemo string `json:"addressMemo"`
	} `json:"address"`
	PostalAddress struct {
		Country struct {
			ID string `json:"id"`
		} `json:"country"`
		Street      string `json:"street"`
		Number      string `json:"number"`
		County      string `json:"county"`
		City        string `json:"city"`
		Postcode    string `json:"postcode"`
		AddressMemo string `json:"addressMemo"`
	} `json:"postalAddress"`
}

func (b Branch) Ref() *Ref {
	return &Ref{ID: b.ID}
}

type ListBranchesRequest struct{}

func (rc RestClient) ListBranches(ctx context.Context, request *ListBranchesRequest) (*BranchIterator, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "branches")

	it, err := rc.list(ctx, &uri)
	return &BranchIterator{it}, err
}

func (rc RestClient) GetBranch(ctx context.Context, id string) (*Branch, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "branches", "id", id)

	response := &Branch{}
	err := rc.get(ctx, &uri, response)
	return response, err
}
