package router_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"ecommerce-platform/controllers/admin_controller"
// 	logger "ecommerce-platform/logger"
// 	"ecommerce-platform/models"
// 	"ecommerce-platform/router"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // MockValidationService represents a mock implementation of validation services
// type MockValidationService struct {
// 	mock.Mock
// }

// func (m *MockValidationService) ValidateReq(ctx *gin.Context, req interface{}) []string {
// 	args := m.Called(ctx, req)
// 	return args.Get(0).([]string)
// }

// func (m *MockValidationService) ValidateEmailPassword(fl validator.FieldLevel) bool {
// 	args := m.Called(fl)
// 	return args.Bool(0)
// }

// // MockAuthService represents a mock implementation of authentication services
// type MockAuthService struct {
// 	mock.Mock
// }

// func (m *MockAuthService) SignUp(ctx *gin.Context, user *models.Users) (*models.Users, string, error) {
// 	args := m.Called(ctx, user)
// 	return args.Get(0).(*models.Users), args.String(1), args.Error(2)
// }

// func (m *MockAuthService) ProcessLogin(ctx *gin.Context, req *models.LoginReq) (string, error) {
// 	args := m.Called(ctx, req)
// 	return args.String(0), args.Error(1)
// }

// func (m *MockAuthService) CheckUserExists(req *models.Users) (bool, error) {
// 	args := m.Called(req)
// 	return args.Bool(0), args.Error(1)
// }

// var (
// 	mockLoggerService   *logger.IAppLogger
// 	mockAuthService     *MockAuthService
// 	mockValidation      *MockValidationService
// 	MockAdminController *admin_controller.AdminControllers
// )

// // SetupTestRouter sets up the router for testing
// func SetupTestRouter() *router.Router {
// 	gin.SetMode(gin.TestMode)

// 	var loggerService logger.IAppLogger // Declare as interface, not pointer to interface
// 	mockAuthService = new(MockAuthService)
// 	mockValidation = new(MockValidationService)
// 	MockAdminController = new(admin_controller.AdminControllers)

// 	return router.NewRouter(
// 		loggerService, // Pass the interface directly
// 		mockAuthService,
// 		mockValidation,
// 		MockAdminController,
// 	)
// }

// func TestHealthCheck(t *testing.T) {
// 	r := SetupTestRouter()
// 	req, _ := http.NewRequest("GET", "/public/health-check", nil)
// 	w := httptest.NewRecorder()
// 	r.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response map[string]string
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "healthy", response["status"])
// 	assert.Equal(t, "API is running smoothly", response["message"])
// }

// // func TestSignUp(t *testing.T) {
// // 	r := SetupTestRouter()

// // 	reqBody := models.SignUpReq{
// // 		UserName:    "Ali",
// // 		Email:       "eeeeeee12@example.com",
// // 		Password:    "testpassword",
// // 		PhoneNumber: "+1234567890",
// // 		Address:     "123 Main St, Anytown, USA",
// // 	}

// // 	mockValidation.On("ValidateReq", mock.Anything, &reqBody).Return([]string{})
// // 	mockAuthService.On("SignUp", mock.Anything, &reqBody).Return(
// // 		&models.Users{
// // 			ID:          2,
// // 			UserName:    "Ali",
// // 			Email:       "eeeeeee12@example.com",
// // 			Password:    "$2a$14$kveS6PTokB7ihwfxel7K6eK27ygoeOOFTduciCJ5p.O2YbYhuifIq",
// // 			PhoneNumber: "+1234567890",
// // 			Role:        "user",
// // 			Address:     "123 Main St, Anytown, USA",
// // 			CreatedAt:   time.Date(2024, 12, 24, 17, 0, 27, 429709320, time.UTC),
// // 			UpdatedAt:   time.Time{},
// // 		},
// // 		200,
// // 		"Signed Up successfully",
// // 		"Signed Up successfully",
// // 	)

// // 	body, _ := json.Marshal(reqBody)
// // 	req, _ := http.NewRequest("POST", "/public/signup", bytes.NewBuffer(body))
// // 	req.Header.Set("Content-Type", "application/json")
// // 	w := httptest.NewRecorder()
// // 	r.Engine.ServeHTTP(w, req)

// // 	assert.Equal(t, http.StatusOK, w.Code)

// // 	// Update response struct to match the API response structure
// // 	var response struct {
// // 		Data struct {
// // 			ID           int       `json:"id"`
// // 			Name         string    `json:"name"`
// // 			Email        string    `json:"email"`
// // 			HashPassword string    `json:"hash_password"`
// // 			PhoneNumber  string    `json:"phone_number"`
// // 			Role         string    `json:"role"`
// // 			Address      string    `json:"address"`
// // 			CreatedAt    time.Time `json:"created_at"`
// // 			UpdatedAt    time.Time `json:"updated_at"`
// // 		} `json:"data"`
// // 		StatusCode int    `json:"status_code"`
// // 		Message    string `json:"message"`
// // 		SubMessage string `json:"sub_message"`
// // 	}
// // 	err := json.Unmarshal(w.Body.Bytes(), &response)
// // 	assert.NoError(t, err)
// // 	assert.Equal(t, 2, response.Data.ID)
// // 	assert.Equal(t, "Ali", response.Data.Name)
// // 	assert.Equal(t, "eeeeeee12@example.com", response.Data.Email)
// // 	assert.Equal(t, "$2a$14$kveS6PTokB7ihwfxel7K6eK27ygoeOOFTduciCJ5p.O2YbYhuifIq", response.Data.HashPassword)
// // 	assert.Equal(t, "+1234567890", response.Data.PhoneNumber)
// // 	assert.Equal(t, "user", response.Data.Role)
// // 	assert.Equal(t, "123 Main St, Anytown, USA", response.Data.Address)
// // 	assert.Equal(t, time.Date(2024, 12, 24, 17, 0, 27, 429709320, time.UTC), response.Data.CreatedAt)
// // 	assert.Equal(t, time.Time{}, response.Data.UpdatedAt)
// // 	assert.Equal(t, 200, response.StatusCode)
// // 	assert.Equal(t, "Signed Up successfully", response.Message)
// // 	assert.Equal(t, "Signed Up successfully", response.SubMessage)
// // }

// func TestLogin(t *testing.T) {
// 	r := SetupTestRouter()

// 	reqBody := models.LoginReq{
// 		Email:    "testuser@example.com",
// 		Password: "testpassword",
// 	}

// 	mockValidation.On("ValidateReq", mock.Anything, &reqBody).Return([]string{})
// 	mockAuthService.On("ProcessLogin", mock.Anything, &reqBody).Return("mocked_token", nil)

// 	body, _ := json.Marshal(reqBody)
// 	req, _ := http.NewRequest("POST", "/public/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	w := httptest.NewRecorder()
// 	r.Engine.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response models.TokenResponse
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "mocked_token", response.Token)
// 	assert.Equal(t, "Login successful", response.Message)
// }

// // func TestLoginInvalidCredentials(t *testing.T) {
// // 	r := SetupTestRouter()

// // 	reqBody := models.LoginReq{
// // 		Email:    "wronguser@example.com",
// // 		Password: "wrongpassword",
// // 	}

// // 	mockValidation.On("ValidateReq", mock.Anything, &reqBody).Return([]string{})
// // 	mockAuthService.On("ProcessLogin", mock.Anything, &reqBody).Return("", assert.AnError)

// // 	body, _ := json.Marshal(reqBody)
// // 	req, _ := http.NewRequest("POST", "/public/login", bytes.NewBuffer(body))
// // 	req.Header.Set("Content-Type", "application/json")
// // 	w := httptest.NewRecorder()
// // 	r.Engine.ServeHTTP(w, req)

// // 	// Check for correct status code
// // 	assert.Equal(t, http.StatusUnauthorized, w.Code)

// // 	// Unmarshal the response
// // 	var response map[string]interface{}
// // 	err := json.Unmarshal(w.Body.Bytes(), &response)
// // 	if err != nil {
// // 		t.Errorf("Failed to unmarshal response: %v", err)
// // 		return
// // 	}

// // 	// Check if the error message is returned for invalid credentials
// // 	if errorMsg, exists := response["error"]; exists {
// // 		// Expecting "user does not exist" message for invalid credentials
// // 		assert.Equal(t, "user does not exist", errorMsg)
// // 	} else {
// // 		// If no error, check that the response includes valid credentials response
// // 		assert.Equal(t, "", response["token"])
// // 		assert.Equal(t, float64(http.StatusUnauthorized), response["status_code"])
// // 		assert.Equal(t, "Invalid credentials", response["message"])
// // 	}
// // }

// // type TestRequest struct {
// // 	Email    string `validate:"alphanum"`
// // 	Password string `validate:"alphanum"`
// // }

// // func TestValidationServiceImpl_ValidateReq(t *testing.T) {
// // 	vs := &MockValidationService{}
// // 	validate := validator.New()

// // 	// Register the custom alphanum validation
// // 	validate.RegisterValidation("alphanum", func(fl validator.FieldLevel) bool {
// // 		return vs.ValidateEmailPassword(fl)
// // 	})

// // 	// Test cases
// // 	tests := []struct {
// // 		name       string
// // 		request    TestRequest
// // 		expectErr  bool
// // 		errorCount int
// // 	}{
// // 		{
// // 			name:       "Valid input",
// // 			request:    TestRequest{Email: "validEmail123", Password: "validPass456"},
// // 			expectErr:  false,
// // 			errorCount: 0,
// // 		},
// // 		{
// // 			name:       "Invalid email",
// // 			request:    TestRequest{Email: "invalid-email@", Password: "validPass456"},
// // 			expectErr:  true,
// // 			errorCount: 1,
// // 		},
// // 		{
// // 			name:       "Invalid password",
// // 			request:    TestRequest{Email: "validEmail123", Password: "invalid pass"},
// // 			expectErr:  true,
// // 			errorCount: 1,
// // 		},
// // 		{
// // 			name:       "Both invalid",
// // 			request:    TestRequest{Email: "invalid email@", Password: "invalid pass"},
// // 			expectErr:  true,
// // 			errorCount: 2,
// // 		},
// // 	}

// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			err := validate.Struct(tt.request)
// // 			if tt.expectErr {
// // 				assert.Error(t, err)
// // 				ValidationErrors := err.(validator.ValidationErrors)
// // 				assert.Equal(t, tt.errorCount, len(ValidationErrors))
// // 			} else {
// // 				assert.NoError(t, err)
// // 			}
// // 		})
// // 	}
// // }
