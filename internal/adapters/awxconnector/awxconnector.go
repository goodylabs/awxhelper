package awxconnector

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/goodylabs/awxhelper/internal/domain/entities"
	"github.com/goodylabs/awxhelper/internal/ports"
	"github.com/goodylabs/awxhelper/pkg/config"
)

type awxconnector struct {
	httpconnector ports.HttpConnector
	httpCfg       ports.HttpConnOpts
}

type listJobTemplatesResponse struct {
	Next    string `json:"next"`
	Results []struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		SummaryFilds struct {
			Labels struct {
				Results []string `json:"results"`
			} `json:"labels"`
		} `json:"summary_fields"`
	} `json:"results"`
}

func NewAwxConnector(httpconnector ports.HttpConnector) ports.AwxConnector {
	return &awxconnector{
		httpconnector: httpconnector,
	}
}

func (a *awxconnector) verifyConnection() error {
	respBody, statusCode, httpErr := a.httpconnector.DoGet(a.httpCfg, "/api/v2/me/")
	if httpErr == nil && statusCode == 200 {
		return nil
	}

	var data any
	if err := a.unmarshalResponseBody(respBody, &data); err != nil {
		return fmt.Errorf("statusCode: %d, failed to unmarshal response: %w, body %v", statusCode, err, respBody)
	}

	pretty, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error: %s", err)
		fmt.Printf("Response body: \n %s", respBody)
		return fmt.Errorf("statusCode: %d, failed to format JSON: %w", statusCode, err)
	}

	fmt.Printf("statusCode: %d\n", statusCode)
	fmt.Println(string(pretty))

	return httpErr
}

func (a *awxconnector) ConfigureConnection(cfg *ports.AwxConfig) error {
	a.httpCfg.BaseURL = cfg.URL
	a.httpCfg.Username = cfg.Username
	a.httpCfg.Password = cfg.Password

	return a.verifyConnection()
}

func (a *awxconnector) ListJobTemplates(prefix string) ([]ports.PrompterItem, error) {
	var jobTemplates []ports.PrompterItem
	nextURL := fmt.Sprintf("/api/v2/job_templates/?name__icontains=%s&page_size=100", prefix)

	for nextURL != "" {
		respBody, statusCode, err := a.httpconnector.DoGet(a.httpCfg, nextURL)
		if err != nil {
			return nil, err
		}
		if statusCode != 200 {
			return nil, fmt.Errorf("failed to list job templates, status %d", statusCode)
		}

		var response listJobTemplatesResponse
		if err := a.unmarshalResponseBody(respBody, &response); err != nil {
			return nil, err
		}

		for _, r := range response.Results {
			jobTemplates = append(jobTemplates, ports.PrompterItem{
				Label: r.Name,
				Value: strconv.Itoa(r.ID),
			})
		}

		nextURL = response.Next
	}

	return jobTemplates, nil
}
func (a *awxconnector) LaunchJob(templateId string, extraVars *entities.ExtraVars) (int, error) {
	templateIdInt, err := strconv.Atoi(templateId)
	if err != nil {
		return 0, fmt.Errorf("invalid template id: %w", err)
	}

	url := fmt.Sprintf("/api/v2/job_templates/%d/launch/", templateIdInt)

	extraVarsValue := map[string]any{}
	if extraVars != nil {
		for k, v := range *extraVars {
			extraVarsValue[k] = v
		}
	}

	launchBody := map[string]any{
		"inventory":  config.INVENTORY_ID,
		"extra_vars": extraVarsValue,
	}

	respBody, statusCode, err := a.httpconnector.DoPost(a.httpCfg, url, launchBody)
	if err != nil {
		return 0, err
	}
	if statusCode != 201 {
		return 0, fmt.Errorf("failed to launch job, status %d: %s", statusCode, string(respBody))
	}

	var response struct {
		ID int `json:"id"`
	}
	err = a.unmarshalResponseBody(respBody, &response)
	return response.ID, err
}

func (a *awxconnector) unmarshalResponseBody(respBody []byte, out any) error {
	if err := json.Unmarshal(respBody, &out); err != nil {
		fmt.Printf("Error: %s", err)
		fmt.Printf("Response body: \n %s", respBody)
		return err
	}
	return nil
}
