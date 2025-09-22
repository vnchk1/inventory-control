package server

import (
	"bytes"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	logging "github.com/vnchk1/inventory-control/internal/logger"
	"github.com/vnchk1/inventory-control/internal/mocks"
	"github.com/vnchk1/inventory-control/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

var logger = logging.NewLogger("debug")

func TestCategoryHandler_Create(t *testing.T) {
	tests := []struct {
		name       string
		reqBody    string
		mockSetup  func(mockService *mocks.MockCategoryUseCase)
		wantStatus int
	}{
		{
			name:    "success",
			reqBody: `{"name":"test category"}`,
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().
					Create(gomock.Any(), &models.Category{Name: "test category"}).
					Return(nil)
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "bad request - invalid json",
			reqBody:    `{"name":`, // сломан JSON
			mockSetup:  func(mockService *mocks.MockCategoryUseCase) {},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:    "internal server error - service failed",
			reqBody: `{"name":"fail category"}`,
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().
					Create(gomock.Any(), &models.Category{Name: "fail category"}).
					Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockCategoryUseCase(ctrl)

			e := echo.New()

			tt.mockSetup(mockService)

			req := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBufferString(tt.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := &CategoryHandler{
				Service: mockService,
				Logger:  logger,
			}

			err := h.Create(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestCategoryHandler_Update(t *testing.T) {
	tests := []struct {
		name       string
		reqBody    string
		mockSetup  func(mockService *mocks.MockCategoryUseCase)
		wantStatus int
	}{
		{
			name:    "success",
			reqBody: `{"id": 1, "name":"test category"}`,
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().Update(gomock.Any(), models.Category{ID: 1, Name: "test category"}).Return(nil)
			},
			wantStatus: http.StatusCreated,
		},
		{
			name:       "bad request - invalid json",
			reqBody:    `{"id": 1, "name":"test category"`,
			mockSetup:  func(mockService *mocks.MockCategoryUseCase) {},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:    "internal server error - service failed",
			reqBody: `{"id": 1, "name":"test category"}`,
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().Update(gomock.Any(), models.Category{ID: 1, Name: "test category"}).Return(errors.New("service error"))
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockCategoryUseCase(ctrl)

			e := echo.New()

			tt.mockSetup(mockService)

			req := httptest.NewRequest(http.MethodPut, "/categories", bytes.NewBufferString(tt.reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			h := &CategoryHandler{
				Service: mockService,
				Logger:  logger,
			}

			err := h.Update(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestCategoryHandler_Delete(t *testing.T) {
	tests := []struct {
		name       string
		reqId      string
		mockSetup  func(mockService *mocks.MockCategoryUseCase)
		wantStatus int
	}{
		{
			name:  "success",
			reqId: "1",
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
			},
			wantStatus: http.StatusNoContent,
		},
		{
			name:       "bad request - invalid id",
			reqId:      "bad",
			mockSetup:  func(mockService *mocks.MockCategoryUseCase) {},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:  "internal server error - not found",
			reqId: "1",
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(errors.New("not found"))
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockCategoryUseCase(ctrl)

			e := echo.New()

			tt.mockSetup(mockService)

			target := "/categories/" + tt.reqId
			req := httptest.NewRequest(http.MethodDelete, target, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.reqId)

			h := &CategoryHandler{
				Service: mockService,
				Logger:  logger,
			}

			err := h.Delete(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}

func TestCategoryHandler_Read(t *testing.T) {
	tests := []struct {
		name         string
		reqId        string
		mockSetup    func(mockService *mocks.MockCategoryUseCase)
		wantStatus   int
		wantCategory models.Category
	}{
		{
			name:  "success",
			reqId: "1",
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().Read(gomock.Any(), gomock.Any()).Return(models.Category{ID: 1, Name: "test category"}, nil)
			},
			wantStatus:   http.StatusOK,
			wantCategory: models.Category{ID: 1, Name: "test category"},
		},
		{
			name:       "bad request - invalid id",
			reqId:      "bad",
			mockSetup:  func(mockService *mocks.MockCategoryUseCase) {},
			wantStatus: http.StatusBadRequest,
		},
		{
			name:  "internal server error - not found",
			reqId: "1",
			mockSetup: func(mockService *mocks.MockCategoryUseCase) {
				mockService.EXPECT().Read(gomock.Any(), gomock.Any()).Return(models.Category{}, errors.New("not found"))
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockCategoryUseCase(ctrl)

			e := echo.New()

			tt.mockSetup(mockService)

			target := "/categories/" + tt.reqId
			req := httptest.NewRequest(http.MethodGet, target, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tt.reqId)

			h := &CategoryHandler{
				Service: mockService,
				Logger:  logger,
			}

			err := h.Read(c)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantStatus, rec.Code)

			if tt.name == "success" {
				assert.Equal(t, tt.wantCategory, models.Category{ID: 1, Name: "test category"})
			}
		})
	}
}
