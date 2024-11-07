CREATE TABLE IF NOT EXISTS users
(
    user_id      BIGINT,
    phone_number TEXT                     NOT NULL,
    alias        TEXT                     NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS debts
(
    collector_id BIGINT,
    debtor_id    BIGINT,
    chat_id      BIGINT,
    amount       BIGINT                            DEFAULT 0,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (collector_id, debtor_id, chat_id)
);

CREATE TABLE IF NOT EXISTS events
(
    event_id     BIGSERIAL,
    collector_id BIGINT                   NOT NULL,
    debtor_id    BIGINT                   NOT NULL,
    chat_id      BIGINT                   NOT NULL,
    amount       BIGINT                            DEFAULT 0,
    description  TEXT                     NOT NULL DEFAULT '',
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (event_id)
);


