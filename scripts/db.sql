CREATE DATABASE banco;
USE banco;

DROP TABLE IF EXISTS `clientes`;
CREATE TABLE `clientes` (
  `cliente_id` int(11) NOT NULL AUTO_INCREMENT,
  `nombre` varchar(100) NOT NULL,
  `fecha_nacimiento` date NOT NULL,
  `ciudad` varchar(100) NOT NULL,
  `codigo_postal` varchar(10) NOT NULL,
  `estatus` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`cliente_id`)
) ENGINE=InnoDB AUTO_INCREMENT=2006 DEFAULT CHARSET=latin1;
INSERT INTO `clientes` VALUES
	(2000,'Steve','1978-12-15','Delhi','110075',1),
	(2001,'Arian','1988-05-21','Newburgh, NY','12550',1),
	(2002,'Hadley','1988-04-30','Englewood, NJ','07631',1),
	(2003,'Ben','1988-01-04','Manchester, NH','03102',0),
	(2004,'Nina','1988-05-14','Clarkston, MI','48348',1),
	(2005,'Osman','1988-11-08','Hyattsville, MD','20782',0);


DROP TABLE IF EXISTS `cuentas`;
CREATE TABLE `cuentas` (
  `cuenta_id` int(11) NOT NULL AUTO_INCREMENT,
  `cliente_id` int(11) NOT NULL,
  `fecha_apertura` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `tipo_cuenta` varchar(10) NOT NULL,
  `cantidad` decimal(10,2) NOT NULL,
  `estatus` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`cuenta_id`),
  KEY `cuentas_FK` (`cliente_id`),
  CONSTRAINT `cuentas_FK` FOREIGN KEY (`cliente_id`) REFERENCES `clientes` (`cliente_id`)
) ENGINE=InnoDB AUTO_INCREMENT=95471 DEFAULT CHARSET=latin1;
INSERT INTO `cuentas` VALUES
	(95470,2000,'2020-08-22 10:20:06', 'ahorro', 6823.23, 1),
	(95471,2002,'2020-08-09 10:27:22', 'chequera', 3342.96, 1),
  (95472,2001,'2020-08-09 10:35:22', 'ahorro', 7000, 1),
  (95473,2001,'2020-08-09 10:38:22', 'ahorro', 5861.86, 1);


DROP TABLE IF EXISTS `transacciones`;
CREATE TABLE `transacciones` (
  `transaccion_id` int(11) NOT NULL AUTO_INCREMENT,
  `cuenta_id` int(11) NOT NULL,
  `cantidad` decimal(10,2) NOT NULL,
  `tipo_transaccion` varchar(10) NOT NULL,
  `fecha_transaccion` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`transaccion_id`),
  KEY `transacciones_FK` (`cuenta_id`),
  CONSTRAINT `transacciones_FK` FOREIGN KEY (`cuenta_id`) REFERENCES `cuentas` (`cuenta_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `usuarios`;
CREATE TABLE `usuarios` (
  `username` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `rol` varchar(20) NOT NULL,
  `cliente_id` int(11) DEFAULT NULL,
  `creado_en` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
INSERT INTO `usuarios` VALUES
  ('admin','abc123','admin', NULL, '2020-08-09 10:27:22'),
  ('2001','abc123','user', 2001, '2020-08-09 10:27:22'),
  ('2000','abc123','user', 2000, '2020-08-09 10:27:22');
