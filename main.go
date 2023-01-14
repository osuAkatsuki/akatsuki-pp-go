package main

/*
#cgo LDFLAGS: -L./lib -lakatsuki_pp_ffi
#include "./lib/akatsuki_pp_ffi.h"
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Calculator struct {
	MapPath string
}

type ScoreParams struct {
	Mode          uint
	Mods          uint
	MaxCombo      uint
	Accuracy      float64
	MissCount     uint
	PassedObjects uint
}

type CalculatePerformanceResult struct {
	PP    float64
	Stars float64
}

func (calculator Calculator) Calculate(params ScoreParams) CalculatePerformanceResult {
	cMapPath := C.CString(calculator.MapPath)
	defer C.free(unsafe.Pointer(cMapPath))

	passedObjects := C.optionu32{t: C.uint(0), is_some: C.uchar(0)}
	if params.PassedObjects > 0 {
		passedObjects = C.optionu32{t: C.uint(params.PassedObjects), is_some: C.uchar(1)}
	}

	rawResult := C.calculate_score(
		cMapPath,
		C.uint(params.Mode),
		C.uint(params.Mods),
		C.uint(params.MaxCombo),
		C.double(params.Accuracy),
		C.uint(params.MissCount),
		passedObjects,
	)
	return CalculatePerformanceResult{PP: float64(rawResult.pp), Stars: float64(rawResult.stars)}
}

func main() {
	calculator := Calculator{MapPath: "./test.osu"}
	scoreParams := ScoreParams{
		Mode:          0,
		Mods:          72,
		MaxCombo:      542,
		Accuracy:      98.60,
		MissCount:     0,
		PassedObjects: 0,
	}

	result := calculator.Calculate(scoreParams)
	fmt.Println(result.PP)
	fmt.Println(result.Stars)
}
