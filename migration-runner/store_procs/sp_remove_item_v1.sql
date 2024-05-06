CREATE PROCEDURE IF NOT EXISTS RemoveItem_v1 (IN p_userId INT, IN p_productId INT) BEGIN
DELETE FROM cart_items
WHERE
    cart_id = (
        SELECT
            id
        FROM
            carts
        WHERE
            user_id = p_userId
    )
    AND product_id = p_productId;

END;