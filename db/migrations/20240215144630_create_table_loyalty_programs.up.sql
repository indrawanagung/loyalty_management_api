create table loyalty_programs (
    id serial not null,
    name varchar not null ,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint loyaltyPrograms_pk primary key (id)
)