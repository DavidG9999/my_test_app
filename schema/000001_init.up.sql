CREATE TABLE
    users (
        id SERIAL primary key not null unique,
        name varchar(255) not null,
        email varchar(255) not null unique,
        password_hash varchar(255) not null
    );

CREATE TABLE
    autos (
        id serial primary key not null unique,
        brand varchar(255) not null,
        model varchar(255) not null,
        state_number varchar(9) not null unique
    );

CREATE TABLE
    contragents (
        id serial primary key not null unique,
        name varchar(255) not null,
        address varchar(255) not null,
        inn_kpp varchar(20) not null unique
    );

CREATE TABLE
    dispetchers (
        id serial primary key not null unique,
        full_name varchar(255) not null
    );

CREATE TABLE
    drivers (
        id serial primary key not null unique,
        full_name varchar(255) not null,
        license varchar(10) not null unique,
        class varchar(255) not null
    );

CREATE TABLE
    mechanics (
        id serial primary key not null unique,
        full_name varchar(255) not null
    );

CREATE TABLE
    organizations (
        id serial primary key not null unique,
        name varchar(255) not null,
        address varchar(255) not null,
        chief varchar(255) not null,
        financial_chief varchar(255) not null,
        inn_kpp varchar(20) not null unique
    );

CREATE TABLE
    accounts (
        id serial primary key not null unique,
        account_number varchar(20) not null unique,
        bank_name varchar(255) not null,
        bank_id_number varchar(9) not null,
        organization_id int references organizations (id) on delete cascade on update cascade not null
    );

CREATE TABLE
    putlist_headers (
        id serial primary key not null unique,
        user_id int references users (id) on delete cascade on update cascade not null,
        number int not null unique,
        account_id int references accounts (id) on delete cascade on update cascade not null,
        date_with date not null,
        date_for date not null,
        auto_id int references autos (id) on delete cascade on update cascade not null,
        driver_id int references drivers (id) on delete cascade on update cascade not null,
        dispetcher_id int references dispetchers (id) on delete cascade on update cascade not null,
        mechanic_id int references mechanics (id) on delete cascade on update cascade not null
    );

CREATE TABLE
    putlist_bodies (
        id serial primary key not null unique,
        putlist_header_number int references putlist_headers (number) on delete cascade on update cascade not null,
        number int int not null,
        contragent_id int references contragents (id) on delete cascade on update cascade not null,
        item varchar(255) not null,
        time_with timestamp not null,
        time_for timestamp not null
    );
