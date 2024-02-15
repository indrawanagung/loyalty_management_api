create table redeemed_point_histories (
    id serial not null ,
    member_id int not null ,
    earned_point bigint not null,
    redeemed_point bigint not null,
    remaining_point bigint not null,
    transaction_date date not null,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint redeemedPointHistories_pk primary key (id),
    constraint redeemedPointHistories_memberships_fk foreign key (member_id) references memberships(id)
)