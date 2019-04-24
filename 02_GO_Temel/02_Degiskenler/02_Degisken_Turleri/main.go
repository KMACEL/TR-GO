package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	//------------------------------------------------------
	// BOOLEAN - True-Flase  İçin Kullanılan Değişken
	//------------------------------------------------------
	/*
		bool    1 byte
	*/

	//------------------------------------------------------
	// bool
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var boolVariable bool
	boolVariable = true // yada false
	variableTypeBool := reflect.TypeOf(boolVariable)
	variableSizeBool := reflect.TypeOf(boolVariable).Size()
	variableValueBool := reflect.ValueOf(boolVariable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeBool)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeBool)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueBool)

	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	//------------------------------------------------------
	// Sözel İfadeler İçin Kullanılan Değişken
	//------------------------------------------------------
	/*
		string
	*/

	//------------------------------------------------------
	// string
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var stringVariable string
	stringVariable = "Merhaba Dünya"
	variableTypeString := reflect.TypeOf(stringVariable)
	variableSizeString := reflect.TypeOf(stringVariable).Size()
	variableValueString := reflect.ValueOf(stringVariable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeString)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeString)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueString)

	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	//------------------------------------------------------
	// Sayısal Değişkenler
	//------------------------------------------------------
	/*
		int     32bit yada 64bit
		int8     8 bit - 1 byte
		int16   16 bit - 2 byte
		int32   32 bit - 4 byte
		int64   64 bit - 8 byte

		'u' unsigned anlamına gelir. '-', 0 dan küçük değerleri almaz.
		Bu alanı '+', 0 dan büyük değerlere ekler.
		uint     32bit yada 64bit
		uint8     8 bit - 1 byte
		uint16   16 bit - 2 byte
		uint32   32 bit - 4 byte
		uint64   64 bit - 8 byte

		byte    8 bit - 1 byte = uint8
		rune    32 bit - 4 byte = int32
	*/

	//------------------------------------------------------
	// int Tipi,
	//		Eğer işletim sistemi 32 bit ise 32 bit değerinde
	//		Eğer işletim sistemi 64 bit ise 64 bit değerinde
	//------------------------------------------------------
	var intVariable int
	intVariable = 265789
	variableTypeInt := reflect.TypeOf(intVariable)
	variableBitInt := reflect.TypeOf(intVariable).Bits()
	variableSizeInt := reflect.TypeOf(intVariable).Size()
	variableValueInt := reflect.ValueOf(intVariable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeInt)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitInt)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeInt)
	fmt.Println("Değişkeninin Değeri                         : ", variableValueInt)

	//------------------------------------------------------
	// int8
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var int8Variable int8
	int8Variable = 125
	variableTypeInt8 := reflect.TypeOf(int8Variable)
	variableBitInt8 := reflect.TypeOf(int8Variable).Bits()
	variableSizeInt8 := reflect.TypeOf(int8Variable).Size()
	variableValueInt8 := reflect.ValueOf(int8Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeInt8)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitInt8)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeInt8)
	fmt.Println("Değişkeninin Değeri                         : ", variableValueInt8)
	fmt.Println("Değerin Alabileceği Aralık                  : ", math.MinInt8, " - ", math.MaxInt8)

	//------------------------------------------------------
	// int16
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var int16Variable int16
	int16Variable = 7489
	variableTypeInt16 := reflect.TypeOf(int16Variable)
	variableBitInt16 := reflect.TypeOf(int16Variable).Bits()
	variableSizeInt16 := reflect.TypeOf(int16Variable).Size()
	variableValueInt16 := reflect.ValueOf(int16Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeInt16)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitInt16)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeInt16)
	fmt.Println("Değişkeninin Değeri                         : ", variableValueInt16)
	fmt.Println("Değerin Alabileceği Aralık                  : ", math.MinInt16, " - ", math.MaxInt16)

	//------------------------------------------------------
	// int32
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var int32Variable int32
	int32Variable = 657875146
	variableTypeInt32 := reflect.TypeOf(int32Variable)
	variableBitInt32 := reflect.TypeOf(int32Variable).Bits()
	variableSizeInt32 := reflect.TypeOf(int32Variable).Size()
	variableValueInt32 := reflect.ValueOf(int32Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeInt32)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitInt32)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeInt32)
	fmt.Println("Değişkeninin Değeri                         : ", variableValueInt32)
	fmt.Println("Değerin Alabileceği Aralık                  : ", math.MinInt32, " - ", math.MaxInt32)

	//------------------------------------------------------
	// int64
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var int64Variable int64
	int64Variable = 8456287954138797414
	variableTypeInt64 := reflect.TypeOf(int64Variable)
	variableBitInt64 := reflect.TypeOf(int64Variable).Bits()
	variableSizeInt64 := reflect.TypeOf(int64Variable).Size()
	variableValueInt64 := reflect.ValueOf(int64Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeInt64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitInt64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeInt64)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueInt64)
	fmt.Println("Değerin Alabileceği Aralık                  : ", math.MinInt64, " - ", math.MaxInt64)

	//------------------------------------------------------
	// uint64 : 'u' unsigned anlamına gelir. '-', 0 dan küçük değerleri almaz.
	// Bu alanı '+', 0 dan büyük değerlere ekler. Bütün sayısal türlerde kullanılır
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var uint64Variable uint64
	uint64Variable = 25
	variableTypeUInt64 := reflect.TypeOf(uint64Variable)
	variableBitUInt64 := reflect.TypeOf(uint64Variable).Bits()
	variableSizeUInt64 := reflect.TypeOf(uint64Variable).Size()
	variableValueUInt64 := reflect.ValueOf(uint64Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeUInt64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitUInt64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeUInt64)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueUInt64)
	fmt.Println("Değerin Alabileceği Aralık                  : ", 0, " - ", uint64(math.MaxUint64))

	//------------------------------------------------------
	// byte = uint8
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var byteVariable byte
	byteVariable = 45
	variableTypeByte := reflect.TypeOf(byteVariable)
	variableBitByte := reflect.TypeOf(byteVariable).Bits()
	variableSizeByte := reflect.TypeOf(byteVariable).Size()
	variableValueByte := reflect.ValueOf(byteVariable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeByte)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitByte)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeByte)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueByte)

	//------------------------------------------------------
	// rune = int32
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var runeVariable rune
	byteVariable = 78
	variableTypeRune := reflect.TypeOf(runeVariable)
	variableBitRune := reflect.TypeOf(runeVariable).Bits()
	variableSizeRune := reflect.TypeOf(runeVariable).Size()
	variableValueRune := reflect.ValueOf(runeVariable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeRune)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitRune)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeRune)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueRune)

	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	//------------------------------------------------------
	// Ondalık İşlemler İçin Kullanılan Değişkenler
	//------------------------------------------------------
	/*
		float32  32 bit - 4 byte
		float64  64 bit - 8 byte
	*/

	//------------------------------------------------------
	// float32
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var float32Variable float32
	float32Variable = 78.36
	variableTypeFloat32 := reflect.TypeOf(float32Variable)
	variableBitFloat32 := reflect.TypeOf(float32Variable).Bits()
	variableSizeFloat32 := reflect.TypeOf(float32Variable).Size()
	variableValueFloat32 := reflect.ValueOf(float32Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeFloat32)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitFloat32)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeFloat32)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueFloat32)
	fmt.Println("Değerin Alabileceği Aralık                  : ", math.SmallestNonzeroFloat32, " - ", math.MaxFloat32)

	//------------------------------------------------------
	// float64
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var float64Variable float64
	float64Variable = 754984988.36444
	variableTypeFloat64 := reflect.TypeOf(float64Variable)
	variableBitFloat64 := reflect.TypeOf(float64Variable).Bits()
	variableSizeFloat64 := reflect.TypeOf(float64Variable).Size()
	variableValueFloat64 := reflect.ValueOf(float64Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeFloat64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitFloat64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeFloat64)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueFloat64)
	fmt.Println("Değerin Alabileceği Aralık                  : ", math.SmallestNonzeroFloat64, " - ", math.MaxFloat64)

	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")
	//------------------------------------------------------
	// Complex Sayılar İçin Kullanılan Değişkenler
	//------------------------------------------------------
	/*
		complex64    64 bit - 8 byte
		complex128  128 bit - 16 byte
	*/

	//------------------------------------------------------
	// complex64
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var complex64Variable complex64
	complex64Variable = complex(8, 9)
	variableTypeComplex64 := reflect.TypeOf(complex64Variable)
	variableBitComplex64 := reflect.TypeOf(complex64Variable).Bits()
	variableSizeComplex64 := reflect.TypeOf(complex64Variable).Size()
	variableValueComplex64 := reflect.ValueOf(complex64Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeComplex64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitComplex64)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeComplex64)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueComplex64)

	//------------------------------------------------------
	// complex128
	//------------------------------------------------------
	fmt.Println("------------------------------------------------------")
	var complex128Variable complex128
	complex128Variable = complex(8, 9)
	variableTypeComplex128 := reflect.TypeOf(complex128Variable)
	variableBitComplex128 := reflect.TypeOf(complex128Variable).Bits()
	variableSizeComplex128 := reflect.TypeOf(complex128Variable).Size()
	variableValueComplex128 := reflect.ValueOf(complex128Variable)

	fmt.Println("Değişkeninin Tipi                           : ", variableTypeComplex128)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BIT)  : ", variableBitComplex128)
	fmt.Println("Değişkeninin Bellekte Kapladığı Alan (BYTE) : ", variableSizeComplex128)
	fmt.Println("Değişkeninin İçinde Aktardığımız Değeri     : ", variableValueComplex128)

}
