package main

/*
#include <ruby/ruby.h>

VALUE vibrant_from_url(VALUE,VALUE);
*/
import "C"

import (
	"fmt"
	"github.com/generaltso/vibrant"
	"image"
	_ "image/jpeg"
	"io"
	"log"
	"net/http"
	"unsafe"
)

func main() {
}

var rb_cVibrant C.VALUE

//export goobj_retain
func goobj_retain(obj unsafe.Pointer) {
	objects[obj] = true
}

//export goobj_free
func goobj_free(obj unsafe.Pointer) {
	delete(objects, obj)
}

//export vibrant_from_url
func vibrant_from_url(dummy C.VALUE, url C.VALUE) C.VALUE {
	fmt.Println(RbGoString(url))
	resp, err := http.Get(RbGoString(url))
	if err != nil {
		log.Fatalln(err)
	}
	return extract(resp.Body)
}

func extract(reader io.Reader) C.VALUE {
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalln(err)
	}
	palette, err := vibrant.NewPaletteFromImage(img)
	if err != nil {
		log.Fatalln(err)
	}
	return to_hash(palette.ExtractAwesome())
}

func to_hash(a map[string]*vibrant.Swatch) C.VALUE {
	hash := C.rb_hash_new()
	for name, swatch := range a {
		C.rb_hash_aset(hash, RbString(name), RbString(swatch.RGBHex()))
	}
	return hash
}

//export Init_vibrant
func Init_vibrant() {
	fmt.Println("Init")
	rb_cVibrant = rb_define_class("Vibrant", C.rb_cObject)
	rb_define_singleton_method(rb_cVibrant, "from_url", C.vibrant_from_url, 1)
}
