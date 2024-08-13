-- +goose Up
-- +goose StatementBegin
CREATE TABLE hn_item (
    id int NOT NULL,
    deleted boolean,
    item_type text,
    item_by text,
    created_time int,
    item_content text,
    parent int,
    kids json,
    item_url text,
    item_score int,
    item_title text,
    descendants int,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE hn_item;
-- +goose StatementEnd
