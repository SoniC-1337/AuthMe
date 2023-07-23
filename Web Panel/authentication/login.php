<?php
// Testing
require_once '../dev_config.php';

// Required includes
// require_once 'config.php';
require_once '../includes/functions.php';

// Start a PHP session (if not already started)
session_start();

// Check if the user is already logged in, redirect to dashboard
if (isUserLoggedIn()) {
    header("Location: ../dashboard");
    exit();
}

if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // Get the form data
    $username = $_POST["username"];
    $password = $_POST["password"];

    // Call the processLogin function
    $login_result = processLogin($username, $password);

    if ($login_result === true) {
        header("Location: ../dashboard");
        exit();
    } else {
        $login_error = $login_result;
    }
}
?>

<!DOCTYPE html>
<html>
<head>
    <title>Login - <?php echo SITE_NAME; ?></title>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.min.js">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js">
</head>
<body>
<div class="container">
    <h1>Login</h1>
    <?php if (!empty($login_error)): ?>
        <div class="alert alert-danger"><?php echo $login_error; ?></div>
    <?php endif; ?>
    <form method="post">
        <div class="form-group">
            <label for="username">Username</label>
            <input type="text" name="username" id="username" class="form-control" required>
        </div>
        <div class="form-group">
            <label for="password">Password</label>
            <input type="password" name="password" id="password" class="form-control" required>
        </div>
        <button type="submit" class="btn btn-primary mt-2">Login</button>
        <a href="./register.php" class="btn btn-secondary mt-2">Register</a>
    </form>
</div>
</body>
</html>
