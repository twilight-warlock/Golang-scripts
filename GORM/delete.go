// Email's ID is `10`
db.Delete(&emailObj)
// DELETE from emails where id = 10;

// Delete with additional conditions
db.Where("name = ?", "Devansh").Delete(&emailObj)
// DELETE from emails where id = 10 AND name = "Devansh";


