package github

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	FullName    string    `json:"full_name"`
	Private     bool      `json:"private"`
	Owner       repoOwner `json:"owner"`
	HtmlUrl     string    `json:"html_url"`
}

type repoOwner struct {
	Login     string `json:"login"`
	Type      string `json:"type"`
	HtmlUrl   string `json:"html_url"`
	AvatarUrl string `json:"avatar_url"`
}
