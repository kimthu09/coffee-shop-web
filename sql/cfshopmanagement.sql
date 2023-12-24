/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

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
  PRIMARY KEY (`foodId`,`categoryId`),
  KEY `categoryId` (`categoryId`),
  CONSTRAINT `CategoryFood_ibfk_1` FOREIGN KEY (`foodId`) REFERENCES `Food` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `CategoryFood_ibfk_2` FOREIGN KEY (`categoryId`) REFERENCES `Category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Customer`;
CREATE TABLE `Customer` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `point` float DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ExportNote`;
CREATE TABLE `ExportNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reason` enum('Damaged','OutOfDate') DEFAULT NULL,
  `createdBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `ExportNote_ibfk_1` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ExportNoteDetail`;
CREATE TABLE `ExportNoteDetail` (
  `exportNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amountExport` int DEFAULT '0',
  PRIMARY KEY (`exportNoteId`,`ingredientId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Feature`;
CREATE TABLE `Feature` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `groupName` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
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
  `totalPrice` int DEFAULT '0',
  `status` enum('InProgress','Done','Cancel') DEFAULT 'InProgress',
  `createdBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `closedBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `closedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `closedBy` (`closedBy`),
  KEY `supplierId` (`supplierId`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `ImportNote_ibfk_1` FOREIGN KEY (`closedBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNote_ibfk_2` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNote_ibfk_3` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ImportNoteDetail`;
CREATE TABLE `ImportNoteDetail` (
  `importNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` float NOT NULL,
  `amountImport` int NOT NULL,
  `totalUnit` float NOT NULL,
  PRIMARY KEY (`importNoteId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `ImportNoteDetail_ibfk_1` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Ingredient`;
CREATE TABLE `Ingredient` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `amount` int DEFAULT '0',
  `measureType` enum('Weight','Volume','Unit') NOT NULL,
  `price` float NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InventoryCheckNote`;
CREATE TABLE `InventoryCheckNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amountDifferent` int NOT NULL,
  `amountAfterAdjust` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createdBy` varchar(12) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `InventoryCheckNote_ibfk_1` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InventoryCheckNoteDetail`;
CREATE TABLE `InventoryCheckNoteDetail` (
  `inventoryCheckNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) NOT NULL,
  `initial` int NOT NULL,
  `difference` int NOT NULL,
  `final` int NOT NULL,
  PRIMARY KEY (`inventoryCheckNoteId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `InventoryCheckNoteDetail_ibfk_1` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `InventoryCheckNoteDetail_ibfk_2` FOREIGN KEY (`inventoryCheckNoteId`) REFERENCES `InventoryCheckNote` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Invoice`;
CREATE TABLE `Invoice` (
  `id` varchar(13) NOT NULL,
  `customerId` varchar(13) NOT NULL,
  `totalPrice` int NOT NULL,
  `amountReceived` int NOT NULL,
  `amountPriceUsePoint` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createdBy` varchar(13) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `customerId` (`customerId`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `Invoice_ibfk_1` FOREIGN KEY (`customerId`) REFERENCES `Customer` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `Invoice_ibfk_2` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InvoiceDetail`;
CREATE TABLE `InvoiceDetail` (
  `invoiceId` varchar(13) NOT NULL,
  `foodId` varchar(13) NOT NULL,
  `sizeName` varchar(13) NOT NULL,
  `amount` int NOT NULL,
  `unitPrice` int NOT NULL,
  `description` text NOT NULL,
  `toppings` json NOT NULL,
  KEY `invoiceId` (`invoiceId`) USING BTREE,
  KEY `foodId` (`foodId`),
  CONSTRAINT `InvoiceDetail_ibfk_1` FOREIGN KEY (`invoiceId`) REFERENCES `Invoice` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `InvoiceDetail_ibfk_2` FOREIGN KEY (`foodId`) REFERENCES `Food` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `MUser`;
CREATE TABLE `MUser` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `roleId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `roleId` (`roleId`),
  CONSTRAINT `MUser_ibfk_1` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Recipe`;
CREATE TABLE `Recipe` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `RecipeDetail`;
CREATE TABLE `RecipeDetail` (
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amountNeed` float NOT NULL,
  PRIMARY KEY (`recipeId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `RecipeDetail_ibfk_1` FOREIGN KEY (`recipeId`) REFERENCES `Recipe` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `RecipeDetail_ibfk_2` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Role`;
CREATE TABLE `Role` (
  `id` varchar(13) NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `RoleFeature`;
CREATE TABLE `RoleFeature` (
  `roleId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `featureId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`roleId`,`featureId`),
  KEY `featureId` (`featureId`),
  CONSTRAINT `RoleFeature_ibfk_1` FOREIGN KEY (`featureId`) REFERENCES `Feature` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `RoleFeature_ibfk_2` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ShopGeneral`;
CREATE TABLE `ShopGeneral` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` text NOT NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `wifiPass` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `accumulatePointPercent` float NOT NULL,
  `usePointPercent` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SizeFood`;
CREATE TABLE `SizeFood` (
  `foodId` varchar(12) NOT NULL,
  `sizeId` varchar(12) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `cost` int NOT NULL,
  `price` int NOT NULL,
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`foodId`,`sizeId`),
  KEY `recipeId` (`recipeId`),
  CONSTRAINT `SizeFood_ibfk_1` FOREIGN KEY (`foodId`) REFERENCES `Food` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `SizeFood_ibfk_2` FOREIGN KEY (`recipeId`) REFERENCES `Recipe` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Supplier`;
CREATE TABLE `Supplier` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `debt` int DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SupplierDebt`;
CREATE TABLE `SupplierDebt` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `supplierId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amount` int NOT NULL,
  `amountLeft` int NOT NULL,
  `type` enum('Debt','Pay') NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createdBy` varchar(9) NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `supplierId` (`supplierId`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `SupplierDebt_ibfk_1` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `SupplierDebt_ibfk_2` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Topping`;
CREATE TABLE `Topping` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` text NOT NULL,
  `cookingGuide` text NOT NULL,
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  `cost` int NOT NULL,
  `price` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `recipeId` (`recipeId`),
  CONSTRAINT `Topping_ibfk_1` FOREIGN KEY (`recipeId`) REFERENCES `Recipe` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;





INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('cs1', '123', 'a@gmail.com', '01234567892', 1444430);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('cs2', '123', 'a@gmail.com', '01234567891', 722214);


INSERT INTO `ExportNote` (`id`, `reason`, `createdBy`, `createdAt`) VALUES
('118xALVIR', NULL, 'za1u8m4Sg', '2023-11-06 16:44:14');
INSERT INTO `ExportNote` (`id`, `reason`, `createdBy`, `createdAt`) VALUES
('7SubAL4Ig', NULL, 'za1u8m4Sg', '2023-11-06 16:44:17');
INSERT INTO `ExportNote` (`id`, `reason`, `createdBy`, `createdAt`) VALUES
('9dPbAL4Ig', NULL, 'za1u8m4Sg', '2023-11-06 16:44:12');
INSERT INTO `ExportNote` (`id`, `reason`, `createdBy`, `createdAt`) VALUES
('dadasdasdsad', NULL, 'za1u8m4Sg', '2023-11-05 03:03:17'),
('DRlxAYVSg', NULL, 'za1u8m4Sg', '2023-11-06 16:44:16'),
('jp9xAL4Ig', NULL, 'za1u8m4Sg', '2023-11-06 16:44:18'),
('Muux0YVSR', NULL, 'za1u8m4Sg', '2023-11-06 16:44:18'),
('oC9x0LVIg', NULL, 'za1u8m4Sg', '2023-11-06 16:44:19'),
('yTjb0L4Sg', NULL, 'za1u8m4Sg', '2023-11-06 16:44:19'),
('YYlbAY4SR', NULL, 'za1u8m4Sg', '2023-11-06 16:44:16'),
('ZcQb0L4IR', NULL, 'za1u8m4Sg', '2023-11-06 16:44:15');

INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`) VALUES
('118xALVIR', 'nvl1', 1);
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`) VALUES
('118xALVIR', 'nvl2', 1);
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`) VALUES
('7SubAL4Ig', 'nvl1', 1);
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`) VALUES
('7SubAL4Ig', 'nvl2', 1),
('9dPbAL4Ig', 'nvl1', 1),
('9dPbAL4Ig', 'nvl2', 1),
('dadasdasdsad', 'nvl1', 10),
('dadasdasdsad', 'nvl2', 100),
('DRlxAYVSg', 'nvl1', 1),
('DRlxAYVSg', 'nvl2', 1),
('jp9xAL4Ig', 'nvl1', 1),
('jp9xAL4Ig', 'nvl2', 1),
('Muux0YVSR', 'nvl1', 1),
('Muux0YVSR', 'nvl2', 1),
('oC9x0LVIg', 'nvl1', 1),
('oC9x0LVIg', 'nvl2', 1),
('yTjb0L4Sg', 'nvl1', 1),
('yTjb0L4Sg', 'nvl2', 1),
('YYlbAY4SR', 'nvl1', 1),
('YYlbAY4SR', 'nvl2', 1),
('ZcQb0L4IR', 'nvl1', 1),
('ZcQb0L4IR', 'nvl2', 1);

INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CAT_CREATE', 'Tạo danh mục', 'Danh mục');
INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CAT_UP_INFO', 'Chỉnh sửa thông tin danh mục', 'Danh mục');
INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CAT_VIEW', 'Xem danh mục', 'Danh mục');
INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CUS_CREATE', 'Tạo khách hàng', 'Khách hàng'),
('CUS_UP_INFO', 'Chỉnh sửa thông tin khách hàng', 'Khách hàng'),
('CUS_VIEW', 'Xem khách hàng', 'Khách hàng'),
('EXP_CREATE', 'Tạo phiếu xuất', 'Phiếu xuất'),
('EXP_VIEW', 'Xem phiếu xuất', 'Phiếu xuất'),
('FOD_CREATE', 'Tạo sản phẩm', 'Sản phẩm'),
('FOD_UP_INFO', 'Chỉnh sửa thông tin sản phẩm', 'Sản phẩm'),
('FOD_UP_STATE', 'Chỉnh sửa trạng thái sản phẩm', 'Sản phẩm'),
('FOD_VIEW', 'Xem sản phẩm', 'Sản phẩm'),
('ICN_CREATE', 'Tạo phiếu kiểm kho', 'Phiếu kiểm kho'),
('ICN_VIEW', 'Xem phiếu kiểm kho', 'Phiếu kiểm kho'),
('IMP_CREATE', 'Tạo phiếu nhập', 'Phiếu nhập'),
('IMP_UP_STATE', 'Chỉnh sửa trạng thái phiếu nhập', 'Phiếu nhập'),
('IMP_VIEW', 'Xem phiếu nhập', 'Phiếu nhập'),
('ING_CREATE', 'Tạo nguyên liệu', 'Nguyên liệu'),
('ING_VIEW', 'Xem nguyên liệu', 'Nguyên liệu'),
('INV_CREATE', 'Tạo hóa đơn', 'Hóa đơn'),
('INV_VIEW', 'Xem hóa đơn', 'Hóa đơn'),
('SUP_CREATE', 'Tạo nhà cung cấp', 'Nhà cung cấp'),
('SUP_PAY', 'Trả nợ nhà cung cấp', 'Nhà cung cấp'),
('SUP_UP_INFO', 'Chỉnh sửa thông tin nhà cung cấp', 'Nhà cung cấp'),
('SUP_VIEW', 'Xem nhà cung cấp', 'Nhà cung cấp'),
('TOP_CREATE', 'Tạo topping', 'Topping'),
('TOP_UP_INFO', 'Chỉnh sửa thông tin topping', 'Topping'),
('TOP_UP_STATE', 'Chỉnh sửa trạng thái topping', 'Topping'),
('TOP_VIEW', 'Xem topping', 'Topping'),
('USE_UP_INFO', 'Chỉnh sửa thông tin người dùng', 'Người dùng'),
('USE_UP_STATE', 'Chỉnh sửa trạng thái người dùng', 'Người dùng'),
('USE_VIEW', 'Xem người dùng', 'Người dùng');

INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `isActive`) VALUES
('foodtest', 'foodtest', '123', '', 1);
INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `isActive`) VALUES
('J1F4ZEVIg', 'foodtest1', '123', '', 1);












INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountPriceUsePoint`, `createdAt`, `createdBy`) VALUES
('4ahGrPVSg', 'cs2', 401230, 40123, 361107, '2023-11-07 05:25:23', 'za1u8m4Sg');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountPriceUsePoint`, `createdAt`, `createdBy`) VALUES
('7qFW9PVSR', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:47', 'za1u8m4Sg');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountPriceUsePoint`, `createdAt`, `createdBy`) VALUES
('CkdWrE4Sg', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:45', 'za1u8m4Sg');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `amountReceived`, `amountPriceUsePoint`, `createdAt`, `createdBy`) VALUES
('GL0GrE4Ig', 'cs2', 401230, 40123, 361107, '2023-11-07 05:25:25', 'za1u8m4Sg'),
('HvFWrP4Sg', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:47', 'za1u8m4Sg'),
('nf5WrEVSR', 'cs1', 401230, 40123, 361107, '2023-11-07 05:24:48', 'za1u8m4Sg');



INSERT INTO `MUser` (`id`, `name`, `email`, `password`, `salt`, `roleId`, `isActive`) VALUES
('g3W21A7SR', 'Nguyễn Văn A', 'admin@gmail.com', '5e107317df151f6e8e0015c4f2ee7936', 'mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms', 'admin', 1);
INSERT INTO `MUser` (`id`, `name`, `email`, `password`, `salt`, `roleId`, `isActive`) VALUES
('za1u8m4Sg', 'Nguyễn Văn U', 'user@gmail.com', 'cb58ac5a2272517d1960565444bde187', 'QYlnGKRgYBxIXzMnnQSVcglbtjPsAhVlxMRMDaqnaquxwADSur', 'user', 1);


INSERT INTO `Recipe` (`id`) VALUES
('1JKVZPVSRm');
INSERT INTO `Recipe` (`id`) VALUES
('6vkZJTVIRz');
INSERT INTO `Recipe` (`id`) VALUES
('GMbtAo4IRz');
INSERT INTO `Recipe` (`id`) VALUES
('J1F4WPVIRM'),
('ZHZd4bVSgz');



INSERT INTO `Role` (`id`, `name`) VALUES
('admin', 'admin');
INSERT INTO `Role` (`id`, `name`) VALUES
('user', 'user');


INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_CREATE');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_UP_INFO');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_VIEW');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CUS_CREATE'),
('admin', 'CUS_UP_INFO'),
('admin', 'CUS_VIEW'),
('admin', 'EXP_CREATE'),
('admin', 'EXP_VIEW'),
('admin', 'FOD_CREATE'),
('admin', 'FOD_UP_INFO'),
('admin', 'FOD_UP_STATE'),
('admin', 'FOD_VIEW'),
('admin', 'ICN_CREATE'),
('admin', 'ICN_VIEW'),
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
('admin', 'USE_VIEW');

INSERT INTO `ShopGeneral` (`id`, `name`, `email`, `phone`, `address`, `wifiPass`, `accumulatePointPercent`, `usePointPercent`) VALUES
('shop', 'Coffee shop', '', '', '', 'coffeeshop123', 0.001, 1);


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

INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('1O-zf04IR', 'rs2_QiVIg', 5300, 1114910, 'Debt', '2023-11-04 19:08:34', 'za1u8m4Sg');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('a5OraA4Ig', 'rs2_QiVIg', 4300, 1108610, 'Debt', '2023-11-04 19:05:41', 'za1u8m4Sg');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('Fd35FpVSg', 'babc', 10, 100028, 'Pay', '2023-11-03 12:39:45', 'za1u8m4Sg');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('Gi6yFnVSg', 'rs2_QiVIg', 1100010, 1100010, 'Debt', '2023-10-31 10:11:51', 'za1u8m4Sg'),
('u_JfK7VSR', 'babc', 100010, 100010, 'Debt', '2023-10-31 10:09:59', 'za1u8m4Sg'),
('vHH2FpVIR', 'babc', 10, 100048, 'Pay', '2023-11-03 12:40:28', 'za1u8m4Sg'),
('vVf2fAVIg', 'rs2_QiVIg', 5300, 1120210, 'Debt', '2023-11-04 19:15:41', 'za1u8m4Sg'),
('x8WlFnVIR', 'babc', -1, 100009, 'Pay', '2023-10-31 10:13:01', 'za1u8m4Sg'),
('XBH5Kt4Ig', 'babc', 9, 100018, 'Pay', '2023-11-03 12:39:23', 'za1u8m4Sg'),
('z8FpKt4IR', 'babc', 10, 100038, 'Pay', '2023-11-03 12:39:59', 'za1u8m4Sg');


DELIMITER //
CREATE TRIGGER update_closedAt
BEFORE UPDATE ON ImportNote
FOR EACH ROW
BEGIN
    IF NEW.status != 'InProgress' THEN
        SET NEW.closedAt = CURRENT_TIMESTAMP;
    END IF;
END;
//
DELIMITER ;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;