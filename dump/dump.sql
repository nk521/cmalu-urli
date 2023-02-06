CREATE TABLE `links` (
  `original_url` varchar(2083) NOT NULL,
  `short_url` varchar(100) NOT NULL,
  `expiry` int NOT NULL,
  PRIMARY KEY (`short_url`)
)
