package topdesk

// Supporting files.
//
// https://developers.topdesk.com/explorer/?page=supporting-files

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
)

type BranchIterator struct {
	*ListIterator
}

func (c *Client) ListBranches(ctx context.Context) (*BranchIterator, error) {
	it, err := c.list(ctx, "branches")
	return &BranchIterator{it}, err
}

func (i *BranchIterator) Branch() (*Branch, error) {
	response := struct {
		ID string `json:"id"`
	}{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetBranch(i.ctx, response.ID)
}

func (c *Client) GetBranch(ctx context.Context, id string) (*Branch, error) {
	uri := *c.endpoint
	uri.Path = path.Join(uri.Path, fmt.Sprintf("branches/id/%s", id))

	branch := &Branch{}
	_, err := c.do(ctx, http.MethodGet, uri.String(), nil, branch)
	return branch, err
}

func (c *Client) SaveBranch(ctx context.Context, branch *Branch) (string, error) {
	method := http.MethodPost
	uri := *c.endpoint
	uri.Path = path.Join(uri.Path, "branches")

	if branch.ID != "" {
		method = http.MethodPut
		uri.Path = path.Join(uri.Path, fmt.Sprintf("id/%s", branch.ID))
	}

	response := &Branch{}
	status, err := c.do(ctx, method, uri.String(), branch, response)
	if err != nil {
		return "", err // TODO: Wrap.
	}

	switch status {
	case http.StatusCreated, http.StatusOK:
		return branch.ID, nil
	default:
		return "", fmt.Errorf("%s", http.StatusText(status))
	}
}

// Branch structure.
//
// https://developers.topdesk.com/explorer/?page=supporting-files#/Branches/retrieveBranches
type Branch struct {
	ID                    string `json:"id"`
	Name                  string `json:"name"`
	Specification         string `json:"specification"`
	ClientReferenceNumber string `json:"clientReferenceNumber"`
	Phone                 string `json:"phone"`
	Fax                   string `json:"fax"`
	Email                 string `json:"email"`
	Website               string `json:"website"`
	BranchType            string `json:"branchType"`
	HeadBranch            struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"headBranch"`
	Address struct {
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
	OptionalFields1 struct {
		Boolean1    bool        `json:"boolean1"`
		Boolean2    bool        `json:"boolean2"`
		Boolean3    bool        `json:"boolean3"`
		Boolean4    bool        `json:"boolean4"`
		Boolean5    bool        `json:"boolean5"`
		Number1     json.Number `json:"number1"`
		Number2     json.Number `json:"number2"`
		Number3     json.Number `json:"number3"`
		Number4     json.Number `json:"number4"`
		Number5     json.Number `json:"number5"`
		Date1       string      `json:"date1"`
		Date2       string      `json:"date2"`
		Date3       string      `json:"date3"`
		Date4       string      `json:"date4"`
		Date5       string      `json:"date5"`
		Text1       string      `json:"text1"`
		Text2       string      `json:"text2"`
		Text3       string      `json:"text3"`
		Text4       string      `json:"text4"`
		Text5       string      `json:"text5"`
		Searchlist1 struct {
			ID string `json:"id"`
		} `json:"searchlist1"`
		Searchlist2 struct {
			ID string `json:"id"`
		} `json:"searchlist2"`
		Searchlist3 struct {
			ID string `json:"id"`
		} `json:"searchlist3"`
		Searchlist4 struct {
			ID string `json:"id"`
		} `json:"searchlist4"`
		Searchlist5 struct {
			ID string `json:"id"`
		} `json:"searchlist5"`
	} `json:"optionalFields1"`
	OptionalFields2 struct {
		Boolean1    bool        `json:"boolean1"`
		Boolean2    bool        `json:"boolean2"`
		Boolean3    bool        `json:"boolean3"`
		Boolean4    bool        `json:"boolean4"`
		Boolean5    bool        `json:"boolean5"`
		Number1     json.Number `json:"number1"`
		Number2     json.Number `json:"number2"`
		Number3     json.Number `json:"number3"`
		Number4     json.Number `json:"number4"`
		Number5     json.Number `json:"number5"`
		Date1       string      `json:"date1"`
		Date2       string      `json:"date2"`
		Date3       string      `json:"date3"`
		Date4       string      `json:"date4"`
		Date5       string      `json:"date5"`
		Text1       string      `json:"text1"`
		Text2       string      `json:"text2"`
		Text3       string      `json:"text3"`
		Text4       string      `json:"text4"`
		Text5       string      `json:"text5"`
		Searchlist1 struct {
			ID string `json:"id"`
		} `json:"searchlist1"`
		Searchlist2 struct {
			ID string `json:"id"`
		} `json:"searchlist2"`
		Searchlist3 struct {
			ID string `json:"id"`
		} `json:"searchlist3"`
		Searchlist4 struct {
			ID string `json:"id"`
		} `json:"searchlist4"`
		Searchlist5 struct {
			ID string `json:"id"`
		} `json:"searchlist5"`
	} `json:"optionalFields2"`
}
