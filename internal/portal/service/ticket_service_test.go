package service

import (
	"testing"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateTicket(t *testing.T) {
	defer func() {
		database.DB.Exec("DELETE FROM ticket_replies")
		database.DB.Exec("DELETE FROM tickets")
		database.DB.Exec("DELETE FROM users")
	}()

	ticketService := TicketService{}
	user := createUserWithBalance("ticket@example.com", 0)

	input := CreateTicketInput{
		Title:   "Test Ticket",
		Message: "This is the first message.",
	}

	ticket, err := ticketService.CreateTicket(user.ID, input)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, input.Title, ticket.Title)
	assert.Equal(t, user.ID, ticket.UserID)
	assert.Equal(t, models.TicketStatusOpen, ticket.Status)
	assert.Len(t, ticket.Replies, 1)
	assert.Equal(t, input.Message, ticket.Replies[0].Message)
}

func TestListAndGetTicket(t *testing.T) {
	defer func() {
		database.DB.Exec("DELETE FROM ticket_replies")
		database.DB.Exec("DELETE FROM tickets")
		database.DB.Exec("DELETE FROM users")
	}()

	ticketService := TicketService{}
	user1 := createUserWithBalance("user1@example.com", 0)
	user2 := createUserWithBalance("user2@example.com", 0)

	// Create tickets for user1
	ticketService.CreateTicket(user1.ID, CreateTicketInput{Title: "Ticket 1", Message: "Msg1"})
	ticketService.CreateTicket(user1.ID, CreateTicketInput{Title: "Ticket 2", Message: "Msg2"})
	// Create ticket for user2
	ticket2, _ := ticketService.CreateTicket(user2.ID, CreateTicketInput{Title: "Ticket 3", Message: "Msg3"})

	// --- Test ListUserTickets ---
	user1Tickets, err := ticketService.ListUserTickets(user1.ID)
	assert.NoError(t, err)
	assert.Len(t, user1Tickets, 2)

	// --- Test GetTicketDetails ---
	details, err := ticketService.GetTicketDetails(user2.ID, ticket2.ID)
	assert.NoError(t, err)
	assert.NotNil(t, details)
	assert.Equal(t, ticket2.Title, details.Title)
	assert.Len(t, details.Replies, 1)

	// --- Test GetTicketDetails (Access Denied) ---
	_, err = ticketService.GetTicketDetails(user1.ID, ticket2.ID)
	assert.Error(t, err)
}

func TestReplyToTicket(t *testing.T) {
	defer func() {
		database.DB.Exec("DELETE FROM ticket_replies")
		database.DB.Exec("DELETE FROM tickets")
		database.DB.Exec("DELETE FROM users")
	}()

	ticketService := TicketService{}
	user := createUserWithBalance("reply@example.com", 0)
	ticket, _ := ticketService.CreateTicket(user.ID, CreateTicketInput{Title: "Reply Test", Message: "Initial msg"})

	replyInput := ReplyToTicketInput{Message: "This is a reply."}
	reply, err := ticketService.ReplyToTicket(user.ID, ticket.ID, replyInput)

	assert.NoError(t, err)
	assert.NotNil(t, reply)
	assert.Equal(t, replyInput.Message, reply.Message)
	assert.Equal(t, user.ID, reply.UserID)
	assert.False(t, reply.IsAdmin)

	// Verify the reply is associated with the ticket
	details, _ := ticketService.GetTicketDetails(user.ID, ticket.ID)
	assert.Len(t, details.Replies, 2)
}
