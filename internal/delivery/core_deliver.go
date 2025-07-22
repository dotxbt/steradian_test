package delivery

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/steradian_test/internal/domain/usecase"
	"github.com/steradian_test/internal/handler"
	"github.com/steradian_test/internal/infrastructure/persistence/repositoryimpl"
)

func CoreDelivery(app *fiber.App, db *sql.DB) {
	api := app.Group("/api")
	carRepo := repositoryimpl.NewCarRepositoryImpl(db)
	carUsecase := usecase.NewCarUsecase(carRepo)
	handler.NewCarHandler(api, carUsecase)

	orderRepo := repositoryimpl.NewOrderRepositoryImpl(db)
	orderUseCase := usecase.NewOrderUsecase(orderRepo)
	handler.NewOrderHandler(api, orderUseCase)
}
