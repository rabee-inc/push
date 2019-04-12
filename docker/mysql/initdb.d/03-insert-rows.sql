INSERT INTO `push`.`tokens` (
    id, app_id, user_id, platform, device_id, token, created_at, updated_at
)
VALUES
    ("c27455ee66ade4d32c21027963c98e24", "sample_app_id", "hoge", "ios", "hoge_ios_device", "hoge_ios_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000)),
    ("9af54a9a8a81094a6aedc3a2459c93eb", "sample_app_id", "hoge", "android", "hoge_android_device", "hoge_android_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000)),
    ("263d6311b5943a3eb184ac6501373371", "sample_app_id", "hoge", "web", "hoge_web_device", "hoge_web_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000)),
    ("a66470d9d13834f6f2dd2e8b167522c8", "sample_app_id", "fuga", "ios", "fuga_ios_device", "fuga_ios_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000)),
    ("2b257ed85277d1acd366bea9d15028c7", "test_app_id", "hoge", "ios", "hoge_ios_device", "hoge_ios_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000));

