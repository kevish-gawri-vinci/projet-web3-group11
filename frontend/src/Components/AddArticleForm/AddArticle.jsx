import React, { useState } from "react";
import axios from "axios";
import "./AddArticle.css";

const AddArticle = () => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [imgurl, setImgUrl] = useState("");
  const [price, setPrice] = useState(0);
  const [message, setMessage] = useState(null);

  const addArticle = async (e) => {
    e.preventDefault();
    try {
      console.log(price)
      const priceInFloat = parseFloat(price)
      console.log(priceInFloat)
      const newArticle = { name, description, imgurl, price: priceInFloat };
      await axios.post("http://localhost:8080/article/add",newArticle, {
          headers: {Authorization: localStorage.getItem("token")}
        }
      )
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
          placeholder="Name"
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="text"
          placeholder="Content"
          onChange={(e) => setDescription(e.target.value)}
        />
        <input
          type="text"
          placeholder="URL of the image"
          onChange={(e) => setImgUrl(e.target.value)}
        />
        <input
          type="text"
          placeholder="Price"
          onChange={(e) => setPrice(e.target.value)}
        />
        <button type="submit">Add Article</button>
      </form>
      <h3>{message}</h3>
    </div>
  );
};
export default AddArticle;
