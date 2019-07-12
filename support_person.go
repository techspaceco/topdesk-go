package topdesk

import (
	"context"
	"path"
)

type PeopleIterator struct {
	*ListIterator
}

func (i PeopleIterator) Person() (*Person, error) {
	person := &Person{}
	err := i.decode(&person)
	return person, err
}

type Person struct {
	ID                          string `json:"id,omitempty"`
	Status                      string `json:"status"`
	Surname                     string `json:"surName"` // Capitalization!?
	FirstName                   string `json:"firstName,omitempty"`
	FirstInitials               string `json:"firstInitials,omitempty"`
	Prefixes                    string `json:"prefixes,omitempty"`
	Gender                      string `json:"gender,omitempty"`
	EmployeeNumber              string `json:"employeeNumber,omitempty"`
	ClientReferenceNumber       string `json:"clientReferenceNumber,omitempty"`
	NetworkLoginName            string `json:"networkLoginName,omitempty"`
	Branch                      *Ref   `json:"branch,omitempty"`
	Location                    *Ref   `json:"location,omitempty"`
	DepartmentFree              string `json:"departmentFree,omitempty"`
	TasLoginName                string `json:"tasLoginName,omitempty"`
	Password                    string `json:"password,omitempty"`
	PhoneNumber                 string `json:"phoneNumber,omitempty"`
	MobileNumber                string `json:"mobileNumber,omitempty"`
	Fax                         string `json:"fax,omitempty"`
	Email                       string `json:"email,omitempty"`
	JobTitle                    string `json:"jobTitle,omitempty"`
	ShowBudgetholder            bool   `json:"showBudgetholder,omitempty"`
	ShowDepartment              bool   `json:"showDepartment,omitempty"`
	ShowBranch                  bool   `json:"showBranch,omitempty"`
	ShowSubsidiaries            bool   `json:"showSubsidiaries,omitempty"`
	ShowAllBranches             bool   `json:"showAllBranches,omitempty"`
	AuthorizeAll                bool   `json:"authorizeAll,omitempty"`
	AuthorizeDepartment         bool   `json:"authorizeDepartment,omitempty"`
	AuthorizeBudgetHolder       bool   `json:"authorizeBudgetHolder,omitempty"`
	AuthorizeBranch             bool   `json:"authorizeBranch,omitempty"`
	AuthorizeSubsidiaryBranches bool   `json:"authorizeSubsidiaryBranches,omitempty"`
}

func (p Person) Ref() *Ref {
	return &Ref{ID: p.ID}
}

type ListPeopleRequest struct{}

func (rc RestClient) ListPeople(ctx context.Context, request *ListPeopleRequest) (*PeopleIterator, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "persons")

	it, err := rc.list(ctx, &uri)
	return &PeopleIterator{it}, err
}

func (rc RestClient) GetPerson(ctx context.Context, id string) (*Person, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "persons", "id", id)

	response := &Person{}
	err := rc.get(ctx, &uri, response)
	return response, err
}
