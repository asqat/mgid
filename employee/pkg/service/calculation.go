package service

import (
	"fmt"
	"github.com/asqat/mgid/employee/pkg/model"
	"sort"
)

var allowedFields = map[string]struct{}{
	"name":   {},
	"job":    {},
	"salary": {},
	"age":    {},
}

type Employees []model.Employee

//Sort sorts employees list by specified field key by ascending/descending order.
func (employees Employees) Sort(key string, asc bool) (Employees, error) {
	if len(employees) <= 1 {
		return employees, nil
	}

	if _, ok := allowedFields[key]; !ok {
		return nil, fmt.Errorf("%s is not allowed flield", key)
	}

	sort.Slice(employees, func(i, j int) bool {
		switch key {
		case "name":
			if asc {
				return employees[i].Name < employees[j].Name
			}
			return employees[i].Name > employees[j].Name
		case "job":
			if asc {
				return employees[i].Job < employees[j].Job
			}
			return employees[i].Job > employees[j].Job
		case "salary":
			if asc {
				return employees[i].Salary < employees[j].Salary
			}
			return employees[i].Salary > employees[j].Salary
		case "age":
			if asc {
				return employees[i].Age < employees[j].Age
			}
			return employees[i].Age > employees[j].Age
		default:
			if asc {
				return employees[i].Name < employees[j].Name
			}
			return employees[i].Name > employees[j].Name
		}
	})

	return employees, nil
}

//TheOldest returns the first element from sorted list by age field.
func (employees Employees) TheOldest() (*model.Employee, error) {
	if len(employees) == 1 {
		return &employees[0], nil
	} else if len(employees) == 0 {
		return &model.Employee{}, nil
	}
	all, err := employees.Sort("age", false)
	if err != nil {
		return nil, err
	}

	return &all[0], nil
}

//TheRichest returns the first element from sorted list by salary field.
func (employees Employees) TheRichest() (*model.Employee, error) {
	if len(employees) == 1 {
		return &employees[0], nil
	} else if len(employees) == 0 {
		return &model.Employee{}, nil
	}
	all, err := employees.Sort("salary", false)
	if err != nil {
		return nil, err
	}

	return &all[0], nil
}

//AverageSalary calculates an average of salary field value.
func (employees Employees) AverageSalary() float64 {
	if len(employees) == 1 {
		return employees[0].Salary
	} else if len(employees) == 0 {
		return 0
	}
	var salaries float64
	for _, employee := range employees {
		salaries += employee.Salary
	}

	return salaries / float64(len(employees))
}

//MedianSalary calculates a median of salary field value.
func (employees Employees) MedianSalary() float64 {
	if len(employees) == 1 {
		return employees[0].Salary
	} else if len(employees) == 0 {
		return 0
	}

	var sal []float64
	for _, emp := range employees {
		sal = append(sal, emp.Salary)
	}

	sort.Float64s(sal)

	n := len(sal)

	return (sal[n/2] + sal[(n-1)/2]) / 2
}
