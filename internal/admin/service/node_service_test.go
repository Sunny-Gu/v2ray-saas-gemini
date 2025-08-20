package service

import (
	"testing"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestNodeCRUD(t *testing.T) {
	defer database.DB.Exec("DELETE FROM nodes")

	nodeService := NodeService{}

	// --- Test CreateNode ---
	createInput := CreateNodeInput{
		Name:         "Test Node",
		ServerIP:     "1.2.3.4",
		Port:         12345,
		ProtocolType: "VMESS",
		Location:     "Test Location",
	}
	node, err := nodeService.CreateNode(createInput)
	assert.NoError(t, err)
	assert.NotNil(t, node)
	assert.Equal(t, createInput.Name, node.Name)

	// --- Test GetNode ---
	retrievedNode, err := nodeService.GetNode(node.ID)
	assert.NoError(t, err)
	assert.Equal(t, node.Name, retrievedNode.Name)

	// --- Test ListNodes ---
	nodes, err := nodeService.ListNodes()
	assert.NoError(t, err)
	assert.Len(t, nodes, 1)

	// --- Test UpdateNode ---
	updateInput := UpdateNodeInput{
		Name:   "Updated Node Name",
		Status: string(models.NodeStatusOnline),
		// Copy other fields to avoid zeroing them out
		ServerIP:     node.ServerIP,
		Port:         node.Port,
		ProtocolType: node.ProtocolType,
		Location:     node.Location,
	}
	updatedNode, err := nodeService.UpdateNode(node.ID, updateInput)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Node Name", updatedNode.Name)
	assert.Equal(t, models.NodeStatusOnline, updatedNode.Status)

	// --- Test DeleteNode ---
	err = nodeService.DeleteNode(node.ID)
	assert.NoError(t, err)

	// Verify it's deleted
	_, err = nodeService.GetNode(node.ID)
	assert.Error(t, err)
	assert.Equal(t, "node not found", err.Error())
}
