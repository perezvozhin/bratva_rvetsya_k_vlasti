package usr_service

import "os"

// возврат значение для функции фетча резюме
// можно брать само резюме (файлом) или его контент - текстом
// смотри логику в checkreadCV()!!
type RetvalsInCV struct {
	file     *os.File
	contains string
}
