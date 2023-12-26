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
  PRIMARY KEY (`exportNoteId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `ExportNoteDetail_ibfk_1` FOREIGN KEY (`exportNoteId`) REFERENCES `ExportNote` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ExportNoteDetail_ibfk_2` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
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



INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('PN001', 'SupCake0001', 2035000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-11-26 04:27:41', '2023-12-26 05:19:16');
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('PN002', 'SupCoffe0001', 2410000, 'InProgress', 'g3W21A7SR', NULL, '2023-12-01 04:32:04', NULL);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('PN003', 'SupIceCr0001', 609500, 'InProgress', 'g3W21A7SR', NULL, '2023-12-26 04:37:49', NULL);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('PN004', 'SupPearl0001', 2499000, 'Cancel', 'g3W21A7SR', 'g3W21A7SR', '2023-12-22 04:43:03', '2023-12-26 05:23:07'),
('PN005', 'SupMilk0001', 706000, 'Cancel', 'g3W21A7SR', 'g3W21A7SR', '2023-10-08 04:45:50', '2023-12-26 05:19:36'),
('PN006', 'SupCacao0001', 1198000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-09-08 05:04:14', '2023-12-26 05:20:04'),
('PN007', 'SupOTea0001', 415000, 'InProgress', 'g3W21A7SR', NULL, '2023-12-26 05:08:05', NULL),
('PN008', 'SupCake0001', 1188000, 'InProgress', 'g3W21A7SR', NULL, '2023-12-26 05:11:35', NULL),
('PN009', 'SupPearl0001', 621000, 'InProgress', 'g3W21A7SR', NULL, '2023-12-26 05:12:35', NULL),
('PN010', 'SupTea0001', 363000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-12-10 05:13:11', '2023-12-26 05:20:25'),
('PN011', 'SupMilk0001', 320000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-08-06 05:15:18', '2023-12-26 05:21:04');

INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`) VALUES
('PN001', 'Ing0001', 45, 5000, 225000);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`) VALUES
('PN001', 'Ing0004', 45, 10000, 450000);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`) VALUES
('PN001', 'Ing0009', 920, 1000, 920000);
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`) VALUES
('PN001', 'Ing0010', 58, 5000, 290000),
('PN001', 'Ing0014', 150, 1000, 150000),
('PN002', 'Ing0002', 230, 5000, 1150000),
('PN002', 'Ing0013', 370, 3000, 1110000),
('PN002', 'Ing0018', 150, 1000, 150000),
('PN003', 'Ing0003', 25, 5000, 125000),
('PN003', 'Ing0005', 2500, 20, 50000),
('PN003', 'Ing0006', 120, 1000, 120000),
('PN003', 'Ing0007', 1200, 200, 240000),
('PN003', 'Ing0008', 42, 1000, 42000),
('PN003', 'Ing0012', 65, 500, 32500),
('PN004', 'Ing0011', 1250, 900, 1125000),
('PN004', 'Ing0016', 80, 4500, 360000),
('PN004', 'Ing0017', 1300, 780, 1014000),
('PN005', 'Ing0001', 45, 6000, 270000),
('PN005', 'Ing0003', 24, 5000, 120000),
('PN005', 'Ing0005', 2600, 10, 26000),
('PN005', 'Ing0010', 58, 5000, 290000),
('PN006', 'Ing0002', 230, 2000, 460000),
('PN006', 'Ing0006', 110, 5000, 550000),
('PN006', 'Ing0009', 940, 200, 188000),
('PN007', 'Ing0003', 25, 5000, 125000),
('PN007', 'Ing0019', 250, 600, 150000),
('PN007', 'Ing020', 70, 2000, 140000),
('PN008', 'Ing0001', 45, 10000, 450000),
('PN008', 'Ing0003', 25, 2000, 50000),
('PN008', 'Ing0004', 45, 5000, 225000),
('PN008', 'Ing0005', 2500, 30, 75000),
('PN008', 'Ing0006', 120, 1200, 144000),
('PN008', 'Ing0014', 150, 1200, 180000),
('PN008', 'Ing0016', 80, 800, 64000),
('PN009', 'Ing0004', 45, 5000, 225000),
('PN009', 'Ing0007', 1200, 200, 240000),
('PN009', 'Ing0010', 58, 2000, 116000),
('PN009', 'Ing0016', 80, 500, 40000),
('PN010', 'Ing0019', 250, 1200, 300000),
('PN010', 'Ing020', 70, 900, 63000),
('PN011', 'Ing0003', 24, 10000, 240000),
('PN011', 'Ing0008', 40, 2000, 80000);

INSERT INTO `Ingredient` (`id`, `name`, `amount`, `measureType`, `price`) VALUES
('Ing0001', 'Đường', 5500, 'Weight', 45);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `measureType`, `price`) VALUES
('Ing0002', 'Hạt Cà Phê', 3000, 'Weight', 230);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `measureType`, `price`) VALUES
('Ing0003', 'Sữa', 10200, 'Volume', 25);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `measureType`, `price`) VALUES
('Ing0004', 'Bột Bánh', 10300, 'Weight', 45),
('Ing0005', 'Trứng', 50, 'Unit', 2500),
('Ing0006', 'Sô Cô La', 5200, 'Weight', 120),
('Ing0007', 'Tinh Dầu Vanilla', 10, 'Volume', 1200),
('Ing0008', 'Nước Dừa', 2500, 'Volume', 42),
('Ing0009', 'Bột Matcha', 1250, 'Weight', 920),
('Ing0010', 'Đường Nâu', 5300, 'Weight', 58),
('Ing0011', 'Dầu Olive', 150, 'Volume', 1250),
('Ing0012', 'Nước Cốt Dừa', 300, 'Volume', 65),
('Ing0013', 'Hạt Hạnh Nhân', 120, 'Weight', 370),
('Ing0014', 'Bơ', 1080, 'Weight', 150),
('Ing0016', 'Bột Baking Soda', 10, 'Weight', 80),
('Ing0017', 'Dầu Hạt Dẻ Cười', 30, 'Volume', 1300),
('Ing0018', 'Hạt Hồ Lô', 50, 'Weight', 150),
('Ing0019', 'Trà Ô long', 1200, 'Weight', 250),
('Ing020', 'Trà Đen', 900, 'Weight', 70);









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




INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupCacao0001', 'NCC Cacao', 'cacao@gmail.com', '0905555555', -1248000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupCake0001', 'NCC Bánh', 'banh@gmail.com', '0943334445', -2035000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupCoffe0001', 'NCC Cà Phê', 'caphe@gmail.com', '0901234567', -80000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupHoney0001', 'NCC Mật Ong', 'matong@gmail.com', '0927777777', -2000),
('SupIceCr0001', 'NCC Kem', 'kem@gmail.com', '0999999999', -60000),
('SupMilk0001', 'NCC Sữa Tuyệt trùng', 'suatuyettrung@gmail.com', '0919876543', -325000),
('SupOTea0001', 'NCC Trà Ôlong', 'olong@gmail.com', '0922333445', -30000),
('SupPearl0001', 'NCC Trân Châu', 'tranchau@gmail.com', '0911122334', -3500),
('SupSugar0001', 'NCC Đường', 'duong@gmail.com', '0921112223', -30000),
('SupTea0001', 'NCC Trà', 'tra@gmail.com', '0922233445', -383000);

INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('PN001', 'SupCake0001', -2035000, -2035000, 'Debt', '2023-12-26 05:16:43', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('PN006', 'SupCacao0001', -1198000, -1248000, 'Debt', '2023-12-26 05:17:40', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('PN010', 'SupTea0001', -363000, -383000, 'Debt', '2023-12-26 05:17:23', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('PN011', 'SupMilk0001', -320000, -325000, 'Debt', '2023-12-26 05:15:59', 'g3W21A7SR');

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