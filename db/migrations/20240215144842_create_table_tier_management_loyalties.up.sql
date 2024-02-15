create table tier_management_loyalties (
    tier_management_id int ,
    loyalty_program_id int,
    constraint tierManagementLoyalties_pk primary key (tier_management_id, loyalty_program_id),
    constraint tierManagementLoyalties_tierManagements_fk foreign key (tier_management_id) references tier_managements(id),
    constraint tierManagementLoyalties_loyaltyPrograms_fk foreign key (loyalty_program_id) references loyalty_programs(id)
)