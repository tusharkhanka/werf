package get

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/werf/werf/cmd/werf/common"
	"github.com/werf/werf/pkg/deploy"
	"github.com/werf/werf/pkg/deploy/helm"
	"github.com/werf/werf/pkg/true_git"
	"github.com/werf/werf/pkg/werf"
)

var cmdData struct {
	helm.GetOptions
}

var commonCmdData common.CmdData

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "get RELEASE_NAME",
		Short:                 "Get release templates and values by release name",
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			common.CmdEnvAnno: common.EnvsDescription(),
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := common.ProcessLogOptions(&commonCmdData); err != nil {
				common.PrintHelp(cmd)
				return err
			}

			if err := common.ValidateArgumentCount(1, args, cmd); err != nil {
				return err
			}
			return runGet(args[0])
		},
	}

	common.SetupTmpDir(&commonCmdData, cmd)
	common.SetupHomeDir(&commonCmdData, cmd)

	common.SetupKubeConfig(&commonCmdData, cmd)
	common.SetupKubeConfigBase64(&commonCmdData, cmd)
	common.SetupKubeContext(&commonCmdData, cmd)
	common.SetupHelmReleaseStorageNamespace(&commonCmdData, cmd)
	common.SetupHelmReleaseStorageType(&commonCmdData, cmd)

	common.SetupLogOptions(&commonCmdData, cmd)

	cmd.Flags().Int32Var(&cmdData.Revision, "revision", 0, "Get the named release by revision (use latest revision by default)")
	cmd.Flags().StringVar(&cmdData.Template, "template", "", "Go template for formatting the output, eg: {{.Release.Name}}")

	return cmd
}

func runGet(releaseName string) error {
	ctx := common.BackgroundContext()

	if err := werf.Init(*commonCmdData.TmpDir, *commonCmdData.HomeDir); err != nil {
		return fmt.Errorf("initialization error: %s", err)
	}

	if err := true_git.Init(true_git.Options{LiveGitOutput: *commonCmdData.LogVerbose || *commonCmdData.LogDebug}); err != nil {
		return err
	}

	helmReleaseStorageType, err := common.GetHelmReleaseStorageType(*commonCmdData.HelmReleaseStorageType)
	if err != nil {
		return err
	}

	deployInitOptions := deploy.InitOptions{
		HelmInitOptions: helm.InitOptions{
			KubeConfig:                  *commonCmdData.KubeConfig,
			KubeConfigBase64:            *commonCmdData.KubeConfigBase64,
			KubeContext:                 *commonCmdData.KubeContext,
			HelmReleaseStorageNamespace: *commonCmdData.HelmReleaseStorageNamespace,
			HelmReleaseStorageType:      helmReleaseStorageType,
			ReleasesMaxHistory:          0,
		},
	}
	if err := deploy.Init(ctx, deployInitOptions); err != nil {
		return err
	}

	if err := helm.Get(common.BackgroundContext(), os.Stdout, releaseName, cmdData.GetOptions); err != nil {
		return err
	}

	return nil
}
