create table earned_point_histories (
    id serial not null ,
    member_id int not null ,
    transaction_date date not null,
    reference_transaction_id int not null ,
    loyalty_program_id int not null ,
    existing_point bigint not null ,
    earned_point bigint not null ,
    balance_point bigint not null ,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint earned_point_histories_pk primary key (id),
    constraint earned_point_histories_memberships_fk foreign key (member_id) references memberships(id),
    constraint earned_point_histories_item_transactions foreign key (reference_transaction_id) references item_transactions(id),
    constraint earned_point_histories_loyalty_programs foreign key (loyalty_program_id) references loyalty_programs(id)
)