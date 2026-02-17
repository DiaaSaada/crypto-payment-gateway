package user_test

import (
	"testing"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/domain/user"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name         string
		username     string
		email        string
		passwordHash string
		wantErr      bool
		expectedErr  error
	}{
		{
			name:         "Valid user",
			username:     "testuser",
			email:        "test@example.com",
			passwordHash: "hashedpassword123",
			wantErr:      false,
		},
		{
			name:         "Empty username",
			username:     "",
			email:        "test@example.com",
			passwordHash: "hashedpassword123",
			wantErr:      true,
			expectedErr:  user.ErrEmptyUsername,
		},
		{
			name:         "Empty email",
			username:     "testuser",
			email:        "",
			passwordHash: "hashedpassword123",
			wantErr:      true,
			expectedErr:  user.ErrInvalidEmail,
		},
		{
			name:         "Empty password hash",
			username:     "testuser",
			email:        "test@example.com",
			passwordHash: "",
			wantErr:      true,
			expectedErr:  user.ErrInvalidPassword,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := user.NewUser(tt.username, tt.email, tt.passwordHash)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewUser() expected error but got none")
				}
				if tt.expectedErr != nil && err != tt.expectedErr {
					t.Errorf("NewUser() error = %v, expected %v", err, tt.expectedErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewUser() unexpected error = %v", err)
				return
			}

			if u.Username != tt.username {
				t.Errorf("Username = %v, want %v", u.Username, tt.username)
			}
			if u.Email != tt.email {
				t.Errorf("Email = %v, want %v", u.Email, tt.email)
			}
			if u.PasswordHash != tt.passwordHash {
				t.Errorf("PasswordHash = %v, want %v", u.PasswordHash, tt.passwordHash)
			}
		})
	}
}
