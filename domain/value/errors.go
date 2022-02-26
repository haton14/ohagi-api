package value

import (
	"errors"
)

var (
	ErrMaxLength         = errors.New("最大長エラー")
	ErrMinLength         = errors.New("最小長エラー")
	ErrMaxRange          = errors.New("最大値エラー")
	ErrMinRange          = errors.New("最小値エラー")
	ErrRegularExpression = errors.New("正規表現エラー")
	ErrCharacterType     = errors.New("文字種エラー")
	ErrEnumValue         = errors.New("列挙値エラー")
	ErrOthers            = errors.New("その他エラー")
)
