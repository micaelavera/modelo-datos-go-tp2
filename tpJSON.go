package main

import (
	"encoding/json"
	"fmt"
	bolt "github.com/coreos/bbolt"
	"log"
	"strconv"
)


type Alumno struct {
	Legajo int
	Nombre string
	Apellido string
}

func CreateUpdate(db *bolt.DB, bucketName string,key []byte, val []byte) error {
	//abre la transaccion de escritura
	tx, err := db.Begin (true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	b, _:=tx.CreateBucketIfNotExists([]byte(bucketName))

	err=b.Put(key, val)
	if err != nil {
		return err
	}

	//cierra transaccion
	if err := tx.Commit(); err != nil {
		return err
	}
	
	return nil
}


func ReadUnique (db *bolt.DB, bucketName string, key []byte)([]byte, error) {
	var buf []byte

	//abre una transaccion de lectura
	err := db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucketb := tx.Bucket([]byte(bucketName))
		buf = b.Get(key)
		return nil
	})

	return buf,err
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
