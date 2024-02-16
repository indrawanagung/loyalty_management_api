create table loyalty_transaction_policies (
    id serial not null ,
    loyalty_program_id int not null ,
    transaction_amount bigint not null,
    qty int not null ,
    first_purchase int not null,
    fixed_point bigint not null ,
    percentage int not null ,
    deleted_at timestamp null,
    constraint loyaltyTransactionPolicies_pk primary key (id),
    constraint loyaltyTransactionPolicies_loyaltyPrograms_fk foreign key (loyalty_program_id) references loyalty_programs(id)
)