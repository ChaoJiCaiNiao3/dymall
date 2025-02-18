DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*创建关联表，用于存储地址信息，通过 user_id 外键关联到 user 表的 id*/
CREATE TABLE `address` (
    `address_id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT,
    `address` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`address_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*创建商品表*/
CREATE TABLE `product` (
    `product_id` INT AUTO_INCREMENT PRIMARY KEY,
    `price` DECIMAL(10, 2) NOT NULL,
    `product_name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
