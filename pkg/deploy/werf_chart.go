package deploy

import (
	"context"

	"github.com/werf/logboek"

	"github.com/werf/werf/pkg/deploy/secret"
	"github.com/werf/werf/pkg/deploy/werf_chart"
)

func PrepareWerfChart(ctx context.Context, projectName, helmChartDir, env string, m secret.Manager, secretValues []string, serviceValues map[string]interface{}) (*werf_chart.WerfChart, error) {
	werfChart, err := werf_chart.InitWerfChart(ctx, projectName, helmChartDir, env, m)
	if err != nil {
		return nil, err
	}

	for _, path := range secretValues {
		if err = werfChart.SetSecretValuesFile(path, m); err != nil {
			return nil, err
		}
	}

	if serviceValues != nil {
		if err = werfChart.SetServiceValues(serviceValues); err != nil {
			return nil, err
		}
	}

	logboek.Context(ctx).Debug().LogF("werf chart: %#v\n", werfChart)

	return werfChart, nil
}
