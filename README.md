SELECT project.id, title, technologies, imagepath, content, enddate, startdate, user_id AS author
FROM project 
LEFT JOIN "user" ON project.user_id = "user".id 
WHERE user_id = user.id 
ORDER BY project.id;
