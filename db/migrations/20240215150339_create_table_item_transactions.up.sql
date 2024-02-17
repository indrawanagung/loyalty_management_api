create table item_transactions (
    id serial not null,
    member_id int not null,
    total_amount bigint not null ,
    transaction_date date,
    transaction_id varchar null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint itemTransaction_pk primary key (id),
    constraint itemTransaction_memberships_fk foreign key (member_id) references memberships(id)
)