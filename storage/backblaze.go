package storage

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	MaxFileSize = 50 * 1024 * 1024 // 50 MB
	B2AuthURL   = "https://api.backblazeb2.com/b2api/v2/b2_authorize_account"
)

type B2Client struct {
	KeyID          string
	ApplicationKey string
	BucketID       string
	BucketName     string
}

type b2AuthResponse struct {
	AccountID          string `json:"accountId"`
	AuthorizationToken string `json:"authorizationToken"`
	APIURL             string `json:"apiUrl"`
	DownloadURL        string `json:"downloadUrl"`
}

type b2UploadURLResponse struct {
	BucketID           string `json:"bucketId"`
	UploadURL          string `json:"uploadUrl"`
	AuthorizationToken string `json:"authorizationToken"`
}

type b2UploadFileResponse struct {
	FileID   string `json:"fileId"`
	FileName string `json:"fileName"`
}

// Authorize authenticates with B2 and returns auth info
func (b *B2Client) authorize() (*b2AuthResponse, error) {
	req, _ := http.NewRequest("GET", B2AuthURL, nil)
	req.SetBasicAuth(b.KeyID, b.ApplicationKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var auth b2AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&auth); err != nil {
		return nil, err
	}
	return &auth, nil
}

func (b *B2Client) getUploadURL(auth *b2AuthResponse) (*b2UploadURLResponse, error) {
	body, _ := json.Marshal(map[string]string{"bucketId": b.BucketID})
	req, _ := http.NewRequest("POST", auth.APIURL+"/b2api/v2/b2_get_upload_url", bytes.NewReader(body))
	req.Header.Set("Authorization", auth.AuthorizationToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var uploadURL b2UploadURLResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadURL); err != nil {
		return nil, err
	}
	return &uploadURL, nil
}

// Upload uploads a file to Backblaze B2 and returns the public URL
func (b *B2Client) Upload(fileData []byte, fileName, contentType, folder string) (string, string, error) {
	if len(fileData) > MaxFileSize {
		return "", "", fmt.Errorf("file too large (max 50MB)")
	}

	auth, err := b.authorize()
	if err != nil {
		return "", "", fmt.Errorf("auth failed: %w", err)
	}

	uploadURL, err := b.getUploadURL(auth)
	if err != nil {
		return "", "", fmt.Errorf("get upload URL failed: %w", err)
	}

	remoteName := folder + "/" + fileName
	sha1Hash := fmt.Sprintf("%x", sha1.Sum(fileData))

	req, _ := http.NewRequest("POST", uploadURL.UploadURL, bytes.NewReader(fileData))
	req.Header.Set("Authorization", uploadURL.AuthorizationToken)
	req.Header.Set("X-Bz-File-Name", remoteName)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(fileData)))
	req.Header.Set("X-Bz-Content-Sha1", sha1Hash)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var result b2UploadFileResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", err
	}

	publicURL := fmt.Sprintf("%s/file/%s/%s", auth.DownloadURL, b.BucketName, remoteName)
	return publicURL, result.FileID, nil
}

// Delete removes a file from B2 by fileID and fileName
func (b *B2Client) Delete(fileID, fileName string) error {
	auth, err := b.authorize()
	if err != nil {
		return err
	}

	body, _ := json.Marshal(map[string]string{"fileId": fileID, "fileName": fileName})
	req, _ := http.NewRequest("POST", auth.APIURL+"/b2api/v2/b2_delete_file_version", bytes.NewReader(body))
	req.Header.Set("Authorization", auth.AuthorizationToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("delete failed: %s", string(bodyBytes))
	}
	return nil
}

// GetFileNameFromURL extracts the remote file path from a B2 public URL
func GetFileNameFromURL(url, bucketName string) string {
	prefix := "/file/" + bucketName + "/"
	idx := strings.Index(url, prefix)
	if idx == -1 {
		return ""
	}
	return url[idx+len(prefix):]
}
