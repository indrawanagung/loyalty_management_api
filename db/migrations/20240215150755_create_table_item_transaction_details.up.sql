create table item_transaction_details (
    id serial not null ,
    item_transaction_id int not null,
    item_name varchar not null,
    price bigint not null ,
    qty int not null,
    item_subtotal bigint,
    created_at varchar not null,
    constraint itemTransactionDetails_pk primary key (id),
    constraint itemTransactionDetails_itemTransactions_fk foreign key (item_transaction_id) references item_transactions(id)

)