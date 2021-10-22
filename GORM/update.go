// Update with conditions
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', 
// updated_at='2021-10-22 21:34:10' WHERE active=true;

// Update attributes with `struct`, will only update non-zero fields
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', 
// age=18, 
// updated_at = '2021-10-22 21:34:10' WHERE id = 111;

