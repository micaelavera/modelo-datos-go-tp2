
create or replace function autorizar_compra(nrotarjeta_c char(16),codseguridad_c
char(4),nrocomercio integer, monto decimal(7,2)) returns boolean as $$
declare
vigente record;
codsegu record;
begin
	select * into vigente from tarjeta where tarjeta.nrotarjeta=nrotarjeta_c;
	if not found then
		insert into rechazo values(nrorechazo, nrotarjeta_c, nrcomercio, fecha, monto, '?tarjeta no valida o no vigente');


	else
		select * into codsegu from tarjeta where tarjeta.codseguridad=codseguridad_c;
		if not found then
		insert into rechazo values(new.nrorechazo, nrotarjeta_c, nrcomercio, new.fecha, monto, '?codigo de seguridad invalido');

		end if;
    else
        select * into 


	end if;
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
