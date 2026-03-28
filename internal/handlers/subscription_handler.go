package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Alexeyts0Y/test_task_em/internal/models"
	"github.com/Alexeyts0Y/test_task_em/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	service service.SubscriptionService
}

func NewHandler(s service.SubscriptionService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(c *gin.Context) {
	var sub models.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		slog.Warn("Invalid create input", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Create(c.Request.Context(), sub)
	if err != nil {
		slog.Error("Failed to create subscription", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create subscription"})
		return
	}

	slog.Info("Subscription created", "id", id)
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	sub, err := h.service.Get(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
			return
		}
		slog.Error("Failed to get subscription", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	c.JSON(http.StatusOK, sub)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	var sub models.Subscription
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.Update(c.Request.Context(), id, sub)
	if err != nil {
		slog.Error("Failed to update subscription", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		return
	}

	slog.Info("Subscription updated", "id", id)
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		slog.Error("Failed to delete subscription", "id", id, "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
		return
	}

	slog.Info("Subscription deleted", "id", id)
	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) List(c *gin.Context) {
	userID := c.Query("user_id")
	serviceName := c.Query("service_name")

	subs, err := h.service.List(c.Request.Context(), userID, serviceName)
	if err != nil {
		slog.Error("Failed to list subscriptions", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, subs)
}

func (h *Handler) CalculateCost(c *gin.Context) {
	var req models.CostRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		slog.Warn("Invalid cost calculation request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total, err := h.service.CalculateTotalCost(c.Request.Context(), req)
	if err != nil {
		slog.Error("Cost calculation failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "calculation error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":      req.UserID,
		"service_name": req.ServiceName,
		"period_start": req.StartDate,
		"period_end":   req.EndDate,
		"total_cost":   total,
	})
}
