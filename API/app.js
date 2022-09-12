const express = require('express');
const mysql = require('mysql');

const bodyparser = require('body-parser');

const app = express();
const cors = require('cors');

app.use(bodyparser.json());
app.use(cors());

const connection = mysql.createConnection({
    host: '34.133.39.78',
    user: 'root',
    password: 'root',
    database: 'sopesDB'
});


app.get('/getRam', (req, res) => {
    const sql = 'SELECT valor FROM ram ORDER BY id DESC LIMIT 1';
    connection.query(sql, (error, results) => {
        if (error) throw error;
        if (results.length > 0) {
            res.json(results);
        } else {
            res.send('Not result');
        }
    });
});

app.get('/getCpu', (req, res) => {
    const sql = 'SELECT valor FROM cpu ORDER BY id DESC LIMIT 1';
    connection.query(sql, (error, results) => {
        if (error) throw error;
        if (results.length > 0) {
            res.json(results);
            console.log
        } else {
            res.send('Not result');
        }
    });
});


connection.connect(function(error){
    if(error) throw error;
    else console.log('Database server running!');
}); 

app.listen(5000, () => {
    console.log('Server running on port 5000');
});


