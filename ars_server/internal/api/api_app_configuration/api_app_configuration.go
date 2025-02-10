package api_app_configuration

import (
	"ars_server/internal/root"
	"context"
)

type ResponseListAppConfigurations struct {
	Object  string        `json:"object"`
	Url     string        `json:"url"`
	HasMore bool          `json:"has_more"`
	Data    []root.Config `json:"data"`
}

func NewResponseListAppConfigurations(data []root.Config, hasMore bool) ResponseListAppConfigurations {
	return ResponseListAppConfigurations{
		Object:  "list",
		Url:     "/v1/app_configurations",
		HasMore: hasMore,
		Data:    data,
	}
}

func ListAppConfigurations(ctx context.Context, configService root.ConfigService) ([]root.Config, bool, error) {
	configs, err := configService.GetAppConfigurationsByScope(ctx, root.ConfigScopeGlobal)
	if err != nil {
		return nil, false, err
	}
	return configs, false, nil
}
