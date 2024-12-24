package models

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	GinContext *gin.Context `json:"-"`
	Data       interface{}  `json:"data,omitempty"`
	Pagination interface{}  `json:"page_info,omitempty"`
	StatusCode int          `json:"status_code"`
	Message    string       `json:"message"`
	SubMessage string       `json:"sub_message"`
}

type ErrorResponse struct {
	GinContext *gin.Context `json:"-"`
	Error      Error        `json:"error"`
}

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
}

type TokenResponse struct {
	Token      string `json:"token"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
