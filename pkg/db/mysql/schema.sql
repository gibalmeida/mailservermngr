
CREATE TABLE `account` (
  `name` char(32) NOT NULL,
  `domain` char(96) NOT NULL,
  `password` char(40) NOT NULL,
  `home_dir` char(160) DEFAULT NULL,
  `quota` char(20) DEFAULT NULL,
  `clear_password` char(16) DEFAULT NULL,
  PRIMARY KEY (`name`,`domain`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `address_alias` (
  `alias` varchar(200) NOT NULL,
  `addresses` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`alias`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `domain` (
  `domain` char(96) NOT NULL,
  PRIMARY KEY (`domain`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `domain_alias` (
  `alias` char(32) NOT NULL,
  `domain` char(96) NOT NULL,
  KEY `alias` (`alias`,`domain`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=utf8mb4_0900_ai_ci;
