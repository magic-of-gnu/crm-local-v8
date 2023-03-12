INSERT INTO public.attendance_values (id,value,name,description,created_at,updated_at) VALUES
	 ('856bd4bc-cd52-4ca9-baee-10b1e78c1cdb', 1,'visited','Visited the lesson',now(),now()),
	 ('fc850071-50d7-4c63-85a2-87e9769f6f1f', 2,'late','Visited but came late',now(),now()),
	 ('4195f773-e32e-4c8a-a35b-67aa53421188', 3,'excused','Absent with an excuse',now(),now()),
	 ('a0046fd8-44b9-44be-9f5d-bd5339dbab5f', 4,'absent','Absent with no excuse',now(),now()),
	 ('8e0334c3-b2a8-464f-aaf4-27b06cbbdc32', 5,'class did not start','Absent with no excuse',now(),now());


INSERT INTO public.payment_statuses (id, name,description,created_at,updated_at) VALUES
	 ('bc8c944e-bcd7-4f7c-88fc-ee41733470e3', 'Paid','Payment has been made',now(), now()),
	 ('1cb5a948-04c4-4820-9588-25cdab29f821', 'Waiting For Payment','Invoice has been sent, wating for payment',now(), now()),
	 ('f1ec986e-4604-46cb-8e4d-cc6f21966db2', 'Missed Payment','Missed payment, for reference Missing payment is 5 days after sending invoice',now(), now()),
	 ('f6db5964-1552-4e5c-807d-62dcac4f46f4', 'Idle','When invoice has not been made',now(), now());



INSERT INTO public.centres (id,name,description,created_at,updated_at) VALUES
	 ('a7aa2031-6184-4a07-baa9-af14bb871a24','Easy Academy','First Centre on Sauran Street','2023-02-14 00:00:00.000','2023-02-14 00:00:00.000');

INSERT INTO public.rooms (id,centre_id,name,num_seats,info,created_at,updated_at) VALUES
	 ('3ea97eb3-6791-46de-a9a8-738462e41d4d', 'a7aa2031-6184-4a07-baa9-af14bb871a24','Rico','0','no additional',now(), now()),
	 ('250ac123-7f1d-4746-8512-7ae8eba80200', 'a7aa2031-6184-4a07-baa9-af14bb871a24','Lucik','6','no additional',now(), now()),
	 ('b3a80867-fdf6-45dc-85c5-c253d6d5b5ab', 'a7aa2031-6184-4a07-baa9-af14bb871a24','Loona','6','no additional',now(), now()),
	 ('af219e63-af0d-41cf-86e6-7ee681280359', 'a7aa2031-6184-4a07-baa9-af14bb871a24','Kiko','6','no additional',now(), now());
