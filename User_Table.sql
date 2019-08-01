CREATE TABLE `User` (
  `UserId` bigint(20) NOT NULL AUTO_INCREMENT,
  `Username` varchar(200) DEFAULT NULL,
  `Email` varchar(1024) DEFAULT NULL,
  `Password` varchar(1024) NULL DEFAULT NULL,
  `Phone` varchar(1024) NULL DEFAULT NULL,
  PRIMARY KEY (`UserId`),
  UNIQUE KEY `ID_UNIQUE` (`UserId`)
);
