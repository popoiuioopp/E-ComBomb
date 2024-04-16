CREATE PROCEDURE GetCartByUserId_v1 (IN p_userId INT) BEGIN DECLARE v_cartId INT DEFAULT NULL;

-- First, get the cart ID to use in subsequent queries
SELECT
    id INTO v_cartId
FROM
    carts
WHERE
    user_id = p_userId
LIMIT
    1;

-- If a cart exists, return the cart info and the item details
IF v_cartId IS NOT NULL THEN
SELECT
    c.id,
    c.created_at,
    c.updated_at,
    c.deleted_at,
    c.user_id
FROM
    carts c
WHERE
    c.id = v_cartId;

SELECT
    ci.product_id,
    ci.quantity,
    p.name,
    p.price,
    p.description
FROM
    cart_items ci
    JOIN products p ON ci.product_id = p.id
WHERE
    ci.cart_id = v_cartId;

END IF;

END;