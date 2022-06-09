package github

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Tags []struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"commit"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	NodeID     string `json:"node_id"`
}

func (c *Client) GetTags(owner, repo string) (Tags, error) {
	tags := new(Tags)
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/tags", nil)
	if err != nil {
		return *tags, fmt.Errorf("Client.GetTags: http.NewRequest: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resp, err := c.client.Do(req)
	if err != nil {
		return *tags, fmt.Errorf("Client.GetTags: c.client.Do: %w", err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(tags); err != nil {
		return *tags, fmt.Errorf("Client.GetTags: json.NewDecoder: %w", err)
	}
	return *tags, err
}

func (c *Client) DownloadTag(owner, repo, tag string) ([]byte, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/tarball/refs/tags/"+tag, nil)
	if err != nil {
		return nil, fmt.Errorf("Client.DownloadTag: http.NewRequest: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Client.DownloadTag: c.client.Do: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Client.DownloadTag: status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Client.DownloadTag: ioutil.ReadAll: %w", err)
	}
	return data, nil
}

func (c *Client) DownloadAndSaveTag(owner, repo, tag, path string) error {
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/tarball/refs/tags/"+tag, nil)
	if err != nil {
		return fmt.Errorf("Client.DownloadAndSaveTag: http.NewRequest: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("Client.DownloadAndSaveTag: c.client.Do: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Client.DownloadTag: status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Client.DownloadAndSaveTag: os.Create: %w", err)
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("Client.DownloadAndSaveTag: io.Copy: %w", err)
	}
	return nil
}

func (c *Client) DownloadAndUntarTag(owner, repo, tag, path string) (string, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", fmt.Errorf("Client.DownloadAndUntarTag: os.MkdirAll: %w", err)
	}
	req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/tarball/refs/tags/"+tag, nil)
	if err != nil {
		return "", fmt.Errorf("Client.DownloadAndSaveTag: http.NewRequest: %w", err)
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Client.DownloadAndSaveTag: c.client.Do: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Client.DownloadTag: status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	rootName, err := Untar(path, resp.Body)
	if err != nil {
		return "", fmt.Errorf("Client.DownloadAndSaveTag: Untar: %w", err)
	}
	return rootName, nil
}
