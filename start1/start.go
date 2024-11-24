package main

import (
	"fmt"
)

func main() {

	/**
	-- FMT --
	* format kelimesinin kısaltmasıdır.
	* fmt paketini import ettiğimizde kullanılır.
	* Go'nun standart kütüphanesidir.
	* IO icin kullanilir.
	**/

	/**
	** Sayısal Veri Tipleri
	** int, float, complex
	 */

	/**
	** Not: Go'da değişken tanımlarken := kullanılır.
	** := kullanırken değişkenin tipi belirtilmez.
	Yani bu konuyla alakalı bir zorunluluk yoktur.
	** Go derleyicisi, atanan değere göre tipi otomatik olarak belirleyebilir:
	*/

	// Bu şekilde sayi değişkeni int tipinde tanımlanır.
	// Go derleyicisi, atanan değere göre tipi otomatik olarak belirler.
	sayi := 10
	fmt.Println(sayi)

	/**
	-> int: Tam sayıları temsil eder. Platform bağımlıdır.
	Platform bağımlılığı, farklı işletim sistemlerinde (örn. Windows, Linux)
	veya donanım mimarilerinde (örn. x86, ARM) farklı tamsayı aralıkları
	gerektiren platformlar arası uyumsuzlukları önlemek için kullanılır.
	32 veya 64 bit olabilir.
	"Platform bağımlı" demek, değişkenin boyutunun (kaç bit olduğunun)
	çalıştığı bilgisayarın işletim sistemi mimarisine göre değişmesi demektir.
	**/

	/**
	-> int8: 8 bit tamsayıları temsil eder.
	-> int16: 16 bit tamsayıları temsil eder.
	-> int32: 32 bit tamsayıları temsil eder.
	-> int64: 64 bit tamsayıları temsil eder.
	**/

	// int8 -> -128 ile 127 arası değerleri alabilir.
	var sayi1 int8 = 10
	fmt.Println(sayi1)

	// int16 -> -32768 ile 32767 arası değerleri alabilir.
	var sayi2 int16 = 100
	fmt.Println(sayi2)

	// int32 -> -2147483648 ile 2147483647 arası değerleri alabilir.
	var sayi3 int32 = 1000
	fmt.Println(sayi3)

	// int64 -> -9223372036854775808 ile 9223372036854775807 arası değerleri alabilir.
	var sayi4 int64 = 10000
	fmt.Println(sayi4)

	/**
	-> uint: Pozitif tam sayıları temsil eder. Platform bağımlıdır.
	32 veya 64 bit olabilir.

	-> uint8: 8 bit pozitif tamsayıları temsil eder.
	-> uint16: 16 bit pozitif tamsayıları temsil eder.
	-> uint32: 32 bit pozitif tamsayıları temsil eder.
	-> uint64: 64 bit pozitif tamsayıları temsil eder.
	**/

	/**
	-- FLOAT --
	-> float32: 32 bit kayan noktalı sayıları temsil eder.
	-> float64: 64 bit kayan noktalı sayıları temsil eder.
	**/

	/**
	-- COMPLEX --
	Matematiksel olarak karmaşık sayıları temsil eder. (a + bi)
	-> complex64: 64 bit karmaşık sayıları temsil eder.
	-> complex128: 128 bit karmaşık sayıları temsil eder.
	**/

}
