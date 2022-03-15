package config

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	ServerAddr           string
	ServerReadingTimeout time.Duration
	ServerWritingTimeout time.Duration
	ServerMaxHeaderBytes int
	ServerPort           string
	ServerHandler        *gin.Engine
}
