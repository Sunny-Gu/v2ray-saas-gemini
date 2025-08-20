package service

import (
	"errors"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"gorm.io/gorm"
)

// NodeService provides node-related services.
type NodeService struct{}

// CreateNodeInput defines the input for creating a new node.
type CreateNodeInput struct {
	Name         string `json:"name" binding:"required"`
	ServerIP     string `json:"server_ip" binding:"required"`
	Port         int    `json:"port" binding:"required"`
	ProtocolType string `json:"protocol_type" binding:"required"`
	Encryption   string `json:"encryption"`
	Location     string `json:"location"`
	Weight       int    `json:"weight"`
}

// CreateNode creates a new node in the database.
func (s *NodeService) CreateNode(input CreateNodeInput) (*models.Node, error) {
	newNode := models.Node{
		Name:         input.Name,
		ServerIP:     input.ServerIP,
		Port:         input.Port,
		ProtocolType: input.ProtocolType,
		Encryption:   input.Encryption,
		Location:     input.Location,
		Weight:       input.Weight,
		Status:       models.NodeStatusOffline, // Default status
	}

	if err := database.DB.Create(&newNode).Error; err != nil {
		return nil, err
	}
	return &newNode, nil
}

// GetNode retrieves a single node by its ID.
func (s *NodeService) GetNode(id uint) (*models.Node, error) {
	var node models.Node
	if err := database.DB.First(&node, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("node not found")
		}
		return nil, err
	}
	return &node, nil
}

// ListNodes retrieves all nodes from the database.
func (s *NodeService) ListNodes() ([]models.Node, error) {
	var nodes []models.Node
	if err := database.DB.Find(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

// UpdateNodeInput defines the input for updating an existing node.
type UpdateNodeInput struct {
	Name         string `json:"name"`
	ServerIP     string `json:"server_ip"`
	Port         int    `json:"port"`
	ProtocolType string `json:"protocol_type"`
	Encryption   string `json:"encryption"`
	Location     string `json:"location"`
	Weight       int    `json:"weight"`
	Status       string `json:"status"`
}

// UpdateNode updates an existing node's information.
func (s *NodeService) UpdateNode(id uint, input UpdateNodeInput) (*models.Node, error) {
	node, err := s.GetNode(id)
	if err != nil {
		return nil, err
	}

	// Update fields
	node.Name = input.Name
	node.ServerIP = input.ServerIP
	node.Port = input.Port
	node.ProtocolType = input.ProtocolType
	node.Encryption = input.Encryption
	node.Location = input.Location
	node.Weight = input.Weight
	node.Status = models.NodeStatusType(input.Status)

	if err := database.DB.Save(&node).Error; err != nil {
		return nil, err
	}
	return node, nil
}

// DeleteNode deletes a node by its ID.
func (s *NodeService) DeleteNode(id uint) error {
	if err := database.DB.Delete(&models.Node{}, id).Error; err != nil {
		return err
	}
	return nil
}
