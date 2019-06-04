package topdesk

// Supporting files.
//
// https://developers.topdesk.com/explorer/?page=supporting-files

import (
	"context"
)

type BranchRef struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type BranchIterator struct {
	*ListIterator
}

func (rc *RestClient) ListBranches(ctx context.Context) (*BranchIterator, error) {
	it, err := rc.list(ctx, "branches")
	return &BranchIterator{it}, err
}

func (i *BranchIterator) Branch() (*Branch, error) {
	response := &BranchRef{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetBranch(i.ctx, response.ID)
}

func (rc *RestClient) GetBranch(ctx context.Context, id string) (*Branch, error) {
	branch := &Branch{}
	err := rc.get(ctx, "branches", id, branch)
	return branch, err
}

func (rc *RestClient) SaveBranch(ctx context.Context, branch *Branch) (string, error) {
	response := *branch
	if err := rc.save(ctx, "branches", branch.ID, branch, &response); err != nil {
		return "", err
	}
	return response.ID, nil
}

func (b Branch) Ref() *BranchRef {
	return &BranchRef{ID: b.ID}
}

// Branch structure.
//
// https://developers.topdesk.com/explorer/?page=supporting-files#/Branches/retrieveBranches
type Branch struct {
	ID                    string     `json:"id,omitempty"`
	Name                  string     `json:"name"`
	Specification         string     `json:"specification"`
	ClientReferenceNumber string     `json:"clientReferenceNumber"`
	Phone                 string     `json:"phone"`
	Fax                   string     `json:"fax"`
	Email                 string     `json:"email"`
	Website               string     `json:"website"`
	BranchType            string     `json:"branchType"`
	HeadBranch            *BranchRef `json:"headBranch"`
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

type CountryIterator struct {
	*ListIterator
}

func (rc *RestClient) ListCountries(ctx context.Context) (*CountryIterator, error) {
	it, err := rc.list(ctx, "countries")
	return &CountryIterator{it}, err
}

func (i *CountryIterator) Country() (*Country, error) {
	country := &Country{}
	err := i.decode(&country)
	return country, err
}

type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PeopleIterator struct {
	*ListIterator
}

func (rc *RestClient) ListPeople(ctx context.Context) (*PeopleIterator, error) {
	it, err := rc.list(ctx, "persons")
	return &PeopleIterator{it}, err
}

func (i *PeopleIterator) Person() (*Person, error) {
	person := &Person{}
	err := i.decode(&person)
	return person, err
}

func (rc *RestClient) GetPerson(ctx context.Context, id string) (*Person, error) {
	person := &Person{}
	err := rc.get(ctx, "persons", id, person)
	return person, err
}

func (rc *RestClient) SavePerson(ctx context.Context, person *Person) (string, error) {
	if len(person.Gender) == 0 { // Requires a value.
		// person.Gender = "UNDEFINED"
	}

	response := *person
	if err := rc.save(ctx, "persons", person.ID, person, &response); err != nil {
		return "", err
	}
	return response.ID, nil
}

type Person struct {
	ID                    string       `json:"id,omitempty"`
	Status                string       `json:"status,omitempty"`
	Surname               string       `json:"surName"` // Capitalization!?
	FirstName             string       `json:"firstName,omitempty"`
	FirstInitials         string       `json:"firstInitials,omitempty"`
	Prefixes              string       `json:"prefixes,omitempty"`
	Gender                string       `json:"gender,omitempty"`
	EmployeeNumber        string       `json:"employeeNumber,omitempty"`
	NetworkLoginName      string       `json:"networkLoginName,omitempty"`
	ClientReferenceNumber string       `json:"clientReferenceNumber,omitempty"`
	Branch                *BranchRef   `json:"branch,omitempty"`
	Location              *LocationRef `json:"location,omitempty"`
	// Department       struct {
	// 	ID string `json:"id,omitempty"`
	// } `json:"department,omitempty"`
	// Language struct {
	// 	ID string `json:"id,omitempty"`
	// } `json:"language,omitempty"`
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

	// TODO(shane): Techspace specific.
	MemberCompany *MemberCompany `json:"personExtraFieldA,omitempty"`
}

type LocationRef struct {
	ID         string     `json:"id"`
	Name       string     `json:"name,omitempty"`
	RoomNumber string     `json:"roomNumber,omitempty"`
	Branch     *BranchRef `json:"branch,omitempty"`
}

type LocationIterator struct {
	*ListIterator
}

func (rc *RestClient) ListLocations(ctx context.Context) (*LocationIterator, error) {
	it, err := rc.list(ctx, "locations")
	return &LocationIterator{it}, err
}

func (i *LocationIterator) Location() (*Location, error) {
	response := &LocationRef{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetLocation(i.ctx, response.ID)
}

func (rc *RestClient) GetLocation(ctx context.Context, id string) (*Location, error) {
	location := &Location{}
	err := rc.get(ctx, "locations", id, location)
	return location, err
}

func (l Location) Ref() *LocationRef {
	return &LocationRef{ID: l.ID}
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

type MemberCompany struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (l MemberCompany) Ref() *MemberCompany {
	return &MemberCompany{ID: l.ID}
}

type MemberCompanyIterator struct {
	*ListIterator
}

func (rc *RestClient) ListMemberCompanies(ctx context.Context) (*MemberCompanyIterator, error) {
	it, err := rc.list(ctx, "personExtraFieldAEntries")
	return &MemberCompanyIterator{it}, err
}

func (i *MemberCompanyIterator) MemberCompany() (*MemberCompany, error) {
	response := &MemberCompany{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return response, nil
}

func (rc *RestClient) SaveMemberCompany(ctx context.Context, company *MemberCompany) (string, error) {
	response := *company
	if err := rc.save(ctx, "personExtraFieldAEntries", "", company, &response); err != nil {
		return "", err
	}
	return response.ID, nil
}
