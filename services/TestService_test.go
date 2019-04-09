package services

import "testing"

func TestTestService_Hello(t *testing.T) {
	type args struct {
		to string
	}
	tests := []struct {
		name string
		s    *TestService
		args args
		want string
	}{
		{
			name: "match name",
			s:    &TestService{},
			args: args{"maze"},
			want: "hi maze",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Hello(tt.args.to); got != tt.want {
				t.Errorf("TestService.Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTest2Service_Hello(t *testing.T) {
	type args struct {
		to string
	}
	tests := []struct {
		name string
		s    *Test2Service
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Hello(tt.args.to); got != tt.want {
				t.Errorf("Test2Service.Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
