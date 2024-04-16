CREATE PROCEDURE UpdateItemQuantity_v1 (
    IN p_cartId INT,
    IN p_productId INT,
    IN p_quantity INT
) BEGIN
UPDATE cart_items
SET
    quantity = p_quantity
WHERE
    cart_id = p_cartId
    AND product_id = p_productId;

END