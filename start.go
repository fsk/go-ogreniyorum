package main

import "fmt"

func main() {

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
}
