SET SQL_SAFE_UPDATES = 0;
SET sql_mode = '';

update articles a
set a.Enabled = true
Where a.ArticleId in
(select articleID from  unique_articles_view
where ((file regexp 'soa |service oriented |webservice') AND (cited_by >= 10))
            OR ((file regexp 'soa |service oriented |webservice') and (year >= 2015))
            OR file like '%microservice%');



select * from articles_view;