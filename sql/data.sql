INSERT INTO users (name, nickname, email, password) VALUES
('User 01', 'user01', 'email01@email.com', '$2a$10$nT9IiQ6PrsGs/6zuYnTmz.udR68hPzD62Nr6jWUiIQn1DsfCwrLTe'),
('User 02', 'user02', 'email02@email.com', '$2a$10$nT9IiQ6PrsGs/6zuYnTmz.udR68hPzD62Nr6jWUiIQn1DsfCwrLTe'),
('User 03', 'user03', 'email03@email.com', '$2a$10$nT9IiQ6PrsGs/6zuYnTmz.udR68hPzD62Nr6jWUiIQn1DsfCwrLTe');

INSERT INTO followers (userId, followerId) VALUES
(1, 2),
(1, 3),
(3, 1);

INSERT INTO posts (title, content, authorId, likes) VALUES
('Title 01', 'Content 01', 1, 0),
('Title 02', 'Content 03', 2, 10),
('Title 03', 'Content 03', 3, 20);
