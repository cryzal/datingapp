package service

import (
	"datingapp/core/entities"
	port "datingapp/core/port/user"
	"datingapp/tests/mock/repository"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	type fields struct {
		UserRepositoryAdapter port.UserAdapter
	}
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.User
		wantErr bool
	}{
		{
			name:    "tes 1",
			fields:  fields{UserRepositoryAdapter: &repository.UserRepository{}},
			args:    args{email: "agung@gmail.com", password: "1234"},
			want:    &entities.User{ID: "1", Email: "agung@gmail.com", Password: "$2a$10$zC.TZJgzB514/AvQ9elfleLz7CcvTYx65UudQwfDKL5aXsOkyHRRq"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &UserService{
				UserRepositoryAdapter: tt.fields.UserRepositoryAdapter,
			}
			got, err := p.Login(tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.ID != tt.want.ID {
				t.Errorf("Email = %v, want %v", got.Email, tt.want.Email)
			}

			if got.Email != tt.want.Email {
				t.Errorf("Email = %v, want %v", got.Email, tt.want.Email)
			}

			if got.Password != tt.want.Password {
				t.Errorf("Password = %v, want %v", got.Email, tt.want.Email)
			}

			if got.Email != tt.want.Email {
				t.Errorf("Email = %v, want %v", got.Email, tt.want.Email)
			}

			if got.Token == "" {
				t.Errorf("Token can't be empty ")
			}

		})
	}
}

func TestUserService_Create(t *testing.T) {
	type fields struct {
		UserRepositoryAdapter port.UserAdapter
	}
	type args struct {
		user *entities.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "tes insert 1",
			fields:  fields{UserRepositoryAdapter: &repository.UserRepository{}},
			args:    args{user: &entities.User{Email: "agung@gmail.com", Password: "1234"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &UserService{
				UserRepositoryAdapter: tt.fields.UserRepositoryAdapter,
			}
			if err := p.Create(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
