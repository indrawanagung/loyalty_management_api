create table add_member_activity_histories (
    id serial not null,
    member_id int not null,
    activity_name varchar not null ,
    transaction_date date not null,
    transaction_id varchar not null,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint addMemberActivityHistories_pk primary key (id),
    constraint addMemberActivityHistories_memberships_fk foreign key (member_id) references memberships(id)
)