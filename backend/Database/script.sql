CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(72) NOT NULL ,
    is_admin BOOLEAN NOT NULL DEFAULT false
);

CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL ,
    description VARCHAR(250) NOT NULL ,
    price FLOAT NOT NULL,
    img_url VARCHAR(100) NOT NULL
);

CREATE TABLE basket_items (
    article_id INTEGER NOT NULL REFERENCES articles(id),
    user_id INTEGER NOT NULL REFERENCES users(id),
    PRIMARY KEY (article_id, user_id),
    quantity INTEGER NOT NULL
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id)
);

CREATE TABLE order_lines (
    order_id INTEGER NOT NULL REFERENCES orders(id),
    article_id INTEGER NOT NULL REFERENCES articles(id),
    quantity INTEGER,
    PRIMARY KEY (order_id, article_id)
);

-- DROIT A WEB3CLIENT
-- GRANT INSERT, UPDATE, DELETE, SELECT ON users, orders, articles, basket_items, order_lines TO web3_client;

--INSERT ADMIN
INSERT INTO users(username, password, is_admin)
VALUES ('ADMIN', '$2a$10$ItBW25aEGAxIEA4e/0WYrObKFjx5aHcpQqVe/LQsbzjt8LzrlfM9i', TRUE);

-- INSERTING DUMMY ARTICLES
INSERT INTO articles (name, description, price, img_url)
VALUES ('article 1', 'A very nice article to have in your home', 12.50, 'https://dummyimage.com/600x400/D7D7/000&text=article1'),
       ('Something Nice', 'Great accessory', 56.99, 'https://dummyimage.com/600x400/62A3/000&text=article2'),
       ('Article Y', 'Futuristic article', 106.45, 'https://dummyimage.com/600x400/000baa/000&text=article3')

