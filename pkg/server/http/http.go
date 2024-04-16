package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/peifengll/SpaceRepetition/pkg/log"
	"net/http"
	"time"
)

type ServerFont struct {
	*gin.Engine
	httpSrv *http.Server
	host    string
	port    int
	logger  *log.Logger
}
type OptionFont func(s *ServerFont)

func NewServerFont(engine *gin.Engine, logger *log.Logger, opts ...OptionFont) *ServerFont {
	s := &ServerFont{
		Engine: engine,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
func WithServerFontHost(host string) OptionFont {
	return func(s *ServerFont) {
		s.host = host
	}
}
func WithServerFontPort(port int) OptionFont {
	return func(s *ServerFont) {
		s.port = port
	}
}

func (s *ServerFont) Start(ctx context.Context) error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s,
	}

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}
func (s *ServerFont) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.logger.Sugar().Fatal("ServerFont forced to shutdown: ", err)
	}

	s.logger.Sugar().Info("ServerFont exiting")
	return nil
}

type ServerAdmin struct {
	*gin.Engine
	httpSrv *http.Server
	host    string
	port    int
	logger  *log.Logger
}
type OptionAdmin func(s *ServerAdmin)

func NewServerAdmin(engine *gin.Engine, logger *log.Logger, opts ...OptionAdmin) *ServerAdmin {
	s := &ServerAdmin{
		Engine: engine,
		logger: logger,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
func WithServerAdminHost(host string) OptionAdmin {
	return func(s *ServerAdmin) {
		s.host = host
	}
}
func WithServerAdminPort(port int) OptionAdmin {
	return func(s *ServerAdmin) {
		s.port = port
	}
}

func (s *ServerAdmin) Start(ctx context.Context) error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s,
	}

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Sugar().Fatalf("listen: %s\n", err)
	}

	return nil
}
func (s *ServerAdmin) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.logger.Sugar().Fatal("ServerAdmin forced to shutdown: ", err)
	}

	s.logger.Sugar().Info("ServerAdmin exiting")
	return nil
}
