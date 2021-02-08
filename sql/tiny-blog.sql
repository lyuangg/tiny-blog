SET NAMES utf8mb4;


CREATE TABLE `tiny_posts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `post_title` varchar(100) NOT NULL DEFAULT '',
  `post_author` int(11) NOT NULL,
  `post_content` text NOT NULL,
  `post_created_at` datetime NOT NULL,
  `post_status` tinyint(4) NOT NULL DEFAULT '0',
  `post_updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `tiny_posts` (`id`, `post_title`, `post_author`, `post_content`, `post_created_at`, `post_status`, `post_updated_at`)
VALUES
	(1,'about',1,'about content', NOW(),1, NOW());



CREATE TABLE `tiny_users` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) NOT NULL DEFAULT '',
  `user_email` varchar(100) NOT NULL DEFAULT '',
  `user_password` varchar(300) NOT NULL DEFAULT '',
  `user_created_at` datetime NOT NULL,
  `user_updated_at` datetime DEFAULT NULL,
  `user_token` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`user_email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



INSERT INTO `tiny_users` (`id`, `user_name`, `user_email`, `user_password`, `user_created_at`, `user_updated_at`, `user_token`)
VALUES
	(1,'tinyblog','tiny@blog.com','aae95a43ac126de6e42f1e2b3b04dcfd',NOW(),NOW(),'');