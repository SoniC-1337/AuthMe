<?php

// Logout the user and clear the session
session_start();
session_unset();
session_destroy();
header("Location: ../authentication/login");
exit();