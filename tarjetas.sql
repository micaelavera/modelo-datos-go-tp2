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
	monto        decimal(7,2),
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
	nroresumen  integer,
	nombre     varchar(64),
	apellido   varchar(64),
	domicilio  varchar(64),
	nrotarjeta char(12),
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









--PRIMARY KEY
alter table cliente  add constraint cliente_pk   primary key (nrocliente);
alter table tarjeta  add constraint tarjeta_pk   primary key (nrotarjeta);
alter table comercio add constraint comercio_pk  primary key (nrocomercio);
alter table compra   add constraint compra_pk    primary key (nrooperacion);
alter table rechazo  add constraint rechazo_pk   primary key (nrorechazo);
alter table cierre   add constraint cierre_pk    primary key (año,mes,terminacion);
--alter table cierre add constraint cierre_pk1   primary key (mes);
--alter table cierre  add constraint cierre_pk2   primary key (terminacion);
alter table cabecera add constraint cabecera_pk  primary key (nroresumen);
alter table detalle  add constraint detalle_pk   primary key (nroresumen,nrolinea);
--alter table detalle add constraint cabecera_pk1 primary key (nrolinea);

--FOREIGN KEY
alter table tarjeta  add constraint tarjeta_fk0 foreign key (nrocliente)  references cliente  (nrocliente);
alter table compra   add constraint compra_fk0  foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
alter table compra   add constraint compra_fk1  foreign key (nrocomercio) references comercio (nrocomercio);
alter table rechazo  add constraint rechazo_fk0 foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
alter table rechazo  add constraint rechazo_fk1 foreign key (nrocomercio) references comercio (nrocomercio);
alter table cabecera add constraint cabecera_fk foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);


--INSERCION DE DATOS

--CLIENTES
insert into cliente values(1,'Jorge','Rodriguez','Godoy Cruz 1064','4584-3863');
insert into cliente values(2,'Mercedes', 'Benz', 'Pte Perón 1223','4665-8989');
insert into cliente values(3,'Megan', 'Ocaranza', 'Tribulato 2345', '4500-7651');
insert into cliente values(4,'Luis', 'Rios', 'Dorrego 1234', '4213-0153');
insert into cliente values(5,'Julio', 'Cortazar', ' Av Balbin 534', '4890-8747');
insert into cliente values(6,'Tomás', 'Aguirre', 'Urquiza 540', '4707-1600');
insert into cliente values(7,'José', 'Avalos', 'Av E Perón 7716', '4589 - 3191');
insert into cliente values(8, 'Prudencia', 'Arzuaga','Rivadavia 416');
insert into cliente values(9, 'Damiana', 'Molina', 'Malvinas 890', '4129-0964');
insert into cliente values(10, 'Ramón', 'Perez', 'Las Heras 460', '4556-8970');
insert into cliente values(11, 'Fernando', 'Álvarez', 'Tribulato 1290', '4908-7822');
insert into cliente values(12, 'Carla', 'Estrella', 'Primera Junta 7865');
insert into cliente values(13, 'Farid', 'Hasan', 'España 438', '4123-9078');
insert into cliente values(14, 'Alicia', 'Castillo', 'Rodriguez Peña 170', '4367-7801');
insert into cliente values(15, 'Elsa', 'López', 'Argüero 1138','4563-2323');
insert into cliente values(16, 'Harry', 'Potter', 'Gutiérrez 908', '4768-9475');
insert into cliente values(17, 'Pedro', 'Moreno', ' Julian Rejala 999',' 4556-9872');
insert into cliente values(18, 'Ofelia', 'Le Brun', 'Paunero 7856','4389-7531');
insert into cliente values(19, 'Fiona', 'Las Margaritas 484', '4895-7939');
insert into cliente values(20,'Orlando', 'Bloom', 'Belgrano 1257',' 4664-1640');






--TARJETAS


