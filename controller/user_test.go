package controller

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Sanketkhote/microService/service/user"
	"github.com/Sanketkhote/microService/service/user/mocks"
	"github.com/Sanketkhote/microService/service/user/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func Test_controller_SaveUser(t *testing.T) {
	userData := model.UserModel{
		Name: "abc",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockuser := mocks.NewMockUser(ctrl)
	type fields struct {
		user user.User
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "200 ok",
			args: args{
				ctx: func() *gin.Context {
					w := httptest.NewRecorder()
					w.Write([]byte("Created"))
					w.WriteHeader(200)
					p, _ := gin.CreateTestContext(w)
					p.Request = httptest.NewRequest("POST", "/user", strings.NewReader(`{"Name":"abc"}`))
					return p
				}(),
			},
			fields: fields{
				user: func() *mocks.MockUser {
					mockuser.EXPECT().SaveUser(userData).Return(false, nil)
					return mockuser
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &controller{
				user: tt.fields.user,
			}
			c.SaveUser(tt.args.ctx)
		})
	}
}
