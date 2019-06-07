package topdesk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type IncidentRef struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Number string `json:"number"`
}

type Incident struct {
	ID          string `json:"id"`
	Status      string `json:"status"`
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
	OnHoldDuration  int         `json:"onHoldDuration"`
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
	TimeSpent                      int     `json:"timeSpent"`
	TimeSpentFirstLine             int     `json:"timeSpentFirstLine"`
	TimeSpentSecondLineAndPartials int     `json:"timeSpentSecondLineAndPartials"`
	Costs                          float64 `json:"costs"`
	EscalationStatus               string  `json:"escalationStatus"`
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
	ExpectedTimeSpent int         `json:"expectedTimeSpent"`
	MainIncident      interface{} `json:"mainIncident"`
	PartialIncidents  []struct {
		Link string `json:"link"`
	} `json:"partialIncidents"`
	OptionalFields1 struct {
		Boolean1    bool   `json:"boolean1"`
		Boolean2    bool   `json:"boolean2"`
		Boolean3    bool   `json:"boolean3"`
		Boolean4    bool   `json:"boolean4"`
		Boolean5    bool   `json:"boolean5"`
		Number1     int    `json:"number1"`
		Number2     int    `json:"number2"`
		Number3     int    `json:"number3"`
		Number4     int    `json:"number4"`
		Number5     int    `json:"number5"`
		Date1       string `json:"date1"`
		Date2       string `json:"date2"`
		Date3       string `json:"date3"`
		Date4       string `json:"date4"`
		Date5       string `json:"date5"`
		Text1       string `json:"text1"`
		Text2       string `json:"text2"`
		Text3       string `json:"text3"`
		Text4       string `json:"text4"`
		Text5       string `json:"text5"`
		Memo1       string `json:"memo1"`
		Memo2       string `json:"memo2"`
		Memo3       string `json:"memo3"`
		Memo4       string `json:"memo4"`
		Memo5       string `json:"memo5"`
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
		Boolean1    bool   `json:"boolean1"`
		Boolean2    bool   `json:"boolean2"`
		Boolean3    bool   `json:"boolean3"`
		Boolean4    bool   `json:"boolean4"`
		Boolean5    bool   `json:"boolean5"`
		Number1     int    `json:"number1"`
		Number2     int    `json:"number2"`
		Number3     int    `json:"number3"`
		Number4     int    `json:"number4"`
		Number5     int    `json:"number5"`
		Date1       string `json:"date1"`
		Date2       string `json:"date2"`
		Date3       string `json:"date3"`
		Date4       string `json:"date4"`
		Date5       string `json:"date5"`
		Text1       string `json:"text1"`
		Text2       string `json:"text2"`
		Text3       string `json:"text3"`
		Text4       string `json:"text4"`
		Text5       string `json:"text5"`
		Memo1       string `json:"memo1"`
		Memo2       string `json:"memo2"`
		Memo3       string `json:"memo3"`
		Memo4       string `json:"memo4"`
		Memo5       string `json:"memo5"`
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
	ExternalLinks []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
		Date string `json:"date"`
	} `json:"externalLinks"`
}

type IncidentIterator struct {
	*ListIterator
}

func (rc *RestClient) ListIncidents(ctx context.Context) (*IncidentIterator, error) {
	it, err := rc.list(ctx, "incidents")
	return &IncidentIterator{it}, err
}

func (i *IncidentIterator) Incident() (*Incident, error) {
	response := &IncidentRef{}
	if err := i.decode(&response); err != nil {
		return nil, err // Wrap this bad boy.
	}
	return i.client.GetIncident(i.ctx, response.ID)
}

func (rc *RestClient) GetIncident(ctx context.Context, id string) (*Incident, error) {
	response := &Incident{}
	err := rc.get(ctx, "incidents", id, response)
	return response, err
}

func (rc *RestClient) GetIncidentNumber(ctx context.Context, number string) (*Incident, error) {
	response := &Incident{}

	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, "incidents")
	if len(number) > 0 {
		uri.Path = path.Join(uri.Path, "number", url.QueryEscape(number))
	}

	status, err := rc.do(ctx, http.MethodGet, uri.String(), nil, response)
	if err != nil {
		return nil, err
	}

	switch status {
	case http.StatusOK:
		return response, err
	default:
		return nil, fmt.Errorf("%s delete %s", http.StatusText(status), uri.String())
	}
}