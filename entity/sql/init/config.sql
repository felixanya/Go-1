/*
配置数据库初始化脚本
*/

/*商品表*/
DELETE FROM `t_common_config` WHERE `key`="charge" AND `subkey`="item_list";
INSERT `t_common_config`(`key`, `subkey`, `value`) values(
    'charge', 'item_list', 
    '{
    "android": {
        "default": [
            {
                "item_id": 1,
                "name": "3万金豆",
                "tag": "",
                "price": 500,
                "coin": 30000,
                "present_coin": 0
            },
            {
                "item_id": 2,
                "name": "3.6万金豆",
                "tag": "",
                "price": 600,
                "coin": 36000,
                "present_coin": 0
            },
            {
                "item_id": 3,
                "name": "6.2万金豆",
                "tag": "",
                "price": 1000,
                "coin": 60000,
                "present_coin": 2000
            },
            {
                "item_id": 4,
                "name": "7.4万金豆",
                "tag": "",
                "price": 1200,
                "coin": 72000,
                "present_coin": 2000
            },
            {
                "item_id": 5,
                "name": "18.8万金豆",
                "tag": "",
                "price": 3000,
                "coin": 180000,
                "present_coin": 8000
            },
            {
                "item_id": 6,
                "name": "64万金豆",
                "tag": "",
                "price": 9800,
                "coin": 588000,
                "present_coin": 52000
            },
            {
                "item_id": 7,
                "name": "125万金豆",
                "tag": "",
                "price": 18800,
                "coin": 1128000,
                "present_coin": 122000
            },
            {
                "item_id": 8,
                "name": "320万金豆",
                "tag": "",
                "price": 44800,
                "coin": 2688000,
                "present_coin": 512000
            }
        ]
    },
    "iphone": {
        "default": [
            {
                "item_id": 1,
                "name": "3万金豆",
                "tag": "",
                "price": 500,
                "coin": 30000,
                "present_coin": 0
            },
            {
                "item_id": 2,
                "name": "3.6万金豆",
                "tag": "",
                "price": 600,
                "coin": 36000,
                "present_coin": 0
            },
            {
                "item_id": 3,
                "name": "6.2万金豆",
                "tag": "",
                "price": 1000,
                "coin": 60000,
                "present_coin": 2000
            },
            {
                "item_id": 4,
                "name": "7.4万金豆",
                "tag": "",
                "price": 1200,
                "coin": 72000,
                "present_coin": 2000
            },
            {
                "item_id": 5,
                "name": "18.8万金豆",
                "tag": "",
                "price": 3000,
                "coin": 180000,
                "present_coin": 8000
            },
            {
                "item_id": 6,
                "name": "64万金豆",
                "tag": "",
                "price": 9800,
                "coin": 588000,
                "present_coin": 52000
            },
            {
                "item_id": 7,
                "name": "125万金豆",
                "tag": "",
                "price": 18800,
                "coin": 1128000,
                "present_coin": 122000
            },
            {
                "item_id": 8,
                "name": "320万金豆",
                "tag": "",
                "price": 44800,
                "coin": 2688000,
                "present_coin": 512000
            }
        ]
    }
}'
);

/*每日最大充值数*/
INSERT `t_common_config`(`key`, `subkey`, `value`) values (
    'charge',
    'day_max',
    '{
    "max_charge": 200000
    }'
);

INSERT `t_common_config`(`key`, `subkey`, `value`) values ( 
    'prop', 
    'interactive',
    '[
            {
                "propID": 1,
                "name": "rose",
                "attrType": 1,
                "attrID":1,
                "attrValue":-100,
                "attrLimit":10000
            },
            {
                "propID": 2,
                "name": "beer",
                "attrType": 1,
                "attrID":1,
                "attrValue":-100,
                "attrLimit":10000
            },
            {
                "propID": 3,
                "name": "bomb",
                "attrType": 1,
                "attrID":1,
                "attrValue":-100,
                "attrLimit":10000
            },
            {
                "propID": 4,
                "name": "grabChicken",
                "attrType": 1,
                "attrID":1,
                "attrValue":-100,
                "attrLimit":10000
            },
            {
                "propID": 5,
                "name": "eggGun",
                "attrType": 1,
                "attrID":1,
                "attrValue":-10000,
                "attrLimit":500000
            }
    ]
    '
); 


INSERT INTO `t_common_config` (`id`, `key`, `subkey`, `value`)
VALUES
  ('71', 'game', 'config', '[ 
{ 
"id":1,
"gameID":1,
"name":"血流麻将",
"type":1,
"minPeople":4,
"maxPeople":4,
"playform":null,
"countryID":null,
"provinceID":null,
"cityID":null,
"channelID":null
},
{ 
"id":2,
"gameID":2,
"name":"血战麻将",
"type":1,
"minPeople":4,
"maxPeople":4,
"playform":null,
"countryID":null,
"provinceID":null,
"cityID":null,
"channelID":null
},
{ 
"id":3,
"gameID":3,
"name":"斗地主",
"type":2,
"minPeople":3,
"maxPeople":3,
"playform":null,
"countryID":null,
"provinceID":null,
"cityID":null,
"channelID":null
},
{ 
"id":4,
"gameID":4,
"name":"二人麻将",
"type":1,
"minPeople":2,
"maxPeople":2,
"playform":null,
"countryID":null,
"provinceID":null,
"cityID":null,
"channelID":null
}]');



INSERT INTO `t_common_config` (`id`, `key`, `subkey`, `value`)
VALUES
  ('72', 'game', 'levelconfig', '[ 
{ 
"id":1,
"gameID":1,
"levelID":1,
"name":"1级场",
"fee":100,
"baseScores":100,
"lowScores":2000,
"highScores":80000,
"realOnlinePeople":1,
"showOnlinePeople":1,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":2,
"gameID":2,
"levelID":1,
"name":"1级场",
"fee":150,
"baseScores":150,
"lowScores":2000,
"highScores":120000,
"realOnlinePeople":1,
"showOnlinePeople":1,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":3,
"gameID":3,
"levelID":1,
"name":"1级场",
"fee":100,
"baseScores":150,
"lowScores":20000,
"highScores":100000,
"realOnlinePeople":1,
"showOnlinePeople":1,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":4,
"gameID":4,
"levelID":1,
"name":"1级场",
"fee":50,
"baseScores":100,
"lowScores":2000,
"highScores":80000,
"realOnlinePeople":1,
"showOnlinePeople":1,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":5,
"gameID":1,
"levelID":2,
"name":"2级场",
"fee":600,
"baseScores":500,
"lowScores":10000,
"highScores":500000,
"realOnlinePeople":1,
"showOnlinePeople":100,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":6,
"gameID":2,
"levelID":2,
"name":"2级场",
"fee":600,
"baseScores":600,
"lowScores":80000,
"highScores":1000000,
"realOnlinePeople":1,
"showOnlinePeople":100,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":7,
"gameID":3,
"levelID":2,
"name":"2级场",
"fee":350,
"baseScores":300,
"lowScores":3000,
"highScores":200000,
"realOnlinePeople":1,
"showOnlinePeople":100,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":8,
"gameID":4,
"levelID":2,
"name":"2级场",
"fee":500,
"baseScores":300,
"lowScores":10000,
"highScores":400000,
"realOnlinePeople":1,
"showOnlinePeople":100,
"status":1,
"tag":null,
"remark":null
},
{ 
"id":9,
"gameID":1,
"levelID":3,
"name":"3级场",
"fee":2500,
"baseScores":2000,
"lowScores":40000,
"highScores":-1,
"realOnlinePeople":1,
"showOnlinePeople":200,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
},
{ 
"id":10,
"gameID":2,
"levelID":3,
"name":"3级场",
"fee":2500,
"baseScores":2500,
"lowScores":30000,
"highScores":-1,
"realOnlinePeople":1,
"showOnlinePeople":200,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
},
{ 
"id":11,
"gameID":3,
"levelID":3,
"name":"3级场",
"fee":1200,
"baseScores":1200,
"lowScores":12000,
"highScores":-1,
"realOnlinePeople":1,
"showOnlinePeople":200,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
},
{ 
"id":12,
"gameID":4,
"levelID":3,
"name":"3级场",
"fee":3000,
"baseScores":1500,
"lowScores":50000,
"highScores":-1,
"realOnlinePeople":1,
"showOnlinePeople":200,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
},
{ 
"id":13,
"gameID":1,
"levelID":4,
"name":"4级场",
"fee":10000,
"baseScores":8000,
"lowScores":160000,
"highScores":-1,
"realOnlinePeople":1,
"showOnlinePeople":500,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
},
{ 
"id":14,
"gameID":2,
"levelID":4,
"name":"4级场",
"fee":8000,
"baseScores":8000,
"lowScores":100000,
"highScores":-1,
"realOnlinePeople":1,
"showOnlinePeople":500,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
},
{ 
"id":15,
"gameID":3,
"levelID":4,
"name":"4级场",
"fee":5000,
"baseScores":5000,
"lowScores":50000,
"highScores":-1,
"realOnlinePeople":1,
"showOnlinePeople":500,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
},
{ 
"id":16,
"gameID":4,
"levelID":4,
"name":"4级场",
"fee":20000,
"baseScores":8000,
"lowScores":100000,
"highScores":-1, 
"realOnlinePeople":1,
"showOnlinePeople":500,
"status":1,
"tag":null,
"isAlms":1,
"remark":null
}]');


INSERT INTO `t_common_config` (`key`, `subkey`, `value`)
VALUES
  ( 'horse', 'config', '[
    {
        "id": 1,
        "prov": 1,
        "city": 0,
        "channel": 0,
        "isOpen": 1,
        "isUseParent": 1,
        "tickTime": 5,
        "sleepTime": 100,
        "horse": [
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯1"
            },
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯2"
            },
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯3"
            },
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯4"
            },
            {
                "isOpen": 1,
                "playType": 2,
                "weekDate": [],
                "beginDate": "2018-07-30",
                "endDate": "2018-09-15",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "指定时间播放"
            }
        ],
        "lastUpdateTime": "2018-08-07 12:08:22"
    },
    {
        "id": 2,
        "prov": 2,
        "city": 0,
        "channel": 0,
        "isOpen": 1,
        "isUseParent": 1,
        "tickTime": 5,
        "sleepTime": 105,
        "horse": [
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯1"
            },
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯2"
            },
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯3"
            },
            {
                "isOpen": 1,
                "playType": 1,
                "weekDate": [
                    1,
                    2,
                    3,
                    4,
                    5,
                    6
                ],
                "beginDate": "",
                "endDate": "",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "循环播放:跑马灯4"
            },
            {
                "isOpen": 1,
                "playType": 2,
                "weekDate": [],
                "beginDate": "2018-07-30",
                "endDate": "2018-09-15",
                "beginTime": "02:00",
                "endTime": "22:00",
                "content": "指定时间播放"
            }
        ],
        "lastUpdateTime": "2018-08-07 12:08:22"
    }
]');



INSERT INTO `t_common_config` (`key`, `subkey`, `value`)
VALUES
  ( 'ad', 'config', '[
    {
        "id":1,
        "prov":0,
        "city":0,
        "channel":0,
        "is_use":1,
        "ad_list":[
            {
                "ad_id":1,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"1"
            },
            {
                "ad_id":2,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"2"
            },
            {
                "ad_id":3,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"3"
            },
            {
                "ad_id":4,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"4"
            }
        ]
    },
    {
        "id":2,
        "prov":0,
        "city":0,
        "channel":1,
        "is_use":1,
        "ad_list":[
            {
                "ad_id":1,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"1"
            },
            {
                "ad_id":2,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"2"
            },
            {
                "ad_id":3,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"3"
            },
            {
                "ad_id":4,
                "ad_tick":5,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"4"
            }
        ]
    }
]');



/*游戏配置*/ 
INSERT INTO `t_game_config` VALUES (1, 1, '血流麻将', 1, 4, 4, NULL, NULL, NULL, NULL, NULL, '2018-08-07 19:01:33', NULL, NULL, NULL);
INSERT INTO `t_game_config` VALUES (2, 2, '血战麻将', 1, 4, 4, NULL, NULL, NULL, NULL, NULL, '2018-08-07 19:03:29', NULL, NULL, NULL);
INSERT INTO `t_game_config` VALUES (3, 3, '斗地主', 2, 3, 3, NULL, NULL, NULL, NULL, NULL, '2018-08-07 20:36:58', NULL, NULL, NULL);
INSERT INTO `t_game_config` VALUES (4, 4, '二人麻将', 1, 2, 2, NULL, NULL, NULL, NULL, NULL, '2018-08-07 20:37:11', NULL, NULL, NULL);

-- /*游戏场次配置*/
INSERT INTO `t_game_level_config` VALUES (1, 1, 1, '新手场', 1, 1, 0, 1000000, 1, 1, 1, NULL, 1, NULL, '2018-08-08 18:17:31', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (2, 2, 1, '新手场', 1, 1, 0, 1000000, 1, 1, 1, NULL, 1, NULL, '2018-08-08 18:17:31', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (3, 3, 1, '新手场', 1, 1, 0, 1000000, 1, 1, 1, NULL, 1, NULL, '2018-08-08 18:17:31', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (4, 4, 1, '新手场', 1, 2, 0, 1000000, 1, 1, 1, NULL, 1, NULL, '2018-08-08 18:17:31', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (5, 1, 2, '中级场', 1, 5, 200, 1000000, 1, 100, 1, NULL, 1, NULL, '2018-08-10 10:40:50', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (6, 2, 2, '中级场', 1, 5, 200, 1000000, 1, 100, 1, NULL, 1, NULL, '2018-08-10 10:40:52', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (7, 3, 2, '中级场', 1, 5, 200, 1000000, 1, 100, 1, NULL, 1, NULL, '2018-08-10 10:41:35', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (8, 4, 2, '中级场', 1, 5, 200, 1000000, 1, 100, 1, NULL, 1, NULL, '2018-08-10 10:43:08', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (9, 1, 3, '大师场', 1, 10, 800, 1000000, 1, 200, 1, NULL, 1, NULL, '2018-08-10 10:43:11', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (10, 2, 3, '大师场', 1, 10, 800, 1000000, 1, 200, 1, NULL, 1, NULL, '2018-08-10 10:43:42', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (11, 3, 3, '大师场', 1, 10, 800, 1000000, 1, 200, 1, NULL, 1, NULL, '2018-08-10 10:45:00', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (12, 4, 3, '大师场', 1, 10, 800, 1000000, 1, 200, 1, NULL, 1, NULL, '2018-08-10 10:45:02', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (14, 2, 4, '土豪场', 1, 100, 100000, 10000000, 1, 500, 1, NULL, 1, NULL, '2018-08-10 10:48:20', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (15, 3, 4, '土豪场', 1, 100, 100000, 10000000, 1, 500, 1, NULL, 1, NULL, '2018-08-10 10:48:50', NULL, NULL, NULL);
INSERT INTO `t_game_level_config` VALUES (16, 4, 4, '土豪场', 1, 100, 100000, 10000000, 1, 500, 1, NULL, 1, NULL, '2018-08-10 10:49:31', NULL, NULL, NULL);