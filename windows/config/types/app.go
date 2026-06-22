package types

type Application struct {
	Name           string `yaml:"name" json:"name"`
	DisplayName    string `yaml:"displayName" json:"displayName"`
	Version        string `yaml:"version" json:"version"`
	Build          string `yaml:"build" json:"build"`
	Environment    string `yaml:"environment" json:"environment"`
	Company        string `yaml:"company" json:"company"`
	Product        string `yaml:"product" json:"product"`
	Description    string `yaml:"description" json:"description"`
	AgentID        string `yaml:"agentId" json:"agentId"`
	TenantID       string `yaml:"tenantId" json:"tenantId"`
	OrganizationID string `yaml:"organizationId" json:"organizationId"`
}
