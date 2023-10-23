USE cartCraftDB;

ALTER TABLE Users
ADD token VARCHAR(255),
ADD refreshed_token VARCHAR(255),
ADD type VARCHAR(100);