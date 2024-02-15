create table loyalty_community_policies (
    id serial not null ,
    loyalty_program_id int not null ,
    name varchar not null ,
    created_at varchar not null ,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint loyaltyCommunityPolicies_pk primary key (id),
    constraint loyaltyCommunityPolicies_loyaltyPrograms_fk foreign key (loyalty_program_id) references loyalty_programs(id)
)