-- test_database

CREATE TABLE `users`(
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `first_name` varchar(255) NOT NULL,
    `last_name` varchar(255) NOT NULL,
    `avatar_url` TEXT,
    PRIMARY KEY(`id`)
);

INSERT INTO `users` (`first_name`, `last_name`, `avatar_url`) VALUES ('kyoya', 'murakami', 'example.com');
INSERT INTO `users` (`first_name`, `last_name`, `avatar_url`) VALUES ('hayato', 'kagiyama', 'example.com');
INSERT INTO `users` (`first_name`, `last_name`, `avatar_url`) VALUES ('shoki', 'shirasugi', 'example.com');
