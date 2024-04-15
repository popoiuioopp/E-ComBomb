CREATE PROCEDURE IF NOT EXISTS GetUser_v1 (p_username VARCHAR(255)) BEGIN
SELECT
    id,
    username,
    password
FROM
    users
WHERE
    username = p_username;

END;