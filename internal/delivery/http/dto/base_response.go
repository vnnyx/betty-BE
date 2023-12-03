package dto

import "github.com/vnnyx/betty-BE/internal/enums"

type Photo struct {
	ID       int64  `json:"id"`
	PhotoURL string `json:"photo_url"`
}

type WebResponse struct {
	Code   int                 `json:"code"`
	Status enums.MessageStatus `json:"status"`
	Data   any                 `json:"data"`
	Error  any                 `json:"error"`
}
