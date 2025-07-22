package repositoryimpl

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/steradian_test/internal/domain/model"
)

type CarRepositoryImp struct {
	DB *sql.DB
}

func NewCarRepositoryImpl(db *sql.DB) *CarRepositoryImp {
	return &CarRepositoryImp{
		DB: db,
	}
}

func (c *CarRepositoryImp) Create(car *model.Car) (*model.Car, error) {
	query := `
	INSERT INTO cars(car_name, day_rate, month_rate, image) 
	VALUES (?,?,?,?) RETURNING car_id
	`
	var carId int
	err := c.DB.QueryRow(
		query,
		&car.CarName,
		&car.DayRate,
		&car.MonthRate,
		&car.Image).Scan(&carId)

	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"Failed to create car, please check your data again",
		)
	}
	car.CarId = &carId
	return car, nil
}

func (c *CarRepositoryImp) FindAll() ([]model.Car, error) {
	rows, err := c.DB.Query("SELECT * FROM cars")
	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"No Data!",
		)
	}
	defer rows.Close()

	cars := []model.Car{}
	for rows.Next() {
		var car model.Car
		err = rows.Scan(
			&car.CarId,
			&car.CarName,
			&car.DayRate,
			&car.MonthRate,
			&car.Image)

		if err != nil {
			//
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (c *CarRepositoryImp) FindById(id int) (*model.Car, error) {
	row := c.DB.QueryRow("SELECT * FROM cars WHERE car_id=?", id)
	var car model.Car
	err := row.Scan(
		&car.CarId,
		&car.CarName,
		&car.DayRate,
		&car.MonthRate,
		&car.Image)

	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"Car not found!",
		)
	}
	return &car, nil
}
func (c *CarRepositoryImp) Update(carReq *model.Car) (*model.Car, error) {
	query := `
	UPDATE cars SET car_name=?, day_rate=?, month_rate=?, image=? 
	WHERE car_id=? RETURNING *
	`
	row := c.DB.QueryRow(
		query,
		&carReq.CarName,
		&carReq.DayRate,
		&carReq.MonthRate,
		&carReq.Image,
		&carReq.CarId)

	var car model.Car
	err := row.Scan(
		&car.CarId,
		&car.CarName,
		&car.DayRate,
		&car.MonthRate,
		&car.Image)

	if err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"Failed to update car, please check your data again",
		)
	}
	return &car, nil
}

func (c *CarRepositoryImp) Delete(carId int) error {
	query := `
	DELETE FROM cars WHERE car_id=?
	`
	_, err := c.DB.Exec(query, carId)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Car not found!",
		)
	}
	return nil
}
