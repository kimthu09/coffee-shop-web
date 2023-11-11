package gincancelnote

//import (
//	"coffee_shop_management_backend/component/appctx"
//	"github.com/h2non/gock"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"gorm.io/gorm"
//	"io/ioutil"
//	"net/http"
//	"strings"
//	"testing"
//)
//
//type mockAppContext struct {
//	mock.Mock
//}
//
//func (m *mockAppContext) GetMainDBConnection() *gorm.DB {
//	args := m.Called()
//	if args.Get(0) == nil {
//		return nil
//	}
//	return args.Get(0).(*gorm.DB)
//}
//
//func (m *mockAppContext) GetSecretKey() string {
//	return ""
//}
//
//func TestCreateCancelNote(t *testing.T) {
//	defer gock.Off()
//
//	type args struct {
//		appCtx appctx.AppContext
//	}
//
//	mockCtx := new(mockAppContext)
//
//	tests := []struct {
//		name           string
//		args           args
//		requestBody    string
//		expectedStatus int
//		expectedBody   string
//	}{
//		{
//			name: "",
//			args: args{
//				appCtx: mockCtx,
//			},
//			requestBody: `{
//                "id": "123",
//                "details": [
//                    {
//                        "ingredientId": "nvl1",
//                        "expiryDate": "08/11/2003",
//                        "reason": "Damaged",
//                        "amountCancel": 10
//                    },
//                    {
//                        "ingredientId": "nvl2",
//                        "expiryDate": "08/12/2003",
//                        "reason": "Damaged",
//                        "amountCancel": 10
//                    }
//                ]
//            }`,
//			expectedStatus: http.StatusOK,
//			expectedBody:   `{"data": "123"}`,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			defer gock.Off()
//
//			gock.New("http://localhost:8080/v1/cancelNotes/").
//				MatchHeader("Authorization", "^foo bar$").
//				MatchHeader("API", "1.[0-9]+").
//				HeaderPresent("Accept").
//				Reply(tt.expectedStatus).
//				BodyString(tt.expectedBody)
//
//			req, _ := http.NewRequest(
//				"POST",
//				"http://localhost:8080/v1/cancelNotes/",
//				strings.NewReader(tt.requestBody))
//			req.Header.Set("Authorization", "foo bar")
//			req.Header.Set("API", "1.0")
//			req.Header.Set("Accept", "text/plain")
//
//			res, err := (&http.Client{}).Do(req)
//
//			assert.NotNil(t, err, "error: %v", err)
//			assert.Equal(t, res.StatusCode, tt.expectedStatus, "%v, %v", res.StatusCode, tt.expectedStatus)
//			body, _ := ioutil.ReadAll(res.Body)
//			assert.Equal(t, string(body), tt.expectedBody, "%v, %v", string(body), tt.expectedBody)
//
//			assert.Equal(t, gock.IsDone(), true)
//		})
//	}
//}
