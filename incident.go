package topdesk

import (
	"context"
	"encoding/json"
	"path"
)

type IncidentIterator struct {
	*ListIterator
}

func (i IncidentIterator) Incident() (*Incident, error) {
	response := &Ref{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetIncident(i.ctx, response.ID)
}

type IncidentStatus string

const (
	IncidentStatusFirstLine  IncidentStatus = "firstLine"
	IncidentStatusSecondLine IncidentStatus = "secondLine"
	IncidentStatusPartial    IncidentStatus = "partial"
)

type Incident struct {
	ID          string `json:"id"`
	Status      IncidentStatus
	Number      string `json:"number"`
	Request     string `json:"request"`
	Requests    string `json:"requests"`
	Action      string `json:"action"`
	Attachments string `json:"attachments"`
	Caller      struct {
		ID          string `json:"id"`
		DynamicName string `json:"dynamicName"`
		Branch      struct {
			ClientReferenceNumber string      `json:"clientReferenceNumber"`
			TimeZone              string      `json:"timeZone"`
			ExtraA                interface{} `json:"extraA"`
			ExtraB                interface{} `json:"extraB"`
			ID                    string      `json:"id"`
			Name                  string      `json:"name"`
		} `json:"branch"`
	} `json:"caller"`
	CallerBranch struct {
		ClientReferenceNumber string      `json:"clientReferenceNumber"`
		TimeZone              string      `json:"timeZone"`
		ExtraA                interface{} `json:"extraA"`
		ExtraB                interface{} `json:"extraB"`
		ID                    string      `json:"id"`
		Name                  string      `json:"name"`
	} `json:"callerBranch"`
	BranchExtraFieldA interface{} `json:"branchExtraFieldA"`
	BranchExtraFieldB interface{} `json:"branchExtraFieldB"`
	BriefDescription  string      `json:"briefDescription"`
	ExternalNumber    string      `json:"externalNumber"`
	Category          struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Subcategory struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"subcategory"`
	CallType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"callType"`
	EntryType struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"entryType"`
	Object struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		Make struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"make"`
		Model struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"model"`
		Branch struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"branch"`
		Location struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"location"`
		Specification string `json:"specification"`
		SerialNumber  string `json:"serialNumber"`
	} `json:"object"`
	Asset struct {
		ID string `json:"id"`
	} `json:"asset"`
	Branch struct {
		ClientReferenceNumber string      `json:"clientReferenceNumber"`
		TimeZone              string      `json:"timeZone"`
		ExtraA                interface{} `json:"extraA"`
		ExtraB                interface{} `json:"extraB"`
		ID                    string      `json:"id"`
		Name                  string      `json:"name"`
	} `json:"branch"`
	Location struct {
		ID     string `json:"id"`
		Branch struct {
			ClientReferenceNumber string      `json:"clientReferenceNumber"`
			TimeZone              string      `json:"timeZone"`
			ExtraA                interface{} `json:"extraA"`
			ExtraB                interface{} `json:"extraB"`
			ID                    string      `json:"id"`
			Name                  string      `json:"name"`
		} `json:"branch"`
		Name string `json:"name"`
		Room string `json:"room"`
	} `json:"location"`
	Impact struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"impact"`
	Urgency struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"urgency"`
	Priority struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"priority"`
	Duration struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"duration"`
	TargetDate string `json:"targetDate"`
	SLA        struct {
		ID string `json:"id"`
	} `json:"sla"`
	OnHold          bool        `json:"onHold"`
	OnHoldDate      interface{} `json:"onHoldDate"`
	OnHoldDuration  json.Number `json:"onHoldDuration"`
	FeedbackMessage interface{} `json:"feedbackMessage"`
	FeedbackRating  interface{} `json:"feedbackRating"`
	Operator        struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Name   string `json:"name"`
	} `json:"operator"`
	OperatorGroup struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"operatorGroup"`
	Supplier struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		ForFirstLine  bool   `json:"forFirstLine"`
		ForSecondLine bool   `json:"forSecondLine"`
	} `json:"supplier"`
	ProcessingStatus struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"processingStatus"`
	Completed     bool        `json:"completed"`
	CompletedDate interface{} `json:"completedDate"`
	Closed        bool        `json:"closed"`
	ClosedDate    interface{} `json:"closedDate"`
	ClosureCode   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"closureCode"`
	TimeSpent                      json.Number `json:"timeSpent"`
	TimeSpentFirstLine             json.Number `json:"timeSpentFirstLine"`
	TimeSpentSecondLineAndPartials json.Number `json:"timeSpentSecondLineAndPartials"`
	Costs                          json.Number `json:"costs"`
	EscalationStatus               string      `json:"escalationStatus"`
	EscalationReason               struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"escalationReason"`
	EscalationOperator struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"escalationOperator"`
	CallDate string `json:"callDate"`
	Creator  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"creator"`
	CreationDate string `json:"creationDate"`
	Modifier     struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"modifier"`
	ModificationDate string `json:"modificationDate"`
	MajorCall        bool   `json:"majorCall"`
	MajorCallObject  struct {
		Name          string `json:"name"`
		ID            string `json:"id"`
		Status        int    `json:"status"`
		MajorIncident bool   `json:"majorIncident"`
	} `json:"majorCallObject"`
	PublishToSsd      bool        `json:"publishToSsd"`
	Monitored         bool        `json:"monitored"`
	ExpectedTimeSpent json.Number `json:"expectedTimeSpent"`
	MainIncident      interface{} `json:"mainIncident"`
	PartialIncidents  []struct {
		Link string `json:"link"`
	} `json:"partialIncidents"`
	ExternalLinks []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Date string `json:"date"`
	} `json:"externalLinks"`
}

func (i Incident) Ref() *Ref {
	return &Ref{ID: i.ID}
}

type ListIncidentsRequest struct {
	ExternalNumber []string
}

func (rc RestClient) ListIncidents(ctx context.Context, request *ListIncidentsRequest) (*IncidentIterator, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "incidents")

	if request != nil {
		query := uri.Query()
		for _, no := range request.ExternalNumber {
			query.Add("external_number", no)
		}

		uri.RawQuery = query.Encode()
	}

	it, err := rc.list(ctx, &uri)
	return &IncidentIterator{it}, err
}

func (rc RestClient) GetIncident(ctx context.Context, id string) (*Incident, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "incidents", "id", id)

	response := &Incident{}
	err := rc.get(ctx, &uri, response)
	return response, err
}

func (rc RestClient) GetIncidentNumber(ctx context.Context, number string) (*Incident, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "incidents", "number", number)

	response := &Incident{}
	err := rc.get(ctx, &uri, response)
	return response, err
}

// CreateIncidentRequest creates an incident from the perspective of an operator.
//
// See https://developers.topdesk.com/documentation/index.html#api-Incident-CreateIncident
type CreateIncidentRequest struct {
	Caller struct {
		DynamicName string `json:"dynamicName"`
		Email       string `json:"email"`
	} `json:"caller"`
	BriefDescription string `json:"briefDescription"`
	Request          string `json:"request"`
	ExternalNumber   string `json:"externalNumber"`
	// TODO: Set the branch, operator group & operator?
}

func (rc RestClient) CreateIncident(ctx context.Context, request *CreateIncidentRequest) (*Incident, error) {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "incidents")

	response := &Ref{}
	if err := rc.create(ctx, &uri, request, response); err != nil {
		return nil, err
	}

	return rc.GetIncident(ctx, response.ID)
}
