## DB, Table 作成

```bash:bash
$ mysql -u root
```

```mysql:mysql
create database tododb default charset utf8;
create user 'todo'@'localhost' identified by 'todo';
// GRANT権限
grant all on `tododb`.* to 'todo'@'localhost';

USE tododb
DROP TABLE IF EXISTS task;
CREATE TABLE IF NOT EXISTS task (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(255) NOT NULL,
    PRIMARY KEY(id)
);

// todo@localhostのパスワード設定
set password for todo@localhost="todoのパスワード";
```

## 実行

```bash
$ MYSQL_TODO_PASSWORD="todoのパスワード" go run main.go &
```
