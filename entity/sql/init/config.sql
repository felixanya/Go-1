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

/*绑定手机奖励*/
INSERT `t_common_config`(`key`, `subkey`, `value`) values (
    'bindphone',
    'reward',
    '{
    "type": 2,
    "num": 5 
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
			"attrLimit":10000,
			"describe":"rose_sm"
		},
		{
			"propID": 2,
			"name": "beer",
			"attrType": 1,
			"attrID":1,
			"attrValue":-100,
			"attrLimit":10000,
			"describe":"beer_sm"
		},
		{
			"propID": 3,
			"name": "bomb",
			"attrType": 1,
			"attrID":1,
			"attrValue":-100,
			"attrLimit":10000,
			"describe":"bomb_sm"
		},
		{
			"propID": 4,
			"name": "grabChicken",
			"attrType": 1,
			"attrID":1,
			"attrValue":-100,
			"attrLimit":10000,
			"describe":"grabChicken_sm"
		},
		{
			"propID": 5,
			"name": "eggGun",
			"attrType": 1,
			"attrID":1,
			"attrValue":-10000,
			"attrLimit":500000,
			"describe":"eggGun_sm"
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
"isAlms":1,
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
"isAlms":1,
"remark":null
},
{ 
"id":3,
"gameID":3,
"levelID":1,
"name":"1级场",
"fee":100,
"baseScores":150,
"lowScores":2000,
"highScores":100000,
"realOnlinePeople":1,
"showOnlinePeople":1,
"status":1,
"tag":null,
"isAlms":1,
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
"isAlms":1,
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
"isAlms":1,
"remark":null
},
{ 
"id":6,
"gameID":2,
"levelID":2,
"name":"2级场",
"fee":600,
"baseScores":600,
"lowScores":8000,
"highScores":1000000,
"realOnlinePeople":1,
"showOnlinePeople":100,
"status":1,
"tag":null,
"isAlms":1,
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
"isAlms":1,
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
"isAlms":1,
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
                "go_type":0,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"1"
            },
            {
                "ad_id":2,
                "ad_tick":5,
                "go_type":0,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"2"
            },
            {
                "ad_id":3,
                "ad_tick":5,
                "go_type":0,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"3"
            },
            {
                "ad_id":4,
                "ad_tick":5,
                "go_type":0,
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
                "go_type":0,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"1"
            },
            {
                "ad_id":2,
                "ad_tick":5,
                "go_type":0,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"2"
            },
            {
                "ad_id":3,
                "ad_tick":5,
                "go_type":0,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"3"
            },
            {
                "ad_id":4,
                "ad_tick":5,
                "go_type":0,
                "pic_url":"http:/www.qq.com/pic123.jpg",
                "go_url":"http:/www.qq.com",
                "ad_param":"4"
            }
        ]
    }
]');


INSERT INTO `t_common_config` (`key`, `subkey`, `value`)
VALUES
  ( 'role', 'init', '[ 
    { 
    "produceId":0, 
    "coins":100000,
    "keyCards":10,
    "ingots":0, 
    "item":"1|1;2|1;3|5"} 
]');


-- /*救济金配置*/
INSERT INTO `t_common_config` (`key`, `subkey`, `value`)
VALUES
  ( 'game', 'alms', '[{"almsCountDown":10,"depositCountDown":10,"getNorm":2000,"getTimes":3,"getNumber":2000,"version":1}]');
