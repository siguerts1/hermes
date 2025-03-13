package handlers

import (
	"hermes/internal/services/thirdparty"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GitHubHandler struct {
	client *thirdparty.GitHubAPI
}

func NewGitHubHandler() *GitHubHandler {
	return &GitHubHandler{
		client: thirdparty.NewGitHubAPI(),
	}
}

func (h *GitHubHandler) GetRepositories(c echo.Context) error {
	username := c.QueryParam("username")
	if username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "username is required"})
	}

	repos, err := h.client.FetchRepositories(username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, repos)
}
