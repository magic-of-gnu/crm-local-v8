INSERT INTO public.attendance_values (id,value,name,description,created_at,updated_at) VALUES
	 (0, 1,'visited','Visited the lesson',now(),now()),
	 (1, 2,'late','Visited but came late',now(),now()),
	 (2, 3,'excused','Absent with an excuse',now(),now()),
	 (3, 4,'absent','Absent with no excuse',now(),now()),
	 (4, 5,'class did not start','Absent with no excuse',now(),now());



INSERT INTO public.payment_statuses (id, name,description,created_at,updated_at) VALUES
	 (0, 'Paid','Payment has been made',now(), now()),
	 (1, 'Waiting For Payment','Invoice has been sent, wating for payment',now(), now()),
	 (2, 'Missed Payment','Missed payment, for reference Missing payment is 5 days after sending invoice',now(), now()),
	 (3, 'Idle','When invoice has not been made',now(), now());



INSERT INTO public.centres (id,name,description,created_at,updated_at) VALUES
	 (0,'Easy Academy','First Centre on Sauran Street','2023-02-14 00:00:00.000','2023-02-14 00:00:00.000');

INSERT INTO public.rooms (id,centre_id,name,num_seats,info,created_at,updated_at) VALUES
	 (0, 0,'Rico','0','no additional',now(), now()),
	 (1, 0,'Lucik','6','no additional',now(), now()),
	 (2, 0,'Loona','6','no additional',now(), now()),
	 (3, 0,'Kiko','6','no additional',now(), now());