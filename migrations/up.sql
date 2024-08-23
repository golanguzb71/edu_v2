CREATE TABLE IF NOT EXISTS USERS
(
    ID           SERIAL PRIMARY KEY,
    PHONE_NUMBER VARCHAR NOT NULL,
    FULL_NAME    VARCHAR NOT NULL,
    CHAT_ID      INT   NOT NULL,
    CREATED_AT   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS COLLECTIONS
(
    ID            SERIAL PRIMARY KEY,
    TITLE         VARCHAR NOT NULL,
    QUESTIONS_URL VARCHAR NOT NULL,
    CREATED_AT    TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS USER_COLLECTION
(
    ID           SERIAL PRIMARY KEY,
    USER_ID      INT REFERENCES USERS (ID) NOT NULL,
    ANSWER_FIELD TEXT[],
    TRUE_COUNT   INT,
    FALSE_COUNT  INT
);

CREATE TABLE IF NOT EXISTS GROUPS
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR NOT NULL,
    teacher_name VARCHAR NOT NULL,
    level        VARCHAR NOT NULL CHECK ( level IN (
                                                    'BEGINNER',
                                                    'ELEMENTARY',
                                                    'PRE_INTERMEDIATE',
                                                    'INTERMEDIATE',
                                                    'UPPER_INTERMEDIATE',
                                                    'ADVANCED',
                                                    'PROFICIENT'
        )),
    created_at   TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS ANSWERS
(
    ID            SERIAL PRIMARY KEY,
    COLLECTION_ID INT REFERENCES COLLECTIONS (ID) NOT NULL,

    ANSWER_FIELD  TEXT[]
);