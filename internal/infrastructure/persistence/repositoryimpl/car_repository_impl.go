package repositoryimpl

import (
	"database/sql"

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
	result, err := c.DB.Exec("INSERT INTO cars(car_name, day_rate, month_rate, image) VALUES (?,?,?,?) RETURNING car_id", &car.CarName, &car.DayRate, &car.MonthRate, &car.Image)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	carId := int(id)
	car.CarId = &carId
	return car, nil
}

func (c *CarRepositoryImp) FindAll() ([]model.Car, error) {
	rows, err := c.DB.Query("SELECT * FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cars []model.Car
	for rows.Next() {
		var car model.Car
		err = rows.Scan(&car.CarId, &car.CarName, &car.DayRate, &car.MonthRate, &car.Image)
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
	err := row.Scan(&car.CarId, &car.CarName, &car.DayRate, &car.MonthRate, &car.Image)
	if err != nil {
		return nil, err
	}
	return &car, nil
}
func (c *CarRepositoryImp) Update(carReq *model.Car) error {
	_, err := c.DB.Exec("UPDATE cars SET car_name=?, day_rate=?, month_rate=?, image=? WHERE car_id=?", &carReq.CarName, &carReq.DayRate, &carReq.MonthRate, &carReq.Image, &carReq.CarId)

	if err != nil {
		return err
	}
	return nil
}

func (c *CarRepositoryImp) Delete(carId int) error {
	_, err := c.DB.Exec("DELETE FROM cars WHERE car_id=?", carId)
	if err != nil {
		return err
	}
	return nil
}
