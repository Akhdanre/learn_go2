create table customer (
    id varchar(100) not null primary key,
    name varchar(100) not null
);
alter table customer
add column email varchar(100);
alter table customer
add column balance int DEFAULT 0,
    add column rating double default 0.0,
    add column created_at timestamp default CURRENT_TIMESTAMP,
    add column birth_data DATE;
alter table customer
add married boolean default false;
desc customer;
insert into customer (
        id,
        name,
        email,
        balance,
        rating,
        birth_data,
        married
    )
values (
        'dani',
        'dani',
        'dani@gmail.com',
        1000000,
        90.0,
        '2003-08-12',
        false
    );
insert into customer (
        id,
        name,
        email,
        balance,
        rating,
        birth_data,
        married
    )
values (
        'oukenzeumasio',
        'oukenzuemasio',
        'oukenzuemasio@gmail.com',
        1000000,
        90.0,
        '2004-05-12',
        false
    ),
    (
        'akrida',
        'akrida',
        'akrida@gmail.com',
        1000000,
        90.0,
        '2004-05-18',
        false
    );
select id,
    name,
    email,
    balance,
    rating,
    birth_data,
    married,
    created_at
from customer;
delete from customer
where id = 'robbani';

insert into customer (
        id,
        name,
        email,
        balance,
        rating,
        birth_data,
        married
    )
values (
        'lofta',
        'lofta',
        NULL,
        1000000,
        90.0,
        NULL,
        true
    );