CREATE PROCEDURE GetOrCreateCart_v1 (IN p_userId INT) BEGIN DECLARE v_cartId INT;

SELECT
    id INTO v_cartId
FROM
    carts
WHERE
    user_id = p_userId;

IF v_cartId IS NULL THEN
INSERT INTO
    carts (user_id)
VALUES
    (p_userId);

SET
    v_cartId = LAST_INSERT_ID ();

END IF;

SELECT
    id,
    user_id
FROM
    carts
WHERE
    id = v_cartId;

END;