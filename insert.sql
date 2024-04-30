CREATE DATABASE `imersao17`;

use `imersao17`;

CREATE TABLE `categories` (
`id` varchar (36) NOT NULL,
`name` varchar (255) NOT NULL,
PRIMARY KEY (`id`)
);

INSERT INTO `imersao17`.`categories` (`id`, `name`) VALUES ('57dd2eb2-7406-4f3c-a919-e4ddc72edd17', 'Category 1');
INSERT INTO `imersao17`.`categories` (`id`, `name`) VALUES ('bce9ddd4-7611-409a-96c4-3551c5f40410', 'Category 2');
 
CREATE TABLE `products`
(
`id` varchar (36) NOT NULL,
`name` varchar (255) NOT NULL,
`description` varchar (255) NOT NULL,
`price` decimal (10,2) NOT NULL,
`category_id` varchar (36) NOT NULL,
PRIMARY KEY (`id`)
);

INSERT INTO `imersao17`.`products` (`id`, `name`, `description`, `price`, `category_id`) VALUES ('5b635289-de29-4415-b5cf-287a6f0d0f2b', 'Product 1', 'Product 1 description', 10.99, '57dd2eb2-7406-4f3c-a919-e4ddc72edd17');
INSERT INTO `imersao17`.`products` (`id`, `name`, `description`, `price`, `category_id`) VALUES ('e563d76b-171a-4574-b3e9-3ba26c0683ec', 'Product 2', 'Product 2 description', 20.99, '57dd2eb2-7406-4f3c-a919-e4ddc72edd17');
INSERT INTO `imersao17`.`products` (`id`, `name`, `description`, `price`, `category_id`) VALUES ('fcc7faea-1e27-43a1-bd9b-0e104b21820f', 'Product 3', 'Product 3 description', 30.99, '57dd2eb2-7406-4f3c-a919-e4ddc72edd17');

