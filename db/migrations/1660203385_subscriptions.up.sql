BEGIN;

CREATE TABLE subscriptions
(
    id              bigint          PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    uuid            uuid            DEFAULT uuid_generate_v4() NOT NULL,
    user_uuid       uuid            NOT NULL,
    classroom_id    bigint          NOT NULL REFERENCES classrooms (id),
    role            varchar         NULL,
    expires_at      timestamp       NULL,
    created_at      timestamp       DEFAULT now() NOT NULL,
    updated_at      timestamp       DEFAULT now() NOT NULL,
    deleted_at      timestamp
);


COMMENT ON COLUMN subscriptions.deleted_at IS 'Timestamp indicating when a subscription was softly deleted, allowing for data recovery. A NULL value means the subscription is active.';

CREATE UNIQUE INDEX subscriptions_uuid_uindex
    ON subscriptions (uuid);

COMMIT;
