package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//NewEngine godoc
// NewEngine creates the gin Engine and returns it
func NewEngine() *gin.Engine {
	engine := gin.New()

	engine.Use(cors.Default())

	return engine
}
