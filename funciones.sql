create or replace function autorizar_compra(nro_tarjeta char(16), cod_seguridad char(4), nrocomercio integer, monto decimal(7,2)) returns boolean as $$
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
$$language plpgsql;


create or replace function alertar_clientes_1min(nro_tarjeta char(16)) returns  text as $$
	declare 
	alertar record;
	begin
--		if not found then
--			insert into alerta values(default,nro_tarjeta, current_timestamp, null, 1,'?dos compras dentro de un minuto');
--		
		select * into alertar from compra c1
		end if;
	end;
$$language plpgsql;

create or replace function alertar_clientes_5min(nrotarjeta char(16)) returns void as $$
declare
begin

end;

$$language plpgsql;

-- si una tarjeta registra dos rechazos por exceso de limite en el mismo dia,
--la tarjeta debe ser suspendida preventivamente
create or replace function alertar_cliente_rechazos(nro_tarjeta char(16)) returns void as $$
declare
	alertar record;

begin
	select * into alertar from rechazo r where r.nrotarjeta=nro_tarjeta and r.f
	--
--	update from tarjeta set estado='suspendida' where ...;

	insert into alerta values();
	
end;

$$language plpgsql;



create or replace function generar_resumen(cliente integer,a integer, m integer) returns void as $$
declare
	numerotarjeta text;
	tertarjeta text;
	resultado record;
	datoscliente record;
	totalcompra decimal;
	cantidadproductos int;
	datoscomercio record;
	i int;
	resumen int;
	
begin
	--obtengo numero de tarjeta
	select t.nrotarjeta into numerotarjeta from tarjeta t,cliente cl where t.nrocliente=cl.nrocliente and cliente=cl.nrocliente;
	--obtengo la terminacion de la tarjeta
	select substring(numerotarjeta,16) into tertarjeta from tarjeta t where t.nrotarjeta=numerotarjeta;
	
	select * into resultado from cierre c where tertarjeta=cast(c.terminacion as char(10)) and c.anio=a and c.mes=m;
	
	--obtengo datos del cliente
	select nombre,apellido,domicilio into datoscliente from cliente c,tarjeta t where cliente=c.nrocliente and t.nrocliente=cliente and t.nrotarjeta=numerotarjeta;
		
	--cantidad de productos comprados
	select count(nrooperacion) into cantidadproductos from compra c where c.nrotarjeta=numerotarjeta;
	
	--total a pagar
	for i in 1..cantidadproductos loop
		select sum(co.monto) into totalcompra from compra co;	
	end loop;
	
	insert into cabecera values (default,datoscliente.nombre,datoscliente.apellido,datoscliente.domicilio,
		numerotarjeta,resultado.fechainicio,resultado.fechacierre,resultado.fechavto,totalcompra);
		
	select nroresumen into resumen from cabecera;
	
	for i in 1..cantidadproductos loop
		select c.nombre,co.fecha,co.monto into datoscomercio from comercio c,compra co 
			where co.nrooperacion=i and c.nrocomercio=co.nrocomercio and co.nrotarjeta=numerotarjeta;
		insert into detalle values(resumen,i,datoscomercio.fecha,datoscomercio.nombre,datoscomercio.monto);
	end loop;	
	
end;
$$language plpgsql;

/*

create or replace function generar_alerta()returns trigger as $$
begin
	insert into alerta values(new.nroalerta, new.nrotarjeta, new.fecha, new.nrorechazo, new.codalerta, new.descripcion);
	return new;
end;
$$language plpgsql;

create trigger generar_alerta_trigger
instead of insert on rechazo
for each row
execute procedure generar_alerta()
/*
