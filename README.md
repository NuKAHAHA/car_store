# Car Store

Car Store is an online store application for selling and buying cars.

## Usage

Build and run `src/cmd/main.go`
`-db-password` must be provided to run!

# Participants
Zhaniya Zhakipova,Amir Khamzin,Nurdaulet Khaimuldin

# Overview
The project "KukaNkuAmr" aims to establish an innovative car store that caters to
the diverse needs and preferences of car enthusiasts. With a focus on providing
exceptional customer service and a wide selection of high-quality vehicles, our
store seeks to redefine the car buying experience.

# Technologies Used
HTML & CSS JavaScript Bootstrap jQuery

# Project features:
## Main Functionality
### HTML Rendering: 
Utilizes goview for HTML rendering with support for Go's native HTML templates.
### Static File Serving: 
Serves static files such as CSS and JavaScript through the /public route.
### Database Connection: 
Establishes a connection to MongoDB using the go.mongodb.org/mongo-driver package.
### Session Management: 
Implements session management using gincontrib/sessions.
## Initialization
###  Configuration: 
Loads project configuration settings using the configuration package.
### API Initialization: Initializes the Gin router and API endpoints.
### MongoDB Connection: Sets up a connection to MongoDB and creates 
necessary repositories.
## Components
### Repositories:
Implements repositories for interacting with MongoDB collections (users, wishlists, cars).
### Services:
Defines services for business logic handling (authentication, wishlist, car).
### Controllers:
Creates controllers for handling HTTP requests and responses.
## Routes
### Authentication Routes:
Handles user authentication and registration through /login and /register endpoints.
### Homepage Routes: 
Serves the homepage and logout functionality for authenticated users.
### Car Routes: 
Manages CRUD operations for cars including listing, adding, editing, and deleting.
### Wishlist Routes:
Facilitates wishlist management for users including adding and removing cars from the wishlist.
## Middleware
### Session Middleware:
Utilizes session middleware for managing user sessions and authentication.
## Design
### Bootstrap: 
Implements responsive design and layout using Bootstrap for enhanced user experience across different devices.
### JavaScript: 
Utilizes JavaScript for dynamic content updates, form validation, and user interaction enhancements.
