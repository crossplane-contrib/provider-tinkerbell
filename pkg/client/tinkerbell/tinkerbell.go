package tinkerbell

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tinkerbell/tink/client"

	"github.com/tinkerbell/tink/protos/hardware"
	"github.com/tinkerbell/tink/protos/template"
	"github.com/tinkerbell/tink/protos/workflow"
)

// ClientParams specify how to connect to a Tinkerbell service
type ClientParams struct {
	GRPCAuthority  string `json:"grpc_authority"`
	CertificateURL string `json:"cert_url"`
}

type tinkClient struct {
	templateClient template.TemplateServiceClient
	workflowClient workflow.WorkflowServiceClient
	hardwareClient hardware.HardwareServiceClient
}

// Client is a Tinkerbell Client
type Client struct {
	Params ClientParams

	tinkClient *tinkClient
}

func (c *Client) Hardware() hardware.HardwareServiceClient {
	return c.tinkClient.hardwareClient
}

func (c *Client) Template() template.TemplateServiceClient {
	return c.tinkClient.templateClient
}

func (c *Client) Workflow() workflow.WorkflowServiceClient {
	return c.tinkClient.workflowClient
}

// NewClient returns a Client using ClientParams
func (p ClientParams) NewClient() (*Client, error) {
	c := &Client{Params: p}

	if err := os.Setenv("TINKERBELL_GRPC_AUTHORITY", p.GRPCAuthority); err != nil {
		return nil, fmt.Errorf("setting TINKERBELL_GRPC_AUTHORITY environment variable: %w", err)
	}

	if err := os.Setenv("TINKERBELL_CERT_URL", p.CertificateURL); err != nil {
		return nil, fmt.Errorf("setting TINKERBELL_CERT_URL environment variable: %w", err)
	}

	conn, err := client.GetConnection()
	if err != nil {
		return nil, fmt.Errorf("creating tink client: %w", err)
	}

	c.tinkClient = &tinkClient{
		templateClient: template.NewTemplateServiceClient(conn),
		workflowClient: workflow.NewWorkflowServiceClient(conn),
		hardwareClient: hardware.NewHardwareServiceClient(conn),
	}

	return c, nil
}

// ClientService creates a Client using credential bytes
func ClientService(secret []byte) (interface{}, error) {
	params := ClientParams{}
	if err := json.Unmarshal(secret, &params); err != nil {
		return nil, err
	}
	return params.NewClient()
}
