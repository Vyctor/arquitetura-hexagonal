package handler

import "encoding/json"

func jsonError(message string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		message,
	}

	result, err := json.Marshal(error)

	if err != nil {
		return []byte(err.Error())
	}

	return result
}
