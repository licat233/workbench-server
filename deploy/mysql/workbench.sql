-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2022-04-06 23:27:15
-- 服务器版本： 8.0.24
-- PHP 版本： 7.4.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `workbench`
--

-- --------------------------------------------------------

--
-- 表的结构 `ad_account`
--

CREATE TABLE `ad_account` (
  `aaid` int UNSIGNED NOT NULL COMMENT '账户id',
  `ad_account_channel_id` int NOT NULL DEFAULT '0' COMMENT '所属渠道id，外键',
  `ad_account_passport` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `ad_account_password` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `ad_account_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '账号名字',
  `ad_account_description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '说明',
  `ad_account_status` tinyint NOT NULL DEFAULT '1' COMMENT '账号状态，默认值1',
  `ad_account_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `ad_account_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ad_account_cookie` text COLLATE utf8mb4_unicode_ci COMMENT 'cookie',
  `ad_account_token` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'token',
  `ad_account_read_code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '阅读代码',
  `ad_account_cv_code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '转化代码',
  `ad_account_other_code` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '其它代码',
  `ad_account_js_file` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'Js文件地址,如果有多个,请以英文";"号分开'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告账户表';

--
-- 转存表中的数据 `ad_account`
--

INSERT INTO `ad_account` (`aaid`, `ad_account_channel_id`, `ad_account_passport`, `ad_account_password`, `ad_account_name`, `ad_account_description`, `ad_account_status`, `ad_account_create_time`, `ad_account_update_time`, `ad_account_cookie`, `ad_account_token`, `ad_account_read_code`, `ad_account_cv_code`, `ad_account_other_code`, `ad_account_js_file`) VALUES
(1, 0, 'Self_White', 'Self_White888...', 'popin祛斑1户', '', 1, '2021-12-24 17:41:37', '2021-12-27 10:21:55', 'var statistic = conversion = () => {\n    window._pop = window._pop || [];\n    window._pop.push([\"_set_conversion_id\", \"Self_White\"]);\n    window._pop.push([\"_set_action\", \"cv\"]);\n\n    let el = document.createElement(\"script\");\n    el.type = \"text/javascript\";\n    el.async = true;\n    el.src = ((\"https:\" == document.location.protocol) ? \"https://\" : \"http://\") + \"api.popin.cc/conversion2.js\";\n    let s = document.getElementsByTagName(\"script\")[0];\n    s.parentNode.insertBefore(el, s);\n    console.log(\"conversion success\")\n}\ntry {\n    $$.statistic = statistic;\n} catch (e) {}\n\n! function () {\n    let pa = document.createElement(\'script\');\n    pa.type = \'text/javascript\';\n    pa.charset = \"utf-8\";\n    pa.async = true;\n    pa.src = window.location.protocol + \"//api.popin.cc/ads/Self_White.js\";\n    let s = document.getElementsByTagName(\'script\')[0];\n    s.parentNode.insertBefore(pa, s);\n}();', NULL, NULL, NULL, '', ''),
(2, 0, 'licat233@gmail.com', '123456789...', 'line直投祛斑1户', '', 0, '2021-12-24 17:41:37', '2021-12-27 10:23:14', '(function (g, d, o) {\n    g._ltq = g._ltq || [];\n    g._lt = g._lt || function () {\n        g._ltq.push(arguments)\n    };\n    var h = location.protocol === \'https:\' ? \'https://d.line-scdn.net\' : \'http://d.line-cdn.net\';\n    var s = d.createElement(\'script\');\n    s.async = 1;\n    s.src = o || h + \'/n/line_tag/public/release/v1/lt.js\';\n    var t = d.getElementsByTagName(\'script\')[0];\n    t.parentNode.insertBefore(s, t);\n})(window, document);\n_lt(\'init\', {\n    customerType: \'lap\',\n    tagId: \'3894da95-7411-4997-8628-d2399eca2834\'\n});\n_lt(\'send\', \'pv\', [\'3894da95-7411-4997-8628-d2399eca2834\']);\n\n$$.statistic = () => {\n    _lt(\'send\', \'cv\', {\n        type: \'Conversion\'\n    }, [\'3894da95-7411-4997-8628-d2399eca2834\']);\n    console.log(\'successful conversion\');\n}', NULL, NULL, NULL, '', '');

-- --------------------------------------------------------

--
-- 表的结构 `ad_campaign`
--

CREATE TABLE `ad_campaign` (
  `id` int UNSIGNED NOT NULL COMMENT '活动id',
  `campaign_uid` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '活动id值，最大64位',
  `adcampaign_status` tinyint NOT NULL DEFAULT '1' COMMENT '状态，默认1',
  `adcampaign_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '活动名称',
  `adcampaign_account_id` int NOT NULL DEFAULT '0' COMMENT '所属账户id，外键',
  `adcampaign_cookie` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'cookie',
  `adcampaign_description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '说明',
  `adcampaign_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `adcampaign_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告活动表';

-- --------------------------------------------------------

--
-- 表的结构 `ad_campaign_cost`
--

CREATE TABLE `ad_campaign_cost` (
  `id` int UNSIGNED NOT NULL COMMENT 'id主键',
  `campaign_id` int NOT NULL COMMENT '活动id，外键',
  `AD_date` datetime NOT NULL COMMENT '广告日期',
  `AD_cost` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '广告消耗',
  `current_cost` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '当日当前消耗',
  `node_cost` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '广告日节点消耗',
  `yday_cost` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '昨天全天消耗',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '插入日期'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='活动消耗表';

-- --------------------------------------------------------

--
-- 表的结构 `ad_channel`
--

CREATE TABLE `ad_channel` (
  `id` int UNSIGNED NOT NULL COMMENT '渠道id，主键',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '渠道昵称',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '说明',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态，默认1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告渠道表';

--
-- 转存表中的数据 `ad_channel`
--

INSERT INTO `ad_channel` (`id`, `name`, `description`, `create_time`, `update_time`, `status`) VALUES
(1, 'POPin', '', '2021-12-24 17:20:58', '2021-12-24 17:20:58', 1),
(2, 'POPin_line', '', '2021-12-24 17:20:58', '2021-12-24 17:38:18', 0),
(3, 'LINE', '', '2021-12-24 17:21:28', '2021-12-24 17:21:28', 1),
(4, 'Dable', '', '2021-12-24 17:21:28', '2021-12-24 17:38:14', 0);

-- --------------------------------------------------------

--
-- 表的结构 `ad_comments`
--

CREATE TABLE `ad_comments` (
  `id` int UNSIGNED NOT NULL,
  `comment_parent_id` int NOT NULL DEFAULT '0' COMMENT '回复对应的上一级回复ID(查找不到，或者没有的话默认为0)',
  `comment_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `comment_content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '留言内容',
  `comment_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `comment_ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Ip地址，每个ip一天最多有5天评论',
  `comment_weburl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '评论的页面url',
  `comment_article_id` int NOT NULL DEFAULT '0' COMMENT '应该为文章表的外键，通过url获取 对应的落地页文章id，若不存在的话，默认为0',
  `comment_article_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文章标题，从网站内ajax上传过来',
  `comment_type` tinyint NOT NULL DEFAULT '0' COMMENT '评论类型：0是评论，1是回复，默认0',
  `like` int NOT NULL DEFAULT '0' COMMENT '点赞人数，最好设置一个ip只能点赞一次，通过redis存储ip信息',
  `website_id` int NOT NULL DEFAULT '0' COMMENT '通过url获取对应的website id，不存在的时候，默认0',
  `parent_content` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '如果parent_id为0，则评论对应的内容从网站内ajax获取上传过来，否则就默认为“”'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='网友留言表，务必做好防sql注入，请传入两个url参数，一个是ad_url，一个是href';

-- --------------------------------------------------------

--
-- 表的结构 `ad_creative`
--

CREATE TABLE `ad_creative` (
  `id` int UNSIGNED NOT NULL,
  `ad_id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '广告创意ID',
  `ad_imp` int NOT NULL COMMENT '广告创意展现数',
  `ad_click` int NOT NULL COMMENT '广告创意点击数',
  `ad_ctr` float NOT NULL COMMENT '创高创意CVR',
  `ad_cv` int NOT NULL COMMENT '广告创意CV',
  `ad_cvr` float NOT NULL COMMENT '广告创意CVR',
  `ad_cpa` decimal(10,2) NOT NULL COMMENT '广告创意CPA',
  `ad_cost` decimal(10,2) NOT NULL COMMENT '广告创意消耗',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `ad_description` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '创意描述',
  `ad_campaign_id` int UNSIGNED DEFAULT NULL COMMENT '所属广告活动id',
  `ad_group_id` int UNSIGNED DEFAULT NULL COMMENT '所属广告组id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告创意表';

-- --------------------------------------------------------

--
-- 表的结构 `ad_group`
--

CREATE TABLE `ad_group` (
  `id` int NOT NULL COMMENT '主键',
  `group_id` varchar(32) CHARACTER SET utf8 NOT NULL COMMENT '广告组id',
  `group_status` tinyint NOT NULL COMMENT '广告组状态，0关闭，1打开，3被封',
  `group_name` varchar(32) CHARACTER SET utf8 NOT NULL COMMENT '广告组名称',
  `group_campaign_uid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '广告组所属活动id',
  `group_cpa` decimal(10,0) NOT NULL COMMENT '广告组CPA',
  `group_max_cap` decimal(10,0) NOT NULL COMMENT '广告组CPA出价',
  `group_description` varchar(200) CHARACTER SET utf8 NOT NULL COMMENT '广告组描述',
  `group_cost` decimal(10,0) NOT NULL COMMENT '广告组消耗',
  `group_ctr` decimal(10,0) NOT NULL COMMENT '广告组点击率',
  `group_click` int NOT NULL COMMENT '广告组点击数',
  `group_cv` int NOT NULL COMMENT '广告组转化数',
  `group_cvr` decimal(10,0) NOT NULL COMMENT '广告组转化率',
  `group_budget` decimal(10,0) NOT NULL COMMENT '广告组预算',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='广告组表';

-- --------------------------------------------------------

--
-- 表的结构 `ad_project_channel`
--

CREATE TABLE `ad_project_channel` (
  `id` int UNSIGNED NOT NULL COMMENT 'id主键',
  `project_id` int UNSIGNED NOT NULL COMMENT '项目id，外键',
  `channel_id` int UNSIGNED NOT NULL COMMENT '渠道id，外键',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '插入时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='项目-渠道 映射表';

-- --------------------------------------------------------

--
-- 表的结构 `article`
--

CREATE TABLE `article` (
  `id` int UNSIGNED NOT NULL COMMENT 'Id主键',
  `title` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '标题',
  `keywords` varchar(127) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'SEO中的关键词',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'SEO中的描述',
  `html` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  `status_id` int NOT NULL DEFAULT '0' COMMENT '状态，0表示未发布/草稿，1表示已发布，2表示已删除，默认 0',
  `release_time` datetime DEFAULT NULL COMMENT '发布时间',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `screenshoot` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '截图路径',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '所属项目id，外键',
  `grade_id` int NOT NULL DEFAULT '0' COMMENT '等级，外键，1优，2良，3中，4差',
  `author_id` int NOT NULL DEFAULT '0' COMMENT '作者id，外键',
  `safehtml` mediumtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '加密html代码',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '对此文章的评论'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='推广文章列表';

-- --------------------------------------------------------

--
-- 表的结构 `article_grade`
--

CREATE TABLE `article_grade` (
  `id` int UNSIGNED NOT NULL COMMENT 'Id外键',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文章等级名称',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '等级描述'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章等级表';

-- --------------------------------------------------------

--
-- 表的结构 `article_status`
--

CREATE TABLE `article_status` (
  `id` int UNSIGNED NOT NULL COMMENT 'id主键',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '状态名称',
  `description` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '关于这个状态的描述'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='文章状态表';

-- --------------------------------------------------------

--
-- 表的结构 `brand`
--

CREATE TABLE `brand` (
  `bid` int UNSIGNED NOT NULL,
  `brand_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '品牌名',
  `brand_teacher_id` int NOT NULL DEFAULT '0' COMMENT '品牌代言老师id',
  `brand_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '联系邮箱',
  `brand_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '品牌创建时间',
  `brand_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '品牌更新时间',
  `brand_official_account_id` int NOT NULL DEFAULT '0' COMMENT '品牌官方账号',
  `brand_description` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '品牌描述',
  `brand_website` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '品牌官网',
  `brand_project_id` int NOT NULL DEFAULT '0' COMMENT '所属项目id，外键',
  `brand_status` int NOT NULL DEFAULT '0' COMMENT '状态，默认0，审核过后才能用',
  `brand_logo` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '品牌logo链接'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='品牌表';

--
-- 转存表中的数据 `brand`
--

INSERT INTO `brand` (`bid`, `brand_name`, `brand_teacher_id`, `brand_email`, `brand_create_time`, `brand_update_time`, `brand_official_account_id`, `brand_description`, `brand_website`, `brand_project_id`, `brand_status`, `brand_logo`) VALUES
(1, '植題-PlantTItle', 1, '', '2021-10-14 16:25:49', '2021-12-21 11:07:47', 0, '', '', 0, 0, 'https://img.alicdn.com/imgextra/i2/917298378/O1CN01WJWT9k2BlAxycWFHs_!!917298378.jpg'),
(2, 'Discoco', 2, '', '2021-12-21 14:17:19', '2021-12-21 14:17:19', 0, '', '', 0, 0, '');

-- --------------------------------------------------------

--
-- 表的结构 `casbin_rule`
--

CREATE TABLE `casbin_rule` (
  `id` bigint UNSIGNED NOT NULL,
  `ptype` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限管理表';

-- --------------------------------------------------------

--
-- 表的结构 `company`
--

CREATE TABLE `company` (
  `cid` int UNSIGNED NOT NULL COMMENT '公司id，主键',
  `company_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '公司名称',
  `company_address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司地址',
  `company_boss_id` int NOT NULL DEFAULT '0' COMMENT '公司负责人id',
  `company_status` tinyint NOT NULL DEFAULT '0' COMMENT '公司状态，默认0，需要审核过后才能用',
  `company_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `company_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `company_description` varchar(1000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司简介',
  `company_logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '公司logo',
  `company_website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='公司数据表';

--
-- 转存表中的数据 `company`
--

INSERT INTO `company` (`cid`, `company_name`, `company_address`, `company_boss_id`, `company_status`, `company_create_time`, `company_update_time`, `company_description`, `company_logo`, `company_website`) VALUES
(0, '未公开', '未公开', 0, 0, '2021-10-11 18:20:53', '2021-12-21 11:11:04', '不方便公开的公司', '', ''),
(1, '广州市享购贸易有限公司', '广州市白云区环洲二路63号四楼自编412A房', 3, 1, '2021-12-18 16:26:59', '2021-12-18 16:27:36', '祛斑', '', ''),
(2, '广州纽德网络科技有限公司', '广州市白云区金沙洲街环洲二路63号7楼707房', 4, 1, '2021-12-18 16:26:59', '2021-12-18 16:27:43', '抗衰老', '', ''),
(3, '其它公司', '其它', 0, 0, '2021-12-21 11:15:29', '2021-12-21 11:15:29', '', '', '');

-- --------------------------------------------------------

--
-- 表的结构 `cost_list`
--

CREATE TABLE `cost_list` (
  `id` int UNSIGNED NOT NULL COMMENT 'Id主键',
  `cost_date` date NOT NULL COMMENT '消耗日期',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `cost_twd` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '消耗台币',
  `cost_cny` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '消耗人民币',
  `fans` int NOT NULL DEFAULT '0' COMMENT '加粉数',
  `cpa` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '每个粉的成本',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `project_id` int NOT NULL DEFAULT '0' COMMENT '所属项目id，外键',
  `channel_id` int NOT NULL DEFAULT '0' COMMENT '所属渠道id，外键',
  `ad_account_id` int NOT NULL DEFAULT '0' COMMENT '所属账户id，外键',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '负责人/操户人_id，外键',
  `brand_id` int NOT NULL DEFAULT '0',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态，默认1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消耗列表';

-- --------------------------------------------------------

--
-- 表的结构 `fans_list`
--

CREATE TABLE `fans_list` (
  `fid` int UNSIGNED NOT NULL COMMENT 'Id主键',
  `fans_date` date NOT NULL COMMENT '加粉日期',
  `fans_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `fans_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `day_fans` int NOT NULL DEFAULT '0' COMMENT '白天的加粉',
  `night_fans` int NOT NULL DEFAULT '0' COMMENT '夜间的加粉',
  `today_fans` int NOT NULL DEFAULT '0' COMMENT '今天全部加粉',
  `active_fans` int NOT NULL DEFAULT '0' COMMENT '活跃的粉丝数',
  `fans_remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `fans_service_account_id` int NOT NULL DEFAULT '0' COMMENT '所属账号id，外键',
  `fans_account_id` int NOT NULL DEFAULT '0' COMMENT '所属广告账户Id，外键',
  `fans_channel_id` int NOT NULL DEFAULT '0' COMMENT '所属广告渠道Id，外键',
  `fans_project_id` int NOT NULL DEFAULT '0' COMMENT '所属项目id，外键',
  `fans_user_id` int NOT NULL DEFAULT '0' COMMENT '流量负责人',
  `fans_status` tinyint NOT NULL DEFAULT '1' COMMENT '状态，默认 1'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='加粉统计表';

-- --------------------------------------------------------

--
-- 表的结构 `feedback`
--

CREATE TABLE `feedback` (
  `id` int UNSIGNED NOT NULL COMMENT 'id主键',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '所属用户id，外键',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '内容',
  `contact` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '联系方式',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '文章查阅状态，默认0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='反馈建议表';

-- --------------------------------------------------------

--
-- 表的结构 `hijack_account`
--

CREATE TABLE `hijack_account` (
  `id` int UNSIGNED NOT NULL COMMENT 'Id主键',
  `url_id` int NOT NULL DEFAULT '0' COMMENT '对应的链接id，外键',
  `account_id` int NOT NULL DEFAULT '0' COMMENT '对应的业务号id，外键',
  `lucky_value` int NOT NULL DEFAULT '100' COMMENT '幸运值',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '上线时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='劫持所用业务号表';

-- --------------------------------------------------------

--
-- 表的结构 `hijack_config`
--

CREATE TABLE `hijack_config` (
  `id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'Id主键',
  `status` tinyint DEFAULT '0' COMMENT '开关状态，默认0',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间，等于开关切换时间',
  `adminer_id` int UNSIGNED DEFAULT '0' COMMENT '管理员id，外键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='流量劫持配置';

-- --------------------------------------------------------

--
-- 表的结构 `hijack_url`
--

CREATE TABLE `hijack_url` (
  `id` int UNSIGNED NOT NULL COMMENT 'id主键',
  `url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '劫持的Url',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `prob` int NOT NULL DEFAULT '0' COMMENT '概率值，百分比，根据秒来划分',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态值，默认 0，需手动开启'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='被劫持的url';

-- --------------------------------------------------------

--
-- 表的结构 `ip_list`
--

CREATE TABLE `ip_list` (
  `id` int UNSIGNED NOT NULL COMMENT 'Id主键',
  `ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'ip地址',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `type_id` int NOT NULL DEFAULT '1' COMMENT '永久白名单2，永久黑名单0，1未设置，默认1',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Ip名单表';

-- --------------------------------------------------------

--
-- 表的结构 `ip_type`
--

CREATE TABLE `ip_type` (
  `id` int UNSIGNED NOT NULL COMMENT 'Id主键',
  `type_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '类型名称，未知、黑名单、白名单'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Ip类型表';

-- --------------------------------------------------------

--
-- 表的结构 `jwt_blacklists`
--

CREATE TABLE `jwt_blacklists` (
  `id` int NOT NULL,
  `uid` int NOT NULL COMMENT '用户id',
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'jwt-token',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '绑定平台',
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '绑定ip',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `expire_time` datetime NOT NULL COMMENT '过期时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='jwt黑名单';

-- --------------------------------------------------------

--
-- 表的结构 `jwt_whitelists`
--

CREATE TABLE `jwt_whitelists` (
  `id` int NOT NULL,
  `uid` int NOT NULL COMMENT '用户id',
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'jwt-token',
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '绑定平台',
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '绑定ip',
  `created_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `expire_time` datetime NOT NULL COMMENT '过期时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='jwt白名单';

-- --------------------------------------------------------

--
-- 表的结构 `landing_page`
--

CREATE TABLE `landing_page` (
  `id` int UNSIGNED NOT NULL,
  `account_id` int UNSIGNED DEFAULT NULL COMMENT '所属广告账号id，外键',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '标题',
  `url` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '链接',
  `article_id` int UNSIGNED DEFAULT NULL COMMENT '文章id，外键',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `js_code` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'Js代码',
  `js_file` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT 'Js文件路径'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='推广网站链接表';

-- --------------------------------------------------------

--
-- 表的结构 `menu`
--

CREATE TABLE `menu` (
  `id` int NOT NULL COMMENT '主键',
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
  `path` varchar(500) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单路径',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '菜单状态，默认关闭0',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `parent_id` int NOT NULL COMMENT '父菜单id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='菜单表';

-- --------------------------------------------------------

--
-- 表的结构 `news`
--

CREATE TABLE `news` (
  `id` int UNSIGNED NOT NULL COMMENT 'id主键',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '通知用户id，外键',
  `content` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '通知内容',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态，默认未读0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='消息通知记录表';

-- --------------------------------------------------------

--
-- 表的结构 `project`
--

CREATE TABLE `project` (
  `pid` int UNSIGNED NOT NULL COMMENT '主键id',
  `project_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '项目名称',
  `project_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '所在位置',
  `project_company_id` int NOT NULL DEFAULT '0' COMMENT '所属公司id',
  `project_admin_id` int NOT NULL DEFAULT '0' COMMENT '项目负责人/项目经理',
  `project_status` tinyint NOT NULL DEFAULT '0' COMMENT '项目状态，默认值0，管理员审核后才能用',
  `project_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `project_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `project_description` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '说明'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='项目表';

--
-- 转存表中的数据 `project`
--

INSERT INTO `project` (`pid`, `project_name`, `project_address`, `project_company_id`, `project_admin_id`, `project_status`, `project_create_time`, `project_update_time`, `project_description`) VALUES
(1, '祛斑', '412', 1, 1, 1, '2021-12-24 16:25:10', '2021-12-24 16:25:10', ''),
(2, '抗衰老', '707', 2, 1, 1, '2021-12-24 16:25:10', '2021-12-24 16:25:10', '');

-- --------------------------------------------------------

--
-- 表的结构 `service_account`
--

CREATE TABLE `service_account` (
  `said` int UNSIGNED NOT NULL COMMENT 'Id主键',
  `sa_user_id` int NOT NULL DEFAULT '0' COMMENT '使用者id，外键',
  `sa_brand_id` int NOT NULL DEFAULT '0' COMMENT '所属品牌id，外键，一般不会修改',
  `sa_account_id` int NOT NULL DEFAULT '0' COMMENT '所在的广告账户id，外键',
  `sa_show_account` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '展示业务号，常用于Line渠道，每个业务号可额外起个昵称',
  `sa_account` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '业务号ID',
  `sa_password` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '业务号密码',
  `sa_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `sa_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `sa_add_link_pc` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '加入好友电脑链接',
  `sa_add_link_mb` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '加入好友手机链接',
  `sa_status` tinyint NOT NULL DEFAULT '0' COMMENT '账号是否可用，0不可用（需要审核），1可用，2满粉了，3出问题了',
  `sa_online` tinyint NOT NULL DEFAULT '0' COMMENT '在线？默认0，默认离线状态',
  `sa_remark` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  `sa_lucky_value` int NOT NULL DEFAULT '0' COMMENT '幸运值，默认 0 ',
  `sa_bind_mobile` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `sa_bind_email` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `sa_type_id` int NOT NULL DEFAULT '0' COMMENT '类型id'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='服务帐户数据表，用于存放line账号';

-- --------------------------------------------------------

--
-- 表的结构 `service_account_type`
--

CREATE TABLE `service_account_type` (
  `id` int UNSIGNED NOT NULL,
  `type_name` varchar(30) CHARACTER SET utf8 NOT NULL,
  `type_remark` varchar(300) CHARACTER SET utf8 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='客服账号类型表';

--
-- 转存表中的数据 `service_account_type`
--

INSERT INTO `service_account_type` (`id`, `type_name`, `type_remark`) VALUES
(1, 'LINE', '台湾LINE账号'),
(2, 'WhatsApp', ''),
(3, 'WeChat', ''),
(4, 'Facebook', '');

-- --------------------------------------------------------

--
-- 表的结构 `sys_log`
--

CREATE TABLE `sys_log` (
  `id` int UNSIGNED NOT NULL,
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求方法',
  `path` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路径，资源',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态，0失败(默认)，1成功',
  `message` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '信息',
  `body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求内容',
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应答内容',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `jwt_token` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '用户jwt_token',
  `platform` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '设备平台'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统操作记录';

-- --------------------------------------------------------

--
-- 表的结构 `sys_role`
--

CREATE TABLE `sys_role` (
  `rid` int UNSIGNED NOT NULL COMMENT '主键id',
  `role_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '角色昵称，分为：游客、客服、客服组长、客服主管、美工、项目经理、文案、推广员、技术员、推广主管、人事、人事主管、会计员、跟单员、总经理',
  `role_status` tinyint NOT NULL DEFAULT '0' COMMENT '状态，默认不可用',
  `role_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `role_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `role_description` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '""' COMMENT '说明',
  `role_icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色图标'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

--
-- 转存表中的数据 `sys_role`
--

INSERT INTO `sys_role` (`rid`, `role_name`, `role_status`, `role_create_time`, `role_update_time`, `role_description`, `role_icon`) VALUES
(0, '未知', 0, '2021-10-11 18:21:52', '2021-11-17 18:49:18', '', ''),
(1, '网推主管', 1, '2021-10-14 12:07:59', '2021-11-08 18:43:28', '超级管理员', ''),
(2, '总经理', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:08', '', ''),
(3, '项目经理', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:09', '', ''),
(4, '客服主管', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:11', '', ''),
(5, '客服组长', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:13', '', ''),
(6, '财务', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:14', '', ''),
(7, '人事主管', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:16', '', ''),
(8, '跟单员', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:17', '', ''),
(9, '文案', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:18', '', ''),
(10, '美工', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:20', '', ''),
(11, '技术', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:22', '', ''),
(12, '客服', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:24', '', ''),
(13, '人事专员', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:28', '', ''),
(14, '前台', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:29', '', ''),
(15, '行政', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:36', '', ''),
(16, '推广专员', 1, '2021-10-14 12:07:59', '2021-11-17 15:50:34', '', '');

-- --------------------------------------------------------

--
-- 表的结构 `teacher`
--

CREATE TABLE `teacher` (
  `tid` int UNSIGNED NOT NULL COMMENT '主键',
  `teacher_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '老师姓名',
  `teacher_email` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '老师邮箱',
  `teacher_mobile` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '老师电话',
  `teacher_create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `teacher_update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `teacher_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '老师头像',
  `teacher_gender` tinyint NOT NULL DEFAULT '1' COMMENT '老师的性别，默认1（女性）',
  `teacher_description` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '老师简介',
  `teacher_website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '老师主页url',
  `teacher_status` tinyint NOT NULL DEFAULT '0' COMMENT '老师状态，默认0（不可用，审核过后才能用）'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='老师数据表';

--
-- 转存表中的数据 `teacher`
--

INSERT INTO `teacher` (`tid`, `teacher_name`, `teacher_email`, `teacher_mobile`, `teacher_create_time`, `teacher_update_time`, `teacher_avatar`, `teacher_gender`, `teacher_description`, `teacher_website`, `teacher_status`) VALUES
(0, '未设置', '未设置', '未设置', '2021-10-14 16:28:00', '2021-10-14 16:28:00', '未设置', 1, NULL, '未设置', 0),
(1, 'Anna', '', '', '2021-10-14 16:26:13', '2021-10-14 16:26:13', '', 1, NULL, '', 0);

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `uid` int UNSIGNED NOT NULL COMMENT 'UID主键',
  `user_real_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '未公开' COMMENT '用户真实姓名',
  `user_avatar` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg' COMMENT '用户头像',
  `user_gender` tinyint NOT NULL DEFAULT '0' COMMENT '用户性别，男是0，女是1',
  `passport` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `user_email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `user_mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '电话号码',
  `user_role_id` int NOT NULL DEFAULT '0' COMMENT '外键，角色id',
  `user_company_id` int NOT NULL DEFAULT '0' COMMENT '所属公司id',
  `user_qq` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户QQ',
  `user_wechat` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户微信',
  `user_status` tinyint NOT NULL DEFAULT '0' COMMENT '账号状态(默认0不可用，需审核；1-可用，2-被删除)',
  `user_create_time` datetime(3) DEFAULT NULL COMMENT '创建日期',
  `user_update_time` datetime(3) DEFAULT NULL COMMENT '升级日期',
  `user_description` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '说明',
  `user_login_count` int NOT NULL DEFAULT '0' COMMENT '登录次数',
  `last_login_time` datetime(3) DEFAULT NULL COMMENT '最后一次登录日期',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除日期'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`uid`, `user_real_name`, `user_avatar`, `user_gender`, `passport`, `password`, `user_email`, `user_mobile`, `user_role_id`, `user_company_id`, `user_qq`, `user_wechat`, `user_status`, `user_create_time`, `user_update_time`, `user_description`, `user_login_count`, `last_login_time`, `deleted_at`) VALUES
(1, '管理员', 'https://pic2.zhimg.com/v2-30c5c15feb31bd7b9a4ce9ee75f2ba5e_is.jpg', 0, 'admin', 'e10adc3949ba59abbe56e057f20f883e', '1697470538@qq.com', '13570923334', 11, 1, '', '', 0, '2021-10-14 16:25:09.000', '2021-12-31 16:27:02.000', '测试号', 459, '2022-03-19 16:22:14.000', NULL),
(3, '安娜', '', 1, 'anna', 'qweqweqw123ed', 'anna@gmail.com', '13345678901', 0, 0, '', '', 0, '2021-10-14 18:52:16.000', '2021-10-14 18:52:16.000', '', 0, NULL, NULL),
(4, '麗莎', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 1, 'lisa', 'qweqweqw123ed', 'lisa@gmail.com', '13345678902', 8, 0, '', '', 0, '2021-10-15 17:08:26.000', '2021-10-15 17:08:26.000', '', 0, NULL, NULL),
(5, '麗莎', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 1, 'lisa2', 'qweqweqw123ed', 'lisa2@gmail.com', '13345678912', 3, 0, '', '', 0, '2021-10-16 12:24:02.000', '2021-10-16 12:24:02.000', '', 0, NULL, NULL),
(6, '游客', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 0, 'demo', '123456', '', '13570923403', 4, 0, '', '', 0, '2021-10-14 16:24:00.000', '2021-10-14 16:28:49.000', '', 0, NULL, NULL),
(9, '李', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 0, 'li', '25f9e794323b453885f5181f1b624d0b', 'li@gmail.com', '13345678952', 5, 0, '', '', 0, '2021-10-19 17:10:27.000', '2021-10-19 17:10:27.000', '', 0, NULL, NULL),
(10, '李', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 0, 'li2', '25f9e794323b453885f5181f1b624d0b', 'li2@gmail.com', '13345278912', 1, 0, '', '', 0, '2021-10-19 18:24:06.000', '2021-10-19 18:24:06.000', '', 0, NULL, NULL),
(11, '李四', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 0, 'li4', '25f9e794323b453885f5181f1b624d0b', 'li4@gmail.com', '13345478912', 6, 0, '', '', 0, '2021-10-20 12:13:08.000', '2021-10-20 12:13:08.000', '', 0, NULL, NULL),
(12, '王二', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 0, 'wang2', '25f9e794323b453885f5181f1b624d0b', 'wang2@gmail.com', '13342478912', 10, 0, '', '', 0, '2021-10-20 12:14:44.000', '2021-10-20 12:14:44.000', '', 0, NULL, NULL),
(13, '啦啦啦', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 0, 'lalala', '65a0ec385ca6a0c1e20d1f8270c28303', 'lalala@gmail.com', '13342418912', 1, 0, '', '', 0, '2021-11-08 18:37:23.000', '2021-11-08 18:37:23.000', '', 0, NULL, NULL),
(14, '无名宝', 'https://img.tukuppt.com/ad_preview/00/09/06/hyVORwxNGY.jpg', 0, 'wuming', '4548120abc49f9136f8d4107b8f0b46c', '123456@qq.com', '13345678910', 12, 1, '', '', 0, NULL, NULL, '', 0, NULL, NULL);

-- --------------------------------------------------------

--
-- 表的结构 `website_config`
--

CREATE TABLE `website_config` (
  `id` int UNSIGNED NOT NULL COMMENT 'id主键',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'catli weibsite' COMMENT '标题',
  `keywords` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '关键词',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `domain` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'localhost:8080' COMMENT '域名',
  `beian` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备案号',
  `domain_duetime` datetime DEFAULT NULL COMMENT '域名到期日',
  `ssl_duetime` datetime DEFAULT NULL COMMENT 'SSL证书到期日',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'licat233@gmail.com' COMMENT '管理员邮箱',
  `adminer_id` int NOT NULL DEFAULT '1' COMMENT '指定网站管理员id，外键'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='网站信息表';

--
-- 转存表中的数据 `website_config`
--

INSERT INTO `website_config` (`id`, `title`, `keywords`, `description`, `domain`, `beian`, `domain_duetime`, `ssl_duetime`, `email`, `adminer_id`) VALUES
(1, '享购工作台', '工作台', '专注于办公管理', 'http://localhost:80/', '', NULL, NULL, 'licat233@gmail.com', 1);

--
-- 转储表的索引
--

--
-- 表的索引 `ad_account`
--
ALTER TABLE `ad_account`
  ADD PRIMARY KEY (`aaid`);

--
-- 表的索引 `ad_campaign`
--
ALTER TABLE `ad_campaign`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ad_campaign_cost`
--
ALTER TABLE `ad_campaign_cost`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ad_channel`
--
ALTER TABLE `ad_channel`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ad_comments`
--
ALTER TABLE `ad_comments`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ad_creative`
--
ALTER TABLE `ad_creative`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ad_group`
--
ALTER TABLE `ad_group`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ad_project_channel`
--
ALTER TABLE `ad_project_channel`
  ADD PRIMARY KEY (`id`),
  ADD KEY `project_id` (`project_id`),
  ADD KEY `channel_id` (`channel_id`);

--
-- 表的索引 `article`
--
ALTER TABLE `article`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `article_grade`
--
ALTER TABLE `article_grade`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `article_status`
--
ALTER TABLE `article_status`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `brand`
--
ALTER TABLE `brand`
  ADD PRIMARY KEY (`bid`),
  ADD KEY `official_account_id_idx` (`brand_official_account_id`);

--
-- 表的索引 `casbin_rule`
--
ALTER TABLE `casbin_rule`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`);

--
-- 表的索引 `company`
--
ALTER TABLE `company`
  ADD PRIMARY KEY (`cid`);

--
-- 表的索引 `cost_list`
--
ALTER TABLE `cost_list`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `fans_list`
--
ALTER TABLE `fans_list`
  ADD PRIMARY KEY (`fid`);

--
-- 表的索引 `feedback`
--
ALTER TABLE `feedback`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `hijack_account`
--
ALTER TABLE `hijack_account`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `hijack_config`
--
ALTER TABLE `hijack_config`
  ADD PRIMARY KEY (`id`),
  ADD KEY `adminer_id_idx` (`adminer_id`);

--
-- 表的索引 `hijack_url`
--
ALTER TABLE `hijack_url`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ip_list`
--
ALTER TABLE `ip_list`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `ip_type`
--
ALTER TABLE `ip_type`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `jwt_blacklists`
--
ALTER TABLE `jwt_blacklists`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uid` (`uid`);

--
-- 表的索引 `jwt_whitelists`
--
ALTER TABLE `jwt_whitelists`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uid` (`uid`);

--
-- 表的索引 `landing_page`
--
ALTER TABLE `landing_page`
  ADD PRIMARY KEY (`id`),
  ADD KEY `account_id_idx` (`account_id`),
  ADD KEY `article_id_idx` (`article_id`);

--
-- 表的索引 `menu`
--
ALTER TABLE `menu`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `news`
--
ALTER TABLE `news`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `project`
--
ALTER TABLE `project`
  ADD PRIMARY KEY (`pid`);

--
-- 表的索引 `service_account`
--
ALTER TABLE `service_account`
  ADD PRIMARY KEY (`said`),
  ADD KEY `sa_type_id` (`sa_type_id`);

--
-- 表的索引 `service_account_type`
--
ALTER TABLE `service_account_type`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `sys_log`
--
ALTER TABLE `sys_log`
  ADD PRIMARY KEY (`id`),
  ADD KEY `create_time` (`create_time`);

--
-- 表的索引 `sys_role`
--
ALTER TABLE `sys_role`
  ADD PRIMARY KEY (`rid`);

--
-- 表的索引 `teacher`
--
ALTER TABLE `teacher`
  ADD PRIMARY KEY (`tid`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`uid`);

--
-- 表的索引 `website_config`
--
ALTER TABLE `website_config`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `ad_account`
--
ALTER TABLE `ad_account`
  MODIFY `aaid` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '账户id', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `ad_campaign`
--
ALTER TABLE `ad_campaign`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '活动id';

--
-- 使用表AUTO_INCREMENT `ad_channel`
--
ALTER TABLE `ad_channel`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '渠道id，主键', AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `ad_comments`
--
ALTER TABLE `ad_comments`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `ad_creative`
--
ALTER TABLE `ad_creative`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `ad_group`
--
ALTER TABLE `ad_group`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `ad_project_channel`
--
ALTER TABLE `ad_project_channel`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id主键';

--
-- 使用表AUTO_INCREMENT `article`
--
ALTER TABLE `article`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id主键';

--
-- 使用表AUTO_INCREMENT `article_grade`
--
ALTER TABLE `article_grade`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id外键';

--
-- 使用表AUTO_INCREMENT `article_status`
--
ALTER TABLE `article_status`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id主键';

--
-- 使用表AUTO_INCREMENT `brand`
--
ALTER TABLE `brand`
  MODIFY `bid` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `casbin_rule`
--
ALTER TABLE `casbin_rule`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `company`
--
ALTER TABLE `company`
  MODIFY `cid` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '公司id，主键', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `cost_list`
--
ALTER TABLE `cost_list`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id主键';

--
-- 使用表AUTO_INCREMENT `fans_list`
--
ALTER TABLE `fans_list`
  MODIFY `fid` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id主键';

--
-- 使用表AUTO_INCREMENT `feedback`
--
ALTER TABLE `feedback`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id主键';

--
-- 使用表AUTO_INCREMENT `hijack_url`
--
ALTER TABLE `hijack_url`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id主键';

--
-- 使用表AUTO_INCREMENT `ip_list`
--
ALTER TABLE `ip_list`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id主键';

--
-- 使用表AUTO_INCREMENT `ip_type`
--
ALTER TABLE `ip_type`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id主键';

--
-- 使用表AUTO_INCREMENT `jwt_blacklists`
--
ALTER TABLE `jwt_blacklists`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `jwt_whitelists`
--
ALTER TABLE `jwt_whitelists`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `landing_page`
--
ALTER TABLE `landing_page`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `menu`
--
ALTER TABLE `menu`
  MODIFY `id` int NOT NULL AUTO_INCREMENT COMMENT '主键';

--
-- 使用表AUTO_INCREMENT `news`
--
ALTER TABLE `news`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id主键';

--
-- 使用表AUTO_INCREMENT `project`
--
ALTER TABLE `project`
  MODIFY `pid` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `service_account`
--
ALTER TABLE `service_account`
  MODIFY `said` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Id主键';

--
-- 使用表AUTO_INCREMENT `service_account_type`
--
ALTER TABLE `service_account_type`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `sys_log`
--
ALTER TABLE `sys_log`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `sys_role`
--
ALTER TABLE `sys_role`
  MODIFY `rid` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id', AUTO_INCREMENT=17;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `uid` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'UID主键', AUTO_INCREMENT=15;

--
-- 使用表AUTO_INCREMENT `website_config`
--
ALTER TABLE `website_config`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id主键', AUTO_INCREMENT=2;

--
-- 限制导出的表
--

--
-- 限制表 `ad_project_channel`
--
ALTER TABLE `ad_project_channel`
  ADD CONSTRAINT `ad_project_channel_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `project` (`pid`),
  ADD CONSTRAINT `ad_project_channel_ibfk_2` FOREIGN KEY (`channel_id`) REFERENCES `ad_channel` (`id`);

--
-- 限制表 `hijack_config`
--
ALTER TABLE `hijack_config`
  ADD CONSTRAINT `adminer_idhijack` FOREIGN KEY (`adminer_id`) REFERENCES `user` (`uid`);

--
-- 限制表 `landing_page`
--
ALTER TABLE `landing_page`
  ADD CONSTRAINT `account_id` FOREIGN KEY (`account_id`) REFERENCES `ad_account` (`aaid`),
  ADD CONSTRAINT `article_id` FOREIGN KEY (`article_id`) REFERENCES `article` (`id`) ON DELETE SET NULL;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
