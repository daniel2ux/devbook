INSERT INTO
    users (name, nick, email, PASSWORD)
VALUES
    (
        "User 1",
        "userone",
        "user1@email.com",
        "$2a$10$4VvzTO/hrykwVFIzZKq5guBGZGde/9V.SOlo4e1SyUD3lKtXjYRmu"
    ),
    (
        "User 2",
        "usertwo",
        "user2@email.com",
        "$2a$10$4VvzTO/hrykwVFIzZKq5guBGZGde/9V.SOlo4e1SyUD3lKtXjYRmu"
    ),
    (
        "User 3",
        "userthree",
        "user3@email.com",
        "$2a$10$4VvzTO/hrykwVFIzZKq5guBGZGde/9V.SOlo4e1SyUD3lKtXjYRmu"
    ),
    (
        "User 4",
        "userfour",
        "user4@email.com",
        "$2a$10$4VvzTO/hrykwVFIzZKq5guBGZGde/9V.SOlo4e1SyUD3lKtXjYRmu"
    ),
    (
        "User 5",
        "userfive",
        "user5@email.com",
        "$2a$10$4VvzTO/hrykwVFIzZKq5guBGZGde/9V.SOlo4e1SyUD3lKtXjYRmu"
    );

-- INSERT INTO
--     followers (user_id, follower_id)
-- VALUES
--     (1, 2),
--     (1, 4),
--     (2, 3),
--     (3, 2);
INSERT INTO
    posts (title, content, author_id)
VALUES
    ("Post User One", "Post by user one", 1),
    ("Post User TWO", "Post by user two", 2),
    ("Post User Three", "Post by user three", 3),
    ("Post User Four", "Post by user four", 4),
    ("Post User Five", "Post by user five", 5)