package service

import (
	"errors"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// TicketService provides ticket-related services.
type TicketService struct{}

// CreateTicketInput defines the input for creating a new ticket.
type CreateTicketInput struct {
	Title   string `json:"title" binding:"required,max=255"`
	Message string `json:"message" binding:"required"`
}

// CreateTicket creates a new support ticket.
func (s *TicketService) CreateTicket(userID uint, input CreateTicketInput) (*models.Ticket, error) {
	ticket := models.Ticket{
		UserID: userID,
		Title:  input.Title,
		Status: models.TicketStatusOpen,
	}
	reply := models.TicketReply{
		UserID:  userID,
		Message: input.Message,
		IsAdmin: false,
	}
	ticket.Replies = append(ticket.Replies, reply)

	if err := database.DB.Create(&ticket).Error; err != nil {
		return nil, errors.New("failed to create ticket")
	}
	return &ticket, nil
}

// ListUserTickets lists all tickets for a specific user.
func (s *TicketService) ListUserTickets(userID uint) ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := database.DB.Where("user_id = ?", userID).Order("updated_at desc").Find(&tickets).Error
	return tickets, err
}

// GetTicketDetails retrieves a ticket and its replies.
func (s *TicketService) GetTicketDetails(userID, ticketID uint) (*models.Ticket, error) {
	var ticket models.Ticket
	// Ensure the user owns the ticket
	err := database.DB.Where("id = ? AND user_id = ?", ticketID, userID).Preload("Replies").First(&ticket).Error
	return &ticket, err
}

// ReplyToTicketInput defines the input for replying to a ticket.
type ReplyToTicketInput struct {
	Message string `json:"message" binding:"required"`
}

// ReplyToTicket adds a new reply to a ticket.
func (s *TicketService) ReplyToTicket(userID, ticketID uint, input ReplyToTicketInput) (*models.TicketReply, error) {
	// First, verify the user owns the ticket
	_, err := s.GetTicketDetails(userID, ticketID)
	if err != nil {
		return nil, errors.New("ticket not found or access denied")
	}

	reply := models.TicketReply{
		TicketID: ticketID,
		UserID:   userID,
		Message:  input.Message,
		IsAdmin:  false,
	}

	if err := database.DB.Create(&reply).Error; err != nil {
		return nil, errors.New("failed to post reply")
	}
	return &reply, nil
}
