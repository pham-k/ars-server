package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func WriteJson[T any](writer http.ResponseWriter, status int, data T) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)

	response, err := json.Marshal(data)
	if err != nil {
		message := fmt.Sprintf("marshal api %v", data)
		slog.Error(message, err)
	}

	_, err = writer.Write(response)
	if err != nil {
		message := fmt.Sprintf("encode response %v", response)
		slog.Error(message, err)
	}
}

func ReadJson(request *http.Request, destination interface{}) error {
	err := json.NewDecoder(request.Body).Decode(destination)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		default:
			return err
		}
	}
	return nil
}
