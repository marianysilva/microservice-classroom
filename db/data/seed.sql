-------------------------------------------------------------------------------------------------------
INSERT INTO public.classrooms (
	code, course_uuid, subject_uuid, "name", description, can_subscribe, "format", starts_at
) VALUES(
	'CRGEO-1', '00000000-0000-0000-0001-111111111111', '00000000-0000-0000-0002-111111111111', 'Classroom GEO 1', 'description - Classroom GEO 1', true, 'online', now()
);

INSERT INTO public.classrooms (
	code, course_uuid, subject_uuid, "name", description, can_subscribe, "format", starts_at
) VALUES(
	'CRGEO-1', '00000000-0000-0000-0001-111111111111', '00000000-0000-0000-0002-222222222222', 'Classroom GEO 2', 'description - Classroom GEO 2', false, 'online', now()
);

INSERT INTO public.classrooms (
	code, course_uuid, "name", description, can_subscribe, "format", starts_at
) VALUES(
	'CRCCO-1', '00000000-0000-0000-0001-222222222222', 'Classroom CCO 1', 'description - Classroom CCO 1', true, 'online', now()
);