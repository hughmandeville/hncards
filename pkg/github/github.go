// This package contains functions for uploading and downloading files from GitHub.
// It uses the GitHub API which requires a Personal Access Token with write permission.
//   https://docs.github.com/en/rest/reference/repos#create-or-update-file-contents

package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Controller for uploading and downloading files to and from GitHub.
type GitHubController struct {
	AccessToken string // GitHub Personal Access Token
	Branch      string // Github branch (e.g. main)
	Owner       string // owner of the GitHub repository (e.g. hughmandeville)
	Repo        string // GitHub repository (e.g. hnui)
}

// Create an new GitHub controller.
func NewGitHubController(accessToken string, branch string, owner string, repo string) *GitHubController {
	return &GitHubController{
		AccessToken: accessToken,
		Branch:      branch,
		Owner:       owner,
		Repo:        repo,
	}
}

// GitHub Get Contents API response.
type GHGetContentsResponse struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	SHA         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

// Get a file from GitHub. Returns its contents and its SHA. The SHA is needed if you update the file later on GitHub.
func (gh *GitHubController) GetFile(path string) (fileData []byte, fileSHA string, err error) {

	// Setup API URL.
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s", gh.Owner, gh.Repo, path, gh.Branch)

	// Create the request.
	client := &http.Client{Timeout: time.Second * 120}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Token "+gh.AccessToken)

	// Make request.
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Unmarshal response body and pull out file contents and SHA.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var data GHGetContentsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}

	// Make sure we get back the correct type and encoding.
	if data.Type != "file" {
		err = fmt.Errorf("Expected type 'file', got '%s'", data.Type)
		return
	}
	if data.Encoding != "base64" {
		err = fmt.Errorf("Expected encoding 'base64', got '%s'", data.Encoding)
		return
	}

	// Decode the file contents.
	fileData, err = base64.StdEncoding.DecodeString(data.Content)
	if err != nil {
		return
	}
	fileSHA = data.SHA
	return
}

// Arugments for the GitHub API contents PUT call.
type GHContentsPutRequest struct {
	Content string `json:"content"`
	Branch  string `json:"branch,omitempty"`
	Message string `json:"message"`
	SHA     string `json:"sha,omitempty"`
}

// Updates a file in GitHub. The SHA of the existing file is needed to replace it.
func (gh *GitHubController) PutFile(path string, data []byte, sha string) (err error) {

	// Setup API URL.
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s", gh.Owner, gh.Repo, path)

	// Base64 encode the file data.
	str := base64.StdEncoding.EncodeToString(data)

	// Generate the payload request.
	payloadData := GHContentsPutRequest{
		Branch:  gh.Branch,
		Content: str,
		Message: "updated file",
		SHA:     sha,
	}
	payload, err := json.Marshal(payloadData)
	if err != nil {
		return
	}

	// Create the request.
	client := &http.Client{Timeout: time.Second * 120}
	req, err := http.NewRequest("PUT", apiURL, bytes.NewReader(payload))
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "Token "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	// Make request.
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Unmarshal response body to check for errors.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		// TO DO: pull error message out of response body JSON.
		err = fmt.Errorf("GitHub API returned %s: %s", resp.Status, body)
		return
	}

	return nil
}
