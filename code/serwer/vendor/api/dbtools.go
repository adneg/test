package api

import "time"

func CreateTableDB() {
	if DB.HasTable(&Dane{}) != true {
		DB.AutoMigrate(&Dane{})

	}

	if DB.HasTable(&Produkty{}) != true {
		DB.AutoMigrate(&Produkty{})
	}

	if DB.HasTable(&Grupy{}) != true {
		DB.AutoMigrate(&Grupy{})
	}
	if DB.HasTable(&Wspw{}) != true {
		DB.AutoMigrate(&Wspw{})

	}

	if DB.HasTable(&NormaOsobowa{}) != true {
		DB.AutoMigrate(&NormaOsobowa{})
		//norma osobowa
		nos := NormaOsobowa{}
		nos.Data = time.Now()
		nos.Norma_osobowa = 66
		DB.Save(&nos)
		// wspolczynnik w
		now := Wspw{}
		now.Data = time.Now()
		now.WspW = 0.95
		DB.Save(&now)

		//DANE TEST
		//		for i := 0; i < 100; i++ {

		//			t := Dane{}
		//			t.Id_grupy = 1
		//			t.Id_produktu = 2
		//			t.Ilosc = 301
		//			t.Nazwa_grupy = "SZWALNIA"
		//			t.Nazwa_produktu = "SPODNIE"
		//			t.Stan_osobowy = 10.5
		//			t.Wsp_produktu = 3.5
		//			t.Data = time.Now()
		//			t.Norma_osobowa = nos.Norma_osobowa
		//			t.WspW = now.WspW
		//			DB.Save(&t)
		//		}

		//GRUPA TEST
		//		w := Grupy{}
		//		w.Nazwa_grupy = "SZWALNIA"
		//		DB.Save(&w)
		//		//PRODUKT TEST
		//		z := Produkty{}
		//		z.Nazwa_produktu = "SPODNIE"
		//		z.Wsp_produktu = 3.34
		//		DB.Save(&z)

	}

}
