package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)



func crearTablas(db *sql.DB) {
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
	nrotarjeta  char(16), fk alter table alerta add constraint alerta_fk0  foreing key (nrotarjeta) references tarjeta (nrotarjeta);
	fecha       timestamp,
	nrorechazo  integer;fk
	codalerta   integer --0:rechazo, 1:compra 1min, 5:compra 5min, 32:límite
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

func crearDB() {
	db, err := sql.Open("postgres", "user = postgres dbname = postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Exec para crear la DB tp2
	_, err = db.Exec(`create database tp2;`)

}
//prueba no sirve haha!
func creatodojunto(db *sql.DB) {
	var n int
	
	for i := 0; i < 2; i++ { {
		if i == 0 {
			fmt.Printf("Enter 1 to create database: \n")
			fmt.Scanf("%d", &n)
			if n == 1 {
				crearDB()
			}
		}

		if i == 1 {
			fmt.Printf("Enter 2 to create tables: \n")
			if n == 2 {
				crearTablas(db)
			}
		}
		i++
	}

}

func main() {

	db, err := sql.Open("postgres", "user = postgres dbname = tp2 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	inicio(db)

	//var n int
	//fmt.Printf("Enter 1 to create database: \n")
	//fmt.Printf("Enter 2 to create tables: \n")

	//fmt.Scanf("%d", &n)
	//if n == 1 {
	//	crearDB()
	//} else if n == 2 {
	//	crearTablas(db)
	//}

	// fmt.Printf(tablas(1))
}
