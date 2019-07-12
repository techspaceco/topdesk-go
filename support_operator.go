package topdesk

import (
	"context"
	"encoding/json"
	"path"
)

type OperatorIterator struct {
	*ListIterator
}

func (i OperatorIterator) Operator() (*Operator, error) {
	response := &Ref{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetOperator(i.ctx, response.ID)
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

func (o Operator) Ref() *Ref {
	return &Ref{ID: o.ID}
}

type ListOperatorsRequest struct{}

func (rc RestClient) ListOperators(ctx context.Context, request *ListOperatorsRequest) (*OperatorIterator, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "operators")

	it, err := rc.list(ctx, &uri)
	return &OperatorIterator{it}, err
}

func (rc RestClient) GetOperator(ctx context.Context, id string) (*Operator, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "operators", "id", id)

	response := &Operator{}
	err := rc.get(ctx, &uri, response)
	return response, err
}
