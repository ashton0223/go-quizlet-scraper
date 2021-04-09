// I have no idea how I got this to work, but it does somehow.
// The C code at the beginning allocates memory for strings and
// stores the pointer in memorry allocated by C.malloc.
// This allows for it to be used after the call to this
// library in a C program.
package main

// #include <string.h>
// #include <stdlib.h>
// #include <stdio.h>
// static void store_ptr(char** str_arr, char* str, int offset, int length) {
//	 //printf("%s", str);
//	 char* newstring;
//	 newstring = (char*)malloc(length*sizeof(char) + 1);
//	 memset(newstring, '\0', sizeof(newstring));
//	 strcpy(newstring, str);
//	 *(str_arr + (offset)) = newstring;
// }
import "C"

import (
	"unsafe"

	"github.com/ashton0223/go-quizlet-scraper/scraper"
)

//export GetStudySetC
func GetStudySetC(url *C.char) (**C.char, **C.char, C.int, *C.char) {
	urlgo := C.GoString(url)
	termArr, defArr, err := scraper.GetStudySet(urlgo)
	test := (**C.char)(C.malloc(C.ulong(uintptr(len(termArr)) * unsafe.Sizeof(uintptr(1)))))
	test2 := (**C.char)(C.malloc(C.ulong(uintptr(len(termArr)) * unsafe.Sizeof(uintptr(1)))))
	for i, _ := range termArr {
		C.store_ptr(test, C.CString(termArr[i]), C.int(i), C.int(len(termArr[i])))
		C.store_ptr(test2, C.CString(defArr[i]), C.int(i), C.int(len(defArr[i])))
	}
	len := C.int(len(termArr))
	reterr := C.CString("")
	if err != nil {
		reterr = C.CString(err.Error())
	}
	return test, test2, len, reterr
}

func main() {}
