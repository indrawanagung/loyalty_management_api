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
INSERT INTO public.memberships
(id, "name", email, "password", phone, birth_date, address, join_date, referral, earned_point, redeemed_point, remained_point, status_id, created_at, updated_at, deleted_at)
VALUES(2, 'ruby', 'ruby@gmail.com', '$2a$04$hTKbPMJUOWox.gd1Of4NqOtsML0pD02Va9CmlCMP4W2tqLOlyTIDG', '+14155552671', '2021-11-22', 'jalan kapuas', '2024-02-17', 'b2be4eb3-8954-420b-8241-38174162c6fa', 0, 0, 0, 1, '1708142034', '', NULL);

-- item transaction
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(1, 1, 100000, '2024-02-16', '1708085310', '1708085310', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(2, 1, 100000, '2024-02-17', '1708100599', '1708100599', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(3, 1, 100000, '2024-02-17', '1708100674', '1708100674', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(4, 1, 100000, '2024-02-17', '1708100718', '1708100718', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(5, 1, 100000, '2024-02-17', '1708100926', '1708100926', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(6, 1, 100000, '2024-02-17', '1708100987', '1708100987', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(7, 1, 100000, '2024-02-17', '1708101029', '1708101029', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(8, 1, 100000, '2024-02-17', '1708101204', '1708101204', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(9, 1, 100000, '2024-02-17', '1708101241', '1708101241', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(10, 1, 100000, '2024-02-17', '1708101253', '1708101253', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(11, 1, 100000, '2024-02-17', '1708123125', '1708123125', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(12, 1, 100000, '2024-02-17', 'TRINV/000000/022024', '1708128261', '', NULL);
INSERT INTO public.item_transactions
(id, member_id, total_amount, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(13, 1, 100000, '2024-02-17', 'TRINV/000013/022024', '1708128370', '', NULL);

-- item transaction detail
INSERT INTO public.item_transaction_details
(id, item_transaction_id, item_name, price, qty, item_subtotal, created_at)
VALUES(1, 1, 'buku golang', 10000, 10, 100000, '1708136593');
INSERT INTO public.item_transaction_details
(id, item_transaction_id, item_name, price, qty, item_subtotal, created_at)
VALUES(2, 2, 'buku golang', 10000, 10, 100000, '1708136598');


-- add activity member
INSERT INTO public.add_member_activity_histories
(id, member_id, activity_name, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(5, 1, 'add activity', '2024-02-17', 'TRACT-000005-022024', '1708131880', '', NULL);
INSERT INTO public.add_member_activity_histories
(id, member_id, activity_name, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(9, 1, 'add activity', '2024-02-17', 'TRACT-000009-022024', '1708132061', '', NULL);
INSERT INTO public.add_member_activity_histories
(id, member_id, activity_name, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(10, 1, 'add activity', '2024-02-17', 'TRACT-000010-022024', '1708132062', '', NULL);
INSERT INTO public.add_member_activity_histories
(id, member_id, activity_name, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(11, 1, 'add activity', '2024-02-17', 'TRACT-000011-022024', '1708132159', '', NULL);
INSERT INTO public.add_member_activity_histories
(id, member_id, activity_name, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(12, 1, 'add activity', '2024-02-17', 'TRACT-000012-022024', '1708132245', '', NULL);
INSERT INTO public.add_member_activity_histories
(id, member_id, activity_name, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(13, 1, 'add activity', '2024-02-17', 'TRACT-000013-022024', '1708132266', '', NULL);
INSERT INTO public.add_member_activity_histories
(id, member_id, activity_name, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(14, 1, 'add activity', '2024-02-17', 'TRACT-000014-022024', '1708132533', '', NULL);

-- eaned point
INSERT INTO public.earned_point_histories
(id, member_id, transaction_date, reference_transaction_id, loyalty_program_id, existing_point, earned_point, balance_point, created_at, updated_at, deleted_at)
VALUES(5, 1, '2024-02-17', 'TRINV/000001/17022024', 1, 18500, 10000, 28500, '1708136593', '', NULL);
INSERT INTO public.earned_point_histories
(id, member_id, transaction_date, reference_transaction_id, loyalty_program_id, existing_point, earned_point, balance_point, created_at, updated_at, deleted_at)
VALUES(6, 1, '2024-02-17', 'TRINV/000002/17022024', 1, 28500, 10000, 38500, '1708136598', '', NULL);

-- redeemed point
INSERT INTO public.redeemed_point_histories
(id, member_id, earned_point, redeemed_point, remaining_point, transaction_date, created_at, updated_at, deleted_at)
VALUES(2, 1, 20000, 1500, 18500, '2024-02-17', '1708134420', '', NULL);

-- add member history
INSERT INTO public.add_member_histories
(id, member_id, person_id, transaction_date, transaction_id, created_at, updated_at, deleted_at)
VALUES(1, 1, 2, '2024-02-17', 'TRMGM-000001-022024', '1708142034', '', NULL);


