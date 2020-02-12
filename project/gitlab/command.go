package gitlab

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/saeedafshari8/flixinit/util"
	"net/http"
)

type GitlabConfig struct {
	Token                                     string `json:-`
	Name                                      string `json:"name"`
	Path                                      string `json:"path,omitempty"`
	NamespaceID                               int32  `json:"namespace_id"`
	Visibility                                string `json:"visibility,omitempty"`
	OnlyAllowMergeIfPipelineSucceeds          bool   `json:"only_allow_merge_if_pipeline_succeeds,omitempty"`
	OnlyAllowMergeIfAllDiscussionsAreResolved bool   `json:"only_allow_merge_if_all_discussions_are_resolved,omitempty"`
	AutoCancelPendingPipelines                string `json:"auto_cancel_pending_pipelines,omitempty"`
	ApprovalsBeforeMerge                      int32  `json:"approvals_before_merge,omitempty"`
	InitializeWithReadme                      bool   `json:"initialize_with_readme,omitempty"`
}

type GitlabProject struct {
	Id                int32         `json:"id"`
	Name              string        `json:"name"`
	NameWithNamespace string        `json:"name_with_namespace"`
	Path              string        `json:"path"`
	PathWithNamespace string        `json:"path_with_namespace"`
	Description       string        `json:"description"`
	DefaultBranch     string        `json:"default_branch"`
	WebUrl            string        `json:"web_url"`
	RepoSshUrl        string        `json:"ssh_url_to_repo"`
	RepoHttpUrl       string        `json:"http_url_to_repo"`
	StarCount         int32         `json:"star_count"`
	Message           GitLabMessage `json:"message"`
}

type GitLabMessage struct {
	Name []string `json:"name"`
}

type GitlabNamespace struct {
	Id                   int32  `json:"id"`
	Name                 string `json:"name"`
	Path                 string `json:"path"`
	Kind                 string `json:"kind"`
	FullPath             string `json:"full_path"`
	WebUrl               string `json:"web_url"`
	BillableMembersCount int32  `json:"billable_members_count"`
}

func GitCreateProject(gitlabConfig GitlabConfig) (*GitlabProject, error) {
	marshaledBody, err := json.Marshal(gitlabConfig)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://git.flix.tech/api/v4/projects?private_token=%s", gitlabConfig.Token)
	req, err := http.NewRequest("POST", url, bytes.NewReader(marshaledBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	ch := make(chan util.ChannelResponse)
	defer close(ch)
	go util.MakeHttpRequest(req, ch)
	channelResponse := <-ch
	if channelResponse.Success {
		data := &GitlabProject{}
		json.Unmarshal(channelResponse.Data, &data)
		return data, nil
	}
	return nil, channelResponse.Error
}

func GitGetNamespaces(token string) ([]GitlabNamespace, error) {
	req, err := http.NewRequest("GET", "https://git.flix.tech/api/v4/namespaces", nil)
	req.Header.Set("PRIVATE-TOKEN", token)

	if err != nil {
		return nil, err
	}

	ch := make(chan util.ChannelResponse)
	defer close(ch)

	go util.MakeHttpRequest(req, ch)

	channelResponse := <-ch
	if channelResponse.Success {
		data := make([]GitlabNamespace, 0)
		json.Unmarshal(channelResponse.Data, &data)
		return data, nil
	}
	return nil, err
}
