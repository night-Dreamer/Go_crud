/*
 Navicat Premium Data Transfer

 Source Server         : Postgresql
 Source Server Type    : PostgreSQL
 Source Server Version : 140001
 Source Host           : localhost:5432
 Source Catalog        : company
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140001
 File Encoding         : 65001

 Date: 26/02/2022 19:42:06
*/


-- ----------------------------
-- Table structure for staff
-- ----------------------------
DROP TABLE IF EXISTS "public"."staff";
CREATE TABLE "public"."staff" (
  "id" varchar(15) COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar(30) COLLATE "pg_catalog"."default",
  "position" varchar(255) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Records of staff
-- ----------------------------
INSERT INTO "public"."staff" VALUES ('001', 'Jack', '管理员');
INSERT INTO "public"."staff" VALUES ('003', 'Bob', '职员');
INSERT INTO "public"."staff" VALUES ('004', 'Micky', '职员');
INSERT INTO "public"."staff" VALUES ('002', 'Tom', '管理员');

-- ----------------------------
-- Primary Key structure for table staff
-- ----------------------------
ALTER TABLE "public"."staff" ADD CONSTRAINT "staff_pkey" PRIMARY KEY ("id");
