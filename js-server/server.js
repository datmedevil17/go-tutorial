const express = require('express');
const cors = require('cors');

const app = express();
const port = process.env.PORT || 3000;

// Middleware
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

// Routes
app.get('/', (req, res) => {
    res.status(200).json({
        message: "Welcome to Simple JS Server!",
        timestamp: new Date().toISOString()
    });
});

app.get('/api/posts', (req, res) => {
    const posts = [
        { id: 1, title: "First Post", body: "This is the first post", author: "John Doe" },
        { id: 2, title: "Second Post", body: "This is the second post", author: "Jane Smith" },
        { id: 3, title: "Third Post", body: "This is the third post", author: "Bob Johnson" }
    ];
    res.status(200).json({ posts, total: posts.length });
});

app.get('/api/posts/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const posts = [
        { id: 1, title: "First Post", body: "This is the first post", author: "John Doe" },
        { id: 2, title: "Second Post", body: "This is the second post", author: "Jane Smith" },
        { id: 3, title: "Third Post", body: "This is the third post", author: "Bob Johnson" }
    ];
    
    const post = posts.find(p => p.id === id);
    if (post) {
        res.status(200).json(post);
    } else {
        res.status(404).json({ error: "Post not found" });
    }
});

app.post('/api/posts', (req, res) => {
    const { title, body, author } = req.body;
    const newPost = {
        id: Date.now(),
        title: title || "Untitled",
        body: body || "No content",
        author: author || "Anonymous"
    };
    res.status(201).json({
        message: "Post created successfully",
        post: newPost
    });
});

app.post('/api/postform', (req, res) => {
    console.log('Form data received:', req.body);
    
    const { title, body, author, category, tags } = req.body;
    
    // Handle form data
    const formPost = {
        id: Date.now(),
        title: title || "Untitled Form Post",
        body: body || "No content provided",
        author: author || "Anonymous User",
        category: category || "General",
        tags: tags ? tags.split(',').map(tag => tag.trim()) : [],
        createdAt: new Date().toISOString(),
        type: "form-submission"
    };
    
    res.status(201).json({
        message: "Form data processed successfully",
        data: formPost,
        receivedFields: Object.keys(req.body)
    });
});

// 404 handler
app.use('*', (req, res) => {
    res.status(404).json({
        error: "Route not found",
        path: req.originalUrl
    });
});

// Start server
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
    console.log(`Available endpoints:`);
    console.log(`  GET  /               - Welcome message`);
    console.log(`  GET  /api/posts      - Get all posts`);
    console.log(`  GET  /api/posts/:id  - Get post by ID`);
    console.log(`  POST /api/posts      - Create new post (JSON)`);
    console.log(`  POST /api/postform   - Upload form data`);
});

module.exports = app;
