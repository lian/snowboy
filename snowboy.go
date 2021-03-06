/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.11
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: snowboy-detect-swig.i

package snowboy

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


typedef _gostring_ swig_type_1;
typedef _gostring_ swig_type_2;
typedef _gostring_ swig_type_3;
typedef _gostring_ swig_type_4;
typedef _gostring_ swig_type_5;
extern void _wrap_Swig_free_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_snowboy_6299b3f4c7855f3b(swig_intgo arg1);
extern uintptr_t _wrap_new_SnowboyDetect_snowboy_6299b3f4c7855f3b(swig_type_1 arg1, swig_type_2 arg2);
extern _Bool _wrap_SnowboyDetect_Reset_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern swig_intgo _wrap_SnowboyDetect_RunDetection__SWIG_0_snowboy_6299b3f4c7855f3b(uintptr_t arg1, swig_type_3 arg2);
extern swig_intgo _wrap_SnowboyDetect_RunDetection__SWIG_1_snowboy_6299b3f4c7855f3b(uintptr_t arg1, swig_voidp arg2, swig_intgo arg3);
extern swig_intgo _wrap_SnowboyDetect_RunDetection__SWIG_2_snowboy_6299b3f4c7855f3b(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3);
extern swig_intgo _wrap_SnowboyDetect_RunDetection__SWIG_3_snowboy_6299b3f4c7855f3b(uintptr_t arg1, uintptr_t arg2, swig_intgo arg3);
extern void _wrap_SnowboyDetect_SetSensitivity_snowboy_6299b3f4c7855f3b(uintptr_t arg1, swig_type_4 arg2);
extern swig_type_5 _wrap_SnowboyDetect_GetSensitivity_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern void _wrap_SnowboyDetect_SetAudioGain_snowboy_6299b3f4c7855f3b(uintptr_t arg1, float arg2);
extern void _wrap_SnowboyDetect_UpdateModel_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern swig_intgo _wrap_SnowboyDetect_NumHotwords_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern void _wrap_SnowboyDetect_ApplyFrontend_snowboy_6299b3f4c7855f3b(uintptr_t arg1, _Bool arg2);
extern swig_intgo _wrap_SnowboyDetect_SampleRate_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern swig_intgo _wrap_SnowboyDetect_NumChannels_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern swig_intgo _wrap_SnowboyDetect_BitsPerSample_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
extern void _wrap_delete_SnowboyDetect_snowboy_6299b3f4c7855f3b(uintptr_t arg1);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_snowboy_6299b3f4c7855f3b(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrSnowboyDetect uintptr

func (p SwigcptrSnowboyDetect) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrSnowboyDetect) SwigIsSnowboyDetect() {
}

func NewSnowboyDetect(arg1 string, arg2 string) (_swig_ret SnowboyDetect) {
	var swig_r SnowboyDetect
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (SnowboyDetect)(SwigcptrSnowboyDetect(C._wrap_new_SnowboyDetect_snowboy_6299b3f4c7855f3b(*(*C.swig_type_1)(unsafe.Pointer(&_swig_i_0)), *(*C.swig_type_2)(unsafe.Pointer(&_swig_i_1)))))
	if Swig_escape_always_false {
		Swig_escape_val = arg1
	}
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) Reset() (_swig_ret bool) {
	var swig_r bool
	_swig_i_0 := arg1
	swig_r = (bool)(C._wrap_SnowboyDetect_Reset_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) RunDetection__SWIG_0(arg2 string) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_SnowboyDetect_RunDetection__SWIG_0_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0), *(*C.swig_type_3)(unsafe.Pointer(&_swig_i_1))))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) RunDetection__SWIG_1(arg2 *float32, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_SnowboyDetect_RunDetection__SWIG_1_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0), C.swig_voidp(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) RunDetection__SWIG_2(arg2 Int16_t, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_SnowboyDetect_RunDetection__SWIG_2_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) RunDetection__SWIG_3(arg2 Int32_t, arg3 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2.Swigcptr()
	_swig_i_2 := arg3
	swig_r = (int)(C._wrap_SnowboyDetect_RunDetection__SWIG_3_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1), C.swig_intgo(_swig_i_2)))
	return swig_r
}

func (p SwigcptrSnowboyDetect) RunDetection(a ...interface{}) int {
	argc := len(a)
	if argc == 1 {
		return p.RunDetection__SWIG_0(a[0].(string))
	}
	if argc == 2 {
		if _, ok := a[0].(*float32); !ok {
			goto check_2
		}
		return p.RunDetection__SWIG_1(a[0].(*float32), a[1].(int))
	}
check_2:
	if argc == 2 {
		if _, ok := a[0].(SwigcptrInt16_t); !ok {
			goto check_3
		}
		return p.RunDetection__SWIG_2(a[0].(Int16_t), a[1].(int))
	}
check_3:
	if argc == 2 {
		return p.RunDetection__SWIG_3(a[0].(Int32_t), a[1].(int))
	}
	panic("No match for overloaded function call")
}

func (arg1 SwigcptrSnowboyDetect) SetSensitivity(arg2 string) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_SnowboyDetect_SetSensitivity_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0), *(*C.swig_type_4)(unsafe.Pointer(&_swig_i_1)))
	if Swig_escape_always_false {
		Swig_escape_val = arg2
	}
}

func (arg1 SwigcptrSnowboyDetect) GetSensitivity() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_SnowboyDetect_GetSensitivity_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

func (arg1 SwigcptrSnowboyDetect) SetAudioGain(arg2 float32) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_SnowboyDetect_SetAudioGain_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0), C.float(_swig_i_1))
}

func (arg1 SwigcptrSnowboyDetect) UpdateModel() {
	_swig_i_0 := arg1
	C._wrap_SnowboyDetect_UpdateModel_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrSnowboyDetect) NumHotwords() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_SnowboyDetect_NumHotwords_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) ApplyFrontend(arg2 bool) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	C._wrap_SnowboyDetect_ApplyFrontend_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0), C._Bool(_swig_i_1))
}

func (arg1 SwigcptrSnowboyDetect) SampleRate() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_SnowboyDetect_SampleRate_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) NumChannels() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_SnowboyDetect_NumChannels_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrSnowboyDetect) BitsPerSample() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_SnowboyDetect_BitsPerSample_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func DeleteSnowboyDetect(arg1 SnowboyDetect) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_SnowboyDetect_snowboy_6299b3f4c7855f3b(C.uintptr_t(_swig_i_0))
}

type SnowboyDetect interface {
	Swigcptr() uintptr
	SwigIsSnowboyDetect()
	Reset() (_swig_ret bool)
	RunDetection(a ...interface{}) int
	SetSensitivity(arg2 string)
	GetSensitivity() (_swig_ret string)
	SetAudioGain(arg2 float32)
	UpdateModel()
	NumHotwords() (_swig_ret int)
	ApplyFrontend(arg2 bool)
	SampleRate() (_swig_ret int)
	NumChannels() (_swig_ret int)
	BitsPerSample() (_swig_ret int)
}


type SwigcptrInt16_t uintptr
type Int16_t interface {
	Swigcptr() uintptr;
}
func (p SwigcptrInt16_t) Swigcptr() uintptr {
	return uintptr(p)
}

type SwigcptrInt32_t uintptr
type Int32_t interface {
	Swigcptr() uintptr;
}
func (p SwigcptrInt32_t) Swigcptr() uintptr {
	return uintptr(p)
}

