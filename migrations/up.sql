CREATE TABLE IF NOT EXISTS USERS
(
    ID           SERIAL PRIMARY KEY,
    PHONE_NUMBER VARCHAR NOT NULL,
    FULL_NAME    VARCHAR NOT NULL,
    CHAT_ID      BIGINT  NOT NULL,
    CREATED_AT   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS COLLECTIONS
(
    ID            SERIAL PRIMARY KEY,
    TITLE         VARCHAR NOT NULL,
    QUESTIONS_URL TEXT[]  NOT NULL,
    CREATED_AT    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS USER_COLLECTION
(
    ID            SERIAL PRIMARY KEY,
    USER_ID       INT REFERENCES USERS (ID)       NOT NULL,
    COLLECTION_ID INT references collections (ID) NOT NULL,
    ANSWER_FIELD  TEXT[],
    CREATED_AT    TIMESTAMP DEFAULT NOW()

);

CREATE TABLE IF NOT EXISTS GROUPS
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR   NOT NULL,
    teacher_name VARCHAR   NOT NULL,
    level        VARCHAR   NOT NULL CHECK ( level IN (
                                                      'BEGINNER',
                                                      'ELEMENTARY',
                                                      'PRE_INTERMEDIATE',
                                                      'INTERMEDIATE',
                                                      'UPPER_INTERMEDIATE',
                                                      'ADVANCED',
                                                      'PROFICIENT'
        )),
    start_time   VARCHAR   NOT NULL,
    started_date TIMESTAMP NOT NULL,
    days_week    VARCHAR   NOT NULL CHECK ( days_week in (
                                                          'EVEN_DAYS',
                                                          'ODD_DAYS',
                                                          'CUSTOM'
        )),
    created_at   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ANSWERS
(
    ID            SERIAL PRIMARY KEY,
    COLLECTION_ID INT REFERENCES COLLECTIONS (ID) NOT NULL UNIQUE,
    ANSWER_FIELD  TEXT[],
    CREATED_AT    TIMESTAMP DEFAULT NOW()
);