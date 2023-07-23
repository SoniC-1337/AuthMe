<?php
// Testing
require_once '../dev_config.php';

// Include the required files
//require_once '../config.php';
require_once '../database/database.php';

// Start a PHP session (if not already started)
session_start();

// Check if the user is already logged in, redirect to dashboard
if (isset($_SESSION['user_id'])) {
    header("Location: ../dashboard");
    exit();
}

// Initialize variables
$username = "";
$password = "";
$confirm_password = "";
$register_error = "";

// Process login form submission
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // Get the form data
    $username = $_POST["username"];
    $password = $_POST["password"];
    $confirm_password = $_POST["confirm_password"];

    // Validate form data (you can add more validation if needed)
    if (empty($username) || empty($password) || empty($confirm_password)) {
        $register_error = "Please enter all fields.";
    } else {
        if ($password != $confirm_password) {
            $register_error = "Passwords do not match.";
        } else {
            $username = mysqli_real_escape_string($mysqli, $username);
            $password = mysqli_real_escape_string($mysqli, $password);
            $confirm_password = mysqli_real_escape_string($mysqli, $confirm_password);

            // Query the database to check if the user exists and if they do return an error
            $query = "SELECT id FROM panel WHERE username = '$username' LIMIT 1";
            $result = $mysqli->query($query);
            if ($result && $result->num_rows == 1) {
                $register_error = "Username already exists.";
            } else {
                // Generate a random salt
                $salt = bin2hex(random_bytes(32));

                // Hash the password with the salt
                $hashed_password = password_hash($password . $salt, PASSWORD_BCRYPT);

                // Insert the new user into the database
                $query = "INSERT INTO panel (username, password, salt) VALUES ('$username', '$hashed_password', '$salt')";
                $result = $mysqli->query($query);

                if ($result) {
                    // Success
                    header("Location: ../login");
                    exit();
                } else {
                    $register_error = "An unknown error occurred.";
                }
            }
        }
    }
}
?>

<!DOCTYPE html>
<html>
    <head>
        <title>Register - <?php echo SITE_NAME; ?></title>
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css">
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.min.js">
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js">
    </head>
    <body>
        <div class="container">
            <h1>Register</h1>
            <?php if (!empty($register_error)): ?>
                <div class="alert alert-danger"><?php echo $register_error; ?></div>
            <?php endif; ?>
            <form method="post">
                <div class="form-group">
                    <label for="username">Username</label>
                    <input type="text" name="username" id="username" class="form-control" value="<?php echo $username; ?>" required>
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" name="password" id="password" class="form-control" required>
                </div>
                <div class="form-group">
                    <label for="confirm_password">Confirm Password</label>
                    <input type="password" name="confirm_password" id="confirm_password" class="form-control" required>
                </div>
                <button type="submit" class="btn btn-primary mt-2">Register</button>
                <a href="./login.php" class="btn btn-secondary mt-2">Login</a>
            </form>
        </div>
    </body>
</html>