CREATE PROCEDURE IF NOT EXISTS GetAllProducts_v1 () BEGIN
SELECT
    id,
    created_at,
    updated_at,
    name,
    price,
    description,
    user_id,
    product_image
FROM
    products;

END;