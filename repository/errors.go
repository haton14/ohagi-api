package repository

import "errors"

var (
	ErrNotFoundRecord = errors.New("レコードが存在しない")
	ErrOthers         = errors.New("予期しないエラー")
	ErrDomainGenerate = errors.New("ドメイン生成エラー")
)
