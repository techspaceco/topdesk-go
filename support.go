package topdesk

// Supporting files.
//
// https://developers.topdesk.com/explorer/?page=supporting-files

import (
	"context"
	"encoding/json"
	"time"
)

type BranchIterator struct {
	*ListIterator
}

func (rc *RestClient) ListBranches(ctx context.Context) (*BranchIterator, error) {
	it, err := rc.list(ctx, "branches")
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
		Date1       time.Time   `json:"date1"`
		Date2       time.Time   `json:"date2"`
		Date3       time.Time   `json:"date3"`
		Date4       time.Time   `json:"date4"`
		Date5       time.Time   `json:"date5"`
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
	it, err := rc.list(ctx, "countries")
	return &PeopleIterator{it}, err
}

func (i *PeopleIterator) Person() (*Person, error) {
	person := &Person{}
	err := i.decode(&person)
	return person, err
}

type Person struct {
	ID               string `json:"id"`
	Status           string `json:"status"`
	Surname          string `json:"surName"` // Capitalization!?
	FirstName        string `json:"firstName"`
	FirstInitials    string `json:"firstInitials"`
	Prefixes         string `json:"prefixes"`
	Gender           string `json:"gender"`
	EmployeeNumber   string `json:"employeeNumber"`
	NetworkLoginName string `json:"networkLoginName"`
	Branch           struct {
		ID string `json:"id"`
	} `json:"branch"`
	Location struct {
		ID string `json:"id"`
	} `json:"location"`
	Department struct {
		ID string `json:"id"`
	} `json:"department"`
	Language struct {
		ID string `json:"id"`
	} `json:"language"`
	DepartmentFree              string `json:"departmentFree"`
	TasLoginName                string `json:"tasLoginName"`
	Password                    string `json:"password"`
	PhoneNumber                 string `json:"phoneNumber"`
	MobileNumber                string `json:"mobileNumber"`
	Fax                         string `json:"fax"`
	Email                       string `json:"email"`
	JobTitle                    string `json:"jobTitle"`
	ShowBudgetholder            bool   `json:"showBudgetholder"`
	ShowDepartment              bool   `json:"showDepartment"`
	ShowBranch                  bool   `json:"showBranch"`
	ShowSubsidiaries            bool   `json:"showSubsidiaries"`
	ShowAllBranches             bool   `json:"showAllBranches"`
	AuthorizeAll                bool   `json:"authorizeAll"`
	AuthorizeDepartment         bool   `json:"authorizeDepartment"`
	AuthorizeBudgetHolder       bool   `json:"authorizeBudgetHolder"`
	AuthorizeBranch             bool   `json:"authorizeBranch"`
	AuthorizeSubsidiaryBranches bool   `json:"authorizeSubsidiaryBranches"`
	OptionalFields1             struct {
		Boolean1    bool      `json:"boolean1"`
		Boolean2    bool      `json:"boolean2"`
		Boolean3    bool      `json:"boolean3"`
		Boolean4    bool      `json:"boolean4"`
		Boolean5    bool      `json:"boolean5"`
		Number1     int       `json:"number1"`
		Number2     int       `json:"number2"`
		Number3     int       `json:"number3"`
		Number4     int       `json:"number4"`
		Number5     int       `json:"number5"`
		Date1       time.Time `json:"date1"`
		Date2       time.Time `json:"date2"`
		Date3       time.Time `json:"date3"`
		Date4       time.Time `json:"date4"`
		Date5       time.Time `json:"date5"`
		Text1       string    `json:"text1"`
		Text2       string    `json:"text2"`
		Text3       string    `json:"text3"`
		Text4       string    `json:"text4"`
		Text5       string    `json:"text5"`
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
		Boolean1    bool      `json:"boolean1"`
		Boolean2    bool      `json:"boolean2"`
		Boolean3    bool      `json:"boolean3"`
		Boolean4    bool      `json:"boolean4"`
		Boolean5    bool      `json:"boolean5"`
		Number1     int       `json:"number1"`
		Number2     int       `json:"number2"`
		Number3     int       `json:"number3"`
		Number4     int       `json:"number4"`
		Number5     int       `json:"number5"`
		Date1       time.Time `json:"date1"`
		Date2       time.Time `json:"date2"`
		Date3       time.Time `json:"date3"`
		Date4       time.Time `json:"date4"`
		Date5       time.Time `json:"date5"`
		Text1       string    `json:"text1"`
		Text2       string    `json:"text2"`
		Text3       string    `json:"text3"`
		Text4       string    `json:"text4"`
		Text5       string    `json:"text5"`
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
	BudgetHolder struct {
		ID string `json:"id"`
	} `json:"budgetHolder"`
	PersonExtraFieldA struct {
		ID string `json:"id"`
	} `json:"personExtraFieldA"`
	PersonExtraFieldB struct {
		ID string `json:"id"`
	} `json:"personExtraFieldB"`
	IsManager bool `json:"isManager"`
	Manager   struct {
		ID string `json:"id"`
	} `json:"manager"`
}

type Location struct {
	ID            string `json:"id"`
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
	OptionalFields1 struct {
		Boolean1    bool      `json:"boolean1"`
		Boolean2    bool      `json:"boolean2"`
		Boolean3    bool      `json:"boolean3"`
		Boolean4    bool      `json:"boolean4"`
		Boolean5    bool      `json:"boolean5"`
		Number1     int       `json:"number1"`
		Number2     int       `json:"number2"`
		Number3     int       `json:"number3"`
		Number4     int       `json:"number4"`
		Number5     int       `json:"number5"`
		Date1       time.Time `json:"date1"`
		Date2       time.Time `json:"date2"`
		Date3       time.Time `json:"date3"`
		Date4       time.Time `json:"date4"`
		Date5       time.Time `json:"date5"`
		Text1       string    `json:"text1"`
		Text2       string    `json:"text2"`
		Text3       string    `json:"text3"`
		Text4       string    `json:"text4"`
		Text5       string    `json:"text5"`
		Memo1       string    `json:"memo1"`
		Memo2       string    `json:"memo2"`
		Memo3       string    `json:"memo3"`
		Memo4       string    `json:"memo4"`
		Memo5       string    `json:"memo5"`
		Searchlist1 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist1"`
		Searchlist2 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist2"`
		Searchlist3 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist3"`
		Searchlist4 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist4"`
		Searchlist5 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist5"`
	} `json:"optionalFields1"`
	OptionalFields2 struct {
		Boolean1    bool      `json:"boolean1"`
		Boolean2    bool      `json:"boolean2"`
		Boolean3    bool      `json:"boolean3"`
		Boolean4    bool      `json:"boolean4"`
		Boolean5    bool      `json:"boolean5"`
		Number1     int       `json:"number1"`
		Number2     int       `json:"number2"`
		Number3     int       `json:"number3"`
		Number4     int       `json:"number4"`
		Number5     int       `json:"number5"`
		Date1       time.Time `json:"date1"`
		Date2       time.Time `json:"date2"`
		Date3       time.Time `json:"date3"`
		Date4       time.Time `json:"date4"`
		Date5       time.Time `json:"date5"`
		Text1       string    `json:"text1"`
		Text2       string    `json:"text2"`
		Text3       string    `json:"text3"`
		Text4       string    `json:"text4"`
		Text5       string    `json:"text5"`
		Memo1       string    `json:"memo1"`
		Memo2       string    `json:"memo2"`
		Memo3       string    `json:"memo3"`
		Memo4       string    `json:"memo4"`
		Memo5       string    `json:"memo5"`
		Searchlist1 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist1"`
		Searchlist2 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist2"`
		Searchlist3 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist3"`
		Searchlist4 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist4"`
		Searchlist5 struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"searchlist5"`
	} `json:"optionalFields2"`
}
