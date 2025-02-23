package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "backend-trackit/models"
)

type AIService struct {
    apiKey  string
    baseURL string
}

type OpenAIRequest struct {
    Model    string    `json:"model"`
    Messages []Message `json:"messages"`
    Store    bool      `json:"store"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type OpenAIResponse struct {
    ID      string `json:"id"`
    Object  string `json:"object"`
    Created int64  `json:"created"`
    Model   string `json:"model"`
    Choices []struct {
        Message Message `json:"message"`
    } `json:"choices"`
}

type OpenAIErrorResponse struct {
    Error struct {
        Message string `json:"message"`
        Type    string `json:"type"`
        Code    string `json:"code"`
    } `json:"error"`
}

func NewAIService(apiKey string) (*AIService, error) {
    if apiKey == "" {
        return nil, fmt.Errorf("OpenAI API key is not set")
    }
    return &AIService{
        apiKey:  apiKey,
        baseURL: "https://api.openai.com/v1/chat/completions",
    }, nil
}

func (s *AIService) GenerateResponse(prompt string) (string, error) {
    request := OpenAIRequest{
        Model: "gpt-4o-mini",
        Store: true,
        Messages: []Message{
            {Role: "user", Content: prompt},
        },
    }

    jsonData, err := json.Marshal(request)
    if err != nil {
        return "", fmt.Errorf("error marshaling request: %w", err)
    }

    log.Printf("Sending request to OpenAI: %s", string(jsonData))

    req, err := http.NewRequest("POST", s.baseURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("error creating request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        var errorResponse OpenAIErrorResponse
        if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
            return "", fmt.Errorf("OpenAI API error, status code: %d", resp.StatusCode)
        }
        return "", fmt.Errorf("OpenAI API error: %s", errorResponse.Error.Message)
    }

    var response OpenAIResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return "", fmt.Errorf("error decoding response: %w", err)
    }

    if len(response.Choices) == 0 {
        return "", fmt.Errorf("no suggestions generated")
    }

    return response.Choices[0].Message.Content, nil
}

func (s *AIService) GenerateTaskSuggestions(task models.Task) (string, error) {
    prompt := generateAIPrompt(task)
    return s.GenerateResponse(prompt)
}

func generateAIPrompt(task models.Task) string {
    return fmt.Sprintf(`Please analyze this task and provide detailed suggestions:

Task Details:
- Title: %s
- Description: %s
- Status: %s
- Priority: %s

Please provide a structured response with the following sections:

1. Task Breakdown:
   - List of subtasks
   - Dependencies between subtasks

2. Time Estimation:
   - Estimated duration for each subtask
   - Total project timeline

3. Potential Challenges:
   - Technical challenges
   - Resource requirements
   - Risk factors

4. Implementation Recommendations:
   - Best practices
   - Tools and technologies
   - Testing strategies

5. Success Criteria:
   - Definition of done
   - Quality metrics
   - Validation steps`, 
    task.Title, 
    task.Description, 
    task.Status, 
    task.Priority)
}