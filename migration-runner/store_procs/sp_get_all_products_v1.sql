CREATE PROCEDURE IF NOT EXISTS GetAllProducts_v1 () BEGIN
SELECT
    id,
    name,
    price,
    description,
    user_id,
    product_image
FROM
    products;

END;