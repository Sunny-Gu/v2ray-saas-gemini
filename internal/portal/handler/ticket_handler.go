package handler

import (
	"net/http"
	"strconv"
	"v2ray-saas-gemini/internal/portal/service"

	"github.com/gin-gonic/gin"
)

// TicketHandler handles ticket-related API requests.
type TicketHandler struct {
	service service.TicketService
}

func NewTicketHandler() *TicketHandler {
	return &TicketHandler{service: service.TicketService{}}
}

func (h *TicketHandler) CreateTicket(c *gin.Context) {
	var input service.CreateTicketInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, _ := c.Get("userID")
	id, _ := userID.(uint)

	ticket, err := h.service.CreateTicket(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, ticket)
}

func (h *TicketHandler) ListUserTickets(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := userID.(uint)

	tickets, err := h.service.ListUserTickets(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tickets"})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) GetTicketDetails(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := userID.(uint)
	ticketID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	ticket, err := h.service.GetTicketDetails(id, uint(ticketID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) ReplyToTicket(c *gin.Context) {
	userID, _ := c.Get("userID")
	id, _ := userID.(uint)
	ticketID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var input service.ReplyToTicketInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reply, err := h.service.ReplyToTicket(id, uint(ticketID), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, reply)
}
