CREATE TABLE IF NOT EXISTS users
(
    user_id      BIGSERIAL,
    tg_user_id   TEXT                     NOT NULL,
    phone_number TEXT                     NOT NULL,
    tg_alias     TEXT                     NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id),
    UNIQUE (tg_user_id)
);

CREATE TABLE IF NOT EXISTS chats
(
    chat_id     BIGSERIAL,
    tg_chat_id  TEXT                     NOT NULL,
    description TEXT                     NOT NULL DEFAULT '',
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (chat_id),
    UNIQUE (tg_chat_id)
);

CREATE TABLE IF NOT EXISTS debts
(
    collector_id BIGSERIAL,
    debtor_id    BIGSERIAL,
    chat_id      BIGSERIAL,
    amount       BIGINT                            DEFAULT 0,
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (collector_id, debtor_id, chat_id)
);

CREATE TABLE IF NOT EXISTS events
(
    event_id     BIGSERIAL,
    collector_id BIGSERIAL,
    debtor_id    BIGSERIAL,
    chat_id      BIGSERIAL,
    amount       BIGINT                            DEFAULT 0,
    description  TEXT                     NULL     DEFAULT '',
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (event_id)
);


