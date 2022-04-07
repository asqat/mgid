package service

import (
	"github.com/asqat/mgid/employee/pkg/model"
	"reflect"
	"testing"
)

var testData = []model.Employee{
	{
		Name:   "Sara",
		Job:    "Developer",
		Salary: 3600,
		Age:    29,
	},
	{
		Name:   "Rick",
		Job:    "Engineer",
		Salary: 6600,
		Age:    41,
	},
	{
		Name:   "Robert",
		Job:    "Sales-manager",
		Salary: 2500,
		Age:    27,
	},
	{
		Name:   "Jasmin",
		Job:    "Manager",
		Salary: 3300,
		Age:    31,
	},
	{
		Name:   "Satoshi",
		Job:    "Developer",
		Salary: 3600,
		Age:    29,
	},
	{
		Name:   "Yuri",
		Job:    "Engineer",
		Salary: 5000,
		Age:    49,
	},
	{
		Name:   "Brian",
		Job:    "Student",
		Salary: 2500,
		Age:    27,
	},
	{
		Name:   "Lola",
		Job:    "Doctor",
		Salary: 3300,
		Age:    31,
	},
}

func TestEmployees_Sort(t *testing.T) {
	type args struct {
		key string
		asc bool
	}
	tests := []struct {
		name string
		e    Employees
		args args
		want Employees
	}{
		{
			name: "EmployeesNameAscSort_Test",
			e:    testData,
			args: args{
				key: "name",
				asc: true,
			},
			want: Employees{
				{
					Name:   "Brian",
					Job:    "Student",
					Salary: 2500,
					Age:    27,
				},
				{
					Name:   "Jasmin",
					Job:    "Manager",
					Salary: 3300,
					Age:    31,
				},
				{
					Name:   "Lola",
					Job:    "Doctor",
					Salary: 3300,
					Age:    31,
				},
				{
					Name:   "Rick",
					Job:    "Engineer",
					Salary: 6600,
					Age:    41,
				},
				{
					Name:   "Robert",
					Job:    "Sales-manager",
					Salary: 2500,
					Age:    27,
				},
				{
					Name:   "Sara",
					Job:    "Developer",
					Salary: 3600,
					Age:    29,
				},
				{
					Name:   "Satoshi",
					Job:    "Developer",
					Salary: 3600,
					Age:    29,
				},
				{
					Name:   "Yuri",
					Job:    "Engineer",
					Salary: 5000,
					Age:    49,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.e.Sort(tt.args.key, tt.args.asc); !reflect.DeepEqual(got, tt.want) {
				if err != nil {
					t.Errorf("testing error: %v", err)
				}
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployees_TheOldest(t *testing.T) {
	tests := []struct {
		name      string
		employees Employees
		want      *model.Employee
		wantErr   bool
	}{
		{
			name:      "TheOldest_Test",
			employees: testData,
			want: &model.Employee{
				Name:   "Yuri",
				Job:    "Engineer",
				Salary: 5000,
				Age:    49,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.employees.TheOldest()
			if (err != nil) != tt.wantErr {
				t.Errorf("TheOldest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TheOldest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployees_TheRichest(t *testing.T) {
	tests := []struct {
		name      string
		employees Employees
		want      *model.Employee
		wantErr   bool
	}{
		{
			name:      "TheRichest_Test",
			employees: testData,
			want: &model.Employee{
				Name:   "Rick",
				Job:    "Engineer",
				Salary: 6600,
				Age:    41,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.employees.TheRichest()
			if (err != nil) != tt.wantErr {
				t.Errorf("TheOldest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TheOldest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployees_AverageSalary(t *testing.T) {
	tests := []struct {
		name      string
		employees Employees
		want      float64
	}{
		{
			name:      "AverageSalary_Test",
			employees: testData,
			want:      3800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.employees.AverageSalary(); got != tt.want {
				t.Errorf("AverageSalary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployees_MedianSalary(t *testing.T) {
	tests := []struct {
		name      string
		employees Employees
		want      float64
	}{
		{
			name:      "MedianSalary_Test",
			employees: testData,
			want:      3450,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.employees.MedianSalary(); got != tt.want {
				t.Errorf("MedianSalary() = %v, want %v", got, tt.want)
			}
		})
	}
}
