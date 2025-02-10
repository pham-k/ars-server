package root

import "context"

type ConfigScope string
type ConfigName string

const (
	ConfigScopeGlobal                   ConfigScope = "global"
	ConfigNameFeatureCashflowCalculator ConfigName  = "feature_cash_flow_calculator"
)

// Config represent the Config object of the system
type Config struct {
	Pid       string `json:"pid"`
	Object    string `json:"object"`
	Scope     string `json:"scope"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	TextValue string `json:"value"`
}

// ConfigService represents a service for managing config.
type ConfigService interface {
	GetAppConfigurationsByScope(ctx context.Context, scope ConfigScope) ([]Config, error)
}
