<?php
require_once '../database/database.php';

// Function to process login form submission
function processLogin($username, $password)
{
    global $mysqli;

    // Validate form data (you can add more validation if needed)
    if (empty($username) || empty($password)) {
        return "Please enter both username and password.";
    } else {
        $username = mysqli_real_escape_string($mysqli, $username);
        $password = mysqli_real_escape_string($mysqli, $password);

        // Query the database to check if the user exists
        $query = "SELECT id, username, password, salt FROM panel WHERE username = '$username' LIMIT 1";
        $result = $mysqli->query($query);

        if ($result && $result->num_rows == 1) {
            $user_data = $result->fetch_assoc();
            $hashed_password = $user_data['password'];
            $salt = $user_data['salt'];

            // Verify the password
            if (password_verify($password . $salt, $hashed_password)) {
                // Password is correct, log the user in
                $_SESSION['user_id'] = $user_data['id'];
                $_SESSION['username'] = $user_data['username'];
                return true;
            } else {
                return "Invalid username or password.";
            }
        } else {
            return "Invalid username or password.";
        }
    }
}

// Function to check if the user is already logged in
function isUserLoggedIn()
{
    return isset($_SESSION['user_id']);
}

// Function to clear the session data and redirect the user to the login page
function logoutUser()
{
    session_unset();
    session_destroy();
    header("Location: ../authentication/login.php");
    exit();
}
?>
