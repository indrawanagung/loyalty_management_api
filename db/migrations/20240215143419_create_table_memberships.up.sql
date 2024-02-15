create table memberships (
    id serial not null,
    name varchar not null,
    email varchar not null,
    phone varchar not null,
    birth_date date not null,
    address varchar,
    join_date date not null ,
    referral varchar,
    earned_point bigint,
    redeemed_point bigint,
    remained_point bigint,
    status_id int not null,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint memberships_pk primary key (id),
    constraint memberships_status_fk foreign key (status_id) references status(id)
)