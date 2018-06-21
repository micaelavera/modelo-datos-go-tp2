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
	anio        integer,
	mes         integer,
	terminacion integer,
	fechainicio date,
	fechavto    date
);

create table cabecera(
	nroresumen serial,
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
	nroresumen      serial,
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

-- Esta tabla no es parte del modelo de datos, pero se incluye para
-- poder probar las funciones.

create table consumo(
	nrotarjeta 	    char(16),
	codseguridad	char(4),
	nrocomercio 	integer,
	monto        	decimal(7,2)
);










--PRIMARY KEY
alter table cliente  add constraint cliente_pk   primary key (nrocliente);
alter table tarjeta  add constraint tarjeta_pk   primary key (nrotarjeta);
alter table comercio add constraint comercio_pk  primary key (nrocomercio);
alter table compra   add constraint compra_pk    primary key (nrooperacion);
alter table rechazo  add constraint rechazo_pk   primary key (nrorechazo);
alter table cierre   add constraint cierre_pk    primary key (anio,mes,terminacion);
--alter table cierre   add constraint cierre_pk    primary key (mes,terminacion);
alter table cabecera add constraint cabecera_pk  primary key (nroresumen);
alter table detalle  add constraint detalle_pk   primary key (nroresumen,nrolinea);
alter table alerta   add constraint alerta_pk    primary key (nroalerta);

--FOREIGN KEY
alter table tarjeta  add constraint tarjeta_fk0 foreign key (nrocliente)  references cliente  (nrocliente);
alter table compra   add constraint compra_fk0  foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
alter table compra   add constraint compra_fk1  foreign key (nrocomercio) references comercio (nrocomercio);
--alter table rechazo  add constraint rechazo_fk0 foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
alter table rechazo  add constraint rechazo_fk1 foreign key (nrocomercio) references comercio (nrocomercio);
alter table cabecera add constraint cabecera_fk foreign key (nrotarjeta)  references tarjeta  (nrotarjeta);
alter table alerta   add constraint alerta_fk0  foreign key (nrotarjeta)  references tarjeta (nrotarjeta);

--INSERCION DE DATOS

--CLIENTES
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
insert into comercio values(1,'Anubis','Av. Pres. Juan Domingo Peron 3497','1613','4463-5343');
insert into comercio values(2,'Si A La Pizza','25 de Mayo 2502','1613','4463-2314');
insert into comercio values(3,'Narrow','Av. Pres. Juan Domingo Peron 1420','1663','4667-7297');
insert into comercio values(4,'Starbucks Coffee','Parana 3745','1640','474898');
insert into comercio values(5,'47 Street','Cruce Ruta 8 y Ruta 202','1613','4667-5770');
insert into comercio values(6,'Frávega','Av. Pres. Juan Domingo Peron 1127','1663','4667-4009');
insert into comercio values(7,'Optica Ivaldi','Av. Pres. Juan Domingo Peron 1645','1663','4667-2332');
insert into comercio values(8,'Farmacity','Av. Constituyentes 6093/99','1617','4587-8243');
insert into comercio values(9,'Cúspide','Cruce Ruta 8 y Ruta 202','1613','1521508092');
insert into comercio values(10,'Garbarino','Av. Dr. Ricardo Balbín 1198','1663',' 4667-6534');
insert into comercio values(11,'Starbucks Coffee','Cruce Ruta 8 y Ruta 202','1613','4667-5434');
insert into comercio values(12,'Bonafide',' Italia 1249','1663','4667-4545');
insert into comercio values(13,'Optica Gris','Av. Arturo Illia 5243','1613','4463-8344');
insert into comercio values(14,'Matu Jean´s','Av. Pres. Juan Domingo ]Peron 3300','1613','4463-9089');
insert into comercio values(15,'Compumundo','Belgrano 1401','1663','4667-3425');
insert into comercio values(16,'Falabella','Parana 3745','1640','4717-8100');
insert into comercio values(17,'McDonald´s','Av. Pres. Juan Domingo Peron 983','1662','4668-0912');
insert into comercio values(18,'M 58','Charlone 1201','1663','4667-4532');
insert into comercio values(19,'Cine Hoyts Unicenter','Parana 3745','1640','4717-8109');
insert into comercio values(20,'Solo Deportes','Av. Pres. Juan Domingo Peron 1317','1663','4667-3453');
