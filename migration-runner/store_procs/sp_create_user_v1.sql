CREATE PROCEDURE IF NOT EXISTS CreateUser_v1 (
    IN p_username VARCHAR(255),
    IN p_password VARCHAR(255)
) BEGIN
INSERT INTO
    users (username, password)
VALUES
    (p_username, p_password);

END;