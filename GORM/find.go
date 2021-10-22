// Get the first record ordered by primary key
db.First(&user, "10")
// SELECT * FROM users WHERE id = 10;

// Get one record, no specified order
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// Get last record, ordered by primary key desc
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);

// Get all records
result := db.Find(&users)
// SELECT * FROM users;

