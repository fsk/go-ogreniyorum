// Basit bir hesap makinesi
package main

import "fmt"

func calculator() {

	fmt.Println("Basit bir hesap makinesi")

	var process int
	var number1, number2 int

	for {
		fmt.Println("1. Toplama")
		fmt.Println("2. Cikarma")
		fmt.Println("3. Carpma")
		fmt.Println("4. Bolme")
		fmt.Println("5. Cikis")

		fmt.Print("Seciminizi yapin: ")
		fmt.Scan(&process)
		fmt.Print("Birinci sayiyi girin: ")
		fmt.Scan(&number1)
		fmt.Print("Ikinci sayiyi girin: ")
		fmt.Scan(&number2)

		if process == 1 {
			fmt.Println("Toplama islemi yapiliyor...")
			fmt.Println("Sonuc: ", number1+number2)
		} else if process == 2 {
			fmt.Println("Cikarma islemi yapiliyor...")
			fmt.Println("Sonuc: ", number1-number2)
		} else if process == 3 {
			fmt.Println("Carpma islemi yapiliyor...")
			fmt.Println("Sonuc: ", number1*number2)
		} else if process == 4 {
			fmt.Println("Bolme islemi yapiliyor...")
			if number2 == 0 {
				fmt.Println("Sifira bolme hatasi!")
				continue
			}
			fmt.Println("Sonuc: ", number1/number2)
		} else if process == 5 {
			fmt.Println("Programdan cikiliyor...")
			break
		} else if process > 5 || process < 1 {
			fmt.Println("Gecersiz secim yapildi!")
			continue
		}
	}
}

func main() {
	calculator()
}

/**
* init fonksiyonu, program ilk calistiginda otomatik olarak calisir.
* init fonksiyonu, package'in icinde olmalidir.
* init fonksiyonu, main fonksiyonundan once calisir.
* init fonksiyonu, birden fazla olabilir.
* init fonksiyonu, return type'i olmaz.
* Global degiskenleri initialize etmek icin kullanilir.
**/
//func init() {
//	calculator()
//}
