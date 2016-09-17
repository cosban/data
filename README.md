# data
> A simple postgres database driver extension to provide simplified query and transaction execution

Created for the sake of learning and memes

# Usage
Using this is actually really simple, you just need to follow these steps:

### Connect

Connect to the database. There's an exported Connect function just pass in the url string form.

### Prepare

You can either manually roll your own Statement objects, or use the Prepare function.    
The prepare function takes in your query along with any arguments which should be injected into it.    

### Execute

Either execute your transactions, i.e. your inserts, updates, deletes, Or your queries which return data.   
For single statement operations, you may prepare and execute/query in one step using the PrepareAnd... methods.    
You can either query for a single row, or multiple. It's up to you.