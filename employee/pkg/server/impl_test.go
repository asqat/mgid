package server

import (
	"context"
	employee "github.com/asqat/mgid/employee/pkg/proto"
	"github.com/asqat/mgid/employee/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	grpcServer := &server{
		data: service.Employees{
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
		},
	}

	server := grpc.NewServer()

	employee.RegisterEmployeeServiceServer(server, grpcServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func Test_server_EmployeesSort(t *testing.T) {
	type fields struct {
		UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
	}
	type args struct {
		in0 context.Context
		in1 *employee.Filter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *employee.Employees
		wantErr bool
	}{
		{
			name: "EmployeeSort_Test_Asc",
			fields: struct {
				UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
			}{UnimplementedEmployeeServiceServer: employee.UnimplementedEmployeeServiceServer{}},
			args: args{
				in0: context.Background(),
				in1: &employee.Filter{
					Field: "name",
					Asc:   true,
					Imitator: &employee.Imitator{
						IsLongLoad: false,
					},
				},
			},
			want: &employee.Employees{
				Employees: []*employee.Employee{
					{
						Name:   "Brian",
						Job:    "Student",
						Salary: 2500,
						Age:    27,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "EmployeeSort_Test_Asc_LongLoading",
			fields: struct {
				UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
			}{UnimplementedEmployeeServiceServer: employee.UnimplementedEmployeeServiceServer{}},
			args: args{
				in0: context.Background(),
				in1: &employee.Filter{
					Field: "name",
					Asc:   true,
					Imitator: &employee.Imitator{
						IsLongLoad: true,
					},
				},
			},
			want: &employee.Employees{
				Employees: []*employee.Employee{
					{
						Name:   "Brian",
						Job:    "Student",
						Salary: 2500,
						Age:    27,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "EmployeeSort_Test_Desc_LongLoading",
			fields: struct {
				UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
			}{UnimplementedEmployeeServiceServer: employee.UnimplementedEmployeeServiceServer{}},
			args: args{
				in0: context.Background(),
				in1: &employee.Filter{
					Field: "name",
					Asc:   false,
					Imitator: &employee.Imitator{
						IsLongLoad: true,
					},
				},
			},
			want: &employee.Employees{
				Employees: []*employee.Employee{
					{
						Name:   "Yuri",
						Job:    "Engineer",
						Salary: 5000,
						Age:    49,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn, err := grpc.DialContext(tt.args.in0, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
			if (err != nil) != tt.wantErr {
				t.Errorf("EmployeesSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			defer conn.Close()

			cli := employee.NewEmployeeServiceClient(conn)

			resp, err := cli.EmployeesSort(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmployeesSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Employees[0].Name != tt.want.Employees[0].Name {
				t.Errorf("EmployeesSort() got = %v, want %v", resp.Employees[0].Name, tt.want.Employees[0].Name)
			}
		})
	}
}

func Test_server_TheOldest(t *testing.T) {
	type fields struct {
		UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
	}
	type args struct {
		in0 context.Context
		emp *employee.Imitator
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *employee.Employees
		wantErr bool
	}{
		{
			name: "TheOldestEmployee_Test",
			fields: struct {
				UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
			}{UnimplementedEmployeeServiceServer: employee.UnimplementedEmployeeServiceServer{}},
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: false,
				},
			},
			want: &employee.Employees{
				Employees: []*employee.Employee{
					{
						Name:   "Yuri",
						Job:    "Engineer",
						Salary: 5000,
						Age:    49,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "TheOldestEmployee_Test_LongLoading",
			fields: struct {
				UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
			}{UnimplementedEmployeeServiceServer: employee.UnimplementedEmployeeServiceServer{}},
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: true,
				},
			},
			want: &employee.Employees{
				Employees: []*employee.Employee{
					{
						Name:   "Yuri",
						Job:    "Engineer",
						Salary: 5000,
						Age:    49,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn, err := grpc.DialContext(tt.args.in0, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
			if (err != nil) != tt.wantErr {
				t.Errorf("TheOldest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			defer conn.Close()

			cli := employee.NewEmployeeServiceClient(conn)

			resp, err := cli.TheOldest(tt.args.in0, tt.args.emp)
			if (err != nil) != tt.wantErr {
				t.Errorf("TheOldest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Name != tt.want.Employees[0].Name {
				t.Errorf("TheOldest() got = %s (%d years), want %v(%d years)",
					resp.Name, resp.Age, tt.want.Employees[0].Name, tt.want.Employees[0].Age)
			}
		})
	}
}

func Test_server_TheRichest(t *testing.T) {
	type fields struct {
		UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
	}
	type args struct {
		in0 context.Context
		emp *employee.Imitator
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *employee.Employees
		wantErr bool
	}{
		{
			name: "TheRichestEmployee_Test",
			fields: struct {
				UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
			}{UnimplementedEmployeeServiceServer: employee.UnimplementedEmployeeServiceServer{}},
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: false,
				},
			},
			want: &employee.Employees{
				Employees: []*employee.Employee{
					{
						Name:   "Rick",
						Job:    "Engineer",
						Salary: 6600,
						Age:    41,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "TheRichestEmployee_Test_LongLoading",
			fields: struct {
				UnimplementedEmployeeServiceServer employee.UnimplementedEmployeeServiceServer
			}{UnimplementedEmployeeServiceServer: employee.UnimplementedEmployeeServiceServer{}},
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: true,
				},
			},
			want: &employee.Employees{
				Employees: []*employee.Employee{
					{
						Name:   "Rick",
						Job:    "Engineer",
						Salary: 6600,
						Age:    41,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn, err := grpc.DialContext(tt.args.in0, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
			if (err != nil) != tt.wantErr {
				t.Errorf("TheRichest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			defer conn.Close()

			cli := employee.NewEmployeeServiceClient(conn)

			resp, err := cli.TheRichest(tt.args.in0, tt.args.emp)
			if (err != nil) != tt.wantErr {
				t.Errorf("TheRichest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Name != tt.want.Employees[0].Name {
				t.Errorf("TheRichest() got = %s (%v $), want %s(%v $)",
					resp.Name, resp.Salary, tt.want.Employees[0].Name, tt.want.Employees[0].Salary)
			}
		})
	}
}

func Test_server_MeanSalary(t *testing.T) {
	type args struct {
		in0 context.Context
		emp *employee.Imitator
	}
	tests := []struct {
		name    string
		args    args
		want    *employee.Salary
		wantErr bool
	}{
		{
			name: "MeanSalary_Test",
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: false,
				},
			},
			want: &employee.Salary{
				Value: 3800,
			},
			wantErr: false,
		},
		{
			name: "MeanSalary_Test_LongLoading",
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: true,
				},
			},
			want: &employee.Salary{
				Value: 3800,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn, err := grpc.DialContext(tt.args.in0, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
			if (err != nil) != tt.wantErr {
				t.Errorf("MeanSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			defer conn.Close()

			cli := employee.NewEmployeeServiceClient(conn)

			resp, err := cli.MeanSalary(tt.args.in0, tt.args.emp)
			if (err != nil) != tt.wantErr {
				t.Errorf("MeanSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Value != tt.want.Value {
				t.Errorf("MeanSalary() got = %v $, want %v $", resp.Value, tt.want.Value)
			}
		})
	}
}

func Test_server_MedianSalary(t *testing.T) {
	type args struct {
		in0 context.Context
		emp *employee.Imitator
	}
	tests := []struct {
		name    string
		args    args
		want    *employee.Salary
		wantErr bool
	}{
		{
			name: "MedianSalary_Test",
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: false,
				},
			},
			want: &employee.Salary{
				Value: 3450,
			},
			wantErr: false,
		},
		{
			name: "MedianSalary_Test_LongLoading",
			args: args{
				in0: context.Background(),
				emp: &employee.Imitator{
					IsLongLoad: true,
				},
			},
			want: &employee.Salary{
				Value: 3450,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn, err := grpc.DialContext(tt.args.in0, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
			if (err != nil) != tt.wantErr {
				t.Errorf("MeanSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			defer conn.Close()

			cli := employee.NewEmployeeServiceClient(conn)

			resp, err := cli.MedianSalary(tt.args.in0, tt.args.emp)
			if (err != nil) != tt.wantErr {
				t.Errorf("MedianSalary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if resp.Value != tt.want.Value {
				t.Errorf("MedianSalary() got = %v $, want %v $", resp.Value, tt.want.Value)
			}
		})
	}
}
