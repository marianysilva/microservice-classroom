BEGIN;

CREATE TABLE classrooms
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    code            varchar         NOT NULL,
    course_uuid     uuid            NOT NULL,
    subject_uuid    uuid            NULL,
    name            varchar         NOT NULL,
    description     text            NULL,
    can_subscribe   bool            DEFAULT false,
    format          varchar         DEFAULT 'online',
    starts_at       timestamp       NOT NULL,
    ends_at         timestamp,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);

COMMENT ON COLUMN classrooms.deleted_at IS 'Timestamp indicating when a classroom was softly deleted, allowing for data recovery. A NULL value means the classroom is active.';

CREATE UNIQUE INDEX classrooms_uuid_uindex
    ON classrooms (uuid);
CREATE UNIQUE INDEX classrooms_code_uindex
    ON classrooms (code, deleted_at) NULLS NOT DISTINCT;

COMMIT;
