package services

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
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
    Temperature float32 `json:"temperature,omitempty"`
    MaxTokens  int     `json:"max_tokens,omitempty"`
}

type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type OpenAIResponse struct {
    Choices []struct {
        Message Message `json:"message"`
    } `json:"choices"`
}

type OpenAIErrorResponse struct {
    Error struct {
        Message string `json:"message"`
        Code    int    `json:"code"`
    } `json:"error"`
}

func NewAIService(apiKey string) (*AIService, error) {
    if apiKey == "" {
        return nil, fmt.Errorf("OpenRouter API key is not set")
    }
    return &AIService{
        apiKey:  apiKey,
        baseURL: "https://openrouter.ai/api/v1/chat/completions",
    }, nil
}

func (s *AIService) GenerateResponse(prompt string) (string, error) {
    request := OpenAIRequest{
        Model:    "deepseek/deepseek-r1:free", 
        Messages: []Message{
            {Role: "user", Content: prompt},
        },
        Temperature: 0.7,
        MaxTokens:   1000,
    }

    jsonData, err := json.Marshal(request)
    if err != nil {
        return "", fmt.Errorf("error marshaling request: %w", err)
    }

    log.Printf("Sending request to OpenRouter: %s", string(jsonData))

    req, err := http.NewRequest("POST", s.baseURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", fmt.Errorf("error creating request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))
    req.Header.Set("X-Title", "TaskAssistant")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    log.Printf("Raw response from OpenRouter: %s", string(body))

    if resp.StatusCode != http.StatusOK {
        var errorResponse OpenAIErrorResponse
        if err := json.Unmarshal(body, &errorResponse); err != nil {
            return "", fmt.Errorf("OpenRouter API error, status code: %d, response: %s", 
                resp.StatusCode, string(body))
        }
        return "", fmt.Errorf("OpenRouter API error [%d]: %s", 
            errorResponse.Error.Code, errorResponse.Error.Message)
    }

    var response OpenAIResponse
    if err := json.Unmarshal(body, &response); err != nil {
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
