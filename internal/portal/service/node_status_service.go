package service

import (
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// NodeStatusService provides services for querying node statuses.
type NodeStatusService struct{}

// NodeStatusOutput defines the publicly visible fields for a node's status.
type NodeStatusOutput struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Status   string `json:"status"`
}

// ListAllNodeStatuses retrieves the status of all nodes.
func (s *NodeStatusService) ListAllNodeStatuses() ([]NodeStatusOutput, error) {
	var nodes []models.Node
	if err := database.DB.Find(&nodes).Error; err != nil {
		return nil, err
	}

	var output []NodeStatusOutput
	for _, node := range nodes {
		output = append(output, NodeStatusOutput{
			Name:     node.Name,
			Location: node.Location,
			Status:   string(node.Status),
		})
	}

	return output, nil
}
