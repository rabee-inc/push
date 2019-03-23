INSERT INTO `push`.`tokens` (
    id, user_id, platform, device_id, token, created_at, updated_at
)
VALUES
    ("eca801a5c449f98de675c0c9d7f1efc6", "hoge", "ios", "hoge_ios_device", "hoge_ios_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000)),
    ("6e5b2fdee49730a903621eab20df3b41", "hoge", "android", "hoge_android_device", "hoge_android_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000)),
    ("edc33b7ebf3b5b3236e204a3fb61bad6", "hoge", "web", "hoge_web_device", "hoge_web_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000)),
    ("947ea7cd8f299e7e1ff6aa3f58ad842d", "fuga", "ios", "fuga_ios_device", "fuga_ios_token", ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000), ROUND(UNIX_TIMESTAMP(CURTIME(4)) * 1000));

