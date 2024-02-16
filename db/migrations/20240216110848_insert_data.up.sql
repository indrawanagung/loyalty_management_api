-- status
INSERT INTO public.status
(id, "name", deleted_at)
VALUES(1, 'active', NULL);
INSERT INTO public.status
(id, "name", deleted_at)
VALUES(2, 'inactive', NULL);

-- tier management
INSERT INTO public.tier_managements
(id, "name", min_point, max_point, created_at, updated_at, deleted_at)
VALUES(1, 'silver', 10000, 100000, '1708075706', '', NULL);
INSERT INTO public.tier_managements
(id, "name", min_point, max_point, created_at, updated_at, deleted_at)
VALUES(2, 'gold', 100001, 1000000, '1708075751', '', NULL);
INSERT INTO public.tier_managements
(id, "name", min_point, max_point, created_at, updated_at, deleted_at)
VALUES(3, 'platinum', 1000001, 10000000, '1708075765', '', NULL);

-- loyalty program
INSERT INTO public.loyalty_programs
(id, "name", created_at, updated_at, deleted_at)
VALUES(1, 'Kodegiri Loyalty', '1708081902', NULL, NULL);

-- loyalty community policy
INSERT INTO public.loyalty_community_policies
(id, loyalty_program_id, "name", fixed_point, created_at, updated_at, deleted_at)
VALUES(1, 1, 'KodegiriClub', 10000, '1708081902', NULL, NULL);

-- loyalty transaction policy
INSERT INTO public.loyalty_transaction_policies
(id, loyalty_program_id, transaction_amount, qty, first_purchase, fixed_point, percentage, deleted_at)
VALUES(1, 1, 50000, 10, 50000, 5000, 10, NULL);

-- tier management loyalty
INSERT INTO public.tier_management_loyalties
(tier_management_id, loyalty_program_id)
VALUES(1, 1);
INSERT INTO public.tier_management_loyalties
(tier_management_id, loyalty_program_id)
VALUES(2, 1);
INSERT INTO public.tier_management_loyalties
(tier_management_id, loyalty_program_id)
VALUES(3, 1);

-- membership
INSERT INTO public.memberships
(id, "name", email, "password", phone, birth_date, address, join_date, referral, earned_point, redeemed_point, remained_point, status_id, created_at, updated_at, deleted_at)
VALUES(1, 'golang', 'golang@gmail.com', '$2a$04$GzTxEOSVmPCCtFYh8.P/0.MiIZlSpy3hNhtrpWD.pjc/t3ToZp8Zi', '+14155552671', '2021-11-22', 'jalan kapuas', '2024-02-16', 'c5ab0b94-be1b-441b-8659-619524ea1626', 0, 0, 0, 1, '1708084936', '', NULL);




