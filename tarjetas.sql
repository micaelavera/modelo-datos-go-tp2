drop database if exists tarjeta;
create database tarjeta;

\c tarjeta

create table cliente(
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
	nrotarjeta   char(12),
	nrocomercio  integer,
	fecha        timestamp,
	monto        decimal(7,2);
	pagado       boolean
);

create table rechazo(
	nrorechazo  integer,
	nrotarjeta  char(12),
	nrocomercio integer,
	fecha       timestamp,
	monto       decimal(7,2),
	motivo      varchar(64)
);



create table cierre(
	año         integer,
	mes         integer,
	terminacion integer,
	fechainicio date,
	fechavto    date
);


create table cabecera(
	nroresume  integer,
	nombre     varchar(64),
	apellido   varchar(64),
	domicilio  varchar(64),
	nrotarjeta varchar(64),
	desde      date,
	hasta      date,
	vence      date,
	total      decimal(8,2)
);



create table detalle (
	nroresumen      integer,
	nrolinea        integer,
	fecha           date,
	nombrecomercio  varchar(64),
	monto           decimal(7,2)
);









--PRIMARY KEY
alter table add constraint cliente_pk   primary key (nrocliente);
alter table add constraint tarjeta_pk   primary key (nrotarjeta);
alter table add constraint comercio_pk  primary key (nrocomercio);
alter table add constraint compra_pk    primary key (nrooperacion);
alter table add constraint rechazo_pk   primary key (nrorechazo);
alter table add constraint cierre_pk0   primary key (año);
alter table add constraint cierre_pk1   primary key (mes);
alter table add constraint cierre_pk2   primary key (terminacion);
alter table add constraint cabecera_pk  primary key (nroresumen);
alter table add constraint detalle_pk0  primary key (nroresumen);
alter table add constraint cabecera_pk1 primary key (nrolinea);

--FOREIGN KEY
alter table add constraint tarjeta_fk0 foreign key (nrocliente)  references 
cliente (nrocliente);
alter table add constraint compra_fk0  foreign key (nrotarjeta)  references
tarjeta (nrotarjeta);
alter table add constraint compra_fk1  foreign key (nrocomercio) references
comerio (nrocomercio);
alter table add constraint rechazo_fk0 foreign key (nrotarjeta)  references  tarjeta(nrotarjeta);
alter table add constraint rechazo_fk1 foreign key (nrocomercio) references  comercio(nrocomercio);
alter table add constraint cabecera_fk foreign key (nrotarjeta)  references  tarjeta(nrotarjeta);

