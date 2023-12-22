/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `CancelNote`;
CREATE TABLE `CancelNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `totalPrice` float DEFAULT '0',
  `createBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `CancelNoteDetail`;
CREATE TABLE `CancelNoteDetail` (
  `cancelNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `expiryDate` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reason` enum('Damaged','OutOfDate') NOT NULL,
  `amountCancel` float DEFAULT '0',
  PRIMARY KEY (`cancelNoteId`,`ingredientId`,`expiryDate`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Category`;
CREATE TABLE `Category` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '',
  `amountProduct` int DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `CategoryFood`;
CREATE TABLE `CategoryFood` (
  `foodId` varchar(12) NOT NULL,
  `categoryId` varchar(12) NOT NULL,
  PRIMARY KEY (`foodId`,`categoryId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Customer`;
CREATE TABLE `Customer` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `debt` float DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `CustomerDebt`;
CREATE TABLE `CustomerDebt` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `customerId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amount` float NOT NULL,
  `amountLeft` float NOT NULL,
  `type` enum('Debt','Pay') NOT NULL,
  `createAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ExportNote`;
CREATE TABLE `ExportNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `totalPrice` float DEFAULT '0',
  `createBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ExportNoteDetail`;
CREATE TABLE `ExportNoteDetail` (
  `exportNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `expiryDate` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amountExport` float DEFAULT '0',
  PRIMARY KEY (`exportNoteId`,`ingredientId`,`expiryDate`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Feature`;
CREATE TABLE `Feature` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Food`;
CREATE TABLE `Food` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` text NOT NULL,
  `cookingGuide` text NOT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ImportNote`;
CREATE TABLE `ImportNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `supplierId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `totalPrice` float DEFAULT '0',
  `status` enum('InProgress','Done','Cancel') DEFAULT 'InProgress',
  `createBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `closeBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `createAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `closeAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ImportNoteDetail`;
CREATE TABLE `ImportNoteDetail` (
  `importNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `expiryDate` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` float NOT NULL,
  `amountImport` float DEFAULT '0',
  PRIMARY KEY (`importNoteId`,`ingredientId`,`expiryDate`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Ingredient`;
CREATE TABLE `Ingredient` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `totalAmount` float DEFAULT '0',
  `measureType` enum('Weight','Volume','Unit') NOT NULL,
  `price` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `IngredientDetail`;
CREATE TABLE `IngredientDetail` (
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `expiryDate` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amount` float DEFAULT '0',
  PRIMARY KEY (`ingredientId`,`expiryDate`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Invoice`;
CREATE TABLE `Invoice` (
  `id` varchar(13) NOT NULL,
  `customerId` varchar(13) NOT NULL,
  `totalPrice` float NOT NULL,
  `amountReceived` float NOT NULL,
  `amountDebt` float NOT NULL,
  `createAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createBy` varchar(13) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InvoiceDetail`;
CREATE TABLE `InvoiceDetail` (
  `invoiceId` varchar(13) NOT NULL,
  `foodId` varchar(13) NOT NULL,
  `sizeName` varchar(13) NOT NULL,
  `amount` float NOT NULL,
  `unitPrice` float NOT NULL,
  `description` text NOT NULL,
  `toppings` json NOT NULL,
  KEY `invoiceId` (`invoiceId`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `MUser`;
CREATE TABLE `MUser` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `roleId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Recipe`;
CREATE TABLE `Recipe` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `RecipeDetail`;
CREATE TABLE `RecipeDetail` (
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `amountNeed` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Role`;
CREATE TABLE `Role` (
  `id` varchar(13) NOT NULL,
  `name` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `RoleFeature`;
CREATE TABLE `RoleFeature` (
  `roleId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `featureId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`roleId`,`featureId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SizeFood`;
CREATE TABLE `SizeFood` (
  `foodId` varchar(12) NOT NULL,
  `sizeId` varchar(12) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `cost` float NOT NULL,
  `price` float NOT NULL,
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`foodId`,`sizeId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Supplier`;
CREATE TABLE `Supplier` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `debt` float DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SupplierDebt`;
CREATE TABLE `SupplierDebt` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `supplierId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amount` float NOT NULL,
  `amountLeft` float NOT NULL,
  `type` enum('Debt','Pay') NOT NULL,
  `createAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createBy` varchar(9) NOT NULL,
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Topping`;
CREATE TABLE `Topping` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` text NOT NULL,
  `cookingGuide` text NOT NULL,
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  `cost` float NOT NULL,
  `price` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `CancelNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('3SultJVIR', 11000, 'za1u8m4Sg', '2023-11-05 02:39:16');
INSERT INTO `CancelNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('IbDrpJVSR', 11000, 'za1u8m4Sg', '2023-11-05 02:40:03');
INSERT INTO `CancelNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('sayhi', 1000, 'za1u8m4Sg', '2023-11-05 02:21:54');
INSERT INTO `CancelNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('ZoCe5JVIg', 2000, 'za1u8m4Sg', '2023-11-05 02:24:33');

INSERT INTO `CancelNoteDetail` (`cancelNoteId`, `ingredientId`, `expiryDate`, `reason`, `amountCancel`) VALUES
('3SultJVIR', 'nvl1', '10/05/2003', 'Damaged', 10);
INSERT INTO `CancelNoteDetail` (`cancelNoteId`, `ingredientId`, `expiryDate`, `reason`, `amountCancel`) VALUES
('3SultJVIR', 'nvl2', '19/05/2003', 'Damaged', 10);
INSERT INTO `CancelNoteDetail` (`cancelNoteId`, `ingredientId`, `expiryDate`, `reason`, `amountCancel`) VALUES
('IbDrpJVSR', 'nvl1', '10/05/2003', 'Damaged', 10);
INSERT INTO `CancelNoteDetail` (`cancelNoteId`, `ingredientId`, `expiryDate`, `reason`, `amountCancel`) VALUES
('IbDrpJVSR', 'nvl2', '19/05/2003', 'Damaged', 10),
('sayhi', 'nvl1', '10/05/2003', 'Damaged', 10),
('ZoCe5JVIg', 'nvl1', '01/05/2003', 'Damaged', 10),
('ZoCe5JVIg', 'nvl1', '10/05/2003', 'Damaged', 10);

INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('1', '1', '', 3);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('2', '2', NULL, 0);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('3', 'danh mục nè', NULL, 1);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('4', 'danh mục nè ú hú hú', 'Má bug gì lắm thế', 4),
('category1', '3', '', 0),
('oxnSztVIR', 'danh mục nè ú hú hú ú nu ú 11', 'bug bay màu bay màu ahaha', 0),
('P2tGkpVSg', 'danh mục nè ú hú hú ú nu ú', NULL, 0),
('uoJZkp4IR', 'danh mục nè ú hú hú ú nu ú nu', NULL, 0),
('YGCnktVIg', 'danh mục nè ú hú hú ú nu ú 1', '', 0),
('zyKFD6HIR', 'danh mục ahihi', 'bug bay màu bay màu ahaha', 0);

INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('foodtest', '1');
INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('foodtest', '4');
INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('J1F4ZEVIg', '3');
INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('J1F4ZEVIg', '4');

INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('cs1', '123', 'a@gmail.com', '01234567892', 1444430);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('cs2', '123', 'a@gmail.com', '01234567891', 722214);


INSERT INTO `CustomerDebt` (`id`, `customerId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('_idZ9P4IR', 'cs1', 361107, 361107, 'Debt', '2023-11-07 05:24:45', 'za1u8m4Sg');
INSERT INTO `CustomerDebt` (`id`, `customerId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('jDKWrEVIg', 'cs1', 361107, 722214, 'Debt', '2023-11-07 05:24:47', 'za1u8m4Sg');
INSERT INTO `CustomerDebt` (`id`, `customerId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('LB5WrE4Sg', 'cs1', 361107, 1444430, 'Debt', '2023-11-07 05:24:48', 'za1u8m4Sg');
INSERT INTO `CustomerDebt` (`id`, `customerId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('uYAMrE4SR', 'cs2', 361107, 722214, 'Debt', '2023-11-07 05:25:25', 'za1u8m4Sg'),
('X-2MrP4IR', 'cs2', 361107, 361107, 'Debt', '2023-11-07 05:25:23', 'za1u8m4Sg'),
('xqFWrEVIR', 'cs1', 361107, 1083320, 'Debt', '2023-11-07 05:24:47', 'za1u8m4Sg');

INSERT INTO `ExportNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('118xALVIR', 1100, 'za1u8m4Sg', '2023-11-06 16:44:14');
INSERT INTO `ExportNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('7SubAL4Ig', 1100, 'za1u8m4Sg', '2023-11-06 16:44:17');
INSERT INTO `ExportNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('9dPbAL4Ig', 1100, 'za1u8m4Sg', '2023-11-06 16:44:12');
INSERT INTO `ExportNote` (`id`, `totalPrice`, `createBy`, `createAt`) VALUES
('dadasdasdsad', 101000, 'za1u8m4Sg', '2023-11-05 03:03:17'),
('DRlxAYVSg', 1100, 'za1u8m4Sg', '2023-11-06 16:44:16'),
('jp9xAL4Ig', 1100, 'za1u8m4Sg', '2023-11-06 16:44:18'),
('Muux0YVSR', 1100, 'za1u8m4Sg', '2023-11-06 16:44:18'),
('oC9x0LVIg', 1100, 'za1u8m4Sg', '2023-11-06 16:44:19'),
('yTjb0L4Sg', 1100, 'za1u8m4Sg', '2023-11-06 16:44:19'),
('YYlbAY4SR', 1100, 'za1u8m4Sg', '2023-11-06 16:44:16'),
('ZcQb0L4IR', 1100, 'za1u8m4Sg', '2023-11-06 16:44:15');

INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `expiryDate`, `amountExport`) VALUES
('118xALVIR', 'nvl1', '19/05/2003', 1);
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `expiryDate`, `amountExport`) VALUES
('118xALVIR', 'nvl2', '19/05/2003', 1);
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `expiryDate`, `amountExport`) VALUES
('7SubAL4Ig', 'nvl1', '19/05/2003', 1);
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `expiryDate`, `amountExport`) VALUES
('7SubAL4Ig', 'nvl2', '19/05/2003', 1),
('9dPbAL4Ig', 'nvl1', '19/05/2003', 1),
('9dPbAL4Ig', 'nvl2', '19/05/2003', 1),
('dadasdasdsad', 'nvl1', '19/05/2003', 10),
('dadasdasdsad', 'nvl2', '19/05/2003', 100),
('DRlxAYVSg', 'nvl1', '19/05/2003', 1),
('DRlxAYVSg', 'nvl2', '19/05/2003', 1),
('jp9xAL4Ig', 'nvl1', '19/05/2003', 1),
('jp9xAL4Ig', 'nvl2', '19/05/2003', 1),
('Muux0YVSR', 'nvl1', '19/05/2003', 1),
('Muux0YVSR', 'nvl2', '19/05/2003', 1),
('oC9x0LVIg', 'nvl1', '19/05/2003', 1),
('oC9x0LVIg', 'nvl2', '19/05/2003', 1),
('yTjb0L4Sg', 'nvl1', '19/05/2003', 1),
('yTjb0L4Sg', 'nvl2', '19/05/2003', 1),
('YYlbAY4SR', 'nvl1', '19/05/2003', 1),
('YYlbAY4SR', 'nvl2', '19/05/2003', 1),
('ZcQb0L4IR', 'nvl1', '19/05/2003', 1),
('ZcQb0L4IR', 'nvl2', '19/05/2003', 1);

INSERT INTO `Feature` (`id`, `description`) VALUES
('CAN_CREATE', 'create new cancel note');
INSERT INTO `Feature` (`id`, `description`) VALUES
('CAN_VIEW', 'view cancel note');
INSERT INTO `Feature` (`id`, `description`) VALUES
('CAT_CREATE', 'create new category');
INSERT INTO `Feature` (`id`, `description`) VALUES
('CAT_UP_INFO', 'update category\'s info'),
('CAT_VIEW', 'view category'),
('CUS_CREATE', 'create new customer'),
('CUS_PAY', 'make a pay for customer'),
('CUS_UP_INFO', 'update info for customer'),
('CUS_VIEW', 'view customer'),
('EXP_CREATE', 'create export note'),
('EXP_VIEW', 'view export note'),
('FOD_CREATE', 'create new food'),
('FOD_UP_INFO', 'update food\'s info'),
('FOD_UP_STATE', 'update food\'s status'),
('FOD_VIEW', 'view food'),
('IMP_CREATE', 'create new import note'),
('IMP_UP_STATE', 'update status of import note'),
('IMP_VIEW', 'view import note'),
('ING_CREATE', 'create new ingredient'),
('ING_VIEW', 'view ingredient'),
('INV_CREATE', 'create new invoice'),
('INV_VIEW', 'view invoice'),
('SUP_CREATE', 'create new supplier'),
('SUP_PAY', 'make a pay for supplier'),
('SUP_UP_INFO', 'update supplier\'s info'),
('SUP_VIEW', 'view supplier'),
('TOP_CREATE', 'create new topping'),
('TOP_UP_INFO', 'update topping\'s info'),
('TOP_UP_STATE', 'update topping\'s status'),
('TOP_VIEW', 'view topping'),
('USE_UP_INFO', 'update info user'),
('USE_UP_STATE', 'update status of user'),
('USE_VIEW', 'view user');

INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `isActive`) VALUES
('foodtest', 'foodtest', '123', '', 1);
INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `isActive`) VALUES
('J1F4ZEVIg', 'foodtest1', '123', '', 1);


INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createBy`, `closeBy`, `createAt`, `closeAt`) VALUES
('1UiW-04IR', 'rs2_QiVIg', 4300, 'InProgress', 'za1u8m4Sg', NULL, '2023-11-04 18:51:54', NULL);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createBy`, `closeBy`, `createAt`, `closeAt`) VALUES
('RcKRaA4SR', 'rs2_QiVIg', 5300, 'Cancel', 'za1u8m4Sg', 'za1u8m4Sg', '2023-11-04 18:50:25', '2023-11-04 18:53:27');
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createBy`, `closeBy`, `createAt`, `closeAt`) VALUES
('s7AawE4Ig', 'tjRaQEVSR', 5300, 'InProgress', 'za1u8m4Sg', NULL, '2023-11-07 04:41:10', NULL);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createBy`, `closeBy`, `createAt`, `closeAt`) VALUES
('sayhi', 'rs2_QiVIg', 5300, 'Done', 'za1u8m4Sg', 'za1u8m4Sg', '2023-11-04 19:15:36', '2023-11-04 19:15:41'),
('test', 'rs2_QiVIg', 5300, 'Done', 'za1u8m4Sg', 'za1u8m4Sg', '2023-11-04 19:08:10', '2023-11-04 19:08:34'),
('ybRkaAVIR', 'rs2_QiVIg', 4300, 'Done', 'za1u8m4Sg', 'za1u8m4Sg', '2023-11-04 18:50:46', '2023-11-04 19:05:41');

INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `expiryDate`, `price`, `amountImport`) VALUES
('1UiW-04IR', 'nvl1', '10/05/2003', 0, 10);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `expiryDate`, `price`, `amountImport`) VALUES
('1UiW-04IR', 'nvl1', '19/05/2003', 200, 20);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `expiryDate`, `price`, `amountImport`) VALUES
('1UiW-04IR', 'nvl3', '19/05/2003', 300, 1);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `expiryDate`, `price`, `amountImport`) VALUES
('RcKRaA4SR', 'nvl1', '10/05/2003', 100, 10),
('RcKRaA4SR', 'nvl1', '19/05/2003', 200, 20),
('RcKRaA4SR', 'nvl3', '19/05/2003', 300, 1),
('s7AawE4Ig', 'nvl1', '10/05/2003', 100, 10),
('s7AawE4Ig', 'nvl1', '19/05/2003', 200, 20),
('s7AawE4Ig', 'nvl11', '19/05/2003', 300, 1),
('sayhi', 'nvl1', '10/05/2003', 100, 10),
('sayhi', 'nvl1', '19/05/2003', 200, 20),
('sayhi', 'nvl11', '19/05/2003', 300, 1),
('test', 'nvl1', '10/05/2003', 100, 10),
('test', 'nvl1', '19/05/2003', 200, 20),
('test', 'nvl4', '19/05/2003', 300, 1),
('ybRkaAVIR', 'nvl1', '10/05/2003', 0, 10),
('ybRkaAVIR', 'nvl1', '19/05/2003', 200, 20),
('ybRkaAVIR', 'nvl3', '19/05/2003', 300, 1);

INSERT INTO `Ingredient` (`id`, `name`, `totalAmount`, `measureType`, `price`) VALUES
('56TJ5tVIR', 'nvl haha11121', 0, 'Weight', 0);
INSERT INTO `Ingredient` (`id`, `name`, `totalAmount`, `measureType`, `price`) VALUES
('BZO1ct4Ig', 'nvl haha121', 0, 'Weight', 0);
INSERT INTO `Ingredient` (`id`, `name`, `totalAmount`, `measureType`, `price`) VALUES
('cceh5p4Sg', 'nvl haha1', 0, 'Unit', 0);
INSERT INTO `Ingredient` (`id`, `name`, `totalAmount`, `measureType`, `price`) VALUES
('nvl1', 'nvl 1', 50, 'Weight', 100),
('nvl11', 'nvl haha12', 1, 'Weight', 300),
('nvl2', 'nvl 2', 70, 'Volume', 1000),
('nvl3', 'nvl 3', 2, 'Volume', 1000000),
('ya-p5p4Sg', 'nvl haha', 0, 'Unit', 0);

INSERT INTO `IngredientDetail` (`ingredientId`, `expiryDate`, `amount`) VALUES
('nvl1', '10/05/2003', 0);
INSERT INTO `IngredientDetail` (`ingredientId`, `expiryDate`, `amount`) VALUES
('nvl1', '19/05/2003', 40);
INSERT INTO `IngredientDetail` (`ingredientId`, `expiryDate`, `amount`) VALUES
('nvl1', '20/05/2003', 10);
INSERT INTO `IngredientDetail` (`ingredientId`, `expiryDate`, `amount`) VALUES
('nvl11', '19/05/2003', 1),
('nvl2', '19/05/2003', 70),
('nvl3', '19/05/2003', 2),
('nvl4', '19/05/2003', 1);

INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountDebt`, `createAt`, `createBy`) VALUES
('4ahGrPVSg', 'cs2', 401230, 40123, 361107, '2023-11-07 05:25:23', 'za1u8m4Sg');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountDebt`, `createAt`, `createBy`) VALUES
('7qFW9PVSR', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:47', 'za1u8m4Sg');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountDebt`, `createAt`, `createBy`) VALUES
('CkdWrE4Sg', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:45', 'za1u8m4Sg');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountDebt`, `createAt`, `createBy`) VALUES
('GL0GrE4Ig', 'cs2', 401230, 40123, 361107, '2023-11-07 05:25:25', 'za1u8m4Sg'),
('HvFWrP4Sg', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:47', 'za1u8m4Sg'),
('nf5WrEVSR', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:48', 'za1u8m4Sg');

INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('CkdWrE4Sg', 'foodtest', 'size M', 10, 30000, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]');
INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('CkdWrE4Sg', 'foodtest', 'size SSS', 10, 10123, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]');
INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('HvFWrP4Sg', 'foodtest', 'size M', 10, 30000, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]');
INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('HvFWrP4Sg', 'foodtest', 'size SSS', 10, 10123, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('7qFW9PVSR', 'foodtest', 'size M', 10, 30000, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('7qFW9PVSR', 'foodtest', 'size SSS', 10, 10123, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('nf5WrEVSR', 'foodtest', 'size M', 10, 30000, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('nf5WrEVSR', 'foodtest', 'size SSS', 10, 10123, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('4ahGrPVSg', 'foodtest', 'size M', 10, 30000, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('4ahGrPVSg', 'foodtest', 'size SSS', 10, 10123, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('GL0GrE4Ig', 'foodtest', 'size M', 10, 30000, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]'),
('GL0GrE4Ig', 'foodtest', 'size SSS', 10, 10123, 'haha', '[{\"id\": \"ZNWdVb4Ig\", \"name\": \"topping7\", \"price\": 10000}]');

INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `isActive`) VALUES
('g3W21A7SR', '123', '31231232131', '', 'b@gmail.com', '5e107317df151f6e8e0015c4f2ee7936', 'mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms', 'user', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `address`, `email`, `password`, `salt`, `roleId`, `isActive`) VALUES
('za1u8m4Sg', 'say hi', '', '', 'c@gmail.com', 'cb58ac5a2272517d1960565444bde187', 'QYlnGKRgYBxIXzMnnQSVcglbtjPsAhVlxMRMDaqnaquxwADSur', 'admin', 1);


INSERT INTO `Recipe` (`id`) VALUES
('1JKVZPVSRm');
INSERT INTO `Recipe` (`id`) VALUES
('6vkZJTVIRz');
INSERT INTO `Recipe` (`id`) VALUES
('GMbtAo4IRz');
INSERT INTO `Recipe` (`id`) VALUES
('J1F4WPVIRM'),
('ZHZd4bVSgz');

INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('GMbtAo4IRz', 'nvl3', 1000);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('6vkZJTVIRz', 'nvl3', 1000);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('ZHZd4bVSgz', 'nvl1', 10);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('ZHZd4bVSgz', 'nvl11', 10),
('1JKVZPVSRm', 'nvl1', 20),
('J1F4WPVIRM', 'nvl1', 40);

INSERT INTO `Role` (`id`, `name`) VALUES
('admin', 'user');
INSERT INTO `Role` (`id`, `name`) VALUES
('fl_9lf4Ig', 'haha');
INSERT INTO `Role` (`id`, `name`) VALUES
('user', 'user');

INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAN_CREATE');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAN_VIEW');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_CREATE');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_UP_INFO'),
('admin', 'CAT_VIEW'),
('admin', 'CUS_CREATE'),
('admin', 'CUS_PAY'),
('admin', 'CUS_UP_INFO'),
('admin', 'CUS_VIEW'),
('admin', 'EXP_CREATE'),
('admin', 'EXP_VIEW'),
('admin', 'FOD_CREATE'),
('admin', 'FOD_UP_INFO'),
('admin', 'FOD_UP_STATE'),
('admin', 'FOD_VIEW'),
('admin', 'IMP_CREATE'),
('admin', 'IMP_UP_STATE'),
('admin', 'IMP_VIEW'),
('admin', 'ING_CREATE'),
('admin', 'ING_VIEW'),
('admin', 'INV_CREATE'),
('admin', 'INV_VIEW'),
('admin', 'SUP_CREATE'),
('admin', 'SUP_PAY'),
('admin', 'SUP_UP_INFO'),
('admin', 'SUP_VIEW'),
('admin', 'TOP_CREATE'),
('admin', 'TOP_UP_INFO'),
('admin', 'TOP_UP_STATE'),
('admin', 'TOP_VIEW'),
('admin', 'USE_UP_INFO'),
('admin', 'USE_UP_STATE'),
('admin', 'USE_VIEW'),
('fl_9lf4Ig', 'IMP_CREATE'),
('user', 'CAN_CREATE'),
('user', 'CAT_CREATE'),
('user', 'CAT_UP_INFO'),
('user', 'CUS_CREATE'),
('user', 'CUS_PAY'),
('user', 'CUS_UP_INFO'),
('user', 'EXP_CREATE'),
('user', 'FOD_CREATE'),
('user', 'FOD_UP_INFO'),
('user', 'FOD_UP_STATE'),
('user', 'IMP_CREATE'),
('user', 'IMP_UP_STATE'),
('user', 'ING_CREATE'),
('user', 'INV_CREATE'),
('user', 'SUP_CREATE'),
('user', 'SUP_PAY'),
('user', 'SUP_UP_INFO'),
('user', 'TOP_CREATE'),
('user', 'TOP_UP_INFO'),
('user', 'TOP_UP_STATE'),
('user', 'USE_UP_INFO'),
('user', 'USE_UP_STATE');

INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('foodtest', 'evzZ1o4SR', 'size M', 16000, 20000, '6vkZJTVIRz');
INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('foodtest', 'MGxpATVSR', 'size SSS', 123, 123, 'GMbtAo4IRz');
INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('J1F4ZEVIg', 'J1K4WEVIgz', 'size S', 12000, 16000, '1JKVZPVSRm');
INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('J1F4ZEVIg', 'J1KVWPVIRZ', 'size M', 16000, 20000, 'J1F4WPVIRM');

INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('babc', 'bug ơi là bug', 'd@gmail.com', '1111111112', 100048);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('gAoHFtVIR', '123', '', '1234567890', 0);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('KxFFFt4Sg', '123', 'a@gmail.com', '12345678901', 0);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('rs2_QiVIg', 'tên đã sửa', 'a@gmail.com', '1111111111', 1114910),
('tjRaQEVSR', 'oos hoos hoos', 'a@gmail.com', '12345678902', 0);

INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('1O-zf04IR', 'rs2_QiVIg', 5300, 1114910, 'Debt', '2023-11-04 19:08:34', 'za1u8m4Sg');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('a5OraA4Ig', 'rs2_QiVIg', 4300, 1108610, 'Debt', '2023-11-04 19:05:41', 'za1u8m4Sg');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('Fd35FpVSg', 'babc', 10, 100028, 'Pay', '2023-11-03 12:39:45', 'za1u8m4Sg');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createAt`, `createBy`) VALUES
('Gi6yFnVSg', 'rs2_QiVIg', 1100010, 1100010, 'Debt', '2023-10-31 10:11:51', 'za1u8m4Sg'),
('u_JfK7VSR', 'babc', 100010, 100010, 'Debt', '2023-10-31 10:09:59', 'za1u8m4Sg'),
('vHH2FpVIR', 'babc', 10, 100048, 'Pay', '2023-11-03 12:40:28', 'za1u8m4Sg'),
('vVf2fAVIg', 'rs2_QiVIg', 5300, 1120210, 'Debt', '2023-11-04 19:15:41', 'za1u8m4Sg'),
('x8WlFnVIR', 'babc', -1, 100009, 'Pay', '2023-10-31 10:13:01', 'za1u8m4Sg'),
('XBH5Kt4Ig', 'babc', 9, 100018, 'Pay', '2023-11-03 12:39:23', 'za1u8m4Sg'),
('z8FpKt4IR', 'babc', 10, 100038, 'Pay', '2023-11-03 12:39:59', 'za1u8m4Sg');

INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('ZNWdVb4Ig', 'topping7', '', '', 'ZHZd4bVSgz', 1, 1000, 10000);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;