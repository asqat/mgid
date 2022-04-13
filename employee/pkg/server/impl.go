package server

import (
	"context"
	employee "github.com/asqat/mgid/employee/pkg/proto"
	"github.com/asqat/mgid/employee/pkg/service"
	"math/rand"
	"time"
)

type server struct {
	employee.UnimplementedEmployeeServiceServer
	data service.Employees
}

func (s *server) EmployeesSort(ctx context.Context, filter *employee.Filter) (*employee.Employees, error) {
	list, err := s.data.Sort(filter.Field, filter.Asc)
	if err != nil {
		return nil, err
	}
	epms := &employee.Employees{}
	for _, item := range list {
		emp := &employee.Employee{}
		emp.Name = item.Name
		emp.Job = item.Job
		emp.Age = uint64(item.Age)
		emp.Salary = item.Salary
		epms.Employees = append(epms.Employees, emp)
	}

	if filter.Imitator.IsLongLoad {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}

	return epms, nil
}

func (s *server) TheOldest(ctx context.Context, imitator *employee.Imitator) (*employee.Employee, error) {
	oldest, err := s.data.TheOldest()
	if err != nil {
		return nil, err
	}

	emp := &employee.Employee{
		Name:   oldest.Name,
		Job:    oldest.Job,
		Salary: oldest.Salary,
		Age:    uint64(oldest.Age),
	}

	if imitator.GetIsLongLoad() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
	return emp, nil
}
func (s *server) TheRichest(ctx context.Context, imitator *employee.Imitator) (*employee.Employee, error) {
	richest, err := s.data.TheRichest()
	if err != nil {
		return nil, err
	}

	emp := &employee.Employee{
		Name:   richest.Name,
		Job:    richest.Job,
		Salary: richest.Salary,
		Age:    uint64(richest.Age),
	}

	if imitator.GetIsLongLoad() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
	return emp, nil
}
func (s *server) MeanSalary(ctx context.Context, imitator *employee.Imitator) (*employee.Salary, error) {
	mean := s.data.AverageSalary()
	salary := &employee.Salary{
		Value: mean,
	}

	if imitator.GetIsLongLoad() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
	return salary, nil
}
func (s *server) MedianSalary(ctx context.Context, imitator *employee.Imitator) (*employee.Salary, error) {
	mean := s.data.MedianSalary()
	salary := &employee.Salary{
		Value: mean,
	}

	if imitator.GetIsLongLoad() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}

	return salary, nil
}
