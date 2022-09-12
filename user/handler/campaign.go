package handler

import (
	"golang-campaign/user/handler/campaign"
	"golang-campaign/user/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {

		response := helper.APIResponse("Error get campaigns", http.StatusBadRequest, "error", campaign.FormatCampaigns(campaigns))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("list of Campaign", http.StatusOK, "success", campaigns)
	c.JSON(http.StatusOK, response)

}
