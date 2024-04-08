-- +goose Up
-- +goose StatementBegin
INSERT INTO stocks."stock" (sku, total_count, reserved)
VALUES (773297411, 150, 10);

INSERT INTO stocks."stock" (sku, total_count, reserved)
VALUES (1002, 200, 20);

INSERT INTO stocks."stock" (sku, total_count, reserved)
VALUES (1003, 250, 30);

INSERT INTO stocks."stock" (sku, total_count, reserved)
VALUES (1004, 300, 40);

INSERT INTO stocks."stock" (sku, total_count, reserved)
VALUES (1005, 350, 50);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
