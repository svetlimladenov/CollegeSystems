CREATE TABLE users (
    id INT NOT NULL,
    username NVARCHAR(255) NOT NULL,
    password NVARCHAR(255) NOT NULL,
    first_name NVARCHAR(255) NOT NULL,
    last_name NVARCHAR(255) NOT NULL,
    email NVARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO CollegeSystems.users (id, username, password, first_name, last_name, email) VALUES (1, 'svetlin', '123456', 'svetlin', 'mladenov', 'svetlin@st6.io')