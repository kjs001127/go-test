package pkg

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-test/gomock-mocks"
	"testing"
)

func TestWithGoMock(t *testing.T) {
	c := gomock.NewController(t)

	i := mocks.NewMockTestInterface(c)

	var capturedArg string

	gomock.InOrder(
		i.
			EXPECT().
			DoSomething(1, gomock.AssignableToTypeOf(capturedArg)). //cmdArg 에 대입 가능한 파라미터인지 검증 가능
			DoAndReturn(func(_ int, arg string) (int, error) {      //DoAndReturn 으로 함수 내용 전체를 모킹하기는 testify 보다 쉬움
				capturedArg = arg
				return 1, nil
			}),
		i.
			EXPECT().
			DoSomething(2, "xyz").
			AnyTimes(), //InOrder 와 조합되어 중간에 몇번 호출되어도 상관 없음!
		i.
			EXPECT().
			DoSomething(3, "cdf").
			Times(2), //두번 호출되지 않으면 assertion fail
	)

	i.DoSomething(1, "def")
	assert.Equal(t, "def", capturedArg)

	i.DoSomething(2, "xyz")
	i.DoSomething(2, "xyz")
	i.DoSomething(2, "xyz")

	i.DoSomething(3, "cdf")
	i.DoSomething(3, "cdf")
}
