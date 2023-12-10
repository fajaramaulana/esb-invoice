/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : localhost:3306
 Source Schema         : esb-invoice

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 10/12/2023 23:21:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for customers
-- ----------------------------
DROP TABLE IF EXISTS `customers`;
CREATE TABLE `customers`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `email` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `address` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_customers_email`(`email`) USING BTREE,
  INDEX `idx_customers_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customers
-- ----------------------------
INSERT INTO `customers` VALUES (1, 'Barrington Publishers', 'barringtonpublisher@gmail.com', '17 Great Suffolk Street London SE1 0NS United Kingdom', '2023-12-10 23:21:12.446', '2023-12-10 23:21:12.446', NULL);
INSERT INTO `customers` VALUES (2, 'Fajar Agus Maulana', 'fajaragusmaulana@gmail.com', 'Kelapa Dua, Tangerang', '2023-12-10 23:21:12.446', '2023-12-10 23:21:12.446', NULL);

-- ----------------------------
-- Table structure for invoice_items
-- ----------------------------
DROP TABLE IF EXISTS `invoice_items`;
CREATE TABLE `invoice_items`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `invoice_id` bigint(20) UNSIGNED NOT NULL,
  `item_id` bigint(20) UNSIGNED NOT NULL,
  `quantity` decimal(10, 2) NOT NULL,
  `unit_price` decimal(10, 2) NOT NULL,
  `total_price` decimal(10, 2) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_invoice_items_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_invoice_items_item`(`item_id`) USING BTREE,
  CONSTRAINT `fk_invoice_items_item` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of invoice_items
-- ----------------------------

-- ----------------------------
-- Table structure for invoices
-- ----------------------------
DROP TABLE IF EXISTS `invoices`;
CREATE TABLE `invoices`  (
  `invoice_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `subject` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `issue_date` datetime(3) NOT NULL,
  `due_date` datetime(3) NOT NULL,
  `customer_id` bigint(20) NOT NULL,
  `payment_status` int(1) NOT NULL DEFAULT 0 COMMENT '\'1 = paid, 0 = unpaid\'',
  `total_item` decimal(10, 2) NOT NULL,
  `subtotal` decimal(10, 2) NOT NULL,
  `tax_rate` decimal(10, 2) NOT NULL,
  `tax_amount` decimal(10, 2) NOT NULL,
  `total_amount` decimal(10, 2) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`invoice_id`) USING BTREE,
  INDEX `idx_invoices_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_customers_invoices`(`customer_id`) USING BTREE,
  CONSTRAINT `fk_customers_invoices` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of invoices
-- ----------------------------

-- ----------------------------
-- Table structure for items
-- ----------------------------
DROP TABLE IF EXISTS `items`;
CREATE TABLE `items`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` longtext CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `price` decimal(10, 2) NOT NULL,
  `type_id` bigint(20) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_items_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_items_type`(`type_id`) USING BTREE,
  CONSTRAINT `fk_items_type` FOREIGN KEY (`type_id`) REFERENCES `types` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of items
-- ----------------------------
INSERT INTO `items` VALUES (1, 'Design', 41.00, 1, '2023-12-10 23:21:12.456', '2023-12-10 23:21:12.456', NULL);
INSERT INTO `items` VALUES (2, 'Development', 57.00, 1, '2023-12-10 23:21:12.456', '2023-12-10 23:21:12.456', NULL);
INSERT INTO `items` VALUES (3, 'Meetings', 4.50, 1, '2023-12-10 23:21:12.456', '2023-12-10 23:21:12.456', NULL);
INSERT INTO `items` VALUES (4, 'Printer', 22.00, 2, '2023-12-10 23:21:12.456', '2023-12-10 23:21:12.456', NULL);
INSERT INTO `items` VALUES (5, 'Monitor', 29.70, 2, '2023-12-10 23:21:12.456', '2023-12-10 23:21:12.456', NULL);

-- ----------------------------
-- Table structure for types
-- ----------------------------
DROP TABLE IF EXISTS `types`;
CREATE TABLE `types`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_types_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of types
-- ----------------------------
INSERT INTO `types` VALUES (1, 'Service', '2023-12-10 23:21:12.451', '2023-12-10 23:21:12.451', NULL);
INSERT INTO `types` VALUES (2, 'Hardware', '2023-12-10 23:21:12.451', '2023-12-10 23:21:12.451', NULL);

SET FOREIGN_KEY_CHECKS = 1;
