package cmd

import (
	"fmt"
	"github.com/saeedafshari8/flixinit/project/gitlab"
	"github.com/saeedafshari8/flixinit/util"
	"github.com/spf13/cobra"
	"log"
)

var (
	gitlabConfig gitlab.GitlabConfig

	cmdGitLab = &cobra.Command{
		Use:   "gitlab",
		Short: "gitlab command generates a new project in the remote repository.",
		Long:  `gitlab command generates a new project in the remote repository.`,
		Run: func(cmd *cobra.Command, args []string) {
			initGitlabConfig(cmd)

			project, err := gitlab.GitCreateProject(gitlabConfig)
			util.LogAndExit(err, util.NetworkError)

			if project.Message.Name == nil {
				log.Println(fmt.Sprintf("Name = %s", project.Name))
				log.Println(fmt.Sprintf("Full name = %s", project.NameWithNamespace))
				log.Println(fmt.Sprintf("SSH_URL = %s", project.RepoSshUrl))
				log.Println(fmt.Sprintf("Http_URL = %s", project.RepoHttpUrl))
			} else {
				for _, msg := range project.Message.Name {
					log.Println(msg)
				}
			}
		},
	}

	namespacesCommand = &cobra.Command{
		Use:   "namespaces",
		Short: "gitlab command get list of existing namespaces.",
		Long:  `gitlab command get list of existing namespaces.`,
		Run: func(cmd *cobra.Command, args []string) {
			token := util.GetValue(cmd, Token)
			util.ValidateRequired(token, Token)

			namespaces, err := gitlab.GitGetNamespaces(token)
			util.LogAndExit(err, util.NetworkError)

			for _, namespace := range namespaces {
				log.Println(fmt.Sprintf("Id = %d, Name=%s, FullPath=%s", namespace.Id, namespace.Name, namespace.FullPath))
			}
		},
	}
)

const (
	Token                                     = "token"
	Name                                      = "name"
	Path                                      = "path"
	NamespaceID                               = "namespace_id"
	Visibility                                = "visibility"
	OnlyAllowMergeIfPipelineSucceeds          = "only_allow_merge_if_pipeline_succeeds"
	OnlyAllowMergeIfAllDiscussionsAreResolved = "only_allow_merge_if_all_discussions_are_resolved"
	AutoCancelPendingPipelines                = "auto_cancel_pending_pipelines"
	ApprovalsBeforeMerge                      = "approvals_before_merge"
	InitializeWithReadme                      = "initialize_with_readme"
)

func init() {
	cmdGitLab.Flags().StringP(AutoCancelPendingPipelines, "", "enabled", "Auto-cancel pending pipelines (Note: this is not a boolean, but enabled/disabled")
	cmdGitLab.Flags().BoolP(OnlyAllowMergeIfPipelineSucceeds, "", true, "Auto-cancel pending pipelines (Note: this is not a boolean, but enabled/disabled")
	cmdGitLab.Flags().BoolP(OnlyAllowMergeIfAllDiscussionsAreResolved, "", true, "Set whether merge requests can only be merged when all the discussions are resolved")
	cmdGitLab.Flags().Int32P(ApprovalsBeforeMerge, "", 1, "How many approvers should approve merge requests by default")
	cmdGitLab.Flags().BoolP(InitializeWithReadme, "r", true, "Initialise by README.md")
	cmdGitLab.Flags().StringP(Name, "n", "", "The name of the new project. Equals path if not provided.")
	cmdGitLab.Flags().Int32P(NamespaceID, "", 0, "Namespace for the new project (defaults to the current user’s namespace)")
	cmdGitLab.Flags().StringP(Path, "p", "", "Repository name for new project. Generated based on name if not provided (generated lowercased with dashes).")
	cmdGitLab.Flags().StringP(Token, "", "", "Gitlab token.")
	cmdGitLab.Flags().StringP(Visibility, "v", "private", "private|internal|public")

	namespacesCommand.Flags().StringP(Token, "", "", "Gitlab token.")

	cmdGitLab.AddCommand(namespacesCommand)
}

func initGitlabConfig(cmd *cobra.Command) {
	//Mandatory flags
	gitlabConfig.Token = util.GetValue(cmd, Token)
	util.ValidateRequired(gitlabConfig.Token, Token)
	gitlabConfig.Name = util.GetValue(cmd, Name)
	util.ValidateRequired(gitlabConfig.Name, Name)
	gitlabConfig.NamespaceID = util.GetValueInt32(cmd, NamespaceID)
	if gitlabConfig.NamespaceID == 0 {
		util.LogMessageAndExit(fmt.Sprintf("%s is mandatory!\n", NamespaceID))
	}

	//Optional flags
	gitlabConfig.Path = util.GetValue(cmd, Path)
	gitlabConfig.Visibility = util.GetValue(cmd, Visibility)
	gitlabConfig.AutoCancelPendingPipelines = util.GetValue(cmd, AutoCancelPendingPipelines)
	gitlabConfig.OnlyAllowMergeIfPipelineSucceeds = util.GetValueBool(cmd, OnlyAllowMergeIfPipelineSucceeds)
	gitlabConfig.OnlyAllowMergeIfAllDiscussionsAreResolved = util.GetValueBool(cmd, OnlyAllowMergeIfAllDiscussionsAreResolved)
	gitlabConfig.ApprovalsBeforeMerge = util.GetValueInt32(cmd, ApprovalsBeforeMerge)
	gitlabConfig.InitializeWithReadme = util.GetValueBool(cmd, InitializeWithReadme)
}
