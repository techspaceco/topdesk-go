package topdesk

// Supporting files.
//
// https://developers.topdesk.com/explorer/?page=supporting-files

import (
	"context"
	"encoding/json"
)

type BranchRef struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

func (b *BranchRef) Ref() *BranchRef {
	return &BranchRef{ID: b.ID}
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
	response := &Branch{}
	if err := rc.save(ctx, "branches", branch.ID, branch.DTO(), &response); err != nil {
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

func (b *Branch) DTO() *BranchDTO {
	head := b.HeadBranch
	if head != nil {
		head = head.Ref()
	}
	return &BranchDTO{
		Name:                  b.Name,
		Specification:         b.Specification,
		ClientReferenceNumber: b.ClientReferenceNumber,
		Phone:                 b.Phone,
		Fax:                   b.Fax,
		Email:                 b.Email,
		Website:               b.Website,
		BranchType:            b.BranchType,
		HeadBranch:            head,
		Address:               b.Address,
		PostalAddress:         b.PostalAddress,
	}
}

type BranchDTO struct {
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
	response := &Person{}
	if err := rc.save(ctx, "persons", person.ID, person.DTO(), response); err != nil {
		return "", err
	}
	return response.ID, nil
}

// Person DTO.
//
// This isn't generated from the swagger which has types for each supported verbs request/reply type.
// Not all fields are supported.
type Person struct {
	ID                    string       `json:"id,omitempty"`
	Status                string       `json:"status"`
	Surname               string       `json:"surName"` // Capitalization!?
	FirstName             string       `json:"firstName,omitempty"`
	FirstInitials         string       `json:"firstInitials,omitempty"`
	Prefixes              string       `json:"prefixes,omitempty"`
	Gender                string       `json:"gender,omitempty"`
	EmployeeNumber        string       `json:"employeeNumber,omitempty"`
	ClientReferenceNumber string       `json:"clientReferenceNumber,omitempty"`
	NetworkLoginName      string       `json:"networkLoginName,omitempty"`
	Branch                *BranchRef   `json:"branch,omitempty"`
	Location              *LocationRef `json:"location,omitempty"`
	// Department       struct {
	// 	ID string `json:"id,omitempty"`
	// } `json:"department,omitempty"`
	// Language struct {
	// 	ID string `json:"id,omitempty"`
	// } `json:"language,omitempty"`
	DepartmentFree              string             `json:"departmentFree,omitempty"`
	TasLoginName                string             `json:"tasLoginName,omitempty"`
	Password                    string             `json:"password,omitempty"`
	PhoneNumber                 string             `json:"phoneNumber,omitempty"`
	MobileNumber                string             `json:"mobileNumber,omitempty"`
	Fax                         string             `json:"fax,omitempty"`
	Email                       string             `json:"email,omitempty"`
	JobTitle                    string             `json:"jobTitle,omitempty"`
	ShowBudgetholder            bool               `json:"showBudgetholder,omitempty"`
	ShowDepartment              bool               `json:"showDepartment,omitempty"`
	ShowBranch                  bool               `json:"showBranch,omitempty"`
	ShowSubsidiaries            bool               `json:"showSubsidiaries,omitempty"`
	ShowAllBranches             bool               `json:"showAllBranches,omitempty"`
	AuthorizeAll                bool               `json:"authorizeAll,omitempty"`
	AuthorizeDepartment         bool               `json:"authorizeDepartment,omitempty"`
	AuthorizeBudgetHolder       bool               `json:"authorizeBudgetHolder,omitempty"`
	AuthorizeBranch             bool               `json:"authorizeBranch,omitempty"`
	AuthorizeSubsidiaryBranches bool               `json:"authorizeSubsidiaryBranches,omitempty"`
	PersonExtraFieldA           *PersonExtraFieldA `json:"personExtraFieldA,omitempty"`
}

func (s *Person) DTO() *PersonDTO {
	// The other way to do this would be composition but you can't initialize struct literals with promoted fields in Go.
	//
	// Person {
	//   PersonDTO
	//   ID string
	//   CustomerReferenceNumber string
	// }
	//
	// Person {
	//   ID: ...,
	//   PersonDTO {
	//     FirstName: ...,
	//     ...
	//   }
	// }

	// Painfully the ref structure returned by GET isn't suitable for a PUT.
	branch := s.Branch
	if branch != nil {
		branch = branch.Ref()
	}
	extraFieldA := s.PersonExtraFieldA
	if extraFieldA != nil {
		extraFieldA = extraFieldA.Ref()
	}
	location := s.Location
	if location != nil {
		location = location.Ref()
	}

	return &PersonDTO{
		Surname:                     s.Surname,
		FirstName:                   s.FirstName,
		FirstInitials:               s.FirstInitials,
		Prefixes:                    s.Prefixes,
		Gender:                      s.Gender,
		EmployeeNumber:              s.EmployeeNumber,
		NetworkLoginName:            s.NetworkLoginName,
		Branch:                      branch,
		Location:                    location,
		DepartmentFree:              s.DepartmentFree,
		TasLoginName:                s.TasLoginName,
		Password:                    s.Password,
		PhoneNumber:                 s.PhoneNumber,
		MobileNumber:                s.MobileNumber,
		Fax:                         s.Fax, // Fax? Where is the pager field :P
		Email:                       s.Email,
		JobTitle:                    s.JobTitle,
		ShowBudgetholder:            s.ShowBudgetholder,
		ShowDepartment:              s.ShowDepartment,
		ShowBranch:                  s.ShowBranch,
		ShowSubsidiaries:            s.ShowSubsidiaries,
		ShowAllBranches:             s.ShowAllBranches,
		AuthorizeAll:                s.AuthorizeAll,
		AuthorizeDepartment:         s.AuthorizeDepartment,
		AuthorizeBudgetHolder:       s.AuthorizeBudgetHolder,
		AuthorizeBranch:             s.AuthorizeBranch,
		AuthorizeSubsidiaryBranches: s.AuthorizeSubsidiaryBranches,
		PersonExtraFieldA:           extraFieldA,
	}
}

// DTO's exist for each Request/Response type in the swagger but it's a pain to work like that in a typed language.
type PersonDTO struct {
	Surname          string       `json:"surName"` // Capitalization!?
	FirstName        string       `json:"firstName,omitempty"`
	FirstInitials    string       `json:"firstInitials,omitempty"`
	Prefixes         string       `json:"prefixes,omitempty"`
	Gender           string       `json:"gender,omitempty"`
	EmployeeNumber   string       `json:"employeeNumber,omitempty"`
	NetworkLoginName string       `json:"networkLoginName,omitempty"`
	Branch           *BranchRef   `json:"branch,omitempty"`
	Location         *LocationRef `json:"location,omitempty"`
	// Department       struct {
	// 	ID string `json:"id,omitempty"`
	// } `json:"department,omitempty"`
	// Language struct {
	// 	ID string `json:"id,omitempty"`
	// } `json:"language,omitempty"`
	DepartmentFree              string             `json:"departmentFree,omitempty"`
	TasLoginName                string             `json:"tasLoginName,omitempty"`
	Password                    string             `json:"password,omitempty"`
	PhoneNumber                 string             `json:"phoneNumber,omitempty"`
	MobileNumber                string             `json:"mobileNumber,omitempty"`
	Fax                         string             `json:"fax,omitempty"`
	Email                       string             `json:"email,omitempty"`
	JobTitle                    string             `json:"jobTitle,omitempty"`
	ShowBudgetholder            bool               `json:"showBudgetholder,omitempty"`
	ShowDepartment              bool               `json:"showDepartment,omitempty"`
	ShowBranch                  bool               `json:"showBranch,omitempty"`
	ShowSubsidiaries            bool               `json:"showSubsidiaries,omitempty"`
	ShowAllBranches             bool               `json:"showAllBranches,omitempty"`
	AuthorizeAll                bool               `json:"authorizeAll,omitempty"`
	AuthorizeDepartment         bool               `json:"authorizeDepartment,omitempty"`
	AuthorizeBudgetHolder       bool               `json:"authorizeBudgetHolder,omitempty"`
	AuthorizeBranch             bool               `json:"authorizeBranch,omitempty"`
	AuthorizeSubsidiaryBranches bool               `json:"authorizeSubsidiaryBranches,omitempty"`
	PersonExtraFieldA           *PersonExtraFieldA `json:"personExtraFieldA,omitempty"`
}

type LocationRef struct {
	ID         string     `json:"id"`
	Name       string     `json:"name,omitempty"`
	RoomNumber string     `json:"roomNumber,omitempty"`
	Branch     *BranchRef `json:"branch,omitempty"`
}

func (l *LocationRef) Ref() *LocationRef {
	return &LocationRef{ID: l.ID}
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

type PersonExtraFieldA struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (l PersonExtraFieldA) Ref() *PersonExtraFieldA {
	return &PersonExtraFieldA{ID: l.ID}
}

type PersonExtraFieldAIterator struct {
	*ListIterator
}

func (rc *RestClient) ListPersonExtraFieldA(ctx context.Context) (*PersonExtraFieldAIterator, error) {
	it, err := rc.list(ctx, "personExtraFieldAEntries")
	return &PersonExtraFieldAIterator{it}, err
}

func (i *PersonExtraFieldAIterator) PersonExtraFieldA() (*PersonExtraFieldA, error) {
	response := &PersonExtraFieldA{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return response, nil
}

func (rc *RestClient) SavePersonExtraFieldA(ctx context.Context, extra *PersonExtraFieldA) (string, error) {
	response := *extra
	if err := rc.save(ctx, "personExtraFieldAEntries", "", extra, &response); err != nil {
		return "", err
	}
	return response.ID, nil
}

type Operator struct {
	ID          string `json:"id"`
	PrincipalID string `json:"principalId"`
	Status      string `json:"status"`
	SurName     string `json:"surName"`
	FirstName   string `json:"firstName"`
	DynamicName string `json:"dynamicName"`
	Initials    string `json:"initials"`
	Prefixes    string `json:"prefixes"`
	BirthName   string `json:"birthName"`
	Title       string `json:"title"`
	Gender      string `json:"gender"`
	Language    struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"language"`
	Branch struct {
		ID                    string `json:"id"`
		Name                  string `json:"name"`
		ClientReferenceNumber string `json:"clientReferenceNumber"`
		TimeZone              string `json:"timeZone"`
		ExtraA                struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"extraA"`
		ExtraB struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"extraB"`
	} `json:"branch"`
	Location struct {
		ID     string `json:"id"`
		Branch struct {
			ID                    string `json:"id"`
			Name                  string `json:"name"`
			ClientReferenceNumber string `json:"clientReferenceNumber"`
			TimeZone              string `json:"timeZone"`
			ExtraA                struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"extraA"`
			ExtraB struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"extraB"`
		} `json:"branch"`
		Name string `json:"name"`
		Room string `json:"room"`
	} `json:"location"`
	Telephone       string `json:"telephone"`
	MobileNumber    string `json:"mobileNumber"`
	FaxNumber       string `json:"faxNumber"`
	Email           string `json:"email"`
	ExchangeAccount string `json:"exchangeAccount"`
	LoginName       string `json:"loginName"`
	LoginPermission bool   `json:"loginPermission"`
	JobTitle        string `json:"jobTitle"`
	Department      struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"department"`
	BudgetHolder struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"budgetHolder"`
	EmployeeNumber     string      `json:"employeeNumber"`
	HourlyRate         json.Number `json:"hourlyRate"`
	NetworkLoginName   string      `json:"networkLoginName"`
	MainframeLoginName string      `json:"mainframeLoginName"`
	HasAttention       bool        `json:"hasAttention"`
	Attention          struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"attention"`
	Comments                 string `json:"comments"`
	Installer                bool   `json:"installer"`
	FirstLineCallOperator    bool   `json:"firstLineCallOperator"`
	SecondLineCallOperator   bool   `json:"secondLineCallOperator"`
	ProblemManager           bool   `json:"problemManager"`
	ProblemOperator          bool   `json:"problemOperator"`
	ChangeCoordinator        bool   `json:"changeCoordinator"`
	ChangeActivitiesOperator bool   `json:"changeActivitiesOperator"`
	RequestForChangeOperator bool   `json:"requestForChangeOperator"`
	ExtensiveChangeOperator  bool   `json:"extensiveChangeOperator"`
	SimpleChangeOperator     bool   `json:"simpleChangeOperator"`
	ScenarioManager          bool   `json:"scenarioManager"`
	PlanningActivityManager  bool   `json:"planningActivityManager"`
	ProjectCoordinator       bool   `json:"projectCoordinator"`
	ProjectActiviesOperator  bool   `json:"projectActiviesOperator"`
	StockManager             bool   `json:"stockManager"`
	ReservationsOperator     bool   `json:"reservationsOperator"`
	ServiceOperator          bool   `json:"serviceOperator"`
	ExternalHelpDeskParty    bool   `json:"externalHelpDeskParty"`
	ContractManager          bool   `json:"contractManager"`
	OperationsOperator       bool   `json:"operationsOperator"`
	OperationsManager        bool   `json:"operationsManager"`
	KnowledgeBaseManager     bool   `json:"knowledgeBaseManager"`
	AccountManager           bool   `json:"accountManager"`
	CreationDate             string `json:"creationDate"`
	Creator                  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"creator"`
	ModificationDate string `json:"modificationDate"`
	Modifier         struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"modifier"`
}

type OperatorIterator struct {
	*ListIterator
}

func (rc *RestClient) ListOperators(ctx context.Context) (*OperatorIterator, error) {
	it, err := rc.list(ctx, "operators")
	return &OperatorIterator{it}, err
}

func (i *OperatorIterator) Operator() (*Operator, error) {
	response := &Operator{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return response, nil
}

func (rc *RestClient) GetOperator(ctx context.Context, id string) (*Operator, error) {
	operator := &Operator{}
	err := rc.get(ctx, "operators", id, operator)
	return operator, err
}
