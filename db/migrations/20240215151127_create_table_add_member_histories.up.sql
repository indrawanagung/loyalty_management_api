create table add_member_histories (
    id serial not null ,
    member_id int not null ,
    person_id int not null,
    transaction_date date not null,
    transaction_id varchar null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint addMemberHistories_pk primary key (id),
    constraint addMemberHistories_membershipMember_fk foreign key (member_id) references memberships(id),
    constraint addMemberHistories_membershipPerson_fk foreign key (person_id) references memberships(id)
)