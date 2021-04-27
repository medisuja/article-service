-- +goose Up
-- +goose StatementBegin
SELECT CREATE TABLE article (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	author VARCHAR(100) NOT NULL,
	title TEXT NOT NULL,
	slug TEXT NOT NULL,
	body LONGTEXT,
	created TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=INNODB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE article;
-- +goose StatementEnd
