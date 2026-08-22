package usr_service

import "os"

// возврат значение для функции фетча резюме
// смотри логику в checkreadCV()!!
type RetvalsInCV struct {
	file     *os.File
	contains string
}
