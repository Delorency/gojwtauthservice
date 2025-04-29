package handlers

import (
	l "auth/internal/logger"
	"auth/internal/schemes"
	"auth/internal/tools"
	e "auth/internal/transport/http/error"
	"auth/internal/transport/http/response"
	"encoding/json"
	"io"
	"net/http"

	v "auth/internal/validator"

	"github.com/go-playground/validator"
)

func (ah *authHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := io.ReadAll(r.Body)
	var req schemes.RefreshRequest

	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		response.NewResponse(
			e.NewError(""),
			http.StatusInternalServerError,
			w,
		)
		ah.logger.Println(l.GetLogEntry(r, http.StatusInternalServerError, bodyBytes))
		return
	}
	if err := validator.New().Struct(req); err != nil {
		data, f := v.HandleValidationErrors(w, err)
		if f {
			response.NewResponse(
				data,
				http.StatusBadRequest,
				w,
			)
			ah.logger.Println(l.GetLogEntry(r, http.StatusBadRequest, bodyBytes))
			return
		}
	}
	ip, _ := tools.GetIp(r)

	req.Ip = ip
	req.UserAgent = r.UserAgent()
}
