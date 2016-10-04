package main

/*
#include <ruby/ruby.h>

VALUE NewGoStruct(VALUE klass, void *p);
void *GetGoStruct(VALUE obj);

VALUE paletteFromUrl(VALUE,VALUE);
VALUE paletteFromFile(VALUE,VALUE);
VALUE paletteSwatches(VALUE);
VALUE swatchColor(VALUE);
VALUE swatchPopulation(VALUE);
VALUE swatchName(VALUE);
*/
import "C"

import (
	"github.com/generaltso/vibrant"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"unsafe"
)

func main() {
}

var rb_cVibrant C.VALUE
var rb_cPalette C.VALUE
var rb_cSwatch C.VALUE

//export goobj_retain
func goobj_retain(obj unsafe.Pointer) {
	objects[obj] = true
}

//export goobj_free
func goobj_free(obj unsafe.Pointer) {
	delete(objects, obj)
}

//export paletteFromFile
func paletteFromFile(dummy C.VALUE, path C.VALUE) C.VALUE {
	file, err := os.Open(RbGoString(path))
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	defer file.Close()
	return extract(file)
}

//export paletteFromUrl
func paletteFromUrl(dummy C.VALUE, url C.VALUE) C.VALUE {
	resp, err := http.Get(RbGoString(url))
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	defer resp.Body.Close()
	return extract(resp.Body)
}

func extract(reader io.Reader) C.VALUE {
	img, _, err := image.Decode(reader)
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	palette, err := vibrant.NewPaletteFromImage(img)
	if err != nil {
		rb_raise(C.rb_eArgError, "'%s'", err)
	}
	return paletteNew(palette.ExtractAwesome())
}

func paletteNew(m map[string]*vibrant.Swatch) C.VALUE {
	p := C.rb_class_new_instance(0, nil, rb_cPalette)
	rb_iv_set(p, "@swatches", paletteSwatches(m))
	return p
}

func paletteSwatches(m map[string]*vibrant.Swatch) C.VALUE {
	ary := C.rb_ary_new()
	for _, s := range m {
		C.rb_ary_push(ary, swatchNew(rb_cSwatch, s))
	}
	return ary
}

func to_hash(a map[string]*vibrant.Swatch) C.VALUE {
	hash := C.rb_hash_new()
	for name, swatch := range a {
		C.rb_hash_aset(hash, RbString(name), swatchNew(rb_cSwatch, swatch))
	}
	return hash
}

func swatchNew(klass C.VALUE, s *vibrant.Swatch) C.VALUE {
	return C.NewGoStruct(klass, unsafe.Pointer(s))
}

//export swatchColor
func swatchColor(self C.VALUE) C.VALUE {
	swatch := (*vibrant.Swatch)(C.GetGoStruct(self))
	return INT2NUM(swatch.Color)
}

//export swatchPopulation
func swatchPopulation(self C.VALUE) C.VALUE {
	swatch := (*vibrant.Swatch)(C.GetGoStruct(self))
	return INT2NUM(swatch.Population)
}

//export swatchName
func swatchName(self C.VALUE) C.VALUE {
	swatch := (*vibrant.Swatch)(C.GetGoStruct(self))
	return RbString(swatch.Name)
}

//export Init_vibrant
func Init_vibrant() {
	sNew := "new"
	str_new := (*C.char)(unsafe.Pointer(&(*(*[]byte)(unsafe.Pointer(&sNew)))[0]))

	rb_cVibrant = rb_define_module("Vibrant")

	rb_cPalette = rb_define_class_under(rb_cVibrant, "Palette", C.rb_cObject)
	C.rb_undef_method(C.rb_class_of(rb_cPalette), str_new)
	rb_define_singleton_method(rb_cPalette, "from_url", C.paletteFromUrl, 1)
	rb_define_singleton_method(rb_cPalette, "from_file", C.paletteFromFile, 1)

	rb_cSwatch = rb_define_class_under(rb_cVibrant, "Swatch", C.rb_cObject)
	C.rb_undef_alloc_func(rb_cSwatch)
	C.rb_undef_method(C.rb_class_of(rb_cSwatch), str_new)
	rb_define_method(rb_cSwatch, "color", C.swatchColor, 0)
	rb_define_method(rb_cSwatch, "population", C.swatchPopulation, 0)
	rb_define_method(rb_cSwatch, "name", C.swatchName, 0)
}
