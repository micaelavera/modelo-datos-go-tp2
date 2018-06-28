package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func CrearDB() {
	db, err := sql.Open("postgres", "user = postgres dbname = postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Exec para crear la DB tp2
	_, err = db.Exec(`create database tp2;`)

}

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

func AgregarPKs(db *sql.DB) {
	_, err := db.Exec(`alter table tarjeta  add constraint tarjeta_pk   primary key (nrotarjeta);
		alter table comercio add constraint comercio_pk  primary key (nrocomercio);
		alter table compra   add constraint compra_pk    primary key (nrooperacion);
		alter table rechazo  add constraint rechazo_pk   primary key (nrorechazo);
		alter table cierre   add constraint cierre_pk    primary key (anio,mes,terminacion);
		alter table cierre   add constraint cierre_pk    primary key (mes,terminacion);
		alter table cabecera add constraint cabecera_pk  primary key (nroresumen);
		alter table detalle  add constraint detalle_pk   primary key (nroresumen,nrolinea);
		alter table alerta   add constraint alerta_pk    primary key (nroalerta);`)

	if err != nil {
		log.Fatal(err)
	}

}
func AgregarFKs(db *sql.DB) {
	_, err := db.Exec(`	--FOREIGN KEY
		alter table tarjeta  add constraint tarjeta_fk0 foreign key (nrocliente)  references cliente  (nrocliente);
		alter table compra   add constraint compra_fk0  foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
		alter table compra   add constraint compra_fk1  foreign key (nrocomercio) references comercio (nrocomercio);
		alter table rechazo  add constraint rechazo_fk0 foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
		alter table rechazo  add constraint rechazo_fk1 foreign key (nrocomercio) references comercio (nrocomercio);
		alter table cabecera add constraint cabecera_fk foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
		alter table alerta   add constraint alerta_fk0  foreign key (nrotarjeta)  references tarjeta (nrotarjeta);
		`)
	if err != nil {
		log.Fatal(err)
	}
	//Alter table nombredelatabla drop constraint nombre_pk;
}
func eliminarPKs(db *sql.DB) {
	_, err := db.Exec(`--DROP PRIMARY KEYs
	alter table tarjeta  drop constraint tarjeta_pk;
	alter table comercio drop constraint comercio_pk;
	alter table compra   drop constraint compra_pk;
	alter table rechazo  drop constraint rechazo_pk;
	alter table cierre   drop constraint cierre_pk;
	alter table cierre   drop constraint cierre_pk;
	alter table cabecera drop constraint cabecera_pk;
	alter table detalle  drop constraint detalle_pk;
	alter table alerta   drop constraint alerta_pk;
	`)
	if err != nil {
		log.Fatal(err)
	}
}
func eliminarFKs(db *sql.DB) {
	_, err := db.Exec(`	-- DROP FOREIGN KEYs
		alter table tarjeta  drop constraint tarjeta_fk0;
		alter table compra   drop constraint compra_fk0;
		alter table compra   drop constraint compra_fk1;
		alter table rechazo  drop constraint rechazo_fk0;
		alter table rechazo  drop constraint rechazo_fk1;
		alter table cabecera drop constraint cabecera_fk;
		alter table alerta   drop constraint alerta_fk0;
		`)
	if err != nil {
		log.Fatal(err)
	}
}

func AlertarClientes_1min() {
	for {
		//funcion del alerta al cliente cada minuto
		time.Sleep(1 * time.Minute)
	}
}

func AlertarClientes_5min() {
	for {
		//funcion del alerta al cliente cada 5 minutos
		time.Sleep(5 * time.Minute)
	}
}

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

	} else if n == 2 {
		CrearTablas(db)
		fmt.Printf("Creando las tablas ...\n")
	} else if n == 3 {
		AgregarPKs(db)
		fmt.Printf("Creando las  Primary Keys ...\n")
	} else if n == 4 {
		AgregarFKs(db)
		fmt.Printf("Creando  las Foreign Keys ...\n")
	}

}

func main() {

	db, err := sql.Open("postgres", "user = postgres dbname = tp2 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*
	   	rows, err := db.Query(`select * from alumno`)
	   if err != nil {
	   log.Fatal(err)
	   }
	   defer rows.Close()
	   var a alumno
	   for rows.Next() {
	   err := rows.Scan(&a.legajo, &a.nombre, &a.apellido)
	   if err != nil {
	   log.Fatal(err)
	   }
	   fmt.Printf("%v %v %v\n", a.legajo, a.nombre, a.apellido)
	   }
	   if err = rows.Err(); err != nil {
	   log.Fatal(err)
	   }
	*/

	var salir = false
	for !salir {
		LeerDatosUsuario(db)
		var respuesta string
		fmt.Printf("¿Desea seguir en la aplicacion S/N?. Respuesta: \n")
		fmt.Scanf("%s", &respuesta)
		if respuesta == "N" || respuesta == "n" {
			salir = true
		}

	}
}
