package server

//type Server struct {
//	echo   *echo.Echo
//	logger *slog.Logger
//}
//
//func NewServer(cfg config.Config, h Handlers, logger *slog.Logger) (err error) {
//	e := echo.New()
//
//	//здесь должна быть мидлваря
//
//	err = e.Start(":" + cfg.ServerPort)
//	if err != nil {
//		err = fmt.Errorf("error starting server: %w", err)
//		return
//	}
//
//	productGroup := e.Group("/products")
//
//	productGroup.GET("/:id", h.Products.Read)
//}
