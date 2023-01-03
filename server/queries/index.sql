SELECT post.slug AS Post,title,cdate,image FROM linked INNER JOIN post ON post.slug = linked.post INNER JOIN category ON category.slug = linked.category WHERE category.slug != infaq
