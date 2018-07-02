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
	NroTarjeta   int
	NroCliente   int
	ValidaDesde  string
	ValidaHasta  string
	CodSeguridad string
	LimiteCompra float32 //nose como poner lo de decimal(8,2)
	Estado       string
}

type Comercio struct {
	NroComercio  int
	Nombre       string
	Domicilio    string
	Codigopostal string
	Telefono     string
}

type Compra struct {
    NroOperacion int
    NroTarjeta   string
    NroComercio  int
    Fecha        string
    Monto        int
    Pagado       bool
}

var clientes = []Cliente{
	{NroCliente: 1, Nombre: "Jose", Apellido: "Argento", Domicilio: "Godoy Cruz 1064", Telefono: "4584-3863"},
	{NroCliente: 2, Nombre: "Mercedes", Apellido: "Benz", Domicilio: "Pte Peron 1223", Telefono: "4665-89892"},
	{NroCliente: 3, Nombre: "Megan", Apellido: "Ocaranza", Domicilio: "Tribulato 2345", Telefono: "4500-7651"},
}

var tarjetas = []Tarjeta{
	{NroTarjeta: 5703068016463339, NroCliente: 1, ValidaDesde: "201106", ValidaHasta: "201606", CodSeguridad: "1234", LimiteCompra: 200000.00, Estado: "anulada"},
	{NroTarjeta: 5578153904072665, NroCliente: 2, ValidaDesde: "201606", ValidaHasta: "201906", CodSeguridad: "1123", LimiteCompra: 200000.00, Estado: "vigente"},
	{NroTarjeta: 5681732770558693, NroCliente: 3, ValidaDesde: "201606", ValidaHasta: "201906", CodSeguridad: "1132", LimiteCompra: 200000.00, Estado: "vigente"},
}
var comercios = []Comercio{
	{NroComercio: 1, Nombre: "Anubis", Domicilio: "Av. Pres. Juan Domingo Peron 3497", Codigopostal: "1613", Telefono: "4463-5343"},
	{NroComercio: 2, Nombre: "Si A La Pizza", Domicilio: "25 de Mayo 2502", Codigopostal: "1613", Telefono: "4463-2314"},
	{NroComercio: 3, Nombre: "Narrow", Domicilio: "Av. Pres. Juan Domingo Peron 1420", Codigopostal: "1663", Telefono: "4667-7297"},
}
var compras = []Compra{
    {NroOperacion:1,NroTarjeta:"5703068016463339",NroComercio:3,Fecha:"2018-06-23",Monto:110,Pagado:true},
    {NroOperacion:2,NroTarjeta:"5578153904072665",NroComercio:1,Fecha:"2018-04-03",Monto:60,Pagado:true},
    {NroOperacion:3,NroTarjeta:"5578153904072665",NroComercio:2,Fecha:"2018-05-03",Monto:50,Pagado:true},
}

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

	nro := Tarjeta{5703068016463339, 1, "201106", "201606", "1234", 200000.00, "anulada"}
	data, err := json.Marshal(nro)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "tarjeta", []byte(strconv.Itoa(nro.NroTarjeta)), data)
	resultado, err := ReadUnique(db, "tarjeta", []byte(strconv.Itoa(nro.NroTarjeta)))
	fmt.Printf("%s\n", resultado)

	nro2 := Tarjeta{5578153904072665, 2, "201106", "201606", "1123", 200000.00, "vigente"}
	data2, err := json.Marshal(nro2)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "tarjeta", []byte(strconv.Itoa(nro2.NroTarjeta)), data2)
	resul, err := ReadUnique(db, "tarjeta", []byte(strconv.Itoa(nro2.NroTarjeta)))
	fmt.Printf("%s\n", resul)

	nro3 := Tarjeta{5681732770558693, 3, "201106", "201606", "1132", 200000.00, "vigente"}
	data3, err := json.Marshal(nro3)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "tarjeta", []byte(strconv.Itoa(nro3.NroTarjeta)), data3)
	resul3, err := ReadUnique(db, "tarjeta", []byte(strconv.Itoa(nro3.NroTarjeta)))
	fmt.Printf("%s\n", resul3)

}

func Comercios(db *bolt.DB) {

	nro := Comercio{1, "Anubis", "Av. Pres. Juan Domingo Peron 3497", "1613", "4463-5343"}
	data, err := json.Marshal(nro)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "comercios", []byte(strconv.Itoa(nro.NroComercio)), data)
	resultado, err := ReadUnique(db, "comercios", []byte(strconv.Itoa(nro.NroComercio)))
	fmt.Printf("%s\n", resultado)

	nro2 := Comercio{2, "Si A La Pizza", "25 de Mayo 2502", "1613", "4463-2314"}
	data2, err := json.Marshal(nro2)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "comercios", []byte(strconv.Itoa(nro2.NroComercio)), data2)
	resul, err := ReadUnique(db, "comercios", []byte(strconv.Itoa(nro2.NroComercio)))
	fmt.Printf("%s\n", resul)

	nro3 := Comercio{3, "Narrow", "Av. Pres. Juan Domingo Peron 1420", "1663", "4667-7297"}
	data3, err := json.Marshal(nro3)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "comercios", []byte(strconv.Itoa(nro3.NroComercio)), data3)
	resul3, err := ReadUnique(db, "comercios", []byte(strconv.Itoa(nro3.NroComercio)))
	fmt.Printf("%s\n", resul3)

}

func Compras(db *bolt.DB) {

	nro := Compra{1, "5703068016463339", 3, "2018-06-23", 110, true}
	data, err := json.Marshal(nro)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "compras", []byte(strconv.Itoa(nro.NroOperacion)), data)
	resultado, err := ReadUnique(db, "compras", []byte(strconv.Itoa(nro.NroOperacion)))
	fmt.Printf("%s\n", resultado)

	nro2 := Compra{2, "5578153904072665", 1, "2018-04-03", 60, true}
	data2, err := json.Marshal(nro2)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "compras", []byte(strconv.Itoa(nro2.NroOperacion)), data2)
	resul, err := ReadUnique(db, "compras", []byte(strconv.Itoa(nro2.NroOperacion)))
	fmt.Printf("%s\n", resul)

	nro3 := Compra{3, "5578153904072665", 2, "2018-05-03", 50,true}
	data3, err := json.Marshal(nro3)
	if err != nil {
		log.Fatal(err)
	}
	CreateUpdate(db, "compras", []byte(strconv.Itoa(nro3.NroOperacion)), data3)
	resul3, err := ReadUnique(db, "compras", []byte(strconv.Itoa(nro3.NroOperacion)))
	fmt.Printf("%s\n", resul3)

}

func main() {

	db, err := bolt.Open("postgres.db", 0600, nil)

	if err != nil {
		log.Fatal(err)

	}
	defer db.Close()
	Clientes(db)
	Tarjetas(db)
	Comercios(db)
	Compras(db)
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
