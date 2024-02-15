create table status (
    id serial not null ,
    name varchar not null,
    deleted_at timestamp null,
    constraint status_pk primary key (id)

)