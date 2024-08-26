INSERT INTO USERS (PHONE_NUMBER, FULL_NAME, CHAT_ID)
VALUES ('123-456-7890', 'John Doe', 123456789.0),
       ('234-567-8901', 'Jane Smith', 234567890.1),
       ('345-678-9012', 'Alice Johnson', 345678901.2),
       ('456-789-0123', 'Bob Brown', 456789012.3),
       ('567-890-1234', 'Charlie Davis', 567890123.4);

INSERT INTO COLLECTIONS (TITLE, questions)
VALUES ('Math Quiz', '1.Which sentences are correct? [qazwsxedc]
 A)I often play football on Sundays.
 B)I play often football on Sundays.
 C)Often I play football on Sundays.');

INSERT INTO USER_COLLECTION (USER_ID, ANSWER_FIELD, collection_id)
VALUES (1, '{"A"}', 1);

INSERT INTO GROUPS (name, teacher_name, level, start_time, started_date, days_week)
VALUES ('Group A', 'Mr. Adams', 'BEGINNER' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group B', 'Ms. Baker', 'ELEMENTARY' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group C', 'Dr. Clark', 'INTERMEDIATE' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group D', 'Mrs. Davis', 'ADVANCED' , '12:00','2024-12-12','ODD_DAYS'),
       ('Group E', 'Prof. Edwards', 'PROFICIENT' , '12:00','2024-12-12','ODD_DAYS');

INSERT INTO ANSWERS (COLLECTION_ID, ANSWER_FIELD)
VALUES (1, '{"A"}');
