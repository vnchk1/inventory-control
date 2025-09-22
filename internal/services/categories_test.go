package services

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vnchk1/inventory-control/internal/mocks"
	"github.com/vnchk1/inventory-control/internal/models"
	"strings"
	"testing"
)

func TestCategoryService_Create(t *testing.T) {
	tests := []struct {
		name      string
		category  *models.Category
		mockSetup func(m *mocks.MockCategoryRepo, t *testing.T)
		wantErr   error
	}{
		{
			name:     "Happy path",
			category: &models.Category{Name: "Books"},
			mockSetup: func(m *mocks.MockCategoryRepo, t *testing.T) {
				m.EXPECT().
					Create(gomock.Any(), gomock.AssignableToTypeOf(&models.Category{})).
					DoAndReturn(func(_ context.Context, c *models.Category) error {
						if c.Name != "Books" {
							t.Errorf("unexpected category name: %s", c.Name)
						}
						return nil
					})
			},
			wantErr: nil,
		},
		{
			name:     "Empty name",
			category: &models.Category{Name: ""},
			mockSetup: func(_ *mocks.MockCategoryRepo, _ *testing.T) {
				// Storage.Create не должен вызываться
			},
			wantErr: models.NewEmptyErr("name"),
		},
		{
			name:     "Too many items",
			category: &models.Category{Name: strings.Repeat("a", 101)},
			mockSetup: func(_ *mocks.MockCategoryRepo, _ *testing.T) {
				// Storage.Create не должен вызываться
			},
			wantErr: models.ErrTooManyItems,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStorage := mocks.NewMockCategoryRepo(ctrl)
			service := &CategoryService{Storage: mockStorage}

			if tt.mockSetup != nil {
				tt.mockSetup(mockStorage, t)
			}

			err := service.Create(context.Background(), tt.category)

			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func TestCategoryService_Update(t *testing.T) {
	tests := []struct {
		name      string
		category  models.Category
		mockSetup func(m *mocks.MockCategoryRepo, t *testing.T)
		wantErr   error
	}{
		{
			name:     "Happy path",
			category: models.Category{ID: 1, Name: "Books"},
			mockSetup: func(m *mocks.MockCategoryRepo, t *testing.T) {
				m.EXPECT().Update(gomock.Any(), models.Category{ID: 1, Name: "Books"}).Return(nil)
			},
			wantErr: nil,
		},
		{
			name:     "Zero ID",
			category: models.Category{ID: 0, Name: "Books"},
			mockSetup: func(_ *mocks.MockCategoryRepo, _ *testing.T) {
				// Storage.Create не должен вызываться
			},
			wantErr: models.NewNegativeErr("id"),
		},
		{
			name:     "Negative ID",
			category: models.Category{ID: -1, Name: "Books"},
			mockSetup: func(_ *mocks.MockCategoryRepo, _ *testing.T) {
				// Storage.Create не должен вызываться
			},
			wantErr: models.NewNegativeErr("id"),
		},
		{
			name:     "Empty name",
			category: models.Category{ID: 1, Name: ""},
			mockSetup: func(_ *mocks.MockCategoryRepo, _ *testing.T) {
				// Storage.Create не должен вызываться
			},
			wantErr: models.NewEmptyErr("name"),
		},
		{
			name:     "Too many items",
			category: models.Category{ID: 1, Name: strings.Repeat("a", 101)},
			mockSetup: func(_ *mocks.MockCategoryRepo, _ *testing.T) {
				// Storage.Create не должен вызываться
			},
			wantErr: models.ErrTooManyItems,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStorage := mocks.NewMockCategoryRepo(ctrl)
			service := &CategoryService{Storage: mockStorage}

			if tt.mockSetup != nil {
				tt.mockSetup(mockStorage, t)
			}

			err := service.Update(context.Background(), tt.category)

			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func TestCategoryService_Read(t *testing.T) {
	tests := []struct {
		name      string
		id        int
		mockSetup func(m *mocks.MockCategoryRepo, t *testing.T)
		want      models.Category
		wantErr   error
	}{
		{
			name: "Happy path",
			id:   1,
			mockSetup: func(m *mocks.MockCategoryRepo, t *testing.T) {
				m.EXPECT().
					Read(gomock.Any(), 1).
					Return(models.Category{ID: 1, Name: "Test category"}, nil)
			},
			want:    models.Category{ID: 1, Name: "Test category"},
			wantErr: nil,
		},
		{
			name:      "Zero ID",
			id:        0,
			mockSetup: func(m *mocks.MockCategoryRepo, t *testing.T) {},
			want:      models.Category{},
			wantErr:   models.NewNegativeErr("id"),
		},
		{
			name:      "Negative ID",
			id:        -1,
			mockSetup: func(m *mocks.MockCategoryRepo, t *testing.T) {},
			want:      models.Category{},
			wantErr:   models.NewNegativeErr("id"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockCategoryRepo(ctrl)
			service := NewCategoryService(mockRepo)

			tt.mockSetup(mockRepo, t)

			result, err := service.Read(context.Background(), tt.id)
			require.Equal(t, tt.want, result)

			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
