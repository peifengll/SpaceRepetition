package server

import (
	"github.com/gin-gonic/gin"
	apiV1 "github.com/peifengll/SpaceRepetition/api/v1"
	"github.com/peifengll/SpaceRepetition/docs"
	"github.com/peifengll/SpaceRepetition/internal/handler"
	"github.com/peifengll/SpaceRepetition/internal/middleware"
	"github.com/peifengll/SpaceRepetition/pkg/jwt"
	"github.com/peifengll/SpaceRepetition/pkg/log"
	"github.com/peifengll/SpaceRepetition/pkg/server/http"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServerFont(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	floderHandler *handler.FloderHandler,
	deckHandler *handler.DeckHandler,
	knowledgeHandler *handler.KnowledgeHandler,
	recordHandler *handler.RecordHandler,
	sectionHandler *handler.SectionHandler,
	announcementsHandler *handler.AnnouncementsHandler,
) *http.ServerFont {
	gin.SetMode(gin.ReleaseMode)
	s := http.NewServerFont(
		gin.Default(),
		logger,
		http.WithServerFontHost(conf.GetString("http_font.host")),
		http.WithServerFontPort(conf.GetInt("http_font.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "This is SpaceRepetition!",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
			strictAuthRouter.POST("/readann", announcementsHandler.SetAnnouncementHaveRead)
			strictAuthRouter.GET("/pushanns", announcementsHandler.GetAnnouncementNoRead)
		}

		decks := v1.Group("/decks").Use(middleware.StrictAuth(jwt, logger))
		{
			// 跟文件夹相关
			decks.GET("/", floderHandler.GetFloder)
			decks.POST("/floder", floderHandler.AddFloder)
			decks.DELETE("/floder/:id", floderHandler.DeleteFloder)
			decks.GET("/floder/:id", floderHandler.GetFloderById)
			decks.PUT("/floder", floderHandler.UpdateFloder)

			//	跟牌组
			decks.POST("/deck", deckHandler.AddDeck)
			//decks.GET("/deck/search", deckHandler.AddDeck)
			decks.GET("/deck/:id", deckHandler.GetDeckById)
			decks.PUT("/deck", deckHandler.UpdateDeck)
			decks.DELETE("/deck/:id", deckHandler.DeleteDeck)

			// 跟牌组里的章节
			decks.POST("/deck/section", sectionHandler.AddSection)
			decks.DELETE("/deck/section/:id", sectionHandler.DeleteSection)
			decks.PUT("/deck/section", sectionHandler.UpdateSection)
			//	跟卡片
			decks.POST("/card", knowledgeHandler.AddCard)
			decks.GET("/card/:id", knowledgeHandler.GetCard)
			decks.PUT("/card", knowledgeHandler.UpdateCard)
			decks.DELETE("/card/:id", knowledgeHandler.DeleteCard)
			decks.GET("/card/search/:content", knowledgeHandler.SearchCards)
			decks.PUT("/card/toreview", knowledgeHandler.ChooseToReview)

			// 复习数据导出功能
			decks.GET("/reviewinfo/export", recordHandler.ExportReviewInfo)
		}

		review := v1.Group("/review").Use(middleware.StrictAuth(jwt, logger))
		{
			// 这个人该复习的所有
			review.GET("/", knowledgeHandler.GetAllReview)
			// 一次复习操作
			review.PUT("/option", knowledgeHandler.ReviewOpt)
			//取消一张卡片的复习调度
			review.PUT("/card/cancel")

		}

	}

	return s
}

func NewHTTPServerAdmin(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	adminHandler *handler.AdminHandler,
	announcementsHandler *handler.AnnouncementsHandler,
) *http.ServerAdmin {
	gin.SetMode(gin.ReleaseMode)
	s := http.NewServerAdmin(
		gin.Default(),
		logger,
		http.WithServerAdminHost(conf.GetString("http_admin.host")),
		http.WithServerAdminPort(conf.GetInt("http_admin.port")),
	)

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "This is SpaceRepetition  Admin!",
		})
	})
	noAuthRouter := s.Group("admin/")
	{
		noAuthRouter.POST("/login", adminHandler.Login)
		noAuthRouter.GET("/announcement", announcementsHandler.ShowAnnouncement)
		AuthRouter := noAuthRouter.Group("").Use(middleware.StrictAuth(jwt, logger))
		{
			AuthRouter.POST("/register", adminHandler.Register)
			AuthRouter.PUT("/privileges", adminHandler.UpdatePrivileges)
			AuthRouter.DELETE("/account", adminHandler.DelAdminAccount)
			AuthRouter.GET("/accounts", adminHandler.ShowAdminAccounts)

			AuthRouter.POST("/announcement", announcementsHandler.AddAnnouncement)
			AuthRouter.DELETE("/announcement", announcementsHandler.DelAnnouncement)
			AuthRouter.PUT("/announcement", announcementsHandler.UpdateAnnouncement)

		}
	}

	return s
}
