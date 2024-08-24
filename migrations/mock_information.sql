INSERT INTO USERS (PHONE_NUMBER, FULL_NAME, CHAT_ID)
VALUES ('123-456-7890', 'John Doe', 123456789.0),
       ('234-567-8901', 'Jane Smith', 234567890.1),
       ('345-678-9012', 'Alice Johnson', 345678901.2),
       ('456-789-0123', 'Bob Brown', 456789012.3),
       ('567-890-1234', 'Charlie Davis', 567890123.4);

INSERT INTO COLLECTIONS (TITLE, QUESTIONS_URL)
VALUES ('Math Quiz', 'https://example.com/math-quiz'),
       ('Science Quiz', 'https://example.com/science-quiz'),
       ('History Quiz', 'https://example.com/history-quiz'),
       ('Geography Quiz', 'https://example.com/geography-quiz'),
       ('Literature Quiz', 'https://example.com/literature-quiz');

INSERT INTO USER_COLLECTION (USER_ID, ANSWER_FIELD, collection_id)
VALUES (1, '{"A", "B", "C"}', 2),
       (2, '{"D", "E", "F"}', 3),
       (3, '{"G", "H", "I"}', 1),
       (4, '{"J", "K", "L"}', 2),
       (5, '{"M", "N", "O"}', 3);

INSERT INTO GROUPS (name, teacher_name, level, start_time, started_date, days_week)
VALUES ('Group A', 'Mr. Adams', 'BEGINNER' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group B', 'Ms. Baker', 'ELEMENTARY' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group C', 'Dr. Clark', 'INTERMEDIATE' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group D', 'Mrs. Davis', 'ADVANCED' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group E', 'Prof. Edwards', 'PROFICIENT' , '12:00','2024-12-12','ODD_DAYS');

INSERT INTO ANSWERS (COLLECTION_ID, ANSWER_FIELD)
VALUES (1, '{"A", "B", "C"}'),
       (2, '{"D", "E", "F"}'),
       (3, '{"G", "H", "I"}'),
       (4, '{"J", "K", "L"}'),
       (5, '{"M", "N", "O"}');
