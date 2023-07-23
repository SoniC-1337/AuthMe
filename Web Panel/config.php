<?php
// Database Configuration
define('DB_HOST', getenv('DB_HOST'));
define('DB_USERNAME', getenv('DB_USERNAME'));
define('DB_PASSWORD', getenv('DB_PASSWORD'));
define('DB_DATABASE', getenv('DB_DATABASE'));

// Site URL
define('BASE_URL', 'https://example.com/authme-panel');

// Other Configurations
define('SITE_NAME', 'AuthMe Panel');
define('SESSION_NAME', 'authme_session');
define('SESSION_TIMEOUT', 3600); // Session timeout in seconds (e.g., 3600 seconds = 1 hour)

// Error Reporting (comment out or adjust as needed for production)
error_reporting(E_ALL);
ini_set('display_errors', 1);