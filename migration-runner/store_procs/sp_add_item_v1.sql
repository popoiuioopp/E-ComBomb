CREATE PROCEDURE AddItem_v1 (
    IN p_cartId INT,
    IN p_productId INT,
    IN p_quantity INT
) BEGIN DECLARE v_itemId INT;

SELECT
    id INTO v_itemId
FROM
    cart_items
WHERE
    cart_id = p_cartId
    AND product_id = p_productId;

IF v_itemId IS NULL THEN
INSERT INTO
    cart_items (cart_id, product_id, quantity)
VALUES
    (p_cartId, p_productId, p_quantity);

ELSE
UPDATE cart_items
SET
    quantity = quantity + p_quantity
WHERE
    id = v_itemId;

END IF;

END;