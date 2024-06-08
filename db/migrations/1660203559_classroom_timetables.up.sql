BEGIN;

CREATE TABLE classroom_timetables
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    user_uuid       uuid            NOT NULL,
    classroom_id    bigint          NOT NULL REFERENCES classrooms (id),
    location        varchar         NULL,
    week_day        varchar         NULL,
    starts_at       timestamp       NULL,
    ends_at         timestamp       NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);


COMMENT ON COLUMN classroom_timetables.deleted_at IS 'Timestamp indicating when a classroom timetable was softly deleted, allowing for data recovery. A NULL value means the classroom timetable is active.';

CREATE UNIQUE INDEX classroom_timetables_uuid_uindex
    ON classroom_timetables (uuid);

COMMIT;
