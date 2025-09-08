package storage

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/vnchk1/inventory-control/internal/models"
	"github.com/vnchk1/inventory-control/testutils"
	"testing"
)

func TestCategoryStorage_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   *models.Category
		wantErr bool
		prepare func(ctx context.Context, db *DB)
		assert  func(ctx context.Context, t *testing.T, db *DB, category *models.Category)
	}{
		{
			name:    "Happy path",
			input:   &models.Category{Name: "Test Category"},
			wantErr: false,
			assert: func(ctx context.Context, t *testing.T, db *DB, category *models.Category) {
				require.NotZero(t, category.ID)

				var name string
				err := db.QueryRow(ctx, "SELECT category_name FROM categories WHERE category_id = $1", category.ID).Scan(&name)
				require.NoError(t, err)
				require.Equal(t, "Test Category", name)
			},
		},
		{
			name:    "Duplicate name",
			input:   &models.Category{Name: "Unique name"},
			wantErr: true,
			prepare: func(ctx context.Context, db *DB) {
				_, err := db.Exec(ctx, "INSERT INTO categories (category_name) VALUES ($1)", "Unique name")
				require.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			testDB, cleanup := testutils.SetupTestContainer(t)
			defer cleanup()

			if tt.prepare != nil {
				tt.prepare(ctx, testDB)
			}

			catStore := NewCategoryStorage(testDB)

			err := catStore.Create(ctx, tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if tt.assert != nil {
				tt.assert(ctx, t, testDB, tt.input)
			}
		})
	}
}

func TestCategoryStorage_Update(t *testing.T) {
	tests := []struct {
		name    string
		input   models.Category
		wantErr bool
	}{
		{
			name: "Happy path",
			input: models.Category{
				ID:   1,
				Name: "Test category 1"},
			wantErr: false,
		},
		{
			name: "Duplicate name",
			input: models.Category{
				ID:   1,
				Name: "Test category 1"},
			wantErr: true,
		},
		{
			name: "Not found",
			input: models.Category{
				ID:   3,
				Name: "Test category 1"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			testDB, cleanup := testutils.SetupTestContainer(t)
			defer cleanup()

			catStore := NewCategoryStorage(testDB)

			category := &models.Category{Name: "Test category"}

			err := catStore.Create(ctx, category)
			require.NoError(t, err)
			require.Equal(t, 1, category.ID)

			if tt.name == "Duplicate name" {
				category1 := &models.Category{Name: "Test category 1"}

				err = catStore.Create(ctx, category1)
				require.NoError(t, err)
				require.Equal(t, 2, category1.ID)
			}

			err = catStore.Update(ctx, tt.input)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

		})
	}
}

func TestCategoryStorage_Read(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		output  models.Category
		wantErr bool
	}{
		{
			name: "Happy path",
			id:   1,
			output: models.Category{
				ID:   1,
				Name: "Test category",
			},
			wantErr: false,
		},
		{
			name:    "Not found",
			id:      2,
			output:  models.Category{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			testDB, cleanup := testutils.SetupTestContainer(t)
			defer cleanup()

			catStore := NewCategoryStorage(testDB)

			category := &models.Category{Name: "Test category"}

			err := catStore.Create(ctx, category)
			require.NoError(t, err)
			require.Equal(t, 1, category.ID)

			result, err := catStore.Read(ctx, tt.id)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.output, result)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestCategoryStorage_Delete(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantErr bool
	}{
		{
			name:    "Happy path",
			id:      1,
			wantErr: false,
		},
		{
			name:    "Not found",
			id:      2,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			testDB, cleanup := testutils.SetupTestContainer(t)
			defer cleanup()

			catStore := NewCategoryStorage(testDB)

			category := &models.Category{Name: "Test category"}

			err := catStore.Create(ctx, category)
			require.NoError(t, err)
			require.Equal(t, 1, category.ID)

			err = catStore.Delete(ctx, tt.id)
			require.Equal(t, tt.wantErr, err != nil)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
