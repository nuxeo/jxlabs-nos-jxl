package root

import (
	"fmt"
	"os"

	"github.com/jenkins-x-labs/helmboot/pkg/cmd/verify"
	"github.com/jenkins-x-labs/jxl/cmd/root/edit"
	"github.com/jenkins-x-labs/jxl/cmd/root/get"
	"github.com/jenkins-x-labs/jxl/cmd/root/start"
	"github.com/jenkins-x-labs/jxl/cmd/root/step"
	"github.com/jenkins-x-labs/jxl/pkg/cmd/upgrade"
	"github.com/jenkins-x-labs/jxl/pkg/cmd/version"
	jxcmd "github.com/jenkins-x/jx/pkg/cmd"
	"github.com/jenkins-x/jx/pkg/cmd/clients"
	"github.com/jenkins-x/jx/pkg/cmd/namespace"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/preview"
	"github.com/jenkins-x/jx/pkg/cmd/promote"
	"github.com/jenkins-x/jx/pkg/cmd/rsh"
	"github.com/jenkins-x/jx/pkg/cmd/stop"
	"github.com/jenkins-x/jx/pkg/cmd/ui"
	"github.com/jenkins-x/jx/pkg/helm"

	"github.com/jenkins-x-labs/helmboot/pkg/cmd"
	"github.com/jenkins-x-labs/helmboot/pkg/common"
	jwcommon "github.com/jenkins-x-labs/jwizard/pkg/cmd/common"
	"github.com/jenkins-x-labs/jwizard/pkg/cmd/create"
	tp "github.com/jenkins-x-labs/trigger-pipeline/pkg/cmd"
	tpcommon "github.com/jenkins-x-labs/trigger-pipeline/pkg/common"
	helmfilePatch "github.com/nuxeo/jxlabs-nos-yaml-patch/pkg/cmd"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "jxl",
		Short: "Experimental labs commands for working with Jenkins X",
		Long: `These are experimental commands that are likely to change and could be removed

Commands can mature to alpha and beta phases before being accepted into the core jx command
Complete documentation is available at https://jenkins-x.io
`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	f := clients.NewFactory()
	commonOptions := opts.NewCommonOptionsWithTerm(f, os.Stdin, os.Stdout, os.Stderr)

	// lets default to helm 3
	commonOptions.SetHelm(helm.NewHelmCLI("helm", helm.V3, "", false))

	common.TopLevelCommand = "boot"
	common.BinaryName = "jxl boot"
	rootCmd.AddCommand(cmd.Main())
	rootCmd.AddCommand(verify.NewCmdVerify())

	tpcommon.TopLevelCommand = "jenkins"
	tpcommon.BinaryName = "jxl jenkins"

	jenkinsCmd := tp.NewCmd()
	jenkinsCmd.Use = "jenkins"
	rootCmd.AddCommand(jenkinsCmd)

	helmfilePatchCmd := helmfilePatch.NewCmd(commonOptions)
	rootCmd.AddCommand(helmfilePatchCmd)

	jwcommon.TopLevelCommand = "project"
	jwcommon.BinaryName = "jxl project"
	rootCmd.AddCommand(create.NewCmdCreateProject(commonOptions))

	rootCmd.AddCommand(upgrade.NewCmdUpgrade(commonOptions))
	rootCmd.AddCommand(version.NewCmdVersion(commonOptions))

	// lets import the jx commands
	rootCmd.AddCommand(jxcmd.NewCmdCompletion(commonOptions))
	rootCmd.AddCommand(jxcmd.NewCmdContext(commonOptions))
	rootCmd.AddCommand(jxcmd.NewCmdEnvironment(commonOptions))
	rootCmd.AddCommand(jxcmd.NewCmdLogs(commonOptions))
	rootCmd.AddCommand(jxcmd.NewCmdPrompt(commonOptions))
	rootCmd.AddCommand(jxcmd.NewCmdRepo(commonOptions))
	rootCmd.AddCommand(jxcmd.NewCmdShell(commonOptions))

	rootCmd.AddCommand(edit.NewCmdEdit(commonOptions))
	rootCmd.AddCommand(get.NewCmdGet(commonOptions))
	rootCmd.AddCommand(namespace.NewCmdNamespace(commonOptions))
	rootCmd.AddCommand(preview.NewCmdPreview(commonOptions))
	rootCmd.AddCommand(promote.NewCmdPromote(commonOptions))
	rootCmd.AddCommand(rsh.NewCmdRsh(commonOptions))
	rootCmd.AddCommand(start.NewCmdStart(commonOptions))
	rootCmd.AddCommand(step.NewCmdStep(commonOptions))
	rootCmd.AddCommand(stop.NewCmdStop(commonOptions))
	rootCmd.AddCommand(ui.NewCmdUI(commonOptions))

}
