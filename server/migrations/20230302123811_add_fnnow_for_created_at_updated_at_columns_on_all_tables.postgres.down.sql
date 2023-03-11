ALTER TABLE public.attendance_values ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.attendance_values ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.attendances ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.attendances ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.centres ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.centres ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.lectures_calendar ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.lectures_calendar ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.courses ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.courses ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.invoices ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.invoices ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.payment_statuses ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.payment_statuses ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.rooms ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.rooms ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.student_courses ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.student_courses ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.students ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.students ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.users ALTER COLUMN created_at DROP DEFAULT;
ALTER TABLE public.users ALTER COLUMN updated_at DROP DEFAULT;

ALTER TABLE public.employees ALTER COLUMN updated_at DROP DEFAULT;
ALTER TABLE public.employees ALTER COLUMN created_at DROP DEFAULT;