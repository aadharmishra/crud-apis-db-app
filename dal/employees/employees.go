package employees

import (
	"context"
	employeesDalModel "crud-apis-db-app/dal/employees/models"
	models "crud-apis-db-app/modules/employees/models"
	"crud-apis-db-app/shared"
	"crud-apis-db-app/utils/common"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeesDal interface {
	IsEmployeeAvailable(ctx context.Context, id int) (bool, error)
	InsertEmployee(ctx context.Context, employees *[]models.Employee) error
	SelectEmployees(ctx context.Context) (*[]models.Employee, error)
	SelectEmployeeById(ctx context.Context, id int) (*models.Employee, error)
	UpdateEmployeesById(ctx context.Context, employees *[]models.Employee) (*[]models.Employee, error)
	DeleteEmployeeById(ctx context.Context, ids []int) (bool, error)
}

type Employees struct {
	Deps *shared.Deps
}

func NewEmployeesDal(deps *shared.Deps) EmployeesDal {
	return &Employees{
		Deps: deps,
	}
}

func (e *Employees) DeleteEmployeeById(ctx context.Context, ids []int) (bool, error) {
	database := e.Deps.Config.Get().Mongodb.Db
	collection := e.Deps.Config.Get().Mongodb.Collection

	filter := bson.M{"id": bson.M{"$in": ids}}

	err := e.Deps.Database.MongoDb.Delete(ctx, database, collection, filter)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (e *Employees) InsertEmployee(ctx context.Context, employees *[]models.Employee) error {
	if employees == nil {
		return errors.New("empty request")
	}

	var dalEmployees []employeesDalModel.Employee
	for _, employee := range *employees {
		dalEmployee := &employeesDalModel.Employee{
			ID:           employee.ID,
			Username:     employee.Username,
			Email:        employee.Email,
			Age:          employee.Age,
			IsAdmin:      employee.IsAdmin,
			Dob:          employee.Dob,
			Details:      employee.Details,
			LastModified: primitive.NewDateTimeFromTime(time.Now()),
		}

		dalEmployees = append(dalEmployees, *dalEmployee)
	}

	database := e.Deps.Config.Get().Mongodb.Db
	collection := e.Deps.Config.Get().Mongodb.Collection

	var employeeList []interface{}
	for _, employee := range dalEmployees {
		employeeList = append(employeeList, employee)
	}

	err := e.Deps.Database.MongoDb.Create(ctx, database, collection, employeeList)
	if err != nil {
		return err
	}

	return nil
}

func (e *Employees) IsEmployeeAvailable(ctx context.Context, id int) (bool, error) {
	database := e.Deps.Config.Get().Mongodb.Db
	collection := e.Deps.Config.Get().Mongodb.Collection

	filter := bson.M{}

	count, err := e.Deps.Database.MongoDb.Exists(ctx, database, collection, filter)
	if err != nil {
		return false, err
	}

	if count <= 0 {
		return false, errors.New("no employee available")
	}

	return true, nil
}

func (e *Employees) SelectEmployeeById(ctx context.Context, id int) (*models.Employee, error) {
	database := e.Deps.Config.Get().Mongodb.Db
	collection := e.Deps.Config.Get().Mongodb.Collection

	filter := bson.M{"id": id}

	cursor, err := e.Deps.Database.MongoDb.Read(ctx, database, collection, filter)
	if err != nil || cursor == nil {
		return nil, err
	}

	var employee models.Employee
	for cursor.Next(ctx) {
		err = cursor.Decode(&employee)
		if err != nil {
			return nil, errors.New("error while mapping response")
		}
	}

	return &employee, nil
}

func (e *Employees) SelectEmployees(ctx context.Context) (*[]models.Employee, error) {
	database := e.Deps.Config.Get().Mongodb.Db
	collection := e.Deps.Config.Get().Mongodb.Collection

	filter := bson.M{}

	cursor, err := e.Deps.Database.MongoDb.Read(ctx, database, collection, filter)
	if err != nil || cursor == nil {
		return nil, err
	}

	var employees []models.Employee
	for cursor.Next(ctx) {
		var employee models.Employee
		err = cursor.Decode(&employee)
		if err != nil {
			return nil, errors.New("error while mapping response")
		}
		employees = append(employees, employee)
	}

	return &employees, nil
}

func (e *Employees) UpdateEmployeesById(ctx context.Context, employees *[]models.Employee) (*[]models.Employee, error) {
	if employees == nil {
		return nil, errors.New("empty request")
	}

	var dalEmployees []employeesDalModel.Employee
	var idList []int
	for _, employee := range *employees {
		dalEmployee := &employeesDalModel.Employee{
			ID:           employee.ID,
			Username:     employee.Username,
			Email:        employee.Email,
			Age:          employee.Age,
			IsAdmin:      employee.IsAdmin,
			Dob:          employee.Dob,
			Details:      employee.Details,
			LastModified: primitive.NewDateTimeFromTime(time.Now()),
		}

		dalEmployees = append(dalEmployees, *dalEmployee)
		idList = append(idList, dalEmployee.ID)
	}

	updateData := common.ConvertStructToBson(dalEmployees)
	writeModels := []mongo.WriteModel{}

	database := e.Deps.Config.Get().Mongodb.Db
	collection := e.Deps.Config.Get().Mongodb.Collection

	var idFilter, filter bson.M

	for _, update := range updateData {
		idFilter = bson.M{"id": update["id"]}
		filter = bson.M{"$or": []bson.M{idFilter}}
		writeModels = append(writeModels, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(bson.M{"$set": update}))
	}

	err := e.Deps.Database.MongoDb.Update(ctx, database, collection, writeModels)
	if err != nil {
		return nil, err
	}

	readFilter := bson.M{"id": bson.M{"$in": idList}}

	cursor, err := e.Deps.Database.MongoDb.Read(ctx, database, collection, readFilter)
	if err != nil || cursor == nil {
		return nil, errors.New("error while getting updated employee")
	}

	var updatedEmployees []models.Employee
	for cursor.Next(ctx) {
		var updatedEmployee models.Employee
		err = cursor.Decode(&updatedEmployee)
		if err != nil {
			return nil, errors.New("error while mapping response")
		}
		updatedEmployees = append(updatedEmployees, updatedEmployee)
	}

	return &updatedEmployees, nil
}
