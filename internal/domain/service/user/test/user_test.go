package test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
	"wordwiz/internal/domain/model"
	user2 "wordwiz/internal/domain/model/user"
	"wordwiz/internal/domain/service/user"
)

func TestService_TryCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	tests := []struct {
		name       string
		setupMocks func(r *MockRepo)
		userID     int
		wantBool   bool
		wantErr    bool
	}{
		{
			name: "User does not exist, created successfully",
			setupMocks: func(r *MockRepo) {
				r.EXPECT().GetByID(gomock.Any(), 1).Return(user2.User{}, model.ErrUserNotFound)
				r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil)
			},
			userID:   1,
			wantBool: true,
			wantErr:  false,
		},
		{
			name: "User already exists",
			setupMocks: func(r *MockRepo) {
				r.EXPECT().GetByID(gomock.Any(), 2).Return(user2.User{ID: 2}, nil)
			},
			userID:   2,
			wantBool: false,
			wantErr:  false,
		},
		{
			name: "Error in GetByID",
			setupMocks: func(r *MockRepo) {
				r.EXPECT().GetByID(gomock.Any(), 3).Return(user2.User{}, errors.New("db error"))
			},
			userID:   3,
			wantBool: false,
			wantErr:  true,
		},
		{
			name: "Error on Create",
			setupMocks: func(r *MockRepo) {
				r.EXPECT().GetByID(gomock.Any(), 4).Return(user2.User{}, model.ErrUserNotFound)
				r.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("insert error"))
			},
			userID:   4,
			wantBool: false,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockRepo := NewMockRepo(ctrl)
			tt.setupMocks(mockRepo)

			s := user.New(mockRepo)

			got, err := s.TryCreateUser(context.TODO(), user2.User{ID: tt.userID})
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tt.wantBool {
				t.Fatalf("unexpected result: got %v, want %v", got, tt.wantBool)
			}
		})
	}
}
