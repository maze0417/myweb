package services

import (
	"reflect"
	"testing"
)

func TestNewBrandService(t *testing.T) {
	type args struct {
		service ITestService
	}
	tests := []struct {
		name string
		args args
		want *BrandService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBrandService(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBrandService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBrandService_GetBrandIdAndStatusByCredentials(t *testing.T) {
	type args struct {
		clientId     string
		clientSecret string
	}
	tests := []struct {
		name       string
		s          *BrandService
		args       args
		wantId     string
		wantStatus int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BrandService{}
			gotId, gotStatus := s.GetBrandIdAndStatusByCredentials(tt.args.clientId, tt.args.clientSecret)
			if gotId != tt.wantId {
				t.Errorf("BrandService.GetBrandIdAndStatusByCredentials() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("BrandService.GetBrandIdAndStatusByCredentials() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
