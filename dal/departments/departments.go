package departments

import (
	"context"
	models "crud-apis-db-app/modules/departments/models"
	"crud-apis-db-app/shared"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/labstack/gommon/log"
)

type DepartmentsDal interface {
	IsDepartmentAvailable(ctx context.Context, id int) (bool, error)
	InsertDepartment(ctx context.Context, departments *[]models.Department) error
	SelectDepartments(ctx context.Context) (*[]models.Department, error)
	SelectDepartmentById(ctx context.Context, id int) (*models.Department, error)
	UpdateDepartmentsById(ctx context.Context, departments *[]models.Department) (*[]models.Department, error)
	DeleteDepartmentById(ctx context.Context, ids []int) (bool, error)
}

type Departments struct {
	Deps *shared.Deps
}

func NewDepartmentsDal(deps *shared.Deps) DepartmentsDal {
	return &Departments{
		Deps: deps,
	}
}

func (e *Departments) IsDepartmentAvailable(ctx context.Context, id int) (bool, error) {
	key := fmt.Sprintf("department:%d", id)
	exists, err := e.Deps.Database.RedisDb.Exists(ctx, key)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (e *Departments) InsertDepartment(ctx context.Context, departments *[]models.Department) error {
	var key string
	var value []byte
	var err error

	for _, department := range *departments {
		key = fmt.Sprintf("department:%d", department.Id)

		department.LastModified = time.Now()

		value, err = json.Marshal(department)
		if err != nil {
			log.Errorf("error while marshalling data", map[string]interface{}{"error": err, "key": key})
			return err
		}

		if len(value) <= 0 {
			log.Warnf("empty data: skipping insertion for key", map[string]interface{}{"key": key})
			return errors.New("empty data error")
		}

		err = e.Deps.Database.RedisDb.Create(ctx, key, value, 0)
		if err != nil {
			log.Errorf("error while inserting data", map[string]interface{}{"error": err, "key": key})
			return err
		}
	}

	return nil
}

func (e *Departments) SelectDepartments(ctx context.Context) (*[]models.Department, error) {
	var keys []string
	var err error
	var value string
	var department models.Department
	var departments []models.Department

	pattern := "department:*"

	keys, err = e.Deps.Database.RedisDb.Keys(ctx, pattern)
	if err != nil {
		return nil, err
	}
	if len(keys) <= 0 {
		return nil, errors.New("no keys found")
	}

	for _, key := range keys {
		value, err = e.Deps.Database.RedisDb.Read(ctx, key)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(value), &department)
		if err != nil {
			return nil, err
		}

		departments = append(departments, department)
	}

	return &departments, nil
}

func (e *Departments) SelectDepartmentById(ctx context.Context, id int) (*models.Department, error) {
	var err error
	var value string
	var department models.Department

	key := fmt.Sprintf("department:%d", id)
	value, err = e.Deps.Database.RedisDb.Read(ctx, key)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(value), &department)
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (e *Departments) UpdateDepartmentsById(ctx context.Context, departments *[]models.Department) (*[]models.Department, error) {
	var key string
	var value []byte
	var err error
	var exists bool
	var updatedValue string
	var updatedDepartment models.Department
	var updatedDepartments []models.Department

	for _, department := range *departments {
		key = fmt.Sprintf("department:%d", department.Id)

		exists, err = e.IsDepartmentAvailable(ctx, department.Id)
		if err != nil {
			return nil, err
		}

		if exists {
			department.LastModified = time.Now()

			value, err = json.Marshal(department)
			if err != nil {
				log.Errorf("error while marshalling data", map[string]interface{}{"error": err, "key": key})
				return nil, err
			}

			if len(value) <= 0 {
				log.Warnf("empty data: skipping insertion for key", map[string]interface{}{"key": key})
				return nil, errors.New("empty data error")
			}

			err = e.Deps.Database.RedisDb.Update(ctx, key, value, 0)
			if err != nil {
				log.Errorf("error while inserting data", map[string]interface{}{"error": err, "key": key})
				return nil, err
			}

			updatedValue, err = e.Deps.Database.RedisDb.Read(ctx, key)
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal([]byte(updatedValue), &updatedDepartment)
			if err != nil {
				return nil, err
			}

			updatedDepartments = append(updatedDepartments, updatedDepartment)
		}
	}

	return &updatedDepartments, nil
}

func (e *Departments) DeleteDepartmentById(ctx context.Context, ids []int) (bool, error) {
	var err error
	var key string

	for _, id := range ids {
		key = fmt.Sprintf("department:%d", id)
		err = e.Deps.Database.RedisDb.Delete(ctx, key)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
