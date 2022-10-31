package operation

import (
	"awesomeProject2/internal/repository/operation"
	"awesomeProject2/models"
	"fmt"
)

func CreateOperation(req models.Operation) (err error) {
	operation := operation.Operation(req)
	err = operation.CreateOperation()
	if err != nil {
		return fmt.Errorf("не получилось создать операцию: %w", err)
	}
	return nil
}
func GetOperations() (operations []*operation.Operation, err error) {
	var operation operation.Operation
	operations, err = operation.GetOperations()
	if err != nil {
		return nil, fmt.Errorf("не получилось получить операции: %w", err)
	}

	return operations, nil
}

func GetOperationByID(id string) (operation *operation.Operation, err error) {
	operation, err = operation.GetOperationByID(id)
	if err != nil {
		return nil, fmt.Errorf("не получилось получить операцию по данному айди: %w", err)
	}
	return operation, nil
}

func UpdateOperationByID(req models.Operation) (err error) {
	operation := operation.Operation(req)
	err = operation.UpdateOperationByID()
	if err != nil {
		return fmt.Errorf("не получилось обновить операцию: %w", err)
	}
	return nil
}

func DeleteOperationByID(id string) (err error) {
	var operation operation.Operation
	err = operation.DeleteOperationByID(id)
	if err != nil {
		return fmt.Errorf("не получилось удалить операцию: %w", err)
	}

	return nil
}
