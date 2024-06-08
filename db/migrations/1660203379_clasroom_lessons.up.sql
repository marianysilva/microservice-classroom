BEGIN;

CREATE TABLE classroom_lessons
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    classroom_id    bigint          NOT NULL REFERENCES classrooms (id),
    lessons_uuid    uuid            NOT NULL,
    starts_at       timestamp       DEFAULT now() NOT NULL,
    ends_at         timestamp,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

COMMENT ON COLUMN classroom_lessons.deleted_at IS 'Timestamp indicating when a classroom lesson was softly deleted, allowing for data recovery. A NULL value means the classroom lesson is active.';

CREATE UNIQUE INDEX classroom_lessons_uuid_uindex
    ON classroom_lessons (uuid);

COMMIT;
