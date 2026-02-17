package user_test

import (
	"context"
	"testing"

	"github.com/DiaaSaada/crypto-payment-gateway/internal/domain/user"
	userRepo "github.com/DiaaSaada/crypto-payment-gateway/internal/repository/user"
)

func TestInMemoryRepository_Create(t *testing.T) {
	repo := userRepo.NewInMemoryRepository()
	ctx := context.Background()

	u, _ := user.NewUser("testuser", "test@example.com", "hashedpassword")

	err := repo.Create(ctx, u)
	if err != nil {
		t.Errorf("Create() unexpected error = %v", err)
	}

	if u.ID == "" {
		t.Error("Create() should generate ID for user")
	}
}

func TestInMemoryRepository_CreateDuplicateEmail(t *testing.T) {
	repo := userRepo.NewInMemoryRepository()
	ctx := context.Background()

	u1, _ := user.NewUser("testuser1", "test@example.com", "hashedpassword")
	u2, _ := user.NewUser("testuser2", "test@example.com", "hashedpassword")

	_ = repo.Create(ctx, u1)
	err := repo.Create(ctx, u2)

	if err != userRepo.ErrUserAlreadyExists {
		t.Errorf("Create() expected ErrUserAlreadyExists, got %v", err)
	}
}

func TestInMemoryRepository_FindByEmail(t *testing.T) {
	repo := userRepo.NewInMemoryRepository()
	ctx := context.Background()

	u, _ := user.NewUser("testuser", "test@example.com", "hashedpassword")
	_ = repo.Create(ctx, u)

	found, err := repo.FindByEmail(ctx, "test@example.com")
	if err != nil {
		t.Errorf("FindByEmail() unexpected error = %v", err)
	}

	if found.Email != u.Email {
		t.Errorf("FindByEmail() email = %v, want %v", found.Email, u.Email)
	}
}

func TestInMemoryRepository_FindByEmailNotFound(t *testing.T) {
	repo := userRepo.NewInMemoryRepository()
	ctx := context.Background()

	_, err := repo.FindByEmail(ctx, "notfound@example.com")

	if err != userRepo.ErrUserNotFound {
		t.Errorf("FindByEmail() expected ErrUserNotFound, got %v", err)
	}
}

func TestInMemoryRepository_FindByID(t *testing.T) {
	repo := userRepo.NewInMemoryRepository()
	ctx := context.Background()

	u, _ := user.NewUser("testuser", "test@example.com", "hashedpassword")
	_ = repo.Create(ctx, u)

	found, err := repo.FindByID(ctx, u.ID)
	if err != nil {
		t.Errorf("FindByID() unexpected error = %v", err)
	}

	if found.ID != u.ID {
		t.Errorf("FindByID() ID = %v, want %v", found.ID, u.ID)
	}
}

func TestInMemoryRepository_Update(t *testing.T) {
	repo := userRepo.NewInMemoryRepository()
	ctx := context.Background()

	u, _ := user.NewUser("testuser", "test@example.com", "hashedpassword")
	_ = repo.Create(ctx, u)

	u.Username = "updateduser"
	err := repo.Update(ctx, u)
	if err != nil {
		t.Errorf("Update() unexpected error = %v", err)
	}

	found, _ := repo.FindByID(ctx, u.ID)
	if found.Username != "updateduser" {
		t.Errorf("Update() username = %v, want %v", found.Username, "updateduser")
	}
}
