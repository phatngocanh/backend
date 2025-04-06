package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/phat-ngoc-anh/backend/internal/controller/http/v1"
)

type Server struct {
	studentHandler *v1.StudentHandler
	invoiceHandler *v1.InvoiceHandler
}

func NewServer(studentHandler *v1.StudentHandler, invoiceHandler *v1.InvoiceHandler) *Server {
	return &Server{
		studentHandler: studentHandler,
		invoiceHandler: invoiceHandler,
	}
}

func (s *Server) Run() {
	router := gin.Default()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	v1.MapRoutes(router, s.studentHandler, s.invoiceHandler)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
