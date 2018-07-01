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
	nrotarjeta   char(16),
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
	nrooperacion serial,
	nrotarjeta   char(16),
	nrocomercio  integer,
	fecha        timestamp,
	monto        decimal(7,2),
	pagado       boolean
);
create table rechazo(
	nrorechazo  serial,
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
	fechacierre date,
	fechavto    date
);
create table cabecera(
	nroresumen  serial,
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
	nroalerta   serial,
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
	_, err := db.Exec(`alter table cliente  add constraint cliente_pk   primary key (nrocliente);
		alter table tarjeta  add constraint tarjeta_pk   primary key (nrotarjeta);
		alter table comercio add constraint comercio_pk  primary key (nrocomercio);
		alter table compra   add constraint compra_pk    primary key (nrooperacion);
		alter table rechazo  add constraint rechazo_pk   primary key (nrorechazo);
		alter table cierre   add constraint cierre_pk  primary key (anio,mes,terminacion);
		alter table cabecera add constraint cabecera_pk  primary key (nroresumen);
		alter table detalle  add constraint detalle_pk   primary key (nroresumen,nrolinea);
		alter table alerta   add constraint alerta_pk    primary key (nroalerta);`)

	if err != nil {
		log.Fatal(err)
	}

}
func AgregarFKs(db *sql.DB) {
	_, err := db.Exec(`alter table tarjeta  add constraint tarjeta_fk0 foreign key (nrocliente)  references cliente  (nrocliente);
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
	_, err := db.Exec(`alter table cliente  drop constraint cliente_pk;
	alter table tarjeta  drop constraint tarjeta_pk;
	alter table comercio drop constraint comercio_pk;
	alter table compra   drop constraint compra_pk;
	alter table rechazo  drop constraint rechazo_pk;
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
	_, err := db.Exec(`alter table tarjeta  drop constraint tarjeta_fk0;
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

func InsertarDatos(db *sql.DB) {
	_, err := db.Exec(`--CLIENTES
    insert into cliente values(1,  'José',      'Argento',      'Godoy Cruz 1064',      '4584-3863');
    insert into cliente values(2,  'Mercedes',  'Benz',         'Pte Perón 1223',       '4665-8989');
    insert into cliente values(3,  'Megan',     'Ocaranza',     'Tribulato 2345',       '4500-7651');
    insert into cliente values(4,  'Luis',      'Rios',         'Dorrego 1234',         '4213-0153');
    insert into cliente values(5,  'Julio',     'Cortazar',     ' Av Balbin 534',       '4890-8747');
    insert into cliente values(6,  'Tomás',     'Aguirre',      'Urquiza 540',          '4707-1600');
    insert into cliente values(7,  'Juan',      'Avalos',       'Av E Perón 7716',      '4589-3191');
    insert into cliente values(8,  'Prudencia', 'Arzuaga',      'Rivadavia 416',        '4665-2333');
    insert into cliente values(9,  'Damiana',   'Molina',       'Malvinas 890',         '4129-0964');
    insert into cliente values(10, 'Ramón',     'Perez',        'Las Heras 460',        '4556-8970');
    insert into cliente values(11, 'Fernando',  'Álvarez',      'Tribulato 1290',       '4908-7822');
    insert into cliente values(12, 'Carla',     'Estrella',     'Primera Junta 7865',   '4444-5625');
    insert into cliente values(13, 'Farid',     'Hasan',        'España 438',           '4123-9078');
    insert into cliente values(14, 'Alicia',    'Castillo',     'Rodriguez Peña 170',   '4367-7801');
    insert into cliente values(15, 'Elsa',      'López',        'Argüero 1138',         '4563-2323');
    insert into cliente values(16, 'Harry',     'Potter',       'Gutiérrez 908',        '4768-9475');
    insert into cliente values(17, 'Pedro',     'Moreno',       'Julian Rejala 999',    '4556-9872');
    insert into cliente values(18, 'Ofelia',    'Le Brun',      'Paunero 7856',         '4389-7531');
    insert into cliente values(19, 'Laura',     'Martin',       'Las Margaritas 484',   '4895-7939');
    insert into cliente values(20, 'Evan',      'Peters',       'Irigoin 296',          '4664-1640');

--TARJETAS

    insert into tarjeta values('5703068016463339' ,  1, '201106', '201606','1234',200000.00,'anulada');
    insert into tarjeta values('5578153904072665' ,  2, '201606', '201906','1123',200000.00,'vigente');
    insert into tarjeta values('5681732770558693' ,  3, '201606', '201906','1132',200000.00,'vigente');
    insert into tarjeta values('5460322592744445' ,  4, '201606', '201906','1231',200000.00,'vigente');
    insert into tarjeta values('5430913178957141' ,  5, '201606', '201906','2122',200000.00,'vigente');
    insert into tarjeta values('5019155187146691' ,  6, '201606', '201906','2433',200000.00,'vigente');
    insert into tarjeta values('5019792017295163' ,  7, '201606', '201906','2344',200000.00,'vigente');
    insert into tarjeta values('5019919938293361' ,  8, '201606', '201906','2235',200000.00,'vigente');
    insert into tarjeta values('6297661287321366' ,  9, '201606', '201906','3236',200000.00,'vigente');
    insert into tarjeta values('6286339466940040' , 10, '201606', '201906','3337',200000.00,'vigente');
    insert into tarjeta values('6270607345512992' , 11, '201606', '201906','3438',200000.00,'vigente');
    insert into tarjeta values('6204296517127663' , 12, '201606', '201906','3239',200000.00,'vigente');
    insert into tarjeta values('6293235484090035' , 13, '201606', '201906','4110',200000.00,'vigente');
    insert into tarjeta values('6921678767563105' , 14, '201606', '201906','4211',200000.00,'vigente');
    insert into tarjeta values('6334546019765631' , 15, '201606', '201906','4412',200000.00,'vigente');
    insert into tarjeta values('6761361843313612' , 16, '201606', '201906','4313',200000.00,'vigente');
    insert into tarjeta values('6528492702135343' , 17, '201606', '201906','5214',200000.00,'vigente');
    insert into tarjeta values('6585963936581775' , 18, '201606', '201906','6215',200000.00,'vigente');
    insert into tarjeta values('3696418858377210' , 19, '201606', '201906','7216',200000.00,'vigente');
    insert into tarjeta values('3689635420613720' , 19, '201606', '201906','8217',200000.00,'vigente');
    insert into tarjeta values('3033446987174022' , 20, '201606', '201906','9218',200000.00,'vigente');
    insert into tarjeta values('3602403813503232' , 20, '201606', '201906','1119',200000.00,'vigente');

--COMERCIOS

    insert into comercio values( 1,'Anubis'              ,'Av. Pres. Juan Domingo Peron 3497','1613','4463-5343' );
    insert into comercio values( 2,'Si A La Pizza'       ,'25 de Mayo 2502'                  ,'1613','4463-2314' );
    insert into comercio values( 3,'Narrow'              ,'Av. Pres. Juan Domingo Peron 1420','1663','4667-7297' );
    insert into comercio values( 4,'Starbucks Coffee'    ,'Parana 3745'                      ,'1640','4748-0098' );
    insert into comercio values( 5,'47 Street'           ,'Cruce Ruta 8 y Ruta 202'          ,'1613','4667-5770' );
    insert into comercio values( 6,'Frávega'             ,'Av. Pres. Juan Domingo Peron 1127','1663','4667-4009' );
    insert into comercio values( 7,'Optica Ivaldi'       ,'Av. Pres. Juan Domingo Peron 1645','1663','4667-2332' );
    insert into comercio values( 8,'Farmacity'           ,'Av. Constituyentes 6093/99'       ,'1617','4587-8243' );
    insert into comercio values( 9,'Cúspide'             ,'Cruce Ruta 8 y Ruta 202'          ,'1613','1521508092');
    insert into comercio values(10,'Garbarino'           ,'Av. Dr. Ricardo Balbín 1198'      ,'1663','4667-6534' );
    insert into comercio values(11,'Starbucks Coffee'    ,'Cruce Ruta 8 y Ruta 202'          ,'1613','4667-5434' );
    insert into comercio values(12,'Bonafide'            ,' Italia 1249'                     ,'1663','4667-4545' );
    insert into comercio values(13,'Optica Gris'         ,'Av. Arturo Illia 5243'            ,'1613','4463-8344' );
    insert into comercio values(14,'Matu Jean´s'         ,'Av. Pres.Juan Domingo ]Peron 3300','1613','4463-9089' );
    insert into comercio values(15,'Compumundo'          ,'Belgrano 1401'                    ,'1663','4667-3425' );
    insert into comercio values(16,'Falabella'           ,'Parana 3745'                      ,'1640','4717-8100' );
    insert into comercio values(17,'McDonald´s'          ,'Av. Pres. Juan Domingo Peron 983' ,'1662','4668-0912' );
    insert into comercio values(18,'M 58'                ,'Charlone 1201'                    ,'1663','4667-4532' );
    insert into comercio values(19,'Cine Hoyts Unicenter','Parana 3745'                      ,'1640','4717-8109' );
    insert into comercio values(20,'Solo Deportes'       ,'Av. Pres. Juan Domingo Peron 1317','1663','4667-3453' );

--CIERRES

    insert into cierre values(2018,1,0 ,'2018-01-27','2018-02-27','2018-03-07');
    insert into cierre values(2018,1,1 ,'2018-01-15','2018-02-15','2018-02-25');
    insert into cierre values(2018,1,2 ,'2018-01-03','2018-02-03','2018-02-15');
    insert into cierre values(2018,1,3 ,'2018-01-07','2018-02-07','2018-02-15');
    insert into cierre values(2018,1,4 ,'2018-01-05','2018-02-05','2018-02-15');
    insert into cierre values(2018,1,5 ,'2018-01-24','2018-02-24','2018-03-07');
    insert into cierre values(2018,1,6 ,'2018-01-28','2018-02-28','2018-03-10');
    insert into cierre values(2018,1,7 ,'2018-01-07','2018-02-07','2018-02-17');
    insert into cierre values(2018,1,8 ,'2018-01-23','2018-02-23','2018-03-05');
    insert into cierre values(2018,1,9 ,'2018-01-09','2018-02-09','2018-02-19');
    insert into cierre values(2018,2,0 ,'2018-02-03','2018-03-03','2018-03-13');
    insert into cierre values(2018,2,1 ,'2018-02-07','2018-03-07','2018-03-17');
    insert into cierre values(2018,2,2 ,'2018-02-19','2018-03-19','2018-03-27');
    insert into cierre values(2018,2,3 ,'2018-02-24','2018-03-24','2018-03-29');
    insert into cierre values(2018,2,4 ,'2018-02-27','2018-03-27','2018-04-02');
    insert into cierre values(2018,2,5 ,'2018-02-02','2018-03-02','2018-03-12');
    insert into cierre values(2018,2,6 ,'2018-02-04','2018-03-04','2018-03-14');
    insert into cierre values(2018,2,7 ,'2018-02-16','2018-03-16','2018-03-26');
    insert into cierre values(2018,2,8 ,'2018-02-13','2018-03-13','2018-03-23');
    insert into cierre values(2018,2,9 ,'2018-02-09','2018-03-09','2018-03-19');
    insert into cierre values(2018,3,0 ,'2018-03-01','2018-04-01','2018-04-11');
    insert into cierre values(2018,3,1 ,'2018-03-04','2018-04-04','2018-04-14');
    insert into cierre values(2018,3,2 ,'2018-03-14','2018-04-14','2018-04-24');
    insert into cierre values(2018,3,3 ,'2018-03-24','2018-04-24','2018-05-05');
    insert into cierre values(2018,3,4 ,'2018-03-18','2018-04-18','2018-04-28');
    insert into cierre values(2018,3,5 ,'2018-03-15','2018-04-15','2018-04-25');
    insert into cierre values(2018,3,6 ,'2018-03-28','2018-04-28','2018-05-03');
    insert into cierre values(2018,3,7 ,'2018-03-12','2018-04-12','2018-04-22');
    insert into cierre values(2018,3,8 ,'2018-03-07','2018-04-07','2018-04-17');
    insert into cierre values(2018,3,9 ,'2018-03-22','2018-04-22','2018-05-02');
    insert into cierre values(2018,4,0 ,'2018-04-26','2018-05-26','2018-06-06');
    insert into cierre values(2018,4,1 ,'2018-04-28','2018-05-28','2018-06-06');
    insert into cierre values(2018,4,2 ,'2018-04-01','2018-05-01','2018-05-11');
    insert into cierre values(2018,4,3 ,'2018-04-12','2018-05-12','2018-05-22');
    insert into cierre values(2018,4,4 ,'2018-04-21','2018-05-21','2018-05-27');
    insert into cierre values(2018,4,5 ,'2018-04-15','2018-05-15','2018-05-26');
    insert into cierre values(2018,4,6 ,'2018-04-19','2018-05-19','2018-05-26');
    insert into cierre values(2018,4,7 ,'2018-04-20','2018-05-20','2018-05-26');
    insert into cierre values(2018,4,8 ,'2018-04-04','2018-05-04','2018-05-13');
    insert into cierre values(2018,4,9 ,'2018-04-06','2018-05-06','2018-05-15');
    insert into cierre values(2018,5,0 ,'2018-05-05','2018-06-05','2018-06-15');
    insert into cierre values(2018,5,1 ,'2018-05-07','2018-06-07','2018-06-17');
    insert into cierre values(2018,5,2 ,'2018-05-09','2018-06-09','2018-06-17');
    insert into cierre values(2018,5,3 ,'2018-05-15','2018-06-15','2018-06-25');
    insert into cierre values(2018,5,4 ,'2018-05-25','2018-06-25','2018-06-30');
    insert into cierre values(2018,5,5 ,'2018-05-18','2018-06-18','2018-06-27');
    insert into cierre values(2018,5,6 ,'2018-05-23','2018-06-23','2018-06-30');
    insert into cierre values(2018,5,7 ,'2018-05-29','2018-06-29','2018-07-07');
    insert into cierre values(2018,5,8 ,'2018-05-30','2018-06-30','2018-07-07');
    insert into cierre values(2018,5,9 ,'2018-05-17','2018-06-17','2018-06-27');
    insert into cierre values(2018,6,0 ,'2018-06-05','2018-07-05','2018-07-15');
    insert into cierre values(2018,6,1 ,'2018-06-07','2018-07-07','2018-07-17');
    insert into cierre values(2018,6,2 ,'2018-06-09','2018-07-09','2018-07-19');
    insert into cierre values(2018,6,3 ,'2018-06-15','2018-07-15','2018-07-25');
    insert into cierre values(2018,6,4 ,'2018-06-18','2018-07-18','2018-07-28');
    insert into cierre values(2018,6,5 ,'2018-06-20','2018-07-20','2018-07-30');
    insert into cierre values(2018,6,6 ,'2018-06-22','2018-07-22','2018-08-04');
    insert into cierre values(2018,6,7 ,'2018-06-26','2018-07-26','2018-08-07');
    insert into cierre values(2018,6,8 ,'2018-06-29','2018-07-29','2018-08-07');
    insert into cierre values(2018,6,9 ,'2018-06-30','2018-07-30','2018-08-07');
    insert into cierre values(2018,7,0 ,'2018-07-05','2018-08-05','2018-08-15');
    insert into cierre values(2018,7,1 ,'2018-07-07','2018-08-07','2018-08-17');
    insert into cierre values(2018,7,2 ,'2018-07-09','2018-08-09','2018-08-19');
    insert into cierre values(2018,7,3 ,'2018-07-12','2018-08-12','2018-08-22');
    insert into cierre values(2018,7,4 ,'2018-07-17','2018-08-17','2018-08-27');
    insert into cierre values(2018,7,5 ,'2018-07-19','2018-08-19','2018-08-29');
    insert into cierre values(2018,7,6 ,'2018-07-22','2018-08-22','2018-08-30');
    insert into cierre values(2018,7,7 ,'2018-07-25','2018-08-25','2018-09-02');
    insert into cierre values(2018,7,8 ,'2018-07-27','2018-08-27','2018-09-02');
    insert into cierre values(2018,7,9 ,'2018-07-29','2018-08-29','2018-09-02');
    insert into cierre values(2018,8,0 ,'2018-08-02','2018-09-02','2018-09-12');
    insert into cierre values(2018,8,1 ,'2018-08-05','2018-09-05','2018-09-12');
    insert into cierre values(2018,8,2 ,'2018-08-07','2018-09-07','2018-09-17');
    insert into cierre values(2018,8,3 ,'2018-08-09','2018-09-09','2018-09-19');
    insert into cierre values(2018,8,4 ,'2018-08-12','2018-09-12','2018-09-22');
    insert into cierre values(2018,8,5 ,'2018-08-15','2018-09-15','2018-09-26');
    insert into cierre values(2018,8,6 ,'2018-08-17','2018-09-17','2018-09-27');
    insert into cierre values(2018,8,7 ,'2018-08-19','2018-09-19','2018-09-29');
    insert into cierre values(2018,8,8 ,'2018-08-22','2018-09-22','2018-09-30');
    insert into cierre values(2018,8,9 ,'2018-08-27','2018-09-27','2018-10-02');
    insert into cierre values(2018,9,0 ,'2018-09-02','2018-10-02','2018-10-12');
    insert into cierre values(2018,9,1 ,'2018-09-05','2018-10-05','2018-10-15');
    insert into cierre values(2018,9,2 ,'2018-09-07','2018-10-07','2018-10-17');
    insert into cierre values(2018,9,3 ,'2018-09-09','2018-10-09','2018-10-19');
    insert into cierre values(2018,9,4 ,'2018-09-12','2018-10-12','2018-10-22');
    insert into cierre values(2018,9,5 ,'2018-09-15','2018-10-15','2018-10-25');
    insert into cierre values(2018,9,6 ,'2018-09-17','2018-10-17','2018-10-25');
    insert into cierre values(2018,9,7 ,'2018-09-19','2018-10-19','2018-10-25');
    insert into cierre values(2018,9,8 ,'2018-09-22','2018-10-22','2018-10-30');
    insert into cierre values(2018,9,9 ,'2018-09-25','2018-10-25','2018-10-30');
    insert into cierre values(2018,10,0,'2018-10-02','2018-11-02','2018-11-12');
    insert into cierre values(2018,10,1,'2018-10-05','2018-11-05','2018-11-15');
    insert into cierre values(2018,10,2,'2018-10-07','2018-11-07','2018-11-17');
    insert into cierre values(2018,10,3,'2018-10-09','2018-11-09','2018-11-19');
    insert into cierre values(2018,10,4,'2018-10-12','2018-11-12','2018-11-22');
    insert into cierre values(2018,10,5,'2018-10-15','2018-11-15','2018-11-25');
    insert into cierre values(2018,10,6,'2018-10-17','2018-11-17','2018-11-27');
    insert into cierre values(2018,10,7,'2018-10-19','2018-11-19','2018-11-29');
    insert into cierre values(2018,10,8,'2018-10-22','2018-11-22','2018-11-29');
    insert into cierre values(2018,10,9,'2018-10-25','2018-11-25','2018-12-05');
    insert into cierre values(2018,11,0,'2018-11-02','2018-12-02','2018-12-12');
    insert into cierre values(2018,11,1,'2018-11-05','2018-12-05','2018-12-15');
    insert into cierre values(2018,11,2,'2018-11-07','2018-12-07','2018-12-17');
    insert into cierre values(2018,11,3,'2018-11-09','2018-12-09','2018-12-19');
    insert into cierre values(2018,11,4,'2018-11-12','2018-12-12','2018-12-22');
    insert into cierre values(2018,11,5,'2018-11-15','2018-12-15','2018-12-25');
    insert into cierre values(2018,11,6,'2018-11-17','2018-12-17','2018-12-27');
    insert into cierre values(2018,11,7,'2018-11-19','2018-12-19','2018-12-29');
    insert into cierre values(2018,11,8,'2018-11-22','2018-12-22','2018-12-30');
    insert into cierre values(2018,11,9,'2018-11-25','2018-12-25','2019-01-02');
    insert into cierre values(2018,12,0,'2018-12-02','2019-01-02','2019-01-12');
    insert into cierre values(2018,12,1,'2018-12-05','2019-01-05','2019-01-15');
    insert into cierre values(2018,12,2,'2018-12-07','2019-01-07','2019-01-17');
    insert into cierre values(2018,12,3,'2018-12-09','2019-01-09','2019-01-19');
    insert into cierre values(2018,12,4,'2018-12-12','2019-01-12','2019-01-22');
    insert into cierre values(2018,12,5,'2018-12-15','2019-01-15','2019-01-25');
    insert into cierre values(2018,12,6,'2018-12-17','2019-01-17','2019-01-27');
    insert into cierre values(2018,12,7,'2018-12-19','2019-01-19','2019-01-29');
    insert into cierre values(2018,12,8,'2018-12-22','2019-01-22','2019-01-29');
    insert into cierre values(2018,12,9,'2018-12-25','2019-01-25','2019-02-02');
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

func AutorizarCompra(db *sql.DB) {
	rows, err := db.Query(`create or replace function autorizar_compra(nro_tarjeta char(16), cod_seguridad char(4), nrocomercio integer, monto decimal(7,2)) returns boolean as $$
declare
	autorizar record;
	pendientes decimal;
begin
	select * into autorizar from tarjeta t where t.nrotarjeta = nro_tarjeta and t.estado= 'vigente';
	if not found then
		insert into rechazo values(default, nro_tarjeta, nrocomercio, current_timestamp, monto,'?tarjeta no valida o no vigente');
	else
		select * into autorizar from tarjeta t where t.codseguridad = cod_seguridad;
		if not found then
			insert into rechazo values(default, nro_tarjeta, nrocomercio,current_timestamp, monto, '?codigo de seguridad invalido');
	    else
  
			select sum(c.monto) as deuda into autorizar from compra c where c.nrotarjeta=nro_tarjeta and c.pagado=false;
			pendientes:=autorizar.deuda;

			if pendientes+monto > (select limitecompra from tarjeta t where t.nrotarjeta=nro_tarjeta) then
				insert into rechazo values(default, nro_tarjeta,nrocomercio,current_timestamp,monto,'?supera limite de tarjeta');
			else
				select * into autorizar from tarjeta t where t.nrotarjeta=nro_tarjeta and t.estado='anulada';
				if found then
					insert into rechazo values(default, nro_tarjeta,nrocomercio,current_timestamp,monto,'?plazo de vigencia expirado');
				else
					select * into autorizar from tarjeta t where t.nrotarjeta=nro_tarjeta and t.estado='suspendida';
					if found then						
						insert into rechazo values(default, nro_tarjeta,nrocomercio,current_timestamp,monto,'?la tarjeta se encuentra suspendida');
					else 	
						insert into compra values(default, nro_tarjeta,nrocomercio,current_timestamp, monto,false);
						return true;		
					end if;
				end if;
			end if;
		end if;
	end if;	
return false;
end;
$$language plpgsql; `)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}

func LeerDatosUsuario(db *sql.DB) {
	var opcion int
	var salir = false
	fmt.Printf("************************ MENU *****************\n")
	fmt.Printf("*	1. Crear la database				 	   \n")
	fmt.Printf("*	2. Crear las tablas						   \n")
	fmt.Printf("*	3. Agregar las Primary Keys				   \n")
	fmt.Printf("*	4. Agregar las Foreign Keys				   \n")
	fmt.Printf("*	5. Insertar los datos en las tablas		   \n")
	fmt.Printf("* 	6. Eliminar las Primary Keys y Foreign Keys\n")
	fmt.Printf("*	7. Salir de la aplicacion				   \n")
	fmt.Printf("***********************************************\n")

	for !salir {
		fmt.Scanf("%d", &opcion)
		switch opcion {
			case 1:
				CrearDB()
				fmt.Printf("Creando database ... \n")
			case 2:
				CrearTablas(db)
				fmt.Printf("Creando las tablas ...\n")
			case 3:
				AgregarPKs(db)
				fmt.Printf("Agregando las Primary keys ...\n")
			case 4:
				AgregarFKs(db)
				fmt.Printf("Agregando las Foreign Keys ... \n")
			case 5:
				InsertarDatos(db)
				fmt.Printf("Insertando datos ...\n")
			case 6:
				eliminarFKs(db)
				eliminarPKs(db)
				fmt.Printf("Eliminando PKs y FKS ...\n")
			case 7:
			salir = true
			default:
			fmt.Printf("Solo opciones entre 1 y 7\n")
		}
	}
}


func main() {
	db, err := sql.Open("postgres", "user = postgres dbname = tp2 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	LeerDatosUsuario(db)
}
