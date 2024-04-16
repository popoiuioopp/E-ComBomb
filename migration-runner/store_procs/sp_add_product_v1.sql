CREATE PROCEDURE IF NOT EXISTS AddProduct_v1 (
    IN p_productName VARCHAR(255),
    IN p_price FLOAT,
    IN p_description LONGTEXT,
    IN p_userId int,
    IN p_productImage LONGTEXT
) BEGIN
INSERT INTO
    products (name, price, description, user_id, product_image)
VALUES
    (
        p_productName,
        p_price,
        p_description,
        p_userId,
        p_productImage
    );

END;