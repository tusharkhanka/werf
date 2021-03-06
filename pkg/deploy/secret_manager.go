package deploy

import (
	"context"
	"os"
	"path/filepath"

	"github.com/werf/logboek"

	"github.com/werf/werf/pkg/deploy/secret"
	"github.com/werf/werf/pkg/deploy/werf_chart"
)

func GetSafeSecretManager(ctx context.Context, projectDir, helmChartDir string, secretValues []string, ignoreSecretKey bool) (secret.Manager, error) {
	isSecretsExists := false
	if _, err := os.Stat(filepath.Join(helmChartDir, werf_chart.SecretDirName)); !os.IsNotExist(err) {
		isSecretsExists = true
	}
	if _, err := os.Stat(filepath.Join(helmChartDir, werf_chart.DefaultSecretValuesFileName)); !os.IsNotExist(err) {
		isSecretsExists = true
	}
	if len(secretValues) > 0 {
		isSecretsExists = true
	}

	if isSecretsExists {
		if ignoreSecretKey {
			logboek.Context(ctx).Default().LogLnDetails("Secrets decryption disabled")
			return secret.NewSafeManager()
		}

		key, err := secret.GetSecretKey(projectDir)
		if err != nil {
			return nil, err
		}

		return secret.NewManager(key)
	} else {
		return secret.NewSafeManager()
	}
}
