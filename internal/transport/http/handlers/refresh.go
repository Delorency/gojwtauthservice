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
	ip, _ := tools.GetIp(r)

	req.Ip = ip
}
