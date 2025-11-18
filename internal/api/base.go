package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/rss-feed/internal/db"
)

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok",
		"message": "System is healthy",
	})

}

func (s *Server) createFeed(c *gin.Context) {
	type createFeedRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
	}

	var req createFeedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	dbParams := db.CreateFeedParams{
		FeedName:    req.Name,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		Url:         req.URL,
	}

	feed, err := s.store.CreateFeed(c, dbParams)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("feed : %v", feed)

	c.JSON(http.StatusOK, gin.H{
		"message": "Feed Created Successfully",
		"data":    feed,
	})
}
func (s *Server) getAllFeeds(c *gin.Context) {

	var req db.GetAllFeedsParams
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	req = db.GetAllFeedsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}
	if err := c.ShouldBindQuery(&req); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	feeds, err := s.store.GetAllFeeds(c, req)
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("feeds : %v", feeds)

	c.JSON(http.StatusOK, gin.H{
		"message": "Feed Fetched Successfully",
		"data":    feeds,
	})
}
func (s *Server) getOneFeed(c *gin.Context) {

	req := c.Param("id")

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(req)
	if err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	feed, err := s.store.GetOneFeedById(c, int32(id))
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("feed : %v", feed)

	c.JSON(http.StatusOK, gin.H{
		"message": "Feed Fetched Successfully",
		"data":    feed,
	})
}
func (s *Server) updateFeed(c *gin.Context) {

	var req db.UpdateFeedParams
	id := c.Param("id")

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	feed, err := s.store.UpdateFeed(c, db.UpdateFeedParams{
		FeedID:      int32(idInt),
		FeedName:    req.FeedName,
		Url:         req.Url,
		Description: req.Description,
	})
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Printf("feed : %v", feed)

	c.JSON(http.StatusOK, gin.H{
		"message": "Feed Created Successfully",
		"data":    feed,
	})
}
func (s *Server) deleteFeed(c *gin.Context) {

	req := c.Param("id")

	if err := c.ShouldBindUri(&req); err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(req)
	if err != nil {
		errResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = s.store.DeleteFeed(c, int32(id))
	if err != nil {
		errResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Feed Deleted Successfully",
	})
}
