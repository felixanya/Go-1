## 通用配置项
-----------------------------------------------------------------------------------
t_common_config
key : ad
subkey : config

## json格式说明
-----------------------------------------------------------------------------------
[
    {
        "id":1,             -- 广告ID
        "prov":0,           -- 省ID
        "city":0,           -- 城市ID
        "channel":0,        -- 渠道ID
        "is_use":1,                -- 是否启用
        "ad_list":[         -- 广告列表
            {
                "ad_id":1,                                  -- 位置ID
                "ad_tick":5,                                -- 轮播时间（单位:秒)
                "pic_url":"http:/www.qq.com/pic123.jpg",    -- 广告图片URL
                "go_url":"http:/www.qq.com",                -- 跳转URL
                "ad_param":"1"                              -- 广告参数
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
 
]


[{
"id":1,
"gameID":1,         -- 游戏ID
"levelID":1,        -- 等级
"name":"新手场",     -- 场次名称
"fee":1,            -- 费用
"baseScores":1,     -- 基础分数
"lowScores":1,      -- 最低分数
"highScores":1000000,   -- 最高分数
"realOnlinePeople":1,   -- 实时在线人数
"showOnlinePeople":1,   -- 显示在线人数
"status":1,             -- 状态
"tag":null,             -- 标签：1.热门；2.New
"remark":null           -- 备注
}]