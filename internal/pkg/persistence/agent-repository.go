package persistence

import (
	models "github.com/mazeyqian/go-gin-gee/internal/pkg/models/agent"
)

type AgentRepository struct{}

var agentRepository *AgentRepository

func GetAgentRepository() *AgentRepository {
	if agentRepository == nil {
		agentRepository = &AgentRepository{}
	}
	return agentRepository
}

func (r *AgentRepository) Mock(res *models.Response) (*models.ResponseData, error) {
	return &res.Data, nil
}
