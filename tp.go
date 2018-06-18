package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

func CrearTablas(db *sql.DB) {
	_, err := db.Exec(`create table cliente(
	nrocliente integer,
	nombre     varchar(64),
	apellido   varchar(64),
	domicilio  varchar(64),
	telefono   char(12)
);
create table tarjeta(
	nrotarjeta   char(12),
	nrocliente   integer,
	validadesde  char(6), --e.g 201106
	validahasta  char(6),
	codseguridad char(4),
	limitecompra decimal(8,2),
	estado       char(10) --'vigente', 'suspendida', 'anulada'
);
create table comercio(
	nrocomercio  integer,
	nombre       varchar(64),
	domicilio    varchar(64),
	codigopostal char(8),
	telefono     char(12)
);
create table compra(
	nrooperacion integer,
	nrotarjeta   char(16),
	nrocomercio  integer,
	fecha        timestamp,
	monto        decimal(7,2),
	pagado       boolean
);
create table rechazo(
	nrorechazo  integer,
	nrotarjeta  char(16),
	nrocomercio integer,
	fecha       timestamp,
	monto       decimal(7,2),
	motivo      varchar(64)
);
create table cierre(
	anio         integer,
	mes         integer,
	terminacion integer,
	fechainicio date,
	fechavto    date
);
create table cabecera(
	nroresumen  integer,
	nombre     varchar(64),
	apellido   varchar(64),
	domicilio  varchar(64),
	nrotarjeta char(16),
	desde      date,
	hasta      date,
	vence      date,
	total      decimal(8,2)
);
create table detalle(
	nroresumen      integer,
	nrolinea        integer,
	fecha           date,
	nombrecomercio  varchar(64),
	monto           decimal(7,2)
);
create table alerta(
	nroalerta   integer,
	nrotarjeta  char(16),
	fecha       timestamp,
	nrorechazo  integer,
	codalerta   integer, --0:rechazo, 1:compra 1min, 5:compra 5min, 32:límite
	descripcion  varchar(64)
);
create table consumo(
	nrotarjeta 	char(16),
	codseguridad	char(4),
	nrocomercio 	integer,
	monto        	decimal(7,2)
);`)
	if err != nil {
		log.Fatal(err)
	}

}

func CrearDB() {
	db, err := sql.Open("postgres", "user = postgres dbname = postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Exec para crear la DB tp2
	_, err = db.Exec(`create database tp2;`)

}
//prueba no sirve haha!
func LeerDatosUsuario(db *sql.DB) {
	var n int
	fmt.Printf("Enter 1 para crear la database: \n")
	fmt.Printf("Enter 2 para crear las tablas: \n")
	fmt.Printf("Enter 3 para agregar las Primary Keys: \n")
	fmt.Printf("Enter 4 para agregar las Foreign Keys: \n")
	fmt.Printf("Enter 5 para insertar los datos en las tablas: \n")
	fmt.Scanf("%d", &n)

	if n == 1 {
		CrearDB()
		
	fmt.Printf("Creando database ... \n")
	
	}else if n == 2 {
		CrearTablas(db)
		fmt.Printf("Creando las tablas ...\n")
	}
	
}


func main() {

	db, err := sql.Open("postgres", "user = postgres dbname = tp2 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	var salir=false
	for !salir{
		LeerDatosUsuario(db)
		var respuesta string
		fmt.Printf("¿Desea seguir en la aplicacion S/N?. Respuesta: \n")
		fmt.Scanf("%s",&respuesta)
		if respuesta=="N"|| respuesta=="n"{
			salir=true
		}

	}
}
