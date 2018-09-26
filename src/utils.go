package src

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func RespondWithErr(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ArrayIntToArrayStr(arrayInt []int64) []string {
	arrayStr := []string{}
	for _, ele := range arrayInt {
		str := strconv.FormatInt(ele, 10)
		arrayStr = append(arrayStr, str)
	}
	return arrayStr
}
