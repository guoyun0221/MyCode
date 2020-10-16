+ A database named "thirty_nine" is needed, and there are three tables in it:

+ + CREATE TABLE `blocked_ip` (
  `IP` varchar(20) NOT NULL,
  PRIMARY KEY (`IP`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

+ + CREATE TABLE `speeches` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `words` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `send_time` datetime DEFAULT NULL,
  `img_name` text CHARACTER SET utf8 COLLATE utf8_general_ci,
  `reply_to` int DEFAULT NULL,
  `at_top` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=245 DEFAULT CHARSET=utf8

+ + CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `level` int unsigned DEFAULT NULL,
  `experience` int unsigned DEFAULT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8
