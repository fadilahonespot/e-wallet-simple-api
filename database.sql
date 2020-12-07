create table customer(
    customer_number varchar(225) not null PRIMARY KEY,
    name varchar(225) not null
);

create table account(
    account_number varchar(225) PRIMARY KEY not null,
    customer_number varchar(225) not null,
    balance numeric(225),
    constraint customer_number FOREIGN KEY(customer_number) REFERENCES customer(customer_number)
);

CREATE SEQUENCE IF NOT EXISTS customer_number_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 10000000000
    START WITH 1001;

CREATE SEQUENCE IF NOT EXISTS account_number_seq
    INCREMENT BY 1
    MINVALUE 1
    MAXVALUE 100000000000
    START WITH 555001;

CREATE VIEW view_detail_customer(account_number, customer_name, balance) as
    SELECT a.account_number,
           c.name,
           a.balance
FROM account a
    JOIN customer c on c.customer_number = a.customer_number;

create function add_customer(name_param character varying, balance_param numeric) returns boolean
    language plpgsql
as
$$
BEGIN
        INSERT INTO customer(customer_number, name)
        VALUES (nextval('customer_number_seq'), name_param);

        INSERT INTO account(account_number, customer_number, balance)
        VALUES (nextval('account_number_seq'), currval('customer_number_seq'), balance_param);

        return true;
    end;
$$;

create function get_customers()
    returns TABLE(account_number character varying, name character varying, balance numeric)
    language plpgsql
as
$$
BEGIN
        return query
        SELECT c.account_number, c.customer_name, c.balance FROM view_detail_customer as c;
    end;
$$;

create function get_detail_customer(account_number_param character varying)
    returns TABLE(account_number character varying, name character varying, balance numeric)
    language plpgsql
as
$$
BEGIN
        return query
        SELECT c.account_number, c.customer_name, c.balance FROM view_detail_customer as c
        WHERE c.account_number = account_number_param;
    end;
$$;

create function is_account_exist(account_number_param character varying) returns boolean
    language plpgsql
as
$$
DECLARE usrid INT;
    BEGIN
        SELECT count(account_number)
              INTO usrid
        FROM account as a
        WHERE a.account_number = account_number_param;

        IF (usrid > 0) THEN
            return true;
        ELSE
            return false;
        end if;
    end;
$$;

create function is_customer_exist(customer_number_param character varying) returns boolean
    language plpgsql
as
$$
DECLARE usrid INT;
    BEGIN
        SELECT count(customer_number)
              INTO usrid
        FROM customer as c
        WHERE c.customer_number = customer_number_param;

        IF (usrid > 0) THEN
            return true;
        ELSE
            return false;
        end if;
    end;
$$;

create function transfer_balance(my_account_number_param character varying, to_account_number_param character varying, amount_param numeric) returns boolean
    language plpgsql
as
$$
DECLARE
        bln numeric;
        bln2 numeric;
    BEGIN
        SELECT balance
        INTO bln
        FROM account as a
        WHERE a.account_number = my_account_number_param;

        IF bln >= amount_param THEN
            UPDATE account as d
                SET balance = bln - amount_param
            WHERE d.account_number = my_account_number_param;
            
            SELECT b.balance
                INTO bln2
            FROM account as b
            WHERE b.account_number = to_account_number_param;
            
            UPDATE account as c
                SET balance = bln2 + amount_param
            WHERE c.account_number = to_account_number_param;
            RETURN true;
        ELSE
            RETURN false;
        end if;
    end;
$$;
