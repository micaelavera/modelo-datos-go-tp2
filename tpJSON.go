package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	bolt "github.com/coreos/bbolt"
)

// solo se marshalean los fields publicos

type Cliente struct {
	NroCliente int
	Nombre     string
	Apellido   string
	Domicilio  string
	Telefono   string
}

type Tarjeta struct {
	NroTarjeta   string
	NroCliente   int
	ValidaDesde  string
	ValidaHasta  string
	CodSeguridad string
	LimiteCompra float32 //nose como poner lo de decimal(8,2)
	Estado       string
}

type Comercio struct {
	Nrocomercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

var clientes = []Cliente{
	{NroCliente: 1, Nombre: "Jose", Apellido: "Argento", Domicilio: "Godoy Cruz 1064", Telefono: "4584-3863"},
	{NroCliente: 2, Nombre: "Mercedes", Apellido: "Benz", Domicilio: "Pte Peron 1223", Telefono: "4665-89892"},
	{NroCliente: 3, Nombre: "Megan", Apellido: "Ocaranza", Domicilio: "Tribulato 2345", Telefono: "4500-7651"},
}

/*
var tarjetas = []Tarjeta{
	{NroTarjeta:"5703068016463339" ,NroCliente:  1, ValidaDesde:"201106", ValidaHasta:"201606",CodSeguridad:"1234",LimiteCompra:200000.00, Estado:"anulada");
    {NroTarjeta:"5578153904072665" ,NroCliente:  2, ValidaDesde:"201606", ValidaHasta:"201906",CodSeguridad:"1123",LimiteCompra:200000.00, Estado:"vigente");
    {NroTarjeta:"5681732770558693" ,NroCliente:  3, ValidaDesde:"201606", ValidaHasta:"201906",CodSeguridad:"1132",LimiteCompra:200000.00, Estado:"vigente");
}
var comercios = []Comerco{
	{Nrocomercio:1, Nombre: "Anubis",			Domicilio: "Av. Pres. Juan Domingo Peron 3497", Codigopostal:"1613",Telefono:"4463-5343" }
	{Nrocomercio:2, Nombre: "Si A La Pizza" ,	domicilio:"25 de Mayo 2502",				 	Codigopostal:"1613",Telefono:"4463-2314" }
	{Nrocomercio:3, Nombre: "Narrow" ,			Domicilio:"Av. Pres. Juan Domingo Peron 1420",	Codigopostal:"1663",Telefono:"4667-7297" }

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

func Clientes(db *bolt.DB) {

	jose := Cliente{1, "Jose", "Argento", "Godoy Cruz 1064", "4584-3863"}
	data, err := json.Marshal(jose)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "cliente", []byte(strconv.Itoa(jose.NroCliente)), data)
	resultado, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(jose.NroCliente)))
	fmt.Printf("%s\n", resultado)

	mercedes := Cliente{2, "Mercedes", "Benz", "Pte Peron 1223", "4665-89892"}
	data2, err := json.Marshal(mercedes)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "cliente", []byte(strconv.Itoa(mercedes.NroCliente)), data2)
	resul, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(mercedes.NroCliente)))
	fmt.Printf("%s\n", resul)

	megan := Cliente{2, "Megan", "Ocaranza", "Tribulato 2345", "4500-7651"}
	data3, err := json.Marshal(megan)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "cliente", []byte(strconv.Itoa(megan.NroCliente)), data3)
	resul3, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(megan.NroCliente)))
	fmt.Printf("%s\n", resul3)

}

func Tarjetas(db *bolt.DB) {

	jose := Cliente{1, "Jose", "Argento", "Godoy Cruz 1064", "4584-3863"}
	data, err := json.Marshal(jose)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "cliente", []byte(strconv.Itoa(jose.NroCliente)), data)
	resultado, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(jose.NroCliente)))
	fmt.Printf("%s\n", resultado)

	mercedes := Cliente{2, "Mercedes", "Benz", "Pte Peron 1223", "4665-89892"}
	data2, err := json.Marshal(mercedes)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "cliente", []byte(strconv.Itoa(mercedes.NroCliente)), data2)
	resul, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(mercedes.NroCliente)))
	fmt.Printf("%s\n", resul)

	megan := Cliente{2, "Megan", "Ocaranza", "Tribulato 2345", "4500-7651"}
	data3, err := json.Marshal(megan)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "cliente", []byte(strconv.Itoa(megan.NroCliente)), data3)
	resul3, err := ReadUnique(db, "cliente", []byte(strconv.Itoa(megan.NroCliente)))
	fmt.Printf("%s\n", resul3)

}

func main() {

	db, err := bolt.Open("postgres.db", 0600, nil)

	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	Clientes(db)
	//	LeerDatosUsuario(db);
	/*
		data, err := json.MarshalIndent(clientes, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)

		var personas []Cliente
		err = json.Unmarshal(data, &personas)
		if err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}

		fmt.Printf("%v\n", personas)

		/*
		   //----------------------------------------
		   	db, err := bolt.Open(".db", 0600, nil)
		   	if err != nil {
		   		log.Fatal(err)
		   	}
		   	defer db.Close()
	*/

}
