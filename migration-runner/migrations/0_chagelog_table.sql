CREATE TABLE
    IF NOT EXISTS database_changelog (
        id int NOT NULL AUTO_INCREMENT,
        filename varchar(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        executed_by varchar(255) NOT NULL,
        UNIQUE (id, filename)
    )