package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	bolt "github.com/coreos/bbolt"
)

type Alumno struct {
	Legajo   int
	Nombre   string
	Apellido string
}

/*
type Cliente struct {
	NroCliente int
	Nombre string
	Apellido string
	Domicilio string
	Telefono string
}

type Tarjeta struct{
	NroTarjeta char(16)
	NroCliente int
	ValidaDesde char(6)
	ValidaHasta char(6)
	CodSeguridad char(4)
	LimiteCompra decimal(8,2)
	Estado char(10)
}
type Comercio struct{
	nrocomercio  integer,
	nombre       varchar(64),
	domicilio    varchar(64),
	codigopostal char(8),
	telefono     char(12)
}

var clientes = []Cliente{
	{NroCliente:1, Nombre: "Jose", Apellido: "Argento", Domicilio: "Godoy Cruz 1064", Telefono: "4584-3863"},
	{NroCliente: 2, Nombre: "Mercedes", Apellido: "Benz", Domicilio: "Pte Peron 1223", Telefono: "4665-89892"},
	{NroCliente: 3, Nombre: "Megan", Apellido: "Ocaranza", Domicilio: "Tribulato 2345", Telefono: "4500-7651"}
}

var tarjetas = []Tarjeta{
	{NroTarjeta:"5703068016463339" ,NroCliente:  1, ValidaDesde:"201106", ValidaHasta:"201606",CodSeguridad:"1234",LimiteCompra:200000.00, Estado:"anulada");
    {NroTarjeta:"5578153904072665" ,NroCliente:  2, ValidaDesde:"201606", ValidaHasta:"201906",CodSeguridad:"1123",LimiteCompra:200000.00, Estado:"vigente");
    {NroTarjeta:"5681732770558693" ,NroCliente:  3, ValidaDesde:"201606", ValidaHasta:"201906",CodSeguridad:"1132",LimiteCompra:200000.00, Estado:"vigente");
}
var comercios = []Comerco{
	{nrocomercio:1, nombre: "Anubis",			domicilio: "Av. Pres. Juan Domingo Peron 3497", codigopostal:"1613",telefono:"4463-5343" }
	{nrocomercio:2, nombre: "Si A La Pizza" ,	domicilio:"25 de Mayo 2502",				 	codigopostal:"1613",telefono:"4463-2314" }
	{nrocomercio:3, nombre: "Narrow" ,			domicilio:"Av. Pres. Juan Domingo Peron 1420",	codigopostal:"1663",telefono:"4667-7297" }

}

}
*/

func CreateUpdate(db *bolt.DB, bucketName string, key []byte, val []byte) error {
	//abre la transaccion de escritura
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _ := tx.CreateBucketIfNotExists([]byte(bucketName))

	err = b.Put(key, val)
	if err != nil {
		return err
	}

	//cierra transaccion
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func ReadUnique(db *bolt.DB, bucketName string, key []byte) ([]byte, error) {
	var buf []byte

	//abre una transaccion de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})

	return buf, err
}

func main() {
	db, err := bolt.Open("guaraní.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cristina := Alumno{1, "Cristina", "Kirchner"}
	data, err := json.Marshal(cristina)
	if err != nil {
		log.Fatal(err)
	}

	CreateUpdate(db, "alumno", []byte(strconv.Itoa(cristina.Legajo)), data)

	resultado, err := ReadUnique(db, "alumno", []byte(strconv.Itoa(cristina.Legajo)))

	fmt.Printf("%s\n", resultado)
}
