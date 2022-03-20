-- name: AddFriend :execresult
INSERT INTO friend (owner, friend_id) VALUES (?,?);

-- name: GetOnesFriends :many
SELECT * 
FROM friend INNER JOIN users ON friend.friend_id = users.id 
WHERE owner = ?
LIMIT ?
OFFSET ?;