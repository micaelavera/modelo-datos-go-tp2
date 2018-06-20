
create or replace function autorizar_compra(nro_tarjeta char(16),cod_seguridad
char(4),nrocomercio integer, monto decimal(7,2)) returns boolean as $$
declare
	autorizar record;
	fecha timestamp;
begin
	select * into autorizar from tarjeta t where t.nrotarjeta=nro_tarjeta and t.estado='vigente';
	
	if not found then
		insert into rechazo values(default, nrotarjeta, nrocomercio,current_date, monto,'?tarjeta no valida o no vigente');
	else
		select * into autorizar from tarjeta t where t.codseguridad=cod_seguridad;
		if not found then
			insert into rechazo values(default, nro_tarjeta, nrocomercio,fecha, monto, '?codigo de seguridad invalido');

	    else
        	select sum(monto) as deuda into autorizar from tarjeta t, compra c where c.nrotarjeta=t.nrotarjeta and c.pagado=false; --falta el monto de la compra actual
			if deuda+monto>t.limitecompra then
				insert into rechazo values(nrorechazo,nro_tarjeta,nrocomercio,fecha,'?supera limite de tarjeta');
			
			else
				select * into autorizar from tarjeta t where t.nrotarjeta=nro_tarjeta and t.estado='anulada';
				if not found then
					insert into rechazo values(nrorechazo, nro_tarjeta,nrocomercio,fecha,'?plazo de vigencia expirado');
				else
					select *  into autorizar from tarjeta t where t.nrotarjeta=nro_tarjeta and t.nrotarjeta='suspendida';
					if not found then

						insert into rechazo values(nrorechazo, nro_tarjeta,nrocomercio,fecha,'?la tarjeta se encuentra suspendida');

					else 
						return true;
						insert into compra values(1,nro_tarjeta,nrocomercio,fecha,monto,true);

				end if;
				end if;
			end if;
		end if;
	end if;	
	return false;
end;
$$language plpgsql;

/*
create or replace function generar_factura (nrocliente integer,desde date,hasta date) returns void as $$
begin
	insert into cabecera values()

$$language plpgsql;




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
