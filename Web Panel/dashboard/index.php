<?php
// dashboard.php

// Testing
require_once '../dev_config.php';

// Required includes
// require_once 'config.php';
require_once '../includes/functions.php';

// Start a PHP session (if not already started)
session_start();

// Check if the user is logged in, otherwise redirect to login page
if (!isUserLoggedIn()) {
    header("Location: ../login.php");
    exit();
}

// Logout user when the Logout button is clicked
if ($_SERVER["REQUEST_METHOD"] == "POST" && isset($_POST["logout"])) {
    logoutUser();
    header("Location: ../login.php");
    exit();
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>Dashboard - <?php echo SITE_NAME; ?></title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.min.js">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js">
</head>
<body>
<div class="container">
    <h1>Dashboard</h1>
    <p>Welcome, <?php echo $_SESSION['username']; ?>!</p>
    <form method="post">
        <button type="submit" class="btn btn-danger" name="logout">Logout</button>
    </form>
</div>
</body>
</html>
