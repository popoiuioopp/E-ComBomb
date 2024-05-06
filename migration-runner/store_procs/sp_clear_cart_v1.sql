CREATE PROCEDURE IF NOT EXISTS ClearCart_v1 (IN p_cartId INT) BEGIN
DELETE FROM cart_items
WHERE
    cart_id = p_cartId;

END;