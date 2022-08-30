package server

import (
	"github.com/gin-gonic/gin"
	"github.com/henribas/cpf/pkg/cpf"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	srv := &Server{router: gin.Default()}
	srv.AddRoutes()
	return srv
}

func (s *Server) Run() error {
	err := s.router.Run(":8080")
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *Server) AddRoutes() {
	v1 := s.router.Group("/cpf")
	{
		v1.GET("/validar", cpf.ValidarEndpoint)
	}

}
