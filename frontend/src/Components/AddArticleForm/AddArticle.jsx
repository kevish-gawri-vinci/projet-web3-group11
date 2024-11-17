import React, { useState } from 'react';

const AddArticle = () => {
    const [title, setTitle] = useState('');
    const [content, setContent] = useState('');
    const [author, setAuthor] = useState('');
    const [message, setMessage] = useState('');
    
    const addArticle = async (e) => {
        e.preventDefault();
        try {
        const newArticle = { title, content, author };
        await axios.post('http://localhost:8080/articles', newArticle);
        setMessage('Article added successfully');
        } catch (err) {
        setMessage('Failed to add article');
        }
    };
    
    return (
        <div>
        <h2>Add Article</h2>
        <form onSubmit={addArticle}>
            <input
            type='text'
            placeholder='Title'
            onChange={(e) => setTitle(e.target.value)}
            />
            <input
            type='text'
            placeholder='Content'
            onChange={(e) => setContent(e.target.value)}
            />
            <input
            type='text'
            placeholder='Author'
            onChange={(e) => setAuthor(e.target.value)}
            />
            <button type='submit'>Add Article</button>
        </form>
        <h3>{message}</h3>
        </div>
    );
}
export default AddArticle;