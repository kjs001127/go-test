package pkg

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-test/mocks/pkg"
	"testing"
)

func TestWithTestify(t *testing.T) {
	i := mocks.NewTestInterface(t)

	var capturedArg string

	i.
		EXPECT().
		DoSomething(1, "abc").
		Maybe() //호출되지 않아도 괜찮음. gomock 의 AnyTimes 와 같다
	i.
		EXPECT().
		DoSomething(2, "cdf").
		Run(func(_ int, arg string) {
			capturedArg = arg //이런식으로 argument capture 해야할듯
		}).
		Return(3, nil) //Run 이 있으면 꼭 Return 도 필요한듯, Run, Return 모두 type-safe 함

	i.DoSomething(2, "cdf")
	assert.Equal(t, "cdf", capturedArg)

	i.
		EXPECT().
		DoSomething(3, mock.Anything).
		Call.
		Return(func(arg int, arg0 string) int { return arg }, nil) //GoMock 의 DoAndReturn 처럼 쓰려면 조금 불편!

	res, _ := i.DoSomething(3, "efg")
	assert.Equal(t, 3, res)

}
