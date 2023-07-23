<?php

// Retrieve database credentials from config.php
$host = DB_HOST;
$username = DB_USERNAME;
$password = DB_PASSWORD;
$database = DB_DATABASE;

// Create a MySQLi connection
$mysqli = new mysqli($host, $username, $password, $database);

// Check connection
if ($mysqli->connect_error) {
    die("Connection failed: " . $mysqli->connect_error);
}

// Set character set to utf8 (optional, adjust as needed)
$mysqli->set_charset("utf8");
