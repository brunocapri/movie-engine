package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	uri        string
	apiKey     string
	httpClient *http.Client
}

type EmbeddingResponse struct {
	Object string `json:"object"`
	Data   []struct {
		Object    string    `json:"object"`
		Index     int       `json:"index"`
		Embedding []float64 `json:"embedding"`
	} `json:"data"`
	Model string `json:"model"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens"`
		TotalTokens  int `json:"total_tokens"`
	} `json:"usage"`
}

type EmbeddingRequest struct {
	Input          string `json:"input"`
	Model          string `json:"model"`
	EncodingFormat string `json:"encoding_format"`
}

func NewClient(uri, apiKey string) *Client {
	return &Client{
		uri:        uri,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c Client) GenerateEmbeddings(input string) (EmbeddingResponse, error) {
	resource := c.uri + "/embeddings"
	er := EmbeddingRequest{
		Input:          input,
		Model:          "text-embedding-3-small",
		EncodingFormat: "float",
	}
	body, err := json.Marshal(er)

	if err != nil {
		return EmbeddingResponse{}, nil
	}

	req, err := http.NewRequest("POST", resource, bytes.NewBuffer(body))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if err != nil {
		return EmbeddingResponse{}, nil
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return EmbeddingResponse{}, nil
	}

	defer res.Body.Close()

	embeddingResponse := EmbeddingResponse{}
	if err := json.NewDecoder(res.Body).Decode(&embeddingResponse); err != nil {
		return EmbeddingResponse{}, err
	}

	return embeddingResponse, nil
}

func (c Client) EmbeddingToString(vector []float64) string {
	var sb strings.Builder
	sb.WriteString("[")

	for i, v := range vector {
		sb.WriteString(fmt.Sprintf("%f", v))
		if i < len(vector)-1 {
			sb.WriteString(", ")
		}
	}

	sb.WriteString("]")

	vectorString := sb.String()

	return vectorString
}
