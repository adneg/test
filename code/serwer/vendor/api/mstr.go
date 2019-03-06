package api

import (
	"time"
)

type PercentProdukt struct {
	Ilosc int
	Sum   float32
}

type Grupy struct {
	Id_grupy    int    `gorm:"primary_key;AUTO_INCREMENT:index"`
	Nazwa_grupy string `gorm:"type:VARCHAR(300);not null;unique"`
}

func (Grupy) TableName() string {
	return "GRUPY"
}

type Produkty struct {
	Id_produktu    int     `gorm:"primary_key;AUTO_INCREMENT:index"`
	Nazwa_produktu string  `gorm:"type:VARCHAR(300);not null;unique"`
	Wsp_produktu   float32 `gorm:"type:REAL(100);not null"`
}

func (Produkty) TableName() string {
	return "PRODUKTY"
}

type NormaOsobowa struct {
	Id_norma_osobowa int       `gorm:"primary_key;AUTO_INCREMENT:index"`
	Norma_osobowa    float32   `gorm:"type:REAL(100);not null"`
	Data             time.Time `gorm:"type:DATETIME;not null"`
}

func (NormaOsobowa) TableName() string {
	return "NORMA_OSOBOWA"
}

type Wspw struct {
	Id_wspw int       `gorm:"primary_key;AUTO_INCREMENT:index"`
	WspW    float32   `gorm:"type:REAL(100);not null"`
	Data    time.Time `gorm:"type:DATETIME;not null"`
}

func (Wspw) TableName() string {
	return "WSPOLCZYNIK_W"
}

type Dane struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT:index"`

	Id_grupy    int    `gorm:"type:BIGINT(1000);not null"`
	Nazwa_grupy string `gorm:"type:VARCHAR(300);not null"`

	Id_produktu    int     `gorm:"type:BIGINT(1000);not null"`
	Nazwa_produktu string  `gorm:"type:VARCHAR(300);not null"`
	Wsp_produktu   float32 `gorm:"type:REAL(100);not null"`

	Stan_osobowy  float32   `gorm:"type:REAL(100);not null"`
	Ilosc         int       `gorm:"type:BIGINT(1000);not null"`
	Data          time.Time `gorm:"type:DATETIME;not null"`
	WspW          float32   `gorm:"type:REAL(100);not null"`
	Norma_osobowa float32   `gorm:"type:REAL(100);not null"`

	ZZZ   float32 `gorm:"type:REAL(100);not null"`
	Index int     `gorm:"-"`
}

func (Dane) TableName() string {
	return "DANE"
}

type Odp struct {
	Stat bool
	Id   int
}

type Mes struct {
	Instrukcja string
	Data       []byte
}
