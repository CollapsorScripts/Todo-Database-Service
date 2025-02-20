package config

import (
	"bytes"
	"encoding/json"
	"testing"
)

const configPathLocal = "../../config/local.yaml"

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "   ")
	if err != nil {
		return in
	}
	return out.String()
}

// ToJSON - конвертирует объект в JSON строку
func ToJSON(object any) string {
	jsonByte, err := json.Marshal(object)
	if err != nil {
		panic(any(err))
	}
	n := len(jsonByte)             //Find the length of the byte array
	result := string(jsonByte[:n]) //convert to string

	return jsonPrettyPrint(result)
}

func Test_MustLoadByPath_Happy(t *testing.T) {
	cfg := MustLoadByPath("../../config/local.yaml")

	t.Logf("Конфигурация local: %s", ToJSON(cfg))

	/*
		cfg = MustLoadByPath("../../config/prod.yaml")

		t.Logf("Конфигурация prod: %s", ToJSON(cfg))

	*/
}

func Test_MustLoadByPath_Bad(t *testing.T) {
	cfg := MustLoadByPath("local.yaml")

	t.Logf("Конфигурация: %+v", cfg)
}
