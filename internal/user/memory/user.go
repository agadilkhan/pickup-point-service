package memory

import (
	"context"
	"sync"
	"time"

	"github.com/agadilkhan/pickup-point-service/internal/user/entity"
	"github.com/agadilkhan/pickup-point-service/internal/user/repository"
	"go.uber.org/zap"
)

type UserMemory struct {
	logger *zap.SugaredLogger
	db     map[string]entity.User
	repo   repository.Repository
	sync.RWMutex
	interval time.Duration
}

func NewUserMemory(logger *zap.SugaredLogger, repo repository.Repository, interval time.Duration,
) *UserMemory {
	return &UserMemory{
		logger:   logger,
		repo:     repo,
		interval: interval,
	}
}

func (um *UserMemory) Run(ctx context.Context) {
	go func() {
		ticker := time.NewTimer(um.interval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				startTime := time.Now()

				um.LoadData(ctx)

				elapsedTime := time.Since(startTime)

				ticker.Reset(um.interval - elapsedTime)
			}
		}
	}()
}

func (um *UserMemory) LoadData(ctx context.Context) {
	um.Lock()
	defer um.Unlock()

	users, err := um.repo.GetUsers(ctx)
	if err != nil {
		um.logger.Errorf("failed to GetUsers err: %v", err)
		return
	}

	newUsers := make(map[string]entity.User)
	for _, user := range *users {
		newUsers[user.Login] = user
	}

	um.db = newUsers
}

func (um *UserMemory) GetUserByLogin(login string) *entity.User {
	um.Lock()
	defer um.Unlock()

	user, ok := um.db[login]
	if !ok {
		um.logger.Errorf("failed to get user by login from memory")
		return nil
	}

	um.logger.Info("get user by login from memory success")

	return &user
}
