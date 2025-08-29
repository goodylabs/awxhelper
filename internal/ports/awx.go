package ports

type AwxConnector interface {
	ConfigureConnection(cfg *AwxConfig) error
	ListJobTemplates(prefix string) ([]PrompterItem, error)
	LaunchJob(templateId string, params map[string]any) (int, error)
	JobProgress(jobId int) ([]Event, error)
}

type AwxConfig struct {
	URL      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Event struct {
	Event     string `json:"event"`
	Task      string `json:"task,omitempty"`
	Changed   bool   `json:"changed,omitempty"`
	Failed    bool   `json:"failed,omitempty"`
	Created   string `json:"created,omitempty"`
	EventData struct {
		Res struct {
			Msg string `json:"msg,omitempty"`
		} `json:"res"`
	} `json:"event_data"`
	SummaryFields struct {
		Job struct {
			Status string `json:"status"`
		} `json:"job"`
	} `json:"summary_fields"`
}
