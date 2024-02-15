create table tier_managements (
    id serial not null,
    name varchar not null,
    min_point bigint not null,
    max_point bigint not null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint tierManagements_pk primary key (id)
)