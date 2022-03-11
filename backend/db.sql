DROP TABLE IF EXISTS tbl_todo;

CREATE TABLE tbl_todo (
    name VARCHAR(100) NOT NULL,
    status VARCHAR(100) NOT NULL,
  
    CONSTRAINT name_unique UNIQUE (name)
);