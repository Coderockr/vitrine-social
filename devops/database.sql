CREATE TABLE categories (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL UNIQUE,
    icon    VARCHAR(255) NOT NULL
);

CREATE TABLE organizations (
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    logo        VARCHAR(255) NOT NULL,
    address     VARCHAR(255) NOT NULL,
    phone       VARCHAR(255) NOT NULL,
    resume      TEXT NOT NULL,
    video       VARCHAR(255) NULL,
    email       VARCHAR(255) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    slug        VARCHAR(255) NOT NULL
);

CREATE TABLE needs (
    id                  SERIAL PRIMARY KEY,
    categoty_id         INTEGER NOT NULL REFERENCES categories (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    organization_id     INTEGER NOT NULL REFERENCES organizations (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    title               VARCHAR(255) NOT NULL,
    description         TEXT NOT NULL,
    required_qtd        INTEGER NOT NULL,
    reached_qtd         INTEGER NOT NULL,
    due_date             DATE NULL,
    status              char(10) NOT NULL
);

CREATE TABLE organizations_images (
    id              SERIAL PRIMARY KEY,
    organization_id INTEGER REFERENCES organizations (id) ON UPDATE CASCADE ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    url             VARCHAR(255) NOT NULL
);

CREATE TABLE needs_images (
    id              SERIAL PRIMARY KEY,
    need_id         INTEGER REFERENCES needs (id) ON UPDATE CASCADE ON DELETE CASCADE,
    name            VARCHAR(255) NOT NULL,
    url             VARCHAR(255) NOT NULL
);
