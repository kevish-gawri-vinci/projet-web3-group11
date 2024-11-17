import React, { useState } from "react";
import axios from "axios";
import "./AddArticle.css";

const AddArticle = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [image, setImage] = useState("");
  const [message, setMessage] = useState(null);

  const addArticle = async (e) => {
    e.preventDefault();
    try {
      const newArticle = { title, content, author };
      await axios.post("http://localhost:8080/articles", newArticle);
      setMessage("Article added successfully");
    } catch (err) {
      setMessage("Failed to add article");
    }
  };

  return (
    <div>
      <form onSubmit={addArticle}>
        <input
          type="text"
          placeholder="Title"
          onChange={(e) => setTitle(e.target.value)}
        />
        <input
          type="text"
          placeholder="Content"
          onChange={(e) => setContent(e.target.value)}
        />
        <input
          type="file"
          accept="image/*"
          onChange={(e) => setImage(e.target.files[0])}
        />
        {image && (
          <div>
            <p>Image Preview:</p>
            <img
              src={URL.createObjectURL(image)}
              alt="Preview"
              style={{ maxWidth: "100%", height: "auto" }}
            />
          </div>
        )}
        <button type="submit">Add Article</button>
      </form>
      <h3>{message}</h3>
    </div>
  );
};
export default AddArticle;
