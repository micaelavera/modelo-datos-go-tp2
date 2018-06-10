drop database if exists tarjeta;
create database tarjeta;

\c tarjeta

create table cliente(
	    nrocliente integer,
		nombre varchar(64), 
		apellido varchar(64),
		domicilio varchar(64), 
		telefono char(12)
);

create table tarjeta(
	    nrotarjeta char(12),
		nrocliente integer,
		validadesde char(6), --e.g 201106
		validahasta char(6),
		codseguridad char(4),
		limitecompra decimal(8,2),
		estado char(10) --'vigente', 'suspendida', 'anulada'
);

create table comercio(
	    nrocomercio integer,
		nombre varchar(64),
		domicilio varchar(64),
		codigopostal char(8),
		telefono char(12)
);

create table compra(
	    nrooperacion integer,
		nrotarjeta char(12),
		nrocomercio integer,
		fecha timestamp,
		monto decimal(7,2);
		pagado boolean
);
