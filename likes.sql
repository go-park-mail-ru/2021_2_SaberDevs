drop table if exists article_likes;

create table if not exists article_likes ( 
Id           serial primary key not null,
Login        varchar(45) references author(Login) on delete cascade,
articleId    int  references articles(Id) on delete cascade,
signum       int 
);