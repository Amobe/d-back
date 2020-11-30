package handler_test

import (
	"net/http"
	"testing"

	"github.com/amobe/d-back/pkg/infra/webserver"

	"github.com/amobe/d-back/pkg/handler"
	"github.com/go-playground/assert/v2"

	"github.com/amobe/d-back/pkg/entity"

	mock_iplimiter "github.com/amobe/d-back/pkg/service/iplimiter/mock"
	"github.com/golang/mock/gomock"
)

func TestHelloHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := http.Request{
		RemoteAddr: "127.0.0.1:65535",
	}

	ipAddress, _ := entity.NewIPAddress("127.0.0.1:65535")
	token := entity.NewRequestToken(1)

	ils := mock_iplimiter.NewMockService(ctrl)
	ils.EXPECT().AcceptRequest(ipAddress).Return(token, nil)

	// given a hello handler
	h := handler.NewHelloHandler(ils)

	// when execute the handle function
	got := h.Handle(&req)

	// then handler should return correct response
	assert.Equal(t, webserver.NewResponse(http.StatusOK, handler.HelloResponse{Index: 1}), got)
}
