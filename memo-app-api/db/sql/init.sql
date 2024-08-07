USE memo_app;

CREATE TABLE memo_list(
	id INT(11) AUTO_INCREMENT NOT NULL, 
  name VARCHAR(30) NOT NULL,
  tag BOOLEAN NOT NULL,
  task_order JSON,
  PRIMARY KEY (id)
);

INSERT INTO memo_list(name, tag, task_order) VALUES
  ('Todoリスト', true, '{"order": [3,2,1]}'),
  ('買い物メモ', false, '{"order": [4,6,5]}'),
  ('行きたい居酒屋', false, '{"order": [8,7]}');

CREATE TABLE task_list(
	id INT(11) AUTO_INCREMENT NOT NULL, 
  name VARCHAR(100) NOT NULL,
  memo_id INT(11) NOT NULL,
  complete BOOLEAN NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO task_list(name, memo_id, complete) VALUES
  ('タスク０', 1, false),
  ('タスク１', 1, false),
  ('タスク２', 1, false),
  ('タスク３', 2, true),
  ('タスク４', 2, false),
  ('タスク５', 2, true),
  ('タスク６', 3, false),
  ('タスク７', 3, true);

CREATE TABLE client_data(
  tab INT(3) NOT NULL
);

INSERT INTO client_data(tab) VALUES
  (0);

CREATE TABLE memo_order(
  `order` JSON
);

INSERT INTO memo_order(`order`) VALUES ('{"order": [3,2,1]}');
