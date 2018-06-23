
create or replace function autorizar_compra(nro_tarjeta char(16), cod_seguridad
char(4), nrocomercio integer, monto_c decimal(7,2)) returns boolean as $$
declare
	autorizar record;
begin
/*
	select * into autorizar from tarjeta t where t.nrotarjeta = nro_tarjeta and t.estado= 'vigente';
	if not found then
		insert into rechazo values(default, nro_tarjeta, nrocomercio, current_timestamp, monto_c,'?tarjeta no valida o no vigente');
	else
		select * into autorizar from tarjeta t where t.codseguridad = cod_seguridad;
		if not found then
			insert into rechazo values(default, nro_tarjeta, nrocomercio,current_timestamp, monto_c, '?codigo de seguridad invalido');
	    else
	
       	select sum(c.monto) as deuda into autorizar from tarjeta t, compra c where c.nrotarjeta=t.nrotarjeta and c.pagado=false; --falta el monto de la compra actual
			if monto>t.limitecompra then
				insert into rechazo values(default, nro_tarjeta,nrocomercio,current_timestamp,'?supera limite de tarjeta');
			else
				select * into autorizar from tarjeta t where t.nrotarjeta=nro_tarjeta and t.estado='anulada';
				if found then
					insert into rechazo values(default, nro_tarjeta,nrocomercio,current_timestamp,monto_c,'?plazo de vigencia expirado');
				else
					select * into autorizar from tarjeta t where t.nrotarjeta=nro_tarjeta and t.estado='suspendida';
					if found then						
						insert into rechazo values(default, nro_tarjeta,nrocomercio,current_timestamp,monto_c,'?la tarjeta se encuentra suspendida');
					else 
*/					
						insert into compra values(default, nro_tarjeta,nrocomercio,current_timestamp, monto_c, true);
						return true;		
--		end if;
--		end if;
--		end if;
--	end if;
--	end if;	
--	return false;
end;
$$language plpgsql;


create or replace function alertar_clientes_1min(nro_tarjeta char(16)) returns void as $$
	declare 
		alertar record;
	begin
		if not found then
			insert into alerta values(default,nro_tarjeta, current_timestamp, null, 1,'?dos compras dentro de un minuto');
		else
			-- falta la comparacion con el codigopostal
			select * into alertar from compra c1
				where not exists(select 1 from compra c2
					where c1.nrocomercio=c2.nrocomercio and 
							c1.nrotarjeta=nro_tarjeta and 
								c2.nrotarjeta=nro_tarjeta);
		end if;
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
*/
