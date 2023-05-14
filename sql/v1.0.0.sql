-- +migrate Up
-- +migrate StatementBegin
CREATE SEQUENCE IF NOT EXISTS customer_pkey_seq;
CREATE TABLE IF NOT EXISTS customer
(
    id      BIGINT NOT NULL             DEFAULT nextval('customer_pkey_seq'::regclass),
    name    VARCHAR(50),
    email   VARCHAR(50),
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted BOOL                        DEFAULT FALSE,
    CONSTRAINT customer_pkey_seq_id PRIMARY KEY (id),
    CONSTRAINT uq_customer_pkey_seq UNIQUE (name, email)
);

CREATE SEQUENCE IF NOT EXISTS user_pkey_seq;
CREATE TABLE IF NOT EXISTS "user"
(
    id        BIGINT NOT NULL             DEFAULT nextval('user_pkey_seq'::regclass),
    user_name VARCHAR(50),
    password  VARCHAR(100),
    created   TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated   TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted   BOOL                        DEFAULT FALSE,
    CONSTRAINT user_pkey_seq_id PRIMARY KEY (id),
    CONSTRAINT uq_user_pkey_seq UNIQUE (user_name)
);

INSERT INTO "user"(user_name, password)
VALUES ('tetsing_user', '0c57d52d6d1bea70dd1b55cfb4a93295f117fe42041323e663cd059a4ba9133d')
ON CONFLICT(user_name) DO NOTHING ;


CREATE SEQUENCE IF NOT EXISTS product_pkey_seq;
CREATE TABLE IF NOT EXISTS product
(
    id      BIGINT NOT NULL             DEFAULT nextval('product_pkey_seq'::regclass),
    name    VARCHAR(50),
    price   FLOAT,
    created TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted BOOL                        DEFAULT FALSE,
    CONSTRAINT product_pkey_seq_id PRIMARY KEY (id)
);

CREATE SEQUENCE IF NOT EXISTS order_pkey_seq;
CREATE TABLE IF NOT EXISTS "order"
(
    id          BIGINT NOT NULL             DEFAULT nextval('order_pkey_seq'::regclass),
    customer_id BIGINT,
    product_id  BIGINT,
    total_order INT,
    amount      FLOAT,
    created     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated     TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted     BOOL                        DEFAULT FALSE,
    CONSTRAINT order_pkey_seq_id PRIMARY KEY (id),
    CONSTRAINT fk_order_to_customer_id FOREIGN KEY (customer_id) REFERENCES customer (id),
    CONSTRAINT fk_order_to_product_id FOREIGN KEY (product_id) REFERENCES product (id)
);

-- +migrate StatementEnd