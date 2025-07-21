package presenters

import (
	"fmt"
	"hello-clean/internal/domain/model"
	"hello-clean/internal/usecase/presenter"
)

type helloPresenter struct{}

func NewHelloPresenter() presenter.HelloPresenter {
	return &helloPresenter{}
}

// helloPresenter構造体がpresenter.HelloPresenterインターフェースを「実装」するために必要なメソッド
func (hp *helloPresenter) FormatHelloMessage(msg *model.HelloMessage) string {
	// 必要であれば、ここでJSON形式に変換したり、HTMLタグを追加可能
	return fmt.Sprintf("%s\n", msg.GetMessage())
}
