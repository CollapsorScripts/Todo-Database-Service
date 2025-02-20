package utilities

import (
	"bytes"
	"databaseService/pkg/logger"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pchchv/captcha"
	"image/color"
	"image/jpeg"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// encode - кодирует  в base64
func encode(bin []byte) []byte {
	e64 := base64.StdEncoding

	maxEncLen := e64.EncodedLen(len(bin))
	encBuf := make([]byte, maxEncLen)

	e64.Encode(encBuf, bin)
	return encBuf
}

// Exists - проверяет существует ли файл
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func format(enc []byte, mime string) string {
	switch mime {
	case "image/gif", "image/jpeg", "image/pjpeg", "image/png", "image/tiff":
		return fmt.Sprintf("data:%s;charset=utf-8;base64,%s", mime, enc)
	default:
	}

	return fmt.Sprintf("data:image/png;base64,%s", enc)
}

// executeExt - извлекает расширение файла
func executeExt(str string) string {
	switch strings.TrimSuffix(str[5:strings.Index(str, ",")], ";base64") {
	case "image/png":
		return "png"
	case "image/jpeg":
		return "jpg"
	default:
		return ""
	}
}

// trimBase64 - отрезает лишнее и возвращает только base64
func trimBase64(str string) string {
	base64str := str[strings.Index(str, ",")+1:]
	return base64str
}

// decodeBase64 - декодирует base64 и возвращает []byte
func decodeBase64(b64 string) []byte {
	b, _ := base64.StdEncoding.DecodeString(b64)
	return b
}

// fromBuffer принимает набор байтов в буфере
// возвращает base64 строку.
func fromBuffer(buf bytes.Buffer) string {
	enc := encode(buf.Bytes())
	mime := http.DetectContentType(buf.Bytes())
	logger.Info("mime: %s", mime)

	return format(enc, mime)
}

// fileFromLocal - файл с локального хранилища
func fileFromLocal(fname string) (string, error) {
	var b bytes.Buffer

	fileExists, _ := exists(fname)
	if !fileExists {
		return "", fmt.Errorf("File does not exist\n")
	}

	file, err := os.Open(fname)
	if err != nil {
		return "", fmt.Errorf("Error opening file\n")
	}

	_, err = b.ReadFrom(file)
	if err != nil {
		return "", fmt.Errorf("Error reading file to buffer\n")
	}

	return fromBuffer(b), nil
}

// FileToBase64 - конвертирует файл в base64
func FileToBase64(filepath string) string {
	result, err := fileFromLocal(filepath)
	if err != nil {
		logger.Error("Ошибка при попытке кодирования файла в base64: %s", err.Error())
		result = ""
	}

	return result
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "   ")
	if err != nil {
		return in
	}
	return out.String()
}

// ToBytesJSON - конвертирует объект в JSON байт
func ToBytesJSON(object any) []byte {
	jsonByte, err := json.Marshal(object)
	if err != nil {
		logger.Error("Ошибка при получении JSON: ", err.Error())
	}
	n := len(jsonByte)             //Find the length of the byte array
	result := string(jsonByte[:n]) //convert to string

	return []byte(jsonPrettyPrint(result))
}

// ToJSON - конвертирует объект в JSON строку
func ToJSON(object any) string {
	jsonByte, err := json.Marshal(object)
	if err != nil {
		logger.Error("Ошибка при получении JSON: ", err.Error())
	}
	n := len(jsonByte)             //Find the length of the byte array
	result := string(jsonByte[:n]) //convert to string

	return jsonPrettyPrint(result)
}

// Compare - сравнивает строку с зашифрованной строкой
func Compare(str, cryptStr string) bool {
	logger.Info("crypted password: %s", MD5(str))
	b := MD5(str) == cryptStr
	logger.Info("сравнение: %t", b)
	return MD5(str) == cryptStr
}

// StrToUint - Конвертирует строку в uint
func StrToUint(s string) uint {
	i, err := strconv.Atoi(s)
	if err != nil {
		logger.Error("%s", err.Error())
		return 0
	}
	return uint(i)
}

// RandInt - возвращает случайное число от min до max
func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// GenerateRandomString - генерирует случайный набор символов (англ алфавит, case uppercase + символ _ и цифры от 0 до 9)
func GenerateRandomString(length int) string {
	alphabet := "QOS4rT08Dm7dZVOPwucfM2haFiNyEjBK3UtC9IqY_lzv6XpWgWsAJebG5H1RxnLbK"

	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(result)
}

// GenerateCaptcha - генерирует капчу и возвращает путь к файлу
func GenerateCaptcha() (*captcha.Captcha, string, error) {
	dataCaptcha, err := captcha.New(450, 150, func(options *captcha.Options) {
		options.BackgroundColor = color.White
		options.Noise = 2
	})
	if err != nil {
		return nil, "", err
	}

	dirPath, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	fileName := fmt.Sprintf("%s.jpg", GenerateRandomString(12))
	pathToImg := filepath.Join(dirPath, fileName)
	file, err := os.Create(pathToImg)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	err = dataCaptcha.WriteJPG(file, &jpeg.Options{Quality: 50})
	if err != nil {
		return nil, "", err
	}

	return dataCaptcha, pathToImg, nil
}

// ChangeEnvAttribute - заменяет значение в attr на prop в .env файле
func ChangeEnvAttribute(attr, prop string) error {
	// Путь к файлу .env
	envFile := ".env"

	// Чтение файла .env
	data, err := os.ReadFile(envFile)
	if err != nil {
		return err
	}

	attrFormat := fmt.Sprintf("%s=", attr)
	newAttrProp := fmt.Sprintf("%s=%s", attr, prop)

	// Разделение файла .env на строки
	lines := strings.Split(string(data), "\n")

	// Флаг для определения замены
	replaced := false

	// Итерация по строкам файла .env
	for i, line := range lines {
		// Проверка, содержит ли строка "attrFormat"
		if strings.HasPrefix(line, attrFormat) {
			// Замена значения attr
			lines[i] = newAttrProp
			replaced = true
			break
		}
	}

	// Если значение BotToken не было заменено, добавьте его в конец файла
	if !replaced {
		lines = append(lines, newAttrProp)
	}

	// Запись измененных строк обратно в файл .env
	err = os.WriteFile(envFile, []byte(strings.Join(lines, "\n")), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// Percentage - рассчитывает процент от числа
func Percentage(percent, num float64) float64 {
	return (percent / 100) * num
}

// GetCryptoKey - создает ключ шифрования и возвращает его хэш
func GetCryptoKey() string {
	//TODO: тут был ключ
	apiToken := strings.Split("", ":")[1]
	hash := MD5(apiToken)

	if len(hash) < 32 {
		for len(hash) < 32 {
			hash += "1"
		}
	}

	return hash[:32]
}

// Transformation - преобразование одной модели в другую
func Transformation(forModel any, toModel any) error {
	encodedJsonModelBytes, err := json.Marshal(forModel)
	if err != nil {
		return err
	}

	err = json.Unmarshal(encodedJsonModelBytes, toModel)
	if err != nil {
		return err
	}

	return nil
}
